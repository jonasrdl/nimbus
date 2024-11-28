package formatter

import (
	"fmt"
	"time"
)

// Format formats a log message with timestamp, level, and message.
func Format(level, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [%s] %s", timestamp, level, message)
}
