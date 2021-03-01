package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IstioRepo interface {
	GetGatewayList(ctx context.Context, ns string, op v1.ListOptions) (*v1beta1.GatewayList, error)
}

type KubeRepo interface {
}

type IstioUsecase struct {
	repo IstioRepo
	log  *log.Helper
}

func NewIstioUsecase(repo IstioRepo, logger log.Logger) *IstioUsecase {
	return &IstioUsecase{repo: repo, log: log.NewHelper("usecase/istio", logger)}
}

func (uc *IstioUsecase) GetGatewayList(ctx context.Context) {
	list, err := uc.repo.GetGatewayList(ctx, "", v1.ListOptions{
		Watch:               false,
		AllowWatchBookmarks: false,
	})
	fmt.Println(list, err)
}
