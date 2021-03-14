package main

import (
	"fmt"

	"github.com/arkits/spring-boot-actuator-cli/cmd"
	"github.com/spf13/viper"
)

var (
	gitCommit string
	gitTag    string
)

func main() {

	assembleVersion()

	cmd.Execute()
}

func assembleVersion() {

	if gitCommit == "" {
		gitCommit = "__UNSET_COMMIT"
	}

	if gitTag == "" {
		gitTag = "UNSET_TAG"
	}

	applicationVersion := fmt.Sprintf("%s:%s", gitTag, gitCommit[0:7])

	viper.Set("application.version", applicationVersion)
	viper.Set("application.gitCommit", gitCommit)
	viper.Set("application.gitTag", gitTag)

}
