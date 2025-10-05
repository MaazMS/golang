package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{8, 2, 3, 9, 5, 1, 8}

	fmt.Println("before sort integer")

	fmt.Println(xi)
	fmt.Println("after sort integer")
	sort.Ints(xi)
	fmt.Println(xi)

	zi := []string{"user", "work", "apple", "z", "rate"}
	fmt.Println("before sort string")
	fmt.Println(zi)
	fmt.Println("after sort string")
	sort.Strings(zi)
	fmt.Println(zi)

}
