// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/cratos/internal/biz"
	"github.com/realotz/cratos/internal/conf"
	"github.com/realotz/cratos/internal/data"
	"github.com/realotz/cratos/internal/server"
	"github.com/realotz/cratos/internal/service"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	istioRepo, err := data.NewIstioRepo(dataData, logger)
	if err != nil {
		return nil, err
	}
	istioUsecase := biz.NewIstioUsecase(istioRepo, logger)
	gatewayService := service.NewGatewayService(istioUsecase, logger)
	kubeRepo, err := data.NewKubeRepo(dataData, logger)
	if err != nil {
		return nil, err
	}
	kubeUsecase := biz.NewKubeUsecase(kubeRepo, logger)
	namespacesService := service.NewNamespacesService(kubeUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, gatewayService, namespacesService)
	grpcServer := server.NewGRPCServer(confServer, gatewayService, namespacesService)
	app := newApp(logger, httpServer, grpcServer)
	return app, nil
}
