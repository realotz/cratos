package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/cratos/api/v1/namespace"
	"github.com/realotz/cratos/internal/data/resources"
	kubev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type KubeRepo interface {
	GetNamespaceV1() corev1.NamespaceInterface
	ListResources(ctx context.Context, option ListOption) (*resources.KubeResourceList, error)
}

type KubeUsecase struct {
	repo KubeRepo
	log  *log.Helper
}

func NewKubeUsecase(repo KubeRepo, logger log.Logger) *KubeUsecase {
	return &KubeUsecase{repo: repo, log: log.NewHelper("usecase/kube", logger)}
}

func (uc *KubeUsecase) GetNamespaceList(ctx context.Context, req *pb.ListOption) (*pb.NamespaceList, error) {
	res, err := uc.repo.ListResources(ctx, ListOption{
		Name:   req.Name,
		Kind:   "Namespace",
		Sort:   req.Sort,
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
	})
	if err != nil {
		return nil, err
	}
	list := &pb.NamespaceList{
		List:  nil,
		Total: res.Total,
	}
	err = res.DeepCopyList(&list.List)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (uc *KubeUsecase) CreateNamespace(ctx context.Context, req *kubev1.Namespace) (*kubev1.Namespace, error) {
	return uc.repo.GetNamespaceV1().Create(ctx, &kubev1.Namespace{
		TypeMeta:   KindGateway,
		ObjectMeta: req.ObjectMeta,
		Spec:       req.Spec,
	}, v1.CreateOptions{})
}

func (uc *KubeUsecase) UpdateNamespace(ctx context.Context, req *kubev1.Namespace) (*kubev1.Namespace, error) {
	return uc.repo.GetNamespaceV1().Update(ctx, &kubev1.Namespace{
		TypeMeta:   KindGateway,
		ObjectMeta: req.ObjectMeta,
		Spec:       req.Spec,
	}, v1.UpdateOptions{})
}

func (uc *KubeUsecase) DelNamespace(ctx context.Context, name string) error {
	return uc.repo.GetNamespaceV1().Delete(ctx, name, v1.DeleteOptions{})
}

func (uc *KubeUsecase) GetNamespace(ctx context.Context, name, version string) (*kubev1.Namespace, error) {
	return uc.repo.GetNamespaceV1().Get(ctx, name, v1.GetOptions{
		ResourceVersion: version,
	})
}
