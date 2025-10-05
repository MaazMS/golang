package main

import "fmt"

func main() {

	fmt.Println(" 1 + 10 = ", yoursum(1, 10))
	fmt.Println(" 2 + 10 = ", yoursum(2, 10))
	fmt.Println(" 3 + 10 = ", yoursum(3, 10))
	fmt.Println(" 4 + 10 = ", yoursum(4, 10))

}

func yoursum(no ...int) int {
	sum := 0
	for _, vls := range no {
		sum += vls
	}
	return sum
}
