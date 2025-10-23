# Exploring Types in Go

## Introduction to Go's Type System

Go is a **statically typed** programming language, which means that:
- Every variable must be declared with a specific type
- Type checking occurs at compile time, not runtime
- Type safety prevents many common programming errors
- The compiler ensures type consistency throughout your program

## Basic Type Declaration

In Go, a variable is declared to hold a value of a specific type. In the example below, the variable with the identifier `name` is of type `string`, and it is assigned the value `"Maaz shaikh"`.

**Important**: Assigning a value of the wrong type results in a compile-time error (for example, `cannot use 10 (type untyped int) as type string in assignment`).

## Basic Example

```go
package main

import (
    "fmt"
)

// Interpreted string literal (supports escape sequences)
var name = "Maaz shaikh"

// Raw string literal (backticks) preserves newlines and characters verbatim
var user = `user
name is
Maaz Shaikh`

// Explicit type declaration
var age int = 30

func main() {
    fmt.Println(name)
    fmt.Printf("type of name: %T\n", name)

    fmt.Println(user)
    fmt.Printf("type of user: %T\n", user)

    fmt.Println(age)
    fmt.Printf("type of age: %T\n", age)

    // Type mismatch example (uncomment to see compile error):
    // name = 10 // cannot use 10 (type untyped int) as type string in assignment

    // Type inference with short declaration inside a function
    city := "Pune"
    fmt.Printf("%s (%T)\n", city, city)
}
```

## Go's Built-in Types

### Numeric Types

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Integer types
    var a int = 42           // Platform-dependent size (32 or 64 bits)
    var b int8 = 127         // 8-bit signed integer (-128 to 127)
    var c int16 = 32767      // 16-bit signed integer
    var d int32 = 2147483647 // 32-bit signed integer
    var e int64 = 9223372036854775807 // 64-bit signed integer
    
    // Unsigned integer types
    var f uint = 42          // Platform-dependent unsigned
    var g uint8 = 255        // 8-bit unsigned integer (0 to 255)
    var h uint16 = 65535    // 16-bit unsigned integer
    var i uint32 = 4294967295 // 32-bit unsigned integer
    var j uint64 = 18446744073709551615 // 64-bit unsigned integer
    
    // Floating-point types
    var k float32 = 3.14     // 32-bit floating-point
    var l float64 = 3.141592653589793 // 64-bit floating-point
    
    // Complex types
    var m complex64 = 1 + 2i  // 64-bit complex number
    var n complex128 = 1 + 2i // 128-bit complex number
    
    // Byte and rune (aliases)
    var o byte = 255         // alias for uint8
    var p rune = 'A'         // alias for int32, represents Unicode code point
    
    fmt.Printf("Integer types: %d, %d, %d, %d, %d\n", a, b, c, d, e)
    fmt.Printf("Unsigned types: %d, %d, %d, %d, %d\n", f, g, h, i, j)
    fmt.Printf("Float types: %.2f, %.15f\n", k, l)
    fmt.Printf("Complex types: %v, %v\n", m, n)
    fmt.Printf("Byte and rune: %d, %c\n", o, p)
}
```

### String Types

```go
package main

import "fmt"

func main() {
    // Interpreted string literals (double quotes)
    var str1 = "Hello, World!\nThis is a new line."
    var str2 = "Path: C:\\Users\\Name\\Documents"
    
    // Raw string literals (backticks) - no escape sequences
    var str3 = `Hello, World!
This is a new line.
Path: C:\Users\Name\Documents`
    
    // String concatenation
    var firstName = "Maaz"
    var lastName = "Shaikh"
    var fullName = firstName + " " + lastName
    
    // String length and indexing
    fmt.Printf("String length: %d\n", len(fullName))
    fmt.Printf("First character: %c\n", fullName[0])
    
    fmt.Println("Interpreted string:", str1)
    fmt.Println("Raw string:", str3)
    fmt.Println("Full name:", fullName)
}
```

### Boolean Type

```go
package main

import "fmt"

func main() {
    var isActive bool = true
    var isComplete bool = false
    
    // Boolean operations
    fmt.Printf("isActive: %t\n", isActive)
    fmt.Printf("isComplete: %t\n", isComplete)
    fmt.Printf("isActive && isComplete: %t\n", isActive && isComplete)
    fmt.Printf("isActive || isComplete: %t\n", isActive || isComplete)
    fmt.Printf("!isActive: %t\n", !isActive)
}
```

## Type Inference

Go can infer types automatically in many cases:

```go
package main

