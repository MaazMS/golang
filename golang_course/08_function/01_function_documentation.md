# Go Functions - Complete Guide

## Overview

Functions in Go are first-class citizens that allow you to group code into reusable blocks. They are essential for code organization, reusability, and maintainability. Go functions support multiple return values, variadic parameters, closures, and can be assigned to variables or passed as arguments.

## Key Features

1. **Code reusability**: Write once, use many times
2. **Code maintainability**: Easier to debug and modify
3. **First-class citizens**: Can be assigned to variables and passed as arguments
4. **Multiple return values**: Return multiple values from a single function
5. **Variadic parameters**: Accept variable number of arguments
6. **Closures**: Functions that capture variables from their lexical scope
7. **Pass by value**: All parameters are passed by value by default

## Why Use Functions?

### Benefits
- **Code reuse**: Avoid duplicating code
- **Modularity**: Break complex problems into smaller parts
- **Maintainability**: Easier to update and debug
- **Testability**: Test individual functions in isolation
- **Readability**: Make code more readable and organized

### When to Use Functions
- Repeated code blocks
- Complex calculations
- Data processing
- API endpoints
- Utility operations
- Business logic

## Basic Syntax

### Function Declaration

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Key Components
- **func**: Keyword to declare a function
- **functionName**: Name of the function (must start with uppercase to be exported)
- **parameters**: Input values (optional)
- **returnType**: Type of return value (optional)
- **function body**: Code to execute

## Basic Examples

### Example 1: Simple Functions

```go
package main

import "fmt"

// Function with no parameters and no return value
func sayHello() {
    fmt.Println("Hello, World!")
}

// Function with parameters and return value
func add(a int, b int) int {
    return a + b
}

// Function with multiple parameters of same type
func multiply(x, y int) int {
    return x * y
}

// Function with multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    // Call functions
    sayHello()
    
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    product := multiply(4, 6)
    fmt.Printf("4 * 6 = %d\n", product)
    
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("10 / 2 = %d\n", quotient)
    }
}
```

### Example 2: Named Return Values

```go
package main

import "fmt"

// Named return values
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // naked return
}

// Multiple named return values
func getPersonInfo() (name string, age int, isActive bool) {
    name = "Alice"
    age = 30
    isActive = true
    return
}

func main() {
    s, p := calculate(5, 3)
    fmt.Printf("Sum: %d, Product: %d\n", s, p)
    
    name, age, active := getPersonInfo()
    fmt.Printf("Name: %s, Age: %d, Active: %t\n", name, age, active)
}
```

### Example 3: Variadic Functions

```go
package main

import "fmt"

// Variadic function - accepts variable number of arguments
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Variadic function with other parameters
func printInfo(prefix string, values ...interface{}) {
    fmt.Printf("%s: ", prefix)
    for i, v := range values {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(v)
    }
    fmt.Println()
}

// Variadic function with slice parameter
func processSlice(items ...string) {
    fmt.Printf("Processing %d items: %v\n", len(items), items)
}

func main() {
    // Call with different number of arguments
    fmt.Printf("Sum of 1, 2, 3: %d\n", sum(1, 2, 3))
    fmt.Printf("Sum of 1, 2, 3, 4, 5: %d\n", sum(1, 2, 3, 4, 5))
    fmt.Printf("Sum of no arguments: %d\n", sum())
    
    // Print info with variadic parameters
    printInfo("Numbers", 1, 2, 3, 4, 5)
    printInfo("Mixed", "hello", 42, true, 3.14)
    
    // Process slice
    processSlice("apple", "banana", "cherry")
}
```

## Advanced Examples

### Example 4: Function as Values

