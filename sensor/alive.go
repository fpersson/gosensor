package main

import (
	"time"

	"log/slog"
)

// alive logs a periodic "alive" message at a given interval using slog.
// It listens for a signal on the `done` channel to terminate the function.
//
// Parameters:
//   - done: A read-only channel (chan bool) used to signal termination of the function.
//   - ticker: A time.Ticker that defines the interval for logging "alive" messages.
//
// Example usage:
//
//	done := make(chan bool)
//	ticker := time.NewTicker(10 * time.Second)
//	go alive(done, *ticker)
//	// ... later ...
//	done <- true
func alive(done <-chan bool, ticker time.Ticker) {
	slog.Info("Calling Alive")
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			slog.Info("alive")
		}
	}
}
