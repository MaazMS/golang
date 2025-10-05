package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH \t\t", runtime.GOARCH)
	fmt.Println("CUPs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	baar()

	fmt.Println("CUPs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines \t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i <= 5; i++ {
		fmt.Println("foo\t", i)
	}
	wg.Done()

}
func baar() {
	for i := 0; i <= 5; i++ {
		fmt.Println("baar \t", i)
	}

}
