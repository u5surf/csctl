package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// TODO viper dependency should not be needed here
	"github.com/spf13/viper"

	"github.com/containership/csctl/pkg/kubeconfig"
)

// Flags
var (
	filename string
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a resource",

	Args: cobra.ExactArgs(1),

	PreRunE: clusterScopedPreRunE,

	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		switch resource {
		case "kubeconfig", "kubecfg":
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			// TODO do this better once proxy client is in place; see issue #7
			proxyBaseURL := viper.GetString("proxyBaseURL")
			serverAddress := fmt.Sprintf("%s/v3/organizations/%s/clusters/%s/k8sapi/proxy",
				proxyBaseURL, organizationID, clusterID)

			account, err := clientset.API().Account().Get()
			if err != nil {
				fmt.Println(err)
				return
			}

			cluster, err := clientset.API().Clusters(organizationID).Get(clusterID)
			if err != nil {
				fmt.Println(err)
				return
			}

			// TODO error handling
			// TODO UUID typecasting
			cfg := kubeconfig.New(&kubeconfig.Config{
				ServerAddress: serverAddress,
				ClusterID:     string(cluster.ID),
				UserID:        string(account.ID),
				Token:         userToken,
			})

			w := os.Stdout
			if filename != "" {
				w, err = os.Create(filename)
				if err != nil {
					fmt.Println(err)
				}
				defer w.Close()
			}

			// TODO implement merging into ~/.kube/config, which should be the new default
			// (instead of stdout)
			err = kubeconfig.Write(cfg, w)
			if err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Printf("Error: invalid resource specified: %q\n", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	bindCommandToClusterScope(exportCmd, false)

	exportCmd.Flags().StringVarP(&filename, "filename", "f", "", "output kubeconfig to file (default is stdout)")
}
