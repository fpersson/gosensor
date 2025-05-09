package main

import (
	"time"

	"log/slog"
)

// alive logs a periodic "alive" message at a given interval using the provided logger.
// It listens for a signal on the `done` channel to terminate the function.
//
// Parameters:
//   - log: The logger to use for logging messages.
//   - done: A read-only channel used to signal termination of the function.
//   - ticker: A ticker that defines the interval for logging "alive" messages.
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
