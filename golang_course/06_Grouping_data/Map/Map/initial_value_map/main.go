package main

import "fmt"

func main()  {
	var dict map[string]string
	fmt.Printf("Zero Value: %#v\n", dict)
	fmt.Printf("Zero Value: %v\n", dict)
	fmt.Printf("Zero Value: %T\n", dict)
	fmt.Printf("# of Keys : %d\n", len(dict))
}
