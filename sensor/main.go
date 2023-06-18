package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/fpersson/gosensor/libsensor"
	libsettings "github.com/fpersson/gosensor/libsettings"
	"github.com/fpersson/gosensor/webservice"
	"github.com/fpersson/gosensor/webservice/model"
	"golang.org/x/exp/slog"
)

const configfile = "tempsensor/settings.json"

// /TODO Move this
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// /TODO Move this
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
	model.LogDir = os.Getenv("LOGDIR")
	loggHandler := slog.NewTextHandler(os.Stdout, nil)

	if model.LogDir != "" {
		f, err := os.OpenFile(model.LogDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("wtf")
		}
		loggHandler = slog.NewTextHandler(f, nil)
	} else {
		fmt.Println("Not using file logger")
	}

	logger := slog.New(loggHandler)

	logger.Info("Start")
	configdir := os.Getenv("CONFIG")

	if configdir == "" {
		configdir = os.Getenv("XDG_DATA_DIRS")
	}

	model.SettingsPath = findConfigFile(configdir)

	conf, err := libsettings.ParseSettings(model.SettingsPath)

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

	myWebservice := webservice.NewWebService(logger)

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
	go myWebservice.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	logger.Info(sig.String())
	done <- true
}
