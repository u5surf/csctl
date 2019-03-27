package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getOrganizationCmd represents the getOrganization command
var getOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Short:   "Get an organization or list of organizations",
	Aliases: resource.Organization().Aliases(),

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.Organization, 1)
		var err error
		if len(args) == 1 {
			var v *types.Organization
			v, err = clientset.API().Organizations().Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.API().Organizations().List()
		}

		if err != nil {
			return err
		}

		orgs := resource.NewOrganizations(resp)

		if len(args) == 1 {
			resource.Organization().DisableListView()
		}

		if mineOnly {
			me, err := getMyAccountID()
			if err != nil {
				return err
			}

			orgs.FilterByOwnerID(me)
		}

		outputResponse(orgs)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getOrganizationCmd)
}
