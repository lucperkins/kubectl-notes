package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use: "completion",
	Short: "Generates auto-complete scripts",
	Args: cobra.ExactArgs(1),
	Run: completionCmdRun,
	ValidArgs: []string{"bash", "zsh"},
}

func completionCmdRun(_ *cobra.Command, args []string) {
	shell := args[0]
	switch  shell {
	case "bash":
		exitOnErr(rootCmd.GenBashCompletion(os.Stdout))
	case "zsh":
		exitOnErr(rootCmd.GenZshCompletion(os.Stdout))
	}
}
