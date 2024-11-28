package nimbus

import (
	"fmt"
	"os"
	"sync"

	"github.com/jonasrdl/nimbus/internal/formatter"
)

var (
	globalLogger *Logger
	once         sync.Once
)

// Logger is the core logging struct, now with support for structured fields.
type Logger struct {
	level     LogLevel
	fields    map[string]interface{}
	formatter formatter.Formatter
	logFile   *os.File
}

// New creates a new Logger with the specified minimum log level and format type.
func New(level LogLevel, formatType, logFilePath string) *Logger {
	var formatterInstance formatter.Formatter
	if formatType == "json" {
		formatterInstance = &formatter.JSONFormatter{}
	} else {
		formatterInstance = &formatter.TextFormatter{}
	}

	var logFile *os.File
	if logFilePath != "" {
		var err error
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
		}
	}

	return &Logger{
		level:     level,
		fields:    make(map[string]interface{}), // initialize fields as an empty map
		formatter: formatterInstance,
		logFile:   logFile,
	}
}

// GetGlobalLogger returns the global logger instance.
func GetGlobalLogger() *Logger {
	once.Do(func() {
		globalLogger = New(InfoLevel, "text", "")
	})
	return globalLogger
}

// SetGlobalLogger allows configuring the global logger instance with a custom level.
func SetGlobalLogger(level LogLevel) {
	once.Do(func() {
		globalLogger = New(level, "text", "")
	})
	globalLogger.level = level
}

// Log logs a message with the specified level and fields.
func (l *Logger) Log(level LogLevel, message string, fields ...interface{}) {
	if level < l.level {
		return
	}

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

	// Use the formatter to format the log message
	logMessage := l.formatter.Format(level.String(), message, allFields)

	// Print the formatted log message
	fmt.Println(logMessage)

	if l.logFile != nil {
		_, err := l.logFile.WriteString(logMessage + "\n")
		if err != nil {
			fmt.Printf("Error writing to log file: %v\n", err)
		}
	}

	// If the level is Fatal, exit the program
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
