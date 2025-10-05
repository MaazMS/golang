# First Go Program Documentation

## Which keyword do you need to use to define a package?

The `package` keyword allows you to define a package for a Go file.

```go 
package main

func main() {
}
```

## What is the purpose of using package main in a program?

1. **To create an executable Go program**
2. The `main` package is special in Go - it tells the Go compiler that this is an executable program, not a library
3. **Every Go program must have exactly one `main` package**

```go
package main

func main() {
}
```

## What is the purpose of func main in a program?

1. **It serves as the entry point** - Go starts executing the program from the `main` function
2. **Go automatically calls the main function** when you run the program
3. **Every executable Go program must have exactly one `main` function**

```go
package main

func main() {
    // This is where your program starts executing
}
```

## Which keyword is used to declare a new function?

The `func` keyword is used to declare functions in Go.

```go
func functionName() {
    // function body
}
```

## What is a function?

A function is a reusable and executable block of code that:

1. **Acts like a mini-program** - it performs a specific task
2. **Helps organize code** into logical, manageable pieces
3. **Can accept parameters** (input values)
4. **Can return values** (output results)
5. **Promotes code reusability** - you can call the same function multiple times

```go
func greet(name string) string {
    return "Hello, " + name + "!"
}
```

## Do you have to call the main function yourself?

**No**, Go calls the `main` function automatically when you run the program.

## Do you have to call a function to execute it?
_(except the main function)_

**Yes**, you must explicitly call a function for it to execute. The `main` function is the only exception - Go calls it automatically.

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello, World!")
}

func main() {
    greet() // You must call greet() for it to execute
}
```

## What does the following program print?

```go
package main

func main() {
}
```

**Answer**: This is a correct program, but it doesn't print anything.

**Explanation**:
1. The program compiles and runs successfully
2. It produces no output because there are no print statements
3. The `main` function is empty, so it executes and exits immediately

## Key Points to Remember

- **Every Go program must have a `main` package**
- **Every executable Go program must have exactly one `main` function**
- **The `main` function is the entry point** - where program execution begins
- **Go automatically calls the `main` function** - you don't need to call it yourself
- **Other functions must be explicitly called** to execute
- **Use `func` keyword** to declare functions 