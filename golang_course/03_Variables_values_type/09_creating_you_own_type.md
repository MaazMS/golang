# Creating Your Own Types in Go

## Introduction

Go is a **statically typed** programming language, which means that types are determined at compile time. Because of this, types are extremely important in Go programming. Go allows you to create your own custom types using the `type` keyword, which helps improve code readability, type safety, and maintainability.

## Understanding Custom Types

### What are Custom Types?

Custom types in Go are user-defined types that are based on existing types (called **underlying types**). When you create a custom type, you must define its underlying type.

### Basic Syntax

```go
type CustomType UnderlyingType
```

## Underlying Types

The **underlying type** is the base type that your custom type is built upon. Understanding underlying types is crucial for type conversion and method inheritance.

### Example of Underlying Types

```go
type T1 string    // T1's underlying type is string
type T2 T1        // T2's underlying type is string (inherited from T1)
```

In this example:
- `T1` has `string` as its underlying type
- `T2` has `string` as its underlying type (inherited from `T1`)

## Basic Custom Type Example

```go
package main

import (
    "fmt"
)

var a int

// Create our own type 'palindrome' with underlying type 'int'
type palindrome int

// Create variable b with TYPE palindrome
var b palindrome

func main() {
    a = 15
    // Assign value to variable b with TYPE palindrome (underlying type is int)
    b = 25
    
    fmt.Println("Original int:", a)
    fmt.Printf("Type of a: %T\n", a)
    
    fmt.Println("Custom palindrome:", b)
    fmt.Printf("Type of b: %T\n", b)
}
```

## Comprehensive Custom Type Examples

### 1. String-Based Custom Types

```go
package main

import "fmt"

// Custom types based on string
type UserID string
type Email string
type Password string

func main() {
    var userID UserID = "user123"
    var email Email = "maaz@example.com"
    var password Password = "securepass123"
    
    fmt.Printf("UserID: %s (Type: %T)\n", userID, userID)
    fmt.Printf("Email: %s (Type: %T)\n", email, email)
    fmt.Printf("Password: %s (Type: %T)\n", password, password)
}
```

### 2. Numeric-Based Custom Types

```go
package main

import "fmt"

// Custom types for different units
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// Custom types for measurements
type Meters float64
type Feet float64

func main() {
    var temp Celsius = 25.5
    var distance Meters = 100.0
    
    fmt.Printf("Temperature: %.1f°C (Type: %T)\n", temp, temp)
    fmt.Printf("Distance: %.1f meters (Type: %T)\n", distance, distance)
}
```

### 3. Boolean-Based Custom Types

```go
package main

import "fmt"

type IsActive bool
type IsVerified bool
type IsPremium bool

func main() {
    var active IsActive = true
    var verified IsVerified = false
    var premium IsPremium = true
    
    fmt.Printf("Active: %t (Type: %T)\n", active, active)
    fmt.Printf("Verified: %t (Type: %T)\n", verified, verified)
    fmt.Printf("Premium: %t (Type: %T)\n", premium, premium)
}
```

## Type Conversion with Custom Types

Custom types require explicit conversion, even if they have the same underlying type:

```go
package main

import "fmt"

type UserID string
type SessionID string

func main() {
    var userID UserID = "user123"
    var sessionID SessionID = "session456"
    
    // This would cause a compile error:
    // userID = sessionID  // cannot use sessionID (type SessionID) as type UserID
    
    // Correct way - explicit conversion
    userID = UserID(sessionID)
    fmt.Printf("Converted UserID: %s\n", userID)
    
    // Convert to underlying type
    var originalString string = string(userID)
    fmt.Printf("Original string: %s\n", originalString)
}
```

## Adding Methods to Custom Types

One of the most powerful features of custom types is the ability to add methods:

```go
package main

import (
    "fmt"
    "strings"
)

type Email string

// Method to validate email format
func (e Email) IsValid() bool {
    return strings.Contains(string(e), "@") && strings.Contains(string(e), ".")
}

// Method to get domain
func (e Email) GetDomain() string {
    parts := strings.Split(string(e), "@")
    if len(parts) == 2 {
        return parts[1]
    }
    return ""
}

// Method to convert to uppercase
func (e Email) ToUpper() Email {
    return Email(strings.ToUpper(string(e)))
}

func main() {
    var email Email = "maaz@example.com"
    
    fmt.Printf("Email: %s\n", email)
    fmt.Printf("Is valid: %t\n", email.IsValid())
    fmt.Printf("Domain: %s\n", email.GetDomain())
    fmt.Printf("Uppercase: %s\n", email.ToUpper())
}
```

## Advanced Custom Type Examples

### 1. Currency Types

