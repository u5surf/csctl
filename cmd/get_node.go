package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

// getNodeCmd represents the getNode command
var getNodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "Get a node or list of nodes",
	Aliases: resource.Node().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Node, 1)
		var err error
		if len(args) == 1 {
			var v *types.Node
			v, err = clientset.Provision().Nodes(organizationID, clusterID, nodePoolID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.Provision().Nodes(organizationID, clusterID, nodePoolID).List()
		}

		if err != nil {
			return err
		}

		nps := resource.NewNodes(resp)

		if len(args) == 1 {
			resource.Node().DisableListView()
		}

		outputResponse(nps)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getNodeCmd)

	bindCommandToNodePoolScope(getNodeCmd, false)
}
