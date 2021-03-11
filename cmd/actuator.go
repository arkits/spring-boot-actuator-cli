package cmd

import (
	"fmt"

	"github.com/arkits/spring-boot-actuator-cli/domain"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Interface with /actuator/info",
	Run: func(cmd *cobra.Command, args []string) {

		domain.SetupConfig(cmd)

		for _, inventory := range domain.CLIConfig.Inventory {
			domain.PrintActuatorInfo(inventory)
		}

	},
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Interface with /actuator/env",
	Run: func(cmd *cobra.Command, args []string) {

		domain.SetupConfig(cmd)

		for _, inventory := range domain.CLIConfig.Inventory {
			domain.PrintActuatorEnv(inventory)
		}
	},
}

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Interface for custom actuator endpoints",
	Run: func(cmd *cobra.Command, args []string) {

		domain.SetupConfig(cmd)

		if !domain.LookupFlagInCmd("endpoint", cmd) {
			fmt.Println("Endpoint not set!")
			return
		}

		endpoint := cmd.Flags().Lookup("endpoint").Value.String()

		for _, inventory := range domain.CLIConfig.Inventory {
			domain.PrintGenericActuatorResponse(inventory, endpoint)
		}
	},
}
