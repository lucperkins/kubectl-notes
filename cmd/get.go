package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(getCmd())
}

func getCmd() *cobra.Command {
	return &cobra.Command{
		Use: "get",
		Short: "Fetch the note from a specified Pod",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Namespace: %s\n", viper.GetString("namespace"))
		},
	}
}
