package main

import "fmt"

func init() {
	fmt.Println(" init() function is executed before the main() function call")
}
func main() {

	fmt.Println("Hello world")

	foo()

}

func foo() {

	fmt.Println("Even Number")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	baar()
}

func baar() {

	fmt.Println("Fuc baar, Program is exect")
}

/*
Flow of control
	1. Sequential flow (top to bottom )
	1. loop iterator
	1. conditional

*/
