package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"github.com/realotz/cratos/internal/data/resources"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	v1alpha3client "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1alpha3"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

// 获取istio rest请求客户端
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
func NewIstioRepo(data *Data, logger log.Logger) (biz.IstioRepo, error) {
	client, _ := getIstioClient(data.cfg.Istio.Config)
	repo := &istioRepo{
		data:   data,
		client: client,
		log:    log.NewHelper("data/istio", logger),
	}
	for _, w := range istioWatchers {
		err := repo.registerWatcher(w)
		if err != nil {
			return nil, err
		}
	}
	return repo, nil
}

func (d *istioRepo) ListResources(ctx context.Context, option biz.ListOption) (*resources.KubeResourceList, error) {
	return d.data.listResources(ctx, option)
}

// 注册Watcher 会启动一个Goroutine 初始检查第一次失败无法启动,后续失败会尝试重连
// 获取watcher event事件后通知数据存储
func (d *istioRepo) registerWatcher(watcherFunc IstioWatcher) error {
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
			d.data.watchKubeEvent(ch)
		}
	}()
	return nil
}

type istioRepo struct {
	data   *Data
	log    *log.Helper
	client *versionedclient.Clientset
}

func (d *istioRepo) GetGatewayV1alpha3(ns string) v1alpha3client.GatewayInterface {
	return d.client.NetworkingV1alpha3().Gateways(ns)
}
