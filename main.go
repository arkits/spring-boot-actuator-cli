package main

import (
	"github.com/arkits/spring-boot-actuator-cli/cmd"
	"github.com/spf13/viper"
)

var (
	version string
)

func main() {

	if version == "" {
		version = "unset"
	}

	viper.Set("application.version", version)

	cmd.Execute()
}
