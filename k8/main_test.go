package k8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func TestNew(t *testing.T) {
	t.Run("with an invalid config", func(t *testing.T) {
		k, err := New("missingconfig")
		assert.Nil(t, k)
		assert.NotNil(t, err)
	})

	t.Run("with a valid config", func(t *testing.T) {
		k, err := New("../fixtures/kubeconfig")
		assert.Nil(t, err)
		assert.NotNil(t, k)
		assert.Equal(t, "", k.arch)
	})

	t.Run("setting the arch option", func(t *testing.T) {
		k, err := New("../fixtures/kubeconfig", Arch("arm"))
		assert.Nil(t, err)
		assert.NotNil(t, k)
		assert.Equal(t, "arm", k.arch)
	})
}

func TestRunBuild(t *testing.T) {
	client := fake.NewSimpleClientset()
	k := &K8{Client: client}
	err := k.RunBuild()
	assert.Nil(t, err)
}
