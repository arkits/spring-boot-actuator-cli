package domain

import (
	"fmt"
)

// PrintKnownActuator contains the logic to hit the right endpoint and print the data based on the passed cmdName
func PrintKnownActuator(inventory Inventory, cmdName string) error {

	var endpoint string
	if cmdName == "actuator" {
		endpoint = "/"
	} else {
		endpoint = cmdName
	}

	strResponse, err := GenericGetActuatorResponse(inventory, endpoint)
	if err != nil {
		ELog(fmt.Sprintf("Error in GenericGetActuatorResponse error='%s'", err.Error()))
		return err
	}

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	switch cmdName {
	case "env":
		PrettyPrintActuatorEnvResponse(strResponse)
	case "actuator":
		PrettyPrintActuatorLinksResponse(strResponse)
	case "health":
		PrettyPrintActuatorHealthResponse(strResponse)
	case "info":
		PrettyPrintActuatorInfoResponse(strResponse)
	default:
		// Haven't added custom pretty-printing support yet... just print as pretty JSON
		PrettyPrintJSON(strResponse)
	}

	return nil

}

// PrintActuatorCustom retrieves data a custom /actuator endpoint and prints it based on the passed params
func PrintActuatorCustom(inventory Inventory, endpoint string) error {

	strResponse, err := GenericGetActuatorResponse(inventory, endpoint)
	if err != nil {
		ELog(fmt.Sprintf("Error in GenericGetActuatorResponse error='%s'", err.Error()))
		return err
	}

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	PrettyPrintJSON(strResponse)

	return nil

}

// GenericGetActuatorResponse retrieves data from a generic actuator endpoint and returns the response as a string
// TODO: perform better error handling to top
func GenericGetActuatorResponse(inventory Inventory, endpoint string) (string, error) {

	// Setup and validate the params
	requestURL, err := GenerateRequestURL(inventory.BaseURL, "/"+CLIConfig.ActuatorEndpointPrefix+"/"+endpoint)
	if err != nil {
		ELog(fmt.Sprintf("Error in GenerateRequestURL error='%s'", err.Error()))
		return "", err
	}

	return MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, inventory.SkipVerifySSL)

}
