## Returning Functions

In Go, you can return a function from another function. This is a powerful feature that enables higher-order functions and functional programming patterns.

### Basic Example: Returning a Function

```go
package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
}

func foo() string {
	return "Hello world"
}

/* Output:
Hello world
*/
```

### Example 1: Returning a Function that Returns an Integer

```go
package main

import "fmt"

func main() {
	y := bar()
	fmt.Printf("Type of y is \t%T\n", y)
	fmt.Println("Function call result:")
	result := y()
	fmt.Println(result)
}

func bar() func() int {
	return func() int {
		return 1997
	}
}

/* Output:
Type of y is    func() int
Function call result:
1997
*/
```

### Example 2: Returning a Function that Returns a String

```go
package main

import "fmt"

func main() {
	a := work()
	i := a()
	fmt.Println(i)
}

func work() func() string {
	return func() string {
		return "Hello golang"
	}
}

/* Output:
Hello golang
*/
```

### Example 3: Direct Function Call Chain

```go
package main

import "fmt"

func main() {
	fmt.Println(smart()())
}

func smart() func() bool {
	return func() bool {
		return true
	}
}

/* Output:
true
*/
```

### Key Points:

1. **Function Types**: When returning a function, you specify the function signature as the return type (e.g., `func() int`, `func() string`)
2. **Anonymous Functions**: The returned functions are typically anonymous functions (closures)
3. **Function Calls**: You can call the returned function immediately or store it in a variable for later use
4. **Chaining**: You can chain function calls directly (e.g., `smart()()`)
5. **Type Safety**: Go's type system ensures that the returned function matches the declared signature