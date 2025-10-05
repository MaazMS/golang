package main

import "fmt"

func main() {
	a := 45
	fmt.Println("a before ", &a)
	fmt.Println("a before ", a)
	foo(a)
	fmt.Println("a after  ", &a)
	fmt.Println("a after  ", a)
}

func foo(b int) {
	fmt.Println("b before ", &b)
	fmt.Println("b before ", b)
	b = 50
	fmt.Println("b after  ", &b)
	fmt.Println("b after  ", b)

}

/* output
a before  0xc0000200a8
a before  45
b before  0xc0000200d0
b before  45
b after   0xc0000200d0
b after   50
a after   0xc0000200a8
a after   45

*/
