package k8

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	imageBase = "gcr.io/dmathieu-191516/bobette"
)

var (
	privileged = true
	optional   = true
)

func (k *K8) imageName() (string, error) {
	if k.arch == "" {
		m, err := k.masterNode()
		if err != nil {
			return "", err
		}
		k.arch = m.Status.NodeInfo.Architecture
	}
	return fmt.Sprintf("%s-%s", imageBase, k.arch), nil
}

// RunBuild starts a pod build
func (k *K8) RunBuild(url string) error {
	env, err := k.buildEnvironment(url)
	if err != nil {
		return err
	}
	imageName, err := k.imageName()
	if err != nil {
		return err
	}

	_, err = k.Client.CoreV1().Pods(defaultNamespace).Create(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "build-",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				corev1.Container{
					Name:  "bobette",
					Image: imageName,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &privileged,
					},
					Env: env,
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	})

	return err
}
