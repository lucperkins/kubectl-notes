package kubectl

import (
	"fmt"
	"os"

	"github.com/lucperkins/ezk8sclient"

	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

type Client struct {
	k8s       *kubernetes.Clientset
	namespace string
}

func NewClient() *Client {
	k8s, err := ezk8sclient.NewClient()
	exitOnErr(err)

	ns := viper.GetString("namespace")

	return &Client{
		k8s:       k8s,
		namespace: ns,
	}
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
