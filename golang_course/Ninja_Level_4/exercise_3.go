package main

import "fmt"

/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]

*/
func main() {
	//  0,   1,  2,   3,  4,  5,  6,  7,  8,  9 index
	no := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(no[0:5])
	fmt.Println(no[5:10])
	fmt.Println(no[2:7])
	fmt.Println(no[1:6])

}
