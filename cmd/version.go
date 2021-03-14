package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of spring-boot-actuator-cli",
	Run:   printCliVersion,
}

func printCliVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("spring-boot-actuator-cli version %s \n", viper.GetString("application.version"))
	fmt.Println("")
	fmt.Println("Here are some cool links: ")
	fmt.Println("- Project Github: https://github.com/arkits/spring-boot-actuator-cli")
}
