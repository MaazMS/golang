# Go Interfaces and Polymorphism - Complete Guide

## Overview

Interfaces in Go are a powerful feature that allows values to be of more than one type. They define a set of method signatures that types must implement to satisfy the interface. Go interfaces are different from other languages in that they are implicit - a type automatically implements an interface if it has all the required methods. This enables polymorphism, where the same method name can be used for different types.

## Key Features

1. **Implicit implementation**: Types automatically implement interfaces
2. **Method signatures**: Define required methods without implementation
3. **Polymorphism**: Same method name for different types
4. **Empty interface**: Can hold any type
5. **Type assertions**: Extract concrete types from interfaces
6. **Composition**: Interfaces can be composed of other interfaces

## Why Use Interfaces?

### Benefits
- **Polymorphism**: Treat different types uniformly
- **Code reusability**: Write functions that work with multiple types
- **Loose coupling**: Depend on behavior, not concrete types
- **Testability**: Easy to mock and test
- **Flexibility**: Easy to add new types that implement interfaces

### When to Use Interfaces
- Define common behavior across types
- Create flexible APIs
- Implement dependency injection
- Enable testing with mocks
- Create plugin architectures

## Basic Syntax

### Interface Declaration
```go
type InterfaceName interface {
    Method1(param1 type1) returnType1
    Method2(param2 type2) returnType2
    // more method signatures
}
```

### Key Components
- **type**: Keyword to declare interface
- **InterfaceName**: Name of the interface
- **interface**: Keyword indicating interface type
- **Method signatures**: Required methods without implementation

## Basic Examples

### Example 1: Simple Interface

```go
package main

import "fmt"

// Define interface
type Speaker interface {
    Speak() string
}

// Person type
type Person struct {
    Name string
    Age  int
}

// Person implements Speaker interface
func (p Person) Speak() string {
    return fmt.Sprintf("Hello, I'm %s and I'm %d years old", p.Name, p.Age)
}

// Dog type
type Dog struct {
    Name string
    Breed string
}

// Dog implements Speaker interface
func (d Dog) Speak() string {
    return fmt.Sprintf("Woof! I'm %s, a %s", d.Name, d.Breed)
}

func main() {
    // Create instances
    person := Person{Name: "Alice", Age: 30}
    dog := Dog{Name: "Buddy", Breed: "Golden Retriever"}
    
    // Use interface
    var speaker Speaker
    speaker = person
    fmt.Println(speaker.Speak())
    
    speaker = dog
    fmt.Println(speaker.Speak())
    
    // Use with function
    makeSpeak(person)
    makeSpeak(dog)
}

// Function that accepts interface
func makeSpeak(s Speaker) {
    fmt.Println("Speaking:", s.Speak())
}
```

**Output:**
```
Hello, I'm Alice and I'm 30 years old
Woof! I'm Buddy, a Golden Retriever
Speaking: Hello, I'm Alice and I'm 30 years old
Speaking: Woof! I'm Buddy, a Golden Retriever
```

### Example 2: Multiple Methods Interface

```go
package main

import "fmt"

// Shape interface with multiple methods
type Shape interface {
    Area() float64
    Perimeter() float64
    String() string
}

// Rectangle type
type Rectangle struct {
    Width  float64
    Height float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Circle type
type Circle struct {
    Radius float64
}

// Circle implements Shape interface
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

func main() {
    shapes := []Shape{
        Rectangle{Width: 10, Height: 5},
        Circle{Radius: 7},
        Rectangle{Width: 3, Height: 4},
    }
    
    for i, shape := range shapes {
        fmt.Printf("Shape %d: %s\n", i+1, shape)
        fmt.Printf("  Area: %.2f\n", shape.Area())
        fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
        fmt.Println()
    }
}
```

**Output:**
```
Shape 1: Rectangle{Width: 10.00, Height: 5.00}
  Area: 50.00
  Perimeter: 30.00

Shape 2: Circle{Radius: 7.00}
  Area: 153.94
  Perimeter: 43.98

Shape 3: Rectangle{Width: 3.00, Height: 4.00}
  Area: 12.00
  Perimeter: 14.00
```

## Advanced Examples

### Example 3: Type Assertions and Type Switches

