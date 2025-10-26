# Go Variadic Parameters - Complete Guide

## Overview

Variadic parameters in Go allow functions to accept a variable number of arguments of the same type. The `...` operator is used to indicate that a function can accept zero or more arguments of the specified type. This feature provides flexibility when you don't know in advance how many arguments a function will receive.

## Key Features

1. **Variable arguments**: Accept zero or unlimited number of arguments
2. **Same type**: All arguments must be of the same type
3. **Slice parameter**: Variadic parameters are treated as slices inside the function
4. **Flexible calling**: Can pass individual arguments or a slice
5. **Zero arguments**: Functions can be called with no arguments

## Syntax

### Basic Syntax
```go
func functionName(parameter ...type) returnType {
    // function body
}
```

### With Other Parameters
```go
func functionName(regularParam type, variadicParam ...type) returnType {
    // function body
}
```

## Basic Examples

### Example 1: Simple Variadic Function

```go
package main

import "fmt"

func main() {
    // Pass multiple values
    result := sum(1, 2, 3, 4, 5)
    fmt.Println("Sum:", result)

    // Pass zero arguments
    emptyResult := sum()
    fmt.Println("Empty sum:", emptyResult)

    // Pass single argument
    singleResult := sum(10)
    fmt.Println("Single sum:", singleResult)
}

func sum(numbers ...int) int {
    fmt.Println("Numbers received:", numbers)
    fmt.Printf("Type of numbers: %T\n", numbers)
    fmt.Printf("Length of numbers: %d\n", len(numbers))

    total := 0
    for _, value := range numbers {
        total += value
    }
    return total
}
```

**Output:**
```
Numbers received: [1 2 3 4 5]
Type of numbers: []int
Length of numbers: 5
Sum: 15

Numbers received: []
Type of numbers: []int
Length of numbers: 0
Empty sum: 0

Numbers received: [10]
Type of numbers: []int
Length of numbers: 1
Single sum: 10
```

### Example 2: Variadic Function with Other Parameters

```go
package main

import "fmt"

func main() {
    // Pass string and multiple integers
    name := "Alice"
    numbers := []int{1, 2, 3, 4, 5, 6}
    result, values := processData(name, numbers...)
    fmt.Printf("Result: %s, Values: %v\n", result, values)

    // Pass string and individual integers
    result2, values2 := processData("Bob", 10, 20, 30)
    fmt.Printf("Result: %s, Values: %v\n", result2, values2)

    // Pass string and no integers
    result3, values3 := processData("Charlie")
    fmt.Printf("Result: %s, Values: %v\n", result3, values3)
}

func processData(name string, numbers ...int) (string, []int) {
    fmt.Printf("Processing data for: %s\n", name)
    fmt.Printf("Numbers: %v\n", numbers)
    fmt.Printf("Type: %T, Length: %d\n", numbers, len(numbers))

    // Process each number
    for i, num := range numbers {
        fmt.Printf("  [%d] = %d\n", i, num)
    }

    return name, numbers
}
```

**Output:**
```
Processing data for: Alice
Numbers: [1 2 3 4 5 6]
Type: []int, Length: 6
  [0] = 1
  [1] = 2
  [2] = 3
  [3] = 4
  [4] = 5
  [5] = 6
Result: Alice, Values: [1 2 3 4 5 6]

Processing data for: Bob
Numbers: [10 20 30]
Type: []int, Length: 3
  [0] = 10
  [1] = 20
  [2] = 30
Result: Bob, Values: [10 20 30]

Processing data for: Charlie
Numbers: []
Type: []int, Length: 0
Result: Charlie, Values: []
```

## Advanced Examples

### Example 3: Variadic Functions with Different Types

```go
package main

import "fmt"

func main() {
    // String variadic function
    result := concatenate("Hello", " ", "World", "!")
    fmt.Println("Concatenated:", result)

    // Float variadic function
    average := calculateAverage(1.5, 2.5, 3.5, 4.5, 5.5)
    fmt.Printf("Average: %.2f\n", average)

    // Interface variadic function
    printValues("String", 42, 3.14, true, []int{1, 2, 3})
}

func concatenate(strs ...string) string {
    result := ""
    for _, str := range strs {
        result += str
    }
    return result
}

func calculateAverage(numbers ...float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    return sum / float64(len(numbers))
}

func printValues(values ...interface{}) {
    fmt.Println("Values received:")
    for i, value := range values {
        fmt.Printf("  [%d] %T: %v\n", i, value, value)
    }
}
```

**Output:**
```
Concatenated: Hello World!
Average: 3.50
Values received:
  [0] string: String
  [1] int: 42
  [2] float64: 3.14
  [3] bool: true
  [4] []int: [1 2 3]
```

