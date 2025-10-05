package main

import (
	"fmt"
)

func main()  {

	// zero value of bool is false
	var condition bool
	fmt.Println(condition)

	// assign the value of bool
	condition = true
	fmt.Println(condition)

	a := 10
	b := 20
	// a is not equal to b their for it is return false
	fmt.Println(a == b)
}