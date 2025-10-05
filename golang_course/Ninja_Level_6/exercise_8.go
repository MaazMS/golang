package main

import "fmt"

/*
● Create a func which returns a func
● assign the returned func to a variable
● call the returned func
*/

func main() {

	vf := Get()
	fmt.Println("call return function ", vf())
}

func Get() func() int {
	return func() int {
		return 1997
	}
}