### Example 4: Variadic Functions with Slices

```go
package main

import "fmt"

func main() {
    // Create slices
    numbers1 := []int{1, 2, 3}
    numbers2 := []int{4, 5, 6}
    numbers3 := []int{7, 8, 9}

    // Pass slices to variadic function
    result := mergeSlices(numbers1, numbers2, numbers3)
    fmt.Printf("Merged slices: %v\n", result)

    // Pass individual numbers
    result2 := mergeSlices(10, 20, 30)
    fmt.Printf("Individual numbers: %v\n", result2)

    // Mix slices and individual numbers
    result3 := mergeSlices(numbers1..., 100, 200, numbers3...)
    fmt.Printf("Mixed: %v\n", result3)
}

func mergeSlices(numbers ...int) []int {
    fmt.Printf("Received %d numbers: %v\n", len(numbers), numbers)
    return numbers
}
```

**Output:**
```
Received 9 numbers: [1 2 3 4 5 6 7 8 9]
Merged slices: [1 2 3 4 5 6 7 8 9]
Received 3 numbers: [10 20 30]
Individual numbers: [10 20 30]
Received 9 numbers: [1 2 3 100 200 7 8 9]
Mixed: [1 2 3 100 200 7 8 9]
```

### Example 5: Variadic Functions for Logging

```go
package main

import (
    "fmt"
    "time"
)

type Logger struct {
    prefix string
}

func NewLogger(prefix string) *Logger {
    return &Logger{prefix: prefix}
}

func (l *Logger) Log(message string, args ...interface{}) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    formattedMessage := fmt.Sprintf(message, args...)
    fmt.Printf("[%s] %s: %s\n", timestamp, l.prefix, formattedMessage)
}

func (l *Logger) Error(message string, args ...interface{}) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    formattedMessage := fmt.Sprintf(message, args...)
    fmt.Printf("[%s] %s ERROR: %s\n", timestamp, l.prefix, formattedMessage)
}

func main() {
    logger := NewLogger("APP")

    // Log with no arguments
    logger.Log("Application started")

    // Log with arguments
    logger.Log("User %s logged in with ID %d", "Alice", 12345)

    // Log with multiple arguments
    logger.Log("Processing %d items in %s seconds", 100, 2.5)

    // Error logging
    logger.Error("Failed to connect to database: %s", "connection timeout")
}
```

**Output:**
```
[2024-01-15 10:30:45] APP: Application started
[2024-01-15 10:30:45] APP: User Alice logged in with ID 12345
[2024-01-15 10:30:45] APP: Processing 100 items in 2.5 seconds
[2024-01-15 10:30:45] APP ERROR: Failed to connect to database: connection timeout
```

## Practical Examples

### Example 6: Mathematical Operations

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Find maximum value
    max := findMax(1, 5, 3, 9, 2, 8, 4)
    fmt.Printf("Maximum: %d\n", max)

    // Find minimum value
    min := findMin(1, 5, 3, 9, 2, 8, 4)
    fmt.Printf("Minimum: %d\n", min)

    // Calculate standard deviation
    stdDev := calculateStandardDeviation(1.0, 2.0, 3.0, 4.0, 5.0)
    fmt.Printf("Standard deviation: %.2f\n", stdDev)

    // Check if all values are positive
    allPositive := allPositive(1, 2, 3, 4, 5)
    fmt.Printf("All positive: %t\n", allPositive)

    allPositive2 := allPositive(1, -2, 3, 4, 5)
    fmt.Printf("All positive: %t\n", allPositive2)
}

func findMax(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    max := numbers[0]
    for _, num := range numbers[1:] {
        if num > max {
            max = num
        }
    }
    return max
}

func findMin(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    min := numbers[0]
    for _, num := range numbers[1:] {
        if num < min {
            min = num
        }
    }
    return min
}

func calculateStandardDeviation(numbers ...float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    // Calculate mean
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    mean := sum / float64(len(numbers))
    
    // Calculate variance
    variance := 0.0
    for _, num := range numbers {
        variance += math.Pow(num-mean, 2)
    }
    variance /= float64(len(numbers))
    
    // Return standard deviation
    return math.Sqrt(variance)
}

