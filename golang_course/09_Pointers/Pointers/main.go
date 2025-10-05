package main

import "fmt"

func main() {

	a := 42
	fmt.Println("The value of a", a)
	fmt.Println("The address of a", &a)        // & gives you the address
	fmt.Println("The value of using *& ", *&a) // give the value

	fmt.Printf("The Type of a variable %T\n", a)
	fmt.Printf("The Type of a variable memory %T\n", &a)

	// b is of type "int pointer" b points to the memory address where an int is stored
	b := &a
	fmt.Printf("The Type of b variable %T\n", b)
	fmt.Println("The address of b", &b)       // & gives you the address
	fmt.Println("The value of using *& ", *b) // give the value

	*b = 45
	fmt.Printf("The Type of b variable %T\n", b)
	fmt.Println("The address of b", &b)       // & gives you the address
	fmt.Println("The value of using *& ", *b) // give the value
}
