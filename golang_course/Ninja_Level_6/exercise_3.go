package main

import "fmt"

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits

func main() {

	defer Cat()
	fmt.Println("func main")
	Sat()
}

func Sat() {
	fmt.Println("Func sat")
}

func Cat() {
	fmt.Println("Func Cat")

}
