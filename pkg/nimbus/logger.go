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

// Logger is the core logging struct.
type Logger struct {
	level LogLevel
}

// New creates a new Logger with the specified minimum log level.
func New(level LogLevel) *Logger {
	return &Logger{level: level}
}

// GetGlobalLogger returns the global logger instance.
// It initializes the logger once with the default InfoLevel.
func GetGlobalLogger() *Logger {
	once.Do(func() {
		globalLogger = New(InfoLevel)
	})
	return globalLogger
}

// SetGlobalLogger allows configuring the global logger instance with a custom log level.
func SetGlobalLogger(level LogLevel) {
	once.Do(func() {
		globalLogger = New(level)
	})
	globalLogger.level = level
}

// Log logs a message with the specified level.
func (l *Logger) Log(level LogLevel, message string, fields ...interface{}) {
	if level < l.level {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s", timestamp, level.String(), message)

	if len(fields) > 0 {
		fmt.Printf(" - ")
		for i := 0; i < len(fields); i += 2 {
			if i+1 < len(fields) {
				fmt.Printf("%s=%v ", fields[i], fields[i+1])
			}
		}
	}
	fmt.Println()

	if level == FatalLevel {
		os.Exit(1)
	}
}

// Convenience methods for each log level.

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
