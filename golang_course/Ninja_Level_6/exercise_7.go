package main

import "fmt"

// ● Assign a func to a variable, then call that func
// declare  var variable func()
var g func()

func main() {

	f := func() {
		for i := 0; i <= 5; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Println("==================================")
	g = f
	g()

}
