package cmd

import (
	"fmt"

	"github.com/arkits/spring-boot-actuator-cli/domain"
	"github.com/spf13/cobra"
)

var inventoryCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Prints the parsed Inventory",
	Run: func(cmd *cobra.Command, args []string) {

		domain.SetupConfig(cmd)

		fmt.Println(">>> Parsed Inventory")

		for _, inventory := range domain.CLIConfig.Inventory {

			fmt.Printf("- Name: %s \n", inventory.Name)
			fmt.Printf("  BaseURL: %s \n", inventory.BaseURL)
			fmt.Printf("  AuthorizationHeader: %s \n", inventory.AuthorizationHeader)
			fmt.Printf("  SkipVerifySSL: %v \n", inventory.SkipVerifySSL)
			fmt.Printf("  Tags: %v \n", inventory.Tags)

			fmt.Println("")

		}

	},
}
