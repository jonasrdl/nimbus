package main

import "nimbus/pkg/nimbus"

// This file demonstrates how to use the nimbus package to log messages.

func main() {
	logger := nimbus.New(nimbus.InfoLevel)

	logger.Info("Application started", "version", "1.0.0")
	logger.Debug("Debugging data", "key", "value")
	logger.Warn("Potential issue detected", "module", "auth")
	logger.Error("An error occurred", "error", "network timeout")
	logger.Fatal("Fatal error, shutting down")
}
