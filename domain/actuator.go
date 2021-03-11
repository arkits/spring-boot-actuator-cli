package domain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// PrintActuatorInfo retrieves data from /actuator/info and prints it out
func PrintActuatorInfo(inventory Inventory) error {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(inventory.BaseURL, "/actuator/info")

	// Make the HTTP call
	strResponseJSON, _ := MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, inventory.SkipVerifySSL)

	// Print out the good stuff
	fmt.Println("")
	PrettyPrintJSON(strResponseJSON)
	fmt.Println("")

	return nil

}

// PrintActuatorEnv retrieves data from /actuator/env and prints it out
func PrintActuatorEnv(inventory Inventory) error {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(inventory.BaseURL, "/actuator/env")

	// Make the HTTP call
	strResponseJSON, _ := MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, inventory.SkipVerifySSL)

	// Print out the good stuff
	fmt.Println("")

	var marshalledResponseJSON map[string]interface{}
	if err := json.Unmarshal([]byte(strResponseJSON), &marshalledResponseJSON); err != nil {
		panic(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"key", "value"})
	t.AppendRows([]table.Row{
		{"activeProfiles", marshalledResponseJSON["activeProfiles"]},
	})
	t.Render()

	return nil

}
