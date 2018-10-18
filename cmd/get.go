package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"k8s.io/client-go/util/jsonpath"

	apitypes "github.com/containership/csctl/cloud/api/types"
	provisiontypes "github.com/containership/csctl/cloud/provision/types"
)

// Flags
var (
	outputFormat string
	mineOnly     bool
)

// TODO this function is beyond terrible
//   - Make an OutputFormatter type
func outputResponse(resp interface{}) {
	switch {
	case outputFormat == "" || outputFormat == "table":
		// Default
		// TODO just dump raw response (json blob) for now
		// TODO this doesn't actually work atm
		fmt.Println(resp)

	case outputFormat == "json":
		j, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Printf("Error formatting JSON: %v\n", err)
			return
		}

		fmt.Println(string(j))

	case outputFormat == "yaml":
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

	case strings.HasPrefix(outputFormat, "jsonpath"):
		fields := strings.SplitN(outputFormat, "=", 2)
		if len(fields) != 2 {
			fmt.Println("Please specify jsonpath using -ojsonpath=<path>")
			return
		}

		template := fields[1]
		outputJSONPath(template, resp)

	case outputFormat == "id":
		// TODO validate that this field is valid for given resource
		outputJSONPath("{.id}", resp)

	default:
		// TODO handle this using cobra itself?
		fmt.Printf("Error: output format %s not supported", outputFormat)
	}
}

// outputJSONPath parses the data provided for the given jsonpath template string
// and outputs the result to stdout or outputs an error if one occurs
func outputJSONPath(template string, data interface{}) {
	jp := jsonpath.New("")
	err := jp.Parse(template)
	if err != nil {
		fmt.Printf("Error parsing jsonpath: %s", err)
		return
	}

	err = jp.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("Error executing jsonpath: %s", err)
	}
}

// filterByOwner takes a list of resources and filters it by owner ID.
// TODO this is super hacky atm
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

// filterMine takes a list of resources and filters it by owner ID of the authorized user.
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
		// TODO make a Command type that knows about aliases
		// TODO lots of de-duplication to be done
		// TODO accept names as well as IDs
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
				id := args[1]
				resp, err = clientset.Provision().CKEClusters(organizationID).Get(id)
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
				id := args[1]
				resp, err = clientset.Provision().NodePools(organizationID, clusterID).Get(id)
			} else {
				resp, err = clientset.Provision().NodePools(organizationID, clusterID).List()
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

		case "plugin", "plugins", "plug", "plugs", "plgn", "plgns":
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			var resp interface{}
			var err error
			if len(args) == 2 {
				id := args[1]
				resp, err = clientset.API().Plugins(organizationID, clusterID).Get(id)
			} else {
				resp, err = clientset.API().Plugins(organizationID, clusterID).List()
			}

			if err != nil {
				fmt.Println(err)
			} else {
				outputResponse(resp)
			}

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
