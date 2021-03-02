package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	v1beta1client "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getIstioClient(cfg string) (*versionedclient.Clientset, error) {
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
	return versionedclient.NewForConfig(config)
}

// NewIstioRepo .
func NewIstioRepo(data *Data, logger log.Logger) biz.IstioRepo {
	client, _ := getIstioClient(data.cfg.Istio.Config)
	return &istioRepo{
		data:   data,
		client: client,
		log:    log.NewHelper("data/istio", logger),
	}
}

type istioRepo struct {
	data   *Data
	log    *log.Helper
	client *versionedclient.Clientset
}

func (d *istioRepo) GetGatewayV1beta1(ns string) v1beta1client.GatewayInterface {
	return d.client.NetworkingV1beta1().Gateways(ns)
}
