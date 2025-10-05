package main

import "fmt"

/*
● create a func with the identifier work that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
● create a func with the identifier easy that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in
*/

func main() {

	noi := []int{1, 2, 3, 4, 5}
	a := work(noi...)
	fmt.Println("sum of variable noi ", a)
	noj := []int{11, 12, 13, 14, 14}
	b := easy(noj)
	fmt.Println("sum of variable noj", b)

}

func work(no ...int) int {
	sum := 0
	for _, s := range no {
		sum += s
	}
	return sum
}

func easy(no []int) int {
	total := 0
	for _, s := range no {
		total += s
	}
	return total
}
