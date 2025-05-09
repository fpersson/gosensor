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

	"log/slog"

	"github.com/fpersson/gosensor/libsensor"
	libsettings "github.com/fpersson/gosensor/libsettings"
	"github.com/fpersson/gosensor/webservice"
	"github.com/fpersson/gosensor/webservice/model"

	"github.com/joho/godotenv"
)

const configfile = "tempsensor/settings.json"

// exists checks if a file or directory exists at the given path.
//
// Parameters:
//   - path: The file or directory path to check.
//
// Returns:
//   - true if the file or directory exists, false otherwise.
func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// findConfigFile searches for the configuration file in the provided paths.
// It returns the first found configuration file path or an empty string if not found.
//
// Parameters:
//   - configpaths: A colon-separated string of paths to search.
//
// Returns:
//   - confile: The path to the found configuration file or an empty string.
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

// main is the entry point of the application. It initializes the logger,
// loads configuration, starts the web service, and periodically reads sensor data.
// It also handles graceful shutdown on receiving termination signals.
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("no .env file found")
	}
	model.HttpDir = os.Getenv("HTTPDIR")
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
					} else {
						logger.Error("Sensor not found", slog.String("path", path))
						os.Exit(1)
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
