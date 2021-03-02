package deployer

import (
	"context"
	"fmt"
	"github.com/urfave/cli"
	v1 "k8s.io.io/api/apps/v1"
	corev1 "k8s.io.io/api/core/v1"
	metav1 "k8s.io.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strings"
	"time"
)

type Service struct {
	Name    string
	ID      string
	Version string
}

var (
	// 部署超时时间
	progressDeadlineSeconds int32 = 600
	// 部署数量
	replicas                      int32 = 1
	revisionHistoryLimit          int32 = 10
	terminationGracePeriodSeconds int64 = 30
	maxSurge                            = intstr.FromString("25%")
	maxUnavailable                      = intstr.FromString("25%")
)

func Deploy(c *cli.Context) {
	var config *rest.Config
	var err error
	service := Service{}
	cfile := c.String("file")
	name := c.String("name")
	image := c.String("image")
	ns := c.String("namespaces")
	ctx := context.Background()

	//env config
	envVars := []corev1.EnvVar{
		{
			Name:  "APP_ID",
			Value: fmt.Sprintf("%s.%s", name, ns),
		},
	}
	envs := c.StringSlice("env")
	for _, v := range envs {
		e := strings.Split(v, "=")
		if len(e) == 2 {
			envVars = append(envVars, corev1.EnvVar{
				Name:  e[0],
				Value: e[1],
			})
		} else {
			log.Panicf("env set error -e APP=Test")
		}
	}

	// port config
	var containerPorts []corev1.ContainerPort
	ports := c.IntSlice("ports")
	for _, v := range ports {
		containerPorts = append(containerPorts, corev1.ContainerPort{
			Name:          fmt.Sprintf("TCP-%d", int32(v)),
			ContainerPort: int32(v),
			Protocol:      "TCP",
		})
	}

	// pod container config
	container := []corev1.Container{
		{
			Name:            name,
			Image:           image,
			ImagePullPolicy: "IfNotPresent",
			Ports:           containerPorts,
			Env:             envVars,
		},
	}

	// init container
	if cfile != "" {
		if config, err = clientcmd.BuildConfigFromFlags("", cfile); err != nil {
			log.Print(err)
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			log.Print(err)
		}
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Print(err)
	}
	deployment, err := clientSet.AppsV1().Deployments(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		// create deployment
		meta := metav1.ObjectMeta{
			Labels: map[string]string{
				"kratos-service-app":     service.Name,
				"kratos-service-version": service.Version,
			},
			Annotations: map[string]string{
				"kratos-service-metadata":  name,
				"kratos-service-protocols": name,
			},
		}
		deployment = &v1.Deployment{
			ObjectMeta: meta,
			Spec: v1.DeploymentSpec{
				ProgressDeadlineSeconds: &progressDeadlineSeconds,
				Replicas:                &replicas,
				RevisionHistoryLimit:    &revisionHistoryLimit,
				Selector: &metav1.LabelSelector{
					MatchLabels: meta.Labels,
				},
				Strategy: v1.DeploymentStrategy{
					Type: "RollingUpdate",
					RollingUpdate: &v1.RollingUpdateDeployment{
						MaxUnavailable: &maxUnavailable,
						MaxSurge:       &maxSurge,
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: meta,
					Spec: corev1.PodSpec{
						Volumes:                       nil,
						InitContainers:                nil,
						Containers:                    container,
						RestartPolicy:                 "Always",
						TerminationGracePeriodSeconds: &terminationGracePeriodSeconds,
						DNSPolicy:                     "ClusterFirst",
					},
				},
			},
		}
		deployment, err = clientSet.AppsV1().Deployments(ns).Create(ctx, deployment, metav1.CreateOptions{})
		if err != nil {
			log.Panic(err)
		}
	} else {
		outContainers := deployment.Spec.Template.Spec.Containers
		// update deployment
		deployment.Spec.Template.Spec.Containers = container
		rollback := func() error {
			deployment.Spec.Template.Spec.Containers = outContainers
			deployment, err = clientSet.AppsV1().Deployments(ns).Update(ctx, deployment, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
			return nil
		}
		deployment, err = clientSet.AppsV1().Deployments(ns).Update(ctx, deployment, metav1.UpdateOptions{})
		if err != nil {
			log.Panic(err)
		}
		nctx, closeFunc := context.WithTimeout(ctx, time.Second*60)
		for true {
			time.Sleep(1 * time.Second)
			log.Print("wait deployment ready!!!")
			select {
			case <-nctx.Done():
				err = rollback()
				if err != nil {
					log.Print("deployer wait timeout rollback error", err)
				} else {
					log.Print("deployer wait timeout")
				}
			default:
				deployment, err = clientSet.AppsV1().Deployments(ns).Get(ctx, name, metav1.GetOptions{})
				if err != nil {
					continue
				}
				if deployment.Status.ReadyReplicas == deployment.Status.Replicas {
					break
				}
			}
		}
		closeFunc()
	}
}
