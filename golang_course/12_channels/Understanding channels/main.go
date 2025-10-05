package main

import "fmt"

func main() {
	//  this code is not going to run.

	// put t2 on c
	/*c := make(chan  int)

	c <- 42
	// off of c
	fmt.Println(<-c) */

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// buffer with channel

	d := make(chan int, 1)

	d <- 50
	fmt.Println(<-d)
}
