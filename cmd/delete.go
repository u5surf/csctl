package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource",
	Long: `Delete a resource

TODO this is a long description`,

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		switch resource {
		case "org", "orgs", "organization", "organizations":
			id := args[1]
			err := clientset.API().Organizations().Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Organization %s successfully deleted\n", id)
			}

		case "cluster", "clusters":
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			clusterID := args[1]
			err := clientset.Provision().CKEClusters(organizationID).Delete(clusterID)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Cluster %s successfully deleted\n", clusterID)
			}

		default:
			fmt.Printf("Error: invalid resource specified: %q\n", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
