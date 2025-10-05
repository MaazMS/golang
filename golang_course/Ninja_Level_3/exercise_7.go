package main

import "fmt"

// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

func main() {

	no := 10
	if no != 10 {
		fmt.Println("number is not equal to 10")
	} else if no == 10 {
		fmt.Println("number is  equal to 10")
	}
}