```go
package main

import "fmt"

type USD float64
type EUR float64
type INR float64

// Method to format USD currency
func (u USD) String() string {
    return fmt.Sprintf("$%.2f", u)
}

// Method to format EUR currency
func (e EUR) String() string {
    return fmt.Sprintf("€%.2f", e)
}

// Method to format INR currency
func (i INR) String() string {
    return fmt.Sprintf("₹%.2f", i)
}

func main() {
    var priceUSD USD = 99.99
    var priceEUR EUR = 85.50
    var priceINR INR = 7500.00
    
    fmt.Printf("USD Price: %s\n", priceUSD)
    fmt.Printf("EUR Price: %s\n", priceEUR)
    fmt.Printf("INR Price: %s\n", priceINR)
}
```

### 2. Status Types

```go
package main

import "fmt"

type OrderStatus string

const (
    StatusPending   OrderStatus = "pending"
    StatusConfirmed OrderStatus = "confirmed"
    StatusShipped   OrderStatus = "shipped"
    StatusDelivered OrderStatus = "delivered"
    StatusCancelled OrderStatus = "cancelled"
)

// Method to check if order is active
func (s OrderStatus) IsActive() bool {
    return s == StatusPending || s == StatusConfirmed || s == StatusShipped
}

// Method to get next status
func (s OrderStatus) GetNextStatus() OrderStatus {
    switch s {
    case StatusPending:
        return StatusConfirmed
    case StatusConfirmed:
        return StatusShipped
    case StatusShipped:
        return StatusDelivered
    default:
        return s
    }
}

func main() {
    var status OrderStatus = StatusPending
    
    fmt.Printf("Current status: %s\n", status)
    fmt.Printf("Is active: %t\n", status.IsActive())
    fmt.Printf("Next status: %s\n", status.GetNextStatus())
}
```

### 3. Measurement Types with Conversion

```go
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

// Convert Celsius to Fahrenheit
func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

// Convert Fahrenheit to Celsius
func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// String method for Celsius
func (c Celsius) String() string {
    return fmt.Sprintf("%.1f°C", c)
}

// String method for Fahrenheit
func (f Fahrenheit) String() string {
    return fmt.Sprintf("%.1f°F", f)
}

func main() {
    var roomTemp Celsius = 25.0
    var bodyTemp Fahrenheit = 98.6
    
    fmt.Printf("Room temperature: %s\n", roomTemp)
    fmt.Printf("Room temperature in Fahrenheit: %s\n", roomTemp.ToFahrenheit())
    
    fmt.Printf("Body temperature: %s\n", bodyTemp)
    fmt.Printf("Body temperature in Celsius: %s\n", bodyTemp.ToCelsius())
}
```

## Type Aliases vs Custom Types

Go also supports type aliases, which are different from custom types:

```go
package main

import "fmt"

// Type alias - same type, different name
type StringAlias = string

// Custom type - different type, same underlying type
type CustomString string

func main() {
    var alias StringAlias = "hello"
    var custom CustomString = "world"
    var original string = "golang"
    
    // Type alias can be assigned directly
    alias = original  // This works
    
    // Custom type requires conversion
    // custom = original  // This would cause an error
    custom = CustomString(original)  // This works
    
    fmt.Printf("Alias: %s (Type: %T)\n", alias, alias)
    fmt.Printf("Custom: %s (Type: %T)\n", custom, custom)
}
```

## Best Practices for Custom Types

### 1. Use Descriptive Names

```go
// Good
type UserID string
type OrderAmount float64
type IsActive bool

// Avoid
type UID string
type Amt float64
type Flag bool
```

### 2. Add Methods for Behavior

```go
type Email string

func (e Email) IsValid() bool {
    // validation logic
    return true
}

func (e Email) GetDomain() string {
    // domain extraction logic
    return "example.com"
}
```

### 3. Use Constants for Status Types

```go
type OrderStatus string

const (
    StatusPending   OrderStatus = "pending"
    StatusConfirmed OrderStatus = "confirmed"
    StatusShipped   OrderStatus = "shipped"
)
```

### 4. Implement String() Method for Custom Types

```go
type Currency float64

func (c Currency) String() string {
    return fmt.Sprintf("$%.2f", c)
}
```

## Common Use Cases

1. **Domain-Specific Types**: Create types that represent concepts in your domain
2. **Type Safety**: Prevent mixing different types that have the same underlying type
3. **Method Attachment**: Add behavior to basic types
4. **API Design**: Create clear interfaces for your APIs
5. **Configuration**: Use custom types for configuration values

## Summary

Custom types in Go provide:

- **Type Safety**: Prevent accidental mixing of different concepts
- **Code Clarity**: Make code more readable and self-documenting
- **Method Attachment**: Add behavior to basic types
- **Domain Modeling**: Represent business concepts clearly
- **API Design**: Create clear and safe interfaces

Creating custom types is a fundamental concept in Go that helps you write more maintainable, safe, and expressive code. Always define the underlying type clearly and consider adding methods to make your types more useful.

## Reference

- [Go Language Specification - Types](https://golang.org/ref/spec#Types)