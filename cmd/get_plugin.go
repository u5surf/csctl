package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getPluginCmd represents the getPlugin command
var getPluginCmd = &cobra.Command{
	Use:     "plugin",
	Short:   "Get a plugin or list of plugins",
	Aliases: resource.Plugin().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Plugin, 1)
		var err error
		if len(args) == 1 {
			var v *types.Plugin
			v, err = clientset.API().Plugins(organizationID, clusterID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.API().Plugins(organizationID, clusterID).List()
		}

		if err != nil {
			return err
		}

		plugs := resource.NewPlugins(resp)

		if len(args) == 1 {
			resource.Plugin().DisableListView()
		}

		outputResponse(plugs)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getPluginCmd)

	bindCommandToClusterScope(getPluginCmd, false)
}
