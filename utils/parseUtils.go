package utils

import "encoding/json"

func ParseStructToMap(data any) (map[string]interface{}, error) {

	var result map[string]interface{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(jsonData, &result)

	return result, nil
}
