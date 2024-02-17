package debug

import (
	"encoding/json"
	"fmt"
)

func TransformData(req interface{}, labels ...string) {
	fmt.Println()
	var label string
	if len(labels) > 0 {
		label = labels[0]
	} else {
		label = "No Label"
	}

	fmt.Println("Type:", label)

	jsonData, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		fmt.Printf("failed to transform data: %v\n", err)
		return
	}

	fmt.Println("Data: ", string(jsonData))
}
