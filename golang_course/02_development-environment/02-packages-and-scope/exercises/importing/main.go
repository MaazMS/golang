package main
// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

import f1 "fmt"
import f2 "fmt"
import f3 "fmt"

func main()  {
	f1.Println("hello")
	f2.Println("hey")
	f3.Println("hi")
}