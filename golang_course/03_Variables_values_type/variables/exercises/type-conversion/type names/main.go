package main

import "fmt"

func main() {

	var speed int // speed is int
	var force int32 // force is int32

	// speed = force  // You can not assign different type to different name
	speed = int(force)
	fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, force)
	fmt.Println()

	//decision := bool(speed) // You can not convert int value to bool
	//fmt.Printf("speed `%T` force `%T` clculate is `%v`:",speed, decision)

	capitalA := 65
	characterA := string(capitalA)
	fmt.Printf("capitalA `%T` characterA `%T` :",capitalA, characterA)

	//capitalB := 66.6
	//characterB := string(capitalB) // Cannot convert expression of type 'float64' to type 'string'
	//fmt.Printf("capitalB `%T` characterB `%T` :",capitalB, characterB)
	fmt.Println()



}