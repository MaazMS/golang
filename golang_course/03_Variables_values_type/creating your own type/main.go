package main

import (
	"fmt"
)

// DECLARE variable no for TYPE int
var a int

// create our own type palindrome and underline type is int
type palindrome int

// create variable b with TYPE palindrome
var b palindrome

func main() {
	a = 15
	// assign value to variable b with TYPE palindrome is int 25
	b = 25
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
