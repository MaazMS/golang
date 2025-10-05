package main

import "fmt"

func main() {

	c := make(chan int)    // bidirectional
	cr := make(<-chan int) // receive
	cd := make(chan<- int) // send

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cd)
}
