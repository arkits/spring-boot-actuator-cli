package domain

import "fmt"

func VLog(text string) {
	if CLIConfig.Verbose {
		fmt.Printf(">>> %s \n", text)
	}
}
