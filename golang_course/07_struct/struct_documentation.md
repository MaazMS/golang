# Go Structs - Complete Guide

## Overview

A struct in Go is a composite data type that allows you to group together values of different types under a single name. Structs are similar to classes in object-oriented languages but without inheritance. They support composition and are the primary way to create custom types in Go.

## Key Features

1. **Composite data type**: Groups different types together
2. **Named fields**: Each field has a name and type
3. **Value type**: Structs are copied when assigned or passed to functions
4. **Composition support**: Can embed other structs
5. **Method support**: Can have methods attached
6. **Zero initialization**: Fields are initialized to their zero values
7. **No inheritance**: Go doesn't support inheritance, but uses composition

## Why Use Structs?

### Advantages
- **Data organization**: Group related data together
- **Type safety**: Compile-time type checking
- **Code clarity**: Self-documenting code structure
- **Method attachment**: Can define methods on structs
- **Composition**: Build complex types from simpler ones

### When to Use Structs
- Grouping related data
- Creating custom types
- Building domain models
- API responses and requests
- Configuration objects

## Basic Syntax

### Struct Declaration

```go
type StructName struct {
    Field1 DataType
    Field2 DataType
    Field3 DataType
}
```

### Field Naming Rules
- Field names must be unique within a struct
- Field names that start with uppercase letters are exported (public)
- Field names that start with lowercase letters are unexported (private)
- Embedded fields don't need explicit names

## Basic Examples

### Example 1: Simple Struct

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Method 1: Positional initialization
    p1 := Person{"Alice", 25, "New York"}
    fmt.Println("Person 1:", p1)

    // Method 2: Named field initialization
    p2 := Person{
        Name: "Bob",
        Age:  30,
        City: "London",
    }
    fmt.Println("Person 2:", p2)

    // Method 3: Partial initialization (remaining fields are zero values)
    p3 := Person{Name: "Charlie"}
    fmt.Println("Person 3:", p3)

    // Accessing fields
    fmt.Printf("Name: %s, Age: %d, City: %s\n", p1.Name, p1.Age, p1.City)

    // Modifying fields
    p1.Age = 26
    p1.City = "Boston"
    fmt.Println("Modified Person 1:", p1)
}
```

### Example 2: Different Data Types in Struct

```go
package main

import "fmt"

type Employee struct {
    ID          int
    Name        string
    Salary      float64
    IsActive    bool
    Skills      []string
    Contact     map[string]string
}

func main() {
    emp := Employee{
        ID:       1001,
        Name:     "John Doe",
        Salary:   75000.50,
        IsActive: true,
        Skills:   []string{"Go", "Python", "JavaScript"},
        Contact: map[string]string{
            "email": "john@example.com",
            "phone": "+1-555-0123",
        },
    }

    fmt.Printf("Employee: %+v\n", emp)
    fmt.Printf("Name: %s\n", emp.Name)
    fmt.Printf("Skills: %v\n", emp.Skills)
    fmt.Printf("Email: %s\n", emp.Contact["email"])
}
```

### Example 3: Struct with Methods

```go
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Method with value receiver (doesn't modify original)
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    fmt.Printf("Rectangle: %+v\n", rect)
    fmt.Printf("Area: %.2f\n", rect.Area())
    fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
    
    // Scale the rectangle
    rect.Scale(2)
    fmt.Printf("Scaled Rectangle: %+v\n", rect)
    fmt.Printf("New Area: %.2f\n", rect.Area())
}
```

## Embedded Structs (Composition)

### Example 4: Basic Embedded Structs

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person        // Embedded struct
    ID     int
    Salary float64
}

type Manager struct {
    Employee      // Embedded struct
    Department string
    TeamSize   int
}

func main() {
    // Create a manager
    mgr := Manager{
        Employee: Employee{
            Person: Person{
                Name: "Alice",
                Age:  35,
            },
            ID:     1001,
            Salary: 100000,
        },
        Department: "Engineering",
        TeamSize:   10,
    }

    // Access fields (promoted fields)
    fmt.Printf("Name: %s\n", mgr.Name)           // Promoted from Person
    fmt.Printf("Age: %d\n", mgr.Age)             // Promoted from Person
    fmt.Printf("ID: %d\n", mgr.ID)               // Promoted from Employee
    fmt.Printf("Salary: %.2f\n", mgr.Salary)     // Promoted from Employee
    fmt.Printf("Department: %s\n", mgr.Department)
    fmt.Printf("Team Size: %d\n", mgr.TeamSize)

    // Access through embedded struct
    fmt.Printf("Person: %+v\n", mgr.Person)
    fmt.Printf("Employee: %+v\n", mgr.Employee)
}
```

