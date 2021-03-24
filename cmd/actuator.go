package cmd

import (
	"fmt"
	"os"

	"github.com/arkits/spring-boot-actuator-cli/domain"
	"github.com/spf13/cobra"
)

var actuatorCmd = &cobra.Command{
	Use:   "actuator",
	Short: "Interface with /actuator - Lists all the available actuator",
	Run:   handleKnownActuatorCmd,
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Interface with /actuator/info",
	Run:   handleKnownActuatorCmd,
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Interface with /actuator/env",
	Run:   handleKnownActuatorCmd,
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Interface with /actuator/health",
	Run:   handleKnownActuatorCmd,
}

var logfileCmd = &cobra.Command{
	Use:   "logfile",
	Short: "Interface with /actuator/logfile",
	Run:   handleActuatorLogfileCmd,
}

func handleKnownActuatorCmd(cmd *cobra.Command, args []string) {

	domain.SetupConfig(cmd)

	for _, inventory := range domain.CLIConfig.Inventory {
		domain.PrintInventoryHeader(inventory)
		domain.GetAndPrintKnownActuator(inventory, cmd.Name())
		domain.PrintInventoryFooter(inventory)
	}

}

var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Interface for custom actuator endpoints",
	Run: func(cmd *cobra.Command, args []string) {

		domain.SetupConfig(cmd)

		if !domain.LookupFlagInCmd("endpoint", cmd) {
			fmt.Fprintf(os.Stderr, ">>> ERROR >>> No endpoint was passed. \n")
			fmt.Fprintf(os.Stderr, "Please define the path ID of the custom actuator endpoint. \n")
			os.Exit(1)
			return
		}

		endpoint := cmd.Flags().Lookup("endpoint").Value.String()

		for _, inventory := range domain.CLIConfig.Inventory {
			domain.PrintInventoryHeader(inventory)
			domain.GetAndPrintActuatorCustom(inventory, endpoint)
			domain.PrintInventoryFooter(inventory)
		}
	},
}

func handleActuatorLogfileCmd(cmd *cobra.Command, args []string) {

	domain.SetupConfig(cmd)

	// currently only support tailing 1 Inventory
	inventory := domain.CLIConfig.Inventory[0]

	domain.GetAndPrintActuatorLogs(inventory)

}
