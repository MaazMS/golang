# Go Slices - Complete Guide

## Overview

A slice in Go is a powerful, lightweight data structure that provides a more flexible alternative to arrays. Slices are built on top of arrays but offer dynamic sizing, making them the preferred choice for most data collection needs in Go. They are reference types and provide efficient access to underlying array elements.

## Key Features

1. **Dynamic size**: Slices can grow and shrink as needed
2. **Lightweight**: More efficient than arrays for most use cases
3. **Reference type**: Slices are passed by reference
4. **Zero-indexed**: Index starts from 0
5. **Homogeneous**: All elements must be of the same type
6. **Built on arrays**: Slices are views into underlying arrays

## Why Use Slices?

### Advantages over Arrays
- **Dynamic sizing**: No need to specify size at declaration
- **Memory efficiency**: Only allocate what you need
- **Built-in functions**: Rich set of operations (append, copy, etc.)
- **Flexibility**: Easy to modify and manipulate

### When to Use Slices vs Arrays
```go
// Use arrays when:
// - Fixed size is known at compile time
// - Memory layout is critical
// - Working with C libraries

// Use slices when:
// - Size is unknown or variable
// - Need dynamic operations
// - Most common use case
```

## Basic Syntax

### Declaration and Initialization

```go
// Method 1: Zero value (nil slice)
var s1 []int

// Method 2: Empty slice
var s2 []int = []int{}

// Method 3: Short declaration with values
s3 := []string{"hello", "world"}

// Method 4: Using make
s4 := make([]int, 5)        // length 5, capacity 5
s5 := make([]int, 5, 10)    // length 5, capacity 10

// Method 5: Slice from array
arr := [5]int{1, 2, 3, 4, 5}
s6 := arr[1:4]              // [2, 3, 4]
```

## Slice Components

### The Three Components
1. **Pointer**: Points to the underlying array element
2. **Length**: Number of elements in the slice
3. **Capacity**: Maximum number of elements from start to end of array

```go
package main

import "fmt"

func main() {
    s := make([]int, 3, 5)
    fmt.Printf("Length: %d\n", len(s))     // 3
    fmt.Printf("Capacity: %d\n", cap(s))   // 5
    fmt.Printf("Slice: %v\n", s)           // [0 0 0]
}
```

## Basic Examples

### Example 1: Basic Slice Operations

```go
package main

import "fmt"

func main() {
    // Create a slice
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Original:", numbers)

    // Access elements
    fmt.Println("First element:", numbers[0])
    fmt.Println("Last element:", numbers[len(numbers)-1])

    // Modify elements
    numbers[0] = 10
    fmt.Println("After modification:", numbers)

    // Append elements
    numbers = append(numbers, 6, 7, 8)
    fmt.Println("After append:", numbers)

    // Slice operations
    subSlice := numbers[1:4]
    fmt.Println("Sub-slice [1:4]:", subSlice)

    // Length and capacity
    fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
}
```

### Example 2: Different Data Types

```go
package main

import "fmt"

func main() {
    // String slice
    names := []string{"Alice", "Bob", "Charlie"}
    fmt.Println("Names:", names)

    // Float slice
    prices := []float64{19.99, 29.99, 39.99}
    fmt.Println("Prices:", prices)

    // Boolean slice
    flags := []bool{true, false, true, false}
    fmt.Println("Flags:", flags)

    // Slice of slices
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("Matrix:", matrix)
}
```

### Example 3: Slice Creation Methods

```go
package main

import "fmt"

func main() {
    // From array
    arr := [5]int{1, 2, 3, 4, 5}
    slice1 := arr[:]        // All elements
    slice2 := arr[1:4]      // Elements 1-3
    slice3 := arr[:3]       // First 3 elements
    slice4 := arr[2:]       // From index 2 to end

    fmt.Println("Array:", arr)
    fmt.Println("Slice1 (all):", slice1)
    fmt.Println("Slice2 [1:4]:", slice2)
    fmt.Println("Slice3 [:3]:", slice3)
    fmt.Println("Slice4 [2:]:", slice4)

    // Using make
    slice5 := make([]int, 3, 6)
    fmt.Printf("Make slice: %v, len=%d, cap=%d\n", slice5, len(slice5), cap(slice5))

    // Nil slice
    var nilSlice []int
    fmt.Printf("Nil slice: %v, len=%d, cap=%d, nil=%t\n", 
        nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
}
```