### Example 5: Embedded Structs with Methods

```go
package main

import "fmt"

type Animal struct {
    Name string
    Age  int
}

func (a Animal) Speak() {
    fmt.Printf("%s makes a sound\n", a.Name)
}

func (a Animal) Info() {
    fmt.Printf("Name: %s, Age: %d\n", a.Name, a.Age)
}

type Dog struct {
    Animal
    Breed string
}

func (d Dog) Speak() {
    fmt.Printf("%s barks\n", d.Name)
}

func (d Dog) Fetch() {
    fmt.Printf("%s fetches the ball\n", d.Name)
}

type Cat struct {
    Animal
    Color string
}

func (c Cat) Speak() {
    fmt.Printf("%s meows\n", c.Name)
}

func (c Cat) Purr() {
    fmt.Printf("%s purrs\n", c.Name)
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy", Age: 3},
        Breed:  "Golden Retriever",
    }

    cat := Cat{
        Animal: Animal{Name: "Whiskers", Age: 2},
        Color:  "Orange",
    }

    // Method calls
    dog.Speak()    // Overrides Animal.Speak()
    dog.Fetch()    // Dog-specific method
    dog.Info()     // Inherited from Animal

    cat.Speak()    // Overrides Animal.Speak()
    cat.Purr()     // Cat-specific method
    cat.Info()     // Inherited from Animal
}
```

## Anonymous Structs

### Example 6: Anonymous Structs

```go
package main

import "fmt"

func main() {
    // Anonymous struct
    person := struct {
        Name string
        Age  int
        City string
    }{
        Name: "Alice",
        Age:  25,
        City: "New York",
    }

    fmt.Printf("Anonymous struct: %+v\n", person)

    // Anonymous struct in function parameter
    processPerson(struct {
        Name string
        Age  int
    }{
        Name: "Bob",
        Age:  30,
    })

    // Anonymous struct slice
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 35},
    }

    for _, p := range people {
        fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
    }

    // Anonymous struct map
    config := map[string]struct {
        Host string
        Port int
    }{
        "database": {"localhost", 5432},
        "redis":    {"localhost", 6379},
        "api":      {"localhost", 8080},
    }

    for service, cfg := range config {
        fmt.Printf("%s: %s:%d\n", service, cfg.Host, cfg.Port)
    }
}

func processPerson(p struct {
    Name string
    Age  int
}) {
    fmt.Printf("Processing: %s, %d years old\n", p.Name, p.Age)
}
```

## Slice of Structs

### Example 7: Slice of Structs

```go
package main

import (
    "fmt"
    "sort"
)

type Student struct {
    ID    int
    Name  string
    Grade float64
}

type Students []Student

func main() {
    // Create slice of structs
    students := Students{
        {ID: 1, Name: "Alice", Grade: 85.5},
        {ID: 2, Name: "Bob", Grade: 92.0},
        {ID: 3, Name: "Charlie", Grade: 78.5},
        {ID: 4, Name: "Diana", Grade: 95.0},
    }

    fmt.Println("Original students:")
    printStudents(students)

    // Sort by grade (descending)
    sort.Slice(students, func(i, j int) bool {
        return students[i].Grade > students[j].Grade
    })

    fmt.Println("\nSorted by grade:")
    printStudents(students)

    // Find student by ID
    id := 2
    if student, found := findStudentByID(students, id); found {
        fmt.Printf("\nFound student with ID %d: %+v\n", id, student)
    } else {
        fmt.Printf("\nStudent with ID %d not found\n", id)
    }

    // Calculate average grade
    avg := calculateAverageGrade(students)
    fmt.Printf("Average grade: %.2f\n", avg)

    // Filter students with grade >= 90
    highPerformers := filterStudentsByGrade(students, 90.0)
    fmt.Println("\nHigh performers (grade >= 90):")
    printStudents(highPerformers)
}

func printStudents(students Students) {
    for _, s := range students {
        fmt.Printf("ID: %d, Name: %s, Grade: %.2f\n", s.ID, s.Name, s.Grade)
    }
}

func findStudentByID(students Students, id int) (Student, bool) {
    for _, s := range students {
        if s.ID == id {
            return s, true
        }
    }
    return Student{}, false
}

func calculateAverageGrade(students Students) float64 {
    if len(students) == 0 {
        return 0
    }
    total := 0.0
    for _, s := range students {
        total += s.Grade
    }
    return total / float64(len(students))
}

func filterStudentsByGrade(students Students, minGrade float64) Students {
    var result Students
    for _, s := range students {
        if s.Grade >= minGrade {
            result = append(result, s)
        }
    }
    return result
}
```

