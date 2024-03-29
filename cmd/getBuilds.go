/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/garypwhite/bapi/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getBuildsCmd represents the getBuilds command
var getBuildsCmd = &cobra.Command{
	Use:   "getBuilds",
	Short: "Get a list of buildkite Builds",
	Long:  `Get a list of Buildkite agents from your configured organization. Will return a long JSON string, recommend parsing with other CLI utils`,
	Run: func(cmd *cobra.Command, args []string) {
		raw, err := api.GetBuildsList()
		if err != nil {
			fmt.Printf("Error fetching Builds:\n%v", err)
		}
		fmt.Print(raw)
	},
}

func init() {
	rootCmd.AddCommand(getBuildsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getBuildsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getBuildsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getBuildsCmd.Flags().StringP("pipeline", "p", "", "Pipeline to scope root command to.")
	getBuildsCmd.Flags().BoolP("all", "a", false, "Include builds that are not running/scheduled. By default this command will only fetch scheduled/running builds.")
	viper.BindPFlags(getBuildsCmd.Flags())
	viper.SetDefault("pipeline", "")
}
