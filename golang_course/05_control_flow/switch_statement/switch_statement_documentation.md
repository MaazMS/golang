# Go Switch Statement - Complete Guide

## Overview

The `switch` statement in Go provides an efficient way to transfer control of code execution based on the value of an expression. It is more powerful and flexible than traditional switch statements in other languages, offering multiple forms and advanced features.

## Key Features

1. **Expression-based switching**: Switch on any expression, not just integers
2. **Type switching**: Switch on the type of an interface value
3. **No fall-through**: Cases don't automatically fall through to the next case
4. **Multiple values**: Each case can handle multiple values
5. **Default case**: Optional default case for unmatched values

## Basic Syntax

```go
switch expression {
case value1, value2:
    // statements
case value3:
    // statements
case value4, value5, value6:
    // statements
default:
    // statements (optional)
}
```

## Basic Examples

### Example 1: Simple Value Switching

```go
package main

import "fmt"

func main() {
    day := "Monday"
    
    switch day {
    case "Monday":
        fmt.Println("Start of the work week")
    case "Tuesday", "Wednesday", "Thursday":
        fmt.Println("Mid week")
    case "Friday":
        fmt.Println("TGIF!")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Invalid day")
    }
}
```

### Example 2: Expression-based Switch

```go
package main

import "fmt"

func main() {
    score := 85
    
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80:
        fmt.Println("Grade: B")
    case score >= 70:
        fmt.Println("Grade: C")
    case score >= 60:
        fmt.Println("Grade: D")
    default:
        fmt.Println("Grade: F")
    }
}
```

### Example 3: Switch with Multiple Values

```go
package main

import "fmt"

func main() {
    month := "March"
    
    switch month {
    case "December", "January", "February":
        fmt.Println("Winter")
    case "March", "April", "May":
        fmt.Println("Spring")
    case "June", "July", "August":
        fmt.Println("Summer")
    case "September", "October", "November":
        fmt.Println("Autumn")
    default:
        fmt.Println("Invalid month")
    }
}
```

## Advanced Examples

### Example 4: Type Switch

```go
package main

import "fmt"

func processValue(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    case float64:
        fmt.Printf("Float: %.2f\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    processValue(42)
    processValue("Hello")
    processValue(true)
    processValue(3.14)
    processValue([]int{1, 2, 3})
}
```

### Example 5: Switch with Initialization

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    switch hour := time.Now().Hour(); {
    case hour < 6:
        fmt.Println("Good night")
    case hour < 12:
        fmt.Println("Good morning")
    case hour < 18:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
    }
}
```

### Example 6: Complex Expression Switch

```go
package main

import "fmt"

func getGrade(score int) string {
    switch {
    case score >= 95:
        return "A+"
    case score >= 90:
        return "A"
    case score >= 85:
        return "B+"
    case score >= 80:
        return "B"
    case score >= 75:
        return "C+"
    case score >= 70:
        return "C"
    case score >= 65:
        return "D+"
    case score >= 60:
        return "D"
    default:
        return "F"
    }
}

func main() {
    scores := []int{98, 87, 76, 65, 45}
    for _, score := range scores {
        fmt.Printf("Score: %d, Grade: %s\n", score, getGrade(score))
    }
}
```

## Type Switch Deep Dive

### Example 7: Advanced Type Switching

```go
package main

import (
    "fmt"
    "reflect"
)

func analyzeType(x interface{}) {
    switch v := x.(type) {
    case nil:
        fmt.Println("Type: nil")
    case int:
        fmt.Printf("Type: int, Value: %d\n", v)
    case int8:
        fmt.Printf("Type: int8, Value: %d\n", v)
    case int16:
        fmt.Printf("Type: int16, Value: %d\n", v)
    case int32:
        fmt.Printf("Type: int32, Value: %d\n", v)
    case int64:
        fmt.Printf("Type: int64, Value: %d\n", v)
    case uint:
        fmt.Printf("Type: uint, Value: %d\n", v)
    case float32:
        fmt.Printf("Type: float32, Value: %.2f\n", v)
    case float64:
        fmt.Printf("Type: float64, Value: %.2f\n", v)
    case string:
        fmt.Printf("Type: string, Value: %s\n", v)
    case bool:
        fmt.Printf("Type: bool, Value: %t\n", v)
    case []int:
        fmt.Printf("Type: []int, Value: %v\n", v)
    case map[string]int:
        fmt.Printf("Type: map[string]int, Value: %v\n", v)
    default:
        fmt.Printf("Type: %T, Value: %v\n", v, v)
    }
}

func main() {
    analyzeType(42)
    analyzeType("Hello")
    analyzeType(3.14)
    analyzeType(true)
    analyzeType([]int{1, 2, 3})
    analyzeType(map[string]int{"a": 1, "b": 2})
    analyzeType(nil)
}
```

## Practical Use Cases

### Example 8: HTTP Status Code Handler

```go
package main

