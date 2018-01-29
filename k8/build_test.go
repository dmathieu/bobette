package k8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestImageName(t *testing.T) {
	t.Run("when the arch is set manually", func(t *testing.T) {
		client := fake.NewSimpleClientset()
		k := &K8{Client: client, arch: "arm"}

		i, err := k.imageName()
		assert.Nil(t, err)
		assert.Equal(t, "gcr.io/dmathieu-191516/bobette-arm", i)
	})

	t.Run("when the arch is set manually to amd64", func(t *testing.T) {
		client := fake.NewSimpleClientset()
		k := &K8{Client: client, arch: "amd64"}

		i, err := k.imageName()
		assert.Nil(t, err)
		assert.Equal(t, "gcr.io/dmathieu-191516/bobette", i)
	})

	t.Run("when the arch is not set manually", func(t *testing.T) {
		node := &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "kubemaster",
				Labels: map[string]string{
					"node-role.kubernetes.io/master": "",
				},
			},
			Status: corev1.NodeStatus{
				NodeInfo: corev1.NodeSystemInfo{
					Architecture: "amd",
				},
			},
		}
		client := fake.NewSimpleClientset(node)
		k := &K8{Client: client}

		i, err := k.imageName()
		assert.Nil(t, err)
		assert.Equal(t, "gcr.io/dmathieu-191516/bobette-amd", i)
	})
}

func TestRunBuild(t *testing.T) {
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
	err := k.RunBuild("https://example.com")
	assert.Nil(t, err)
}
