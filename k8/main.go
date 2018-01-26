package k8

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	defaultNamespace = "default"
)

// K8 allows our cluster manipulations
type K8 struct {
	arch   string
	Client kubernetes.Interface
}

// Arch sets a CPU architecture for the kubernetes client
func Arch(arch string) func(*K8) {
	return func(k *K8) {
		k.arch = arch
	}
}

// New creates a new K8 instance
func New(config string, opts ...func(*K8)) (*K8, error) {
	c, err := client(config)
	if err != nil {
		return nil, err
	}
	k := &K8{Client: c}
	for _, opt := range opts {
		opt(k)
	}

	return k, nil
}

func client(configPath string) (*kubernetes.Clientset, error) {
	c, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(c)
}
