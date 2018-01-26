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

func (k *K8) imageName() string {
	base := "gcr.io/dmathieu-191516/bobette"
	if k.arch == "" {
		return base
	}
	return fmt.Sprintf("%s-%s", base, k.arch)
}

// RunBuild starts a pod build
func (k *K8) RunBuild(url string) error {
	env, err := k.buildEnvironment(url)
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
					Image: k.imageName(),
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
