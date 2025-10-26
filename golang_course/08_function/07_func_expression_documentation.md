## Function Expression

A function expression in Go is when you assign a function to a variable. This is also known as a **function literal** or **anonymous function**.

### Key Points:
1. You can assign a function to a variable
2. This is called a function expression or function literal
3. The function can be called using the variable name

### Example:

```go
package main

import "fmt"

func main() {
	// This is NOT a function expression - it's calling a function and assigning its return value
	// x := foo()  // This would assign the return value of foo() to x

	// This IS a function expression - assigning a function to a variable
	f := func() {
		fmt.Println("function f")
	}
	f() // Call the function using the variable

	// Function expression with parameters
	x := func(num int) {
		fmt.Println("function x with parameter:", num)
	}
	x(1997) // Call the function with an argument
}
```

### Explanation:
- `f := func() { ... }` creates an anonymous function and assigns it to variable `f`
- `x := func(num int) { ... }` creates an anonymous function with a parameter and assigns it to variable `x`
- You can call these functions using the variable names: `f()` and `x(1997)`

### Common Use Cases:
- **Callbacks**: Passing functions as arguments to other functions
- **Closures**: Functions that capture variables from their surrounding scope
- **Event handlers**: Assigning behavior to events
- **Functional programming**: Creating higher-order functions
