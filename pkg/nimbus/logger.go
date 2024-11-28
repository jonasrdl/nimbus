package nimbus

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	globalLogger *Logger
	once         sync.Once
)

// Logger is the core logging struct, now with support for structured fields.
type Logger struct {
	level  LogLevel
	fields map[string]interface{}
}

// New creates a new Logger with the specified minimum log level.
func New(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		fields: make(map[string]interface{}), // initialize fields as an empty map
	}
}

// GetGlobalLogger returns the global logger instance.
func GetGlobalLogger() *Logger {
	once.Do(func() {
		globalLogger = New(InfoLevel)
	})
	return globalLogger
}

// SetGlobalLogger allows configuring the global logger instance with a custom level.
func SetGlobalLogger(level LogLevel) {
	once.Do(func() {
		globalLogger = New(level)
	})
	globalLogger.level = level
}

// Log logs a message with the specified level and fields.
func (l *Logger) Log(level LogLevel, message string, fields ...interface{}) {
	if level < l.level {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] %s", timestamp, level.String(), message)

	// Combine additional fields passed to the Log function with the predefined fields in the Logger instance
	allFields := make(map[string]interface{})
	for key, value := range l.fields {
		allFields[key] = value
	}

	// Process the fields passed into the Log method (should be key-value pairs)
	if len(fields) > 0 {
		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				key, ok := fields[i].(string)
				if ok {
					allFields[key] = fields[i+1]
				}
			}
		}
	}

	// Add the fields to the log message
	if len(allFields) > 0 {
		logMessage += " - "
		for key, value := range allFields {
			logMessage += fmt.Sprintf("%s=%v ", key, value)
		}
	}

	// Remove the trailing space
	if len(logMessage) > 0 && logMessage[len(logMessage)-1] == ' ' {
		logMessage = logMessage[:len(logMessage)-1]
	}

	// Print the final log message
	fmt.Println(logMessage)

	if level == FatalLevel {
		os.Exit(1)
	}
}

// WithFields allows adding structured fields to every log message in a logger instance.
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	// Return a new Logger with the original level and the new fields.
	newLogger := *l
	newLogger.fields = fields
	return &newLogger
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.Log(DebugLevel, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.Log(InfoLevel, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.Log(WarnLevel, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.Log(ErrorLevel, msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.Log(FatalLevel, msg, fields...)
}

// Global logger convenience functions

func Debug(msg string, fields ...interface{}) {
	GetGlobalLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	GetGlobalLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	GetGlobalLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	GetGlobalLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	GetGlobalLogger().Fatal(msg, fields...)
}