```go
package main

import "fmt"

// Animal interface
type Animal interface {
    MakeSound() string
    Move() string
}

// Dog type
type Dog struct {
    Name string
}

func (d Dog) MakeSound() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running"
}

// Cat type
type Cat struct {
    Name string
}

func (c Cat) MakeSound() string {
    return "Meow!"
}

func (c Cat) Move() string {
    return "Sneaking"
}

// Bird type
type Bird struct {
    Name string
}

func (b Bird) MakeSound() string {
    return "Tweet!"
}

func (b Bird) Move() string {
    return "Flying"
}

func main() {
    animals := []Animal{
        Dog{Name: "Buddy"},
        Cat{Name: "Whiskers"},
        Bird{Name: "Tweety"},
    }
    
    for _, animal := range animals {
        // Type switch
        switch a := animal.(type) {
        case Dog:
            fmt.Printf("Dog %s: %s %s\n", a.Name, a.MakeSound(), a.Move())
        case Cat:
            fmt.Printf("Cat %s: %s %s\n", a.Name, a.MakeSound(), a.Move())
        case Bird:
            fmt.Printf("Bird %s: %s %s\n", a.Name, a.MakeSound(), a.Move())
        default:
            fmt.Printf("Unknown animal: %s %s\n", a.MakeSound(), a.Move())
        }
    }
    
    // Type assertion
    var animal Animal = Dog{Name: "Rex"}
    if dog, ok := animal.(Dog); ok {
        fmt.Printf("Type assertion successful: %s\n", dog.Name)
    }
    
    // Type assertion with error handling
    if cat, ok := animal.(Cat); ok {
        fmt.Printf("This won't print: %s\n", cat.Name)
    } else {
        fmt.Println("Type assertion failed - not a cat")
    }
}
```

**Output:**
```
Dog Buddy: Woof! Running
Cat Whiskers: Meow! Sneaking
Bird Tweety: Tweet! Flying
Type assertion successful: Rex
Type assertion failed - not a cat
```

### Example 4: Empty Interface

```go
package main

import "fmt"

func main() {
    // Empty interface can hold any type
    var data interface{}
    
    // Assign different types
    data = 42
    printType(data)
    
    data = "Hello, World!"
    printType(data)
    
    data = 3.14159
    printType(data)
    
    data = []int{1, 2, 3, 4, 5}
    printType(data)
    
    data = map[string]int{"a": 1, "b": 2}
    printType(data)
    
    // Use with function
    processData(42)
    processData("Hello")
    processData([]string{"a", "b", "c"})
}

func printType(data interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", data, data)
}

func processData(data interface{}) {
    switch v := data.(type) {
    case int:
        fmt.Printf("Processing integer: %d\n", v)
    case string:
        fmt.Printf("Processing string: %s\n", v)
    case []string:
        fmt.Printf("Processing string slice: %v\n", v)
    default:
        fmt.Printf("Processing unknown type: %T\n", v)
    }
}
```

**Output:**
```
Type: int, Value: 42
Type: string, Value: Hello, World!
Type: float64, Value: 3.14159
Type: []int, Value: [1 2 3 4 5]
Type: map[string]int, Value: map[a:1 b:2]
Processing integer: 42
Processing string: Hello
Processing string slice: [a b c]
```

### Example 5: Interface Composition

```go
package main

import "fmt"

// Reader interface
type Reader interface {
    Read() string
}

// Writer interface
type Writer interface {
    Write(string)
}

// ReadWriter interface (composed of Reader and Writer)
type ReadWriter interface {
    Reader
    Writer
}

// File type
type File struct {
    name    string
    content string
}

// File implements Reader interface
func (f *File) Read() string {
    return f.content
}

// File implements Writer interface
func (f *File) Write(data string) {
    f.content += data
}

// Console type
type Console struct {
    buffer string
}

// Console implements Reader interface
func (c *Console) Read() string {
    return c.buffer
}

// Console implements Writer interface
func (c *Console) Write(data string) {
    c.buffer += data
    fmt.Printf("Console output: %s\n", data)
}

func main() {
    // Create instances
    file := &File{name: "test.txt", content: "Initial content\n"}
    console := &Console{}
    
    // Use as ReadWriter
    var rw ReadWriter
    rw = file
    rw.Write("New content\n")
    fmt.Printf("File content: %s", rw.Read())
    
    rw = console
    rw.Write("Hello from console\n")
    rw.Write("Another line\n")
    fmt.Printf("Console buffer: %s", rw.Read())
    
    // Use individual interfaces
    var reader Reader = file
    var writer Writer = console
    
    fmt.Printf("Reader says: %s", reader.Read())
    writer.Write("Direct write to console\n")
}
```