## Advanced Examples

### Example 8: Struct with JSON Tags

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`                    // Don't include in JSON
    IsActive bool   `json:"is_active"`
    Profile  struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
        Age       int    `json:"age"`
    } `json:"profile"`
}

func main() {
    user := User{
        ID:       1,
        Username: "johndoe",
        Email:    "john@example.com",
        Password: "secret123",
        IsActive: true,
        Profile: struct {
            FirstName string `json:"first_name"`
            LastName  string `json:"last_name"`
            Age       int    `json:"age"`
        }{
            FirstName: "John",
            LastName:  "Doe",
            Age:       30,
        },
    }

    // Convert to JSON
    jsonData, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        fmt.Printf("Error marshaling JSON: %v\n", err)
        return
    }
    fmt.Println("JSON representation:")
    fmt.Println(string(jsonData))

    // Convert from JSON
    jsonString := `{
        "id": 2,
        "username": "janedoe",
        "email": "jane@example.com",
        "is_active": true,
        "profile": {
            "first_name": "Jane",
            "last_name": "Doe",
            "age": 25
        }
    }`

    var newUser User
    err = json.Unmarshal([]byte(jsonString), &newUser)
    if err != nil {
        fmt.Printf("Error unmarshaling JSON: %v\n", err)
        return
    }
    fmt.Printf("\nUnmarshaled user: %+v\n", newUser)
}
```

### Example 9: Struct with Validation

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

type Product struct {
    ID          int
    Name        string
    Price       float64
    Description string
    Category    string
}

func (p Product) Validate() error {
    var errors []string

    if p.ID <= 0 {
        errors = append(errors, "ID must be positive")
    }

    if strings.TrimSpace(p.Name) == "" {
        errors = append(errors, "Name cannot be empty")
    }

    if p.Price < 0 {
        errors = append(errors, "Price cannot be negative")
    }

    if strings.TrimSpace(p.Category) == "" {
        errors = append(errors, "Category cannot be empty")
    }

    if len(errors) > 0 {
        return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
    }

    return nil
}

func (p *Product) SetName(name string) error {
    if strings.TrimSpace(name) == "" {
        return errors.New("name cannot be empty")
    }
    p.Name = strings.TrimSpace(name)
    return nil
}

func (p *Product) SetPrice(price float64) error {
    if price < 0 {
        return errors.New("price cannot be negative")
    }
    p.Price = price
    return nil
}

func main() {
    // Valid product
    product1 := Product{
        ID:          1,
        Name:        "Laptop",
        Price:       999.99,
        Description: "High-performance laptop",
        Category:    "Electronics",
    }

    if err := product1.Validate(); err != nil {
        fmt.Printf("Validation error: %v\n", err)
    } else {
        fmt.Println("Product 1 is valid")
    }

    // Invalid product
    product2 := Product{
        ID:          -1,
        Name:        "",
        Price:       -50.0,
        Description: "Invalid product",
        Category:    "",
    }

    if err := product2.Validate(); err != nil {
        fmt.Printf("Validation error: %v\n", err)
    }

    // Using setter methods
    product3 := Product{ID: 3}
    if err := product3.SetName("  "); err != nil {
        fmt.Printf("SetName error: %v\n", err)
    }

    if err := product3.SetPrice(-100); err != nil {
        fmt.Printf("SetPrice error: %v\n", err)
    }
}
```

### Example 10: Struct with Interface

```go
package main

import "fmt"

type Shape interface {
    Area() float64
    Perimeter() float64
}

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

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    shapes := []Shape{
        Rectangle{Width: 10, Height: 5},
        Circle{Radius: 7},
        Rectangle{Width: 3, Height: 4},
    }

    for i, shape := range shapes {
        fmt.Printf("Shape %d: ", i+1)
        printShapeInfo(shape)
    }
}
```

## Struct Patterns and Best Practices

### Pattern 1: Builder Pattern

```go
type UserBuilder struct {
    user User
}

func NewUserBuilder() *UserBuilder {
    return &UserBuilder{}
}

func (b *UserBuilder) SetID(id int) *UserBuilder {
    b.user.ID = id
    return b
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
    b.user.Name = name
    return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
    b.user.Email = email
    return b
}

