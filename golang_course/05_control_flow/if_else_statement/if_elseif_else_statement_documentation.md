# If-Else If-Else Statement in Go

## Overview
The `if-else if-else` statement allows you to check multiple conditions in sequence. It provides a way to handle multiple possible scenarios with different execution paths.

## Key Concepts

### If-Else If Chain
1. The `if` statement is executed if the given condition evaluates to `true`.
2. If the condition is not fulfilled, the program checks the next `else if` condition.
3. The `else if` statement comes after the `if` condition.
4. You can have multiple `else if` statements.
5. If an `else if` condition is not fulfilled, the program moves to the next condition.

### If-Else If-Else Chain
1. This is important because if none of the `else if` conditions are fulfilled, the program will not enter any of those blocks.
2. The `else` statement is used to execute code when all conditions are false.
3. The `else` block serves as a default case.

## Syntax

### If-Else If
```go
if condition1 {
    // statements executed if condition1 is true
} else if condition2 {
    // statements executed if condition2 is true
} else if condition3 {
    // statements executed if condition3 is true
}
```

### If-Else If-Else
```go
if condition1 {
    // statements executed if condition1 is true
} else if condition2 {
    // statements executed if condition2 is true
} else if condition3 {
    // statements executed if condition3 is true
} else {
    // statements executed if all conditions are false
}
```

## Examples

### Grade Classification
```go
package main

import "fmt"

func main() {
    score := 85
    
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else if score >= 60 {
        fmt.Println("Grade: D")
    } else {
        fmt.Println("Grade: F")
    }
}
```

### User Role Authorization
```go
package main

import "fmt"

func main() {
    userRole := "moderator"
    
    if userRole == "admin" {
        fmt.Println("Full access granted")
    } else if userRole == "moderator" {
        fmt.Println("Moderate access granted")
    } else if userRole == "user" {
        fmt.Println("Basic access granted")
    } else {
        fmt.Println("Access denied - invalid role")
    }
}
```

### Temperature Classification
```go
package main

import "fmt"

func main() {
    temperature := 25
    
    if temperature > 30 {
        fmt.Println("Hot weather")
    } else if temperature > 20 {
        fmt.Println("Warm weather")
    } else if temperature > 10 {
        fmt.Println("Cool weather")
    } else {
        fmt.Println("Cold weather")
    }
}
```

## Important Notes
- Only the first condition that evaluates to `true` will execute its block
- Once a condition is met, the remaining conditions are not evaluated
- The `else` block is optional but serves as a default case
- All conditions must evaluate to boolean values
- The `else if` and `else` keywords must be on the same line as the previous closing brace
