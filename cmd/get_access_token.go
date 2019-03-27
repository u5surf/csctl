package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getAccessTokenCmd represents the getAccessToken command
var getAccessTokenCmd = &cobra.Command{
	Use:     "access-token",
	Short:   "Get an access token or list of access tokens",
	Aliases: resource.AccessToken().Aliases(),

	Args: cobra.MaximumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.AccessToken, 1)
		var err error
		if len(args) == 1 {
			var v *types.AccessToken
			v, err = clientset.API().AccessTokens().Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.API().AccessTokens().List()
		}

		if err != nil {
			return err
		}

		tokens := resource.NewAccessTokens(resp)

		if len(args) == 1 {
			resource.AccessToken().DisableListView()
		}

		outputResponse(tokens)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getAccessTokenCmd)
}