func allPositive(numbers ...int) bool {
    for _, num := range numbers {
        if num <= 0 {
            return false
        }
    }
    return true
}
```

**Output:**
```
Maximum: 9
Minimum: 1
Standard deviation: 1.58
All positive: true
All positive: false
```

### Example 7: String Processing

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Join strings with separator
    result := joinStrings("-", "apple", "banana", "cherry")
    fmt.Printf("Joined: %s\n", result)

    // Find longest string
    longest := findLongest("short", "medium length", "very long string", "tiny")
    fmt.Printf("Longest: %s\n", longest)

    // Check if all strings contain substring
    allContain := allContain("go", "golang", "google", "goroutine")
    fmt.Printf("All contain 'go': %t\n", allContain)

    allContain2 := allContain("go", "golang", "python", "goroutine")
    fmt.Printf("All contain 'go': %t\n", allContain2)

    // Count total characters
    totalChars := countTotalCharacters("hello", "world", "golang")
    fmt.Printf("Total characters: %d\n", totalChars)
}

func joinStrings(separator string, strings ...string) string {
    return strings.Join(strings, separator)
}

func findLongest(strings ...string) string {
    if len(strings) == 0 {
        return ""
    }
    
    longest := strings[0]
    for _, str := range strings[1:] {
        if len(str) > len(longest) {
            longest = str
        }
    }
    return longest
}

func allContain(substring string, strings ...string) bool {
    for _, str := range strings {
        if !strings.Contains(str, substring) {
            return false
        }
    }
    return true
}

func countTotalCharacters(strings ...string) int {
    total := 0
    for _, str := range strings {
        total += len(str)
    }
    return total
}
```

**Output:**
```
Joined: apple-banana-cherry
Longest: very long string
All contain 'go': true
All contain 'go': false
Total characters: 18
```

## Important Rules and Best Practices

### 1. Variadic Parameter Rules
- Only one variadic parameter per function
- Must be the last parameter
- All arguments must be of the same type
- Can accept zero or more arguments

### 2. Calling Variadic Functions
```go
// Pass individual arguments
result := sum(1, 2, 3, 4, 5)

// Pass no arguments
result := sum()

// Pass slice using spread operator
numbers := []int{1, 2, 3, 4, 5}
result := sum(numbers...)

// Mix individual arguments and slice
result := sum(1, 2, numbers..., 6, 7)
```

### 3. Common Patterns

#### Pattern 1: Optional Parameters
```go
func createUser(name string, options ...string) {
    email := "default@example.com"
    phone := "000-000-0000"
    
    for i := 0; i < len(options); i += 2 {
        if i+1 < len(options) {
            switch options[i] {
            case "email":
                email = options[i+1]
            case "phone":
                phone = options[i+1]
            }
        }
    }
    
    fmt.Printf("User: %s, Email: %s, Phone: %s\n", name, email, phone)
}
```

#### Pattern 2: Builder Pattern
```go
type QueryBuilder struct {
    table string
    conditions []string
}

func NewQueryBuilder(table string) *QueryBuilder {
    return &QueryBuilder{table: table}
}

func (qb *QueryBuilder) Where(conditions ...string) *QueryBuilder {
    qb.conditions = append(qb.conditions, conditions...)
    return qb
}

func (qb *QueryBuilder) Build() string {
    query := "SELECT * FROM " + qb.table
    if len(qb.conditions) > 0 {
        query += " WHERE " + strings.Join(qb.conditions, " AND ")
    }
    return query
}
```

## Performance Considerations

### Example 8: Performance Comparison

```go
package main

import (
    "fmt"
    "time"
)

// Variadic function
func sumVariadic(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Slice function
func sumSlice(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    // Test data
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Benchmark variadic function
    start := time.Now()
    for i := 0; i < 1000000; i++ {
        sumVariadic(numbers...)
    }
    variadicTime := time.Since(start)
    
    // Benchmark slice function
    start = time.Now()
    for i := 0; i < 1000000; i++ {
        sumSlice(numbers)
    }
    sliceTime := time.Since(start)
    
    fmt.Printf("Variadic function time: %v\n", variadicTime)
    fmt.Printf("Slice function time: %v\n", sliceTime)
    fmt.Printf("Performance difference: %.2fx\n", float64(variadicTime)/float64(sliceTime))
}
```

## Common Use Cases

1. **Mathematical operations**: Sum, average, min, max
2. **String processing**: Concatenation, formatting
3. **Logging**: Flexible logging with multiple arguments
4. **Configuration**: Optional parameters
5. **Data processing**: Flexible data handling
6. **API design**: Flexible function interfaces

## Conclusion

Variadic parameters in Go provide:

1. **Flexibility**: Accept variable number of arguments
2. **Simplicity**: Easy to use and understand
3. **Performance**: Efficient implementation
4. **Type safety**: All arguments must be same type
5. **Versatility**: Can be used in many scenarios

Understanding variadic parameters is essential for writing flexible and reusable Go functions, especially when dealing with functions that need to handle varying numbers of arguments.
