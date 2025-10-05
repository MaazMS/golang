## Callback  
1. callback, which is passing a function in as an argument to another function.   
1. functional programming not something that is recommended in go, however, it is good to be aware of callbacks.  

````go
package main

import "fmt"

func main() {

	no := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// function call with variadic parameters no
	x := sum(no...)

	fmt.Println("Sum of number\t", x)

	// function "call with variadic parameters no1 and return int func sum " and pass variadic parameters
	y := even(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println("even  number\t", y)

}

// function call with variadic argument no1 and return int
func sum(no1 ...int) int {
	n := 0
	for _, v := range no1 {
		n += v
	}
	return n
}

// function even "with variadic parameters and return int func sum " is defined Anonymous func
//y1 ...int  pass variadic argument
// return type is int

func even(f func(x ...int) int, y1 ...int) int {

	var eveni []int
	for _, vls := range y1 {
		if vls%2 == 0 {
			eveni = append(eveni, vls)
		}
	}
	// call anonymous function call with variadic parameter
	total := f(eveni...)
	return total
}

/* output
Sum of number    45
even  number     20

*/

````