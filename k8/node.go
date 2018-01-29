package k8

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *K8) masterNode() (corev1.Node, error) {
	if k.master.ObjectMeta.Name != "" {
		return k.master, nil
	}

	l, err := k.Client.CoreV1().Nodes().List(metav1.ListOptions{
		LabelSelector: "node-role.kubernetes.io/master",
	})
	if err != nil {
		return corev1.Node{}, err
	}

	if len(l.Items) == 0 {
		return corev1.Node{}, errors.New("no master node found")
	}
	k.master = l.Items[0]
	return k.master, err
}
