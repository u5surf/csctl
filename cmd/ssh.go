package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	username string
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH to nodes in a cluster",
	Long: `SSH to nodes in a cluster.

TODO this is a long description`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh called")
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)

	sshCmd.Flags().StringVarP(&username, "username", "u", "containership", "Username to login with")
}
