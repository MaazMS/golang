# Complete Guide to Error Handling in Go

Go takes a unique approach to error handling that differs significantly from traditional exception-based systems. This guide covers Go's error handling philosophy, patterns, and best practices.

## Table of Contents
1. [Go's Error Handling Philosophy](#gos-error-handling-philosophy)
2. [Error Types and Values](#error-types-and-values)
3. [Basic Error Checking](#basic-error-checking)
4. [Error Handling Patterns](#error-handling-patterns)
5. [Custom Error Types](#custom-error-types)
6. [Panic and Recover](#panic-and-recover)
7. [Logging and Error Reporting](#logging-and-error-reporting)
8. [Best Practices](#best-practices)
9. [Common Patterns](#common-patterns)

## Go's Error Handling Philosophy

### Why Go Doesn't Have Exceptions

Go's designers believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code. It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.

### Go's Approach

Go takes a different approach. For plain error handling, Go's multi-value returns make it easy to report an error without overloading the return value. A canonical error type, coupled with Go's other features, makes error handling pleasant but quite different from that in other languages.

Go also has a couple of built-in functions to signal and recover from truly exceptional conditions. The recovery mechanism is executed only as part of a function's state being torn down after an error, which is sufficient to handle catastrophe but requires no extra control structures and, when used well, can result in clean error-handling code.

### Key Principles

1. **Errors are values**: Errors in Go are just values that can be passed around
2. **Explicit error handling**: Errors must be explicitly checked and handled
3. **No exceptions**: Go doesn't use try-catch-finally patterns
4. **Panic for exceptional cases**: Use panic only for truly exceptional conditions
5. **Recover for cleanup**: Use recover to handle panics and perform cleanup

## Error Types and Values

### The Error Interface

```go
type error interface {
    Error() string
}
```

### Built-in Error Types

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Creating errors using errors.New()
	err1 := errors.New("something went wrong")
	fmt.Println("Error 1:", err1)

	// Creating errors using fmt.Errorf()
	err2 := fmt.Errorf("failed to process %s", "data")
	fmt.Println("Error 2:", err2)

	// Checking if error is nil
	if err1 != nil {
		fmt.Println("Error exists:", err1.Error())
	}
}
```

### Error Values

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

/* Output:
Result: 5
Error: division by zero
*/
```

## Basic Error Checking

### Understanding Error States

#### `if err == nil`
When `err == nil`, it means that an error is a reference type, and it's not pointing to anything. It doesn't exist. There's no error.

#### `if err != nil`
When `err != nil`, it means there is an error. We're pointing to something that exists.

### Basic Error Checking Pattern

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Example: Opening a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Example: Reading from file
	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
}
```

### Multiple Error Checks

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func processFile(filename string) error {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read file info
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	// Process file
	fmt.Printf("Processing file: %s (size: %d bytes)\n", filename, info.Size())
	return nil
}

func main() {
	err := processFile("example.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

## Error Handling Patterns

### 1. Early Return Pattern

```go
package main

import (
	"fmt"
	"os"
)

func readConfig(filename string) (string, error) {
	// Check if file exists
	if _, err := os.Stat(filename); err != nil {
		return "", fmt.Errorf("config file not found: %w", err)
	}

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open config: %w", err)
	}
	defer file.Close()

	// Read file
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		return "", fmt.Errorf("failed to read config: %w", err)
	}

	return string(data[:n]), nil
}

func main() {
	config, err := readConfig("config.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Config: %s\n", config)
}
```

### 2. Error Wrapping

```go
package main

import (
	"errors"
	"fmt"
)

func validateUser(name, email string) error {
	if name == "" {
		return fmt.Errorf("validation failed: %w", errors.New("name is required"))
	}
	if email == "" {
		return fmt.Errorf("validation failed: %w", errors.New("email is required"))
	}
	return nil
}

func createUser(name, email string) error {
	if err := validateUser(name, email); err != nil {
		return fmt.Errorf("user creation failed: %w", err)
	}
	// Create user logic here
	return nil
}

func main() {
	err := createUser("", "user@example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

### 3. Error Aggregation

```go
package main

import (
	"errors"
	"fmt"
	"strings"
)

type MultiError struct {
	Errors []error
}

func (me *MultiError) Error() string {
	var messages []string
	for _, err := range me.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

func (me *MultiError) Add(err error) {
	if err != nil {
		me.Errors = append(me.Errors, err)
	}
}

func (me *MultiError) HasErrors() bool {
	return len(me.Errors) > 0
}

func validateData(data map[string]string) error {
	var multiErr MultiError

	if data["name"] == "" {
		multiErr.Add(errors.New("name is required"))
	}
	if data["email"] == "" {
		multiErr.Add(errors.New("email is required"))
	}
	if data["age"] == "" {
		multiErr.Add(errors.New("age is required"))
	}

	if multiErr.HasErrors() {
		return &multiErr
	}
	return nil
}

func main() {
	data := map[string]string{
		"name":  "",
		"email": "",
		"age":   "25",
	}

	err := validateData(data)
	if err != nil {
		fmt.Printf("Validation errors: %v\n", err)
	}
}
```

## Custom Error Types

### Creating Custom Error Types

```go
package main

import (
	"fmt"
	"time"
)

// Custom error type with additional context
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
	Time    time.Time
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s (value: %v)", 
		ve.Field, ve.Message, ve.Value)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "age cannot be negative",
			Time:    time.Now(),
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "age cannot be greater than 150",
			Time:    time.Now(),
		}
	}
	return nil
}

func main() {
	err := validateAge(-5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

### Error with Context

```go
package main

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	URL        string
	Message    string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d error for %s: %s", 
		he.StatusCode, he.URL, he.Message)
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
			Message:    "request failed",
		}
	}

	// Read response body
	data := make([]byte, 1024)
	n, err := resp.Body.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return data[:n], nil
}

func main() {
	data, err := fetchData("https://httpbin.org/status/404")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Data: %s\n", string(data))
}
```

## Panic and Recover

### Understanding Panic

Panic is used for truly exceptional conditions that should not occur in normal program execution.

```go
package main

import (
	"fmt"
	"os"
)

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to open %s: %v", filename, err))
	}
	return file
}

func main() {
	// This will panic if file doesn't exist
	file := mustOpen("nonexistent.txt")
	defer file.Close()
	fmt.Println("File opened successfully")
}
```

### Using Recover

```go
package main

import (
	"fmt"
	"log"
)

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	// This will cause a panic
	panic("something went wrong")
}

