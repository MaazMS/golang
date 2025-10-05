## Bool type  
1. A bool is a binary TYPE having two possible values of either “true” and “false.”   
1. When you use an equality comparison operator, this is an expression which will
   evaluate to a boolean value.  
   
```bash
a. ==
b. <=
c. >=
d. !=
e. <
f. >
``` 

```go
package main

import (
	"fmt"
)

func main()  {

	// zero value of bool is false
	var condition bool
	fmt.Println(condition)

	// assign the value of bool
	condition = true
	fmt.Println(condition)

	a := 10
	b := 20 
	// a is not equal to b their for it is return false
	fmt.Println(a == b)
}
```