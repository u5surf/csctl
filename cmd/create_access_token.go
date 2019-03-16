package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var accessTokenOpts options.AccessTokenCreate

// createAccessTokenCmd represents the createAccessToken command
var createAccessTokenCmd = &cobra.Command{
	Use:     "access-token",
	Short:   "Create an access token",
	Aliases: resource.AccessToken().Aliases(),

	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := accessTokenOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := accessTokenOpts.CreateAccessTokenRequest()

		accessToken, err := clientset.API().AccessTokens().Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Access token %s created successfully. Copy the token below as it will not be displayed again:\n%s\n",
			string(*accessToken.Name),
			string(accessToken.Token))

		return nil
	},
}

func init() {
	createCmd.AddCommand(createAccessTokenCmd)
	createAccessTokenCmd.PersistentFlags().StringVarP(&accessTokenOpts.Name, "name", "n", "", "access token name")
}