## Advanced Examples

### Example 4: Slice Manipulation

```go
package main

import "fmt"

func main() {
    // Create initial slice
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Original:", numbers)

    // Insert at beginning
    numbers = append([]int{0}, numbers...)
    fmt.Println("After prepend:", numbers)

    // Insert at middle
    index := 3
    numbers = append(numbers[:index], append([]int{99}, numbers[index:]...)...)
    fmt.Println("After insert at index 3:", numbers)

    // Remove element at index
    index = 2
    numbers = append(numbers[:index], numbers[index+1:]...)
    fmt.Println("After remove at index 2:", numbers)

    // Reverse slice
    for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
        numbers[i], numbers[j] = numbers[j], numbers[i]
    }
    fmt.Println("After reverse:", numbers)
}
```

### Example 5: Slice Functions

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Original:", numbers)

    // Sort slice
    sort.Ints(numbers)
    fmt.Println("Sorted:", numbers)

    // Search in sorted slice
    index := sort.SearchInts(numbers, 25)
    fmt.Printf("25 found at index: %d\n", index)

    // Copy slice
    original := []int{1, 2, 3, 4, 5}
    copied := make([]int, len(original))
    copy(copied, original)
    copied[0] = 99
    fmt.Println("Original:", original)
    fmt.Println("Copied:", copied)

    // Slice comparison
    slice1 := []int{1, 2, 3}
    slice2 := []int{1, 2, 3}
    fmt.Printf("Slices equal: %t\n", equalSlices(slice1, slice2))
}

func equalSlices(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
```

### Example 6: Slice as Stack

```go
package main

import "fmt"

type Stack struct {
    items []int
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    index := len(s.items) - 1
    item := s.items[index]
    s.items = s.items[:index]
    return item, true
}

func (s *Stack) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func main() {
    stack := &Stack{}

    // Push elements
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    fmt.Println("Stack after push:", stack.items)

    // Peek at top
    if top, ok := stack.Peek(); ok {
        fmt.Println("Top element:", top)
    }

    // Pop elements
    for !stack.IsEmpty() {
        if item, ok := stack.Pop(); ok {
            fmt.Println("Popped:", item)
        }
    }
}
```

### Example 7: Slice as Queue

```go
package main

import "fmt"

type Queue struct {
    items []int
}

func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

func (q *Queue) Front() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func main() {
    queue := &Queue{}

    // Enqueue elements
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)

    fmt.Println("Queue after enqueue:", queue.items)

    // Front element
    if front, ok := queue.Front(); ok {
        fmt.Println("Front element:", front)
    }

    // Dequeue elements
    for !queue.IsEmpty() {
        if item, ok := queue.Dequeue(); ok {
            fmt.Println("Dequeued:", item)
        }
    }
}
```

## Slice Internals and Memory Management

### Example 8: Understanding Slice Internals

```go
package main

import "fmt"

func main() {
    // Create slice with make
    s := make([]int, 3, 5)
    fmt.Printf("Initial: %v, len=%d, cap=%d\n", s, len(s), cap(s))

    // Append within capacity
    s = append(s, 1, 2)
    fmt.Printf("After append: %v, len=%d, cap=%d\n", s, len(s), cap(s))

    // Append beyond capacity (reallocation)
    s = append(s, 3, 4, 5)
    fmt.Printf("After reallocation: %v, len=%d, cap=%d\n", s, len(s), cap(s))

    // Slice sharing
    original := []int{1, 2, 3, 4, 5}
    shared := original[1:4]
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Shared: %v\n", shared)

    // Modifying shared slice affects original
    shared[0] = 99
    fmt.Printf("After modification:\n")
    fmt.Printf("Original: %v\n", original)
    fmt.Printf("Shared: %v\n", shared)
}
```

### Example 9: Slice Memory Optimization

```go
package main

import "fmt"

