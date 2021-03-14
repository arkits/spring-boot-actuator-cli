package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"golang.org/x/term"
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

	reader := makeDynamicStructReader(ActuatorEnvProperties{}, actuatorEnvResponseStr)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	// activeProfiles table
	t := makeTable()

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

	width, _ := getTerminalSize()
	col1WidthMax := width / 3
	col2WidthMax := width - col1WidthMax

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight, WidthMax: col1WidthMax},
		{Number: 2, Align: text.AlignLeft, WidthMax: col2WidthMax},
	})

	for _, propertySources := range reader.GetField("PropertySources").Interface().([]ActuatorEnvPropertySources) {

		// omit a propertySource if there no properties in it
		if len(propertySources.Properties) == 0 {
			continue
		}

		var propertySourceType string
		var applicationConfigFilename string
		var isKnownPropertySourceType bool

		if strings.HasPrefix(propertySources.Name, "server.ports") {

			propertySourceType = propertySources.Name
			isKnownPropertySourceType = true

		} else if strings.HasPrefix(propertySources.Name, "systemProperties") {

			propertySourceType = propertySources.Name
			isKnownPropertySourceType = true

		} else if strings.HasPrefix(propertySources.Name, "systemEnvironment") {

			propertySourceType = propertySources.Name
			isKnownPropertySourceType = true

		} else if strings.HasPrefix(propertySources.Name, "applicationConfig") {

			// propertySources.Name = "applicationConfig: [file:/data/config/application.yml] (document #3)"

			isKnownPropertySourceType = true

			propertySourceType = "applicationConfig"

			applicationConfigFilename = GetStrBetween(propertySources.Name, "file:", "]")

		} else if strings.HasPrefix(propertySources.Name, "Config resource 'class path resource [application.yml]'") {

			// propertySources.Name = "Config resource 'class path resource [application.yml]' via location 'optional:classpath:/'"

			isKnownPropertySourceType = true

			propertySourceType = "applicationConfig"

			cpResourceFilename := GetStrBetween(propertySources.Name, "[", "]")

			cpRoute := GetStrBetween(propertySources.Name, "optional:", "'")

			applicationConfigFilename = cpRoute + "/" + cpResourceFilename

		} else {
			propertySourceType = propertySources.Name
			isKnownPropertySourceType = false
		}

		// construct the propertySource header
		propertySourceHeaderStr := fmt.Sprintf("sourceType: '%s' \n", propertySourceType)

		// conditionally append the filename
		if applicationConfigFilename != "" {
			propertySourceHeaderStr = propertySourceHeaderStr + text.WrapHard(fmt.Sprintf("filename: '%s'", applicationConfigFilename), col1WidthMax) + "\n"
		}

		// append len of the properties
		propertySourceHeaderStr = propertySourceHeaderStr + fmt.Sprintf("len: %v", len(propertySources.Properties))

		// bold the header str
		propertySourceHeaderStr = text.Bold.Sprint(propertySourceHeaderStr)

		if isKnownPropertySourceType {
			t.AppendRow(table.Row{propertySourceHeaderStr, propertySourceHeaderStr}, rowConfigAutoMerge)
		} else {
			// if its an unknown property then we print it as-is
			// however the text can get too long and the col widths are not equal,
			// causes the rowConfigAutoMerge will fail...
			// work around: just print it once on the the bigger right col
			t.AppendRow(table.Row{"", propertySourceHeaderStr}, rowConfigAutoMerge)

		}

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

type ActuatorLinks struct {
	Links map[string]interface{} `json:"_links"`
}

func PrettyPrintActuatorLinksResponse(actuatorResponse string) {

	reader := makeDynamicStructReader(ActuatorLinks{}, actuatorResponse)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := makeTable()

	t.AppendHeader(table.Row{
		text.Bold.Sprint("Available Actuators"), text.Bold.Sprint("Available Actuators"),
	}, rowConfigAutoMerge)

	t.AppendRow(table.Row{
		text.Bold.Sprint("href"),
		text.Bold.Sprint("templated"),
	})

	t.AppendSeparator()

	// Iterate through each element in _links
	for _, link := range reader.GetField("Links").Interface().(map[string]interface{}) {

		var href string
		var templated string

		for v_k, v_v := range link.(map[string]interface{}) {

			if v_k == "href" {
				href = fmt.Sprintf("%v", v_v)
			}

			if v_k == "templated" {
				templated = fmt.Sprintf("%v", v_v)
			}

		}

		t.AppendRow(table.Row{href, templated})

	}

	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()

}

func makeTable() table.Writer {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	// Get window size and set and the allowed row length
	width, height := getTerminalSize()
	VLog(fmt.Sprintf("width=%v height=%v", width, height))

	t.SetAllowedRowLength(width)

	return t

}

func makeDynamicStructReader(structToExtent interface{}, jsonStr string) dynamicstruct.Reader {

	instance := dynamicstruct.ExtendStruct(structToExtent).
		Build().
		New()

	err := json.Unmarshal([]byte(jsonStr), &instance)
	if err != nil {
		ELog(fmt.Sprintf("Error in parsing JSON response error='%s'", err.Error()))
	}

	return dynamicstruct.NewReader(instance)

}

func getTerminalSize() (int, int) {

	// Get window size and set and the allowed row length
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Printf(">>> Caught an error from terminal.GetSize: %s", err.Error())
	}

	return width, height

}
