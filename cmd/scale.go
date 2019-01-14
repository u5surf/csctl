package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
)

const countUnset int32 = -1

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
		np, err := clientset.Provision().NodePools(organizationID, clusterID).Get(nodePoolID)
		if err != nil {
			return errors.Wrapf(err, "getting node pool %q", nodePoolID)
		}

		targetCount, err := validateAndGetTargetCount(*np.Count, upCount, downCount, toCount)
		if err != nil {
			return err
		}

		req := types.NodePoolScaleRequest{
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

func validateAndGetTargetCount(from, up, down, to int32) (int32, error) {
	if !exactlyOneSet(up, down, to) {
		return 0, errors.New("must specify exactly one of --up, --down, or --to")
	}

	var target int32
	switch {
	case up != countUnset:
		if up < 1 {
			return 0, errors.New("can't scale up by less than 1 node")
		}

		target = from + up

	case down != countUnset:
		if down < 1 {
			return 0, errors.New("can't scale down by less than 1 node")
		}

		target = from - down

	case to != countUnset:
		target = to
	}

	if target < 0 {
		return 0, errors.New("can't scale below 0")
	}

	return target, nil
}

func exactlyOneSet(v1, v2, v3 int32) bool {
	switch {
	case v1 == countUnset && v2 == countUnset && v3 == countUnset:
		return false
	case v1 != countUnset && (v2 != countUnset || v3 != countUnset):
		return false
	case v2 != countUnset && (v1 != countUnset || v3 != countUnset):
		return false
	}

	return true
}

func init() {
	rootCmd.AddCommand(scaleCmd)

	bindCommandToNodePoolScope(scaleCmd, false)

	scaleCmd.Flags().Int32VarP(&upCount, "up", "u", countUnset, "scale up by N nodes")
	scaleCmd.Flags().Int32VarP(&downCount, "down", "d", countUnset, "scale down by N nodes")
	scaleCmd.Flags().Int32VarP(&toCount, "to", "t", countUnset, "scale to N nodes")
}
