package main

/*
create a func with the identifier foo that returns an int
	○ create a func with the identifier bar that returns an int and a string
	○ call both funcs
	○ print out their results
*/
import "fmt"

func main() {

	a := foo()
	fmt.Println("func foo  return int type", a)
	x, y := bar()
	fmt.Println("func foo  return int type and string", x, y)
}

func bar() (int, string) {
	return 25, "Hello"
}

func foo() int {
	return 100
}
