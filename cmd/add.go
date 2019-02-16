package cmd

import (
	"github.com/lucperkins/kubectl-notes/pkg/kubectl"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a note to a specified Pod",
	Args: cobra.ExactArgs(1),
	Run: addCmdRun,
}

func addCmdRun(_ *cobra.Command, args []string) {
	client := kubectl.NewClient()

	requirePodName()

	note := args[0]

	client.AddNote(note)
}
