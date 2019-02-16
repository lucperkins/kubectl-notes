package cmd

import (
	"github.com/lucperkins/kubectl-notes/pkg/kubectl"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Pod's note",
	Run:   deleteCmdRun,
}

func deleteCmdRun(_ *cobra.Command, _ []string) {
	client := kubectl.NewClient()

	requirePodName()

	client.DeleteNote()
}
