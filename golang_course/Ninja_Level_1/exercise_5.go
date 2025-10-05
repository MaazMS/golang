package main

/*
1. at the package level scope, using the “var” keyword, create a VARIABLE with the
IDENTIFIER “y4”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x4”
2. in func main
a. this should already be done
	i. print out the value of the variable “x4”
	ii. print out the type of the variable “x4”
	iii. assign your own VALUE to the VARIABLE “x4” using the “=” OPERATOR
	iv. print out the value of the variable “x4”
b. now do this
i. now use CONVERSION to convert the TYPE of the VALUE stored in “x4” to the UNDERLYING TYPE
	1. then use the “=” operator to ASSIGN that value to “y4”
	2. print out the value stored in “y4”
	3. print out the type of “y4”
*/

import (
	"fmt"
)

type number int

var x4 number
var y4 int

func main() {

	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 10
	fmt.Println(x4)
	//  CONVERSION to convert the TYPE of the VALUE stored in “x4” to the UNDERLYING TYPE int
	y4 = int(x4)
	y4 = 20
	fmt.Println(y4)
	fmt.Printf("%T\n", y4)

}
