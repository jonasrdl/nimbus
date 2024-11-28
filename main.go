package main

import "github.com/jonasrdl/nimbus/pkg/nimbus"

// This file demonstrates how to use the nimbus package to log messages.

func main() {
	// Example 1: Using the global logger
	nimbus.SetGlobalLogger(nimbus.DebugLevel)
	nimbus.Info("Global logger: Application started", "version", "1.0.0")
	nimbus.Debug("Global logger: Debugging enabled", "module", "main")
	nimbus.Warn("Global logger: Potential issue detected", "module", "auth")
	nimbus.Error("Global logger: An error occurred", "error", "network timeout")
	// Uncomment Fatal if you want to terminate the program
	// nimbus.Fatal("Global logger: Fatal error, shutting down")

	// Example 2: Using an instance-based logger
	logger := nimbus.New(nimbus.InfoLevel, "text")
	logger.Info("Instance-based logger: Application started", "version", "1.0.0")
	logger.Debug("Instance-based logger: Debugging enabled (this will not be shown)")
	logger.Warn("Instance-based logger: Potential issue detected", "module", "auth")
	logger.Error("Instance-based logger: An error occurred", "error", "disk full")
	// Uncomment Fatal if you want to terminate the program
	// logger.Fatal("Instance-based logger: Fatal error, shutting down")

	// Example 3: Using WithFields to add persistent fields
	loggerWithFields := logger.WithFields(map[string]interface{}{
		"user_id": 1234,
		"session": "abc1234",
	})
	loggerWithFields.Info("User logged in")

	// Example 4: Using JSON formatter
	jsonLogger := nimbus.New(nimbus.InfoLevel, "json")
	//nolint:lll
	// {"level":"INFO","message":"JSON logger: Application","status":"running","timestamp":"2024-11-28 11:22:56","version":"1.0.0"}
	jsonLogger.Info("JSON logger: Application", "version", "1.0.0", "status", "running")
}
