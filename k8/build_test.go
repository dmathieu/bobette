package k8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func TestRunBuild(t *testing.T) {
	client := fake.NewSimpleClientset()
	k := &K8{Client: client}
	err := k.RunBuild(&BuildConfig{})
	assert.Nil(t, err)
}
