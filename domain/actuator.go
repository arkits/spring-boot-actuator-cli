package domain

import (
	"fmt"
	"time"
)

// GetAndPrintKnownActuator contains the logic to hit the right endpoint and print the data based on the passed cmdName
func GetAndPrintKnownActuator(inventory Inventory, cmdName string) error {

	var endpoint string
	if cmdName == "actuator" {
		endpoint = "/"
	} else {
		endpoint = cmdName
	}

	strResponse, err := GetGenericActuatorResponse(inventory, endpoint)
	if err != nil {
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

// GetAndPrintActuatorCustom retrieves data a custom /actuator endpoint and prints it based on the passed params
func GetAndPrintActuatorCustom(inventory Inventory, endpoint string) error {

	strResponse, err := GetGenericActuatorResponse(inventory, endpoint)
	if err != nil {
		return err
	}

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	PrettyPrintJSON(strResponse)

	return nil

}

// GetGenericActuatorResponse retrieves data from a generic actuator endpoint and returns the response as a string
// TODO: perform better error handling to top
func GetGenericActuatorResponse(inventory Inventory, endpoint string) (string, error) {

	// Setup and validate the params
	requestURL, err := GenerateRequestURL(inventory.BaseURL, "/"+CLIConfig.ActuatorEndpointPrefix+"/"+endpoint)
	if err != nil {
		ELog(fmt.Sprintf("Error in GenerateRequestURL error='%s'", err.Error()))
		return "", err
	}

	response, err := MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, "", inventory.SkipVerifySSL)
	if err != nil {
		ELog(fmt.Sprintf("Error in MakeHTTPCall error='%s'", err.Error()))
		return "", err
	}

	defer response.Body.Close()

	VLog(fmt.Sprintf("[response] Proto: %s Status: %s", response.Proto, response.Status))

	responseBodyStr, err := ResponseBodyToStr(response)
	if err != nil {
		ELog(fmt.Sprintf("Error in ResponseBodyToStr error='%s'", err.Error()))
		return "", err
	}

	if response.StatusCode != 200 {
		err := fmt.Errorf("HTTP response from target was not 2XX - response.StatusCode=%v responseBodyStr=%v", response.StatusCode, responseBodyStr)
		ELog(fmt.Sprint(err))
		return "", err
	}

	return responseBodyStr, nil

}

func GetAndPrintActuatorLogs(inventory Inventory) {

	contentLength := 0

	for {

		rangeHeader := fmt.Sprintf("bytes=%v-", contentLength)

		// Setup and validate the params
		requestURL, err := GenerateRequestURL(inventory.BaseURL, "/"+CLIConfig.ActuatorEndpointPrefix+"/"+"logfile")
		if err != nil {
			ELog(fmt.Sprintf("Error in GenerateRequestURL error='%s'", err.Error()))
		}

		response, err := MakeHTTPCall("GET", requestURL, inventory.AuthorizationHeader, rangeHeader, inventory.SkipVerifySSL)
		if err != nil {
			ELog(fmt.Sprintf("Error in MakeHTTPCall error='%s'", err.Error()))
		}

		defer response.Body.Close()

		VLog(fmt.Sprintf("[response] Proto: %s Status: %s", response.Proto, response.Status))

		if response.StatusCode == 206 {

			responseBodyStr, err := ResponseBodyToStr(response)
			if err != nil {
				ELog(fmt.Sprintf("Error in ResponseBodyToStr error='%s'", err.Error()))
			}

			contentLength = contentLength + len(responseBodyStr)
			VLog(fmt.Sprintf("contentLength=%v", contentLength))

			fmt.Print(responseBodyStr)

		}

		// Run once and exit out
		if !CLIConfig.Tail {
			break
		} else {
			response.Body.Close()
			time.Sleep(1 * time.Second)
		}

	}

}
