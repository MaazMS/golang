# IOTA in Go - Complete Beginner's Guide

## What is IOTA?

`iota` is a special predeclared identifier in Go that represents successive untyped integer constants. It's like an automatic counter that starts at 0 and increments by 1 for each constant in a `const` declaration block.

Think of `iota` as a smart counter that automatically assigns sequential numbers to your constants, making it perfect for creating enumerated values, flags, and other sequential constants.

## How IOTA Works

### Basic IOTA Usage

```go
package main

import "fmt"

func main() {
    const (
        a = iota  // 0
        b = iota  // 1
        c = iota  // 2
        d = iota  // 3
    )
    
    fmt.Println("a =", a)  // Output: a = 0
    fmt.Println("b =", b)  // Output: b = 1
    fmt.Println("c =", c)  // Output: c = 2
    fmt.Println("d =", d)  // Output: d = 3
}
```

### IOTA Auto-Increment (Simplified Syntax)

```go
package main

import "fmt"

func main() {
    const (
        a = iota  // 0
        b         // 1 (automatically iota)
        c         // 2 (automatically iota)
        d         // 3 (automatically iota)
    )
    
    fmt.Println("a =", a)  // Output: a = 0
    fmt.Println("b =", b)  // Output: b = 1
    fmt.Println("c =", c)  // Output: c = 2
    fmt.Println("d =", d)  // Output: d = 3
}
```

## IOTA Resets

IOTA resets to 0 whenever a new `const` declaration starts:

```go
package main

import "fmt"

func main() {
    const (
        first  = iota  // 0
        second = iota  // 1
        third  = iota  // 2
    )
    
    const (
        alpha = iota  // 0 (resets!)
        beta  = iota  // 1
        gamma = iota  // 2
    )
    
    fmt.Println("First group:")
    fmt.Println("first =", first)   // 0
    fmt.Println("second =", second) // 1
    fmt.Println("third =", third)   // 2
    
    fmt.Println("\nSecond group:")
    fmt.Println("alpha =", alpha)  // 0
    fmt.Println("beta =", beta)     // 1
    fmt.Println("gamma =", gamma)   // 2
}
```

## Practical Examples

### Example 1: Days of the Week

```go
package main

import "fmt"

func main() {
    const (
        Sunday    = iota  // 0
        Monday    = iota  // 1
        Tuesday   = iota  // 2
        Wednesday = iota  // 3
        Thursday  = iota  // 4
        Friday    = iota  // 5
        Saturday  = iota  // 6
    )
    
    fmt.Println("Days of the week:")
    fmt.Printf("Sunday = %d\n", Sunday)
    fmt.Printf("Monday = %d\n", Monday)
    fmt.Printf("Tuesday = %d\n", Tuesday)
    fmt.Printf("Wednesday = %d\n", Wednesday)
    fmt.Printf("Thursday = %d\n", Thursday)
    fmt.Printf("Friday = %d\n", Friday)
    fmt.Printf("Saturday = %d\n", Saturday)
}
```

### Example 2: User Roles

```go
package main

import "fmt"

func main() {
    const (
        Guest = iota  // 0
        User          // 1
        Moderator     // 2
        Admin         // 3
        SuperAdmin    // 4
    )
    
    userRole := User
    
    switch userRole {
    case Guest:
        fmt.Println("Welcome, guest!")
    case User:
        fmt.Println("Welcome back, user!")
    case Moderator:
        fmt.Println("Moderator privileges active")
    case Admin:
        fmt.Println("Admin panel access granted")
    case SuperAdmin:
        fmt.Println("Full system access")
    }
}
```

### Example 3: File Permissions

```go
package main

import "fmt"

func main() {
    const (
        Read    = 1 << iota  // 1 (binary: 001)
        Write                // 2 (binary: 010)
        Execute              // 4 (binary: 100)
    )
    
    // Combining permissions
    readWrite := Read | Write  // 3 (binary: 011)
    allPerms := Read | Write | Execute  // 7 (binary: 111)
    
    fmt.Printf("Read permission: %d (binary: %03b)\n", Read, Read)
    fmt.Printf("Write permission: %d (binary: %03b)\n", Write, Write)
    fmt.Printf("Execute permission: %d (binary: %03b)\n", Execute, Execute)
    fmt.Printf("Read + Write: %d (binary: %03b)\n", readWrite, readWrite)
    fmt.Printf("All permissions: %d (binary: %03b)\n", allPerms, allPerms)
}
```

