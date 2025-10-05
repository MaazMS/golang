package main

import (
	"fmt"
)

// DECLARE a variable to be of a certain TYPE
var name string
var number int
var fnumber float64
var condition bool

func main() {

	fmt.Print("var name string is :\t", name)
	fmt.Printf("\t%T\t\n", name)

	fmt.Print("var number int :\t", number)
	fmt.Printf("\t%T\t\n", number)

	fmt.Print("var fnumber float64  :\t", fnumber)
	fmt.Printf("\t%T\t\n", fnumber)

	fmt.Print("var condition bool :\t", condition)
	fmt.Printf("\t%T\t\n", condition)

}
