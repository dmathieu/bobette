package k8

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestSetSecret(t *testing.T) {
	url := "https://example.com"

	t.Run("when the secret does not exist yet", func(t *testing.T) {
		client := fake.NewSimpleClientset()
		k := &K8{Client: client}

		err := k.SetSecret(url, "foo=bar")
		assert.Nil(t, err)

		s, err := k.GetSecret(url)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(s.Data))
	})

	t.Run("when the secret already exists", func(t *testing.T) {
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))),
				Namespace: defaultNamespace,
			},
			Data: map[string][]byte{
				"hello": []byte("world"),
			},
		}
		client := fake.NewSimpleClientset(secret)
		k := &K8{Client: client}

		err := k.SetSecret(url, "foo=bar")
		assert.Nil(t, err)

		s, err := k.GetSecret(url)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(s.Data))
	})

	t.Run("when setting several values", func(t *testing.T) {
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))),
				Namespace: defaultNamespace,
			},
			Data: map[string][]byte{},
		}
		client := fake.NewSimpleClientset(secret)
		k := &K8{Client: client}

		err := k.SetSecret(url, "foo=bar", "hello=world")
		assert.Nil(t, err)

		s, err := k.GetSecret(url)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(s.Data))
	})

	t.Run("when removing a value", func(t *testing.T) {
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))),
				Namespace: defaultNamespace,
			},
			Data: map[string][]byte{
				"foo":   []byte("bar"),
				"hello": []byte("world"),
			},
		}
		client := fake.NewSimpleClientset(secret)
		k := &K8{Client: client}

		err := k.SetSecret(url, "foo=")
		assert.Nil(t, err)

		s, err := k.GetSecret(url)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(s.Data))
	})
}

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
				Name:      strings.ToLower(fmt.Sprintf("bobette-%s", base64.StdEncoding.EncodeToString([]byte(url)))),
				Namespace: defaultNamespace,
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
