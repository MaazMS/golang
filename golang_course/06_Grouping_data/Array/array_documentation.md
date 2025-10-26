# Go Arrays - Complete Guide

## Overview

An array in Go is a fixed-length sequence of elements of the same type. Arrays are value types in Go, meaning when you assign an array to another variable or pass it to a function, a complete copy is made. Arrays provide a foundation for slices and are useful when you need a fixed-size collection with known dimensions.

## Key Features

1. **Fixed length**: Size must be specified at declaration and cannot change
2. **Homogeneous elements**: All elements must be of the same type
3. **Zero-indexed**: Index starts from 0
4. **Value type**: Arrays are copied when assigned or passed to functions
5. **Mutable**: Elements can be modified after creation
6. **Memory efficient**: Contiguous memory allocation

## Why Use Arrays?

### Advantages
- **Predictable memory usage**: Fixed size means known memory footprint
- **Performance**: Direct memory access, no overhead
- **Type safety**: Compile-time size checking
- **Foundation for slices**: Arrays underlie slice implementation

### When to Use Arrays vs Slices
```go
// Use arrays when:
// - Size is known at compile time
// - Memory layout is critical
// - Working with C libraries
// - Need value semantics (copying)

// Use slices when:
// - Size is unknown or variable
// - Need dynamic operations
// - Most common use case
```

## Basic Syntax

### Declaration and Initialization

```go
// Method 1: Declaration then assignment
var arr1 [5]int
arr1[0] = 1
arr1[1] = 2
arr1[2] = 3
arr1[3] = 4
arr1[4] = 5

// Method 2: Declaration with initialization
var arr2 [5]int = [5]int{1, 2, 3, 4, 5}

// Method 3: Short declaration with initialization
arr3 := [5]int{1, 2, 3, 4, 5}

// Method 4: Partial initialization (remaining elements are zero values)
arr4 := [5]int{1, 2, 3} // [1, 2, 3, 0, 0]

// Method 5: Initialize specific indices
arr5 := [5]int{0: 10, 2: 30, 4: 50} // [10, 0, 30, 0, 50]

// Method 6: Let compiler determine size
arr6 := [...]int{1, 2, 3, 4, 5} // [5]int{1, 2, 3, 4, 5}
```

## Basic Examples

### Example 1: Basic Array Operations

```go
package main

import "fmt"

func main() {
    // Create an array
    numbers := [5]int{10, 20, 30, 40, 50}
    fmt.Println("Original array:", numbers)

    // Access elements
    fmt.Println("First element:", numbers[0])
    fmt.Println("Last element:", numbers[len(numbers)-1])

    // Modify elements
    numbers[0] = 100
    numbers[4] = 500
    fmt.Println("After modification:", numbers)

    // Array length
    fmt.Println("Array length:", len(numbers))

    // Iterate over array
    fmt.Println("Array elements:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("  [%d] = %d\n", i, numbers[i])
    }

    // Range loop
    fmt.Println("Using range:")
    for index, value := range numbers {
        fmt.Printf("  [%d] = %d\n", index, value)
    }
}
```

### Example 2: Different Data Types

```go
package main

import "fmt"

func main() {
    // String array
    names := [3]string{"Alice", "Bob", "Charlie"}
    fmt.Println("Names:", names)

    // Float array
    prices := [4]float64{19.99, 29.99, 39.99, 49.99}
    fmt.Println("Prices:", prices)

    // Boolean array
    flags := [3]bool{true, false, true}
    fmt.Println("Flags:", flags)

    // Array of arrays (2D array)
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("Matrix:", matrix)

    // Access 2D array elements
    fmt.Println("Element at [1][2]:", matrix[1][2])
}
```

### Example 3: Array Initialization Patterns

```go
package main

import "fmt"

func main() {
    // Zero values
    var zeroArray [5]int
    fmt.Println("Zero array:", zeroArray)

    // Partial initialization
    partialArray := [5]int{1, 2}
    fmt.Println("Partial array:", partialArray)

    // Specific index initialization
    specificArray := [5]int{0: 100, 2: 300, 4: 500}
    fmt.Println("Specific indices:", specificArray)

    // Compiler-determined size
    autoSizeArray := [...]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Printf("Auto size array: %v (length: %d)\n", autoSizeArray, len(autoSizeArray))

    // Mixed initialization
    mixedArray := [5]int{1, 2, 3} // Last two elements are zero
    fmt.Println("Mixed array:", mixedArray)
}
```

## Advanced Examples

### Example 4: Array Functions and Operations

