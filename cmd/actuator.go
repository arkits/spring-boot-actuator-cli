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

		domain.SetupConfig()

		// targetBase := cmd.Flags().Lookup("target").Value.String()
		// viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))

		// domain.GetActuatorInfo(targetBase)
	},
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Interface with /actuator/env",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("")

		targetBase := cmd.Flags().Lookup("target").Value.String()

		domain.GetActuatorInfo(targetBase)
	},
}
