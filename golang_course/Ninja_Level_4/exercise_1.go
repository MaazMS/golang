package main

import "fmt"

/*
● Using a COMPOSITE LITERAL:
	● create an ARRAY which holds 5 VALUES of TYPE int
	● assign VALUES to each index position.
● Range over the array and print the values out.
● Using format printing
	○ print out the TYPE of the array

*/

func main() {

	var a [5]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	for i, v := range a {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", a)
}
