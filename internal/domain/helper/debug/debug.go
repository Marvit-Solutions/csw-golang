package debug

import (
	"encoding/json"
)

func TransformData(req interface{}) string {
	jsonData, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}
