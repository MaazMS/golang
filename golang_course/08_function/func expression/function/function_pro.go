package main

import "fmt"

func calculateArea(l, b int) int {

	are := l * b

	return are
}

func main() {

	fmt.Println(calculateArea(10, 10))
}
