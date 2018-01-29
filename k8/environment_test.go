package k8

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
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
		k := &K8{Client: client, master: corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "master",
			},
			Status: corev1.NodeStatus{
				NodeInfo: corev1.NodeSystemInfo{
					Architecture: "amd",
				},
			},
		}}

		d, err := k.buildEnvironment(url)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(d))
		assert.Equal(t, "REPO_URL", d[0].Name)

		names := []string{}
		for _, v := range d {
			names = append(names, v.Name)
		}
		sort.Strings(names)
		assert.Equal(t, []string{"ARCH", "REPO_URL"}, names)
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
		k := &K8{Client: client, master: corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "master",
			},
			Status: corev1.NodeStatus{
				NodeInfo: corev1.NodeSystemInfo{
					Architecture: "amd",
				},
			},
		}}

		d, err := k.buildEnvironment(url)
		assert.Nil(t, err)
		assert.Equal(t, 4, len(d))

		names := []string{}
		for _, v := range d {
			names = append(names, v.Name)
		}
		sort.Strings(names)
		assert.Equal(t, []string{"ARCH", "FOO", "HELLO", "REPO_URL"}, names)
	})
}
