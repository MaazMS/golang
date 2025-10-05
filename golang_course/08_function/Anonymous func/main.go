package main

import "fmt"

func main() {
	foo()

	// Anonymous func  with no parameter
	func() {
		fmt.Println("Anonymous func  with no parameter")
	}()

	// Anonymous func  with parameter
	func(x int) {
		fmt.Println("Anonymous func  with parameter \t", x)
	}(45)

}

func foo() {
	fmt.Println("func foo")
}
