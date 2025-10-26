## Callbacks

A callback is a function that is passed as an argument to another function. While functional programming patterns are not commonly recommended in Go, it's important to understand callbacks as they are useful in certain scenarios, especially when working with libraries or APIs.

### Understanding Callbacks

In this example, we'll demonstrate how to use callbacks by:
1. Creating a `sum` function that calculates the total of variadic integer parameters
2. Creating an `even` function that filters even numbers and applies a callback function to them

### Example: Using Callbacks with Variadic Parameters

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	// Direct function call with variadic parameters
	total := sum(numbers...)
	fmt.Println("Sum of numbers:\t", total)

	// Using callback: pass the sum function as an argument to even function
	evenSum := even(sum, numbers...)
	fmt.Println("Sum of even numbers:\t", evenSum)
}

// sum calculates the total of variadic integer parameters
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// even filters even numbers and applies the callback function to them
// f: callback function that takes variadic int parameters and returns int
// nums: variadic integer parameters to filter
func even(f func(nums ...int) int, nums ...int) int {
	var evenNumbers []int
	
	// Filter even numbers
	for _, num := range nums {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	
	// Apply callback function to filtered even numbers
	result := f(evenNumbers...)
	return result
}

/* Output:
Sum of numbers:     45
Sum of even numbers: 20
*/
```

### Key Concepts:

1. **Callback Function**: A function passed as an argument to another function
2. **Function Signature**: The callback function signature must match what the receiving function expects
3. **Variadic Parameters**: Using `...` to accept a variable number of arguments
4. **Function Composition**: Combining multiple functions to achieve complex behavior
5. **Filter and Apply Pattern**: Common pattern where you filter data and then apply a function to the filtered results

### When to Use Callbacks:

- **Library Integration**: When working with libraries that require callback functions
- **Event Handling**: For handling events or asynchronous operations
- **Custom Iterators**: When you need custom iteration logic
- **Functional Programming**: For implementing functional programming patterns (though not idiomatic Go)

### Go Best Practices:

- Use callbacks sparingly in Go
- Prefer interfaces for polymorphism over function callbacks
- Keep callback signatures simple and clear
- Document callback function requirements clearly