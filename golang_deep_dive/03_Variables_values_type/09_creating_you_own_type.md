## Creating  your own type  
1. golang is static type language,    
1. Their for it is type is more important in golang.   
1. Create our own type with the help of `type` keyword.  
1. you mast be defined underline type of your type.   

```go
package main

import (
	"fmt"
)

var a int

// create our own type palindrome and underline type is int
type palindrome int

// create variable b with TYPE palindrome
var b palindrome

func main() {
	a = 15
	// assign value to variable b with TYPE palindrome is int 25
	b = 25
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}

```