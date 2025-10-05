package main

import "fmt"

/*
1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
IDENTIFIERS “x” and “y” and “z”
a. 42
b. `James Bond`
c. true
Now print the values stored in those variables using
a. a single print statement
b. multiple print statements
*/

func main() {

	//  short declaration operator, ASSIGN these VALUES to VARIABLES IDENTIFIERS “x” and “y” and “z”
	x := 42
	y := `James Bond`
	z := true

	// a single print statement
	fmt.Println(x, y, z)

	// multiple print statements
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
