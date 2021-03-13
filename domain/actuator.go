package domain

import (
	"fmt"
)

// GenericGetActuatorResponse retrieves data from a generic actuator endpoint and returns the response as a string
// TODO: perform better error handling to top
func GenericGetActuatorResponse(inventory Inventory, endpoint string) (string, error) {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(inventory.BaseURL, "/"+CLIConfig.ActuatorEndpointPrefix+"/"+endpoint)

	return MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, inventory.SkipVerifySSL)

}

// PrintActuatorCustom retrieves data a custom /actuator endpoint and prints it based on the passed params
func PrintActuatorCustom(inventory Inventory, endpoint string) error {

	strResponse, _ := GenericGetActuatorResponse(inventory, endpoint)

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	PrettyPrintJSON(strResponse)

	return nil

}

// PrintActuatorInfo retrieves data from /actuator/info and prints it out
func PrintActuatorInfo(inventory Inventory) error {

	strResponse, _ := GenericGetActuatorResponse(inventory, "info")

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	PrettyPrintJSON(strResponse)

	return nil

}

// PrintActuatorEnv retrieves data from /actuator/env and prints it out
func PrintActuatorEnv(inventory Inventory) error {

	strResponse, _ := GenericGetActuatorResponse(inventory, "env")

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	PrettyPrintActuatorEnvResponse(strResponse)

	return nil

}
