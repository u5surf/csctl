package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteClusterCmd represents the deleteCluster command
var deleteClusterCmd = &cobra.Command{
	Use:     "cluster",
	Short:   "Delete one or more clusters",
	Aliases: resource.CKECluster().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.Provision().CKEClusters(organizationID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Cluster %s delete successfully initiated\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteClusterCmd)

	bindCommandToOrganizationScope(deleteClusterCmd, false)
}
