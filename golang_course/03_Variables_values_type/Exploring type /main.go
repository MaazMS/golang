package main

import (
	"fmt"
)

// DECLARE the VARIABLE with the IDENTIFIER "name"
// is of TYPE string
// and I ASSIGN the VALUE "Maaz shaikh"

var name = "Maaz shaikh"

var user = `user 
name is
Maaz Shaikh `

func main() {

	fmt.Print(name)
	fmt.Printf("%T\n", name)

	fmt.Print(user)
	fmt.Printf("%T\n", user)

	// cannot use 10 (type untyped int) as type string in assignment
	//name =  10
	//fmt.Print(user)
	//fmt.Printf("%T\n",user)

}
