package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the current version of csctl",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO implement versioning")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
