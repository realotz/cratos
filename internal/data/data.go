package data

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/realotz/cratos/internal/biz"
	"github.com/realotz/cratos/internal/conf"
	"github.com/realotz/cratos/internal/data/resources"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/watch"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewIstioRepo)

type Data struct {
	cfg *conf.Data
	db  *gorm.DB
	log *log.Helper
}

func NewData(cfg *conf.Data, logger log.Logger) (*Data, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&resources.KubeResources{})
	if err != nil {
		return nil, err
	}
	return &Data{cfg: cfg, db: db, log: log.NewHelper("data/base", logger)}, nil
}

// 查询资源
func (d *Data) listResources(params biz.ListOption) (*resources.KubeResourceList, error) {
	kubeList := &resources.KubeResourceList{
		Items: make([]resources.KubeResources, 0),
		Total: 0,
	}
	db := d.db.Model(&resources.KubeResources{})
	if params.Name != "" {
		db = db.Where("name like ?", "%"+params.Name+"%")
	}
	if params.Namespace != "" {
		db = db.Where("namespace = ?", params.Namespace)
	}
	if params.Kind != "" {
		db = db.Where("kind = ?", params.Kind)
	}
	if params.Sort != "" {
		db = db.Order(params.Sort)
	} else {
		db = db.Order("id desc")
	}
	db.Count(&kubeList.Total)
	if params.Limit != 0 {
		db.Limit(params.Limit)
	}
	if params.Offset != 0 {
		db.Offset(params.Limit)
	}
	if err := db.Find(&kubeList.Items).Error; err != nil {
		return nil, err
	}
	return kubeList, nil
}

// k8s资源入库
func (d *Data) watchKubeEvent(ev watch.Event) {
	resource, err := resources.Event(ev).Converts()
	if err != nil {
		d.log.Error(err)
		return
	}
	switch ev.Type {
	// 新增,修改,修改版本
	case watch.Added, watch.Modified, watch.Bookmark:
		var outResources resources.KubeResources
		row := d.db.Where("kind = ? and api_version=?", resource.Kind, resource.APIVersion).
			First(&outResources, datatypes.JSONQuery("mete_data").Equals(resource.Name, "name"),
				datatypes.JSONQuery("mete_data").Equals(resource.Namespace, "namespace"))
		if errors.Is(row.Error, gorm.ErrRecordNotFound) {
			resource.CreatedAt = time.Now()
			d.db.Create(&resource)
		} else if outResources.ID > 0 {
			resource.ID = outResources.ID
			d.db.Updates(&resource)
		} else {
			d.log.Error(row.Error)
		}
	case watch.Deleted:
		err = d.db.Where("kind = ? and api_version=?", resource.Kind, resource.APIVersion).
			Delete(&resource, datatypes.JSONQuery("mete_data").Equals(resource.Name, "name"),
				datatypes.JSONQuery("mete_data").Equals(resource.Namespace, "namespace")).Error
		if err != nil {
			d.log.Error(err)
			return
		}
	default:
		d.log.Error(ev.Type, ev.Object)
	}
}
