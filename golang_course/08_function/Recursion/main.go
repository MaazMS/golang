package main

import "fmt"

func main() {

	fmt.Println(4 * 3 * 2 * 1)
	r1 := recursionFactorial(4)
	fmt.Println("recursion Factorial", r1)
	r2 := loopFactorial(4)
	fmt.Println(r1)
	fmt.Println("loop Factorial", r2)

}

func recursionFactorial(no int) int {

	if no == 0 {
		return 1
	}
	return no * recursionFactorial(no-1)
	// behind the seen
	// return no * recursionFactorial(4 -1) * recursionFactorial(3 -1) * recursionFactorial(2 -1) * 1
}

func loopFactorial(no int) int {

	total := 1

	for ; no > 0; no-- {
		total *= no
	}
	return total
}
