package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteOrganization(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s", id)
	return apiClient.DeleteResource(path)
}

func deleteCluster(orgID string, clusterID string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", orgID, clusterID)
	return provisionClient.DeleteResource(path)
}

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
			orgID := args[1]
			err := deleteOrganization(orgID)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Organization %s successfully deleted\n", orgID)
			}

		case "cluster", "clusters":
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			clusterID := args[1]
			err := deleteCluster(organizationID, clusterID)
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
