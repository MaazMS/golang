# Identifiers and Keywords

## Identifiers

**Identifiers** are user-defined names for program components in Go. They serve as labels for various elements in your code.

### What can be an identifier?
- Variable names
- Function names  
- Constants
- Statement labels (used with `goto`)
- Package names
- Type names (including structures)
- Method names
- Field names

### Identifier Rules
- Must start with a letter or underscore (`_`)
- Can contain letters, digits, and underscores
- Cannot start with a digit
- Case-sensitive (e.g., `name` and `Name` are different)
- Cannot be a Go keyword

### Exported vs Unexported Identifiers
- Identifiers starting with an uppercase letter are exported (visible to other packages): `Println`, `User`, `HTTPServer`.
- Identifiers starting with a lowercase letter are unexported (package-private): `println`, `user`, `server`.

### Unicode Support
- Identifiers may use Unicode letters and digits (e.g., `π`, `用户数`).
- Prefer ASCII for public APIs to maximize portability and readability.

### Examples
```go
var userName string        // Valid identifier
var _private int          // Valid identifier  
var 2invalid int         // Invalid - starts with digit
var user-name string      // Invalid - contains hyphen
```

### The Blank Identifier `_`
- `_` discards values you don't need and suppresses “declared and not used” errors.
- Common uses:
```go
v, _ := someFunc()           // ignore second return value
for _, item := range xs { _ = item } // ignore index
var _ = someExpensiveInit    // keep side effects, ignore value
```
- Import for side effects only:
```go
import _ "github.com/lib/pq" // register driver without direct reference
```

## Keywords

**Keywords** are reserved words in the Go programming language that have special meanings and cannot be used as identifiers.

### Key Points about Keywords
1. Keywords are reserved words with predefined meanings in the language
2. They cannot be used as identifiers (variable names, function names, etc.)
3. Using keywords as identifiers will result in a compile-time error
4. Go has **25 keywords** in total

### Complete List of Go Keywords (25)

| Keyword | Purpose |
|---------|---------|
| `break` | Exit from a loop or switch |
| `case` | Used in switch statements |
| `chan` | Channel type for communication |
| `const` | Declare constants |
| `continue` | Skip to next iteration of loop |
| `default` | Default case in switch |
| `defer` | Execute function when surrounding function returns |
| `else` | Alternative branch in if statement |
| `fallthrough` | Continue to next case in switch |
| `for` | Loop construct |
| `func` | Function declaration |
| `go` | Start a goroutine |
| `goto` | Jump to a label |
| `if` | Conditional statement |
| `import` | Import packages |
| `interface` | Define interface type |
| `map` | Map type (key-value pairs) |
| `package` | Package declaration |
| `range` | Iterate over arrays, slices, maps |
| `return` | Return from function |
| `select` | Choose from multiple channel operations |
| `struct` | Define structure type |
| `switch` | Multi-way conditional |
| `type` | Define new types |
| `var` | Declare variables |

### Example of Keyword Usage Error
```go
// This will cause a compile error
var var int = 10  // Error: "var" is a keyword

// Correct usage
var variable int = 10  // Valid
```

![](https://media.geeksforgeeks.org/wp-content/uploads/20191118101819/Golang-Keywords.png)      

## Predeclared Identifiers (Not Keywords)
These names exist in the universe scope. They are not reserved words, but shadowing them is discouraged for readability.

### Types
`bool, byte, complex64, complex128, error, float32, float64, int, int8, int16, int32, int64, rune, string, uint, uint8, uint16, uint32, uint64, uintptr`

### Constants
`true, false, iota, nil`

### Built-in Functions
`append, cap, close, complex, copy, delete, imag, len, make, new, panic, print, println, real, recover`

### Notes
- Avoid names like `len`, `error`, or `string` for your own identifiers to prevent confusion.
- Prefer `fmt.Print*` over `print`/`println` in production code.