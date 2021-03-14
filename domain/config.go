package domain

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config defines the utility's config
type Config struct {
	Verbose                bool
	ActuatorEndpointPrefix string
	SkipPrettyPrint        bool
	Inventory              []Inventory
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

	viper.BindPFlag("Verbose", cmd.Flags().Lookup("verbose"))
	viper.BindPFlag("SkipPrettyPrint", cmd.Flags().Lookup("skip-pretty-print"))

	if LookupFlagInCmd("actuator-base", cmd) {
		viper.BindPFlag("ActuatorEndpointPrefix", cmd.Flags().Lookup("actuator-base"))
	} else {
		viper.SetDefault("ActuatorEndpointPrefix", "actuator")
	}

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

	// Build out the Inventory
	// From here we attempt to figure what Inventory the user wishes to use

	// If the specific flag is passed, then we use that to build the Inventory
	if LookupFlagInCmd("specific", cmd) {

		// Specific flag's parsing logic...

		// Get the input string
		specific := cmd.Flags().Lookup("specific").Value.String()

		// Multiple name can be pass with a ','
		specificNames := strings.Split(specific, ",")

		var specificInventory []Inventory

		// Filter out the inventory based on the specific names
		for _, name := range specificNames {

			// TODO: this needs to be switched with a hash-map lookup to handle larger inventories
			inventory := getInventoryByName(name)

			if inventory.BaseURL != "" {
				specificInventory = append(specificInventory, inventory)
			}

		}

		CLIConfig.Inventory = specificInventory

	}

	// Check if flags for impromptu definition of an Inventory were passed
	// Assume that if url was passed, then it is an impromptu definition
	if LookupFlagInCmd("url", cmd) {

		VLog("url was set... assuming impromptu definition")

		var impromptuInventory Inventory
		impromptuInventory.Name = "unset"
		impromptuInventory.BaseURL = cmd.Flags().Lookup("url").Value.String()

		if LookupFlagInCmd("auth-header", cmd) {
			impromptuInventory.AuthorizationHeader = cmd.Flags().Lookup("auth-header").Value.String()
		}

		if LookupFlagInCmd("skip-verify-ssl", cmd) {
			impromptuInventory.SkipVerifySSL = true
		}

		var singleInventory []Inventory
		singleInventory = append(singleInventory, impromptuInventory)

		CLIConfig.Inventory = singleInventory

	}

}

// LookupFlagInCmd returns whether a flag is set in the cmd
func LookupFlagInCmd(flagName string, cmd *cobra.Command) bool {

	// I'm not sure about this...
	value := cmd.Flags().Lookup(flagName).Value.String()

	if value == "" {
		return false
	}
	return true
}

func getInventoryByName(name string) Inventory {

	var inventory Inventory

	for _, i := range CLIConfig.Inventory {
		if i.Name == name {
			return i
		}
	}

	return inventory

}
