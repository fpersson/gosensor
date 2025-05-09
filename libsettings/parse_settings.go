// Package libsettings provides utilities for managing and parsing configuration settings.
// It includes functionality to read JSON configuration files and define the structure of settings.
package libsettings

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ParseSettings reads a JSON file and unmarshals its content into a Settings struct.
//
// Parameters:
//   - file: The path to the JSON file to be parsed.
//
// Returns:
//   - Settings: The parsed settings object.
//   - error: An error object if any issues occur during file reading or unmarshalling.
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
