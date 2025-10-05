package main

import "fmt"

// Build and use an anonymous func

func main() {

	func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)

		}
	}()
}
