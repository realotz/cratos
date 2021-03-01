// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/realotz/cratos/internal/biz"
	"github.com/realotz/cratos/internal/conf"
	"github.com/realotz/cratos/internal/data"
	"github.com/realotz/cratos/internal/server"
	"github.com/realotz/cratos/internal/service"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
