package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sba-cli",
		Short: "A CLI utility to work with Spring Boot's Actuator Endpoint",
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
	addCommonFlags(customCmd)

	customCmd.Flags().StringP("endpoint", "E", "", "Endpoint prefix of the custom endpoint")

	// Add the commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(inventoryCmd)
	rootCmd.AddCommand(customCmd)

}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("verbose", "V", false, "Set whether to output verbose log")

	cmd.Flags().StringP("specific", "S", "", "Name of a specific Inventory")

	// Flags for impromptu definition of an Inventory
	cmd.Flags().StringP("baseURL", "U", "", "URL of the target Spring Boot app")
	cmd.Flags().StringP("authorizationHeader", "H", "", "Authorization Header to use when making the HTTP call")
	cmd.Flags().BoolP("skipVerifySSL", "K", false, "Skip verification of SSL")
	cmd.Flags().BoolP("skip-pretty-print", "", false, "Skip any pretty printing")
}
