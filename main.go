package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func flattenJSON(data map[string]interface{}, prefix string, result map[string]interface{}) {
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		switch v := value.(type) {
		case map[string]interface{}:
			flattenJSON(v, fullKey, result)
		case []interface{}:
			for i, item := range v {
				itemKey := fmt.Sprintf("%s[%d]", fullKey, i)
				if itemMap, ok := item.(map[string]interface{}); ok {
					flattenJSON(itemMap, itemKey, result)
				} else {
					result[itemKey] = item
				}
			}
		default:
			result[fullKey] = value
		}
	}
}

func main() {
	jsonData := ``

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		log.Fatal("error unmarshaling JSON: ", err)
	}

	flatData := make(map[string]interface{})
	flattenJSON(data, "", flatData)

	flatJSON, err := json.MarshalIndent(flatData, "", " ")
	if err != nil {
		log.Fatal("failed to transform flat date into flat json: ", err)
	}

	fmt.Printf("flattened JSON:")
	// for k, v := range flatData {
	// 	fmt.Printf("%s: %v\n", k, v)
	// }

	fmt.Println(string(flatJSON))

	file, err := os.Create("flat_contract.json")
	if err != nil {
		log.Fatal("error creating file ", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(flatJSON))
	if err != nil {
		log.Fatal("failed to write to file ", err)
	}
}
