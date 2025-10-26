## Closures

A closure is a function that captures and retains access to variables from its surrounding lexical scope, even after the outer function has finished executing. This allows the inner function to "close over" the variables it needs.

**Important Note**: The examples below demonstrate variable scope rather than true closures. True closures in Go typically involve functions that capture variables from their enclosing scope.

## Variable Scope Examples

### Package-Level Variable Scope

Variables declared at the package level are accessible throughout the entire package.

```go
package main

import "fmt"

var x int

func main() {
	fmt.Println("Package variable:", x)
	x++
	fmt.Println("Package variable incremented in main:", x)
	foo()
}

func foo() {
	fmt.Println("Package variable accessed in foo:", x)
	x++
	fmt.Println("Package variable incremented in foo:", x)
}

/* Output:
Package variable: 0
Package variable incremented in main: 1
Package variable accessed in foo: 1
Package variable incremented in foo: 2
*/
```

### Function-Level Variable Scope

Variables declared within a function are only accessible within that function.

```go
package main

import "fmt"

func main() {
	foo()
}

func foo() {
	var y int
	fmt.Println("Function-level variable y:", y)
	y++
	fmt.Println("Function variable y incremented:", y)
}

/* Output:
Function-level variable y: 0
Function variable y incremented: 1
*/
```

### Block-Level Variable Scope

Variables declared within a block (enclosed in curly braces) are only accessible within that block.

```go
package main

import "fmt"

func main() {
	{
		z := 10
		fmt.Println("Block variable z:", z)
		z++
		fmt.Println("Block variable z incremented:", z)
	}
	// fmt.Println("Block variable z is not accessible outside of block:", z)
}

/* Output:
Block variable z: 10
Block variable z incremented: 11
*/
```

## True Closure Example

Here's an example of a true closure that captures variables from its enclosing scope:

```go
package main

import "fmt"

func main() {
	// Create a closure that captures the variable 'name'
	greeter := createGreeter("Alice")
	
	// Call the closure multiple times
	fmt.Println(greeter()) // Hello, Alice! Count: 1
	fmt.Println(greeter()) // Hello, Alice! Count: 2
	fmt.Println(greeter()) // Hello, Alice! Count: 3
}

// createGreeter returns a closure that captures 'name' and maintains 'count'
func createGreeter(name string) func() string {
	count := 0
	return func() string {
		count++
		return fmt.Sprintf("Hello, %s! Count: %d", name, count)
	}
}

/* Output:
Hello, Alice! Count: 1
Hello, Alice! Count: 2
Hello, Alice! Count: 3
*/
```

## Key Concepts:

1. **Variable Scope**: Determines where a variable can be accessed
2. **Package Scope**: Variables accessible throughout the entire package
3. **Function Scope**: Variables accessible only within the function where they're declared
4. **Block Scope**: Variables accessible only within the block where they're declared
5. **Closure**: A function that captures variables from its lexical scope
6. **Lexical Scoping**: Variables are accessible based on where they're defined in the code

## Scope Rules in Go:

- **Package-level variables**: Accessible from any function in the same package
- **Function-level variables**: Only accessible within the function where they're declared
- **Block-level variables**: Only accessible within the block where they're declared
- **Shadowing**: Inner scopes can declare variables with the same name as outer scopes

## Best Practices:

- Use the most restrictive scope possible for variables
- Avoid package-level variables when possible
- Use closures for maintaining state in functional programming patterns
- Be aware of variable shadowing to avoid confusion