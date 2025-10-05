package main

import "fmt"

func main() {

	var no1 int = 300
	var no2 int = 400

	if no1%2 == 0 {
		fmt.Println(" no1 is even ")

		if no1 > 100 {
			fmt.Println("no1 is greater then 100")
		}
	}

	if no2 == 100 {
		fmt.Println("no2 is equal to 100")
	} else if no2 == 200 {
		fmt.Println(" no2 is equal to 200")
	} else if no2 == 300 {
		fmt.Println("no2 is equal to 300 ")
	} else {
		fmt.Println(" no2 is greater then 300 ")
	}

}
