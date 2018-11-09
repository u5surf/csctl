package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteNodePoolCmd represents the deleteNodePool command
var deleteNodePoolCmd = &cobra.Command{
	Use:     "nodepool",
	Short:   "Delete a node pool",
	Aliases: resource.NodePool().Aliases(),

	Args: cobra.ExactArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		err := clientset.Provision().NodePools(organizationID, clusterID).Delete(id)
		if err != nil {
			return err
		}

		fmt.Printf("Node pool %s delete successfully initiated\n", id)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteNodePoolCmd)

	bindCommandToClusterScope(deleteNodePoolCmd, false)
}
