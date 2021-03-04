package data

import (
	"context"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"sync"
)

type IstioWatcher func(client *versionedclient.Clientset) (watch.Interface, error)

var (
	istioWatchers = make([]IstioWatcher, 0)
	istioMutex    sync.Mutex
)

// 注册istio相关资源观察者
func RegisterIstioWatcher(f IstioWatcher) {
	istioMutex.Lock()
	defer istioMutex.Unlock()
	istioWatchers = append(istioWatchers, f)
}

func init() {
	// gateway 网关观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.NetworkingV1alpha3().Gateways("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
	// EnvoyFilters 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.NetworkingV1alpha3().EnvoyFilters("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
	// VirtualServices 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.NetworkingV1alpha3().VirtualServices("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
	// DestinationRules 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.NetworkingV1alpha3().DestinationRules("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
	// Sidecars 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.NetworkingV1alpha3().Sidecars("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
	// ServiceEntries 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.NetworkingV1alpha3().ServiceEntries("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})

	// AuthorizationPolicies 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.SecurityV1beta1().AuthorizationPolicies("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})

	// PeerAuthentications 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.SecurityV1beta1().PeerAuthentications("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})

	// RequestAuthentications 观察者
	RegisterIstioWatcher(func(client *versionedclient.Clientset) (watch.Interface, error) {
		return client.SecurityV1beta1().RequestAuthentications("").Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
}
