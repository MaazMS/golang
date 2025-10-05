package main

import "fmt"

func main() {

	i := 10

	j := &i
	fmt.Println("retun memeory Address of i variable \t", *j)
	fmt.Println("value of variable i                 \t", i)

	var k *int
	k = &i
	fmt.Println("store memory address of i variable \t", k)
	*k = 20
	fmt.Println("Access variable and change value\t", *k)

}
