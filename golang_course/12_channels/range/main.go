package main

// Range stops reading from a channel when the channel is closed.

import "fmt"

func main() {

	c := make(chan int)
	// send
	go func() {
		for i := 0; i <= 5; i++ {
			c <- 45
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

/*
because i am not close channel
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main.func1(0xc000094060)
        /home/maaz/github/LearningGO/golang_deep_dive/12_channels/range/add.go:11 +0x45
main.main()
        /home/maaz/github/LearningGO/golang_deep_dive/12_channels/range/add.go:1
*/
