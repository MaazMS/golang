package main

import "fmt"

func main() {

	var s1 [2]string
	s1[0] = "zero"
	s1[1] = "end"

	fmt.Println(s1[0])
	fmt.Println(s1[1])

	fullname := [3]string{"Md", "Maaz", "Shaikh"}
	for i := 0; i < 3; i++ {
		fmt.Println(fullname[i])
	}

}
