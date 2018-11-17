package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deletePluginCmd represents the deletePlugin command
var deletePluginCmd = &cobra.Command{
	Use:     "plugin",
	Short:   "Delete a plugin",
	Aliases: resource.Plugin().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.API().Plugins(organizationID, clusterID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Plugin %s delete successfully initiated\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deletePluginCmd)

	bindCommandToClusterScope(deletePluginCmd, false)
}
