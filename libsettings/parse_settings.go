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

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return result, err
	}

	json.Unmarshal(byteValue, &result)

	return result, nil
}
