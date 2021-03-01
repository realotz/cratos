package data

import (
	"github.com/google/wire"
	"github.com/realotz/cratos/internal/conf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewIstioRepo)

type Data struct {
	cfg *conf.Data
	db  *gorm.DB
}

func NewData(cfg *conf.Data) (*Data, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Data{cfg: cfg, db: db}, nil
}
