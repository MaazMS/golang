### Package (Go)  

**What is a package?**  
- A package is a collection of Go source files in the same directory that are compiled together.  
- Files in a package share the same package name (e.g., `package fmt`, `package main`).  
- There are two kinds of packages:  
  - **main package**: builds an executable and must define `func main()`.  
  - **library package**: reusable code that is imported by other packages.  

**Purpose of packages**  
- **Organization**: Group related code into a single unit for clarity.  
- **Reusability**: Share functionality across programs by importing packages.  
- **Encapsulation**: Control visibility with exported identifiers (names starting with an uppercase letter are exported).  
- **Namespace separation**: Avoid name collisions by qualifying with package names (e.g., `fmt.Println`).  

**Key points**  
- All files in the same directory should use the same package name.  
- `package main` is used only when building a runnable program (it contains `main()`).  
- Non-`main` packages provide functions, types, and variables to be imported elsewhere.  

**Example**  
- Executable: `package main` → has `func main()` → run with `go run .`  
- Library: `package mathutil` → imported as `import "module/mathutil"`  


## What is a package?  
- All source files that belong to the same package must live in the same directory.  
- All source files in a directory (module-aware context) should declare the same package name.  
- The `package` clause must be the first non-comment line in a Go source file.  

## How to execute multiple Go source files in the same package  
1. `go run .` — runs all files in the current module directory that are part of the `main` package.    
2. `go run *.go` — runs all `.go` files in the current directory (shell expands the glob).  
3. `go run main.go hey.go bye.go` — explicitly list files (all must be in the same package).  

Example:  
```
$ go run .
Hi
goodbye

$ go run *.go
Hi
goodbye

$ go run main.go hey.go bye.go
Hi
goodbye
```

## What are the types of packages in Go?  
There are two kinds of packages: executable and library.  
1. Executable package  
   - Must be declared as `package main`.  
   - Must define exactly one `func main()` entry point.  
   - Produces a runnable binary (`go run`, `go build`, `go install`).  
2. Library package  
   - Declared with any non-`main` name (for example, `package mathutil`).  
   - Compiles to a library that other packages can import.  

## Differences: library vs executable packages    

| Library | Executable |
|---|---|
| Created for reusability | Created to run |
| Not directly executed | Executed as a program |
| Importable by other packages | Not importable |
| Package name can be any non-`main` | Package name must be `main` |
| Does not have `func main()` | Must have `func main()` |

## Where to store source files that belong to a package?  
1. In a single directory per package (for example, `internal/mathutil/`).  

## Why is a package clause used in a Go source file?  
1. It declares which package the file belongs to, enabling the compiler and tooling to group related files.  

## Where should you put the package clause in a Go source file?  
1. As the first non-comment line in the file.    

## How many times can you use a package clause in a single source file?  
1. Once.  

## Package clause examples  
1. `package main`  
2. `package mathutil`  

## Working within a single package  
1. All files in the same package can call each other's unexported identifiers.  
2. Exported identifiers start with an uppercase letter (for example, `Add`, `HTTPClient`).  

## Which package type can `go run` execute?
1. Only `main` (executable) packages.  

## Which package types can `go build` compile?  
1. Both executable and library packages (it produces a binary only for `main` packages).  

## Which package type is used for reusability and can be imported?  
1. Library packages.  

## Go Package Management

Package management is the process of managing external libraries and packages that your Go application depends on. For detailed module information, see `08_go_module.md`.

### What is Package Management?

Package management involves:
- **Finding and installing** external packages/libraries
- **Managing versions** of dependencies
- **Resolving conflicts** between different package versions
- **Ensuring reproducibility** across different environments
- **Updating and maintaining** dependencies over time

### Benefits

- **Code Reuse**: Leverage existing, well-tested solutions
- **Faster Development**: Focus on business logic rather than reinventing wheels
- **Community Support**: Benefit from community contributions and bug fixes
- **Best Practices**: Use packages that follow Go best practices
- **Maintenance**: Automatic updates and security patches

### Example: With vs Without Package Management

**Without Package Management:**
```go
func main() {
    // HTTP server implementation
    // JSON parsing
    // Database connection
    // Authentication
    // Logging
    // Configuration management
    // ... hundreds of lines of code
}
```

**With Package Management:**
```go
package main

import (
    "github.com/gin-gonic/gin"           // Web framework
    "github.com/spf13/viper"            // Configuration
    "go.uber.org/zap"                  // Logging
    "gorm.io/gorm"                     // Database ORM
    "github.com/golang-jwt/jwt/v4"     // JWT authentication
)

func main() {
    // Focus on business logic
    r := gin.Default()
    r.GET("/api/users", getUsers)
    r.Run(":8080")
}
```

### Security Considerations

Before using any external package:

1. **Is it trustworthy?** - Check author reputation and community adoption
2. **Is it secure?** - Check for vulnerabilities: `go list -json -m all | nancy sleuth`
3. **Is it well-maintained?** - Recent commits, active issue resolution
4. **Is it compatible?** - License, Go version, platform support


### Example of package 
Purpose: A minimal Go executable that prints a greeting and then a farewell.  
Structure: 
package main in all three files makes them part of the same executable program.    
hey.go defines hey() which prints "Hi ".  
bye.go defines bye() which prints "goodbye".  
main.go defines main(), the entry point, which calls hey() and then bye().  
Imports: Both hey.go and bye.go import fmt to print to the console.  
Execution: Running the program prints two lines in order:  
Hi  
goodbye  
You can run it from the Package folder with:  
go run . 

### Short answer
1. Use package main when you want to build and run an executable. It must contain a func main().  
2. Use a different package name when you’re creating a reusable library to be imported by other code.  

### When to use package main
1. You are building a command you can run (e.g., go run . or go build → binary).  
2. The package must define func main().
3. All files in that directory must also declare package main.  

### When to use another package name
1. You’re writing reusable code (a library) to be imported elsewhere.  
2. Example: package mathutil in a folder, then another main package imports it with import "module/mathutil" and calls its exported functions.  
3. In the same directory, all files must have the same package name; you can’t mix main and non-main together.  

### Typical structure
1. Executables: put package main under cmd/appname/.  
2. Libraries: put reusable packages under directories like pkg/, internal/, or any folder named for the package (the folder name usually matches the package name).    

### In your current code
1. It’s package main because you’re building a small executable that runs main() and prints messages.    
2. If you renamed it (e.g., package greetings), it would no longer run directly; you’d need a separate main package that imports greetings and calls its functions.  