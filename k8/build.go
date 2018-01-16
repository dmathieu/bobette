package k8

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BuildConfig is the build configuration, passed to the pod
type BuildConfig struct {
	RepositoryURL string `json:"repository_url"`
}

func (k *K8) imageName() string {
	base := "gcr.io/dmathieu-191516/bobette"
	if k.arch == "" {
		return base
	}
	return fmt.Sprintf("%s-%s", base, k.arch)
}

// RunBuild starts a pod build
func (k *K8) RunBuild(c *BuildConfig) error {
	config, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = k.Client.CoreV1().Pods("default").Create(&corev1.Pod{
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
					Env: []corev1.EnvVar{
						corev1.EnvVar{
							Name:  "BOBETTE_CONFIG",
							Value: base64.StdEncoding.EncodeToString(config),
						},
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	})

	return err
}
