# Type Conversion in Go (Not Casting)

## Introduction

Go is a **statically typed** programming language, which means that types are determined at compile time. Because of this, types are extremely important in Go programming. When you need to change a value from one type to another, you must use **explicit type conversion** (not casting, as in other languages).

## Understanding Type Conversion

### What is Type Conversion?

Type conversion in Go is the process of explicitly converting a value from one type to another compatible type. Unlike some languages that use "casting," Go requires explicit conversion syntax.

### Key Points About Type Conversion

1. **Explicit Conversion Required**: Go requires explicit type conversion - you cannot implicitly convert types
2. **Type Compatibility**: Both types must be compatible for conversion to work
3. **Compile-time Safety**: Type conversion is checked at compile time
4. **Syntax**: Use `Type(value)` syntax for conversion

## Basic Type Conversion Example

```go
package main

import (
    "fmt"
)

// Declare variable no with TYPE int
var no int

// Create our own TYPE palindrome with underlying type int
type palindrome int

// Create variable b with TYPE palindrome
var b palindrome

func main() {
    no = 100
    fmt.Println("Original int:", no)
    fmt.Printf("Type of no: %T\n", no)

    b = 500
    fmt.Println("Custom palindrome:", b)
    fmt.Printf("Type of b: %T\n", b)

    // Type conversion: no = b
    // This would cause a compile error:
    // cannot use b (type palindrome) as type int in assignment
    
    // Correct way - explicit conversion
    // int(b) because variable no is declared as TYPE int
    no = int(b)
    fmt.Println("After conversion:", no)
    fmt.Printf("Type of no after conversion: %T\n", no)
}
```

## Comprehensive Type Conversion Examples

### 1. Numeric Type Conversions

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = 3.14
    var b byte = 255
    
    // int to float64
    var floatFromInt float64 = float64(i)
    fmt.Printf("int %d -> float64 %.2f\n", i, floatFromInt)
    
    // float64 to int (truncates decimal part)
    var intFromFloat int = int(f)
    fmt.Printf("float64 %.2f -> int %d\n", f, intFromFloat)
    
    // int to byte
    var byteFromInt byte = byte(i)
    fmt.Printf("int %d -> byte %d\n", i, byteFromInt)
    
    // byte to int
    var intFromByte int = int(b)
    fmt.Printf("byte %d -> int %d\n", b, intFromByte)
}
```

### 2. String Conversions

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num int = 123
    var floatNum float64 = 45.67
    var boolVal bool = true
    
    // Number to string using strconv package
    strFromInt := strconv.Itoa(num)
    fmt.Printf("int %d -> string '%s'\n", num, strFromInt)
    
    // Float to string
    strFromFloat := strconv.FormatFloat(floatNum, 'f', 2, 64)
    fmt.Printf("float64 %.2f -> string '%s'\n", floatNum, strFromFloat)
    
    // Bool to string
    strFromBool := strconv.FormatBool(boolVal)
    fmt.Printf("bool %t -> string '%s'\n", boolVal, strFromBool)
    
    // String to number
    str := "456"
    if intFromStr, err := strconv.Atoi(str); err == nil {
        fmt.Printf("string '%s' -> int %d\n", str, intFromStr)
    }
    
    // String to float
    strFloat := "78.90"
    if floatFromStr, err := strconv.ParseFloat(strFloat, 64); err == nil {
        fmt.Printf("string '%s' -> float64 %.2f\n", strFloat, floatFromStr)
    }
}
```

### 3. Custom Type Conversions

```go
package main

import "fmt"

// Custom types
type UserID string
type SessionID string
type Age int
type Temperature float64

func main() {
    var userID UserID = "user123"
    var sessionID SessionID = "session456"
    var age Age = 25
    var temp Temperature = 36.5
    
    // Convert between custom types with same underlying type
    var convertedSession SessionID = SessionID(userID)
    fmt.Printf("UserID '%s' -> SessionID '%s'\n", userID, convertedSession)
    
    // Convert custom type to underlying type
    var originalString string = string(userID)
    fmt.Printf("UserID '%s' -> string '%s'\n", userID, originalString)
    
    // Convert underlying type to custom type
    var newUserID UserID = UserID("newuser789")
    fmt.Printf("string 'newuser789' -> UserID '%s'\n", newUserID)
    
    // Convert between different numeric custom types
    var ageAsFloat float64 = float64(age)
    fmt.Printf("Age %d -> float64 %.1f\n", age, ageAsFloat)
    
    // Convert float custom type to int custom type
    var tempAsInt int = int(temp)
    fmt.Printf("Temperature %.1f -> int %d\n", temp, tempAsInt)
}
```

