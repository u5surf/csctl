package cmd

import (
	"encoding/json"
	"fmt"

	//"github.com/ghodss/yaml"
	"github.com/go-resty/resty"
	//"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/containership/csctl/pkg/cloud/api/types"
)

// Flags
var (
	outputFormat string
)

func getResource(url string, output interface{}) error {
	token := viper.GetString("token")
	if token == "" {
		return errors.New("please provide a token")
	}

	authHeader := fmt.Sprintf("JWT %s", token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(output).
		Get(url)

	if err != nil {
		return errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership Cloud responded with status %d: %s", resp.StatusCode(), resp.Body())
	}

	return nil
}

func listOrganizations() ([]types.Organization, error) {
	url := fmt.Sprintf("%s/v3/organizations", viper.Get("apiBaseURL"))
	orgs := make([]types.Organization, 0)
	return orgs, getResource(url, &orgs)
}

func getOrganization(id string) (*types.Organization, error) {
	url := fmt.Sprintf("%s/v3/organizations/%s", viper.Get("apiBaseURL"), id)
	var org types.Organization
	return &org, getResource(url, &org)
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
			fmt.Printf("Error formatting JSON: %v", err)
			return
		}

		fmt.Println(string(j))

	case "yaml":
		//y, err := yaml.JSONToYAML([]byte(resp))
		//if err != nil {
		//fmt.Println("(Error converting to YAML)")
		//return
		//}

		//fmt.Println(string(y))

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
		case "org", "orgs", "organization", "organizations":
			var resp interface{}
			var err error
			if len(args) == 2 {
				orgId := args[1]
				resp, err = getOrganization(orgId)
			} else {
				resp, err = listOrganizations()
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

			/*
				case "cluster", "clusters":
					if organizationID == "" {
						fmt.Println("organization is required")
						return
					}

					var resp interface{}
					var err error
					if len(args) == 2 {
						clusterId := args[1]
						resp, err = getCluster(orgId, clusterId)
					} else {
						resp, err = listClusters(orgId)
					}

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

							resp, err := getNodePools(organizationID, clusterID)

							if err != nil {
								fmt.Println(err)
							} else {
								outputResponse(resp)
							}
			*/

		default:
			fmt.Println("Error: invalid resource specified: %q", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format")
}
