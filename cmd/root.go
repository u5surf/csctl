package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/containership/csctl/cloud"
)

// Flags / config options
var (
	cfgFile      string
	debugEnabled bool

	// TODO support names, not just IDs, and resolve appropriately
	organizationID string
	clusterID      string
	nodePoolID     string

	userToken string
)

var (
	clientset cloud.Interface
)

func orgScopedPreRunE(cmd *cobra.Command, args []string) error {
	organizationID = viper.GetString("organization")
	if organizationID == "" {
		return errors.New("please specify an organization via --organization or config file")
	}

	return nil
}

func clusterScopedPreRunE(cmd *cobra.Command, args []string) error {
	if err := orgScopedPreRunE(cmd, args); err != nil {
		return err
	}

	clusterID = viper.GetString("cluster")
	if clusterID == "" {
		return errors.New("please specify a cluster via --cluster or config file")
	}

	return nil
}

func nodePoolScopedPreRunE(cmd *cobra.Command, args []string) error {
	if err := clusterScopedPreRunE(cmd, args); err != nil {
		return err
	}

	nodePoolID = viper.GetString("node-pool")
	if nodePoolID == "" {
		return errors.New("please specify a node pool via --node-pool or config file")
	}

	return nil
}

func bindCommandToOrganizationScope(cmd *cobra.Command, persistent bool) {
	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVar(&organizationID, "organization", "", "organization to use")
	viper.BindPFlag("organization", flagset.Lookup("organization"))
}

func bindCommandToClusterScope(cmd *cobra.Command, persistent bool) {
	bindCommandToOrganizationScope(cmd, persistent)

	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVar(&clusterID, "cluster", "", "cluster to use")
	viper.BindPFlag("cluster", flagset.Lookup("cluster"))
}

func bindCommandToNodePoolScope(cmd *cobra.Command, persistent bool) {
	bindCommandToClusterScope(cmd, persistent)

	var flagset *pflag.FlagSet
	if persistent {
		flagset = cmd.PersistentFlags()
	} else {
		flagset = cmd.Flags()
	}

	flagset.StringVar(&nodePoolID, "node-pool", "", "node pool to use")
	viper.BindPFlag("node-pool", flagset.Lookup("node-pool"))
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csctl",
	Short: "Command line client for interacting with Containership",
	Long: `TODO

This is a long description`,
	SilenceUsage: true,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		userToken = viper.GetString("token")
		if userToken == "" {
			return errors.New("please specify a token in your config file")
		}

		var err error
		clientset, err = cloud.New(cloud.Config{
			Token:            userToken,
			APIBaseURL:       viper.GetString("apiBaseURL"),
			ProvisionBaseURL: viper.GetString("provisionBaseURL"),
			DebugEnabled:     debugEnabled,
		})

		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.containership/csctl.yaml)")

	rootCmd.PersistentFlags().BoolVar(&debugEnabled, "debug", false, "enable/disable debug mode (trace all HTTP requests)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		containershipDir := path.Join(home, ".containership")

		// Search config in ~/.containership directory with name "csctl.yaml"
		viper.AddConfigPath(containershipDir)
		// Note that function expects the extension to be omitted
		viper.SetConfigName("csctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// TODO debug logging. In the meantime, this is commented out in order
		// to avoid extraneous output that breaks easy piping
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
