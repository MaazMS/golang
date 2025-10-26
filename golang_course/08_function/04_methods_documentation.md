# Go Methods - Complete Guide

## Overview

Methods in Go are functions that are attached to a specific type through a receiver. They allow you to define behavior for types, making your code more organized and object-oriented. Methods are different from functions in that they have a receiver that gives them access to the properties and fields of the type they are attached to.

## Key Features

1. **Receiver attachment**: Methods are attached to types through receivers
2. **Type-specific behavior**: Methods provide behavior specific to a type
3. **Value and pointer receivers**: Can use both value and pointer receivers
4. **Method sets**: Different method sets for value and pointer receivers
5. **Package restriction**: Receiver type must be in the same package
6. **Method overloading**: Can have methods with the same name on different types

## Why Use Methods?

### Benefits
- **Code organization**: Group related functionality with types
- **Object-oriented programming**: Encapsulate behavior with data
- **Type safety**: Methods are type-specific and safe
- **Interface implementation**: Methods enable interface implementation
- **Code reusability**: Methods can be called on multiple instances

### When to Use Methods
- Define behavior for custom types
- Implement interfaces
- Encapsulate type-specific operations
- Create object-oriented designs
- Group related functionality

## Basic Syntax

### Method Declaration
```go
func (receiverName receiverType) methodName(parameters) returnType {
    // method body
}
```

### Key Components
- **func**: Keyword to declare a method
- **receiverName**: Name of the receiver (can be any valid identifier)
- **receiverType**: Type that the method is attached to
- **methodName**: Name of the method
- **parameters**: Input parameters (optional)
- **returnType**: Return type (optional)

## Basic Examples

### Example 1: Simple Method

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{
        Name: "Alice",
        Age:  30,
    }
    
    fmt.Println("Person:", p)
    p.Speak()
    p.Introduce()
}

// Method with value receiver
func (p Person) Speak() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Method with value receiver
func (p Person) Introduce() {
    fmt.Printf("Nice to meet you! I'm %s.\n", p.Name)
}
```

**Output:**
```
Person: {Alice 30}
Hello, my name is Alice and I am 30 years old.
Nice to meet you! I'm Alice.
```

### Example 2: Value vs Pointer Receivers

```go
package main

import "fmt"

type Counter struct {
    value int
}

func main() {
    c := Counter{value: 0}
    
    fmt.Printf("Initial value: %d\n", c.GetValue())
    
    // Value receiver - doesn't modify original
    c.IncrementValue()
    fmt.Printf("After IncrementValue: %d\n", c.GetValue())
    
    // Pointer receiver - modifies original
    c.IncrementPointer()
    fmt.Printf("After IncrementPointer: %d\n", c.GetValue())
    
    // Reset counter
    c.Reset()
    fmt.Printf("After Reset: %d\n", c.GetValue())
}

// Value receiver - cannot modify the original
func (c Counter) IncrementValue() {
    c.value++ // This doesn't affect the original
    fmt.Println("IncrementValue called (value receiver)")
}

// Pointer receiver - can modify the original
func (c *Counter) IncrementPointer() {
    c.value++ // This modifies the original
    fmt.Println("IncrementPointer called (pointer receiver)")
}

// Value receiver - can read but not modify
func (c Counter) GetValue() int {
    return c.value
}

// Pointer receiver - can modify
func (c *Counter) Reset() {
    c.value = 0
    fmt.Println("Reset called (pointer receiver)")
}
```

**Output:**
```
Initial value: 0
IncrementValue called (value receiver)
After IncrementValue: 0
IncrementPointer called (pointer receiver)
After IncrementPointer: 1
Reset called (pointer receiver)
After Reset: 0
```

### Example 3: Methods on Different Types

```go
package main

import "fmt"

// Custom type for temperature
type Celsius float64
type Fahrenheit float64

func main() {
    c := Celsius(25.0)
    f := Fahrenheit(77.0)
    
    fmt.Printf("Temperature in Celsius: %.2f°C\n", c)
    fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", f)
    
    // Convert temperatures
    fahrenheit := c.ToFahrenheit()
    celsius := f.ToCelsius()
    
    fmt.Printf("25°C = %.2f°F\n", fahrenheit)
    fmt.Printf("77°F = %.2f°C\n", celsius)
    
    // Check if temperatures are hot
    fmt.Printf("Is 25°C hot? %t\n", c.IsHot())
    fmt.Printf("Is 77°F hot? %t\n", f.IsHot())
}

// Method on Celsius type
func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

// Method on Celsius type
func (c Celsius) IsHot() bool {
    return c > 30
}

