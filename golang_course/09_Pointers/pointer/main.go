package main

import "fmt"

func main() {

	a := 200
	fmt.Println("a before ", &a)
	fmt.Println("a before ", a)
	bar(&a)
	fmt.Println("a after  ", &a)
	fmt.Println("a after  ", a)

}

func bar(b *int) {
	fmt.Println("b before ", *b)
	fmt.Println("b before ", b)
	*b = 500 // assign the of b datatype is *int
	fmt.Println("b after  ", *b)
	fmt.Println("b after  ", b)

}

/* output
a before  0xc0000200a8
a before  200
b before  200
b before  0xc0000200a8
b after   500
b after   0xc0000200a8
a after   0xc0000200a8
a after   500
*/
