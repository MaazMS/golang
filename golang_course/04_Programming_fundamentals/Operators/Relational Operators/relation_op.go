package main

import "fmt"

func main() {

	no1 := 10
	no2 := 20

	res1 := no1 == no2
	fmt.Println("no1 and no2 equal to ", res1)

	res2 := no1 != no2
	fmt.Println("no1 and no2 Not equal to ", res2)

	res3 := no1 > no2
	fmt.Println("no1 greater then no2", res3)

	res4 := no1 < no2
	fmt.Println("no1 less then no2 ", res4)

	res5 := no1 >= no2
	fmt.Println("no1 greater then equal no2 ", res5)

	res6 := no1 <= no2
	fmt.Println("no1  less then equal no2 ", res6)

}
