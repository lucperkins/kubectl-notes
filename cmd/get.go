package cmd

import (
	"github.com/lucperkins/kubectl-notes/pkg/kubectl"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch the note (if any) from a Kubernetes Pod",
	Run:   getCmdRun,
}

func getCmdRun(_ *cobra.Command, _ []string) {
	client := kubectl.NewClient()

	requirePodName()

	client.GetNote()
}