## Advanced IOTA Patterns

### 1. Starting from a Different Value

```go
package main

import "fmt"

func main() {
    const (
        _ = iota        // Skip 0
        January         // 1
        February        // 2
        March           // 3
        April           // 4
        May             // 5
        June            // 6
    )
    
    fmt.Println("Months (1-indexed):")
    fmt.Printf("January = %d\n", January)
    fmt.Printf("February = %d\n", February)
    fmt.Printf("March = %d\n", March)
}
```

### 2. Skipping Values

```go
package main

import "fmt"

func main() {
    const (
        Small  = iota  // 0
        Medium = iota  // 1
        _              // Skip 2
        _              // Skip 3
        Large  = iota  // 4
        XLarge = iota  // 5
    )
    
    fmt.Printf("Small = %d\n", Small)
    fmt.Printf("Medium = %d\n", Medium)
    fmt.Printf("Large = %d\n", Large)
    fmt.Printf("XLarge = %d\n", XLarge)
}
```

### 3. Custom Values with IOTA

```go
package main

import "fmt"

func main() {
    const (
        Apple  = (iota + 1) * 10  // 10
        Orange                    // 20
        Banana                    // 30
        Grape                     // 40
    )
    
    fmt.Printf("Apple = %d\n", Apple)   // 10
    fmt.Printf("Orange = %d\n", Orange) // 20
    fmt.Printf("Banana = %d\n", Banana) // 30
    fmt.Printf("Grape = %d\n", Grape)   // 40
}
```

### 4. String Constants with IOTA

```go
package main

import "fmt"

func main() {
    const (
        StatusPending = "pending"
        StatusActive  = "active"
        StatusPaused  = "paused"
        StatusStopped = "stopped"
    )
    
    // Using iota for indexing
    const (
        _ = iota
        Pending
        Active
        Paused
        Stopped
    )
    
    statuses := []string{"", StatusPending, StatusActive, StatusPaused, StatusStopped}
    
    fmt.Printf("Status %d: %s\n", Pending, statuses[Pending])
    fmt.Printf("Status %d: %s\n", Active, statuses[Active])
    fmt.Printf("Status %d: %s\n", Paused, statuses[Paused])
    fmt.Printf("Status %d: %s\n", Stopped, statuses[Stopped])
}
```

## Real-World Examples

### Example 1: HTTP Status Codes

```go
package main

import "fmt"

func main() {
    const (
        StatusOK                  = 200
        StatusCreated             = 201
        StatusBadRequest          = 400
        StatusUnauthorized        = 401
        StatusForbidden           = 403
        StatusNotFound            = 404
        StatusInternalServerError = 500
    )
    
    // Using iota for status categories
    const (
        _ = iota
        SuccessCategory  // 1
        ClientErrorCategory  // 2
        ServerErrorCategory  // 3
    )
    
    responseCode := 404
    
    if responseCode >= 200 && responseCode < 300 {
        fmt.Printf("Success category: %d\n", SuccessCategory)
    } else if responseCode >= 400 && responseCode < 500 {
        fmt.Printf("Client error category: %d\n", ClientErrorCategory)
    } else if responseCode >= 500 {
        fmt.Printf("Server error category: %d\n", ServerErrorCategory)
    }
}
```

### Example 2: Log Levels

```go
package main

import "fmt"

func main() {
    const (
        LogDebug = iota  // 0
        LogInfo          // 1
        LogWarning       // 2
        LogError         // 3
        LogFatal         // 4
    )
    
    currentLogLevel := LogInfo
    
    if currentLogLevel >= LogInfo {
        fmt.Println("This message will be logged (Info level)")
    }
    
    if currentLogLevel >= LogWarning {
        fmt.Println("This warning will be logged")
    } else {
        fmt.Println("Warning level not reached")
    }
}
```

