# Go Defer Statements - Complete Guide

## Overview

The `defer` keyword in Go is used to delay the execution of a statement, function, method, or anonymous function until the surrounding function returns. This is particularly useful for cleanup operations, resource management, and ensuring that certain code executes regardless of how the function exits.

## Key Features

1. **Delayed execution**: Statements are executed when the surrounding function returns
2. **LIFO order**: Multiple defer statements are executed in Last In, First Out order
3. **Guaranteed execution**: Defer statements execute even if the function panics
4. **Resource cleanup**: Perfect for closing files, releasing locks, etc.
5. **Flexible usage**: Can defer any function call, method call, or statement

## Why Use Defer?

### Benefits
- **Resource cleanup**: Automatically close files, connections, etc.
- **Error handling**: Ensure cleanup happens even on errors
- **Code clarity**: Keep cleanup code close to resource acquisition
- **Panic safety**: Cleanup happens even if function panics
- **Debugging**: Add logging or tracing statements

### When to Use Defer
- File operations
- Database connections
- Mutex unlocking
- Resource cleanup
- Logging and debugging
- Panic recovery

## Basic Syntax

### Defer Statement
```go
defer statement
defer function()
defer method()
defer func() { /* anonymous function */ }()
```

## Basic Examples

### Example 1: Simple Defer Usage

```go
package main

import "fmt"

func main() {
    fmt.Println("Starting main function")
    defer foo()
    bar()
    fmt.Println("Ending main function")
}

func foo() {
    fmt.Println("Deferred function foo()")
}

func bar() {
    fmt.Println("Function bar()")
}
```

**Output:**
```
Starting main function
Function bar()
Ending main function
Deferred function foo()
```

### Example 2: Multiple Defer Statements (LIFO Order)

```go
package main

import "fmt"

func main() {
    fmt.Println("Starting main function")
    
    defer fmt.Println("Defer 1")
    defer fmt.Println("Defer 2")
    defer fmt.Println("Defer 3")
    
    fmt.Println("Middle of main function")
    fmt.Println("Ending main function")
}
```

**Output:**
```
Starting main function
Middle of main function
Ending main function
Defer 3
Defer 2
Defer 1
```

### Example 3: Defer with Function Parameters

```go
package main

import "fmt"

func main() {
    fmt.Println("Starting main function")
    
    // Parameters are evaluated when defer is called
    x := 10
    defer printValue("Deferred with x =", x)
    
    x = 20
    defer printValue("Deferred with x =", x)
    
    fmt.Println("x is now", x)
    fmt.Println("Ending main function")
}

func printValue(message string, value int) {
    fmt.Printf("%s %d\n", message, value)
}
```

**Output:**
```
Starting main function
x is now 20
Ending main function
Deferred with x = 20
Deferred with x = 10
```

## Advanced Examples

### Example 4: File Operations with Defer

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Open file
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close() // Ensure file is closed
    
    // Read file content
    data := make([]byte, 100)
    n, err := file.Read(data)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
    fmt.Println("File operations completed")
}

func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Will be called when function returns
    
    // Process file content
    fmt.Printf("Processing file: %s\n", filename)
    
    // Simulate some processing
    for i := 0; i < 3; i++ {
        fmt.Printf("Processing step %d\n", i+1)
    }
    
    return nil
}
```

### Example 5: Database Operations with Defer

```go
package main

import "fmt"

// Simulated database connection
type Database struct {
    connected bool
}

func (db *Database) Connect() error {
    fmt.Println("Connecting to database...")
    db.connected = true
    return nil
}

func (db *Database) Close() {
    if db.connected {
        fmt.Println("Closing database connection...")
        db.connected = false
    }
}

func (db *Database) Query(sql string) ([]string, error) {
    if !db.connected {
        return nil, fmt.Errorf("database not connected")
    }
    fmt.Printf("Executing query: %s\n", sql)
    return []string{"result1", "result2", "result3"}, nil
}

func main() {
    db := &Database{}
    
    // Connect to database
    if err := db.Connect(); err != nil {
        fmt.Printf("Error connecting: %v\n", err)
        return
    }
    defer db.Close() // Ensure connection is closed
    
    // Perform database operations
    results, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Printf("Query error: %v\n", err)
        return
    }
    
    fmt.Printf("Query results: %v\n", results)
    fmt.Println("Database operations completed")
}
```

### Example 6: Mutex Operations with Defer

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock() // Ensure mutex is unlocked
    
    c.count++
    fmt.Printf("Count incremented to: %d\n", c.count)
}

func (c *Counter) GetCount() int {
    c.mu.Lock()
    defer c.mu.Unlock() // Ensure mutex is unlocked
    
    return c.count
}

func main() {
    counter := &Counter{}
    
    // Simulate concurrent access
    for i := 0; i < 5; i++ {
        go func(id int) {
            for j := 0; j < 3; j++ {
                counter.Increment()
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    // Wait for goroutines to complete
    time.Sleep(2 * time.Second)
    
    fmt.Printf("Final count: %d\n", counter.GetCount())
}
```

## Practical Examples