func main() {
	fmt.Println("Before risky function")
	riskyFunction()
	fmt.Println("After risky function")
}
```

### Panic Recovery Pattern

```go
package main

import (
	"fmt"
	"log"
)

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	result = a / b
	return result, nil
}

func main() {
	result, err := safeDivide(10, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %d\n", result)
}
```

## Logging and Error Reporting

### Logging Levels

Go provides different logging levels for different types of messages:

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Basic logging
	fmt.Println("Basic print statement")

	// Logging to standard output
	log.Println("This is a log message")

	// Logging to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("This message goes to the log file")

	// Fatal logging (exits program)
	// log.Fatalln("This will exit the program")

	// Panic logging (can be recovered)
	// log.Panicln("This will panic but can be recovered")
}
```

### Structured Logging

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
	fatal *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatal: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Error(err error) {
	l.error.Printf("Error occurred: %v\n", err)
}

func (l *Logger) Fatal(err error) {
	l.fatal.Printf("Fatal error: %v\n", err)
	os.Exit(1)
}

func main() {
	logger := NewLogger()

	logger.Info("Application started")
	
	// Simulate an error
	err := fmt.Errorf("database connection failed")
	logger.Error(err)

	// Uncomment to see fatal behavior
	// logger.Fatal(fmt.Errorf("critical system failure"))
}
```

## Best Practices

### 1. Always Check Errors

```go
// Good
file, err := os.Open("file.txt")
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
defer file.Close()

// Bad
file, _ := os.Open("file.txt") // Ignoring error
```

### 2. Use Error Wrapping

```go
// Good
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Bad
if err != nil {
    return fmt.Errorf("operation failed: %v", err) // Loses error context
}
```

### 3. Create Meaningful Error Messages

```go
// Good
return fmt.Errorf("failed to connect to database %s: %w", dbName, err)

// Bad
return fmt.Errorf("error: %v", err)
```

### 4. Use Defer for Cleanup

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close() // Always closes file
    
    // Process file
    return nil
}
```

### 5. Don't Panic for Normal Errors

```go
// Good
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Bad
func divide(a, b float64) float64 {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

## Common Patterns

### 1. Retry Pattern

```go
package main

import (
	"fmt"
	"time"
)

func retryOperation(operation func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	return fmt.Errorf("operation failed after %d retries: %w", maxRetries, err)
}

func main() {
	err := retryOperation(func() error {
		// Simulate an operation that might fail
		return fmt.Errorf("operation failed")
	}, 3)
	
	if err != nil {
		fmt.Printf("Final error: %v\n", err)
	}
}
```

### 2. Circuit Breaker Pattern

```go
package main

import (
	"fmt"
	"time"
)

type CircuitBreaker struct {
	failures    int
	maxFailures int
	timeout     time.Duration
	lastFailure time.Time
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

func (cb *CircuitBreaker) Call(operation func() error) error {
	if cb.failures >= cb.maxFailures {
		if time.Since(cb.lastFailure) < cb.timeout {
			return fmt.Errorf("circuit breaker is open")
		}
		cb.failures = 0 // Reset
	}

	err := operation()
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		return err
	}

	cb.failures = 0
	return nil
}

func main() {
	cb := NewCircuitBreaker(3, 5*time.Second)
	
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error {
			return fmt.Errorf("operation failed")
		})
		fmt.Printf("Attempt %d: %v\n", i+1, err)
	}
}
```

### 3. Error Context Pattern

```go
package main

