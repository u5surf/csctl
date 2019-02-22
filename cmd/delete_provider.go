package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteProviderCmd represents the deleteProvider command
var deleteProviderCmd = &cobra.Command{
	Use:     "provider",
	Short:   "Delete a provider",
	Aliases: resource.Provider().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.API().Providers(organizationID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Provider %s successfully deleted\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteProviderCmd)

	bindCommandToOrganizationScope(deleteProviderCmd, false)
}
