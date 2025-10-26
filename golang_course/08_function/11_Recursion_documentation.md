## Recursion

Recursion is a programming technique where a function calls itself to solve a problem. While recursion can be elegant and intuitive for certain problems, it's important to understand its trade-offs compared to iterative solutions.

### Key Points About Recursion:

1. **Recursion vs Iteration**: Most recursive solutions can be rewritten using loops (iteration)
2. **Memory Impact**: Recursion uses the call stack, which can lead to stack overflow for deep recursion
3. **Performance**: Recursive solutions may have higher overhead due to function call overhead
4. **Base Case**: Every recursive function must have a base case to prevent infinite recursion

### Example: Factorial Calculation

This example demonstrates both recursive and iterative approaches to calculating factorials.

```go
package main

import "fmt"

func main() {
	// Direct calculation
	fmt.Println("Direct calculation:", 4*3*2*1)
	
	// Recursive approach
	recursiveResult := factorialRecursive(4)
	fmt.Println("Recursive factorial:", recursiveResult)
	
	// Iterative approach
	iterativeResult := factorialIterative(4)
	fmt.Println("Iterative factorial:", iterativeResult)
}

// factorialRecursive calculates factorial using recursion
func factorialRecursive(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: n! = n * (n-1)!
	return n * factorialRecursive(n-1)
	
	// Execution trace for factorialRecursive(4):
	// factorialRecursive(4) = 4 * factorialRecursive(3)
	// factorialRecursive(3) = 3 * factorialRecursive(2)
	// factorialRecursive(2) = 2 * factorialRecursive(1)
	// factorialRecursive(1) = 1 * factorialRecursive(0)
	// factorialRecursive(0) = 1 (base case)
	// Result: 4 * 3 * 2 * 1 * 1 = 24
}

// factorialIterative calculates factorial using a loop
func factorialIterative(n int) int {
	result := 1
	
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

/* Output:
Direct calculation: 24
Recursive factorial: 24
Iterative factorial: 24
*/
```

### When to Use Recursion:

**Good for:**
- **Tree/Graph Traversal**: Natural fit for hierarchical data structures
- **Divide and Conquer**: Problems that can be broken into smaller subproblems
- **Mathematical Definitions**: When the problem naturally follows a recursive definition
- **Backtracking**: Exploring all possible solutions

**Avoid when:**
- **Simple Iteration**: When a loop is more straightforward
- **Deep Recursion**: Risk of stack overflow
- **Performance Critical**: When function call overhead matters
- **Memory Constrained**: When stack space is limited

### Example: Tree Traversal (Good Use Case)

```go
package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Create a simple binary tree
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 4},
			Right: &TreeNode{Value: 5},
		},
		Right: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 6},
			Right: &TreeNode{Value: 7},
		},
	}
	
	fmt.Println("In-order traversal:")
	inOrderTraversal(root)
	fmt.Println()
}

// inOrderTraversal demonstrates a good use case for recursion
func inOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	
	// Recursively traverse left subtree
	inOrderTraversal(node.Left)
	
	// Process current node
	fmt.Printf("%d ", node.Value)
	
	// Recursively traverse right subtree
	inOrderTraversal(node.Right)
}

/* Output:
In-order traversal:
4 2 5 1 6 3 7
*/
```

### Best Practices:

1. **Always Define Base Case**: Ensure recursion terminates
2. **Consider Stack Depth**: Be aware of maximum recursion depth
3. **Use Tail Recursion When Possible**: Some compilers can optimize tail recursion
4. **Profile Performance**: Measure actual performance impact
5. **Consider Iterative Alternatives**: Evaluate if a loop would be simpler
6. **Document Recursive Logic**: Make the recursive structure clear

### Common Pitfalls:

- **Missing Base Case**: Leads to infinite recursion and stack overflow
- **Incorrect Recursive Call**: Wrong parameters or logic in recursive step
- **Stack Overflow**: Deep recursion exceeding available stack space
- **Performance Issues**: Unnecessary function call overhead