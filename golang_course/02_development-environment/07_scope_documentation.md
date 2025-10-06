## What is scope?
Scope is the region of source code where an identifier (variable, constant, type, function, method, label, or package name) is visible and can be referenced. Go uses lexical (static) scoping and block structure.

At a high level:
- Identifiers declared at the package level are visible to all files in the same package.
- Imports are file-scoped: each file must import the packages it uses.
- Identifiers declared inside a block (e.g., a function, `if`, `for`, or `{ ... }`) are only visible within that block and its nested blocks.
- Whether a name is exported to other packages depends on capitalization: an identifier starting with an uppercase letter is exported; lowercase is unexported.

---

## Package scope
Package scope refers to identifiers declared outside of any function in a `.go` file that belongs to a package. Such identifiers are visible to all other files in the same package.

Example:
```go
package main

import "fmt"

// Package-scope declarations
const appName = "demo"
var version = "v1"

func greet() { // visible within package main
    fmt.Println("Hello from", appName, version)
}

func main() {
    greet() // OK: greet, appName, version are in package scope
}
```

Notes:
- Exported identifiers (starting with an uppercase letter, e.g., `Greet`, `AppName`) are visible to other packages that import this package.
- Unexported identifiers (lowercase) are only visible within the same package.

---

## File scope (imports and file-level items)
Imports in Go are file-scoped. If a file references a package, that file must import it—even if other files in the same package already import it.

Example (two files in the same package):
```go
// file: a.go
package main

import "fmt" // required in this file because fmt is used here

func printA() {
    fmt.Println("A")
}
```

```go
// file: b.go
package main

// This file does not use fmt, so it does not need to import fmt.
func printB() {}
```

Key points:
- “We import packages into a file, not into a whole package.” Each file manages its own imports.
- Unused imports cause a compile error. Remove or use them.

---

## Block and function scope
Identifiers declared inside a function or an inner block are only visible within that block (and nested blocks).

Example:
```go
package main

import "fmt"

func main() {
    var a = 1 // function scope (visible in all of main)
    if a == 1 {
        b := 2 // block scope (visible only inside this if block)
        fmt.Println(a, b)
    }
    // fmt.Println(b) // compile error: b is not in scope here
}
```

Short variable declarations (`:=`) are allowed only inside functions. At package scope, use `var`, `const`, `type`, or `func`.

---

## Name collisions and shadowing
You cannot declare two identifiers with the same name in the same scope. Go also allows shadowing, where an inner declaration uses the same name as an outer one; the inner name hides the outer within the inner scope. Use shadowing judiciously as it can reduce clarity.

Example (collision):
```go
package p

var x = 1
var x = 2 // compile error: x redeclared in this block
```

Example (shadowing):
```go
package main

import "fmt"

var value = 10 // package scope

func main() {
    value := 5 // shadows package-scope value inside main
    fmt.Println(value) // prints 5
}
```

---

## Exported vs. unexported identifiers
- An identifier whose name starts with an uppercase letter is exported and can be used by other packages that import this package.
- An identifier starting with a lowercase letter is unexported and is visible only within the same package.

Example:
```go
package lib

// Exported type
type Counter struct{ N int }

// Unexported function (package-only)
func increment(c *Counter) { c.N++ }

// Exported function
func NewCounter() *Counter { return &Counter{} }
```

---

## Import aliases and special imports
You can rename an import with an alias, import for side effects only, or (rarely) import into the current namespace.

Examples:
```go
import f "fmt"       // alias: use f.Println, f.Errorf, etc.
import _ "net/http/pprof" // blank import: for side effects (init functions)
import . "fmt"       // dot import: brings names into the file’s scope (discouraged)
```

Guidelines:
- Prefer normal imports; use aliases when avoiding name conflicts or improving clarity.
- Avoid dot imports in production code—they can obscure where names come from.
- Blank imports are acceptable when you deliberately need package side effects.

---

## Common pitfalls related to scope
- Unused identifiers and imports cause compile errors; remove or use them.
- Assuming an import in one file applies to other files in the same package—it does not.
- Accidentally shadowing variables, leading to subtle bugs. Consider naming to avoid shadowing, or enable linters that warn about it.

---

## Additional practical examples

Multiple files sharing package-scope identifiers:
```go
// file: config.go
package main

var debug = true // visible to all files in package main
```

```go
// file: run.go
package main

import "fmt"

func run() {
    if debug {
        fmt.Println("debug mode")
    }
}
```

Function-local constants and variables are not visible outside:
```go
func compute() int {
    const factor = 2
    x := 10
    return factor * x
}
// fmt.Println(factor) // compile error: factor not in scope here
```

---

## FAQ
### Can I declare two functions with the same name in one package?
No. Two top-level declarations in the same package cannot have the same name.

### Do I need to import `fmt` in every file that calls `fmt.Println`?
Yes. Imports are file-scoped; each file that uses `fmt` must import it.

### How do I rename an imported package?
Use an alias:
```go
import f "fmt"

func main() {
    f.Println("hello")
}
```