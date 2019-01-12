package cmd

import (
	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgradeCmd command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a cluster node pool or resource",
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
