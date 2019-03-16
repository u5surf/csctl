package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteAccessTokenCmd represents the deleteAccessToken command
var deleteAccessTokenCmd = &cobra.Command{
	Use:     "access-token",
	Short:   "Delete one or more access tokens",
	Aliases: resource.AccessToken().Aliases(),

	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.API().AccessTokens().Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Access token %s deleted successfully\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteAccessTokenCmd)
}
