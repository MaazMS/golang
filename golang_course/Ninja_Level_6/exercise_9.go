package main

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
*/
import "fmt"

func main() {

	run(hello)
	run(world)
}

func hello() {
	fmt.Println("function hello")
}
func world() {
	fmt.Println("function world")
}

func run(f func()) {
	fmt.Println("function run")
	f()
}
