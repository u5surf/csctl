package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags / config options
var (
	cfgFile string

	// TODO support names, not just IDs, and resolve appropriately
	organizationID string
	clusterID      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csctl",
	Short: "Command line client for interacting with Containership",
	Long: `TODO

This is a long description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.csctl.yaml)")

	rootCmd.PersistentFlags().StringVar(&organizationID, "organization", "", "Organization")
	// TODO alias and binding not working as expected
	viper.RegisterAlias("organization", "org")
	viper.BindPFlag("organization", rootCmd.Flags().Lookup("organization"))

	rootCmd.PersistentFlags().StringVar(&clusterID, "cluster", "", "Cluster")
	viper.BindPFlag("cluster", rootCmd.Flags().Lookup("cluster"))

	viper.SetDefault("apiBaseURL", "https://api.containership.io")
	viper.SetDefault("provisionBaseURL", "https://provision.containership.io")
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
