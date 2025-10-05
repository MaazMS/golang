## Iota  
1. Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.    
1. It is reset to 0 whenever the reserved word const appears in the source.   
1. It can be used to construct a set of related constants.  

```go
package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)
const (
	l = iota
	m
	n
)

func main()  {

	fmt.Println(a)  // a = 0
	fmt.Println(b) // b = 1
	fmt.Println(c) // c = 2
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)

	fmt.Println(l)  // l = 0
	fmt.Println(m) // m = 1
	fmt.Println(n) // n = 2
	fmt.Printf("%T\n",l)
	fmt.Printf("%T\n",m)
	fmt.Printf("%T\n",n)



}
```