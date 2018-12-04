package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// deleteAutoscalingPolicyCmd represents the deleteAutoscalingPolicy command
var deleteAutoscalingPolicyCmd = &cobra.Command{
	Use:     "autoscaling-policy",
	Short:   "Delete one or more autoscaling policies",
	Aliases: resource.AutoscalingPolicy().Aliases(),

	Args: cobra.MinimumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		for _, id := range args {
			err := clientset.Provision().AutoscalingPolicies(organizationID, clusterID).Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Autoscaling policy %s deleted successfully\n", id)
		}

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteAutoscalingPolicyCmd)

	bindCommandToClusterScope(deleteAutoscalingPolicyCmd, false)
}
