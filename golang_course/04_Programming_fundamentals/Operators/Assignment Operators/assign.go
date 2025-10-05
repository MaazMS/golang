package main

import "fmt"

func main() {

	var no1 = 10
	var no2 = 2

	no1 = no2
	fmt.Println("Simple Assignment no1 = no2 \t", no1)

	no1 += no2
	fmt.Println("Simple Assignment no1 += no2 \t", no1)

	no1 -= no2
	fmt.Println("Simple Assignment no1 -= no2 \t", no1)

	no1 *= no2
	fmt.Println("Simple Assignment no1 *= no2 \t", no1)

	no1 /= no2
	fmt.Println("Simple Assignment no1 /= no2 \t", no1)

	no1 %= no2
	fmt.Println("Simple Assignment no1 %= no2 \t", no1)
}
