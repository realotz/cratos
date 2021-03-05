package data

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"sync"
)

type KubeWatcher func(client *kubernetes.Clientset) (watch.Interface, error)

var (
	kubeWatchers = make([]KubeWatcher, 0)
	kubeMutex    sync.Mutex
)

// 注册Kube相关资源观察者
func RegisterKubeWatcher(f KubeWatcher) {
	kubeMutex.Lock()
	defer kubeMutex.Unlock()
	kubeWatchers = append(kubeWatchers, f)
}

func init() {
	// Namespace 网关观察者
	RegisterKubeWatcher(func(client *kubernetes.Clientset) (watch.Interface, error) {
		return client.CoreV1().Namespaces().Watch(context.Background(), v1.ListOptions{
			Watch: true,
		})
	})
}
