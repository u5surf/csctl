package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource",
	Long: `Delete a resource

TODO this is a long description`,

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		resourceName := args[0]
		switch {
		case resource.Organization().HasAlias(resourceName):
			id := args[1]
			err := clientset.API().Organizations().Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Organization %s successfully deleted\n", id)
			}

		case resource.CKECluster().HasAlias(resourceName):
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			id := args[1]
			err := clientset.Provision().CKEClusters(organizationID).Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Cluster %s delete successfully initiated\n", clusterID)
			}

		case resource.NodePool().HasAlias(resourceName):
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			id := args[1]
			err := clientset.Provision().NodePools(organizationID, clusterID).Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Node pool %s delete successfully initiated\n", id)
			}

		case resource.Plugin().HasAlias(resourceName):
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			id := args[1]
			err := clientset.API().Plugins(organizationID, clusterID).Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Plugin %s successfully deleted\n", id)
			}

		case resource.Template().HasAlias(resourceName):
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			id := args[1]
			err := clientset.Provision().Templates(organizationID).Delete(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Template %s successfully deleted\n", id)
			}

		case resource.Account().HasAlias(resourceName):
			fmt.Println("Error: accounts cannot be deleted")

		default:
			fmt.Printf("Error: invalid resource name specified: %q\n", resourceName)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
