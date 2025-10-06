// The main package builds an executable program.
package main

// main is the entry point of a Go program. When you run `go run` or
// execute the compiled binary, Go starts here.
//
// Because `hey` and `bye` are defined in other files within the same
// package (same folder, same `package main`), we can call them directly.
func main() {
	hey() // call the greeting function
	bye() // call the farewell function
}
