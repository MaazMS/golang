package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Declare variable with certain TYPE int8 , range is -128 to 127
	var a int8
	//var b int8
	var c int8
	//var d int8

	a = 127
	// constant 128 overflows int8
	//b = 128

	c = -128
	// constant -129 overflows int8
	//d = -129

	fmt.Println(" variable a ", a)
	//fmt.Println(" variable b ", b)
	fmt.Println(" variable c ", c)
	//fmt.Println(" variable d ", d)

	// architecture of golang
	// linux and amd64 word size of computer is 64
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	var  a1 uint
	a1 = 255
	fmt.Println(a1)

	var  b2 int32
	b2 = 255
	fmt.Println(b2)
}
