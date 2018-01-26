package k8

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestMasterNode(t *testing.T) {
	t.Run("with no nodes", func(t *testing.T) {
		client := fake.NewSimpleClientset()
		k := &K8{Client: client}

		_, err := k.masterNode()
		assert.Equal(t, errors.New("no master node found"), err)
	})

	t.Run("with nodes, but no master", func(t *testing.T) {
		node := &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name:   "kubemaster",
				Labels: map[string]string{},
			},
		}

		client := fake.NewSimpleClientset(node)
		k := &K8{Client: client}

		_, err := k.masterNode()
		assert.Equal(t, errors.New("no master node found"), err)
	})

	t.Run("with a master node", func(t *testing.T) {
		node := &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "kubemaster",
				Labels: map[string]string{
					"node-role.kubernetes.io/master": "",
				},
			},
		}

		client := fake.NewSimpleClientset(node)
		k := &K8{Client: client}

		n, err := k.masterNode()
		assert.Nil(t, err)
		assert.Equal(t, "kubemaster", n.ObjectMeta.Name)
	})

	t.Run("with several master nodes", func(t *testing.T) {
		node1 := &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "kubemaster",
				Labels: map[string]string{
					"node-role.kubernetes.io/master": "",
				},
			},
		}
		node2 := &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "kubemaster2",
				Labels: map[string]string{
					"node-role.kubernetes.io/master": "",
				},
			},
		}

		client := fake.NewSimpleClientset(node1, node2)
		k := &K8{Client: client}

		n, err := k.masterNode()
		assert.Nil(t, err)
		assert.Equal(t, "kubemaster", n.ObjectMeta.Name)
	})
}
