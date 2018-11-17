package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteTemplateCmd represents the deleteTemplate command
var deleteTemplateCmd = &cobra.Command{
	Use:     "template",
	Short:   "Delete a template",
	Aliases: resource.Template().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.Provision().Templates(organizationID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Template %s successfully deleted\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteTemplateCmd)

	bindCommandToOrganizationScope(deleteTemplateCmd, false)
}
