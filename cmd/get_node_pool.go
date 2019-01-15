package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

// getNodePoolCmd represents the getNodePool command
var getNodePoolCmd = &cobra.Command{
	Use:     "node-pool",
	Short:   "Get a node pool or list of node pools",
	Aliases: resource.NodePool().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.NodePool, 1)
		var err error
		if len(args) == 1 {
			var v *types.NodePool
			v, err = clientset.Provision().NodePools(organizationID, clusterID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.Provision().NodePools(organizationID, clusterID).List()
		}

		if err != nil {
			return err
		}

		nps := resource.NewNodePools(resp)
		outputResponse(nps)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getNodePoolCmd)

	bindCommandToClusterScope(getNodePoolCmd, false)
}
