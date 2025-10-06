// Every Go file starts by declaring its package. All files in the same
// folder that declare `package main` are compiled together as one program.
package main

// We import the `fmt` package so we can print text to the screen.
import "fmt"

// bye prints a simple farewell message to the console.
func bye() {
	fmt.Println("goodbye")
}
