package main

import "fmt"

func main() {

	no1 := 10
	no2 := &no1

	fmt.Println("no1 := 10 \t", no1)
	fmt.Println("no2 := &no1 \t", *no2)

	*no2 = 20
	fmt.Println(" assign no1 value using * no2 = 20  \t", no1)

}
