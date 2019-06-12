package if_json_has

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


func FileRead (src, projectName string) bool{
	jsonFile, err := os.Open(src)
	if err != nil {
		log.Fatalln("os.Open()", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var packageJSON map[string]interface{}

	err = json.Unmarshal([]byte(byteValue), &packageJSON)

	if err != nil {
		log.Fatalln("json.Unmarshal()", err)
	}
	if dependencies, ok  := packageJSON["dependencies"].(map[string]interface{}); ok {
		for k := range dependencies {
			if k == projectName {
				return true
			}
		}
	} else {
		return ok
	}

	return false
}