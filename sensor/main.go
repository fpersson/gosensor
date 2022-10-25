package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fpersson/gosensor/libsensor"
	libsettings "github.com/fpersson/gosensor/libsettings"
)

const configfile = "tempsensor/settings.json"

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

///TODO flytta och testa denna funktionen
func findConfigFile(configpaths string) (confile string) {
	v := strings.Split(configpaths, ":")

	for _, element := range v {
		file := element + "/" + configfile
		fmt.Println(file)
		if exists(file) {
			return file
		}
	}

	return ""
}

func main() {
	fmt.Println("Start")
	configdir := os.Getenv("CONFIG")

	if configdir == "" {
		configdir = os.Getenv("XDG_DATA_DIRS")
	}

	settingsfile := findConfigFile(configdir)

	conf, err := libsettings.ParseSettings(settingsfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	device := os.Getenv("DEVICE")

	if device == "" {
		device = "/sys/bus/w1/devices/"
	}

	path := device + conf.Sensor + "/w1_slave"
	interval, err := strconv.ParseInt(conf.Influx.Interval, 10, 0)

	if err != nil {
		interval = 1
	}

	for {
		val, err := libsensor.ReadSensor(path)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		err = libsensor.Post(conf, val)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
