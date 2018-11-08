package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// Flags
var (
	outputFormat string
	mineOnly     bool
)

func outputResponse(d resource.Displayable) {
	var err error
	switch {
	case outputFormat == "" || outputFormat == "table":
		err = d.Table(os.Stdout)

	case outputFormat == "json":
		err = d.JSON(os.Stdout)

	case outputFormat == "yaml":
		err = d.YAML(os.Stdout)

	case strings.HasPrefix(outputFormat, "jsonpath"):
		fields := strings.SplitN(outputFormat, "=", 2)
		if len(fields) != 2 {
			err = errors.New("Please specify jsonpath using -ojsonpath=<path>")
			break
		}

		template := fields[1]
		err = d.JSONPath(os.Stdout, template)

	default:
		// TODO handle this using cobra itself?
		err = errors.Errorf("output format %s not supported", outputFormat)
	}

	if err != nil {
		fmt.Println(err)
	}
}

func getMyAccountID() (string, error) {
	me, err := clientset.API().Account().Get()
	if err != nil {
		return "", errors.Wrap(err, "retrieving account")
	}

	return string(me.ID), nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource",
	Long: `Get a resource

TODO this is a long description`,

	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "output format")
	getCmd.PersistentFlags().BoolVarP(&mineOnly, "mine", "m", false, "only list resources your user owns")
}
