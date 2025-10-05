# WaitGroup 
1. A WaitGroup waits for a collection of goroutines to finish.  
1. The main goroutine calls Add to set the number of goroutines to wait for.  
1. Then each of the goroutines runs and calls Done when finished.  
1. At the same time, Wait can be used to block until all goroutines have finished.  
1. Writing concurrent code is super easy.  
```go
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

``` 


````bash
OS               linux
ARCH             amd64
CUPs             4
Goroutines       1
baar     0
baar     1
baar     2
baar     3
baar     4
baar     5
CUPs             4
Goroutines       2
foo      0
foo      1
foo      2
foo      3
foo      4
foo      5
````