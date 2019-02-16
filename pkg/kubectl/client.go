package kubectl

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

type Client struct {
	k8s       *kubernetes.Clientset
	namespace string
}

func NewClient() *Client {
	flags := genericclioptions.NewConfigFlags()
	restConfig, err := flags.ToRESTConfig()

	exitOnErr(err)

	k8s, err := kubernetes.NewForConfig(restConfig)

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
