# Constants in Go - Complete Beginner's Guide

## What are Constants?

Constants in Go are values that cannot be changed once they are declared. They are like variables, but their value is fixed and cannot be modified during the program's execution. Think of them as "read-only" values that remain the same throughout your program.

## Why Use Constants?

1. **Prevent accidental changes**: Once set, constants cannot be modified
2. **Improve code readability**: Give meaningful names to values
3. **Better performance**: Compiler can optimize constant values
4. **Magic number elimination**: Replace hardcoded values with meaningful names

## Types of Constants

Go supports several types of constants:
- **Boolean constants**: `true`, `false`
- **String constants**: `"Hello World"`
- **Numeric constants**: `42`, `3.14`, `1 + 2i`
- **Rune constants**: `'A'`, `'中'`

## Declaring Constants

### 1. Single Constant Declaration

```go
package main

import "fmt"

func main() {
    // Typed constant - explicitly specify the type
    const pi float64 = 3.14159
    const name string = "Go Programming"
    const isActive bool = true
    
    fmt.Println("Pi:", pi)
    fmt.Println("Name:", name)
    fmt.Println("Active:", isActive)
}
```

### 2. Untyped Constants (Recommended)

```go
package main

import "fmt"

func main() {
    // Untyped constants - Go infers the type
    const pi = 3.14159
    const name = "Go Programming"
    const isActive = true
    
    fmt.Printf("Pi: %v (Type: %T)\n", pi, pi)
    fmt.Printf("Name: %v (Type: %T)\n", name, name)
    fmt.Printf("Active: %v (Type: %T)\n", isActive, isActive)
}
```

### 3. Multiple Constants Declaration

```go
package main

import "fmt"

func main() {
    // Declaring multiple constants
    const (
        companyName = "TechCorp"
        maxUsers    = 1000
        isPublic    = true
        version     = "1.0.0"
    )
    
    fmt.Println("Company:", companyName)
    fmt.Println("Max Users:", maxUsers)
    fmt.Println("Public:", isPublic)
    fmt.Println("Version:", version)
}
```

## Practical Examples

### Example 1: Mathematical Constants

```go
package main

import "fmt"

func main() {
    const (
        pi     = 3.14159
        e      = 2.71828
        golden = 1.61803
    )
    
    radius := 5.0
    area := pi * radius * radius
    
    fmt.Printf("Area of circle with radius %.1f: %.2f\n", radius, area)
    fmt.Printf("Euler's number: %.5f\n", e)
    fmt.Printf("Golden ratio: %.5f\n", golden)
}
```

### Example 2: Application Configuration

```go
package main

import "fmt"

func main() {
    const (
        appName    = "MyGoApp"
        version    = "2.1.0"
        maxRetries = 3
        timeout    = 30
    )
    
    fmt.Printf("Application: %s v%s\n", appName, version)
    fmt.Printf("Max retries: %d\n", maxRetries)
    fmt.Printf("Timeout: %d seconds\n", timeout)
}
```

### Example 3: HTTP Status Codes

```go
package main

import "fmt"

func main() {
    const (
        statusOK        = 200
        statusNotFound  = 404
        statusError     = 500
        statusCreated   = 201
    )
    
    responseCode := 200
    
    if responseCode == statusOK {
        fmt.Println("Request successful!")
    } else if responseCode == statusNotFound {
        fmt.Println("Resource not found")
    } else if responseCode == statusError {
        fmt.Println("Internal server error")
    }
}
```

## Constants vs Variables

| Feature | Constants | Variables |
|---------|-----------|-----------|
| Declaration | `const` | `var` |
| Can be changed | ❌ No | ✅ Yes |
| Must be initialized | ✅ Yes | ❌ No (can be zero value) |
| Memory allocation | Compile time | Runtime |
| Performance | Better | Slower |

### Comparison Example

```go
package main

import "fmt"

func main() {
    // Constants
    const maxSize = 100
    const appName = "Calculator"
    
    // Variables
    var currentSize = 50
    var userName = "John"
    
    fmt.Printf("Max size: %d (constant)\n", maxSize)
    fmt.Printf("Current size: %d (variable)\n", currentSize)
    
    // This would cause a compilation error:
    // maxSize = 200  // ❌ Cannot assign to maxSize
    
    // This is allowed:
    currentSize = 75  // ✅ Variables can be changed
    fmt.Printf("Updated current size: %d\n", currentSize)
}
```

## Best Practices

### 1. Use Descriptive Names

```go
// ❌ Bad
const a = 3.14
const b = "error"

// ✅ Good
const pi = 3.14
const errorMessage = "error"
```

### 2. Group Related Constants

```go
// ✅ Good - Group related constants
const (
    // Database configuration
    dbHost     = "localhost"
    dbPort     = 5432
    dbName     = "myapp"
    dbUser     = "admin"
    
    // API configuration
    apiVersion = "v1"
    apiTimeout = 30
)
```

### 3. Use Constants for Magic Numbers

```go
// ❌ Bad - Magic numbers
func calculateTax(amount float64) float64 {
    return amount * 0.15  // What is 0.15?
}

// ✅ Good - Named constants
const taxRate = 0.15

func calculateTax(amount float64) float64 {
    return amount * taxRate
}
```

## Advanced Constant Features

### 1. Constant Expressions

```go
package main

import "fmt"

func main() {
    const (
        a = 10
        b = 20
        c = a + b        // 30
        d = a * b        // 200
        e = b / a        // 2
    )
    
    fmt.Printf("a = %d\n", a)
    fmt.Printf("b = %d\n", b)
    fmt.Printf("c = a + b = %d\n", c)
    fmt.Printf("d = a * b = %d\n", d)
    fmt.Printf("e = b / a = %d\n", e)
}
```

### 2. String Constants

```go
package main

import "fmt"

func main() {
    const (
        greeting = "Hello, "
        name     = "World"
        message  = greeting + name  // String concatenation
    )
    
    fmt.Println(message)  // Output: Hello, World
}
```

## Common Mistakes to Avoid

### 1. Trying to Modify Constants

```go
package main

func main() {
    const maxUsers = 100
    
    // ❌ This will cause a compilation error
    // maxUsers = 200  // cannot assign to maxUsers
}
```

### 2. Not Initializing Constants

```go
// ❌ This will cause a compilation error
// const name  // missing value in const declaration

// ✅ Correct
const name = "Go"
```

### 3. Using Variables When Constants Are Better

```go
// ❌ Bad - Using variable for fixed value
var maxRetries = 3

// ✅ Good - Using constant for fixed value
const maxRetries = 3
```

## Summary

- **Constants** are immutable values that cannot be changed
- Use `const` keyword to declare constants
- Constants must be initialized when declared
- Prefer untyped constants for flexibility
- Use descriptive names for better code readability
- Group related constants together
- Replace magic numbers with named constants

Constants are essential for writing maintainable and readable Go code. They help prevent bugs and make your code more self-documenting.
