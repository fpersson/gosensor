package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseSettings(file string) (Settings, error) {
	var result Settings

	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	json.Unmarshal(byteValue, &result)

	return result, nil
}
