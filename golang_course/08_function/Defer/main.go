package main

import "fmt"

func main() {

	fmt.Println("func main")
	defer foo()
	bar()
}

func foo() {
	fmt.Println("func foo")
}
func bar() {
	fmt.Println("func bar")

}
