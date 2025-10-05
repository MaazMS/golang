package main

import "fmt"

func main() {

	fmt.Println("for loop program")
	for i := 1; i <= 15; i++ {

		fmt.Println(i)

		if i == 8 {

			// skip two iterations
			i = i + 2
			fmt.Println("skip two iterations \t", i)
			continue
		}

		if i == 13 {
			break
		}

	}
	goto lable1

lable1:
	fmt.Println("This is label identifier")

}
