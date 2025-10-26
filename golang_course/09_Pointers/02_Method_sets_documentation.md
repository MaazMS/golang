## Method Sets

A method set is the collection of methods that are associated with a type. When you define methods for a type, those methods become part of the type's method set.

**IMPORTANT**: The method set of a type determines which interfaces that type implements. This is crucial for understanding polymorphism in Go.

### Method Set Rules

The method set depends on the receiver type:

| Receiver Type | Method Set Available For |
|---------------|-------------------------|
| `(t T)`       | Both `T` and `*T`       |
| `(t *T)`      | Only `*T`               |

### Understanding Method Sets

- **Value Receiver `(t T)`**: Methods with value receivers are available for both the type and its pointer
- **Pointer Receiver `(t *T)`**: Methods with pointer receivers are only available for the pointer type

### Example 1: Value Receiver with Value Type

```go
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// Value receiver method
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
	// Value receiver with value type
	c := circle{radius: 5}
	info(c) // Works: value receiver is available for value type
}

/* Output:
Area: 78.54
*/
```

### Example 2: Value Receiver with Pointer Type

```go
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// Value receiver method
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
	// Value receiver with pointer type
	c := circle{radius: 10}
	info(&c) // Works: value receiver is available for pointer type
}

/* Output:
Area: 314.16
*/
```

### Example 3: Pointer Receiver with Pointer Type

```go
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

// Value receiver method for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Pointer receiver method for square
func (s *square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
	// Pointer receiver with pointer type
	sq := square{length: 3.5}
	info(&sq) // Works: pointer receiver is available for pointer type
}

/* Output:
Area: 12.25
*/
```

### Example 4: What Doesn't Work

```go
package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type shape interface {
	area() float64
}

// Pointer receiver method
func (s *square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
	sq := square{length: 3.5}
	// info(sq) // This would cause a compile error!
	// Cannot use sq (type square) as type shape in argument to info:
	// square does not implement shape (area method has pointer receiver)
	
	info(&sq) // This works
}

/* Output:
Area: 12.25
*/
```

### Key Concepts:

1. **Method Set**: Collection of methods associated with a type
2. **Interface Implementation**: A type implements an interface if its method set contains all methods required by the interface
3. **Value Receiver**: Available for both value and pointer types
4. **Pointer Receiver**: Only available for pointer types
5. **Automatic Dereferencing**: Go automatically dereferences pointers when calling value receiver methods

### Best Practices:

1. **Consistency**: Use the same receiver type for all methods of a type
2. **Pointer Receivers**: Use when methods need to modify the receiver or when working with large structs
3. **Value Receivers**: Use for small structs or when methods don't need to modify the receiver
4. **Interface Design**: Design interfaces with value receiver methods when possible for maximum flexibility

### Common Pitfalls:

- **Mixing Receiver Types**: Can lead to confusion about which methods are available
- **Pointer Receiver Limitation**: Methods with pointer receivers are not available for value types
- **Interface Implementation**: Forgetting that pointer receivers limit interface implementation to pointer types only