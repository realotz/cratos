package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/cratos/api/v1/gateway"
	"github.com/realotz/cratos/internal/data/resources"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1alpha3client "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IstioRepo interface {
	GetGatewayV1alpha3(ns string) v1alpha3client.GatewayInterface
	ListResources(ctx context.Context, option ListOption) (*resources.KubeResourceList, error)
}

type ListOption struct {
	Name      string
	Namespace string
	Kind      string
	Sort      string
	Limit     int
	Offset    int
}

var KindGateway = v1.TypeMeta{
	Kind:       "Gateway",
	APIVersion: "networking.istio.io/v1alpha3",
}

type IstioUsecase struct {
	repo IstioRepo
	log  *log.Helper
}

func NewIstioUsecase(repo IstioRepo, logger log.Logger) *IstioUsecase {
	return &IstioUsecase{repo: repo, log: log.NewHelper("usecase/istio", logger)}
}

func (uc *IstioUsecase) GetGatewayList(ctx context.Context, req *pb.ListOption) (*pb.GatewayList, error) {
	res, err := uc.repo.ListResources(ctx, ListOption{
		Name:      req.Name,
		Namespace: req.Namespace,
		Kind:      "Gateway",
		Sort:      req.Sort,
		Limit:     int(req.Limit),
		Offset:    int(req.Offset),
	})
	if err != nil {
		return nil, err
	}
	list := &pb.GatewayList{
		List:  nil,
		Total: res.Total,
	}
	err = res.DeepCopyList(&list.List)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (uc *IstioUsecase) CreateGateway(ctx context.Context, req *pb.Gateway) (*v1alpha3.Gateway, error) {
	return uc.repo.GetGatewayV1alpha3(req.Metadata.GetNamespace()).Create(ctx, &v1alpha3.Gateway{
		TypeMeta: KindGateway,
		ObjectMeta: v1.ObjectMeta{
			Name:      req.Metadata.GetName(),
			Namespace: req.Metadata.GetNamespace(),
			Labels:    req.Metadata.GetLabels(),
		},
		Spec: *req.Spec,
	}, v1.CreateOptions{})
}

func (uc *IstioUsecase) UpdateGateway(ctx context.Context, req *pb.Gateway) (*v1alpha3.Gateway, error) {
	return uc.repo.GetGatewayV1alpha3(req.Metadata.GetNamespace()).Update(ctx, &v1alpha3.Gateway{
		TypeMeta: KindGateway,
		ObjectMeta: v1.ObjectMeta{
			Name:            req.Metadata.GetName(),
			Namespace:       req.Metadata.GetNamespace(),
			Labels:          req.Metadata.GetLabels(),
			ResourceVersion: req.Metadata.GetResourceVersion(),
		},
		Spec: *req.Spec,
	}, v1.UpdateOptions{})
}

func (uc *IstioUsecase) DelGateway(ctx context.Context, ns string, name string) error {
	return uc.repo.GetGatewayV1alpha3(ns).Delete(ctx, name, v1.DeleteOptions{})
}

func (uc *IstioUsecase) GetGateway(ctx context.Context, ns string, name, version string) (*v1alpha3.Gateway, error) {
	return uc.repo.GetGatewayV1alpha3(ns).Get(ctx, name, v1.GetOptions{
		ResourceVersion: version,
	})
}
