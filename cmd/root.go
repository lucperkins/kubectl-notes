package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "kubectl-notes",
		Short: "A kubectl extension for manipulating Pod notes",
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("namespace", "n", "default", "The Kubernetes namespace")
	viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
