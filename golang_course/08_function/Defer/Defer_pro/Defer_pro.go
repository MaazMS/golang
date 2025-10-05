package main

import "fmt"

func add(a, b int) int {

	addition := a + b
	fmt.Println(addition)
	return 0
}
func main() {

	add(10, 10)

	defer add(1, 1)
	defer fmt.Println("defer statement")

	fmt.Println("Hello")
}
