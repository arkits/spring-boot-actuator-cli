package domain

import (
	"fmt"
)

func GetActuatorInfo(targetBase string) error {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(targetBase, "/actuator/info")
	authorizationHeader := ""

	// Make the HTTP call
	response, _ := MakeHTTPCall("GET", requestURL, authorizationHeader)

	fmt.Println("")

	// Print out the relevant data
	PrettyPrintJSON(response)

	return nil

}

func GetActuatorEnv(targetBase string) error {

	// Setup and validate the params
	requestURL, _ := GenerateRequestURL(targetBase, "/actuator/env")
	authorizationHeader := ""

	// Make the HTTP call
	response, _ := MakeHTTPCall("GET", requestURL, authorizationHeader)

	fmt.Println("")

	// Print out the relevant data
	PrettyPrintJSON(response)

	return nil

}
