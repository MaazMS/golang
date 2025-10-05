package main

import (
	"fmt"
	"sync"
)

/*
● in addition to the main goroutine, launch two additional goroutines
	○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/

var wg sync.WaitGroup

func main() {

	fmt.Println("function main")
	//runtime.Gosched()
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

}

func foo() {
	fmt.Println("function foo")
	wg.Done()
}
func bar() {
	fmt.Println("function bar")
	wg.Done()
}
