package main

import "fmt"

func main() {

	n, err := fmt.Println(" fmt package ", 123, true)

	fmt.Println(n)
	fmt.Println(err)
	no, _ := fmt.Println("avoid fmt error", 123, true)
	fmt.Println(no)

}
