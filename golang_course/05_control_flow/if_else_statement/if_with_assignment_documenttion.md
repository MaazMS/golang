# If Statement with Assignment in Go

## Overview
Go allows you to combine variable assignment with conditional statements using the `if` statement. This is a powerful feature that enables you to declare and assign variables within the condition, and those variables are only available within the scope of the `if` block.

## Key Concepts
1. You can declare and assign variables in the `if` statement condition.
2. The assigned variables are only available within the `if` block and its `else` block.
3. This pattern is commonly used with function calls that return multiple values (like error handling).
4. The condition can use the assigned variables immediately.

## Syntax
```go
if variable := expression; condition {
    // statements using the variable
    // variable is available here
} else {
    // variable is also available here
}
// variable is not available here
```

## Examples

### Basic Assignment with Condition
```go
package main

import "fmt"

func main() {
    if x := 10; x > 5 {
        fmt.Printf("x is %d and it's greater than 5\n", x)
    } else {
        fmt.Printf("x is %d and it's not greater than 5\n", x)
    }
    // x is not available here
}
```

### Function Call with Error Handling
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    
    if num, err := strconv.Atoi(str); err == nil {
        fmt.Printf("Successfully converted '%s' to %d\n", str, num)
    } else {
        fmt.Printf("Failed to convert '%s': %v\n", str, err)
    }
}
```

### Multiple Variable Assignment
```go
package main

import "fmt"

func divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func main() {
    a, b := 10, 2
    
    if result, success := divide(a, b); success {
        fmt.Printf("%d divided by %d equals %d\n", a, b, result)
    } else {
        fmt.Println("Division by zero is not allowed")
    }
}
```

### Map Lookup with Assignment
```go
package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 92,
    }
    
    name := "Alice"
    if score, exists := scores[name]; exists {
        fmt.Printf("%s's score is %d\n", name, score)
    } else {
        fmt.Printf("No score found for %s\n", name)
    }
}
```

### Channel Operations
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 42
    
    if value, ok := <-ch; ok {
        fmt.Printf("Received value: %d\n", value)
    } else {
        fmt.Println("Channel is closed")
    }
}
```

### Type Assertion with Assignment
```go
package main

import "fmt"

func main() {
    var i interface{} = "Hello, World!"
    
    if str, ok := i.(string); ok {
        fmt.Printf("The interface contains a string: %s\n", str)
    } else {
        fmt.Println("The interface does not contain a string")
    }
}
```

## Important Notes
- Variables declared in the `if` statement are only available within the `if` and `else` blocks
- This pattern is particularly useful for error handling in Go
- The assignment and condition are separated by a semicolon `;`
- You can declare multiple variables in the assignment part
- This is a common Go idiom for handling functions that return multiple values
- The scope of the variables is limited to the `if-else` block
