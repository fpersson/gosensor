package syscmd

import (
	"bufio"
	"os"
	"strings"
)

const OsReleasePath = "/etc/os-release"

func GetOsOsReleaseHTML() (string, error) {
	osinfo, err := ParseOsRelease(OsReleasePath)
	var result string

	if err != nil {
		result = "<b>Hostsystem: </b>Unkown <b>OS version:</b> unkown<br/>"
	}

	result = "<b>Hostsystem: </b>" + osinfo["NAME"] + " <b>OS version:</b> " + osinfo["VERSION_ID"] + "<br/>"
	return result, err
}

func ParseOsRelease(file string) (osrelease map[string]string, err error) {
	var result = make(map[string]string)
	readFile, err := os.Open(file)

	if err != nil {
		return result, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		value := strings.Split(line, "=")
		result[value[0]] = strings.Trim(value[1], "\"")
	}

	return result, nil
}
