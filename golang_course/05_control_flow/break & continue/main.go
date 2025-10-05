package main

import "fmt"

func main() {

	no := 1
	for {

		no++
		if no > 50 {
			break
		}
		if no%2 != 0 {
			fmt.Println(no)
			continue
		}

	}
}
