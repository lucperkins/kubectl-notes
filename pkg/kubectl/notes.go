package kubectl

import (
	"fmt"
	"k8s.io/api/core/v1"
	"os"

	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

var (
	getOpts = metav1.GetOptions{}
)

func getPodName() string {
	return viper.GetString("pod")
}

func (c *Client) pods() corev1.PodInterface {
	return c.k8s.CoreV1().Pods(c.namespace)
}

func (c *Client) getPod(name string) (*v1.Pod, error) {
	return c.pods().Get(name, getOpts)
}

func (c *Client) GetNote() {
	p := getPodName()
	pod, err := c.getPod(p)
	exitOnErr(err)
	fmt.Println(pod.Annotations["note"])
}

func (c *Client) AddNote(note string) {
	p := getPodName()
	pod, err := c.getPod(p)
	exitOnErr(err)
	pod.Annotations["note"] = note
	_, err = c.pods().Update(pod)
	exitOnErr(err)

	fmt.Printf("Successfully added note to %s/%s\n", c.namespace, p)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
