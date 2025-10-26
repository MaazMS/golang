# Go Anonymous Functions - Complete Guide

## Overview

Anonymous functions in Go are functions that have no name and are defined inline. They are also known as lambda functions or closures. Anonymous functions can be self-executing (immediately invoked) or assigned to variables for later use. They are particularly useful for short, one-time operations and for creating closures that capture variables from their lexical scope.

## Key Features

1. **No name**: Functions defined without a name
2. **Self-executing**: Can be called immediately after definition
3. **Closures**: Can capture variables from surrounding scope
4. **Inline definition**: Defined where they are used
5. **Flexible usage**: Can be assigned to variables or passed as arguments
6. **Lexical scoping**: Access variables from enclosing scope

## Why Use Anonymous Functions?

### Benefits
- **Conciseness**: Define functions inline without naming
- **Closures**: Capture variables from surrounding scope
- **Immediate execution**: Execute code right after definition
- **Flexibility**: Can be passed as arguments or assigned to variables
- **Encapsulation**: Keep related code together

### When to Use Anonymous Functions
- Short, one-time operations
- Creating closures
- Callback functions
- Event handlers
- Immediate execution scenarios
- Functional programming patterns

## Basic Syntax

### Anonymous Function Declaration
```go
// Self-executing anonymous function
func() {
    // function body
}()

// Anonymous function with parameters
func(param1 type1, param2 type2) {
    // function body
}(arg1, arg2)

// Anonymous function assigned to variable
variable := func() {
    // function body
}
```

## Basic Examples

### Example 1: Simple Anonymous Functions

```go
package main

import "fmt"

func main() {
    // Anonymous function with no parameters
    func() {
        fmt.Println("Anonymous function with no parameters")
    }()
    
    // Anonymous function with parameters
    func(x int) {
        fmt.Printf("Anonymous function with parameter: %d\n", x)
    }(45)
    
    // Anonymous function with multiple parameters
    func(name string, age int) {
        fmt.Printf("Name: %s, Age: %d\n", name, age)
    }("Alice", 30)
    
    // Anonymous function with return value
    result := func(a, b int) int {
        return a + b
    }(10, 20)
    fmt.Printf("Result: %d\n", result)
}
```

**Output:**
```
Anonymous function with no parameters
Anonymous function with parameter: 45
Name: Alice, Age: 30
Result: 30
```

### Example 2: Anonymous Functions Assigned to Variables

```go
package main

import "fmt"

func main() {
    // Assign anonymous function to variable
    greet := func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }
    
    // Call the function
    greet("Alice")
    greet("Bob")
    
    // Anonymous function that returns a function
    multiplier := func(factor int) func(int) int {
        return func(x int) int {
            return x * factor
        }
    }
    
    // Create specific multipliers
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Printf("Double of 5: %d\n", double(5))
    fmt.Printf("Triple of 5: %d\n", triple(5))
    
    // Anonymous function with multiple return values
    calculator := func(a, b int) (int, int, int) {
        return a + b, a - b, a * b
    }
    
    sum, diff, product := calculator(10, 5)
    fmt.Printf("Sum: %d, Difference: %d, Product: %d\n", sum, diff, product)
}
```

**Output:**
```
Hello, Alice!
Hello, Bob!
Double of 5: 10
Triple of 5: 15
Sum: 15, Difference: 5, Product: 50
```

## Advanced Examples

### Example 3: Closures and Variable Capture

```go
package main

import "fmt"

func main() {
    // Counter closure
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()
    
    fmt.Printf("Count: %d\n", counter())
    fmt.Printf("Count: %d\n", counter())
    fmt.Printf("Count: %d\n", counter())
    
    // Multiple counters
    counter1 := createCounter()
    counter2 := createCounter()
    
    fmt.Printf("Counter1: %d\n", counter1())
    fmt.Printf("Counter1: %d\n", counter1())
    fmt.Printf("Counter2: %d\n", counter2())
    fmt.Printf("Counter1: %d\n", counter1())
    
    // Closure with parameters
    accumulator := func(initial int) func(int) int {
        sum := initial
        return func(value int) int {
            sum += value
            return sum
        }
    }(10)
    
    fmt.Printf("Accumulator: %d\n", accumulator(5))
    fmt.Printf("Accumulator: %d\n", accumulator(3))
    fmt.Printf("Accumulator: %d\n", accumulator(7))
}

// Function that returns a closure
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

**Output:**
```
Count: 1
Count: 2
Count: 3
Counter1: 1
Counter1: 2
Counter2: 1
Counter1: 3
Accumulator: 15
Accumulator: 18
Accumulator: 25
```

### Example 4: Anonymous Functions in Collections

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter even numbers using anonymous function
    evenNumbers := filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Printf("Even numbers: %v\n", evenNumbers)
    
    // Filter odd numbers using anonymous function
    oddNumbers := filter(numbers, func(n int) bool {
        return n%2 != 0
    })
    fmt.Printf("Odd numbers: %v\n", oddNumbers)
    
    // Map numbers to their squares using anonymous function
    squares := mapNumbers(numbers, func(n int) int {
        return n * n
    })
    fmt.Printf("Squares: %v\n", squares)
    
    // Map numbers to strings using anonymous function
    numberStrings := mapNumbersToStrings(numbers, func(n int) string {
        return fmt.Sprintf("Number: %d", n)
    })
    fmt.Printf("Number strings: %v\n", numberStrings)
}

// Generic filter function
func filter(numbers []int, predicate func(int) bool) []int {
    var result []int
    for _, n := range numbers {
        if predicate(n) {
            result = append(result, n)
        }
    }
    return result
}

// Generic map function
func mapNumbers(numbers []int, mapper func(int) int) []int {
    var result []int
    for _, n := range numbers {
        result = append(result, mapper(n))
    }
    return result
}

// Generic map function for strings
func mapNumbersToStrings(numbers []int, mapper func(int) string) []string {
    var result []string
    for _, n := range numbers {
        result = append(result, mapper(n))
    }
    return result
}
```

