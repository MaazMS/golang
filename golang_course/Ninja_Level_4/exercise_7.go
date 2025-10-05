package main

import (
	"fmt"
)

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
	"James", "Bond", "Shaken, not stirred"
	"Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record.
*/

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Printf("slcie string x\t%s\n", x)
	fmt.Printf("slcie string y\t%s\n", y)

	z := [][]string{x, y}
	fmt.Printf("slcie string z\t%s\n", z)

	for _, v := range z {
		for j, vls := range v {
			fmt.Println(j, vls)
		}
	}

}
