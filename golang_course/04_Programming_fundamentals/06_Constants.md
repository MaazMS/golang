## Constants   
1. There are boolean constants, rune constants, integer constants, floating-point constants, complex constants,      
   and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants.   

## There are two type is constant.   
1. declare and assign value to constant use assigment operator.  
1. typed constant   `const int a =40`  
1. untyped constant `constant a = 40`  

```go
package main

import (
	"fmt"
)

const (
	no = 40
	name = "Maaz"
	fb =10.0
)
func main()  {

	fmt.Print(no)
	fmt.Printf("\t%T\n", no)

	fmt.Print(name)
	fmt.Printf("\t%T\n", name)

	fmt.Print(fb)
	fmt.Printf("\t%T\n", fb)
}

```