import "fmt"

func handleHTTPStatus(code int) string {
    switch code {
    case 200:
        return "OK"
    case 201:
        return "Created"
    case 204:
        return "No Content"
    case 400:
        return "Bad Request"
    case 401:
        return "Unauthorized"
    case 403:
        return "Forbidden"
    case 404:
        return "Not Found"
    case 500:
        return "Internal Server Error"
    case 502:
        return "Bad Gateway"
    case 503:
        return "Service Unavailable"
    default:
        if code >= 200 && code < 300 {
            return "Success"
        } else if code >= 300 && code < 400 {
            return "Redirection"
        } else if code >= 400 && code < 500 {
            return "Client Error"
        } else if code >= 500 {
            return "Server Error"
        }
        return "Unknown Status"
    }
}

func main() {
    statusCodes := []int{200, 404, 500, 201, 999}
    for _, code := range statusCodes {
        fmt.Printf("Status %d: %s\n", code, handleHTTPStatus(code))
    }
}
```

### Example 9: File Extension Handler

```go
package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func getFileType(filename string) string {
    ext := strings.ToLower(filepath.Ext(filename))
    
    switch ext {
    case ".go":
        return "Go source file"
    case ".py":
        return "Python script"
    case ".js":
        return "JavaScript file"
    case ".html", ".htm":
        return "HTML document"
    case ".css":
        return "CSS stylesheet"
    case ".json":
        return "JSON data"
    case ".xml":
        return "XML document"
    case ".txt":
        return "Text file"
    case ".md":
        return "Markdown document"
    case ".pdf":
        return "PDF document"
    case ".jpg", ".jpeg", ".png", ".gif":
        return "Image file"
    case ".mp4", ".avi", ".mov":
        return "Video file"
    case ".mp3", ".wav", ".flac":
        return "Audio file"
    default:
        return "Unknown file type"
    }
}

func main() {
    files := []string{
        "main.go",
        "style.css",
        "index.html",
        "data.json",
        "README.md",
        "photo.jpg",
        "video.mp4",
        "unknown.xyz",
    }
    
    for _, file := range files {
        fmt.Printf("%s -> %s\n", file, getFileType(file))
    }
}
```

## Important Rules and Best Practices

### 1. No Fall-through
Unlike C/C++, Go switch statements don't fall through to the next case automatically:

```go
// ❌ This won't work as expected in Go
switch x {
case 1:
    fmt.Println("One")
    // No fall-through to case 2
case 2:
    fmt.Println("Two")
}
```

### 2. Multiple Values in Cases
```go
switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
}
```

### 3. Expression Switch
When no expression is provided, it's equivalent to `switch true`:

```go
score := 85
switch { // equivalent to switch true
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
}
```

### 4. Type Assertion in Type Switch
```go
var i interface{} = "hello"

switch v := i.(type) {
case string:
    fmt.Printf("String: %s\n", v)
case int:
    fmt.Printf("Integer: %d\n", v)
}
```

## Performance Considerations

1. **Switch vs if-else**: Switch is generally more efficient for multiple conditions
2. **Type switches**: More efficient than type assertions with if statements
3. **Expression evaluation**: Expression is evaluated only once
4. **Compile-time optimization**: Go compiler optimizes switch statements

## Common Patterns

### Pattern 1: State Machine
```go
type State int

const (
    Idle State = iota
    Running
    Paused
    Stopped
)

func (s State) String() string {
    switch s {
    case Idle:
        return "Idle"
    case Running:
        return "Running"
    case Paused:
        return "Paused"
    case Stopped:
        return "Stopped"
    default:
        return "Unknown"
    }
}
```

### Pattern 2: Command Pattern
```go
func handleCommand(cmd string) {
    switch cmd {
    case "start":
        startService()
    case "stop":
        stopService()
    case "restart":
        restartService()
    case "status":
        showStatus()
    default:
        fmt.Println("Unknown command")
    }
}
```

## Error Handling
```go
func processError(err error) {
    switch e := err.(type) {
    case *CustomError:
        fmt.Printf("Custom error: %s\n", e.Message)
    case *ValidationError:
        fmt.Printf("Validation error: %s\n", e.Field)
    case nil:
        fmt.Println("No error")
    default:
        fmt.Printf("Unknown error: %v\n", err)
    }
}
```

## Conclusion

The `switch` statement in Go is a powerful and flexible control structure that goes beyond traditional switch statements. It supports value switching, type switching, and expression-based switching, making it suitable for a wide range of programming scenarios. Understanding its various forms and best practices is essential for writing clean, efficient Go code.
