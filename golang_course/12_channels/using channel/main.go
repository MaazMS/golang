package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	baar(c)

	fmt.Println("about to exit code")
}

// send channel
func foo(c chan<- int) {
	c <- 500
}

// receive channel
func baar(c <-chan int) {
	fmt.Println("receive channel", <-c)
}
