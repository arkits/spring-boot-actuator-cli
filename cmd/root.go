package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sba-cli",
		Short: "A CLI app to interface with Spring Boot's Actuator Endpoint",
	}

	inventoryFileName string
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	addCommonFlags(infoCmd)
	addCommonFlags(envCmd)
	addCommonFlags(inventoryCmd)

	// Add the commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(inventoryCmd)

}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("verbose", "V", false, "Set whether to output verbose log")

	// Flags for impromptu definition of an Inventory
	cmd.Flags().StringP("baseURL", "U", "", "URL of the target Spring Boot app")
	cmd.Flags().StringP("authorizationHeader", "H", "", "Authorization Header to use when making the HTTP call")
	cmd.Flags().BoolP("skipVerifySSL", "K", false, "Skip verification of SSL")
}