```go
package main

import "fmt"

func main() {
    numbers := [5]int{64, 34, 25, 12, 22}
    fmt.Println("Original:", numbers)

    // Sum of array
    sum := sumArray(numbers)
    fmt.Println("Sum:", sum)

    // Find maximum
    max := maxArray(numbers)
    fmt.Println("Maximum:", max)

    // Find minimum
    min := minArray(numbers)
    fmt.Println("Minimum:", min)

    // Reverse array
    reversed := reverseArray(numbers)
    fmt.Println("Reversed:", reversed)

    // Search in array
    index := searchArray(numbers, 25)
    if index != -1 {
        fmt.Printf("Found 25 at index %d\n", index)
    } else {
        fmt.Println("25 not found")
    }
}

func sumArray(arr [5]int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func maxArray(arr [5]int) int {
    if len(arr) == 0 {
        return 0
    }
    max := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    return max
}

func minArray(arr [5]int) int {
    if len(arr) == 0 {
        return 0
    }
    min := arr[0]
    for _, v := range arr {
        if v < min {
            min = v
        }
    }
    return min
}

func reverseArray(arr [5]int) [5]int {
    result := arr // Copy array
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return result
}

func searchArray(arr [5]int, target int) int {
    for i, v := range arr {
        if v == target {
            return i
        }
    }
    return -1
}
```

### Example 5: Multi-dimensional Arrays

```go
package main

import "fmt"

func main() {
    // 2D array
    var matrix [3][4]int
    fmt.Println("Empty 2D array:", matrix)

    // Initialize 2D array
    matrix2 := [3][4]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    fmt.Println("Initialized 2D array:", matrix2)

    // Access elements
    fmt.Println("Element at [1][2]:", matrix2[1][2])

    // Iterate over 2D array
    fmt.Println("2D array elements:")
    for i := 0; i < len(matrix2); i++ {
        for j := 0; j < len(matrix2[i]); j++ {
            fmt.Printf("  [%d][%d] = %d\n", i, j, matrix2[i][j])
        }
    }

    // 3D array
    var cube [2][3][4]int
    fmt.Printf("3D array dimensions: %dx%dx%d\n", 
        len(cube), len(cube[0]), len(cube[0][0]))

    // Initialize 3D array
    cube[0][0][0] = 1
    cube[1][2][3] = 24
    fmt.Println("3D array:", cube)
}
```

### Example 6: Array as Stack

```go
package main

import "fmt"

type ArrayStack struct {
    items [10]int
    top   int
}

func (s *ArrayStack) Push(item int) bool {
    if s.top >= len(s.items) {
        return false // Stack overflow
    }
    s.items[s.top] = item
    s.top++
    return true
}

func (s *ArrayStack) Pop() (int, bool) {
    if s.top <= 0 {
        return 0, false // Stack underflow
    }
    s.top--
    return s.items[s.top], true
}

func (s *ArrayStack) Peek() (int, bool) {
    if s.top <= 0 {
        return 0, false
    }
    return s.items[s.top-1], true
}

func (s *ArrayStack) IsEmpty() bool {
    return s.top == 0
}

func (s *ArrayStack) IsFull() bool {
    return s.top >= len(s.items)
}

func main() {
    stack := &ArrayStack{}

    // Push elements
    for i := 1; i <= 5; i++ {
        if stack.Push(i * 10) {
            fmt.Printf("Pushed: %d\n", i*10)
        } else {
            fmt.Println("Stack overflow!")
        }
    }

    // Peek at top
    if top, ok := stack.Peek(); ok {
        fmt.Printf("Top element: %d\n", top)
    }

    // Pop elements
    for !stack.IsEmpty() {
        if item, ok := stack.Pop(); ok {
            fmt.Printf("Popped: %d\n", item)
        }
    }
}
```

## Array Value Semantics

### Example 7: Array Copying and Value Semantics

```go
package main

import "fmt"

func main() {
    // Original array
    original := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Original:", original)

    // Assignment creates a copy
    copy := original
    fmt.Println("Copy:", copy)

    // Modify the copy
    copy[0] = 100
    fmt.Println("After modifying copy:")
    fmt.Println("Original:", original) // Unchanged
    fmt.Println("Copy:", copy)         // Changed

    // Function parameter (also creates copy)
    modifyArray(original)
    fmt.Println("After function call:")
    fmt.Println("Original:", original) // Still unchanged

    // Array comparison
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{1, 2, 4}

    fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2) // true
    fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3) // false
}

func modifyArray(arr [5]int) {
    arr[0] = 999 // This doesn't affect the original
    fmt.Println("Inside function:", arr)
}
```