```go
package main

import "fmt"

// Function type
type MathOperation func(int, int) int

func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func subtract(a, b int) int {
    return a - b
}

// Function that takes another function as parameter
func calculate(a, b int, operation MathOperation) int {
    return operation(a, b)
}

// Function that returns a function
func getOperation(op string) MathOperation {
    switch op {
    case "add":
        return add
    case "multiply":
        return multiply
    case "subtract":
        return subtract
    default:
        return add
    }
}

func main() {
    // Assign function to variable
    var operation MathOperation = add
    result := operation(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    // Pass function as argument
    sum := calculate(10, 5, add)
    product := calculate(10, 5, multiply)
    difference := calculate(10, 5, subtract)
    
    fmt.Printf("10 + 5 = %d\n", sum)
    fmt.Printf("10 * 5 = %d\n", product)
    fmt.Printf("10 - 5 = %d\n", difference)
    
    // Get function from another function
    addFunc := getOperation("add")
    multiplyFunc := getOperation("multiply")
    
    fmt.Printf("Using add function: %d\n", addFunc(7, 3))
    fmt.Printf("Using multiply function: %d\n", multiplyFunc(7, 3))
}
```

### Example 5: Closures

```go
package main

import "fmt"

// Function that returns a closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Function that returns a closure with parameter
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Function that returns multiple closures
func createCounters() (func() int, func() int) {
    count1 := 0
    count2 := 0
    
    counter1 := func() int {
        count1++
        return count1
    }
    
    counter2 := func() int {
        count2++
        return count2
    }
    
    return counter1, counter2
}

func main() {
    // Create counter closure
    c1 := counter()
    c2 := counter()
    
    fmt.Printf("Counter 1: %d\n", c1()) // 1
    fmt.Printf("Counter 1: %d\n", c1()) // 2
    fmt.Printf("Counter 2: %d\n", c2()) // 1
    fmt.Printf("Counter 1: %d\n", c1()) // 3
    
    // Create multiplier closure
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Printf("Double 5: %d\n", double(5))   // 10
    fmt.Printf("Triple 5: %d\n", triple(5))   // 15
    
    // Create multiple counters
    counter1, counter2 := createCounters()
    fmt.Printf("Counter 1: %d\n", counter1()) // 1
    fmt.Printf("Counter 2: %d\n", counter2()) // 1
    fmt.Printf("Counter 1: %d\n", counter1()) // 2
    fmt.Printf("Counter 2: %d\n", counter2()) // 2
}
```

### Example 6: Recursive Functions

```go
package main

import "fmt"

// Factorial using recursion
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// Fibonacci using recursion
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// Binary search using recursion
func binarySearch(arr []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    mid := (left + right) / 2
    if arr[mid] == target {
        return mid
    } else if arr[mid] > target {
        return binarySearch(arr, target, left, mid-1)
    } else {
        return binarySearch(arr, target, mid+1, right)
    }
}

// Tree traversal using recursion
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    result := inorderTraversal(root.Left)
    result = append(result, root.Value)
    result = append(result, inorderTraversal(root.Right)...)
    return result
}

func main() {
    // Factorial
    fmt.Printf("Factorial of 5: %d\n", factorial(5))
    
    // Fibonacci
    fmt.Printf("Fibonacci sequence: ")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
    fmt.Println()
    
    // Binary search
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    target := 7
    index := binarySearch(arr, target, 0, len(arr)-1)
    if index != -1 {
        fmt.Printf("Found %d at index %d\n", target, index)
    } else {
        fmt.Printf("%d not found\n", target)
    }
    
    // Tree traversal
    root := &TreeNode{
        Value: 1,
        Left: &TreeNode{
            Value: 2,
            Left:  &TreeNode{Value: 4},
            Right: &TreeNode{Value: 5},
        },
        Right: &TreeNode{
            Value: 3,
            Left:  &TreeNode{Value: 6},
            Right: &TreeNode{Value: 7},
        },
    }
    
    result := inorderTraversal(root)
    fmt.Printf("Inorder traversal: %v\n", result)
}
```

## Function Types and Patterns

### Example 7: Function Types and Interfaces

