package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	upCount   int32
	downCount int32
	toCount   int32
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scale a node pool",
	Long:  `Scale a node pool up/down by N nodes or to a target node count`,

	Args: cobra.NoArgs,

	PreRunE: nodePoolScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		if !exactlyOneSet(upCount, downCount, toCount) {
			return errors.New("must specify exactly one of --up, --down, or --to")
		}

		np, err := clientset.Provision().NodePools(organizationID, clusterID).Get(nodePoolID)
		if err != nil {
			return errors.Wrapf(err, "getting node pool %q", nodePoolID)
		}

		targetCount := *np.Count

		switch {
		case upCount != 0:
			targetCount += upCount
		case downCount != 0:
			targetCount -= upCount
		case toCount != 0:
			targetCount = upCount
		}

		req := types.ScaleNodePoolRequest{
			Count: &targetCount,
		}

		_, err = clientset.Provision().NodePools(organizationID, clusterID).Scale(nodePoolID, &req)
		if err != nil {
			return errors.Wrapf(err, "patching node pool %q", nodePoolID)
		}

		fmt.Printf("Scale node pool %q to %d initiated successfully\n", nodePoolID, targetCount)

		return nil
	},
}

func exactlyOneSet(v1, v2, v3 int32) bool {
	switch {
	case v1 == 0 && v2 == 0 && v3 == 0:
		return false
	case v1 != 0 && (v2 != 0 || v3 != 0):
		return false
	case v2 != 0 && (v1 != 0 || v3 != 0):
		return false
	case v3 != 0 && (v1 != 0 || v2 != 0):
		return false
	}

	return true
}

func init() {
	rootCmd.AddCommand(scaleCmd)

	bindCommandToNodePoolScope(scaleCmd, false)

	scaleCmd.Flags().Int32VarP(&upCount, "up", "u", 0, "scale up by N nodes")
	scaleCmd.Flags().Int32VarP(&downCount, "down", "d", 0, "scale down by N nodes")
	scaleCmd.Flags().Int32VarP(&toCount, "to", "t", 0, "scale to N nodes")
}
