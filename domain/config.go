package domain

import (
	"fmt"

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
func SetupConfig() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Error reading config file! - %s", err)
	}

	err := viper.Unmarshal(&CLIConfig)
	if err != nil {
		fmt.Errorf("Unable to decode into struct, %v", err)
	}

}
