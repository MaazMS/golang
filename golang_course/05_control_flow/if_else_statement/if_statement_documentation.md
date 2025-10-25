# If Statement in Go

## Overview
The `if` statement is a fundamental control flow construct in Go that allows you to execute a block of code conditionally based on a boolean expression.

## Key Concepts
1. The block of code and statements are executed if the given condition evaluates to `true`.
2. If the condition evaluates to `false`, the block of code and statements are not executed.
3. The condition must be a boolean expression.

## Syntax
```go
if condition {
    // statements
    // statements
}
```

## Examples

### Basic If Statement
```go
package main

import "fmt"

func main() {
    age := 18
    
    if age >= 18 {
        fmt.Println("You are eligible to vote")
    }
}
```

### If Statement with Boolean Variables
```go
package main

import "fmt"

func main() {
    isLoggedIn := true
    
    if isLoggedIn {
        fmt.Println("Welcome back!")
    }
}
```

### If Statement with Function Calls
```go
package main

import "fmt"

func isEven(number int) bool {
    return number%2 == 0
}

func main() {
    num := 4
    
    if isEven(num) {
        fmt.Println("The number is even")
    }
}
```

## Important Notes
- The condition must evaluate to a boolean value
- Go does not require parentheses around the condition
- The opening brace `{` must be on the same line as the `if` statement
- Go uses strict formatting rules for braces
