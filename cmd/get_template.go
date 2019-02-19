package cmd

import (
	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/spf13/cobra"
)

// getTemplateCmd represents the getTemplate command
var getTemplateCmd = &cobra.Command{
	Use:     "template",
	Short:   "Get a template or list of templates",
	Aliases: resource.Template().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Template, 1)
		var err error
		if len(args) == 1 {
			var v *types.Template
			v, err = clientset.Provision().Templates(organizationID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.Provision().Templates(organizationID).List()
		}

		if err != nil {
			return err
		}

		templates := resource.NewTemplates(resp)

		if len(args) == 1 {
			templates.DisableItemListView()
		}

		if mineOnly {
			me, err := getMyAccountID()
			if err != nil {
				return err
			}

			templates.FilterByOwnerID(me)
		}

		// Non-CKE is deprecated. TODO consider filtering even earlier
		// since consumers should never care about non-CKE clusters.
		templates.FilterByEngine(types.TemplateEngineContainershipKubernetesEngine)

		outputResponse(templates)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getTemplateCmd)

	bindCommandToOrganizationScope(getTemplateCmd, false)
}
