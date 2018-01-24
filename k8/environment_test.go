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

func TestBuildEnvironment(t *testing.T) {
	url := "https://example.com"

	t.Run("with no secret set", func(t *testing.T) {
		client := fake.NewSimpleClientset()
		k := &K8{Client: client}

		d, err := k.buildEnvironment(url)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(d))
		assert.Equal(t, "REPO_URL", d[0].Name)
		assert.Equal(t, url, d[0].Value)
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

		d, err := k.buildEnvironment(url)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(d))
		assert.Equal(t, "REPO_URL", d[0].Name)
		assert.Equal(t, "FOO", d[1].Name)
		assert.Equal(t, "HELLO", d[2].Name)
	})
}
