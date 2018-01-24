package k8

import (
	"encoding/base64"
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
	secretName := fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))

	_, err := k.Client.CoreV1().Pods("default").Create(&corev1.Pod{
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
							Name:  "REPO_URL",
							Value: url,
						},
						corev1.EnvVar{
							Name: "REPO_AUTH",
							ValueFrom: &corev1.EnvVarSource{
								SecretKeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: secretName,
									},
									Key:      "repo_auth",
									Optional: &optional,
								},
							},
						},
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever,
		},
	})

	return err
}
