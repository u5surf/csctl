package cmd

import (
	provisiontypes "github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/spf13/cobra"
)

// getClusterCmd represents the getCluster command
var getClusterCmd = &cobra.Command{
	Use:     "cluster",
	Short:   "Get a cluster or list of clusters",
	Aliases: resource.CKECluster().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]provisiontypes.CKECluster, 1)
		var err error
		if len(args) == 1 {
			var v *provisiontypes.CKECluster
			v, err = clientset.Provision().CKEClusters(organizationID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.Provision().CKEClusters(organizationID).List()
		}

		if err != nil {
			return err
		}

		clusters := resource.NewCKEClusters(resp)

		if len(args) == 1 {
			clusters.DisableItemListView()
		}

		if mineOnly {
			me, err := getMyAccountID()
			if err != nil {
				return err
			}

			clusters.FilterByOwnerID(me)
		}

		outputResponse(clusters)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getClusterCmd)

	bindCommandToOrganizationScope(getClusterCmd, false)
}
