package domain

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/text"
)

// VLog is a utility function to pretty verbose logging statements based on the passed string and the application's config
func VLog(log string) {
	if CLIConfig.Verbose {
		fmt.Printf(">>> %s \n", log)
	}
}

// ELog is utility function to print error logging statements
func ELog(log string) {
	t := text.FgRed.Sprintf(">>> ERR >>> %s", log)
	fmt.Println(t)
	os.Exit(1)
}

// GetStrBetween returns the trimmed string based on the input, start and end
// eg -
// 		input="class path resource [application.yml]"
// 		start="[" 	end="]"
// 		return="application.yml"
func GetStrBetween(input string, start string, end string) string {
	var toReturn string

	trimmedLSlice := strings.Split(input, start)

	if len(trimmedLSlice) > 0 {
		trimmedLStr := trimmedLSlice[1]

		trimmedRSlice := strings.Split(trimmedLStr, end)

		// assume: if we got till here then it's safe to directly read from index
		toReturn = trimmedRSlice[0]

	}

	return toReturn
}
