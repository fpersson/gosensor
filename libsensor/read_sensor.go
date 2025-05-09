// Package libsensor provides utilities for reading sensor data and posting it
// to an InfluxDB instance. It includes functions for sensor data parsing and
// database interaction.
package libsensor

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadSensor reads the sensor data from the specified file and parses its value.
//
// Parameters:
//   - sensor: The file path to the sensor data.
//
// Returns:
//   - float64: The parsed sensor value.
//   - error: An error object if any issues occur during file reading or parsing.
func ReadSensor(sensor string) (float64, error) {
	content, err := os.ReadFile(sensor)
	if err != nil {
		return 0, fmt.Errorf("failed to read sensor file %s: %w", sensor, err)
	}

	value, err := parseValue(string(content))
	if err != nil {
		return 0, fmt.Errorf("failed to parse sensor data: %w", err)
	}

	return value, nil
}

// parseValue parses the raw sensor data string and extracts the sensor value.
//
// Parameters:
//   - s: The raw sensor data as a string.
//
// Returns:
//   - float64: The extracted and scaled sensor value.
//   - error: An error object if parsing fails.
func parseValue(s string) (float64, error) {
	lines := strings.Split(s, "\n")
	if len(lines) < 2 {
		return 0, fmt.Errorf("unexpected sensor data format: %s", s)
	}

	value := strings.Split(lines[1], "=")
	if len(value) < 2 {
		return 0, fmt.Errorf("missing '=' in sensor data: %s", lines[1])
	}

	floatValue, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert sensor value to float: %w", err)
	}

	return floatValue / 1000, nil
}
