package k8

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8) buildEnvironment(url string) ([]corev1.EnvVar, error) {
	d := []corev1.EnvVar{
		corev1.EnvVar{
			Name:  "REPO_URL",
			Value: url,
		},
	}

	secretName := fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))
	s, err := k.Client.CoreV1().Secrets("default").Get(secretName, metav1.GetOptions{})
	if err != nil {
		if err.(*errors.StatusError).ErrStatus.Code == http.StatusNotFound {
			return d, nil
		}
		return nil, err
	}

	for key := range s.Data {
		d = append(d, corev1.EnvVar{
			Name: strings.ToUpper(key),
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secretName,
					},
					Key: key,
				},
			},
		})
	}

	return d, nil
}
