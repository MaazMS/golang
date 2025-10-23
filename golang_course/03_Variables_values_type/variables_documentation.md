# Variables in Go

## Introduction to Variables

A **variable** is like a storage box for a value in your program. The name of a variable allows you to access and manipulate its stored value. In Go, every variable has a specific type, making it a **statically typed** programming language.

## Understanding Variables in Go

### What are Variables?

Variables in Go are containers that store values of a specific type. Unlike dynamic programming languages such as Python or JavaScript, Go requires you to declare the type of a variable when you create it, and this type cannot be changed later.

### Key Characteristics of Go Variables

1. **Static Typing**: Every variable has a fixed type that cannot be changed after declaration
2. **Strong Typing**: Variables can only store values that match their declared type
3. **Type Safety**: Go prevents type mismatches at compile time
4. **Memory Efficiency**: Variables occupy specific amounts of memory based on their type

## Basic Variable Concepts

## Zero Values

In Go, variables declared without an initial value are given their **zero value**:

```go
package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    var p *int
    
    fmt.Printf("Zero values:\n")
    fmt.Printf("int: %d\n", i)           // 0
    fmt.Printf("float64: %.2f\n", f)     // 0.00
    fmt.Printf("bool: %t\n", b)          // false
    fmt.Printf("string: '%s'\n", s)      // "" (empty string)
    fmt.Printf("pointer: %v\n", p)       // <nil>
}
```

## Variable Scope and Lifetime

### Global vs Local Variables

```go
package main

import "fmt"

// Global variables (package-level)
var globalName string = "Global Variable"
var globalCounter int = 0

func main() {
    // Local variables (function-level)
    var localName string = "Local Variable"
    localCounter := 0
    
    fmt.Printf("Global name: %s\n", globalName)
    fmt.Printf("Local name: %s\n", localName)
    
    // Modify global variable
    globalCounter++
    localCounter++
    
    fmt.Printf("Global counter: %d\n", globalCounter)
    fmt.Printf("Local counter: %d\n", localCounter)
    
    // Call another function
    anotherFunction()
}

func anotherFunction() {
    // Can access global variables
    fmt.Printf("From another function - Global name: %s\n", globalName)
    fmt.Printf("From another function - Global counter: %d\n", globalCounter)
    
    // Cannot access local variables from main()
    // fmt.Printf("Local name: %s\n", localName)  // This would cause an error
    
    // Create new local variables
    var localVar string = "Function Local Variable"
    fmt.Printf("Function local variable: %s\n", localVar)
}
```

## Variable Assignment and Reassignment

### Basic Assignment Operations

```go
package main

import "fmt"

func main() {
    // Initial assignment
    var name string = "Maaz"
    var age int = 25
    
    fmt.Printf("Initial - Name: %s, Age: %d\n", name, age)
    
    // Reassignment (same type)
    name = "Maaz Shaikh"
    age = 30
    
    fmt.Printf("After reassignment - Name: %s, Age: %d\n", name, age)
    
    // Multiple assignment
    var x, y int = 10, 20
    fmt.Printf("Before swap - x: %d, y: %d\n", x, y)
    
    // Swap values
    x, y = y, x
    fmt.Printf("After swap - x: %d, y: %d\n", x, y)
    
    // Assignment with expressions
    var sum int = x + y
    var product int = x * y
    
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
```

## Constants vs Variables

### Understanding the Difference

```go
package main

import "fmt"

const (
    // Constants (cannot be changed)
    PI = 3.14159
    MAX_SIZE = 100
    APP_NAME = "Go Variables Demo"
)

func main() {
    // Variables (can be changed)
    var radius float64 = 5.0
    var area float64 = PI * radius * radius
    
    fmt.Printf("App: %s\n", APP_NAME)
    fmt.Printf("Radius: %.2f\n", radius)
    fmt.Printf("Area: %.2f\n", area)
    
    // Variables can be reassigned
    radius = 10.0
    area = PI * radius * radius
    fmt.Printf("New radius: %.2f\n", radius)
    fmt.Printf("New area: %.2f\n", area)
    
    // Constants cannot be reassigned
    // PI = 3.14  // This would cause a compile error
}
```


## Best Practices for Variables

### 1. Use Descriptive Names

```go
// Good variable names
var userName string = "maaz"
var userAge int = 25
var isUserActive bool = true

// Avoid unclear names
var n string = "maaz"        // What does 'n' mean?
var a int = 25               // What does 'a' represent?
var f bool = true            // What does 'f' stand for?
```

### 2. Choose Appropriate Types

```go
// Use appropriate types for your data
var userID int64 = 1234567890        // Use int64 for large numbers
var price float64 = 99.99            // Use float64 for currency
var isPremium bool = true            // Use bool for true/false values
var message string = "Hello"         // Use string for text
```

### 3. Initialize Variables When Possible

```go
// Good - initialize with meaningful values
var userName string = "Guest"
var userCount int = 0
var isLoggedIn bool = false

// Avoid - rely on zero values when initial value matters
var userName string  // Will be empty string
var userCount int    // Will be 0
var isLoggedIn bool  // Will be false
```

### 4. Use Short Declaration When Appropriate

```go
func processUser() {
    // Use short declaration for local variables
    name := "Maaz"
    age := 25
    isActive := true
    
    // Use var for package-level variables
    // var globalConfig = loadConfig()
    
    fmt.Printf("Processing user: %s, age: %d, active: %t\n", name, age, isActive)
}
```

## Common Variable Patterns

### 1. Variable Swapping

```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
    
    // Swap without temporary variable
    a, b = b, a
    fmt.Printf("After swap: a=%d, b=%d\n", a, b)
}
```

### 2. Variable Shadowing

```go
package main

import "fmt"

var name string = "Global Name"

func main() {
    fmt.Printf("Global name: %s\n", name)
    
    // Local variable shadows global variable
    name := "Local Name"
    fmt.Printf("Local name: %s\n", name)
    
    // Access global variable
    fmt.Printf("Global name (still): %s\n", name)
}
```

### 3. Variable Initialization with Functions

```go
package main

import (
    "fmt"
    "time"
)

func getCurrentTime() string {
    return time.Now().Format("15:04:05")
}

func main() {
    // Initialize variable with function result
    currentTime := getCurrentTime()
    fmt.Printf("Current time: %s\n", currentTime)
    
    // Initialize multiple variables
    name, age, isActive := "Maaz", 25, true
    fmt.Printf("Name: %s, Age: %d, Active: %t\n", name, age, isActive)
}
```

## Summary

Variables in Go:

- **Are statically typed**: Each variable has a fixed type
- **Provide type safety**: Prevent type mismatches at compile time
- **Occupy memory**: Variables are stored in computer memory
- **Have zero values**: Uninitialized variables get default values
- **Have scope**: Variables exist within their declared scope
- **Are created at runtime**: Variables don't exist at compile time
- **Support reassignment**: Values can be changed (but not types)

Understanding variables is fundamental to Go programming. They provide the foundation for storing and manipulating data in your programs while maintaining type safety and memory efficiency.

## Key Takeaways

1. **Static Typing**: Go variables have fixed types that cannot be changed
2. **Type Safety**: Variables can only store values of their declared type
3. **Memory Management**: Variables occupy specific amounts of memory
4. **Scope Rules**: Variables exist within their declared scope
5. **Zero Values**: Uninitialized variables get default values
6. **Runtime Creation**: Variables are created when the program runs

Mastering variables in Go is essential for writing safe, efficient, and maintainable programs.  