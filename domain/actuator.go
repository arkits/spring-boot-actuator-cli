package domain

import (
	"fmt"
)

// PrintActuatorInfo retrieves data from /actuator/info and prints it out
func PrintActuatorInfo(inventory Inventory) error {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(inventory.BaseURL, "/actuator/info")

	// Make the HTTP call
	response, _ := MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, inventory.SkipVerifySSL)

	// Print out the good stuff
	fmt.Println("")
	PrettyPrintJSON(response)
	fmt.Println("")

	return nil

}

// PrintActuatorEnv retrieves data from /actuator/env and prints it out
func PrintActuatorEnv(inventory Inventory) error {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(inventory.BaseURL, "/actuator/env")

	// Make the HTTP call
	response, _ := MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, inventory.SkipVerifySSL)

	fmt.Println("")

	// Print out the good stuff
	fmt.Println("")
	PrettyPrintJSON(response)
	fmt.Println("")

	return nil

}
