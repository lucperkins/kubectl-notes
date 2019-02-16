package cmd

import (
	"github.com/lucperkins/kubectl-docs/pkg/kubectl"
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
	client, err := kubectl.NewClient()

	exitOnErr(err)

	requirePodName()

	note := args[0]

	client.AddNote(note)
}
