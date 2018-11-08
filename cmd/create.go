package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
	Long: `Create a resource

TODO this is a long description`,
	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
