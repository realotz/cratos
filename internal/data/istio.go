package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type istioRepo struct {
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

// NewIstioRepo .
func NewIstioRepo(data *Data, logger log.Logger) biz.KubeRepo {
	client, _ := getKubeClient(data.cfg.Istio.Config)
	return &istioRepo{
		data:   data,
		client: client,
		log:    log.NewHelper("data/istio", logger),
	}
}
