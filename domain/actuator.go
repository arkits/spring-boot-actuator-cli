package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
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

type ActuatorEnvProperties struct {
	ActiveProfiles  []string                     `json:"activeProfiles"`
	PropertySources []ActuatorEnvPropertySources `json:"propertySources"`
}

type ActuatorEnvPropertySources struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

// PrintActuatorEnv retrieves data from /actuator/env and prints it out
func PrintActuatorEnv(inventory Inventory) error {

	strResponse, _ := GenericGetActuatorResponse(inventory, "env")

	if CLIConfig.SkipPrettyPrint {
		fmt.Println(strResponse)
		return nil
	}

	instance := dynamicstruct.ExtendStruct(ActuatorEnvProperties{}).
		Build().
		New()

	err := json.Unmarshal([]byte(strResponse), &instance)
	if err != nil {
		log.Fatal(err)
	}

	reader := dynamicstruct.NewReader(instance)

	// activeProfiles table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetAllowedRowLength(200)
	t.AppendHeader(table.Row{
		"Active Profiles",
	})
	for _, profileName := range reader.GetField("ActiveProfiles").Interface().([]string) {
		t.AppendRows([]table.Row{
			{profileName},
		})
	}
	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()

	// propertySources table
	t.AppendHeader(table.Row{
		"Property Sources",
	})

	for _, propertySources := range reader.GetField("PropertySources").Interface().([]ActuatorEnvPropertySources) {

		// omit a propertySource if there no properties in it
		if len(propertySources.Properties) == 0 {
			continue
		}

		var propertySourceType string
		var applicationConfigFilename string

		// parse out the propertySource.Name
		if strings.HasPrefix(propertySources.Name, "applicationConfig") {

			propertySourceType = "applicationConfig"

			applicationConfigFilenames := strings.Split(propertySources.Name, "file:") // propertySources.Name = "applicationConfig: [file:/data/config/application.yml] (document #3)"

			if len(applicationConfigFilenames) > 0 {

				filenameWithDocumentStr := applicationConfigFilenames[1] // filenameWithDocumentStr="/data/config/application.yml] (document #3)"

				filenameParsed := strings.Split(filenameWithDocumentStr, "]") // spilt the end "] (document #3)"

				applicationConfigFilename = text.WrapSoft(filenameParsed[0], 45)
			}

		} else {
			propertySourceType = propertySources.Name
		}

		// construct the propertySource header
		propertySourceHeaderStr := fmt.Sprintf("sourceType: '%s' \n", propertySourceType)

		// conditionally append the filename
		if applicationConfigFilename != "" {
			propertySourceHeaderStr = propertySourceHeaderStr + fmt.Sprintf("filename: '%s' \n", applicationConfigFilename)
		}

		// append len of the properties
		propertySourceHeaderStr = propertySourceHeaderStr + fmt.Sprintf("len: %v", len(propertySources.Properties))

		t.AppendRows([]table.Row{
			{propertySourceHeaderStr},
		})
		t.AppendSeparator()

		// traverse the property map
		for k, v := range propertySources.Properties {

			for v_k, v_v := range v.(map[string]interface{}) {

				// if the map is element's key is not "value", continue
				if v_k != "value" {
					continue
				}

				var prettyVK string

				switch v_v.(type) {
				case string:
					prettyVK = fmt.Sprintf("%q", v_v)
				default:
					prettyVK = fmt.Sprintf("%v", v_v)
				}

				prettyVK = text.WrapHard(prettyVK, 80)

				t.AppendRows([]table.Row{
					{k, prettyVK},
				})
			}
		}

		t.AppendSeparator()

	}

	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()

	return nil

}
