package cmd

import (
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
