package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	//"github.com/olekukonko/tablewriter"

	apitypes "github.com/containership/csctl/cloud/api/types"
	provisiontypes "github.com/containership/csctl/cloud/provision/types"
)

// Flags
var (
	outputFormat string
	mineOnly     bool
)

/*
// TODO hack
func listNodes(orgID string, clusterID string) ([]interface{}, error) {
	path := "/api/v1/nodes"
	nodes := make([]interface{}, 0)
	return nodes, proxyClient.KubernetesGet(orgID, clusterID, path, &nodes)
}

func getNode(orgID string, clusterID string, nodeID string) (*interface{}, error) {
	path := fmt.Sprintf("/api/v1/nodes/%s", nodeID)
	var node interface{}
	return &node, proxyClient.KubernetesGet(orgID, clusterID, path, &node)
}
*/

// TODO this function is beyond terrible
func outputResponse(resp interface{}) {
	switch outputFormat {
	case "", "table":
		// Default
		// TODO just dump raw response (json blob) for now
		fmt.Println(resp)

	case "json":
		j, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Printf("Error formatting JSON: %v\n", err)
			return
		}

		fmt.Println(string(j))

	case "yaml":
		j, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error in intermediate parsing to JSON: %v\n", err)
			return
		}

		y, err := yaml.JSONToYAML([]byte(j))
		if err != nil {
			fmt.Printf("Error converting to YAML: %v\n", err)
			return
		}

		fmt.Println(string(y))

	case "jsonpath":
		fallthrough
	default:
		// TODO
		fmt.Printf("(output format %s not supported)", outputFormat)
	}
}

// TODO hack
func filterByOwner(list interface{}, ownerID apitypes.UUID) ([]interface{}, error) {
	var res []interface{}
	switch l := list.(type) {
	case []apitypes.Organization:
		for _, o := range l {
			if o.OwnerID == ownerID {
				res = append(res, o)
			}
		}
	case []provisiontypes.CKECluster:
		for _, o := range l {
			// TODO typecast hack
			if string(o.OwnerID) == string(ownerID) {
				res = append(res, o)
			}
		}
	default:
		return nil, errors.Errorf("cannot filter by owner on type %T", l)
	}

	return res, nil
}

func filterMine(list interface{}) ([]interface{}, error) {
	me, err := clientset.API().Account().Get()
	if err != nil {
		return nil, errors.Wrap(err, "error getting account")
	}

	return filterByOwner(list, me.ID)
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resource>",
	Short: "Get a resource",
	Long: `Get a resource

TODO this is a long description`,

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		switch resource {
		case "organization", "organizations", "org", "orgs":
			var resp interface{}
			var err error
			if len(args) == 2 {
				id := args[1]
				resp, err = clientset.API().Organizations().Get(id)
			} else {
				resp, err = clientset.API().Organizations().List()
			}

			if err == nil && mineOnly {
				resp, err = filterMine(resp)
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

		case "cluster", "clusters":
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			var resp interface{}
			var err error
			if len(args) == 2 {
				clusterID := args[1]
				resp, err = clientset.Provision().CKEClusters(organizationID).Get(clusterID)
			} else {
				resp, err = clientset.Provision().CKEClusters(organizationID).List()
			}

			// TODO gross
			if err == nil && mineOnly {
				resp, err = filterMine(resp)
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

		case "account", "acct", "user", "usr":
			// Accounts are essentially users
			// A user can only get their own account
			resp, err := clientset.API().Account().Get()
			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

		case "nodepool", "nodepools", "np", "nps":
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			var resp interface{}
			var err error
			if len(args) == 2 {
				nodePoolID := args[1]
				resp, err = clientset.Provision().NodePools(organizationID, clusterID).Get(nodePoolID)
			} else {
				resp, err = clientset.Provision().NodePools(organizationID, clusterID).List()
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

			/*
				case "node", "nodes", "no", "nos":
					if organizationID == "" || clusterID == "" {
						fmt.Println("organization and cluster are required")
						return
					}

					var resp interface{}
					var err error
					if len(args) == 2 {
						nodeID := args[1]
						resp, err = getNode(organizationID, clusterID, nodeID)
					} else {
						resp, err = listNodes(organizationID, clusterID)
					}

					if err != nil {
						fmt.Println(err)
					} else {
						outputResponse(resp)
					}
			*/

		default:
			fmt.Printf("Error: invalid resource specified: %q\n", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "output format")
	getCmd.Flags().BoolVarP(&mineOnly, "mine", "m", false, "only list resources your user owns")
}
