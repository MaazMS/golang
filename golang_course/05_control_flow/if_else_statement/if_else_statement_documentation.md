# If-Else Statement in Go

## Overview
The `if-else` statement provides an alternative execution path when the initial condition is not met. It allows you to execute one block of code if the condition is true, and a different block if the condition is false.

## Key Concepts
1. The block of code and statements are executed if the given condition evaluates to `true`.
2. If the condition evaluates to `false`, the `else` block is executed.
3. The `else` keyword must be on the same line as the closing brace `}` of the `if` block.
4. The `else` statement is not executed if the given condition is fulfilled.

## Syntax
```go
if condition {
    // statements executed if condition is true
    // statements
} else {
    // statements executed if condition is false
    // statements
}
```

## Examples

### Basic If-Else Statement
```go
package main

import "fmt"

func main() {
    age := 16
    
    if age >= 18 {
        fmt.Println("You are eligible to vote")
    } else {
        fmt.Println("You are not eligible to vote yet")
    }
}
```

### If-Else with String Comparison
```go
package main

import "fmt"

func main() {
    userRole := "admin"
    
    if userRole == "admin" {
        fmt.Println("Access granted to admin panel")
    } else {
        fmt.Println("Access denied - admin privileges required")
    }
}
```

### If-Else with Numeric Conditions
```go
package main

import "fmt"

func main() {
    score := 85
    
    if score >= 90 {
        fmt.Println("Grade: A")
    } else {
        fmt.Println("Grade: Below A")
    }
}
```

## Important Notes
- The `else` keyword must be on the same line as the closing brace of the `if` block
- Only one of the two blocks (if or else) will be executed
- Both blocks are mutually exclusive
- The condition must evaluate to a boolean value
