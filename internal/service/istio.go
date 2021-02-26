package service

import pb "github.com/realotz/cratos/api/v1"

type IstioService struct {
	pb.UnimplementedIstioServiceServer
}

func NewIstioService() *IstioService {
	return &IstioService{}
}
