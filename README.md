# nimbus

Nimbus is a lightweight and flexible logging library for Go. It offers structured logging, customizable log levels, and the ability to add fields to logs for better traceability.

## Installation

To use Nimbus in your Go project, simply import it:

```go
import "github.com/jonasrdl/nimbus"
```

Then, add it to your project's dependencies:

```bash
go get github.com/jonasrdl/nimbus
```

## Usage

### Create a new logger
```go
package main

import (
	"github.com/jonasrdl/nimbus"
)

func main() {
	// Create a new logger with the desired log level
	logger := nimbus.New(nimbus.InfoLevel)

	// Log messages with different log levels
	logger.Info("This is an info message")
	logger.Debug("This is a debug message") // Won't be logged due to the log level
	logger.Error("An error occurred", "code", 500)
}
```

### Global logger
You can also use a global logger instance for simpler access throughout your application:
```go
package main

import (
	"github.com/jonasrdl/nimbus"
)

func main() {
	// Use the global logger
	nimbus.Info("This is a global info message")
	nimbus.Debug("This is a global debug message")
}
```

### Structured logging
Attach key-value pairs to your log messages for more context:
```go
package main

import (
	"github.com/jonasrdl/nimbus"
)

func main() {
	// Log with fields for structured logging
	nimbus.Info("User logged in", "user_id", 1234, "session", "abc1234")
}
```

### With fields
You can create a new logger instance with additional fields:
```go
package main

import (
	"github.com/jonasrdl/nimbus"
)

func main() {
	// Create a logger with predefined fields
	logger := nimbus.New(nimbus.InfoLevel)
	loggerWithFields := logger.WithFields(map[string]interface{}{
		"app_version": "1.0.0",
	})

	// Log with the additional fields
	loggerWithFields.Info("This log includes version information")
}
```
