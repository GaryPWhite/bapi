package cmd

import (
	"fmt"

	"github.com/garypwhite/bapi/api"
	"github.com/spf13/cobra"
)

// getAgentsCmd represents the getAgents command
var getAgentsCmd = &cobra.Command{
	Use:   "getAgents",
	Short: "Get a list of Buildkite Agents",
	Long:  `Get a list of Buildkite agents from your configured organization. Will return a long JSON string, recommend parsing with other CLI utils`,
	Run: func(cmd *cobra.Command, args []string) {
		raw, err := api.GetAgentList()
		if err != nil {
			fmt.Printf("Error fetching agents:\n%v", err)
		}
		fmt.Print(raw)
	},
}

func init() {
	rootCmd.AddCommand(getAgentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAgentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAgentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
