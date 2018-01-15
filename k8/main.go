package k8

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// K8 allows our cluster manipulations
type K8 struct {
	Arch   string
	Client kubernetes.Interface
}

// Arch sets a CPU architecture for the kubernetes client
func Arch(arch string) func(*K8) {
	return func(k *K8) {
		k.Arch = arch
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

func (k *K8) imageName() string {
	base := "gcr.io/dmathieu-191516/bobette"
	if k.Arch == "" {
		return base
	}
	return fmt.Sprintf("%s-%s", base, k.Arch)
}

// RunBuild starts a pod build
func (k *K8) RunBuild() error {
	podName := "build"

	_, err := k.Client.CoreV1().Pods("default").Create(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				corev1.Container{
					Name:  podName,
					Image: k.imageName(),
					Env:   []corev1.EnvVar{},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	})

	return err
}
