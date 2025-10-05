package main

/*
Write a program that prints a number in decimal, binary, and hex

*/

import (
	"fmt"
)

func main()  {
	a := 15

	fmt.Printf("decimal number %d\n",a)
	fmt.Printf("binary number %b\n",a)
	fmt.Printf("hex number %#X\n",a)


}