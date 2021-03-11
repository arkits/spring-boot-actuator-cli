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

	// Add the commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(inventoryCmd)

}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("target", "T", "", "URI of the target Spring Boot app")
	cmd.Flags().BoolP("verbose", "v", false, "Set whether to output verbose log")
}
