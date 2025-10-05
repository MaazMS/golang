package main

/*
Print every rune code point of the uppercase alphabet three times.
*/
import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c\n", i)
		}
	}
}
