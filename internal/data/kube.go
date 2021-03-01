package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type kubeRepo struct {
	data   *Data
	log    *log.Helper
	client *clientset.Clientset
}

func getKubeClient(cfg string) (*clientset.Clientset, error) {
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
	return clientset.NewForConfig(config)
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
