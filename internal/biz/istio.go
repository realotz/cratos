package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/cratos/api/v1"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	v1beta1client "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IstioRepo interface {
	GetGatewayV1beta1(ns string) v1beta1client.GatewayInterface
}

const ApiVersion = "networking.istio.io/v1beta1"
const KindGateway = "Gateway"

type IstioUsecase struct {
	repo IstioRepo
	log  *log.Helper
}

func NewIstioUsecase(repo IstioRepo, logger log.Logger) *IstioUsecase {
	return &IstioUsecase{repo: repo, log: log.NewHelper("usecase/istio", logger)}
}

func (uc *IstioUsecase) GetGatewayList(ctx context.Context, req *pb.ListOption) ([]v1beta1.Gateway, error) {
	res, err := uc.repo.GetGatewayV1beta1(req.Namespace).List(ctx, v1.ListOptions{
		Watch:               false,
		AllowWatchBookmarks: false,
		Limit:               req.Limit,
		Continue:            req.Continue,
		LabelSelector:       req.LabelSelector,
		FieldSelector:       req.FieldSelector,
	})
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

func (uc *IstioUsecase) CreateGateway(ctx context.Context, req *pb.Gateway) (*v1beta1.Gateway, error) {
	return uc.repo.GetGatewayV1beta1(req.Metadata.GetNamespace()).Create(ctx, &v1beta1.Gateway{
		TypeMeta: v1.TypeMeta{
			Kind:       KindGateway,
			APIVersion: ApiVersion,
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      req.Metadata.GetName(),
			Namespace: req.Metadata.GetNamespace(),
			Labels:    req.Metadata.GetLabels(),
		},
		Spec: *req.Spec,
	}, v1.CreateOptions{})
}

func (uc *IstioUsecase) UpdateGateway(ctx context.Context, req *pb.Gateway) (*v1beta1.Gateway, error) {
	return uc.repo.GetGatewayV1beta1(req.Metadata.GetNamespace()).Update(ctx, &v1beta1.Gateway{
		TypeMeta: v1.TypeMeta{
			Kind:       KindGateway,
			APIVersion: ApiVersion,
		},
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
	return uc.repo.GetGatewayV1beta1(ns).Delete(ctx, name, v1.DeleteOptions{})
}

func (uc *IstioUsecase) GetGateway(ctx context.Context, ns string, name, version string) (*v1beta1.Gateway, error) {
	return uc.repo.GetGatewayV1beta1(ns).Get(ctx, name, v1.GetOptions{
		ResourceVersion: version,
	})
}
