package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x)

	y := baar()
	fmt.Printf("Type of y is \t%T\n", y)
	fmt.Printf("function call")
	y()

	a := work()
	i := a()
	fmt.Println(i)

	fmt.Println(smart()())
}

func foo() string {
	return "Hello world"
}

func baar() func() int {
	return func() int {
		return 1997
	}
}

func work() func() string {
	return func() string {
		return "Hello golang"
	}
}

func smart() func() bool {
	return func() bool {
		return true
	}
}
