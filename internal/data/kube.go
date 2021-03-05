package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"github.com/realotz/cratos/internal/data/resources"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type kubeRepo struct {
	data   *Data
	log    *log.Helper
	client *kubernetes.Clientset
}

func getKubeClient(cfg string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if cfg != "" {
		if config, err = clientcmd.BuildConfigFromFlags("", cfg); err != nil {
			return nil, err
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			return nil, err
		}
	}
	return kubernetes.NewForConfig(config)
}

// NewKubeRepo .
func NewKubeRepo(data *Data, logger log.Logger) biz.KubeRepo {
	client, _ := getKubeClient(data.cfg.Istio.Config)
	return &kubeRepo{
		data:   data,
		client: client,
		log:    log.NewHelper("data/kube", logger),
	}
}

// 注册Watcher 会启动一个Goroutine 初始检查第一次失败无法启动,后续失败会尝试重连
// 获取watcher event事件后通知数据存储
func (d *kubeRepo) registerWatcher(watcherFunc KubeWatcher) error {
	watcher, err := watcherFunc(d.client)
	if err != nil {
		d.log.Errorf("kube watcher conn error %v", err)
		return err
	}
	go func() {
		for true {
			ch := <-watcher.ResultChan()
			if ch.Object == nil {
				// 重新获取watcher
				watcher, err = watcherFunc(d.client)
				if err != nil {
					d.log.Errorf("kube watcher conn error %v", err)
					time.Sleep(time.Second * 1)
				}
			}
			fmt.Println(ch.Object)
			d.data.watchKubeEvent(ch)
		}
	}()
	return nil
}

func (k *kubeRepo) GetNamespaceV1() corev1.NamespaceInterface {
	return k.client.CoreV1().Namespaces()
}

func (d *kubeRepo) ListResources(ctx context.Context, option biz.ListOption) (*resources.KubeResourceList, error) {
	return d.data.listResources(ctx, option)
}
