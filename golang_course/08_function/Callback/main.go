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
