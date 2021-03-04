package resources

import (
	"errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/watch"
	"reflect"
)

type Event watch.Event

// 转换方法
type ConvertsFunc func(interface{}) (*KubeResources, error)

// 转换器集合
var eventType = make(map[string]ConvertsFunc)

// k8s资产缓存
type KubeResources struct {
	gorm.Model
	Name       string `json:"name,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
	MeteData   datatypes.JSON
	Spec       datatypes.JSON
	Status     datatypes.JSON
}

func (e Event) Converts() (*KubeResources, error) {
	objType := reflect.TypeOf(e.Object)
	if c, ok := eventType[objType.String()]; ok {
		return c(e.Object)
	} else {
		return nil, errors.New("event converts not full")
	}
}
