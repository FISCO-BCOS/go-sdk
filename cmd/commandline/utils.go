package commandline

import (
	"encoding/json"
	"fmt"
)

const (
	DefaultSuccessMessage string = "success"
)

func ListToJSONStr(list interface{}, name string) (string, error) {
	var mapObject = make(map[string]interface{})
	mapObject[name] = list
	jsonBytes, err := json.Marshal(mapObject)
	if err != nil {
		return "", fmt.Errorf("ListToJsonStr failed, json.Marshal err: %v", err)
	}
	return string(jsonBytes), nil
}
