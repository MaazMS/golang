## The var keyword  
1.  where var can be used  
    1. any place within the package   
    
```go
package main

import "fmt"

// DECLARE the variable "No"
// ASSIGN the value 10
// declare & assign = initialization
var No = 10

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {

	fmt.Println(No)
	foo()
}

func foo() {
	fmt.Println("again ", No)

}
```
    