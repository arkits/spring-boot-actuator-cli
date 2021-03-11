package domain

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
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
