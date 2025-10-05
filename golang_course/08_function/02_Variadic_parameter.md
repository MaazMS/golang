## Variadic parameter  
1. `...` Inside parameter is called Variadic parameter.  
1. `...` operator category in golang specification. 
1. zero and unlimited
```go
package main

import "fmt"

func main() {

	// pass multiple vau
	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)

	// passing in zero
	y := foo()
	fmt.Println(y)

	//
	name := "Maaz"
	no := []int {1, 2, 3, 4, 5, 6}
	a,b := bar(name, no...)
	fmt.Println(a, b)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

	no := 0
	for _, v := range x {
		no += v

	}
	return no
}

func foo(no ...int) []int {
	fmt.Println(no)
	fmt.Printf("%T\n", no)
	fmt.Println(len(no))
	return no
}

func bar(nm string, no ...int) (string, []int)  {

	fmt.Println(nm)
	for _, k := range no {
		fmt.Println(k)
	}
	return nm, no
}

``` 
