package k8

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
)

func (k *K8) buildEnvironment(url string) ([]corev1.EnvVar, error) {
	d := []corev1.EnvVar{
		corev1.EnvVar{
			Name:  "REPO_URL",
			Value: url,
		},
	}

	m, err := k.masterNode()
	if err != nil {
		return d, err
	}
	d = append(d, corev1.EnvVar{
		Name:  "ARCH",
		Value: m.Status.NodeInfo.Architecture,
	})

	s, err := k.GetSecret(url)
	if err != nil {
		return d, err
	}

	for key := range s.Data {
		d = append(d, corev1.EnvVar{
			Name: strings.ToUpper(key),
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: k.secretName(url),
					},
					Key: key,
				},
			},
		})
	}

	return d, nil
}
