package main

import "fmt"

func show(s1 string) {

	for i := 0; i <= 5; i++ {
		fmt.Println(s1)
	}
}

func main() {

	go show("hello")

	show("Maaz")
}
