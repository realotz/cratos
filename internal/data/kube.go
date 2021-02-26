package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"k8s.io/client-go/kubernetes"
)

type kubeRepo struct {
	data   *Data
	log    *log.Helper
	client *kubernetes.Clientset
}

// NewKubeRepo .
func NewKubeRepo(data *Data, logger log.Logger) biz.IstioRepo {
	client, _ := getKubeClient(data.cfg.Istio.Config)
	return &kubeRepo{
		data:   data,
		client: client,
		log:    log.NewHelper("data/kube", logger),
	}
}
