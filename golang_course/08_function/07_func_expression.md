## func expression
1. function, and we're going to assign it to a variable.    
1. It  known as a funk expression.  

```go
package main

import "fmt"

func main() {

	// This is not function assign it to a variable
	// x :=foo()

	// This is function assign it to a variable

	f := func() {
		fmt.Println("function f")
	}
	f()

	// This is function assign it to a variable
	x := func(no int) {
		fmt.Println("function x", no)
	}
	x(1997)
}

```
