package cmd

import (
	"fmt"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	provisionBaseURL = "https://stage-api.containership.io"
)

var (
	orgID   string
	orgName string

	clusterID   string
	clusterName string
)

func getResource(url string) (string, error) {
	token := viper.GetString("token")
	if token == "" {
		return "", errors.New("please provide a token")
	}

	authHeader := fmt.Sprintf("JWT %s", token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).Get(url)
	if err != nil {
		return "", errors.Wrap(err, "error requesting resource")
	}

	return string(resp.Body()), err
}

func getClusters(organizationID string) (string, error) {
	url := fmt.Sprintf("%s/v3/organizations/%s/clusters", provisionBaseURL, organizationID)
	return getResource(url)
}

func getOrganizations() (string, error) {
	url := fmt.Sprintf("%s/v3/organizations", provisionBaseURL)
	return getResource(url)
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resource>",
	Short: "Get a resource",
	Long: `Get a resource

TODO this is a long description`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		switch resource {
		case "cluster", "clusters":
			if orgID == "" && orgName == "" {
				fmt.Println("organization is required")
				return
			}

			resp, err := getClusters(orgID)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(resp)

		case "org", "orgs", "organization", "organizations":
			resp, err := getOrganizations()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(resp)

		default:
			fmt.Println("Error: invalid resource specified: %q", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("output", "o", "json", "Output format")
	getCmd.Flags().StringVar(&orgID, "org", "", "Organization")
	getCmd.Flags().StringVar(&clusterID, "cluster", "", "Cluster")
}
