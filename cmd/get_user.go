package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// getUserCmd represents the getUser command
var getUserCmd = &cobra.Command{
	Use:     "user",
	Short:   "Get a list of users",
	Aliases: resource.User().Aliases(),

	Args: cobra.NoArgs,

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		// Note that there is no route to get an individual user
		resp, err := clientset.API().Users(organizationID).List()

		if err != nil {
			return err
		}

		users := resource.NewUsers(resp)

		if len(args) == 1 {
			resource.User().DisableListView()
		}

		outputResponse(users)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getUserCmd)

	bindCommandToOrganizationScope(getUserCmd, false)
}
