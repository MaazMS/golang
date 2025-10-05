package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 10; i++ {

		for j := 0; j < 3; j++ {

			fmt.Printf("outer  for loop %d\n\t\t\t and inner for loop %d\n", i, j)
		}
		fmt.Println("==================================================")
	}
}
