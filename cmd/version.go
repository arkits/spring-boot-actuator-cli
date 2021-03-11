package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of spring-boot-actuator-cli",
	Run:   printCliVersion,
}

func printCliVersion(cmd *cobra.Command, args []string) {
	fmt.Println("spring-boot-actuator-cli v0.0.1")
	fmt.Println("")
	fmt.Println("Here are some cool links: ")
	fmt.Println("- Project Github: https://github.com/arkits/spring-boot-actuator-cli")
}