func (b *UserBuilder) Build() User {
    return b.user
}

func main() {
    user := NewUserBuilder().
        SetID(1).
        SetName("John Doe").
        SetEmail("john@example.com").
        Build()
    
    fmt.Printf("Built user: %+v\n", user)
}
```

### Pattern 2: Factory Pattern

```go
type AnimalType int

const (
    Dog AnimalType = iota
    Cat
    Bird
)

type Animal struct {
    Name string
    Type AnimalType
}

func NewAnimal(name string, animalType AnimalType) *Animal {
    return &Animal{
        Name: name,
        Type: animalType,
    }
}

func NewDog(name string) *Animal {
    return NewAnimal(name, Dog)
}

func NewCat(name string) *Animal {
    return NewAnimal(name, Cat)
}

func NewBird(name string) *Animal {
    return NewAnimal(name, Bird)
}
```

### Pattern 3: Struct Composition

```go
type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

type Contact struct {
    Email string
    Phone string
}

type Person struct {
    Name    string
    Age     int
    Address Address
    Contact Contact
}

func (p Person) FullAddress() string {
    return fmt.Sprintf("%s, %s, %s %s", p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

func (p Person) ContactInfo() string {
    return fmt.Sprintf("Email: %s, Phone: %s", p.Contact.Email, p.Contact.Phone)
}
```

## Performance Considerations

### Example 11: Struct vs Pointer Performance

```go
package main

import (
    "fmt"
    "time"
)

type LargeStruct struct {
    Data [1000]int
}

func processByValue(s LargeStruct) int {
    sum := 0
    for _, v := range s.Data {
        sum += v
    }
    return sum
}

func processByPointer(s *LargeStruct) int {
    sum := 0
    for _, v := range s.Data {
        sum += v
    }
    return sum
}

func main() {
    largeStruct := LargeStruct{}
    for i := range largeStruct.Data {
        largeStruct.Data[i] = i
    }

    // Benchmark by value
    start := time.Now()
    for i := 0; i < 100000; i++ {
        processByValue(largeStruct)
    }
    valueTime := time.Since(start)

    // Benchmark by pointer
    start = time.Now()
    for i := 0; i < 100000; i++ {
        processByPointer(&largeStruct)
    }
    pointerTime := time.Since(start)

    fmt.Printf("By value time: %v\n", valueTime)
    fmt.Printf("By pointer time: %v\n", pointerTime)
    fmt.Printf("Pointer is %.2fx faster\n", float64(valueTime)/float64(pointerTime))
}
```

## Error Handling with Structs

### Example 12: Struct with Error Handling

```go
package main

import (
    "errors"
    "fmt"
)

type Database struct {
    Host     string
    Port     int
    Username string
    Password string
}

func (d *Database) Connect() error {
    if d.Host == "" {
        return errors.New("host cannot be empty")
    }
    if d.Port <= 0 || d.Port > 65535 {
        return errors.New("port must be between 1 and 65535")
    }
    if d.Username == "" {
        return errors.New("username cannot be empty")
    }
    if d.Password == "" {
        return errors.New("password cannot be empty")
    }
    
    // Simulate connection
    fmt.Printf("Connecting to %s:%d as %s\n", d.Host, d.Port, d.Username)
    return nil
}

func (d *Database) Query(sql string) ([]map[string]interface{}, error) {
    if sql == "" {
        return nil, errors.New("SQL query cannot be empty")
    }
    
    // Simulate query
    fmt.Printf("Executing query: %s\n", sql)
    return []map[string]interface{}{
        {"id": 1, "name": "Alice"},
        {"id": 2, "name": "Bob"},
    }, nil
}

func main() {
    db := &Database{
        Host:     "localhost",
        Port:     5432,
        Username: "admin",
        Password: "secret",
    }

    if err := db.Connect(); err != nil {
        fmt.Printf("Connection error: %v\n", err)
        return
    }

    results, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Printf("Query error: %v\n", err)
        return
    }

    fmt.Printf("Query results: %+v\n", results)
}
```

## Conclusion

Structs in Go are powerful and flexible data structures that provide:

1. **Data organization**: Group related data together
2. **Type safety**: Compile-time type checking
3. **Composition**: Build complex types from simpler ones
4. **Method attachment**: Define behavior on data
5. **Value semantics**: Predictable copying behavior
6. **Interface support**: Polymorphic behavior

Understanding structs is essential for effective Go programming, as they are the primary way to create custom types and organize data in Go applications. They provide a solid foundation for building complex data structures and domain models.
