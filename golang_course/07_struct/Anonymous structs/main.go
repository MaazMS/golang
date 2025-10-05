package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		Name string
		age  int
	}{
		Name: "Maaz",
		age:  25,
	}
	fmt.Println(p1)
}
