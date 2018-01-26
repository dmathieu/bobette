package k8

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetSecret(t *testing.T) {
	url := "https://example.com"

	t.Run("with no secret set", func(t *testing.T) {
		client := fake.NewSimpleClientset()
		k := &K8{Client: client}

		s, err := k.GetSecret(url)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(s.Data))
	})

	t.Run("with secrets set", func(t *testing.T) {
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url))),
				Namespace: "default",
			},
			Data: map[string][]byte{
				"foo":   []byte("bar"),
				"hello": []byte("world"),
			},
		}
		client := fake.NewSimpleClientset(secret)
		k := &K8{Client: client}

		s, err := k.GetSecret(url)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(s.Data))
	})
}