```go
package main

import "fmt"

// Function type
type Handler func(string) string

// Interface with function
type Processor interface {
    Process(string) string
}

// Function that implements Processor interface
type HandlerFunc func(string) string

func (f HandlerFunc) Process(s string) string {
    return f(s)
}

// Function that takes function type
func applyHandler(s string, handler Handler) string {
    return handler(s)
}

// Function that returns function type
func createHandler(prefix string) Handler {
    return func(s string) string {
        return prefix + ": " + s
    }
}

func main() {
    // Use function type
    upperHandler := func(s string) string {
        return "UPPER: " + s
    }
    
    result := applyHandler("hello", upperHandler)
    fmt.Println(result)
    
    // Use function that returns function
    errorHandler := createHandler("ERROR")
    infoHandler := createHandler("INFO")
    
    fmt.Println(errorHandler("Something went wrong"))
    fmt.Println(infoHandler("Process completed"))
    
    // Use function as interface
    var processor Processor = HandlerFunc(func(s string) string {
        return "Processed: " + s
    })
    
    fmt.Println(processor.Process("data"))
}
```

### Example 8: Defer and Panic/Recover

```go
package main

import (
    "fmt"
    "os"
)

// Function with defer
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Will be called when function returns
    
    // Process file
    fmt.Printf("Processing file: %s\n", filename)
    return nil
}

// Function with panic and recover
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    fmt.Println("Starting risky operation")
    panic("Something went wrong!")
    fmt.Println("This won't be printed")
}

// Function with multiple defers
func multipleDefers() {
    fmt.Println("Function start")
    
    defer fmt.Println("Defer 1")
    defer fmt.Println("Defer 2")
    defer fmt.Println("Defer 3")
    
    fmt.Println("Function middle")
    fmt.Println("Function end")
}

// Function with defer and return value
func deferWithReturn() (result int) {
    defer func() {
        result++ // This will modify the return value
    }()
    
    return 5
}

func main() {
    // Test defer
    err := processFile("nonexistent.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Test panic and recover
    riskyOperation()
    fmt.Println("Program continues after panic")
    
    // Test multiple defers
    multipleDefers()
    
    // Test defer with return value
    result := deferWithReturn()
    fmt.Printf("Defer with return: %d\n", result)
}
```

## Practical Examples

### Example 9: Utility Functions

```go
package main

import (
    "fmt"
    "math"
    "strings"
    "time"
)

// String utilities
func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isPalindrome(s string) bool {
    cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
    return cleaned == reverseString(cleaned)
}

// Math utilities
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Time utilities
func formatDuration(d time.Duration) string {
    if d < time.Minute {
        return fmt.Sprintf("%.0f seconds", d.Seconds())
    } else if d < time.Hour {
        return fmt.Sprintf("%.0f minutes", d.Minutes())
    } else {
        return fmt.Sprintf("%.0f hours", d.Hours())
    }
}

func main() {
    // String utilities
    text := "A man a plan a canal Panama"
    fmt.Printf("Is '%s' a palindrome? %t\n", text, isPalindrome(text))
    
    // Math utilities
    fmt.Printf("Is 17 prime? %t\n", isPrime(17))
    fmt.Printf("GCD of 48 and 18: %d\n", gcd(48, 18))
    
    // Time utilities
    duration := 2*time.Hour + 30*time.Minute + 45*time.Second
    fmt.Printf("Duration: %s\n", formatDuration(duration))
}
```

### Example 10: API and HTTP Functions

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// User struct
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// API response struct
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Error   string      `json:"error,omitempty"`
}

// Function to create API response
func createResponse(success bool, data interface{}, err error) APIResponse {
    response := APIResponse{
        Success: success,
        Data:    data,
    }
    if err != nil {
        response.Error = err.Error()
    }
    return response
}

// Function to handle JSON encoding
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    return json.NewEncoder(w).Encode(data)
}

// Function to validate user input
func validateUser(user User) error {
    if user.Name == "" {
        return fmt.Errorf("name is required")
    }
    if user.Email == "" {
        return fmt.Errorf("email is required")
    }
    if !strings.Contains(user.Email, "@") {
        return fmt.Errorf("invalid email format")
    }
    return nil
}

