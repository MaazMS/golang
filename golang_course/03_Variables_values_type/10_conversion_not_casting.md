## Conversion of type
1. golang is static type language,
1. Their for it is type is more important in golang.
1. Change type one to another type we must be fist check both variable type.  
1. left side variable TYPE and right side variable TYPE.   
1. For conversion of TYPE right side variable is use left side variable TYPE.      
1. var no int  `// it is left side variable.`   
1. type palindrome  int `// CREATE our own TYPE palindrome and underline type is int`
1. var b palindrome `// create variable b with TYPE palindrome and it is right side variable`
1. Example `no = int(b)`
```go
package main

import (
	"fmt"
)

// DECLARE variable no for TYPE int
var no int

// CREATE our own TYPE palindrome and underline type is int
type palindrome int

// create variable b with TYPE palindrome
var b palindrome

func main() {

	no = 100
	fmt.Println(no)
	fmt.Printf("%T\n", no)

	b = 500
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// conversion of  no = b
	// cannot use b (type palindrome) as type int in assignment
	// int(b) because DECLARE variable no for TYPE int
	no = int(b)
	fmt.Println(no)
	fmt.Printf("%T\n", no)
}

```