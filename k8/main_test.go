package k8

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestRunBuild(t *testing.T) {
	client := fake.NewSimpleClientset(&v1.Pod{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "test-app",
			Namespace: "default",
			Labels:    map[string]string{},
		},
		Status: v1.PodStatus{},
	})

	k := &K8{client}
	err := k.RunBuild()
	assert.Nil(t, err)
}
