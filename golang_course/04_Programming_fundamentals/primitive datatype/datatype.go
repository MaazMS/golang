package main

import "fmt"

func inttype() {

	var no int = -10
	var no1 uint = 100

	fmt.Printf("%v %T\t\n", no, no)
	fmt.Printf("%v %T\t\n", no1, no1)
}

func floattype() {

	var no float32 = 10.6556
	var no1 float64 = 1097.655698798798

	fmt.Printf("%v %T\t\n", no, no)
	fmt.Printf("%v %T\t\n", no1, no1)
}

func complextype() {

	var no complex64 = complex(16, 6)
	var no1 complex128 = complex(19, 6)

	fmt.Printf("%v %T\t\n", no, no)
	fmt.Printf("%v %T\t\n", no1, no1)
}

func booltype() {

	var flag bool
	fmt.Printf("%v %T default value \t\n", flag, flag)

	var value bool = true
	fmt.Printf("%v %T true is assign to value \t\n", value, value)

	var code bool = false
	fmt.Printf("%v %T false is assign to  code \t\n", code, code)
}

func stringtype() {

	var s1 string = "Maaz"
	fmt.Printf("%v %T\t\n", s1, s1)

}

func main() {

	inttype()
	floattype()
	complextype()
	booltype()
	stringtype()

}
