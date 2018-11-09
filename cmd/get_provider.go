package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getProviderCmd represents the getProvider command
var getProviderCmd = &cobra.Command{
	Use:     "provider",
	Short:   "Get a provider or list of providers",
	Aliases: resource.Provider().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Provider, 1)
		var err error
		if len(args) == 1 {
			var v *types.Provider
			v, err = clientset.API().Providers(organizationID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.API().Providers(organizationID).List()
		}

		if err != nil {
			return err
		}

		providers := resource.NewProviders(resp)

		outputResponse(providers)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getProviderCmd)

	bindCommandToOrganizationScope(getProviderCmd, false)
}
