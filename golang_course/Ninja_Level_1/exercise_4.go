package main

/*
1. Create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR”
keyword
3. in func main
a. print out the value of the variable “x2”
b. print out the type of the variable “x2”
c. assign 42 to the VARIABLE “x2” using the “=” OPERATOR
d. print out the value of the variable “x2”
*/

import (
	"fmt"
)

//Declare variable a with TYPE int
var a int

// Declare variable user with underlying type be an int
type user int

var x2 user

func main()  {

	fmt.Println(x2)
	fmt.Printf("%T\n", x2)
	x2 = 42
	fmt.Println(x2)
}
