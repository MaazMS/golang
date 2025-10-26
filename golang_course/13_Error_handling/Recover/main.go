package main

import (
	"fmt"
)

func main() {

	f()
	fmt.Println("Progam exit")

}

func f() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered f in", r)
		}
	}()
	fmt.Println("calling g")
	g(0)
	fmt.Println("Program g return normally ")

}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking ")
		panic(fmt.Sprintf("%v", i))

	}
	defer fmt.Println("Defer in g ", i)
	fmt.Println("Printing i in g ", i)
	g(i + 1)
}

/* output
calling g
Printing i in g  0
Printing i in g  1
Printing i in g  2
Printing i in g  3
panicking
Defer in g  3
Defer in g  2
Defer in g  1
Defer in g  0
*/
