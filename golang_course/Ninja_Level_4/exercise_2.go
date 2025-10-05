package main

import "fmt"

/*
● Using a COMPOSITE LITERAL:
	● create a SLICE of TYPE int
	● assign 10 VALUES
● Range over the slice and print the values out.
	● Using format printing
○ print out the TYPE of the slice
*/

func main() {

	no := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range no {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", no)
}
