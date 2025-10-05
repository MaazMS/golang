package main

import "fmt"

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:
*/

func main() {

	f := enclosed()
	fmt.Println("function return f variable call", f())
	g := enclosed()
	fmt.Println("function return g variable call", g())
	fmt.Println("function return g variable call", g())
	fmt.Println("function return g variable call", g())
	fmt.Println("function return g variable call", g())
	fmt.Println("function return g variable call", g())
	fmt.Println("function return g variable call", g())

}
func enclosed() func() int {
	x := 0
	return func() int {
		x++
		return x
	}

}
