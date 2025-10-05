package main

import "fmt"

func main() {
	intno1 := 10
	intno2 := 20

	floatno1 := 10.0
	floatno2 := 20.0

	res1 := intno1 != intno2 && floatno1 != floatno2
	fmt.Println(res1)
	res2 := intno1 == intno2 || floatno1 == floatno2
	fmt.Println(res2)
	res3 := !(intno1 != intno2)
	fmt.Println(res3)
}
