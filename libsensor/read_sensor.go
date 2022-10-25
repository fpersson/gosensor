package libsensor

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadSensor(sensor string) (float64, error) {
	content, err := ioutil.ReadFile(sensor)
	if err != nil {
		return 0, err
	}

	value, err := parse_value(string(content))

	if err != nil {
		return 0, err
	}

	return value, err
}

func parse_value(s string) (float64, error) {
	lines := strings.Split(s, "\n")

	if len(lines) < 2 {
		return 0, nil
	}

	value := strings.Split(lines[1], "=")
	float_value, err := strconv.ParseFloat(value[1], 64)

	if err != nil {
		return 0, err
	}

	return float_value / 1000, nil
}
