package main

/*
1. Use var to DECLARE three VARIABLES. The variables should have package level
scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the
variables and make sure the variables are of the following TYPE (meaning they can
store VALUES of that TYPE).
a. identifier “x1” type int
b. identifier “y1” type string
c. identifier “z1” type bool
2. in func main
a. print out the values for each identifier
b. The compiler assigned values to the variables. What are these values called?
*/

import (
	"fmt"
)

var x1 int
var y1 string
var z1 bool

func main() {

	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)
	fmt.Println("This value is called zero vaule")

}
