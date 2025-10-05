package main

import (
	"fmt"
)

func main() {

	no := 10

	if no == 10 {
		fmt.Println("no is 10")
	}

	x := 20
	if x == 10 {
		fmt.Println("x is 10")
	} else if x == 20 {
		fmt.Println("x is 20")
	}

	y := 20
	if y == 1 {
		fmt.Println("y is 1")
	} else if y == 2 {
		fmt.Println("y is 2")
	} else {
		fmt.Println("y is 20")
	}

}
