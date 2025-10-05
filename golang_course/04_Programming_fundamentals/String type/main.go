package main

import (
	"fmt"
)

func main() {

	// string value in double quote""
	s := "Hello World"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)

	// string value in back ticks ``
	multis := `Hello 
		world`
	fmt.Printf("%s\n", multis)
	fmt.Printf("%q\n", multis)

	// Iterator string using range
	for index , v := range s {
		fmt.Printf("The index number of %c is %d\n", v, index)
	}
	// Iterator string using length of string
	for i := 0 ;i < len(s) ; i++ {
		fmt.Printf("%c\n", s[i])
	}

	// A string value is a (possibly empty) sequence of bytes.
	// byte which is an alias for uint8
	bs := []byte(s)
	fmt.Println(s)
	fmt.Printf("%T\n", bs)
}
