## Mutex  
1. A “mutex” is a mutual exclusion lock.  
1. Mutexes allow us to lock our code so that only one goroutine can access that locked chunk of code at a time.  
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	var mu sync.Mutex

	for i := 0; i<= gs; i++{
		go func(){
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count", counter)

}

```
```bash
// without using mutex 
CPUs 4
Goroutines 1
Goroutines 2
Goroutines 3
Goroutines 4
Goroutines 5
Goroutines 6
Goroutines 7
Goroutines 8
Goroutines 9
Goroutines 10
Goroutines 11
Goroutines 12
Goroutines 13
Goroutines 14
Goroutines 15
Goroutines 16
Goroutines 17
Goroutines 18
Goroutines 19
Goroutines 20
Goroutines 21
Goroutines 22
Goroutines 23
Goroutines 24
Goroutines 25
Goroutines 26
Goroutines 27
Goroutines 28
Goroutines 29
Goroutines 30
Goroutines 31
panic: sync: negative WaitGroup counter

goroutine 19 [running]:
sync.(*WaitGroup).Add(0xc0000200e0, 0xffffffffffffffff)
        /usr/local/go/src/sync/waitgroup.go:74 +0x147
sync.(*WaitGroup).Done(...)
        /usr/local/go/src/sync/waitgroup.go:99
main.main.func1(0xc0000200d0, 0xc0000200e0)
        /home/maaz/github/LearningGO/golang_deep_dive/11_Concurrency/Mutex/main.go:27 +0x65
created by main.main
        /home/maaz/github/LearningGO/golang_deep_dive/11_Concurrency/Mutex/main.go:20 +0x1d1
```