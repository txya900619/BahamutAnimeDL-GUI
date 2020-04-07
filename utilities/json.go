package utilities

import (
	"encoding/json"
	"log"
)

func ToJson(anyStruct interface{}) string {
	jsonStruct, err := json.Marshal(anyStruct)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonStruct)
}