**Output:**
```
Even numbers: [2 4 6 8 10]
Odd numbers: [1 3 5 7 9]
Squares: [1 4 9 16 25 36 49 64 81 100]
Number strings: [Number: 1 Number: 2 Number: 3 Number: 4 Number: 5 Number: 6 Number: 7 Number: 8 Number: 9 Number: 10]
```

### Example 5: Anonymous Functions with Error Handling

```go
package main

import (
    "errors"
    "fmt"
    "strconv"
)

func main() {
    // Anonymous function for safe division
    safeDivide := func(a, b float64) (float64, error) {
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    }
    
    result, err := safeDivide(10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }
    
    result, err = safeDivide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("10 / 0 = %.2f\n", result)
    }
    
    // Anonymous function for string to int conversion
    parseInt := func(s string) (int, error) {
        return strconv.Atoi(s)
    }
    
    numbers := []string{"123", "456", "abc", "789"}
    for _, numStr := range numbers {
        if num, err := parseInt(numStr); err != nil {
            fmt.Printf("Failed to parse '%s': %v\n", numStr, err)
        } else {
            fmt.Printf("Parsed '%s' as %d\n", numStr, num)
        }
    }
    
    // Anonymous function for validation
    validateEmail := func(email string) error {
        if len(email) == 0 {
            return errors.New("email cannot be empty")
        }
        if !contains(email, "@") {
            return errors.New("email must contain @")
        }
        return nil
    }
    
    emails := []string{"alice@example.com", "bob@", "", "invalid-email"}
    for _, email := range emails {
        if err := validateEmail(email); err != nil {
            fmt.Printf("Invalid email '%s': %v\n", email, err)
        } else {
            fmt.Printf("Valid email: %s\n", email)
        }
    }
}

// Helper function
func contains(s, substr string) bool {
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}
```

**Output:**
```
10 / 2 = 5.00
Error: division by zero
Parsed '123' as 123
Parsed '456' as 456
Failed to parse 'abc': strconv.Atoi: parsing "abc": invalid syntax
Parsed '789' as 789
Valid email: alice@example.com
Invalid email 'bob@': email must contain @
Invalid email '': email cannot be empty
Invalid email 'invalid-email': email must contain @
```

## Practical Examples

### Example 6: Anonymous Functions in HTTP Handlers

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // HTTP handler using anonymous function
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from anonymous function!")
    })
    
    // HTTP handler with logging using anonymous function
    http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
        fmt.Fprintf(w, "Request logged!")
    })
    
    // HTTP handler with error handling using anonymous function
    http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Recovered from panic: %v\n", r)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        
        // Simulate some processing that might panic
        panic("Something went wrong!")
    })
    
    fmt.Println("Server starting on :8080")
    // Note: In a real application, you would use http.ListenAndServe(":8080", nil)
    fmt.Println("Server would start here...")
}
```

### Example 7: Anonymous Functions for Sorting

```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
        {"Diana", 28},
    }
    
    fmt.Println("Original order:")
    printPeople(people)
    
    // Sort by age using anonymous function
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    fmt.Println("\nSorted by age:")
    printPeople(people)
    
    // Sort by name using anonymous function
    sort.Slice(people, func(i, j int) bool {
        return people[i].Name < people[j].Name
    })
    fmt.Println("\nSorted by name:")
    printPeople(people)
    
    // Sort by age descending using anonymous function
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age > people[j].Age
    })
    fmt.Println("\nSorted by age (descending):")
    printPeople(people)
}

func printPeople(people []Person) {
    for _, person := range people {
        fmt.Printf("  %s (%d years old)\n", person.Name, person.Age)
    }
}
```

**Output:**
```
Original order:
  Alice (30 years old)
  Bob (25 years old)
  Charlie (35 years old)
  Diana (28 years old)

Sorted by age:
  Bob (25 years old)
  Diana (28 years old)
  Alice (30 years old)
  Charlie (35 years old)

Sorted by name:
  Alice (30 years old)
  Bob (25 years old)
  Charlie (35 years old)
  Diana (28 years old)

