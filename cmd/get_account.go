package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource"
)

// getAccountCmd represents the getAccount command
var getAccountCmd = &cobra.Command{
	Use:     "account",
	Short:   "Get your account",
	Aliases: resource.Account().Aliases(),

	RunE: func(cmd *cobra.Command, args []string) error {
		// A user can only get their own account
		var resp = make([]types.Account, 1)
		v, err := clientset.API().Account().Get()
		resp[0] = *v
		if err != nil {
			return err
		}

		accounts := resource.NewAccounts(resp)

		outputResponse(accounts)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getAccountCmd)
}
