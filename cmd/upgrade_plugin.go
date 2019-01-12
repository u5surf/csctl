package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

var (
	version string
)

// upgradePluginCmd represents the upgradePlugin command
var upgradePluginCmd = &cobra.Command{
	Use:     "plugin",
	Short:   "Upgrade a plugin",
	Aliases: resource.Plugin().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		if version == "" {
			return errors.New("please supply target version using --version")
		}

		req := &types.PluginUpgradeRequest{
			Version: &version,
		}

		id := args[0]

		err := clientset.API().Plugins(organizationID, clusterID).Upgrade(id, req)
		if err != nil {
			return err
		}

		fmt.Printf("Plugin %s upgrade to %s successfully initiated\n", id, version)

		return nil
	},
}

func init() {
	upgradeCmd.AddCommand(upgradePluginCmd)

	bindCommandToClusterScope(upgradePluginCmd, false)

	upgradePluginCmd.Flags().StringVarP(&version, "version", "v", "", "plugin version to upgrade to")
}
