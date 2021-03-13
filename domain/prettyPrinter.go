package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
)

// PrettyPrintJSON prints a JSON string in a pretty format using the colorjson library
func PrettyPrintJSON(jsonStr string) {

	colorJsonFormatter := colorjson.NewFormatter()
	colorJsonFormatter.Indent = 2

	// Create an intersting JSON object to marshal in a pretty format
	var obj map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &obj)

	s, _ := colorJsonFormatter.Marshal(obj)
	fmt.Println(string(s))

}

type ActuatorEnvProperties struct {
	ActiveProfiles  []string                     `json:"activeProfiles"`
	PropertySources []ActuatorEnvPropertySources `json:"propertySources"`
}

type ActuatorEnvPropertySources struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

// PrettyPrintActuatorEnvResponse pretty prints the response from /actuator/env
func PrettyPrintActuatorEnvResponse(actuatorEnvResponseStr string) {
	instance := dynamicstruct.ExtendStruct(ActuatorEnvProperties{}).
		Build().
		New()

	err := json.Unmarshal([]byte(actuatorEnvResponseStr), &instance)
	if err != nil {
		log.Fatal(err)
	}

	reader := dynamicstruct.NewReader(instance)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	// activeProfiles table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetAllowedRowLength(200)
	t.AppendHeader(table.Row{
		text.Bold.Sprint("Active Profiles"),
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
		text.Bold.Sprint("Property Sources"),
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight},
		{Number: 2, Align: text.AlignLeft},
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

				applicationConfigFilename = filenameParsed[0]
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

		// bold the header str
		propertySourceHeaderStr = text.Bold.Sprint(propertySourceHeaderStr)

		t.AppendRow(table.Row{propertySourceHeaderStr, propertySourceHeaderStr}, rowConfigAutoMerge)
		t.AppendSeparator()

		// traverse the property map
		for k, v := range propertySources.Properties {

			for v_k, v_v := range v.(map[string]interface{}) {

				// if the map is element's key is not "value", continue
				if v_k != "value" {
					continue
				}

				var prettyVK string

				// Pretty-print based on the type
				switch v_v.(type) {
				case string:
					prettyVK = fmt.Sprintf("%q", v_v)
				default:
					prettyVK = fmt.Sprintf("%v", v_v)
				}

				// Enforce word wrap
				prettyVK = text.WrapHard(prettyVK, 80)

				t.AppendRow(table.Row{k, prettyVK}, rowConfigAutoMerge)

			}
		}

		// Extra padding in the bottom of each propertySource listing... to improve read-ability
		t.AppendRow(table.Row{""})

		t.AppendSeparator()

	}

	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()
}
