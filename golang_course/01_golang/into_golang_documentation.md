# Introduction to Go (Golang)

## Who invented Go?

Go is a statically typed, compiled programming language designed at Google in 2007 and developed by:

- **Robert Griesemer** (Hotspot, JVM)
- **Rob Pike** (Unix, UTF-8) 
- **Ken Thompson** (B, C, Unix, UTF-8)
- And a few other engineers

**Timeline:**
- **2007**: Development started at Google
- **2009**: First public release
- **2012**: Released as an open-source programming language

### The Problem Before Go

At Google (and other big tech companies), they had to write large-scale systems that serve millions of users:

- **C++**: Fast but complicated and slow to compile
- **Python**: Easy to use but too slow for huge systems
- **Java**: Complex type system

So they wanted a new language that combines the best of both worlds.

### Why Go Was Created (The Goals)

The creators of Go wanted a language that was:

1. **Fast like C++ ⚡**
   - Compiles quickly
   - Runs efficiently

2. **Simple like Python 🐍**
   - Easy to learn
   - Clean syntax

3. **Great for concurrency 🔄**
   - Concurrency = running many things at the same time (like handling thousands of users at once)
   - Go has goroutines → lightweight threads for multitasking

4. **Safe & maintainable 🛠️**
   - Designed to make code easier to read, review, and manage at scale

**Reference**: [Why did you create a new language?](https://golang.org/doc/faq#creating_a_new_language)



## Go Documentation

- [The Go Programming Language Specification](https://golang.org/ref/spec)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)

## golang.org vs pkg.go.dev

**golang.org** (Official Go website):
- Standard library documentation
- Source code
- Language specification

**pkg.go.dev** (Package discovery site):
- Standard library AND third-party packages
- Module information
- Version history
- Dependencies

## Go Playground

The Go Playground is an online tool that allows you to:
- **Format code** - Automatically format Go code
- **Share code** - Share code snippets with others
- **Run code** - Execute Go code in the browser
- **Test code** - Experiment with Go features

**Access**: [Go Playground](https://play.golang.org/)

## Go Features

1. **Performant** - Fast execution and compilation
2. **Multi-core support** - Designed for modern multi-core processors
3. **Concurrency** - Built-in support for concurrent programming with goroutines
4. **Compiled** - Compiles to native machine code
5. **Network programming** - Excellent support for network applications
6. **Clean syntax** - Simple and readable code
7. **Powerful standard library** - Rich set of built-in packages
8. **Garbage collected** - Automatic memory management
9. **Portable** - Compiles on many operating systems
10. **Open source** - Free and open-source software

## Go vs Other Languages

**Comparison highlights:**
- **vs C++**: Simpler syntax, faster compilation, garbage collection
- **vs Java**: No JVM overhead, faster startup, simpler deployment
- **vs Python**: Much faster execution, static typing, better concurrency
- **vs Node.js**: Better performance, static typing, built-in concurrency

**Reference**: [Go vs Other Languages](https://talks.golang.org/2014/gocon-tokyo.slide#1)

## Companies Using Go

Many major companies use Go in production:
- **Google** - Internal systems and services
- **Docker** - Containerization platform
- **Kubernetes** - Container orchestration
- **Uber** - Microservices and backend systems
- **Netflix** - Streaming services
- **Dropbox** - File storage and sync
- **SoundCloud** - Audio streaming platform

**Complete list**: [GoUsers](https://github.com/golang/go/wiki/GoUsers)

## What is a Package?

A package is **pre-written code that you can use in your project**. It's imported at the beginning of your source code.

### Key Concepts

**package main**:
- Every Go application is structured as packages
- `main` is a special package
- It's the entry point of the application

**import**:
- Used to add additional libraries
- `fmt` library is used to format strings and print output

**func main()**:
- It's the entry point of the application
- Go automatically calls this function when the program starts

### Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Key Points to Remember

- **Go was created at Google** by Robert Griesemer, Rob Pike, and Ken Thompson
- **Go combines** the speed of C++ with the simplicity of Python
- **Go excels at concurrency** with goroutines and channels
- **Go is compiled** but has fast compilation times
- **Go has a rich standard library** and excellent tooling
- **Go is widely used** by major tech companies worldwide


