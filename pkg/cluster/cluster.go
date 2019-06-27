package cluster

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Cluster struct {
	firstNodeIP string
}

func New() *Cluster {
	return &Cluster{}
}

func CreatePod(name string, client kubernetes.Interface) error {
	labels := map[string]string{
		"app":        "consuladm",
		"controller": name,
	}
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Labels:      labels,
			Annotations: map[string]string{},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "consuladm",
					Image:   "docker.io/nirmoy/consuladm",
					Command: []string{"consuladm", "init"},
				},
			},
		},
	}
	_, err := client.CoreV1().Pods("default").Create(pod)
	return err
}
