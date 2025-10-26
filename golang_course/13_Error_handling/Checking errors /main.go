package main

import (
	"fmt"
)

func main() {

	var Name string
	fmt.Print("Your Name\t")
	_, err := fmt.Scan(&Name)
	if err != nil {
		panic(err)
	}
	fmt.Println(Name)
	n, err := fmt.Println("Hello, playground")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The length of return n int size of string ", n)
	fmt.Println("program is end")

}
