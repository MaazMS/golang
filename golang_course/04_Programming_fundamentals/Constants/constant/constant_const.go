package main

import "fmt"

func constdeclare1() {

	const rightangle int = 90
	const name string = "Maaz"
	const code bool = true

	fmt.Printf("%v %T \t\n", rightangle, rightangle)
	fmt.Printf("%v %T \t\n", name, name)
	fmt.Printf("%v %T \t\n\n", code, code)
}
func constdeclare2() {

	const strightangle = 180
	const lastname = "Shaikh"
	const correct = false

	fmt.Printf("%v %T \t\n", strightangle, strightangle)
	fmt.Printf("%v %T \t\n", lastname, lastname)
	fmt.Printf("%v %T \t\n", correct, correct)
}
func main() {

	constdeclare1()
	constdeclare2()
}

/*
gopls requires a module at the root of your workspace.
You can work with multiple modules by opening each one as a workspace folder.
Improvements to this workflow will be coming soon (https://github.com/golang/go/issues/32394),
and you can learn more here: https://github.com/golang/go/issues/36899.



{
	"resource": "/home/maaz/github/LearningGO/code/src/constant/constant_const.go",
	"owner": "_generated_diagnostic_collection_name_#0",
	"severity": 8,
	"message": "gopls requires a module at the root of your workspace.\nYou can work with multiple modules by opening each one as a workspace folder.\nImprovements to this workflow will be coming soon (https://github.com/golang/go/issues/32394),\nand you can learn more here: https://github.com/golang/go/issues/36899.",
	"startLineNumber": 1,
	"startColumn": 1,
	"endLineNumber": 1,
	"endColumn": 13
}

*/
