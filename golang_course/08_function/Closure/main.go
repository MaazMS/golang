package main

import "fmt"

var x int

func main() {

	fmt.Println("package variable x ", x)
	x++
	fmt.Println("package variable x increment in func main", x)
	foo()
	{
		z := 10
		fmt.Println("block variable z ", z)
		z++
		fmt.Println("block variable z increment in block ", z)
	}
	// fmt.Println("block variable z is not use outside of block ",z)

}
func foo() {
	var y int
	fmt.Println("package variable x without increment in func foo", x)
	x++
	fmt.Println("package variable x increment in func foo", x)
	fmt.Println("function level variable y", y)
	y++
	fmt.Println("function variable y increment in func foo", y)

}
