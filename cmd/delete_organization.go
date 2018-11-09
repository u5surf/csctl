package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteOrganizationCmd represents the deleteOrganization command
var deleteOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Short:   "Delete an organization",
	Aliases: resource.Organization().Aliases(),

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		err := clientset.API().Organizations().Delete(id)
		if err != nil {
			return err
		}

		fmt.Printf("Organization %s successfully deleted\n", id)
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteOrganizationCmd)
}
