package main

import "fmt"

type userinfo interface {
	calculateArea() int
}

func calculateArea(l, b int) int {

	are := l * b
	return are
}
func main() {

	fmt.Println(calculateArea(10, 10))
}
