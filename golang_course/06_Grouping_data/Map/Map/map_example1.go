package main

import "fmt"

func main() {
	var menu map[string]float64
	menu = map[string]float64{
		"eggs":  70.0,
		"Milk":  30.0,
		"water": 15.0,
	}
	fmt.Print(menu, "\n")

	menu["food"] = 20.01	   // Add a new key-value pair
	fmt.Print(menu, "\n")

	menu["food"] = 40.01        // Update value
	fmt.Print(menu, "\n")

	foodPrice := menu["food"]    // Get value for specific kay
	fmt.Print(foodPrice, "\n")

	for key, value := range menu{
		fmt.Print(key, "\t",value , "\n")
	}
}