### Example 7: HTTP Server with Defer

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    defer func() {
        duration := time.Since(start)
        fmt.Printf("Request completed in %v\n", duration)
    }()
    
    fmt.Printf("Handling request: %s %s\n", r.Method, r.URL.Path)
    
    // Simulate some processing
    time.Sleep(100 * time.Millisecond)
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Request processed successfully"))
}

func main() {
    http.HandleFunc("/", handleRequest)
    
    fmt.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}
```

### Example 8: Error Handling with Defer

```go
package main

import (
    "fmt"
    "os"
)

func processData(filename string) error {
    // Open file
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Process file
    fmt.Printf("Processing file: %s\n", filename)
    
    // Simulate processing that might fail
    if filename == "error.txt" {
        return fmt.Errorf("processing failed for %s", filename)
    }
    
    fmt.Println("File processed successfully")
    return nil
}

func main() {
    files := []string{"data.txt", "error.txt", "config.txt"}
    
    for _, filename := range files {
        fmt.Printf("\nProcessing %s:\n", filename)
        if err := processData(filename); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
    }
}
```

### Example 9: Defer with Panic Recovery

```go
package main

import "fmt"

func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    fmt.Println("Starting risky operation")
    
    // Simulate a panic
    panic("Something went wrong!")
    
    fmt.Println("This won't be printed")
}

func safeOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic in safe operation: %v\n", r)
        }
    }()
    
    fmt.Println("Starting safe operation")
    
    // This will cause a panic
    var slice []int
    fmt.Println(slice[10]) // This will panic
    
    fmt.Println("This won't be printed")
}

func main() {
    fmt.Println("Testing panic recovery with defer")
    
    riskyOperation()
    fmt.Println("Program continues after risky operation")
    
    safeOperation()
    fmt.Println("Program continues after safe operation")
}
```

### Example 10: Defer with Return Values

```go
package main

import "fmt"

func deferWithReturn() (result int) {
    defer func() {
        result++ // This modifies the return value
        fmt.Printf("Defer: result is now %d\n", result)
    }()
    
    fmt.Println("Function body: setting result to 5")
    return 5
}

func deferWithNamedReturn() (result int) {
    defer func() {
        result *= 2 // This modifies the return value
        fmt.Printf("Defer: result is now %d\n", result)
    }()
    
    fmt.Println("Function body: setting result to 10")
    result = 10
    return
}

func main() {
    fmt.Println("Testing defer with return values")
    
    result1 := deferWithReturn()
    fmt.Printf("Final result 1: %d\n", result1)
    
    fmt.Println()
    
    result2 := deferWithNamedReturn()
    fmt.Printf("Final result 2: %d\n", result2)
}
```

## Common Patterns and Best Practices

### Pattern 1: Resource Cleanup
```go
func processResource() error {
    resource := acquireResource()
    defer releaseResource(resource)
    
    // Use resource
    return useResource(resource)
}
```

### Pattern 2: Logging and Debugging
```go
func processData(data []int) {
    defer func() {
        fmt.Printf("Processed %d items\n", len(data))
    }()
    
    // Process data
    for _, item := range data {
        // Process item
    }
}
```

### Pattern 3: Timing Operations
```go
func timedOperation() {
    start := time.Now()
    defer func() {
        fmt.Printf("Operation took %v\n", time.Since(start))
    }()
    
    // Perform operation
}
```

### Pattern 4: Error Handling
```go
func safeOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("operation failed: %v", r)
        }
    }()
    
    // Perform operation that might panic
    return nil
}
```

## Important Rules and Considerations

### 1. Defer Execution Order
- Multiple defer statements execute in LIFO (Last In, First Out) order
- Defer statements are executed when the function returns, not when they are encountered

### 2. Parameter Evaluation
- Parameters to deferred functions are evaluated when the defer statement is executed
- Use pointers or closures to capture current values

### 3. Return Value Modification
- Deferred functions can modify named return values
- This is useful for cleanup and error handling

### 4. Panic Recovery
- Defer statements execute even when a function panics
- Use defer with recover() to handle panics gracefully

### 5. Performance Considerations
- Defer has a small performance cost
- Use defer for cleanup operations where the cost is justified
- Avoid defer in performance-critical loops

## Common Mistakes to Avoid

### Mistake 1: Defer in Loop
```go
// ❌ Bad: defer in loop
func badExample() {
    for i := 0; i < 10; i++ {
        defer fmt.Println(i) // All defers execute at end of function
    }
}

// ✅ Good: defer outside loop
func goodExample() {
    for i := 0; i < 10; i++ {
        func() {
            defer fmt.Println(i) // Each defer executes immediately
        }()
    }
}
```

### Mistake 2: Ignoring Errors
```go
// ❌ Bad: ignoring error
func badExample() {
    file, _ := os.Open("file.txt")
    defer file.Close() // Might panic if file is nil
}

// ✅ Good: check error
func goodExample() {
    file, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer file.Close()
}
```

## Conclusion

Defer statements in Go provide:

1. **Resource management**: Automatic cleanup of resources
2. **Error safety**: Cleanup happens even on errors or panics
3. **Code clarity**: Keep cleanup code close to resource acquisition
4. **Flexibility**: Can defer any function call or statement
5. **Reliability**: Ensures important operations always execute

Understanding defer is essential for writing robust Go programs, especially when dealing with resources, error handling, and cleanup operations.
