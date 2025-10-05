package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS     \t", runtime.GOOS)
	fmt.Println("GOARCH   \t", runtime.GOARCH)
	fmt.Println("Compiler \t", runtime.Compiler)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
