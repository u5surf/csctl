package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteOrganizationCmd represents the deleteOrganization command
var deleteOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Short:   "Delete one or more organizations",
	Aliases: resource.Organization().Aliases(),

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.API().Organizations().Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Organization %s successfully deleted\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteOrganizationCmd)
}
