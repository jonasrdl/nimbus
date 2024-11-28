package formatter

import (
	"encoding/json"
	"fmt"
	"time"
)

// Formatter interface that can be implemented for different log formats.
type Formatter interface {
	Format(level, message string, fields map[string]interface{}) string
}

// TextFormatter formats logs in plain text.
type TextFormatter struct{}

// Format for text logs, which includes timestamp, level, and message.
func (f *TextFormatter) Format(level, message string, fields map[string]interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] %s", timestamp, level, message)

	if len(fields) > 0 {
		logMessage += " - "
		for key, value := range fields {
			logMessage += fmt.Sprintf("%s=%v ", key, value)
		}
	}

	if len(logMessage) > 0 && logMessage[len(logMessage)-1] == ' ' {
		logMessage = logMessage[:len(logMessage)-1]
	}

	return logMessage
}

// JSONFormatter formats logs in JSON.
type JSONFormatter struct{}

// Format for JSON logs, which includes timestamp, level, message, and fields as a JSON object.
func (f *JSONFormatter) Format(level, message string, fields map[string]interface{}) string {
	logData := make(map[string]interface{})
	logData["level"] = level
	logData["message"] = message
	logData["timestamp"] = time.Now().Format("2006-01-02 15:04:05")

	for key, value := range fields {
		logData[key] = value
	}

	logMessage, _ := json.Marshal(logData)
	return string(logMessage)
}
