package main

import "fmt"

func main()  {

	var safe bool // initial value is false
	safe, speed := true, 50  // update safe value false to true
	// safe := true // No new variables on left side of :=
	fmt.Println(safe, speed )
}