### 4. Pointer Type Conversions

```go
package main

import "fmt"

func main() {
    var x int = 42
    var ptr *int = &x
    
    // Convert pointer to uintptr (for low-level operations)
    var addr uintptr = uintptr(unsafe.Pointer(ptr))
    fmt.Printf("Pointer %p -> uintptr %d\n", ptr, addr)
    
    // Convert uintptr back to pointer
    var newPtr *int = (*int)(unsafe.Pointer(addr))
    fmt.Printf("uintptr %d -> Pointer %p\n", addr, newPtr)
    fmt.Printf("Value at new pointer: %d\n", *newPtr)
}
```

## Type Conversion Rules and Limitations

### 1. Compatible Types

```go
package main

import "fmt"

func main() {
    // These conversions are allowed:
    var i int = 42
    var f float64 = float64(i)    // int -> float64
    var b byte = byte(i)          // int -> byte
    var r rune = rune(i)          // int -> rune
    
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float64: %.2f\n", f)
    fmt.Printf("byte: %d\n", b)
    fmt.Printf("rune: %c\n", r)
    
    // These conversions would cause compile errors:
    // var s string = string(i)  // int -> string (not allowed)
    // var b2 bool = bool(i)    // int -> bool (not allowed)
}
```

### 2. String Conversions with Numbers

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num int = 65
    
    // Direct conversion int -> string gives Unicode character
    var char string = string(num)
    fmt.Printf("int %d -> string '%s' (Unicode character)\n", num, char)
    
    // To convert number to string representation, use strconv
    var numStr string = strconv.Itoa(num)
    fmt.Printf("int %d -> string '%s' (number as string)\n", num, numStr)
    
    // Convert string to int
    if convertedNum, err := strconv.Atoi(numStr); err == nil {
        fmt.Printf("string '%s' -> int %d\n", numStr, convertedNum)
    }
}
```

## Advanced Conversion Examples

### 1. Interface Type Assertions

```go
package main

import "fmt"

func main() {
    var i interface{} = "Hello, World!"
    
    // Type assertion (conversion from interface to specific type)
    if str, ok := i.(string); ok {
        fmt.Printf("Interface -> string: '%s'\n", str)
    }
    
    // Type switch for multiple conversions
    switch v := i.(type) {
    case string:
        fmt.Printf("Interface -> string: '%s'\n", v)
    case int:
        fmt.Printf("Interface -> int: %d\n", v)
    case float64:
        fmt.Printf("Interface -> float64: %.2f\n", v)
    default:
        fmt.Printf("Interface -> unknown type: %T\n", v)
    }
}
```

### 2. Slice Conversions

```go
package main

import "fmt"

func main() {
    // Convert between different slice types
    var intSlice []int = []int{1, 2, 3, 4, 5}
    
    // Convert []int to []float64
    var floatSlice []float64 = make([]float64, len(intSlice))
    for i, v := range intSlice {
        floatSlice[i] = float64(v)
    }
    
    fmt.Printf("int slice: %v\n", intSlice)
    fmt.Printf("float64 slice: %v\n", floatSlice)
    
    // Convert []byte to string
    var byteSlice []byte = []byte{'H', 'e', 'l', 'l', 'o'}
    var str string = string(byteSlice)
    fmt.Printf("[]byte %v -> string '%s'\n", byteSlice, str)
    
    // Convert string to []byte
    var newByteSlice []byte = []byte(str)
    fmt.Printf("string '%s' -> []byte %v\n", str, newByteSlice)
}
```

### 3. Custom Type with Methods

```go
package main

import (
    "fmt"
    "strconv"
)

type Temperature float64

// Method to convert to Celsius
func (t Temperature) ToCelsius() Temperature {
    return t
}

// Method to convert to Fahrenheit
func (t Temperature) ToFahrenheit() Temperature {
    return Temperature(t*9/5 + 32)
}

