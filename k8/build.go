package k8

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	privileged = true
	optional   = true
)

func (k *K8) imageName() (string, error) {
	name := "gcr.io/dmathieu-191516/bobette"
	if k.arch == "" {
		m, err := k.masterNode()
		if err != nil {
			return "", err
		}
		k.arch = m.Status.NodeInfo.Architecture
	}

	switch k.arch {
	case "amd64":
		// Amd64 is the default image. Do nothing
	default:
		name = fmt.Sprintf("%s-%s", name, k.arch)
	}

	return name, nil
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
