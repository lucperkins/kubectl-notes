package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringP("namespace", "n", "default", "The Kubernetes namespace")
	exitOnErr(viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace")))

	rootCmd.PersistentFlags().StringP("pod", "p", "", "The Kubernetes Pod")
	exitOnErr(viper.BindPFlag("pod", rootCmd.PersistentFlags().Lookup("pod")))

	rootCmd.SetVersionTemplate("kubectl-notes v{{ .Version }}\n")
}

var rootCmd = &cobra.Command{
	Use: "kubectl-notes",
	Short: "A kubectl extension for working with notes for Kubernetes Pods",
	Version: VERSION,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func requirePodName() {
	if viper.GetString("pod") == "" {
		fmt.Println("You must specify a Pod using the --pod or -p flag")
		os.Exit(1)
	}
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
