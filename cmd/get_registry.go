package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getRegistryCmd represents the getRegistry command
var getRegistryCmd = &cobra.Command{
	Use:     "registry",
	Short:   "Get a registry or list of registries",
	Aliases: resource.Registry().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Registry, 1)
		var err error
		if len(args) == 1 {
			var v *types.Registry
			v, err = clientset.API().Registries(organizationID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.API().Registries(organizationID).List()
		}

		if err != nil {
			return err
		}

		regs := resource.NewRegistries(resp)
		outputResponse(regs)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getRegistryCmd)

	bindCommandToClusterScope(getRegistryCmd, false)
}