func main() {
    // Inefficient: creates new slice each time
    var inefficient []int
    for i := 0; i < 1000; i++ {
        inefficient = append(inefficient, i)
    }
    fmt.Printf("Inefficient: len=%d, cap=%d\n", len(inefficient), cap(inefficient))

    // Efficient: pre-allocate capacity
    efficient := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        efficient = append(efficient, i)
    }
    fmt.Printf("Efficient: len=%d, cap=%d\n", len(efficient), cap(efficient))

    // Most efficient: set length and assign directly
    mostEfficient := make([]int, 1000)
    for i := 0; i < 1000; i++ {
        mostEfficient[i] = i
    }
    fmt.Printf("Most efficient: len=%d, cap=%d\n", len(mostEfficient), cap(mostEfficient))
}
```

## Common Patterns and Best Practices

### Pattern 1: Slice Filtering

```go
func filterEven(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

func filterEvenInPlace(numbers []int) []int {
    writeIndex := 0
    for _, num := range numbers {
        if num%2 == 0 {
            numbers[writeIndex] = num
            writeIndex++
        }
    }
    return numbers[:writeIndex]
}
```

### Pattern 2: Slice Mapping

```go
func mapToStrings(numbers []int) []string {
    result := make([]string, len(numbers))
    for i, num := range numbers {
        result[i] = fmt.Sprintf("Number: %d", num)
    }
    return result
}
```

### Pattern 3: Slice Reduction

```go
func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func max(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}
```

## Slice vs Array Comparison

### Example 10: Performance Comparison

```go
package main

import (
    "fmt"
    "time"
)

func arraySum(arr [1000]int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func sliceSum(slice []int) int {
    sum := 0
    for _, v := range slice {
        sum += v
    }
    return sum
}

func main() {
    // Array
    arr := [1000]int{}
    for i := range arr {
        arr[i] = i
    }

    // Slice
    slice := make([]int, 1000)
    for i := range slice {
        slice[i] = i
    }

    // Benchmark array
    start := time.Now()
    for i := 0; i < 100000; i++ {
        arraySum(arr)
    }
    arrayTime := time.Since(start)

    // Benchmark slice
    start = time.Now()
    for i := 0; i < 100000; i++ {
        sliceSum(slice)
    }
    sliceTime := time.Since(start)

    fmt.Printf("Array time: %v\n", arrayTime)
    fmt.Printf("Slice time: %v\n", sliceTime)
    fmt.Printf("Performance difference: %.2fx\n", float64(sliceTime)/float64(arrayTime))
}
```

## Error Handling and Edge Cases

### Example 11: Safe Slice Operations

```go
package main

import (
    "errors"
    "fmt"
)

var ErrIndexOutOfRange = errors.New("index out of range")
var ErrEmptySlice = errors.New("slice is empty")

func safeGet(slice []int, index int) (int, error) {
    if index < 0 || index >= len(slice) {
        return 0, ErrIndexOutOfRange
    }
    return slice[index], nil
}

func safePop(slice []int) ([]int, int, error) {
    if len(slice) == 0 {
        return slice, 0, ErrEmptySlice
    }
    lastIndex := len(slice) - 1
    last := slice[lastIndex]
    return slice[:lastIndex], last, nil
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}

    // Safe access
    if value, err := safeGet(numbers, 2); err == nil {
        fmt.Println("Value at index 2:", value)
    } else {
        fmt.Println("Error:", err)
    }

    if value, err := safeGet(numbers, 10); err == nil {
        fmt.Println("Value at index 10:", value)
    } else {
        fmt.Println("Error:", err)
    }

    // Safe pop
    for len(numbers) > 0 {
        var last int
        var err error
        numbers, last, err = safePop(numbers)
        if err == nil {
            fmt.Println("Popped:", last)
        } else {
            fmt.Println("Error:", err)
        }
    }
}
```

## Conclusion

Slices are the most commonly used data structure in Go for collections. They provide:

1. **Flexibility**: Dynamic sizing and easy manipulation
2. **Efficiency**: Built on arrays with optimized operations
3. **Simplicity**: Easy to use with rich built-in functions
4. **Memory management**: Automatic handling of underlying arrays
5. **Performance**: Good balance of speed and memory usage

Understanding slices is essential for effective Go programming, as they are used in almost every Go program for data collection and manipulation.
