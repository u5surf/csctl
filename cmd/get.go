package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	apitypes "github.com/containership/csctl/cloud/api/types"
	provisiontypes "github.com/containership/csctl/cloud/provision/types"
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
	Use:   "get <resource>",
	Short: "Get a resource",
	Long: `Get a resource

TODO this is a long description`,

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		resourceName := args[0]
		switch {
		case resource.Organization().HasAlias(resourceName):
			var resp = make([]apitypes.Organization, 1)
			var err error
			if len(args) == 2 {
				var v *apitypes.Organization
				v, err = clientset.API().Organizations().Get(args[1])
				resp[0] = *v
			} else {
				resp, err = clientset.API().Organizations().List()
			}

			if err != nil {
				fmt.Println(err)
				return
			}

			orgs := resource.NewOrganizations(resp)

			if mineOnly {
				me, err := getMyAccountID()
				if err != nil {
					fmt.Println(err)
					return
				}

				orgs.FilterByOwnerID(me)
			}

			outputResponse(orgs)

		case resource.CKECluster().HasAlias(resourceName):
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			var resp = make([]provisiontypes.CKECluster, 1)
			var err error
			if len(args) == 2 {
				var v *provisiontypes.CKECluster
				v, err = clientset.Provision().CKEClusters(organizationID).Get(args[1])
				resp[0] = *v
			} else {
				resp, err = clientset.Provision().CKEClusters(organizationID).List()
			}

			if err != nil {
				fmt.Println(err)
				return
			}

			clusters := resource.NewCKEClusters(resp)

			if mineOnly {
				me, err := getMyAccountID()
				if err != nil {
					fmt.Println(err)
					return
				}

				clusters.FilterByOwnerID(me)
			}

			outputResponse(clusters)

		case resource.Account().HasAlias(resourceName):
			// A user can only get their own account
			var resp = make([]apitypes.Account, 1)
			v, err := clientset.API().Account().Get()
			resp[0] = *v
			if err != nil {
				fmt.Println(err)
				return
			}

			accounts := resource.NewAccounts(resp)
			outputResponse(accounts)

		case resource.NodePool().HasAlias(resourceName):
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			var resp = make([]provisiontypes.NodePool, 1)
			var err error
			if len(args) == 2 {
				var v *provisiontypes.NodePool
				v, err = clientset.Provision().NodePools(organizationID, clusterID).Get(args[1])
				resp[0] = *v
			} else {
				resp, err = clientset.Provision().NodePools(organizationID, clusterID).List()
			}

			if err != nil {
				fmt.Println(err)
			}

			nps := resource.NewNodePools(resp)
			outputResponse(nps)

		case resource.Plugin().HasAlias(resourceName):
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			var resp = make([]apitypes.Plugin, 1)
			var err error
			if len(args) == 2 {
				var v *apitypes.Plugin
				v, err = clientset.API().Plugins(organizationID, clusterID).Get(args[1])
				resp[0] = *v
			} else {
				resp, err = clientset.API().Plugins(organizationID, clusterID).List()
			}

			if err != nil {
				fmt.Println(err)
			}

			plugs := resource.NewPlugins(resp)
			outputResponse(plugs)

		default:
			fmt.Printf("Error: invalid resource name specified: %q\n", resourceName)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "output format")
	getCmd.Flags().BoolVarP(&mineOnly, "mine", "m", false, "only list resources your user owns")
}
