package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "kubectl-docs",
	RunE: rootRunE,
}

func rootRunE(_ *cobra.Command, _ []string) error {
	log.Println("Hello world")

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
