# Complete Guide to Concurrency in Go

Go is designed for concurrency, and understanding concurrent programming is essential for writing efficient Go applications. This guide covers goroutines, synchronization, and common concurrency patterns.

## Table of Contents
1. [Concurrency vs Parallelism](#concurrency-vs-parallelism)
2. [Goroutines](#goroutines)
3. [WaitGroup](#waitgroup)
4. [Race Conditions](#race-conditions)
5. [Mutex](#mutex)
6. [Runtime Scheduling](#runtime-scheduling)
7. [Best Practices](#best-practices)
8. [Common Patterns](#common-patterns)

## Concurrency vs Parallelism

Understanding the difference between concurrency and parallelism is crucial for effective Go programming.

### Key Distinctions

**Concurrency** is about dealing with lots of things at once. It's the composition of independently executing processes that may or may not run simultaneously.

**Parallelism** is about doing lots of things at once. It's the simultaneous execution of (possibly related) computations.

### The Go Philosophy

- **Concurrency**: Go excels at concurrent programming through goroutines and channels
- **Parallelism**: Go can achieve parallelism when running on multiple CPU cores
- **Design**: Go is designed for concurrency, which enables parallelism

### Visual Representation

```
Concurrency:     Parallelism:
Task A ──┐       Task A ──┐
Task B ──┼───    Task B ──┼───
Task C ──┘       Task C ──┘
```

### References
- [Go Blog: Concurrency is not Parallelism](https://blog.golang.org/waza-talk)
- [Understanding Concurrency and Parallelism in Go](https://spiralscout.com/blog/understanding-concurrency-and-parallelism-in-golang)

## Goroutines

Goroutines are Go's lightweight threads that enable concurrent execution. They are managed by the Go runtime scheduler.

### What are Goroutines?

A goroutine is a function or method that executes independently and simultaneously in connection with other goroutines. Every concurrently executing activity in Go is called a goroutine.

### Key Characteristics

1. **Lightweight**: Goroutines are much lighter than OS threads
2. **Managed**: Goroutines are managed by the Go runtime scheduler
3. **Background Execution**: Goroutines always work in the background
4. **Hierarchical**: All goroutines work under the main goroutine
5. **Lifecycle**: If the main goroutine terminates, all goroutines in the program also terminate

### Creating Goroutines

Goroutines are created by using the `go` keyword as a prefix to a function or method call.

#### Syntax
```go
func functionName() {
    // statements
}

go functionName()
```

#### Example: Basic Goroutine

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Start a goroutine
	go sayHello("World")
	
	// Main goroutine continues
	sayHello("Go")
	
	// Wait to see goroutine output
	time.Sleep(1 * time.Second)
}

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello %s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

/* Output:
Hello Go: 0
Hello World: 0
Hello Go: 1
Hello World: 1
Hello Go: 2
Hello World: 2
*/
```

### Goroutine Lifecycle

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Initial goroutines:", runtime.NumGoroutine())
	
	// Start multiple goroutines
	for i := 0; i < 3; i++ {
		go worker(i)
	}
	
	fmt.Println("After starting goroutines:", runtime.NumGoroutine())
	
	// Wait for goroutines to complete
	time.Sleep(2 * time.Second)
	fmt.Println("After goroutines complete:", runtime.NumGoroutine())
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

/* Output:
Initial goroutines: 1
After starting goroutines: 4
Worker 0 starting
Worker 1 starting
Worker 2 starting
Worker 0 done
Worker 1 done
Worker 2 done
After goroutines complete: 1
*/
```

## WaitGroup

A WaitGroup waits for a collection of goroutines to finish. It's essential for coordinating goroutines in concurrent programs.

### How WaitGroup Works

1. **Add**: The main goroutine calls `Add` to set the number of goroutines to wait for
2. **Done**: Each goroutine calls `Done` when finished
3. **Wait**: `Wait` blocks until all goroutines have finished

### Example: Using WaitGroup

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Initial Goroutines:", runtime.NumGoroutine())

	// Add 1 to WaitGroup counter
	wg.Add(1)
	go foo()
	
	// Run bar in main goroutine
	bar()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	
	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines finished")
}

func foo() {
	for i := 0; i <= 5; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // Decrement WaitGroup counter
}

func bar() {
	for i := 0; i <= 5; i++ {
		fmt.Println("bar:", i)
	}
}

/* Output:
OS: linux
ARCH: amd64
CPUs: 4
Initial Goroutines: 1
bar: 0
bar: 1
bar: 2
bar: 3
bar: 4
bar: 5
CPUs: 4
Goroutines: 2
foo: 0
foo: 1
foo: 2
foo: 3
foo: 4
foo: 5
All goroutines finished
*/
```

### Multiple Goroutines with WaitGroup

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	
	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All workers completed")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done is called even if function panics
	
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

/* Output:
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 4 starting
Worker 5 starting
Worker 1 done
Worker 2 done
Worker 3 done
Worker 4 done
Worker 5 done
All workers completed
*/
```

## Race Conditions

A race condition occurs when two or more goroutines have access to the same resource (such as a variable or data structure) and attempt to read and write to that resource without proper synchronization.

### Example: Race Condition

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	const goroutines = 100

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			
			// Race condition: multiple goroutines accessing counter
			temp := counter
			runtime.Gosched() // Yield to other goroutines
			temp++
			counter = temp
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter) // Will be less than 100
}

/* Output (example):
Final counter: 87
*/
```

### Detecting Race Conditions

Use the `-race` flag when running Go programs:

```bash
go run -race main.go
```

This will detect and report race conditions in your code.

## Mutex

A mutex (mutual exclusion lock) allows us to lock our code so that only one goroutine can access that locked chunk of code at a time.

### Example: Using Mutex to Fix Race Conditions

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Initial Goroutines:", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const goroutines = 100
	var mu sync.Mutex

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			
			mu.Lock() // Lock the critical section
			temp := counter
			runtime.Gosched() // Yield to other goroutines
			temp++
			counter = temp
			mu.Unlock() // Unlock the critical section
		}()
	}

	wg.Wait()
	fmt.Println("Final Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final counter:", counter) // Will be exactly 100
}

/* Output:
CPUs: 4
Initial Goroutines: 1
Final Goroutines: 1
Final counter: 100
*/
```

### Mutex Best Practices

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // Use defer to ensure unlock
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", counter.Value())
}

/* Output:
Final value: 1000
*/
```

## Runtime Scheduling

### Gosched

`runtime.Gosched()` yields the processor, allowing other goroutines to run. It's useful for demonstrating goroutine scheduling.

```go
package main

import "runtime"

func main() {
	go say("World")
	say("Hello")
}

func say(s string) {
	for i := 0; i <= 3; i++ {
		runtime.Gosched() // Yield to other goroutines
		println(i, s)
	}
}

/* Output without Gosched:
0 Hello
1 Hello
2 Hello
3 Hello
*/

/* Output with Gosched:
0 Hello
0 World
1 Hello
1 World
2 Hello
2 World
3 Hello
3 World
*/
```

### Runtime Information

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	
	// Get current goroutine
	fmt.Println("Current Goroutine ID:", runtime.GoID())
}

/* Output:
OS: linux
Architecture: amd64
Number of CPUs: 4
Number of Goroutines: 1
Current Goroutine ID: 1
*/
```

## Best Practices

### 1. Always Use WaitGroup for Coordination

```go
// Good
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()

// Bad - using sleep
go func() {
    // work
}()
time.Sleep(1 * time.Second) // Unreliable
```

### 2. Use Mutex for Shared Resources

```go
// Good
type SafeData struct {
    mu   sync.Mutex
    data int
}

func (s *SafeData) Update(value int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data = value
}

// Bad - unprotected shared data
var data int
func update(value int) {
    data = value // Race condition
}
```

### 3. Use Defer with Mutex

```go
// Good
mu.Lock()
defer mu.Unlock()
// critical section

// Bad - manual unlock
mu.Lock()
// critical section
mu.Unlock() // Might be forgotten
```

### 4. Avoid Race Conditions

```go
// Good - use channels for communication
ch := make(chan int)
go func() {
    ch <- 42
}()
value := <-ch

// Bad - shared memory
var value int
go func() {
    value = 42 // Race condition
}()
```

### 5. Proper Error Handling

```go
// Good
go func() {
    defer wg.Done()
    if err := doWork(); err != nil {
        // handle error
    }
}()

// Bad - ignoring errors
go func() {
    defer wg.Done()
    doWork() // Error ignored
}()
```

## Common Patterns

### 1. Worker Pool Pattern

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= 9; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}
```

### 2. Fan-out/Fan-in Pattern

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	input := make(chan int)
	output := make(chan int)
	
	// Fan-out: distribute work
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()
	
	// Start multiple workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range input {
				output <- n * n
			}
		}()
	}
	
	// Fan-in: collect results
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Print results
	for result := range output {
		fmt.Println("Result:", result)
	}
}
```

### 3. Pipeline Pattern

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int)
	squares := make(chan int)
	
	// Stage 1: Generate numbers
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	// Stage 2: Square numbers
	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()
	
	// Stage 3: Print results
	for result := range squares {
		fmt.Println("Square:", result)
	}
}
```

## Summary

Concurrency in Go is built around:

1. **Goroutines**: Lightweight threads for concurrent execution
2. **Channels**: Communication mechanism between goroutines
3. **WaitGroup**: Coordination primitive for waiting on goroutines
4. **Mutex**: Mutual exclusion for protecting shared resources
5. **Runtime**: Built-in scheduler and runtime information

Key principles:
- **Don't communicate by sharing memory; share memory by communicating**
- **Use channels for communication, mutexes for synchronization**
- **Always coordinate goroutines properly**
- **Detect and fix race conditions**
- **Follow established patterns for common problems**

Understanding these concepts is essential for writing efficient, safe, and maintainable concurrent Go programs.
