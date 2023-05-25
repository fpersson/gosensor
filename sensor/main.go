package main

import (
	"errors"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/fpersson/gosensor/libsensor"
	libsettings "github.com/fpersson/gosensor/libsettings"
	"golang.org/x/exp/slog"
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
		if exists(file) {
			return file
		}
	}

	return ""
}

func main() {
	loggHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggHandler)

	logger.Info("Start")
	configdir := os.Getenv("CONFIG")

	if configdir == "" {
		configdir = os.Getenv("XDG_DATA_DIRS")
	}

	settingsfile := findConfigFile(configdir)

	conf, err := libsettings.ParseSettings(settingsfile)

	if err != nil {
		logger.Info(err.Error())
		os.Exit(0)
	}

	device := os.Getenv("DEVICE")

	if device == "" {
		device = "/sys/bus/w1/devices/"
	}

	path := device + conf.Sensor + "/w1_slave"

	logger.Info(path)
	interval, err := strconv.ParseInt(conf.Influx.Interval, 10, 0)

	if err != nil {
		interval = 1
	}

	found := true
	posted := true

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				val, err := libsensor.ReadSensor(path)

				if err != nil {
					if found {
						logger.Info(err.Error())
						found = false
						posted = true
					}
				} else {
					found = true
					err = libsensor.Post(conf, val)
					if err != nil {
						if posted {
							logger.Info(err.Error())
							posted = false
							found = true
						}
					}
				}
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	logger.Info(sig.String())
	done <- true
}
