package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

// getTemplateCmd represents the getTemplate command
var getTemplateCmd = &cobra.Command{
	Use:     "template",
	Short:   "Get a template or list of templates",
	Aliases: resource.Template().Aliases(),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Template, 1)
		var err error
		if len(args) == 2 {
			var v *types.Template
			v, err = clientset.Provision().Templates(organizationID).Get(args[1])
			resp[0] = *v
		} else {
			resp, err = clientset.Provision().Templates(organizationID).List()
		}

		if err != nil {
			return err
		}

		templates := resource.NewTemplates(resp)

		if mineOnly {
			me, err := getMyAccountID()
			if err != nil {
				return err
			}

			templates.FilterByOwnerID(me)
		}

		outputResponse(templates)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getTemplateCmd)

	bindCommandToOrganizationScope(getTemplateCmd, false)
}
