package debug

import (
	"encoding/json"
	"fmt"
)

func TransformData(req interface{}) {
	jsonData, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		fmt.Errorf("failed to transform data: %v", err)
	}

	fmt.Println(string(jsonData))
}
