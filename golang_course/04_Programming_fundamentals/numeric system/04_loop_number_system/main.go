package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loop of number system for decimal, binary, hexadecimal")
	for i := 1; i <= 10; i++ {
		fmt.Printf("decimal Number %d\t binary number %b\t hexadecimal number %#X\t", i, i, i)
		fmt.Println()
	}
}