import (
	"fmt"
	"context"
)

type ErrorContext struct {
	Operation string
	Context   map[string]interface{}
	Err       error
}

func (ec *ErrorContext) Error() string {
	return fmt.Sprintf("operation '%s' failed: %v", ec.Operation, ec.Err)
}

func (ec *ErrorContext) Unwrap() error {
	return ec.Err
}

func processWithContext(ctx context.Context, data string) error {
	// Simulate processing
	if data == "" {
		return &ErrorContext{
			Operation: "processWithContext",
			Context: map[string]interface{}{
				"data": data,
				"user": "unknown",
			},
			Err: fmt.Errorf("data cannot be empty"),
		}
	}
	return nil
}

func main() {
	err := processWithContext(context.Background(), "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

## Summary

Go's error handling approach emphasizes:

1. **Explicit Error Handling**: Errors must be explicitly checked and handled
2. **Errors as Values**: Errors are just values that can be passed around
3. **No Exceptions**: Go doesn't use try-catch-finally patterns
4. **Panic for Exceptional Cases**: Use panic only for truly exceptional conditions
5. **Recover for Cleanup**: Use recover to handle panics and perform cleanup

Key principles:
- **Always check errors**: Never ignore errors
- **Use error wrapping**: Preserve error context with `fmt.Errorf` and `%w`
- **Create meaningful errors**: Provide clear, actionable error messages
- **Use defer for cleanup**: Ensure resources are properly cleaned up
- **Don't panic for normal errors**: Use panic only for exceptional conditions

This approach results in more predictable, maintainable code where error handling is explicit and errors are treated as first-class values.
