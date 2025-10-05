package main

import "fmt"

func main() {

	speed := 100 // speed is int
	force := 2.5 // force is float64

	fmt.Println(speed, force)
	clculate := speed *  int(force)
	fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force, clculate)

	//clculate := float64(speed) *  force // Cannot use 'float64(speed) * force' (type float64) as type int
	//fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force, clculate)
	fmt.Println()

	clculate = int (float64(speed) *  force)
	fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force, clculate)

}