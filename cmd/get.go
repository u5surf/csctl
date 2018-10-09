package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	//"github.com/olekukonko/tablewriter"
	//"github.com/pkg/errors"

	provisiontypes "github.com/containership/csctl/pkg/cloud/api/provision/types"
	apitypes "github.com/containership/csctl/pkg/cloud/api/types"
)

// Flags
var (
	outputFormat string
)

func listOrganizations() ([]apitypes.Organization, error) {
	path := "/v3/organizations"
	orgs := make([]apitypes.Organization, 0)
	return orgs, apiClient.GetResource(path, &orgs)
}

func getOrganization(id string) (*apitypes.Organization, error) {
	path := fmt.Sprintf("/v3/organizations/%s", id)
	var org apitypes.Organization
	return &org, apiClient.GetResource(path, &org)
}

func listClusters(orgID string) ([]provisiontypes.CKECluster, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters", orgID)
	clusters := make([]provisiontypes.CKECluster, 0)
	return clusters, provisionClient.GetResource(path, &clusters)
}

func getCluster(orgID string, clusterID string) (*provisiontypes.CKECluster, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", orgID, clusterID)
	var cluster provisiontypes.CKECluster
	return &cluster, provisionClient.GetResource(path, &cluster)
}

func getAccount() (*apitypes.Account, error) {
	path := "/v3/account"
	var account apitypes.Account
	return &account, apiClient.GetResource(path, &account)
}

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
			fmt.Println("(Error in intermediate parsing to JSON: %v\n", err)
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
				orgID := args[1]
				resp, err = getOrganization(orgID)
			} else {
				resp, err = listOrganizations()
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
				resp, err = getCluster(organizationID, clusterID)
			} else {
				resp, err = listClusters(organizationID)
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

		case "account", "acct", "user", "usr":
			// Accounts are essentially users
			// A user can only get their own account
			resp, err := getAccount()
			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

			/*
				case "nodepool", "nodepools", "np", "nps":
					if organizationID == "" || clusterID == "" {
						fmt.Println("organization and cluster are required")
						return
					}

					resp, err := getNodePools(organizationID, clusterID)

					if err != nil {
						fmt.Println(err)
					} else {
						outputResponse(resp)
					}
			*/

		default:
			fmt.Println("Error: invalid resource specified: %q\n", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format")
}
