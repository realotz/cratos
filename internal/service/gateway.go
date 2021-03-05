package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/cratos/api/v1/gateway"
	"github.com/realotz/cratos/internal/biz"
	"github.com/realotz/cratos/pkg/utils"
	"net/http"
)

func NewGatewayService(uc *biz.IstioUsecase, logger log.Logger) *GatewayService {
	return &GatewayService{uc: uc, log: log.NewHelper("service/gateway", logger)}
}

type GatewayService struct {
	pb.UnimplementedGatewaysServer
	uc  *biz.IstioUsecase
	log *log.Helper
}

func (s *GatewayService) Get(ctx context.Context, kind *pb.GetKind) (*pb.Gateway, error) {
	if kind.Name == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if kind.Namespace == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	res, err := s.uc.GetGateway(ctx, kind.Namespace, kind.Name, kind.Version)
	if err != nil {
		return nil, err
	}
	gateway := &pb.Gateway{
		ApiVersion: res.APIVersion,
		Kind:       res.Kind,
		Spec:       res.Spec.DeepCopy(),
	}
	_ = utils.StudentToStudentForJson(res.ObjectMeta, &gateway.Metadata)
	return gateway, nil
}

func (s *GatewayService) Delete(ctx context.Context, kind *pb.DeleteKind) (*pb.Response, error) {
	if kind.Name == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if kind.Namespace == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	err := s.uc.DelGateway(ctx, kind.Namespace, kind.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Response{}, nil
}

func (s *GatewayService) Create(ctx context.Context, gateway *pb.Gateway) (*pb.Response, error) {
	if gateway.Metadata.GetName() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if gateway.Metadata.GetNamespace() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	_, err := s.uc.CreateGateway(ctx, gateway)
	if err != nil {
		return nil, err
	}
	return &pb.Response{}, nil
}

func (s *GatewayService) Update(ctx context.Context, gateway *pb.Gateway) (*pb.Response, error) {
	if gateway.Metadata.GetName() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入name")
	}
	if gateway.Metadata.GetNamespace() == "" {
		return nil, errors.Errorf(http.StatusBadRequest, pb.Errors_NotName, "错误参数，未传入namespace")
	}
	_, err := s.uc.UpdateGateway(ctx, gateway)
	if err != nil {
		return nil, err
	}
	return &pb.Response{}, nil
}

func (s *GatewayService) CheckInfo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}

// 获取gateway列表
func (s *GatewayService) List(ctx context.Context, req *pb.ListOption) (*pb.GatewayList, error) {
	list, err := s.uc.GetGatewayList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}
