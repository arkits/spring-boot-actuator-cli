package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/text"
	"github.com/jedib0t/go-pretty/v6/table"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"golang.org/x/term"
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

func MakeTable() table.Writer {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	// Get window size and set and the allowed row length
	width, height := GetTerminalSize()
	VLog(fmt.Sprintf("[term] width=%v height=%v", width, height))

	t.SetAllowedRowLength(width)

	return t

}

func MakeDynamicStructReader(structToExtent interface{}, jsonStr string) dynamicstruct.Reader {

	instance := dynamicstruct.ExtendStruct(structToExtent).
		Build().
		New()

	err := json.Unmarshal([]byte(jsonStr), &instance)
	if err != nil {
		ELog(fmt.Sprintf("Error in parsing JSON response error='%s'", err.Error()))
	}

	return dynamicstruct.NewReader(instance)

}

func GetTerminalSize() (int, int) {

	// Get window size and set and the allowed row length
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Printf(">>> Caught an error from terminal.GetSize: %s", err.Error())
	}

	return width, height

}
