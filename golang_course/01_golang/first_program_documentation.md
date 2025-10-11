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


1.  Every Go source file starts with a `package` declaration.
1. The package name usually matches the directory name. The special package name `main` is used for programs that are meant to be compiled as executables.
1. A program’s entry point is the `main` function inside `package main`.
1. If you use code from another package, you must `import` that package in the file that uses it.

## main() Function 
1. `func main()` is the entry point of an executable program.    
2. It does not take arguments and does not return a value.  
3. Go calls `main()` automatically; you never call it yourself.  
3. Each executable (command) you build must have exactly one `package main` with a `main()` function in the build target.

## init() Function   
1. `init()` is an optional function that runs automatically to initialize a package before it is used.  
2. It takes no arguments and returns no value.  
3. A package may have zero or more `init()` functions (even multiple per file); all run after package-level variables are initialized and before `main()` (for the main package).  
4. You cannot call `init()` yourself; it is invoked by the runtime in import order.   
5. Typical use cases: register handlers, validate configuration, set up package-level state. Avoid heavy work or long-running tasks.

Note: You normally do not write semicolons in Go; the lexer inserts them automatically according to simple rules.   
Note: The program starts at `func main()`; when `main()` returns (or the process calls `os.Exit`), the program ends.

## Hello World and Flow of Control  
Flow of control:
    1. Sequential flow (top to bottom)
    1. Iteration (loops)
    1. Conditionals

Example:
```go
package main

import "fmt"

func init() {
    // Runs before main.
    fmt.Println("init: setup")
}

func main() {
    fmt.Println("hello, world") // sequential
    for i := 0; i < 3; i++ {    // iteration
        if i%2 == 0 {           // conditional
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }
}
```

## Key Points to Remember

- **Every Go program must have a `main` package**
- **Every executable Go program must have exactly one `main` function**
- **The `main` function is the entry point** - where program execution begins
- **Go automatically calls the `main` function** - you don't need to call it yourself
- **Other functions must be explicitly called** to execute
- **Use `func` keyword** to declare functions 