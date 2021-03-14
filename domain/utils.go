package domain

import (
	"fmt"
	"strings"
)

func VLog(text string) {
	if CLIConfig.Verbose {
		fmt.Printf(">>> %s \n", text)
	}
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
