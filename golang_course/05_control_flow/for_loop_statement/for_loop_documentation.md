# For Loop in Go

## Overview
Go language contains only a single loop construct: the **for loop**. Unlike other programming languages that have multiple loop types (while, do-while), Go uses the for loop for all iteration needs.

The for loop allows you to execute code for a specific number of iterations or until a specific condition is met.

## Syntax
```go
for initialization; condition; post {
    // statements
    // statements
}
```

## Components

### 1. Initialization
- **Optional**: The initialization statement is executed once before the loop begins
- **Purpose**: Used to declare and initialize loop variables
- **Scope**: Variables declared here are only accessible within the loop

### 2. Condition
- **Required**: A boolean expression that determines whether the loop continues
- **Execution**: Evaluated before each iteration
- **Behavior**: 
  - If `true`: Loop continues to the next iteration
  - If `false`: Loop terminates

### 3. Post Statement
- **Optional**: Executed after each iteration completes
- **Purpose**: Typically used for incrementing or decrementing loop variables
- **Timing**: Runs after the loop body but before the condition is checked again

## Basic Examples

### Example 1: Simple Counter Loop
```go
package main

import "fmt"

func main() {
    // Count from 1 to 5
    for i := 1; i <= 5; i++ {
        fmt.Printf("Iteration %d\n", i)
    }
}
```

### Example 2: Loop with Different Increment
```go
package main

import "fmt"

func main() {
    // Count by 2s from 0 to 10
    for i := 0; i <= 10; i += 2 {
        fmt.Printf("Even number: %d\n", i)
    }
}
```

### Example 3: Reverse Loop
```go
package main

import "fmt"

func main() {
    // Count down from 5 to 1
    for i := 5; i >= 1; i-- {
        fmt.Printf("Countdown: %d\n", i)
    }
}
```

## Alternative For Loop Syntax

### 1. While-Style Loop
```go
// Traditional while loop equivalent
i := 0
for i < 5 {
    fmt.Printf("Value: %d\n", i)
    i++
}
```

### 2. Infinite Loop
```go
// Infinite loop (use with caution)
for {
    // This will run forever
    // Use break to exit
}
```

### 3. Range-Based Loop
```go
// Iterate over slices, arrays, maps, strings
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

## Best Practices

1. **Use meaningful variable names**: Instead of `i`, use descriptive names like `index`, `counter`, or `item`
2. **Keep loops simple**: Avoid complex logic inside loops
3. **Use range when possible**: For iterating over collections, prefer range-based loops
4. **Be careful with infinite loops**: Always ensure there's a way to exit
5. **Consider performance**: For large datasets, consider the efficiency of your loop logic

## Common Use Cases

- **Iterating through arrays and slices**
- **Processing collections of data**
- **Implementing algorithms that require repetition**
- **Reading data from files or databases**
- **Implementing retry logic**
- **Creating user interfaces with repeated elements**
  
# Nested For Loops in Go

## Overview
Nested for loops are loops that contain other loops within their body. This creates a hierarchical structure where the inner loop executes completely for each iteration of the outer loop. Nested loops are essential for working with multi-dimensional data structures and complex algorithms.

## Basic Syntax
```go
for outer_initialization; outer_condition; outer_post {
    // Outer loop statements
    
    for inner_initialization; inner_condition; inner_post {
        // Inner loop statements
        // Inner loop body
    }
    
    // More outer loop statements
}
```

## How Nested Loops Work

1. **Outer loop starts**: The outer loop begins its first iteration
2. **Inner loop executes completely**: The inner loop runs from start to finish
3. **Outer loop continues**: After the inner loop completes, the outer loop moves to its next iteration
4. **Process repeats**: Steps 2-3 repeat until the outer loop condition becomes false

## Basic Examples

### Example 1: Simple Nested Loop
```go
package main

import "fmt"

func main() {
    fmt.Println("Nested Loop Example:")
    
    for i := 1; i <= 3; i++ {
        fmt.Printf("Outer loop iteration: %d\n", i)
        
        for j := 1; j <= 2; j++ {
            fmt.Printf("  Inner loop iteration: %d\n", j)
        }
        
        fmt.Println("---")
    }
}
```
## When to Use Nested Loops

- **2D/3D array processing**
- **Matrix operations**
- **Pattern generation**
- **Sorting algorithms**
- **Searching in multi-dimensional data**
- **Generating combinations or permutations**
- **Image processing**
- **Game development (grid-based games)**

# Infinite For Loops in Go

## Overview
Infinite loops are loops that run indefinitely without a natural termination condition. While they might seem dangerous, they are essential for many real-world applications like servers, event loops, and interactive programs that need to run continuously.

## Syntax for Infinite Loops

### Method 1: Empty Condition
```go
for {
    // This will run forever
    // statements
}
```

### Method 2: Always True Condition
```go
for true {
    // This will also run forever
    // statements
}
```

### Method 3: While-Style with Always True
```go
condition := true
for condition {
    // This will run forever
    // statements
}