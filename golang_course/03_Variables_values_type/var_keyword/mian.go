package main

import "fmt"

// DECLARE the variable "No"
// ASSIGN the value 10
// declare & assign = initilization
var No = 10

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {

	fmt.Println(No)
	foo()
}

func foo() {
	fmt.Println("again ", No)

}
