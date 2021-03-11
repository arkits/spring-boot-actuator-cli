package domain

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config defines the utility's config
type Config struct {
	Verbose   bool        `yaml:"verbose"`
	Inventory []Inventory `yaml:"inventory"`
}

// Inventory defines properties related to single Spring Boot application
type Inventory struct {
	Name                string
	BaseURL             string
	AuthorizationHeader string
	SkipVerifySSL       bool
}

// CLIConfig stores the parsed Config
var CLIConfig Config

// SetupConfig sets up the Config based on the config file and any passed params
func SetupConfig(cmd *cobra.Command) {

	viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Error reading config file - %v", err)
	}

	err := viper.Unmarshal(&CLIConfig)
	if err != nil {
		fmt.Errorf("Unable to decode into struct - %v", err)
	}

	// Check if flags for impromptu definition of an Inventory were passed
	// Assume that if baseURL was passed, then it is an impromptu definition
	if lookupFlagInCmd("baseURL", cmd) {

		fmt.Println("baseURL was there")

		var impromptuInventory Inventory
		impromptuInventory.Name = ""
		impromptuInventory.BaseURL = cmd.Flags().Lookup("baseURL").Value.String()

		if lookupFlagInCmd("authorizationHeader", cmd) {
			impromptuInventory.AuthorizationHeader = cmd.Flags().Lookup("authorizationHeader").Value.String()
		}

		fmt.Printf("skipVerifySSL=%v", cmd.Flags().Lookup("skipVerifySSL").Value.String())

		if lookupFlagInCmd("skipVerifySSL", cmd) {
			// impromptuInventory.SkipVerifySSL = false
		}

		var singleInventory []Inventory
		singleInventory = append(singleInventory, impromptuInventory)

		CLIConfig.Inventory = singleInventory

	}

}

func lookupFlagInCmd(flagName string, cmd *cobra.Command) bool {

	// I'm not sure about this...
	value := cmd.Flags().Lookup(flagName).Value.String()

	if value == "" {
		return false
	}
	return true
}
