package domain

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintInventoryHeader(inventory Inventory) {
	fmt.Printf(">>> %v \n", inventory.Name)
}

func PrintInventoryFooter(inventory Inventory) {
	fmt.Print("\n")
}

// PrettyPrintJSON prints a JSON string in a pretty format using the colorjson library
func PrettyPrintJSON(jsonStr string) {

	fmt.Println(PrettyJSON(jsonStr))

}

func PrettyJSON(jsonStr string) string {

	colorJsonFormatter := colorjson.NewFormatter()
	colorJsonFormatter.Indent = 2

	// Create an intersting JSON object to marshal in a pretty format
	var obj map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &obj)

	s, _ := colorJsonFormatter.Marshal(obj)

	return string(s)

}

// PrettyPrintActuatorEnvResponse pretty prints the response from /actuator/env
func PrettyPrintActuatorEnvResponse(actuatorEnvResponseStr string) {

	reader := MakeDynamicStructReader(ActuatorEnvProperties{}, actuatorEnvResponseStr)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	// activeProfiles table
	t := MakeTable()

	t.AppendHeader(table.Row{
		text.Bold.Sprint("Active Profiles"),
	})
	for _, profileName := range reader.GetField("ActiveProfiles").Interface().([]string) {
		t.AppendRows([]table.Row{
			{profileName},
		})
	}

	renderAndResetTable(t)

	// propertySources table
	t.AppendHeader(table.Row{
		text.Bold.Sprint("Property Sources"),
	})

	width, _ := GetTerminalSize()
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
			// however the text can get too long and the col widths would not be equal,
			// but equal col widths are required for rowConfigAutoMerge..
			// work around: just print it once in the bigger col on the right
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

	renderAndResetTable(t)

}

func PrettyPrintActuatorLinksResponse(actuatorResponse string) {

	reader := MakeDynamicStructReader(ActuatorLinks{}, actuatorResponse)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := MakeTable()

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

	renderAndResetTable(t)

}

// PrettyPrintActuatorHealthResponse pretty prints the response from /actuator/health
func PrettyPrintActuatorHealthResponse(actuatorResponse string) {

	reader := MakeDynamicStructReader(ActuatorHealthProperties{}, actuatorResponse)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := MakeTable()

	t.AppendHeader(table.Row{
		text.Bold.Sprint("Health"), text.Bold.Sprint("Health"),
	}, rowConfigAutoMerge)

	status := reader.GetField("Status").String()
	if status == "UP" {
		t.AppendRow(table.Row{
			"status", text.FgGreen.Sprint(status),
		})
	} else {
		t.AppendRow(table.Row{
			"status", text.FgRed.Sprint(status),
		})
	}

	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()

}

// PrettyPrintActuatorInfoResponse pretty prints the response from /actuator/info
func PrettyPrintActuatorInfoResponse(actuatorResponse string) {

	reader := MakeDynamicStructReader(ActuatorInfoProperties{}, actuatorResponse)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := MakeTable()

	// Parse Git info
	if reader.HasField("Git") {

		gitInfo := reader.GetField("Git").Interface().(ActuatorInfoGitProperties)

		if gitInfo.Branch == "" {

			VLog("[pp] gitInfo.Branch was empty... Skipping parsing of GitInfo")

		} else {
			t.AppendHeader(table.Row{
				text.Bold.Sprint("Git Info"), text.Bold.Sprint("Git Info"),
			}, rowConfigAutoMerge)

			t.AppendSeparator()

			t.AppendRow(table.Row{
				"branch", gitInfo.Branch,
			}, rowConfigAutoMerge)

			// Parse commit related info in response.git.commit
			for k, v := range gitInfo.Commit {
				if k == "id" {
					switch v.(type) {
					case string:
						// Target's info.git.mode config was set to DEFAULT
						// Sample -
						// "commit": {
						// 	"id": "01dbf9f",
						// 	"time": "2021-03-14 23:30:28+0000"
						// }
						t.AppendRow(table.Row{
							"commit.ID", fmt.Sprintf("%s", v),
						}, rowConfigAutoMerge)

					default:
						// Target's info.git.mode config was set to FULL
						// Sample -
						// "commit": {
						// 	"time": "2021-03-14 23:30:28+0000",
						// 	"message": {
						// 		"full": "dev: wip pretty-printing git info\n",
						// 		"short": "dev: wip pretty-printing git info"
						// 	},
						// 	"id": {
						// 		"describe": "0.0.2-5-g01dbf9f-dirty",
						// 		"abbrev": "01dbf9f",
						// 		"full": "01dbf9f76c23701dbccf44cd4b8e44abd6ec8640"
						// 	},
						// 	"user": {
						// 		"email": "arkits@outlook.com",
						// 		"name": "Archit Khode"
						// 	}
						// },
						for v_k, v_v := range v.(map[string]interface{}) {
							t.AppendRow(table.Row{
								fmt.Sprintf("commit.%v", v_k), fmt.Sprintf("%v", v_v),
							}, rowConfigAutoMerge)
						}

					}
				}
			}

			renderAndResetTable(t)

		}

	}

	t.AppendHeader(table.Row{
		text.Bold.Sprint("Raw /actuator/info Response"),
	}, rowConfigAutoMerge)

	t.AppendRow(table.Row{
		PrettyJSON(actuatorResponse),
	}, rowConfigAutoMerge)

	renderAndResetTable(t)

}

func renderAndResetTable(t table.Writer) {

	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()

}