### Example 8: Array Pointers

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Original:", arr)

    // Pass array by pointer to modify it
    modifyArrayByPointer(&arr)
    fmt.Println("After pointer modification:", arr)

    // Array of pointers
    var ptrArray [3]*int
    a, b, c := 10, 20, 30
    ptrArray[0] = &a
    ptrArray[1] = &b
    ptrArray[2] = &c

    fmt.Println("Array of pointers:")
    for i, ptr := range ptrArray {
        fmt.Printf("  [%d] = %d (address: %p)\n", i, *ptr, ptr)
    }
}

func modifyArrayByPointer(arr *[5]int) {
    arr[0] = 999 // This modifies the original array
    fmt.Println("Inside function:", *arr)
}
```

## Array vs Slice Comparison

### Example 9: Performance and Memory Comparison

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

    // Memory usage
    fmt.Printf("Array size: %d bytes\n", len(arr)*8) // 8 bytes per int
    fmt.Printf("Slice size: %d bytes (slice header) + %d bytes (underlying array)\n", 
        24, len(slice)*8) // 24 bytes for slice header
}
```

## Common Patterns and Best Practices

### Pattern 1: Array Initialization

```go
// Initialize with specific pattern
func initializeArray() [10]int {
    var arr [10]int
    for i := range arr {
        arr[i] = i * i // Square numbers
    }
    return arr
}

// Initialize with function
func initializeWithFunction() [5]int {
    arr := [5]int{}
    for i := range arr {
        arr[i] = fibonacci(i)
    }
    return arr
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

### Pattern 2: Array Validation

```go
func isValidArray(arr [5]int) bool {
    // Check if array contains only positive numbers
    for _, v := range arr {
        if v <= 0 {
            return false
        }
    }
    return true
}

func isSortedArray(arr [5]int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[i-1] {
            return false
        }
    }
    return true
}
```

### Pattern 3: Array Conversion

```go
func arrayToSlice(arr [5]int) []int {
    // Convert array to slice
    return arr[:]
}

func sliceToArray(slice []int) ([5]int, bool) {
    // Convert slice to array (if size matches)
    if len(slice) != 5 {
        return [5]int{}, false
    }
    var arr [5]int
    copy(arr[:], slice)
    return arr, true
}
```

## Error Handling and Edge Cases

### Example 10: Safe Array Operations

```go
package main

import (
    "errors"
    "fmt"
)

var ErrIndexOutOfRange = errors.New("index out of range")
var ErrEmptyArray = errors.New("array is empty")

func safeGet(arr [5]int, index int) (int, error) {
    if index < 0 || index >= len(arr) {
        return 0, ErrIndexOutOfRange
    }
    return arr[index], nil
}

func safeSet(arr *[5]int, index int, value int) error {
    if index < 0 || index >= len(arr) {
        return ErrIndexOutOfRange
    }
    arr[index] = value
    return nil
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}

    // Safe access
    if value, err := safeGet(arr, 2); err == nil {
        fmt.Println("Value at index 2:", value)
    } else {
        fmt.Println("Error:", err)
    }

    if value, err := safeGet(arr, 10); err == nil {
        fmt.Println("Value at index 10:", value)
    } else {
        fmt.Println("Error:", err)
    }

    // Safe modification
    if err := safeSet(&arr, 1, 99); err == nil {
        fmt.Println("Array after safe set:", arr)
    } else {
        fmt.Println("Error:", err)
    }
}
```

## Array Literals and Constants

### Example 11: Array Constants and Literals

```go
package main

import "fmt"

func main() {
    // Array literal
    literal := [3]int{1, 2, 3}
    fmt.Println("Literal:", literal)

    // Array with computed values
    computed := [5]int{}
    for i := range computed {
        computed[i] = i * 2
    }
    fmt.Println("Computed:", computed)

    // Array with function calls
    funcArray := [3]int{factorial(1), factorial(2), factorial(3)}
    fmt.Println("Function array:", funcArray)

    // Array with mixed initialization
    mixed := [5]int{1, 2} // Last three elements are zero
    fmt.Println("Mixed:", mixed)
}

func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

## Conclusion

Arrays in Go are fundamental data structures that provide:

1. **Fixed size**: Predictable memory usage
2. **Value semantics**: Copying behavior
3. **Type safety**: Compile-time checking
4. **Performance**: Direct memory access
5. **Foundation**: Base for slices and other structures

While slices are more commonly used in Go programs, arrays are essential for:
- Fixed-size collections
- Performance-critical code
- Working with C libraries
- Understanding slice internals
- Value semantics requirements

Understanding arrays is crucial for effective Go programming, especially when working with low-level operations or when you need the specific characteristics that arrays provide.
