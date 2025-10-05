package main

import (
	"fmt"
)

// DECLARE variable no for TYPE int
var no int

// CREATE our own TYPE palindrome and underline type is int
type palindrome int

// create variable b with TYPE palindrome
var b palindrome

func main() {

	no = 100
	fmt.Println(no)
	fmt.Printf("%T\n", no)

	b = 500
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// conversion of  no = b
	// cannot use b (type palindrome) as type int in assignment
	// int(b) because DECLARE variable no for TYPE int
	no = int(b)
	fmt.Println(no)
	fmt.Printf("%T\n", no)
}
