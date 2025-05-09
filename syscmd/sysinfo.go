// Package syscmd provides utilities for retrieving and parsing system information.
// It includes functions to extract OS release details and format them for display.
package syscmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// OsReleasePath is the default file path for the OS release information.
const OsReleasePath = "/etc/os-release"

// GetOsOsReleaseHTML retrieves the operating system name and version in HTML format.
//
// Returns:
//   - string: A formatted HTML string containing the OS name and version.
//   - error: An error object if the OS release file cannot be parsed.
func GetOsOsReleaseHTML() (string, error) {
	osinfo, err := ParseOsRelease(OsReleasePath)
	var result string

	if err != nil {
		result = "<b>Hostsystem: </b>Unkown <b>OS version:</b> unkown<br/>"
	}

	result = "<b>Hostsystem: </b>" + osinfo["NAME"] + " <b>OS version:</b> " + osinfo["VERSION_ID"] + "<br/>"
	return result, err
}

// ParseOsRelease parses the OS release file and extracts key-value pairs.
//
// Parameters:
//   - file: The path to the OS release file.
//
// Returns:
//   - map[string]string: A map containing OS release information (e.g., NAME, VERSION_ID).
//   - error: An error object if the file cannot be opened or parsed.
func ParseOsRelease(file string) (osrelease map[string]string, err error) {
	var result = make(map[string]string)
	readFile, err := os.Open(file)

	if err != nil {
		return result, fmt.Errorf("could not open: %s", file)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// OpenSuSE has changed the format of the os-release file
		if strings.Contains(line, "=") {
			value := strings.Split(line, "=")
			result[value[0]] = strings.Trim(value[1], "\"")
		}
	}

	return result, nil
}
