package main

import (
	"fmt"
)

func main() {

	// basic func
	foo()
	// function pass an argument
	user("Maaz")
	// func return
	r := amount(300)
	fmt.Println(r)
	// func multiple return
	x, y := website()
	fmt.Println(x, y)

}

func foo() {
	fmt.Println("function foo")
}

// function take parameter
func user(us string) {
	fmt.Println("User ", us)
}

// func take parameter and return type is int
func amount(no int) int {
	New_balance := no + 50
	return New_balance
}

// func take parameter and multiple return type string, string
func website() (string, string) {

	ecommerce := "Amazon"
	searchEngine := "Google"

	return ecommerce, searchEngine
}
