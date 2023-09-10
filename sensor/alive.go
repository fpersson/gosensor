package main

import (
	"time"

	"golang.org/x/exp/slog"
)

// print "alive" to logs at given interval
// log, which logger to use
// done, readonly chanal
// ticker
func alive(log slog.Logger, done <-chan bool, ticker time.Ticker) {
	log.Info("Calling Alive")
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			log.Info("alive")
		}
	}
}
