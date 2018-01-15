package k8

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// K8 allows our cluster manipulations
type K8 struct {
}

func (k *K8) config() (*rest.Config, error) {
	return rest.InClusterConfig()
}

func (k *K8) client() (*kubernetes.Clientset, error) {
	config, err := k.config()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}
