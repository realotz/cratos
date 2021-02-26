package data

import (
	"github.com/google/wire"
	"github.com/realotz/cratos/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewIstioRepo, NewKubeRepo)

type Data struct {
	cfg *conf.Data
}

func NewData(cfg *conf.Data) *Data {
	return &Data{cfg: cfg}
}
