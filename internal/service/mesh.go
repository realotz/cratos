package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"

	pb "github.com/realotz/cratos/api/v1"
)

type MeshService struct {
	pb.UnimplementedMeshServer
	uc  *biz.IstioUsecase
	log *log.Helper
}

func NewMeshService(uc *biz.IstioUsecase, logger log.Logger) pb.MeshServer {
	return &MeshService{uc: uc, log: log.NewHelper("service/mesh", logger)}
}

func (s *MeshService) CheckInfo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	s.uc.GetGatewayList(ctx)
	return &pb.Response{}, nil
}
func (s *MeshService) GetGatewayList(ctx context.Context, req *pb.GetGatewayOption) (*pb.GatewayList, error) {
	return &pb.GatewayList{}, nil
}
func (s *MeshService) CreateGateway(ctx context.Context, req *pb.Gateway) (*pb.Response, error) {
	return &pb.Response{}, nil
}
