package cluster

import (
	"container/list"
	"fmt"
	"math"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes"
)

type member struct {
	podName    string
	isFirstPod bool
}

type Cluster struct {
	firstNodeName string
	memeberList   *list.List
	kubeClient    kubernetes.Interface
}

func New(kubeClient kubernetes.Interface) *Cluster {
	fmt.Println("Called New")
	return &Cluster{kubeClient: kubeClient, memeberList: &list.List{}}
}

func getRandomPodName(clusterName string) string {
	suffix := utilrand.String(10)
	return clusterName + "-" + suffix
}

func (c *Cluster) Sync(name string, size int32) error {
	fmt.Println("size:", size)
	if c.memeberList.Len() == 0 {
		podName, err := c.createPod(name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		c.firstNodeName = podName
		c.setPodFlag(podName)
		size -= 1
	}
	diff := c.memeberList.Len() - int(size)
	fmt.Println("diff:", diff)
	if diff == 0 {
		return nil
	}

	for i := 0; i < int(math.Abs(float64(size))); i++ {
		fmt.Println(i, diff)
		if diff > 0 {
			c.deletePod()
		} else {
			c.createPod(name)
		}
	}
	return nil
}

func (c *Cluster) createService() error {
	labels := map[string]string{
		"app":        "consuladm",
		"controller": "example-consul-operator",
	}
	ports := []corev1.ServicePort{{
		Name:       "server",
		Port:       8500,
		TargetPort: intstr.FromInt(8500),
		Protocol:   corev1.ProtocolTCP,
	}}
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "consul",
			Labels: labels,
		},
		Spec: corev1.ServiceSpec{
			Ports:     ports,
			Selector:  labels,
			ClusterIP: "",
		},
	}
	_, err := c.kubeClient.CoreV1().Services("default").Create(svc)
	return err
}

func (c *Cluster) CreateServices() {
	c.createService()
}

func (c *Cluster) createPod(name string) (string, error) {

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
					Image:   "docker.io/nirmoy/consuladm:latest",
					Command: []string{},
				},
			},
			Hostname:  podName,
			Subdomain: name,
		},
	}

	if c.memeberList.Len() == 0 {
		pod.Spec.Containers[0].Command = []string{"consuladm", "init"}
	} else {
		masterpod := c.firstNodeName + "." + name + ".default.svc"
		pod.Spec.Containers[0].Command = []string{"consuladm", "join", masterpod}
	}
	_, err := c.kubeClient.CoreV1().Pods("default").Create(pod)
	if err == nil {
		c.memeberList.PushBack(&member{podName: podName})
	}

	return podName, err
}

func (c *Cluster) setPodFlag(podName string) {
	for e := c.memeberList.Front(); e != nil; e = e.Next() {
		if e.Value.(*member).podName == podName {
			e.Value.(*member).isFirstPod = true
		}
	}
}

func (c *Cluster) deletePod() {
	fmt.Println("deleting a pod")
	for e := c.memeberList.Front(); e != nil; e = e.Next() {
		if !e.Value.(*member).isFirstPod {
			fmt.Println("deleting  pod:", e.Value.(*member).podName)
			opts := metav1.NewDeleteOptions(10)
			c.kubeClient.CoreV1().Pods("delete").Delete(e.Value.(*member).podName, opts)
			c.memeberList.Remove(e)
		}
	}
}
