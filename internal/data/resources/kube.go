package resources

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
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

type KubeResourceList struct {
	Items []KubeResources
	Total int64
}

// 转换
func (r *KubeResourceList) DeepCopyList(impl interface{}) error {
	if len(r.Items) == 0 {
		return nil
	}
	stuRef := reflect.ValueOf(impl).Elem()
	cTypeRef := reflect.TypeOf(impl).Elem().Elem()
	for _, v := range r.Items {
		d := reflect.New(cTypeRef.Elem())
		if d.Elem().FieldByName("Kind").CanSet() {
			d.Elem().FieldByName("Kind").SetString(v.Kind)
		}

		if d.Elem().FieldByName("ApiVersion").CanSet() {
			d.Elem().FieldByName("ApiVersion").SetString(v.APIVersion)
		}

		if d.Elem().FieldByName("Metadata").CanSet() {
			meta := reflect.New(d.Elem().FieldByName("Metadata").Type().Elem()).Interface()
			_ = jsoniter.Unmarshal(v.MeteData, meta)
			d.Elem().FieldByName("Metadata").Set(reflect.ValueOf(meta))
		}

		if d.Elem().FieldByName("Spec").CanSet() {
			spec := reflect.New(d.Elem().FieldByName("Spec").Type().Elem()).Interface()
			_ = jsoniter.Unmarshal(v.Spec, spec)
			d.Elem().FieldByName("Spec").Set(reflect.ValueOf(spec))
			stuRef.Set(reflect.Append(stuRef, d))
		}

	}
	return nil
}

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
