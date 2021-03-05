package resources

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/security/v1beta1"
	kubev1 "k8s.io/api/core/v1"
)

type Gateway v1alpha3.Gateway

func init() {
	eventType["*v1alpha3.Gateway"] = v1alpha3Gateway
	eventType["*v1alpha3.EnvoyFilter"] = v1alpha3EnvoyFilter
	eventType["*v1alpha3.VirtualService"] = v1alpha3VirtualService
	eventType["*v1alpha3.DestinationRule"] = v1alpha3DestinationRule
	eventType["*v1alpha3.Sidecar"] = v1alpha3Sidecar
	eventType["*v1alpha3.ServiceEntry"] = v1alpha3ServiceEntry
	eventType["*v1beta1.RequestAuthentication"] = v1beta1RequestAuthentication
	eventType["*v1beta1.PeerAuthentication"] = v1beta1PeerAuthentication
	eventType["*v1beta1.AuthorizationPolicy"] = v1beta1AuthorizationPolicy

	// kube
	eventType["*v1.Namespace"] = v1Namespace
}

// Namespace
func v1Namespace(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*kubev1.Namespace)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1beta1.PeerAuthentication %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "Namespace"
		kindObj.TypeMeta.APIVersion = "v1"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// PeerAuthentication
func v1beta1AuthorizationPolicy(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1beta1.PeerAuthentication)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1beta1.PeerAuthentication %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "PeerAuthentication"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1beta1"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// PeerAuthentication
func v1beta1PeerAuthentication(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1beta1.PeerAuthentication)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1beta1.PeerAuthentication %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "PeerAuthentication"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1beta1"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// RequestAuthentication
func v1beta1RequestAuthentication(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1beta1.RequestAuthentication)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1beta1.RequestAuthentication %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "RequestAuthentication"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1beta1"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// gateway
func v1alpha3Gateway(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1alpha3.Gateway)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1alpha3.Gateway %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "Gateway"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1alpha3"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// EnvoyFilter
func v1alpha3EnvoyFilter(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1alpha3.EnvoyFilter)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1alpha3.EnvoyFilter %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "EnvoyFilter"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1alpha3"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// EnvoyFilter
func v1alpha3VirtualService(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1alpha3.VirtualService)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1alpha3.VirtualService %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "VirtualService"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1alpha3"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// DestinationRule
func v1alpha3DestinationRule(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1alpha3.DestinationRule)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1alpha3.DestinationRule %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "DestinationRule"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1alpha3"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// Sidecar
func v1alpha3Sidecar(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1alpha3.Sidecar)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1alpha3.Sidecar %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "Sidecar"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1alpha3"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}

// ServiceEntry
func v1alpha3ServiceEntry(obj interface{}) (*KubeResources, error) {
	kindObj, ok := obj.(*v1alpha3.ServiceEntry)
	if !ok {
		return nil, errors.New(fmt.Sprintf("vent object assertion failed *v1alpha3.ServiceEntry %v", obj))
	}
	if kindObj.TypeMeta.Kind == "" {
		kindObj.TypeMeta.Kind = "ServiceEntry"
		kindObj.TypeMeta.APIVersion = "networking.istio.io/v1alpha3"
	}
	resources := &KubeResources{
		Name:       kindObj.ObjectMeta.Name,
		Namespace:  kindObj.ObjectMeta.Namespace,
		Kind:       kindObj.Kind,
		APIVersion: kindObj.APIVersion,
	}
	resources.MeteData, _ = jsoniter.Marshal(kindObj.ObjectMeta)
	resources.Spec, _ = jsoniter.Marshal(kindObj.Spec)
	resources.Status, _ = jsoniter.Marshal(kindObj.Status)
	return resources, nil
}