// Method on Fahrenheit type
func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// Method on Fahrenheit type
func (f Fahrenheit) IsHot() bool {
    return f > 86
}
```

**Output:**
```
Temperature in Celsius: 25.00°C
Temperature in Fahrenheit: 77.00°F
25°C = 77.00°F
77°F = 25.00°C
Is 25°C hot? false
Is 77°F hot? false
```

## Advanced Examples

### Example 4: Methods with Multiple Parameters

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    fmt.Printf("Rectangle: %.2f x %.2f\n", rect.Width, rect.Height)
    fmt.Printf("Area: %.2f\n", rect.Area())
    fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
    
    // Scale the rectangle
    rect.Scale(2.0)
    fmt.Printf("After scaling by 2: %.2f x %.2f\n", rect.Width, rect.Height)
    
    // Check if rectangle is square
    fmt.Printf("Is square? %t\n", rect.IsSquare())
    
    // Compare with another rectangle
    otherRect := Rectangle{Width: 20, Height: 10}
    fmt.Printf("Is equal to 20x10? %t\n", rect.Equals(otherRect))
}

// Method to calculate area
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method to calculate perimeter
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Method to scale rectangle
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Method to check if rectangle is square
func (r Rectangle) IsSquare() bool {
    return r.Width == r.Height
}

// Method to compare rectangles
func (r Rectangle) Equals(other Rectangle) bool {
    return r.Width == other.Width && r.Height == other.Height
}
```

**Output:**
```
Rectangle: 10.00 x 5.00
Area: 50.00
Perimeter: 30.00
After scaling by 2: 20.00 x 10.00
Is square? false
Is equal to 20x10? true
```

### Example 5: Methods with Interfaces

```go
package main

import "fmt"

// Interface definition
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

type Square struct {
    Side float64
}

func main() {
    // Create shapes
    circle := Circle{Radius: 5}
    square := Square{Side: 4}
    
    // Use methods directly
    fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", 
        circle.Area(), circle.Perimeter())
    fmt.Printf("Square - Area: %.2f, Perimeter: %.2f\n", 
        square.Area(), square.Perimeter())
    
    // Use through interface
    shapes := []Shape{circle, square}
    for i, shape := range shapes {
        fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n", 
            i+1, shape.Area(), shape.Perimeter())
    }
}

// Methods for Circle
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Methods for Square
func (s Square) Area() float64 {
    return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
    return 4 * s.Side
}
```

**Output:**
```
Circle - Area: 78.54, Perimeter: 31.42
Square - Area: 16.00, Perimeter: 16.00
Shape 1 - Area: 78.54, Perimeter: 31.42
Shape 2 - Area: 16.00, Perimeter: 16.00
```

### Example 6: Methods with Error Handling

```go
package main

import (
    "errors"
    "fmt"
)

type BankAccount struct {
    AccountNumber string
    Balance       float64
    Owner         string
}

func main() {
    account := BankAccount{
        AccountNumber: "123456789",
        Balance:       1000.0,
        Owner:         "Alice",
    }
    
    fmt.Printf("Account: %s, Balance: $%.2f\n", account.AccountNumber, account.Balance)
    
    // Deposit money
    if err := account.Deposit(500.0); err != nil {
        fmt.Printf("Deposit error: %v\n", err)
    } else {
        fmt.Printf("After deposit: $%.2f\n", account.Balance)
    }
    
    // Withdraw money
    if err := account.Withdraw(200.0); err != nil {
        fmt.Printf("Withdrawal error: %v\n", err)
    } else {
        fmt.Printf("After withdrawal: $%.2f\n", account.Balance)
    }
    
    // Try to withdraw more than balance
    if err := account.Withdraw(2000.0); err != nil {
        fmt.Printf("Withdrawal error: %v\n", err)
    } else {
        fmt.Printf("After withdrawal: $%.2f\n", account.Balance)
    }
    
    // Get account info
    info := account.GetInfo()
    fmt.Printf("Account info: %s\n", info)
}

// Method to deposit money
func (a *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }
    a.Balance += amount
    return nil
}

// Method to withdraw money
func (a *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }
    if amount > a.Balance {
        return errors.New("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

// Method to get account information
func (a BankAccount) GetInfo() string {
    return fmt.Sprintf("Account %s: %s has $%.2f", 
        a.AccountNumber, a.Owner, a.Balance)
}
```

**Output:**
```
Account: 123456789, Balance: $1000.00
After deposit: $1500.00
After withdrawal: $1300.00
Withdrawal error: insufficient funds
Account info: Account 123456789: Alice has $1300.00
```

## Method Sets and Rules

### Example 7: Method Sets

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    pPtr := &p
    
    // Both value and pointer can call value receiver methods
    p.Speak()      // Value calling value receiver method
    pPtr.Speak()   // Pointer calling value receiver method
    
    // Both value and pointer can call pointer receiver methods
    p.GrowOlder()  // Value calling pointer receiver method
    pPtr.GrowOlder() // Pointer calling pointer receiver method
    
    fmt.Printf("Final age: %d\n", p.Age)
}

