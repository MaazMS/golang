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
