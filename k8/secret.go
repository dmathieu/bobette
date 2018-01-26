package k8

import (
	"encoding/base64"
	"fmt"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8) secretName(url string) string {
	return fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))
}

// GetSecret returns the specified url's secrets
func (k *K8) GetSecret(url string) (*corev1.Secret, error) {
	s, err := k.Client.CoreV1().Secrets("default").Get(k.secretName(url), metav1.GetOptions{})
	if err != nil {
		if err.(*errors.StatusError).ErrStatus.Code == http.StatusNotFound {
			return &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k.secretName(url),
					Namespace: "default",
				},
			}, nil
		}
		return nil, err
	}

	return s, err
}
