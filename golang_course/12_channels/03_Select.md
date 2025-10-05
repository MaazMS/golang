## Select using channel  
1. Select statements pull the value from whatever channel has a value ready to be pulled.   
1. create three variable with type channel.  
1. go routine send function is used to `send channel` , see the send function parameter.  
1. Then close the channel inside send function.  
1. receive function is used to `receive channel` , see the receive function parameter.
1. Inside `for loop` select keyword is used to check the type of the case, case type is even , odd and quite.  

````go
package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true
}

// receive channel
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case <-quit:
			return
		}
	}
}

````