package main

import (
	"fmt"
	"sort"
)

func main() {

	s1 := []string{"A", "C", "B"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
