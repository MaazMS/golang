# Loop Control Statements in Go

## Overview
Loop control statements allow you to change the normal flow of loop execution. Go provides three main control statements: `break`, `continue`, and `goto`. These statements give you fine-grained control over loop behavior.

## Break Statement

### Purpose
The `break` statement is used to **terminate the loop immediately** and transfer control to the statement following the loop.

### Usage
- **Inside single loops**: Terminates the current loop
- **Inside nested loops**: Terminates the innermost loop by default
- **With labels**: Can terminate specific outer loops when used with labels

### Syntax
```go
break
// or
break label
```

### Examples

#### Example 1: Basic Break
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            fmt.Println("Breaking at 5")
            break
        }
        fmt.Printf("Value: %d\n", i)
    }
    fmt.Println("Loop ended")
}
```

#### Example 2: Break in Nested Loops
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 3; j++ {
            if j == 2 {
                fmt.Println("Breaking inner loop")
                break // Only breaks inner loop
            }
            fmt.Printf("  Inner loop: %d\n", j)
        }
    }
}
```

#### Example 3: Break with Labels
```go
package main

import "fmt"

func main() {
outer:
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop: %d\n", i)
        for j := 1; j <= 3; j++ {
            if j == 2 {
                fmt.Println("Breaking outer loop")
                break outer // Breaks the outer loop
            }
            fmt.Printf("  Inner loop: %d\n", j)
        }
    }
    fmt.Println("Both loops ended")
}
```

## Continue Statement

### Purpose
The `continue` statement **skips the current iteration** and moves to the next iteration of the loop.

### Behavior
- **Skips remaining code**: Jumps to the next iteration
- **Loop continues**: The loop itself is not terminated
- **Post statement executes**: The post statement (increment/decrement) still runs

### Syntax
```go
continue
// or
continue label
```

### Examples

#### Example 1: Basic Continue
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            fmt.Println("Skipping iteration 3")
            continue
        }
        fmt.Printf("Value: %d\n", i)
    }
}
```

#### Example 2: Continue with Labels
```go
package main

import "fmt"

func main() {
outer:
    for i := 1; i <= 2; i++ {
        fmt.Printf("Outer: %d\n", i)
        for j := 1; j <= 3; j++ {
            if j == 2 {
                fmt.Println("Continuing outer loop")
                continue outer
            }
            fmt.Printf("  Inner: %d\n", j)
        }
    }
}
```

#### Example 3: Processing Only Even Numbers
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i%2 != 0 {
            continue // Skip odd numbers
        }
        fmt.Printf("Even number: %d\n", i)
    }
}
```

## Goto Statement

### Purpose
The `goto` statement **transfers control to a labeled statement** within the same function.

### Important Notes
- **Function scope only**: `goto` works only within the same function
- **Avoid in loops**: Using `goto` inside loops can create infinite loops
- **Use sparingly**: Generally discouraged in modern programming
- **Label placement**: Labels must be placed before statements

### Syntax
```go
goto label
...
label: statement
```

### Examples

#### Example 1: Basic Goto
```go
package main

import "fmt"

func main() {
    fmt.Println("Before goto")
    goto skip
    fmt.Println("This won't be printed")
    
skip:
    fmt.Println("After goto")
}
```

#### Example 2: Error Handling with Goto
```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    for i := 0; i < 5; i++ {
        if processData(i) {
            fmt.Printf("Successfully processed %d\n", i)
        } else {
            fmt.Printf("Failed to process %d, retrying...\n", i)
            goto retry
        }
        continue
        
    retry:
        fmt.Printf("Retrying %d\n", i)
        i-- // Decrement to retry the same index
    }
}

func processData(n int) bool {
    // Simulate processing with random success
    return rand.Intn(2) == 1
}
```

#### Example 3: Cleanup with Goto
```go
package main

import "fmt"

func main() {
    fmt.Println("Starting process")
    
    // Simulate some work
    if true {
        fmt.Println("Work completed successfully")
        goto cleanup
    }
    
    fmt.Println("This won't be printed")
    
cleanup:
    fmt.Println("Cleaning up resources")
    fmt.Println("Process finished")
}
```

## Best Practices

### 1. Prefer Break and Continue
```go
// Good: Using break and continue
for i := 1; i <= 10; i++ {
    if i == 5 {
        break
    }
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}

// Avoid: Using goto for simple cases
for i := 1; i <= 10; i++ {
    if i == 5 {
        goto end
    }
    if i%2 == 0 {
        goto next
    }
    fmt.Println(i)
next:
}
end:
```

### 2. Use Labels Meaningfully
```go
// Good: Descriptive labels
outerLoop:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if someCondition {
                break outerLoop
            }
        }
    }

// Avoid: Generic labels
loop1:
    for i := 1; i <= 3; i++ {
        // ...
    }
```

### 3. Limit Goto Usage
```go
// Good: Use goto for cleanup
func processFile() error {
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    
    // Process file
    if someError {
        goto cleanup
    }
    
cleanup:
    file.Close()
    return nil
}

// Avoid: Complex goto patterns
func badExample() {
    goto start
middle:
    goto end
start:
    goto middle
end:
}
```

## Common Patterns

### 1. Early Exit Pattern
```go
for i := 1; i <= 10; i++ {
    if shouldExit(i) {
        break
    }
    process(i)
}
```

### 2. Skip Pattern
```go
for i := 1; i <= 10; i++ {
    if shouldSkip(i) {
        continue
    }
    process(i)
}
```

### 3. Nested Loop Control
```go
outer:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if found {
                break outer
            }
        }
    }
```

## Performance Considerations

- **Break and Continue**: Minimal performance impact
- **Goto**: Very fast but can make code hard to follow
- **Labels**: No runtime cost, compile-time feature

## When to Use Each Statement

### Use `break` when:
- You need to exit a loop early
- You've found what you're looking for
- An error condition is met

### Use `continue` when:
- You want to skip the current iteration
- You need to process only certain items
- You want to avoid nested conditions

### Use `goto` when:
- You need to implement cleanup logic
- You're handling complex error scenarios
- You're implementing state machines (sparingly)
