package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bapi",
	Short: "A Buildkite API CLI client",
	Long: `A utility for getting useful information from the Buildktite API, at the command line.
	Created to serve purpose in fetching specific queries commonly used in operation and maintainence of Buildkite API.

	Examples:

	getAgents 	  - return all agents currently active for $BAPI_ORGANIZATION
	jobsFromAgent - return recent jobs from $BAPI_AGENT
	stopAgent     - stop $BAPI_AGENT after job finishes
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := valiateInitalFlags(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func valiateInitalFlags() error {
	if !viper.IsSet("organization") {
		return errors.New("BAPI_ORGANIZATION not found in config, env, or command line")
	}
	if !viper.IsSet("token") {
		return errors.New("BAPI_TOKEN not set in config, env, or command line")
	}
	return nil
}

func init() {
	initConfig()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bapi.yaml)")
	rootCmd.PersistentFlags().StringP("organization", "o", "", "organization slug from buildkite")
	rootCmd.PersistentFlags().StringP("params", "p", "", "parameters for the request (filters)")
	rootCmd.PersistentFlags().String("token", "", "Rest API token")

	viper.BindPFlag("organization", rootCmd.PersistentFlags().Lookup("organization"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".bapi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bapi")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.SetEnvPrefix("BAPI")
	viper.AutomaticEnv() // read in environment variables that match
}
