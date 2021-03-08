package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	v1 "k8s.io/api/core/v1"
	"net/http"

	pb "github.com/realotz/cratos/api/v1/namespace"
)

type NamespacesService struct {
	pb.UnimplementedNamespacesServer
	uc  *biz.KubeUsecase
	log *log.Helper
}

func NewNamespacesService(uc *biz.KubeUsecase, logger log.Logger) *NamespacesService {
	return &NamespacesService{uc: uc, log: log.NewHelper("service/namespace", logger)}
}

func (s *NamespacesService) ListTags(ctx context.Context, request *pb.ListOption) (*pb.TagsList, error) {
	list, err := s.uc.GetNamespaceTags(ctx, request)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *NamespacesService) List(ctx context.Context, req *pb.ListOption) (*pb.NamespaceList, error) {
	list, err := s.uc.GetNamespaceList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (s *NamespacesService) Get(ctx context.Context, req *pb.GetKind) (*v1.Namespace, error) {
	if req.Name == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if req.Namespace == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	return s.uc.GetNamespace(ctx, req.Name, req.Version)
}
func (s *NamespacesService) Create(ctx context.Context, req *v1.Namespace) (*pb.Response, error) {
	if req.GetName() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if req.GetNamespace() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	_, err := s.uc.CreateNamespace(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{}, nil
}
func (s *NamespacesService) Update(ctx context.Context, req *v1.Namespace) (*pb.Response, error) {
	if req.GetName() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if req.GetNamespace() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	_, err := s.uc.UpdateNamespace(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{}, nil
}
func (s *NamespacesService) Delete(ctx context.Context, req *pb.DeleteKind) (*pb.Response, error) {
	if req.Name == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if req.Namespace == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	err := s.uc.DelNamespace(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Response{}, nil
}
