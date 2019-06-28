package cluster

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes"
)

type member struct {
	podName    string
	isFirstPod bool
}

type Cluster struct {
	firstNodeIP string
	memeberList []member
}

func New() *Cluster {
	return &Cluster{}
}

func getRandomPodName(clusterName string) string {
	suffix := utilrand.String(10)
	return clusterName + "-" + suffix
}

func (c *Cluster) Sync(name string, client kubernetes.Interface, size int32) {
	c.createFirstPod(name, client)
}

func (c *Cluster) createFirstPod(name string, client kubernetes.Interface) error {

	if len(c.memeberList) > 0 {
		return nil
	}
	podName := getRandomPodName(name)
	labels := map[string]string{
		"app":        "consuladm",
		"controller": name,
	}
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        podName,
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
	if err != nil {
		c.memeberList = append(c.memeberList, member{podName: podName})
	}
	return err
}
