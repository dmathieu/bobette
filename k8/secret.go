package k8

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8) secretName(url string) string {
	return strings.ToLower(fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url))))
}

// SetSecret sets a config value in the specified url's secret
func (k *K8) SetSecret(url string, args ...string) error {
	if len(args) == 0 {
		return errors.New("no config provided. Need config vars in the format 'key=value'")
	}

	data := map[string][]byte{}
	for _, e := range args {
		d := strings.Split(e, "=")
		if len(d) == 1 || len(d[1]) == 0 {
			data[d[0]] = nil
		} else {
			data[d[0]] = []byte(d[1])
		}
	}

	s, err := k.Client.CoreV1().Secrets(defaultNamespace).Get(k.secretName(url), metav1.GetOptions{})
	if err != nil && err.(*kerrors.StatusError).ErrStatus.Code == http.StatusNotFound {
		// We need to create the secret
		_, err = k.Client.CoreV1().Secrets(defaultNamespace).Create(&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      k.secretName(url),
				Namespace: defaultNamespace,
			},
			Data: data,
		})
		return err
	}

	for key, value := range data {
		if value == nil {
			delete(s.Data, key)
		} else {
			s.Data[key] = value
		}
	}
	_, err = k.Client.CoreV1().Secrets(defaultNamespace).Update(s)
	return err
}

// GetSecret returns the specified url's secret
func (k *K8) GetSecret(url string) (*corev1.Secret, error) {
	s, err := k.Client.CoreV1().Secrets(defaultNamespace).Get(k.secretName(url), metav1.GetOptions{})
	if err != nil {
		if err.(*kerrors.StatusError).ErrStatus.Code == http.StatusNotFound {
			return &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k.secretName(url),
					Namespace: defaultNamespace,
				},
			}, nil
		}
		return nil, err
	}

	return s, err
}