Sorted by age (descending):
  Charlie (35 years old)
  Alice (30 years old)
  Diana (28 years old)
  Bob (25 years old)
```

### Example 8: Anonymous Functions for Timing and Profiling

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Anonymous function for timing operations
    timeOperation := func(name string, operation func()) {
        start := time.Now()
        operation()
        duration := time.Since(start)
        fmt.Printf("%s took %v\n", name, duration)
    }
    
    // Time different operations
    timeOperation("Sleep 100ms", func() {
        time.Sleep(100 * time.Millisecond)
    })
    
    timeOperation("Sleep 200ms", func() {
        time.Sleep(200 * time.Millisecond)
    })
    
    // Time calculation
    timeOperation("Calculate sum", func() {
        sum := 0
        for i := 0; i < 1000000; i++ {
            sum += i
        }
        fmt.Printf("Sum: %d\n", sum)
    })
    
    // Anonymous function for retry logic
    retryOperation := func(maxRetries int, operation func() error) error {
        for i := 0; i < maxRetries; i++ {
            if err := operation(); err != nil {
                fmt.Printf("Attempt %d failed: %v\n", i+1, err)
                if i == maxRetries-1 {
                    return fmt.Errorf("operation failed after %d attempts", maxRetries)
                }
                time.Sleep(100 * time.Millisecond)
            } else {
                fmt.Printf("Operation succeeded on attempt %d\n", i+1)
                return nil
            }
        }
        return nil
    }
    
    // Simulate retry operation
    retryOperation(3, func() error {
        // Simulate random failure
        if time.Now().UnixNano()%2 == 0 {
            return fmt.Errorf("random failure")
        }
        return nil
    })
}
```

**Output:**
```
Sleep 100ms took 100.123ms
Sleep 200ms took 200.456ms
Sum: 499999500000
Calculate sum took 2.345ms
Attempt 1 failed: random failure
Attempt 2 failed: random failure
Operation succeeded on attempt 3
```

## Common Patterns and Best Practices

### Pattern 1: Immediate Execution
```go
// Execute immediately after definition
func() {
    fmt.Println("Immediate execution")
}()

// With parameters
func(x int) {
    fmt.Printf("Value: %d\n", x)
}(42)
```

### Pattern 2: Assignment to Variables
```go
// Assign to variable for later use
greet := func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

greet("Alice")
greet("Bob")
```

### Pattern 3: Return from Functions
```go
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
```

### Pattern 4: Closures for State
```go
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Important Rules and Considerations

### 1. Variable Capture
- Anonymous functions capture variables by reference
- Variables must be in scope when the function is defined
- Captured variables persist for the lifetime of the function

### 2. Memory Management
- Closures can cause memory leaks if not handled properly
- Be careful with large captured variables
- Consider the lifetime of captured variables

### 3. Performance
- Anonymous functions have minimal overhead
- Use them judiciously in performance-critical code
- Consider function inlining for simple cases

## Common Mistakes to Avoid

### Mistake 1: Capturing Loop Variables
```go
// ❌ Bad: All functions will use the final value of i
var funcs []func() int
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() int {
        return i // This will always return 3
    })
}

// ✅ Good: Capture the value properly
var funcs []func() int
for i := 0; i < 3; i++ {
    i := i // Create a new variable
    funcs = append(funcs, func() int {
        return i // This will return the correct value
    })
}
```

### Mistake 2: Memory Leaks
```go
// ❌ Bad: Large variable captured in closure
func badExample() func() {
    largeData := make([]byte, 1024*1024) // 1MB
    return func() {
        fmt.Println("Using large data")
        // largeData is kept in memory even if not used
    }
}

// ✅ Good: Don't capture large variables unnecessarily
func goodExample() func() {
    return func() {
        fmt.Println("No large data captured")
    }
}
```

### Mistake 3: Overusing Anonymous Functions
```go
// ❌ Bad: Complex anonymous function
func badExample() {
    result := func(a, b, c int) int {
        // Complex logic here
        if a > b {
            if c > 0 {
                return a + c
            } else {
                return a - c
            }
        } else {
            if c > 0 {
                return b + c
            } else {
                return b - c
            }
        }
    }(1, 2, 3)
    fmt.Println(result)
}

// ✅ Good: Extract to named function
func goodExample() {
    result := calculate(1, 2, 3)
    fmt.Println(result)
}

func calculate(a, b, c int) int {
    if a > b {
        if c > 0 {
            return a + c
        } else {
            return a - c
        }
    } else {
        if c > 0 {
            return b + c
        } else {
            return b - c
        }
    }
}
```

## Conclusion

Anonymous functions in Go provide:

1. **Conciseness**: Define functions inline without naming
2. **Closures**: Capture variables from surrounding scope
3. **Flexibility**: Can be assigned, passed, or executed immediately
4. **Functional programming**: Enable functional programming patterns
5. **Event handling**: Perfect for callbacks and event handlers

Understanding anonymous functions is essential for writing concise, functional Go code. They are particularly useful for short operations, closures, and functional programming patterns while maintaining readability and performance.