// Function to process user data
func processUser(user User) (User, error) {
    if err := validateUser(user); err != nil {
        return User{}, err
    }
    
    // Simulate processing
    user.Name = strings.Title(strings.ToLower(user.Name))
    user.Email = strings.ToLower(user.Email)
    
    return user, nil
}

func main() {
    // Test user processing
    user := User{
        ID:    1,
        Name:  "john doe",
        Email: "JOHN@EXAMPLE.COM",
    }
    
    processedUser, err := processUser(user)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Processed user: %+v\n", processedUser)
    
    // Test API response
    response := createResponse(true, processedUser, nil)
    jsonData, _ := json.MarshalIndent(response, "", "  ")
    fmt.Printf("API Response:\n%s\n", string(jsonData))
}
```

## Best Practices

### 1. Function Naming
```go
// Good: Clear, descriptive names
func calculateTotalPrice(items []Item) float64
func validateEmailAddress(email string) bool
func getUserByID(id int) (*User, error)

// Bad: Unclear names
func calc(items []Item) float64
func check(email string) bool
func get(id int) (*User, error)
```

### 2. Parameter Design
```go
// Good: Few parameters, clear types
func processUser(name, email string) error

// Bad: Too many parameters
func processUser(name, email, phone, address, city, state, zip string) error

// Better: Use struct for many parameters
type UserData struct {
    Name, Email, Phone, Address, City, State, Zip string
}
func processUser(data UserData) error
```

### 3. Return Values
```go
// Good: Clear return types
func divide(a, b int) (float64, error)

// Good: Named return values for clarity
func parseConfig(filename string) (config Config, err error)

// Bad: Unclear return types
func process(data interface{}) interface{}
```

### 4. Error Handling
```go
// Good: Always handle errors
result, err := riskyOperation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Good: Use custom error types
type ValidationError struct {
    Field string
    Value string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("invalid %s: %s", e.Field, e.Value)
}
```

## Performance Considerations

### Example 11: Function Performance

```go
package main

import (
    "fmt"
    "time"
)

// Function with value receiver
func processByValue(data [1000]int) int {
    sum := 0
    for _, v := range data {
        sum += v
    }
    return sum
}

// Function with pointer receiver
func processByPointer(data *[1000]int) int {
    sum := 0
    for _, v := range data {
        sum += v
    }
    return sum
}

// Function with slice parameter
func processSlice(data []int) int {
    sum := 0
    for _, v := range data {
        sum += v
    }
    return sum
}

func main() {
    // Create test data
    var arr [1000]int
    for i := range arr {
        arr[i] = i
    }
    
    slice := arr[:]
    
    // Benchmark value parameter
    start := time.Now()
    for i := 0; i < 100000; i++ {
        processByValue(arr)
    }
    valueTime := time.Since(start)
    
    // Benchmark pointer parameter
    start = time.Now()
    for i := 0; i < 100000; i++ {
        processByPointer(&arr)
    }
    pointerTime := time.Since(start)
    
    // Benchmark slice parameter
    start = time.Now()
    for i := 0; i < 100000; i++ {
        processSlice(slice)
    }
    sliceTime := time.Since(start)
    
    fmt.Printf("Value parameter time: %v\n", valueTime)
    fmt.Printf("Pointer parameter time: %v\n", pointerTime)
    fmt.Printf("Slice parameter time: %v\n", sliceTime)
}
```

## Conclusion

Functions in Go are powerful and flexible tools that provide:

1. **Code organization**: Break complex problems into manageable pieces
2. **Reusability**: Write once, use many times
3. **Maintainability**: Easier to debug and update
4. **Flexibility**: First-class citizens with multiple patterns
5. **Performance**: Efficient execution with pass-by-value semantics

Understanding functions is essential for effective Go programming, as they are the building blocks of Go applications and provide the foundation for creating maintainable, reusable code.
