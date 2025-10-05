package main

import "fmt"

/*
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

const (
	number = 10
	name string = "Maaz"
)

func main()  {

	fmt.Println("Untyped constant", number)
	fmt.Println("yped constant", name)

}