# Complete Guide to Go Data Types for Beginners

## Table of Contents
1. [Introduction to Go's Type System](#introduction-to-gos-type-system)
2. [Primitive Data Types](#primitive-data-types)
   - [Boolean Type](#boolean-type)
   - [Numeric Types](#numeric-types)
   - [String Type](#string-type)
3. [Composite Data Types](#composite-data-types)
   - [Arrays](#arrays)
   - [Slices](#slices)
   - [Maps](#maps)
   - [Structs](#structs)
4. [Advanced Data Types](#advanced-data-types)
   - [Pointers](#pointers)
   - [Channels](#channels)
   - [Interfaces](#interfaces)
5. [Type Conversion and Type Assertions](#type-conversion-and-type-assertions)
6. [Best Practices](#best-practices)

---

## Introduction to Go's Type System

Go is a **statically typed** programming language, which means:
- Every variable must be declared with a specific type
- Type checking occurs at compile time, not runtime
- Type safety prevents many common programming errors
- The compiler ensures type consistency throughout your program

### Key Concepts:
- **Zero Values**: Every type has a default value when declared without initialization
- **Type Inference**: Go can automatically determine the type in many cases
- **Explicit Conversion**: Go requires explicit type conversion (no implicit casting)

---

## Primitive Data Types

### Boolean Type

The `bool` type represents boolean values with only two possible values: `true` or `false`.

**Zero Value**: `false`

```go
package main

import "fmt"

func main() {
    // Declaration and initialization
    var isActive bool = true
    var isComplete bool = false
    
    // Zero value
    var defaultBool bool
    fmt.Printf("Default bool value: %t\n", defaultBool) // false
    
    // Boolean operations
    fmt.Printf("isActive: %t\n", isActive)
    fmt.Printf("isComplete: %t\n", isComplete)
    fmt.Printf("AND operation: %t\n", isActive && isComplete)
    fmt.Printf("OR operation: %t\n", isActive || isComplete)
    fmt.Printf("NOT operation: %t\n", !isActive)
    
    // Comparison operations return boolean
    a := 10
    b := 20
    fmt.Printf("a == b: %t\n", a == b)
    fmt.Printf("a < b: %t\n", a < b)
    fmt.Printf("a > b: %t\n", a > b)
}
```

### Numeric Types

Go provides several numeric types for different use cases:

#### Integer Types

| Type | Size | Range |
|------|------|-------|
| `int8` | 8 bits | -128 to 127 |
| `int16` | 16 bits | -32,768 to 32,767 |
| `int32` | 32 bits | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64 bits | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `int` | Platform dependent | 32 or 64 bits |

#### Unsigned Integer Types

| Type | Size | Range |
|------|------|-------|
| `uint8` | 8 bits | 0 to 255 |
| `uint16` | 16 bits | 0 to 65,535 |
| `uint32` | 32 bits | 0 to 4,294,967,295 |
| `uint64` | 64 bits | 0 to 18,446,744,073,709,551,615 |
| `uint` | Platform dependent | 32 or 64 bits |

#### Special Types
- `byte` - alias for `uint8`
- `rune` - alias for `int32` (represents Unicode code points)

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Integer types
    var a int = 42
    var b int8 = 127
    var c int16 = 32767
    var d int32 = 2147483647
    var e int64 = 9223372036854775807
    
    // Unsigned integer types
    var f uint = 42
    var g uint8 = 255
    var h uint16 = 65535
    var i uint32 = 4294967295
    var j uint64 = 18446744073709551615
    
    // Special types
    var k byte = 255        // alias for uint8
    var l rune = 'A'        // alias for int32
    
    fmt.Printf("Signed integers: %d, %d, %d, %d, %d\n", a, b, c, d, e)
    fmt.Printf("Unsigned integers: %d, %d, %d, %d, %d\n", f, g, h, i, j)
    fmt.Printf("Byte: %d, Rune: %c (%d)\n", k, l, l)
    
    // Zero values
    var zeroInt int
    var zeroUint uint
    fmt.Printf("Zero int: %d, Zero uint: %d\n", zeroInt, zeroUint)
}
```

#### Floating-Point Types

| Type | Size | Precision |
|------|------|-----------|
| `float32` | 32 bits | ~7 decimal digits |
| `float64` | 64 bits | ~15 decimal digits |

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Floating-point types
    var pi32 float32 = 3.14159
    var pi64 float64 = 3.141592653589793
    
    // Zero values
    var zeroFloat32 float32
    var zeroFloat64 float64
    fmt.Printf("Zero float32: %f, Zero float64: %f\n", zeroFloat32, zeroFloat64)
    
    // Mathematical operations
    fmt.Printf("Pi (32-bit): %.5f\n", pi32)
    fmt.Printf("Pi (64-bit): %.15f\n", pi64)
    fmt.Printf("Square root of 2: %.10f\n", math.Sqrt(2))
    
    // Special floating-point values
    fmt.Printf("Positive infinity: %f\n", math.Inf(1))
    fmt.Printf("Negative infinity: %f\n", math.Inf(-1))
    fmt.Printf("Not a number: %f\n", math.NaN())
}
```

#### Complex Types

| Type | Size | Description |
|------|------|-------------|
| `complex64` | 64 bits | Complex number with float32 real and imaginary parts |
| `complex128` | 128 bits | Complex number with float64 real and imaginary parts |

```go
package main

import "fmt"

func main() {
    // Complex numbers
    var c1 complex64 = 1 + 2i
    var c2 complex128 = 3 + 4i
    
    // Zero values
    var zeroComplex64 complex64
    var zeroComplex128 complex128
    fmt.Printf("Zero complex64: %v\n", zeroComplex64)
    fmt.Printf("Zero complex128: %v\n", zeroComplex128)
    
    // Complex operations
    fmt.Printf("Complex64: %v\n", c1)
    fmt.Printf("Complex128: %v\n", c2)
    fmt.Printf("Real part: %.2f\n", real(c2))
    fmt.Printf("Imaginary part: %.2f\n", imag(c2))
}
```

### String Type

Strings in Go are sequences of bytes and are immutable.

**Zero Value**: `""` (empty string)

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // String literals
    str1 := "Hello, World!"           // Interpreted string literal
    str2 := `This is a raw string
    with multiple lines
    and no escape sequences`         // Raw string literal
    
    // Zero value
    var emptyString string
    fmt.Printf("Empty string: '%s'\n", emptyString)
    
    // String operations
    fmt.Printf("String: %s\n", str1)
    fmt.Printf("Length: %d\n", len(str1))
    fmt.Printf("Raw string: %s\n", str2)
    
    // String concatenation
    firstName := "John"
    lastName := "Doe"
    fullName := firstName + " " + lastName
    fmt.Printf("Full name: %s\n", fullName)
    
    // String indexing (returns bytes, not runes)
    fmt.Printf("First character: %c\n", str1[0])
    fmt.Printf("Last character: %c\n", str1[len(str1)-1])
    
    // String iteration
    fmt.Println("Character by character:")
    for i, char := range str1 {
        fmt.Printf("Index %d: %c\n", i, char)
    }
    
    // String methods
    fmt.Printf("Uppercase: %s\n", strings.ToUpper(str1))
    fmt.Printf("Lowercase: %s\n", strings.ToLower(str1))
    fmt.Printf("Contains 'World': %t\n", strings.Contains(str1, "World"))
}
```

---

## Composite Data Types

### Arrays

Arrays are fixed-size sequences of elements of the same type.

**Zero Value**: Array with zero values for each element

```go
package main

import "fmt"

func main() {
    // Array declaration and initialization
    var numbers [5]int                    // Zero-initialized array
    var fruits = [3]string{"apple", "banana", "orange"}
    var matrix = [2][3]int{{1, 2, 3}, {4, 5, 6}}
    
    // Zero values
    fmt.Printf("Zero array: %v\n", numbers)
    
    // Array initialization
    numbers = [5]int{1, 2, 3, 4, 5}
    fmt.Printf("Numbers: %v\n", numbers)
    fmt.Printf("Fruits: %v\n", fruits)
    fmt.Printf("Matrix: %v\n", matrix)
    
    // Array operations
    fmt.Printf("Length: %d\n", len(numbers))
    fmt.Printf("First element: %d\n", numbers[0])
    fmt.Printf("Last element: %d\n", numbers[len(numbers)-1])
    
    // Array iteration
    fmt.Println("Array elements:")
    for i, value := range numbers {
        fmt.Printf("Index %d: %d\n", i, value)
    }
    
    // Array modification
    numbers[0] = 100
    fmt.Printf("Modified array: %v\n", numbers)
}
```

### Slices

Slices are dynamic arrays that can grow and shrink.

**Zero Value**: `nil`

```go
package main

import "fmt"

func main() {
    // Slice declaration
    var numbers []int                     // nil slice
    var fruits = []string{"apple", "banana", "orange"}
    
    // Zero values
    fmt.Printf("Nil slice: %v (length: %d, capacity: %d)\n", numbers, len(numbers), cap(numbers))
    
    // Slice creation
    numbers = []int{1, 2, 3, 4, 5}
    fmt.Printf("Numbers: %v (length: %d, capacity: %d)\n", numbers, len(numbers), cap(numbers))
    
    // Slice operations
    fmt.Printf("First element: %d\n", numbers[0])
    fmt.Printf("Last element: %d\n", numbers[len(numbers)-1])
    
    // Slice iteration
    fmt.Println("Slice elements:")
    for i, value := range numbers {
        fmt.Printf("Index %d: %d\n", i, value)
    }
    
    // Slice modification
    numbers[0] = 100
    fmt.Printf("Modified slice: %v\n", numbers)
    
    // Slice operations
    numbers = append(numbers, 6, 7, 8)
    fmt.Printf("After append: %v (length: %d, capacity: %d)\n", numbers, len(numbers), cap(numbers))
    
    // Slice slicing
    subSlice := numbers[1:4]
    fmt.Printf("Sub-slice [1:4]: %v\n", subSlice)
    
    // Make slice with specific length and capacity
    newSlice := make([]int, 3, 5)
    fmt.Printf("Made slice: %v (length: %d, capacity: %d)\n", newSlice, len(newSlice), cap(newSlice))
}
```

### Maps

Maps are key-value pairs, similar to dictionaries in other languages.

**Zero Value**: `nil`

```go
package main

import "fmt"

func main() {
    // Map declaration
    var ages map[string]int               // nil map
    var scores = map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 92,
    }
    
    // Zero values
    fmt.Printf("Nil map: %v\n", ages)
    
    // Map initialization
    ages = make(map[string]int)
    ages["Alice"] = 25
    ages["Bob"] = 30
    ages["Carol"] = 28
    
    fmt.Printf("Ages: %v\n", ages)
    fmt.Printf("Scores: %v\n", scores)
    
    // Map operations
    fmt.Printf("Alice's age: %d\n", ages["Alice"])
    fmt.Printf("Alice's score: %d\n", scores["Alice"])
    
    // Check if key exists
    if age, exists := ages["Alice"]; exists {
        fmt.Printf("Alice's age: %d\n", age)
    }
    
    // Map iteration
    fmt.Println("All ages:")
    for name, age := range ages {
        fmt.Printf("%s: %d\n", name, age)
    }
    
    // Map modification
    ages["Alice"] = 26
    fmt.Printf("Updated ages: %v\n", ages)
    
    // Delete from map
    delete(ages, "Bob")
    fmt.Printf("After deletion: %v\n", ages)
}
```

### Structs

Structs are collections of fields with different types.

**Zero Value**: Struct with zero values for each field

```go
package main

import "fmt"

// Define a struct
type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address
}

type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

func main() {
    // Struct declaration
    var person Person                    // Zero-initialized struct
    
    // Zero values
    fmt.Printf("Zero struct: %+v\n", person)
    
    // Struct initialization
    person = Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john@example.com",
        Address: Address{
            Street: "123 Main St",
            City:   "New York",
            State:  "NY",
            Zip:    "10001",
        },
    }
    
    fmt.Printf("Person: %+v\n", person)
    
    // Accessing struct fields
    fmt.Printf("Name: %s\n", person.Name)
    fmt.Printf("Age: %d\n", person.Age)
    fmt.Printf("Email: %s\n", person.Email)
    fmt.Printf("Address: %s, %s, %s %s\n", 
        person.Address.Street, 
        person.Address.City, 
        person.Address.State, 
        person.Address.Zip)
    
    // Modifying struct fields
    person.Age = 31
    person.Email = "john.doe@example.com"
    fmt.Printf("Updated person: %+v\n", person)
    
    // Anonymous struct
    employee := struct {
        ID       int
        Position string
        Salary   float64
    }{
        ID:       1001,
        Position: "Software Engineer",
        Salary:   75000.0,
    }
    
    fmt.Printf("Employee: %+v\n", employee)
}
```

---

## Advanced Data Types

### Pointers

Pointers store the memory address of a value.

**Zero Value**: `nil`

```go
package main

import "fmt"

func main() {
    // Pointer declaration
    var p *int                           // nil pointer
    var x int = 42
    
    // Zero values
    fmt.Printf("Nil pointer: %v\n", p)
    
    // Getting address of variable
    p = &x
    fmt.Printf("Pointer value: %v\n", p)
    fmt.Printf("Value at pointer: %d\n", *p)
    
    // Modifying value through pointer
    *p = 100
    fmt.Printf("Modified x through pointer: %d\n", x)
    
    // Pointer to pointer
    var pp **int = &p
    fmt.Printf("Pointer to pointer: %v\n", pp)
    fmt.Printf("Value through pointer to pointer: %d\n", **pp)
    
    // Pointer arithmetic (limited in Go)
    var arr = [3]int{1, 2, 3}
    var ptr *int = &arr[0]
    fmt.Printf("First element: %d\n", *ptr)
    ptr = &arr[1]
    fmt.Printf("Second element: %d\n", *ptr)
}
```

### Channels

Channels are used for communication between goroutines.

**Zero Value**: `nil`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Channel declaration
    var ch chan int                      // nil channel
    ch = make(chan int)                  // Buffered channel
    ch2 := make(chan string, 2)          // Buffered channel with capacity 2
    
    // Zero values
    fmt.Printf("Nil channel: %v\n", ch)
    
    // Channel operations
    go func() {
        ch <- 42                         // Send value to channel
        ch2 <- "Hello"
        ch2 <- "World"
    }()
    
    // Receive from channel
    value := <-ch
    fmt.Printf("Received: %d\n", value)
    
    // Receive from buffered channel
    msg1 := <-ch2
    msg2 := <-ch2
    fmt.Printf("Messages: %s, %s\n", msg1, msg2)
    
    // Channel with select
    ch3 := make(chan int)
    go func() {
        time.Sleep(1 * time.Second)
        ch3 <- 100
    }()
    
    select {
    case val := <-ch3:
        fmt.Printf("Received from ch3: %d\n", val)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout")
    }
}
```

### Interfaces

Interfaces define behavior through method signatures.

**Zero Value**: `nil`

```go
package main

import "fmt"y

// Define an interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implement interface for Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Implement interface for Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

func main() {
    // Interface usage
    var shape Shape                       // nil interface
    
    // Zero values
    fmt.Printf("Nil interface: %v\n", shape)
    
    // Create shapes
    rect := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 4}
    
    // Use interface
    shape = rect
    fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", 
        shape.Area(), shape.Perimeter())
    
    shape = circle
    fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", 
        shape.Area(), shape.Perimeter())
    
    // Type assertion
    if rect, ok := shape.(Rectangle); ok {
        fmt.Printf("It's a rectangle: %+v\n", rect)
    } else if circle, ok := shape.(Circle); ok {
        fmt.Printf("It's a circle: %+v\n", circle)
    }
    
    // Type switch
    switch s := shape.(type) {
    case Rectangle:
        fmt.Printf("Rectangle with width %.2f and height %.2f\n", s.Width, s.Height)
    case Circle:
        fmt.Printf("Circle with radius %.2f\n", s.Radius)
    default:
        fmt.Printf("Unknown shape: %T\n", s)
    }
}
```

---

## Type Conversion and Type Assertions

### Type Conversion

Go requires explicit type conversion:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Numeric conversions
    var i int = 42
    var f float64 = float64(i)
    var b byte = byte(i)
    
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float64: %.2f\n", f)
    fmt.Printf("byte: %d\n", b)
    
    // String conversions
    var num int = 123
    var str string = strconv.Itoa(num)
    fmt.Printf("Number as string: %s\n", str)
    
    // String to number
    str2 := "456"
    if num2, err := strconv.Atoi(str2); err == nil {
        fmt.Printf("String as number: %d\n", num2)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    
    // Float conversions
    var pi float64 = 3.14159
    var pi32 float32 = float32(pi)
    fmt.Printf("float64: %.5f\n", pi)
    fmt.Printf("float32: %.5f\n", pi32)
}
```

### Type Assertions

Type assertions are used with interfaces:

```go
package main

import "fmt"

func main() {y
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

---

## Best Practices

### 1. Use Appropriate Types
```go
// Good: Use specific types when precision matters
var userID uint64 = 123456789
var temperature float32 = 25.5

// Avoid: Using generic types when specific ones are better
var id int = 123456789  // Could overflow on 32-bit systems
```

### 2. Handle Zero Values
```go
// Good: Check for zero values
var name string
if name == "" {
    name = "Unknown"
}

// Good: Use zero values effectively
var numbers []int  // nil slice
if numbers == nil {
    numbers = make([]int, 0, 10)
}
```

### 3. Use Type Inference Wisely
```go
// Good: Use type inference for simple cases
name := "John"
age := 25

// Good: Use explicit types for clarity
var userID uint64 = 123456789
var isActive bool = true
```

### 4. Handle Type Conversion Errors
```go
// Good: Always handle conversion errors
str := "123"
if num, err := strconv.Atoi(str); err == nil {
    fmt.Printf("Number: %d\n", num)
} else {
    fmt.Printf("Error: %v\n", err)
}
```

### 5. Use Interfaces for Flexibility
```go
// Good: Use interfaces for flexible code
type Writer interface {
    Write([]byte) (int, error)
}

func processData(w Writer, data []byte) {
    w.Write(data)
}
```

---

## Summary

Go's type system provides:
- **Static typing** for compile-time error detection
- **Type inference** for cleaner code
- **Zero values** for safe initialization
- **Explicit conversions** to prevent accidental type mixing
- **Type safety** to prevent runtime errors

Understanding Go's data types is fundamental to writing safe, efficient, and maintainable Go programs. Each type has its specific use case and understanding when to use which type is crucial for effective Go programming.