**Output:**
```
File content: Initial content
New content
Console output: Hello from console
Console output: Another line
Console buffer: Hello from console
Another line
Reader says: Initial content
New content
Console output: Direct write to console
```

## Polymorphism Examples

### Example 6: Polymorphic Functions

```go
package main

import "fmt"

// Drawer interface
type Drawer interface {
    Draw() string
}

// Rectangle type
type Rectangle struct {
    Width  int
    Height int
}

func (r Rectangle) Draw() string {
    return fmt.Sprintf("Drawing rectangle %dx%d", r.Width, r.Height)
}

// Circle type
type Circle struct {
    Radius int
}

func (c Circle) Draw() string {
    return fmt.Sprintf("Drawing circle with radius %d", c.Radius)
}

// Triangle type
type Triangle struct {
    Base   int
    Height int
}

func (t Triangle) Draw() string {
    return fmt.Sprintf("Drawing triangle with base %d and height %d", t.Base, t.Height)
}

func main() {
    shapes := []Drawer{
        Rectangle{Width: 10, Height: 5},
        Circle{Radius: 7},
        Triangle{Base: 8, Height: 6},
    }
    
    // Polymorphic function
    drawAll(shapes)
    
    // Individual drawing
    var drawer Drawer = Rectangle{Width: 20, Height: 10}
    fmt.Println(drawer.Draw())
}

// Polymorphic function - works with any type that implements Drawer
func drawAll(shapes []Drawer) {
    for i, shape := range shapes {
        fmt.Printf("Shape %d: %s\n", i+1, shape.Draw())
    }
}
```

**Output:**
```
Shape 1: Drawing rectangle 10x5
Shape 2: Drawing circle with radius 7
Shape 3: Drawing triangle with base 8 and height 6
Drawing rectangle 20x10
```

### Example 7: Database Operations with Interfaces

```go
package main

import "fmt"

// User struct
type User struct {
    ID   int
    Name string
    Email string
}

// Repository interface
type Repository interface {
    Save(user User) error
    FindByID(id int) (User, error)
    Delete(id int) error
}

// InMemoryRepository implements Repository
type InMemoryRepository struct {
    users map[int]User
}

func NewInMemoryRepository() *InMemoryRepository {
    return &InMemoryRepository{
        users: make(map[int]User),
    }
}

func (r *InMemoryRepository) Save(user User) error {
    r.users[user.ID] = user
    fmt.Printf("Saved user %d in memory\n", user.ID)
    return nil
}

func (r *InMemoryRepository) FindByID(id int) (User, error) {
    user, exists := r.users[id]
    if !exists {
        return User{}, fmt.Errorf("user %d not found", id)
    }
    fmt.Printf("Found user %d in memory\n", id)
    return user, nil
}

func (r *InMemoryRepository) Delete(id int) error {
    delete(r.users, id)
    fmt.Printf("Deleted user %d from memory\n", id)
    return nil
}

// DatabaseRepository implements Repository
type DatabaseRepository struct {
    connection string
}

func NewDatabaseRepository(conn string) *DatabaseRepository {
    return &DatabaseRepository{connection: conn}
}

func (r *DatabaseRepository) Save(user User) error {
    fmt.Printf("Saved user %d to database (%s)\n", user.ID, r.connection)
    return nil
}

func (r *DatabaseRepository) FindByID(id int) (User, error) {
    fmt.Printf("Found user %d in database (%s)\n", id, r.connection)
    return User{ID: id, Name: "Database User", Email: "db@example.com"}, nil
}

func (r *DatabaseRepository) Delete(id int) error {
    fmt.Printf("Deleted user %d from database (%s)\n", id, r.connection)
    return nil
}

// UserService uses Repository interface
type UserService struct {
    repo Repository
}

func NewUserService(repo Repository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) error {
    user := User{ID: 1, Name: name, Email: email}
    return s.repo.Save(user)
}

func (s *UserService) GetUser(id int) (User, error) {
    return s.repo.FindByID(id)
}

func main() {
    // Use in-memory repository
    memRepo := NewInMemoryRepository()
    userService := NewUserService(memRepo)
    
    userService.CreateUser("Alice", "alice@example.com")
    user, _ := userService.GetUser(1)
    fmt.Printf("User: %+v\n", user)
    
    fmt.Println()
    
    // Use database repository
    dbRepo := NewDatabaseRepository("postgres://localhost:5432/mydb")
    userService = NewUserService(dbRepo)
    
    userService.CreateUser("Bob", "bob@example.com")
    user, _ = userService.GetUser(1)
    fmt.Printf("User: %+v\n", user)
}
```