### Example 3: Database Operations

```go
package main

import "fmt"

func main() {
    const (
        OpCreate = iota  // 0
        OpRead          // 1
        OpUpdate        // 2
        OpDelete        // 3
    )
    
    operation := OpRead
    
    switch operation {
    case OpCreate:
        fmt.Println("Creating new record...")
    case OpRead:
        fmt.Println("Reading record...")
    case OpUpdate:
        fmt.Println("Updating record...")
    case OpDelete:
        fmt.Println("Deleting record...")
    default:
        fmt.Println("Unknown operation")
    }
}
```

## IOTA with Bit Flags

### Creating Bit Flags

```go
package main

import "fmt"

func main() {
    const (
        FlagRead    = 1 << iota  // 1 (0001)
        FlagWrite                // 2 (0010)
        FlagExecute              // 4 (0100)
        FlagDelete               // 8 (1000)
    )
    
    // Combining flags
    readWrite := FlagRead | FlagWrite  // 3 (0011)
    allFlags := FlagRead | FlagWrite | FlagExecute | FlagDelete  // 15 (1111)
    
    fmt.Printf("Read flag: %d (binary: %04b)\n", FlagRead, FlagRead)
    fmt.Printf("Write flag: %d (binary: %04b)\n", FlagWrite, FlagWrite)
    fmt.Printf("Read+Write: %d (binary: %04b)\n", readWrite, readWrite)
    fmt.Printf("All flags: %d (binary: %04b)\n", allFlags, allFlags)
    
    // Checking flags
    if allFlags&FlagRead != 0 {
        fmt.Println("Read permission is set")
    }
    if allFlags&FlagWrite != 0 {
        fmt.Println("Write permission is set")
    }
}
```

## Common Mistakes to Avoid

### 1. Forgetting IOTA Resets

```go
// ❌ Wrong expectation
const (
    a = iota  // 0
    b = iota  // 1
)
const (
    c = iota  // 0 (not 2!)
    d = iota  // 1 (not 3!)
)
```

### 2. Mixing IOTA with Regular Constants

```go
// ❌ Confusing
const (
    a = iota  // 0
    b = 100   // 100
    c = iota  // 2 (not 1!)
)

// ✅ Clear
const (
    a = iota  // 0
    b = iota  // 1
    c = iota  // 2
)
```

### 3. Not Understanding Bit Shifts

```go
// ❌ Wrong expectation
const (
    a = 1 << iota  // 1
    b = 1 << iota  // 2
    c = 1 << iota  // 4
)

// The values are correct, but make sure you understand bit shifting!
```

## Best Practices

### 1. Use Descriptive Names

```go
// ❌ Bad
const (
    a = iota
    b = iota
    c = iota
)

// ✅ Good
const (
    StatusPending = iota
    StatusActive
    StatusInactive
)
```

### 2. Group Related Constants

```go
// ✅ Good - Group related constants
const (
    // User roles
    RoleGuest = iota
    RoleUser
    RoleModerator
    RoleAdmin
)

const (
    // File permissions
    PermissionRead = 1 << iota
    PermissionWrite
    PermissionExecute
)
```

### 3. Document Your Constants

```go
// ✅ Good - Documented constants
const (
    // HTTP status codes
    StatusOK = 200
    StatusNotFound = 404
    StatusInternalError = 500
)

const (
    // User permission levels
    PermissionGuest = iota  // 0 - No special permissions
    PermissionUser         // 1 - Basic user permissions
    PermissionModerator    // 2 - Can moderate content
    PermissionAdmin        // 3 - Full administrative access
)
```

## Summary

- **IOTA** is an automatic counter that starts at 0 and increments by 1
- IOTA resets to 0 with each new `const` declaration
- Use IOTA for creating enumerated values and sequential constants
- IOTA is perfect for bit flags using bit shifting (`1 << iota`)
- Group related constants together for better organization
- Use descriptive names for your constants
- IOTA makes your code more maintainable and less error-prone

IOTA is a powerful feature that helps you create clean, maintainable code with automatically numbered constants. It's especially useful for enums, flags, and any situation where you need sequential values.
