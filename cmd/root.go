package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sba-cli",
		Short: "A CLI utility to work with Spring Boot's Actuator Endpoint",
	}

	// inventoryFileName string // TODO: implement ability to pass the name/path to inventory file
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	addCommonFlags(inventoryCmd)
	addCommonFlags(actuatorCmd)
	addCommonFlags(infoCmd)
	addCommonFlags(envCmd)
	addCommonFlags(customCmd)

	customCmd.Flags().StringP("endpoint", "E", "", "Endpoint prefix of the custom endpoint")

	// Add the commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(inventoryCmd)
	rootCmd.AddCommand(actuatorCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(customCmd)

}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("verbose", "V", false, "Set whether to output verbose log")

	cmd.Flags().StringP("specific", "S", "", "Name of a specific Inventory")

	// Flags for impromptu definition of an Inventory
	cmd.Flags().StringP("url", "U", "", "URL of the target Spring Boot app")
	cmd.Flags().StringP("auth-header", "H", "", "Authorization Header to use when making the HTTP call")

	// Other Misc Behavior
	cmd.Flags().StringP("actuator-base", "B", "actuator", "Base of the actuator endpoint")
	cmd.Flags().BoolP("skip-verify-ssl", "K", false, "Skip verification of SSL")
	cmd.Flags().BoolP("skip-pretty-print", "", false, "Skip any pretty printing")

}