**Output:**
```
Saved user 1 in memory
Found user 1 in memory
User: {ID:1 Name:Alice Email:alice@example.com}

Saved user 1 to database (postgres://localhost:5432/mydb)
Found user 1 in database (postgres://localhost:5432/mydb)
User: {ID:1 Name:Database User Email:db@example.com}
```

## Important Rules and Best Practices

### 1. Interface Implementation
- Types automatically implement interfaces if they have all required methods
- No explicit declaration needed
- Interface satisfaction is checked at compile time

### 2. Method Sets
- Value receivers: Can be called on both value and pointer
- Pointer receivers: Can be called on both value and pointer
- Go automatically handles conversions

### 3. Empty Interface
- `interface{}` can hold any type
- Use type assertions or type switches to extract concrete types
- Avoid overusing empty interfaces

### 4. Interface Composition
- Interfaces can be composed of other interfaces
- Promotes code reusability
- Follows interface segregation principle

## Common Patterns

### Pattern 1: Dependency Injection
```go
type Logger interface {
    Log(message string)
}

type Service struct {
    logger Logger
}

func NewService(logger Logger) *Service {
    return &Service{logger: logger}
}
```

### Pattern 2: Plugin Architecture
```go
type Plugin interface {
    Name() string
    Execute() error
}

type PluginManager struct {
    plugins []Plugin
}

func (pm *PluginManager) Register(plugin Plugin) {
    pm.plugins = append(pm.plugins, plugin)
}
```

### Pattern 3: Strategy Pattern
```go
type PaymentProcessor interface {
    Process(amount float64) error
}

type CreditCardProcessor struct{}
func (c CreditCardProcessor) Process(amount float64) error {
    // Process credit card payment
    return nil
}

type PayPalProcessor struct{}
func (p PayPalProcessor) Process(amount float64) error {
    // Process PayPal payment
    return nil
}
```

## Common Mistakes to Avoid

### Mistake 1: Overusing Empty Interface
```go
// ❌ Bad: Overusing empty interface
func processData(data interface{}) {
    // Hard to work with
}

// ✅ Good: Use specific interfaces
func processData(data Reader) {
    // Easy to work with
}
```

### Mistake 2: Not Handling Type Assertions
```go
// ❌ Bad: Panic if type assertion fails
value := data.(string)

// ✅ Good: Handle type assertion safely
if value, ok := data.(string); ok {
    // Use value
} else {
    // Handle error
}
```

### Mistake 3: Interface Too Large
```go
// ❌ Bad: Interface with too many methods
type BadInterface interface {
    Method1()
    Method2()
    Method3()
    Method4()
    Method5()
    // ... many more
}

// ✅ Good: Small, focused interfaces
type Reader interface {
    Read() ([]byte, error)
}

type Writer interface {
    Write([]byte) (int, error)
}
```

## Performance Considerations

### Example 8: Interface Performance

```go
package main

import "fmt"

type Calculator interface {
    Add(a, b int) int
}

type SimpleCalculator struct{}

func (s SimpleCalculator) Add(a, b int) int {
    return a + b
}

func main() {
    calc := SimpleCalculator{}
    
    // Direct method call
    result1 := calc.Add(5, 3)
    fmt.Printf("Direct call: %d\n", result1)
    
    // Interface call
    var calcInterface Calculator = calc
    result2 := calcInterface.Add(5, 3)
    fmt.Printf("Interface call: %d\n", result2)
    
    // Interface calls have a small overhead
    // but are generally negligible for most applications
}
```

## Conclusion

Interfaces and polymorphism in Go provide:

1. **Flexibility**: Write code that works with multiple types
2. **Polymorphism**: Same method name for different types
3. **Loose coupling**: Depend on behavior, not concrete types
4. **Testability**: Easy to mock and test
5. **Composition**: Build complex interfaces from simpler ones

Understanding interfaces is essential for writing flexible, maintainable Go code. They enable powerful patterns like dependency injection, plugin architectures, and strategy patterns while maintaining type safety and performance.