// Value receiver method
func (p Person) Speak() {
    fmt.Printf("Hello, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

// Pointer receiver method
func (p *Person) GrowOlder() {
    p.Age++
    fmt.Printf("%s is now %d years old.\n", p.Name, p.Age)
}
```

**Output:**
```
Hello, I'm Alice and I'm 30 years old.
Hello, I'm Alice and I'm 30 years old.
Alice is now 31 years old.
Alice is now 32 years old.
Final age: 32
```

## Important Rules and Restrictions

### 1. Package Restriction
- The receiver type must be defined in the same package as the method
- You cannot define methods on types from other packages
- You cannot define methods on built-in types (int, string, etc.)

### 2. Method Sets
- **Value receiver methods**: Can be called on both value and pointer
- **Pointer receiver methods**: Can be called on both value and pointer
- Go automatically handles the conversion

### 3. Method Overloading
- You can have methods with the same name on different types
- You cannot have methods with the same name on the same type with different parameters

## Common Patterns and Best Practices

### Pattern 1: Getter and Setter Methods
```go
type User struct {
    name  string
    email string
}

func (u *User) GetName() string {
    return u.name
}

func (u *User) SetName(name string) {
    u.name = name
}

func (u *User) GetEmail() string {
    return u.email
}

func (u *User) SetEmail(email string) {
    u.email = email
}
```

### Pattern 2: Validation Methods
```go
type Product struct {
    Name  string
    Price float64
}

func (p Product) IsValid() bool {
    return p.Name != "" && p.Price > 0
}

func (p Product) Validate() error {
    if p.Name == "" {
        return errors.New("product name cannot be empty")
    }
    if p.Price <= 0 {
        return errors.New("product price must be positive")
    }
    return nil
}
```

### Pattern 3: String Representation
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}
```

## Methods vs Functions Comparison

| Aspect | Methods | Functions |
|--------|---------|-----------|
| **Receiver** | Contains receiver | No receiver |
| **Parameters** | Can accept both value and pointer receivers | Cannot accept both value and pointer |
| **Overloading** | Same name on different types allowed | Same name not allowed |
| **Type-specific** | Attached to specific type | Independent |
| **Access** | Can access receiver fields | Cannot access type fields |

## Common Mistakes to Avoid

### Mistake 1: Defining Methods on Built-in Types
```go
// ❌ This will cause a compile error
func (i int) Double() int {
    return i * 2
}

// ✅ Use a custom type instead
type MyInt int

func (i MyInt) Double() MyInt {
    return i * 2
}
```

### Mistake 2: Defining Methods on Types from Other Packages
```go
// ❌ This will cause a compile error
func (f os.File) CustomMethod() {
    // ...
}

// ✅ Use composition or embedding instead
type MyFile struct {
    *os.File
}

func (f MyFile) CustomMethod() {
    // ...
}
```

### Mistake 3: Using Value Receivers When Modification is Needed
```go
// ❌ This won't modify the original
func (p Person) SetAge(age int) {
    p.Age = age // This doesn't affect the original
}

// ✅ Use pointer receiver for modification
func (p *Person) SetAge(age int) {
    p.Age = age // This modifies the original
}
```

## Performance Considerations

### Example 8: Performance Comparison

```go
package main

import "fmt"

type LargeStruct struct {
    Data [1000]int
}

// Value receiver method
func (l LargeStruct) ProcessValue() int {
    sum := 0
    for _, v := range l.Data {
        sum += v
    }
    return sum
}

// Pointer receiver method
func (l *LargeStruct) ProcessPointer() int {
    sum := 0
    for _, v := range l.Data {
        sum += v
    }
    return sum
}

func main() {
    large := LargeStruct{}
    for i := range large.Data {
        large.Data[i] = i
    }
    
    // Both methods will have similar performance for this operation
    // but pointer receiver avoids copying the struct
    result1 := large.ProcessValue()
    result2 := large.ProcessPointer()
    
    fmt.Printf("Value receiver result: %d\n", result1)
    fmt.Printf("Pointer receiver result: %d\n", result2)
}
```

## Conclusion

Methods in Go provide:

1. **Type-specific behavior**: Attach functionality to types
2. **Code organization**: Group related functionality with data
3. **Interface implementation**: Enable polymorphism
4. **Encapsulation**: Control access to type fields
5. **Flexibility**: Support both value and pointer receivers

Understanding methods is essential for writing object-oriented Go code and implementing interfaces effectively. Methods are the foundation for creating clean, organized, and maintainable Go programs.
