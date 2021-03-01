package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (d *istioRepo) GetGatewayList(ctx context.Context, ns string, op v1.ListOptions) (*v1beta1.GatewayList, error) {
	return d.client.NetworkingV1beta1().Gateways(ns).List(ctx, op)
}
