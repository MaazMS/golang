// This allows Go to create an executable program
package main

import "fmt"

// The main function is the entry point of the program.
// The Go runtime will execute this function first when the program starts.

func init() {
	fmt.Println(" init() function is executed before the main() function call")
}

func main() {
	fmt.Println("Hello golang")
}
