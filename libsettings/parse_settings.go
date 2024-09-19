package libsettings

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ParseSettings(file string) (Settings, error) {
	var result Settings

	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(jsonFile)

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return Settings{}, err
	}

	return result, nil
}