import "fmt"

func main() {
    // Type inference examples
    var name = "Maaz"        // inferred as string
    var age = 25             // inferred as int
    var salary = 50000.50    // inferred as float64
    var isEmployed = true    // inferred as bool
    
    // Short declaration with type inference
    city := "Pune"           // inferred as string
    population := 3000000    // inferred as int
    temperature := 25.5      // inferred as float64
    
    fmt.Printf("Name: %s (%T)\n", name, name)
    fmt.Printf("Age: %d (%T)\n", age, age)
    fmt.Printf("Salary: %.2f (%T)\n", salary, salary)
    fmt.Printf("Employed: %t (%T)\n", isEmployed, isEmployed)
    fmt.Printf("City: %s (%T)\n", city, city)
    fmt.Printf("Population: %d (%T)\n", population, population)
    fmt.Printf("Temperature: %.1f (%T)\n", temperature, temperature)
}
```

## Zero Values

In Go, variables declared without an initial value are given their **zero value**:

```go
package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    var p *int
    
    fmt.Printf("Zero values:\n")
    fmt.Printf("int: %d\n", i)           // 0
    fmt.Printf("float64: %.2f\n", f)     // 0.00
    fmt.Printf("bool: %t\n", b)          // false
    fmt.Printf("string: '%s'\n", s)      // "" (empty string)
    fmt.Printf("pointer: %v\n", p)       // <nil>
}
```

## Type Conversion

Go requires explicit type conversion (casting):

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Numeric type conversions
    var i int = 42
    var f float64 = float64(i)    // int to float64
    var b byte = byte(i)          // int to byte
    
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float64: %.2f\n", f)
    fmt.Printf("byte: %d\n", b)
    
    // String conversions
    var num int = 123
    var str string = strconv.Itoa(num)  // int to string
    fmt.Printf("Number as string: %s\n", str)
    
    // String to number
    str2 := "456"
    if num2, err := strconv.Atoi(str2); err == nil {
        fmt.Printf("String as number: %d\n", num2)
    }
    
    // Float conversions
    var pi float64 = 3.14159
    var pi32 float32 = float32(pi)
    fmt.Printf("float64: %.5f\n", pi)
    fmt.Printf("float32: %.5f\n", pi32)
}
```

## Type Assertions (for Interfaces)

```go
package main

import "fmt"

func main() {
    var i interface{} = "Hello, World!"
    
    // Type assertion
    if str, ok := i.(string); ok {
        fmt.Printf("String value: %s\n", str)
    }
    
    // Type switch
    switch v := i.(type) {
    case string:
        fmt.Printf("It's a string: %s\n", v)
    case int:
        fmt.Printf("It's an int: %d\n", v)
    case float64:
        fmt.Printf("It's a float64: %.2f\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## Type Safety Benefits

```go
package main

import "fmt"

func main() {
    var name string = "Maaz"
    var age int = 25
    
    // This would cause a compile-time error:
    // name = age  // cannot use age (type int) as type string in assignment
    
    // Correct way - explicit conversion
    name = fmt.Sprintf("Age: %d", age)
    fmt.Println(name)
    
    // Type safety prevents runtime errors
    var numbers []int = []int{1, 2, 3, 4, 5}
    // numbers[10] = 100  // This would cause a runtime panic (index out of range)
    
    // Safe access with bounds checking
    if len(numbers) > 10 {
        numbers[10] = 100
    } else {
        fmt.Println("Index out of bounds")
    }
}
```

## Best Practices

1. **Use explicit types** when clarity is important
2. **Use type inference** for simple cases to reduce verbosity
3. **Always handle type conversion errors** when converting strings to numbers
4. **Use meaningful variable names** that indicate their type and purpose
5. **Leverage Go's type safety** to catch errors at compile time

## Summary

Go's type system provides:
- **Static typing** for compile-time error detection
- **Type inference** for cleaner code
- **Zero values** for safe initialization
- **Explicit conversions** to prevent accidental type mixing
- **Type safety** to prevent runtime errors

Understanding Go's type system is fundamental to writing safe, efficient, and maintainable Go programs.