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
	Tail                   bool
	Inventory              []Inventory
}

// Inventory defines properties related to single Spring Boot application
type Inventory struct {
	Name                string
	BaseURL             string
	AuthorizationHeader string
	SkipVerifySSL       bool
	Tags                []string
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

	if LookupFlagInCmd("tail", cmd) {
		viper.BindPFlag("Tail", cmd.Flags().Lookup("tail"))
	} else {
		viper.SetDefault("Tail", false)
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

	// If the specific flag is passed, then we use that to build the Inventory
	if LookupFlagInCmd("tag", cmd) {

		// Tag flag's parsing logic...
		// Get the input string
		inputStr := cmd.Flags().Lookup("tag").Value.String()

		// Multiple name can be pass with a ','
		taqs := strings.Split(inputStr, ",")

		var specificInventory []Inventory

		// Iterate through each tag
		for _, tag := range taqs {

			// Get all the inventories that match the tag
			inventories := getAllInventoriesByTag(tag)

			specificInventory = mergeInventories(specificInventory, inventories)

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
	exists := cmd.Flags().Lookup(flagName)
	if exists == nil {
		return false
	}

	value := cmd.Flags().Lookup(flagName).Value.String()
	if value == "" {
		return false
	}

	return true
}

// getInventoryByName is a utility function to retrive an Inventory based on the Name
func getInventoryByName(name string) Inventory {

	var inventory Inventory

	for _, i := range CLIConfig.Inventory {
		if i.Name == name {
			return i
		}
	}

	return inventory

}

// getAllInventoriesByTag is a utility function to retrive an Inventory based on a single Tag
func getAllInventoriesByTag(targetTag string) []Inventory {

	var inventories []Inventory

	// Iterate through the Inventories
	for _, i := range CLIConfig.Inventory {

		// Iterate through the each tag in a single Inventory
		for _, tag := range i.Tags {

			if tag == targetTag {
				inventories = append(inventories, i)
				break
			}

		}
	}

	return inventories

}

// mergeInventories is a utility function to merge 2 Slices of Inventories without adding duplicates
func mergeInventories(oldInventory []Inventory, newInventory []Inventory) []Inventory {

	mergedInventory := oldInventory

	for _, newI := range newInventory {

		for _, oldI := range newInventory {

			// assumption that Name will be unique
			if oldI.Name != newI.Name {
				mergedInventory = append(mergedInventory, newI)
				break
			}

		}

	}

	return mergedInventory

}
