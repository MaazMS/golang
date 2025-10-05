## Recursion  
1. I will say that anything you do with a recursion, you could also do with a loop.  
1. it's also going to have negative impacts on memory usage in your program.  
1. function calls itself is called Recursion.  

```go
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
/* output
24
recursion Factorial 24
24
loop Factorial 24

*/
```