package k8

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// K8 allows our cluster manipulations
type K8 struct {
	Client kubernetes.Interface
}

// New creates a new K8 instance
func New(config string) (*K8, error) {
	c, err := client(config)
	if err != nil {
		return nil, err
	}
	return &K8{c}, nil
}

func client(configPath string) (*kubernetes.Clientset, error) {
	c, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(c)
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
					Image: "gcr.io/",
					Env:   []corev1.EnvVar{},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	})

	return err
}
