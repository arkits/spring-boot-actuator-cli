package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of spring-boot-actuator-cli",
	Run:   printCliVersion,
}

func printCliVersion(cmd *cobra.Command, args []string) {
	// About
	fmt.Println(text.Bold.Sprintf("sba-cli version: %s", viper.GetString("application.version")))
	fmt.Println("A CLI utility to work with Spring Boot's Actuator Endpoint")
	fmt.Println("")

	// Build info
	fmt.Println(text.Bold.Sprintf("Build Info:"))
	fmt.Printf("- application.gitTag=%s \n", viper.GetString("application.gitTag"))
	fmt.Printf("- application.gitCommit=%s \n", viper.GetString("application.gitCommit"))
	fmt.Println("")

	// Shilling
	fmt.Println(text.Bold.Sprintf("Here are some cool links:"))
	fmt.Println("- Project Github: https://github.com/arkits/spring-boot-actuator-cli")
}