// Method to convert to string
func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C", t)
}

func main() {
    var temp Temperature = 25.0
    
    // Convert to different numeric types
    var asFloat float64 = float64(temp)
    var asInt int = int(temp)
    
    fmt.Printf("Temperature: %s\n", temp)
    fmt.Printf("As float64: %.2f\n", asFloat)
    fmt.Printf("As int: %d\n", asInt)
    fmt.Printf("In Fahrenheit: %.1f°F\n", temp.ToFahrenheit())
}
```

## Common Conversion Patterns

### 1. Error Handling with Conversions

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Safe string to number conversion
    str := "123"
    if num, err := strconv.Atoi(str); err == nil {
        fmt.Printf("Successfully converted '%s' to %d\n", str, num)
    } else {
        fmt.Printf("Failed to convert '%s': %v\n", str, err)
    }
    
    // Safe string to float conversion
    strFloat := "45.67"
    if floatNum, err := strconv.ParseFloat(strFloat, 64); err == nil {
        fmt.Printf("Successfully converted '%s' to %.2f\n", strFloat, floatNum)
    } else {
        fmt.Printf("Failed to convert '%s': %v\n", strFloat, err)
    }
}
```

### 2. Type Conversion in Functions

```go
package main

import "fmt"

type UserID string
type SessionID string

// Function that accepts UserID
func processUser(id UserID) {
    fmt.Printf("Processing user: %s\n", id)
}

// Function that accepts SessionID
func processSession(id SessionID) {
    fmt.Printf("Processing session: %s\n", id)
}

func main() {
    var userID UserID = "user123"
    var sessionID SessionID = "session456"
    
    // Convert and pass to functions
    processUser(UserID(sessionID))  // Convert SessionID to UserID
    processSession(SessionID(userID))  // Convert UserID to SessionID
}
```

## Best Practices for Type Conversion

### 1. Always Use Explicit Conversion

```go
// Good - explicit conversion
var i int = 42
var f float64 = float64(i)

// Bad - implicit conversion (not allowed in Go)
// var f float64 = i  // This would cause a compile error
```

### 2. Handle Conversion Errors

```go
// Good - handle errors
if num, err := strconv.Atoi("123"); err == nil {
    // Use num
} else {
    // Handle error
}

// Bad - ignore errors
// num, _ := strconv.Atoi("123")  // Don't ignore errors
```

### 3. Use Appropriate Conversion Functions

```go
// For string conversions, use strconv package
var str string = strconv.Itoa(123)        // int to string
var num int, _ = strconv.Atoi("456")      // string to int

// For numeric conversions, use type conversion
var f float64 = float64(123)              // int to float64
var i int = int(45.67)                    // float64 to int
```

### 4. Be Aware of Data Loss

```go
var f float64 = 3.14159
var i int = int(f)  // i will be 3 (decimal part is lost)
fmt.Printf("Original: %.5f, Converted: %d\n", f, i)
```

## Common Pitfalls and Solutions

### 1. String Conversion Confusion

```go
var num int = 65

// This gives Unicode character 'A', not "65"
var char string = string(num)
fmt.Printf("Wrong: %s\n", char)  // Output: A

// This gives "65" as string
var numStr string = strconv.Itoa(num)
fmt.Printf("Correct: %s\n", numStr)  // Output: 65
```

### 2. Pointer Conversion Issues

```go
var x int = 42
var ptr *int = &x

// This won't work - can't convert *int to int directly
// var y int = int(ptr)  // Compile error

// Correct way
var y int = *ptr  // Dereference the pointer
fmt.Printf("Value: %d\n", y)
```

## Summary

Type conversion in Go:

- **Requires explicit syntax**: `Type(value)`
- **Is checked at compile time**: Prevents runtime errors
- **Must be compatible**: Types must be convertible
- **Is not casting**: Go doesn't have implicit casting
- **Uses strconv package**: For string conversions
- **Handles errors**: Always check conversion errors

Understanding type conversion is essential for working with Go's type system safely and effectively. Always use explicit conversions and handle errors properly to write robust Go programs.

## Reference

- [Go Language Specification - Conversions](https://golang.org/ref/spec#Conversions)
- [strconv package documentation](https://golang.org/pkg/strconv/)