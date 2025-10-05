## Array 
1. An array is a numbered sequence of elements of a single type, called the element type.  
1. The number of elements is called the length of the array.  
1. It is never negative.  
1. The length is part of the array's type.  
1. Array types are always one-dimensional but may be composed to form multi-dimensional types.  
*. Array is allow you to group together value of same TYPE and fixed size.   

```go
package main

import (
	"fmt"
)

func main()  {

	fmt.Println("Array example")
	var no [5] int
	fmt.Println(no)
	no[3] = 100
	fmt.Println(no)
	fmt.Println(len(no))
}

```

```bash
Array example
[0 0 0 0 0]
[0 0 0 100 0]
5
``` 

## Iterate an array   
1. Iterate an array using range 

```go
package main

import (
	"fmt"
) 
func main()  {

	var a [5] int
	a [0] = 0
	a [1] = 1
	a [2] = 2
	a [3] = 3
	a [4] = 4

	for i, v := range a{
		fmt.Println(i,  v)

	}
}
```
```bash
0 0
1 1
2 2
3 3
4 4
```