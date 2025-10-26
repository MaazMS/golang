# Go Maps - Complete Guide

## Overview

A map in Go is a built-in data type that stores key-value pairs. It provides an efficient way to associate values with unique keys, making it ideal for lookups, caching, and data organization. Maps are reference types and are more flexible than arrays or slices for certain use cases.

## Key Features

1. **Key-value storage**: Maps store data as key-value pairs
2. **Fast lookups**: O(1) average time complexity for operations
3. **Dynamic size**: Maps can grow and shrink as needed
4. **Reference type**: Maps are passed by reference
5. **Type safety**: Keys and values must be of specified types

## Why Use Maps?

### Problem with Slices/Arrays
```go
// ❌ Inefficient: O(n) time complexity
func findPrice(slice []string, item string) float64 {
    for i, s := range slice {
        if s == item {
            return prices[i] // Need separate slice for prices
        }
    }
    return 0.0
}
```

### Solution with Maps
```go
// ✅ Efficient: O(1) time complexity
func findPrice(menu map[string]float64, item string) float64 {
    return menu[item] // Direct lookup
}
```

## Basic Syntax

### Declaration and Initialization

```go
// Method 1: Declaration then initialization
var menu map[string]float64
menu = map[string]float64{
    "eggs":  70.0,
    "milk":  30.0,
    "water": 15.0,
}

// Method 2: Short declaration (recommended)
menu := map[string]float64{
    "eggs":  70.0,
    "milk":  30.0,
    "water": 15.0,
}

// Method 3: Empty map initialization
menu := map[string]float64{}
```

## Basic Examples

### Example 1: Simple Map Operations

```go
package main

import "fmt"

func main() {
    // Create a map
    ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Carol": 35,
    }

    // Access values
    fmt.Println("Alice's age:", ages["Alice"])

    // Add new key-value pair
    ages["David"] = 28

    // Update existing value
    ages["Alice"] = 26

    // Check if key exists
    if age, exists := ages["Eve"]; exists {
        fmt.Println("Eve's age:", age)
    } else {
        fmt.Println("Eve not found")
    }

    // Delete a key
    delete(ages, "Bob")

    // Iterate over map
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}
```

### Example 2: Map with Different Key Types

```go
package main

import "fmt"

func main() {
    // String keys
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }

    // Integer keys
    scores := map[int]string{
        90: "A",
        80: "B",
        70: "C",
        60: "D",
    }

    // Boolean keys
    settings := map[bool]string{
        true:  "Enabled",
        false: "Disabled",
    }

    fmt.Println("Red color:", colors["red"])
    fmt.Println("Score 90:", scores[90])
    fmt.Println("True setting:", settings[true])
}
```

### Example 3: Nested Maps

```go
package main

import "fmt"

func main() {
    // Map of maps
    students := map[string]map[string]int{
        "Alice": {
            "math":    95,
            "science": 87,
            "english": 92,
        },
        "Bob": {
            "math":    78,
            "science": 85,
            "english": 88,
        },
    }

    // Access nested values
    fmt.Println("Alice's math score:", students["Alice"]["math"])

    // Add new student
    students["Carol"] = map[string]int{
        "math":    88,
        "science": 91,
        "english": 85,
    }

    // Iterate over nested map
    for name, subjects := range students {
        fmt.Printf("\n%s's scores:\n", name)
        for subject, score := range subjects {
            fmt.Printf("  %s: %d\n", subject, score)
        }
    }
}
```

## Advanced Examples

### Example 4: Map Operations and Utilities

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    inventory := map[string]int{
        "apples":  50,
        "bananas": 30,
        "oranges": 25,
        "grapes":  40,
    }

    // Check if map is empty
    if len(inventory) == 0 {
        fmt.Println("Inventory is empty")
    } else {
        fmt.Printf("Inventory has %d items\n", len(inventory))
    }

    // Get all keys
    keys := make([]string, 0, len(inventory))
    for k := range inventory {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    fmt.Println("Sorted keys:", keys)

    // Get all values
    values := make([]int, 0, len(inventory))
    for _, v := range inventory {
        values = append(values, v)
    }
    fmt.Println("Values:", values)

    // Find items with quantity > 30
    highStock := make(map[string]int)
    for item, qty := range inventory {
        if qty > 30 {
            highStock[item] = qty
        }
    }
    fmt.Println("High stock items:", highStock)
}
```

### Example 5: Map as Cache

```go
package main

import (
    "fmt"
    "time"
)

type Cache struct {
    data map[string]interface{}
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

func (c *Cache) Set(key string, value interface{}) {
    c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) Delete(key string) {
    delete(c.data, key)
}

func (c *Cache) Clear() {
    c.data = make(map[string]interface{})
}

func main() {
    cache := NewCache()

    // Store some data
    cache.Set("user:1", map[string]string{
        "name":  "Alice",
        "email": "alice@example.com",
    })

    cache.Set("config:theme", "dark")
    cache.Set("config:language", "en")

    // Retrieve data
    if user, exists := cache.Get("user:1"); exists {
        fmt.Println("User found:", user)
    }

    if theme, exists := cache.Get("config:theme"); exists {
        fmt.Println("Theme:", theme)
    }

    // Check non-existent key
    if _, exists := cache.Get("user:999"); !exists {
        fmt.Println("User 999 not found")
    }
}
```

### Example 6: Map for Counting

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "hello world hello go world go hello"
    words := strings.Fields(text)

    // Count word frequency
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println("Word frequency:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }

    // Find most frequent word
    maxCount := 0
    mostFrequent := ""
    for word, count := range wordCount {
        if count > maxCount {
            maxCount = count
            mostFrequent = word
        }
    }
    fmt.Printf("Most frequent word: %s (%d times)\n", mostFrequent, maxCount)
}
```

## Map Key Types and Constraints

### Valid Key Types (Comparable Types)
```go
// ✅ Valid key types
var m1 map[string]int           // string
var m2 map[int]string           // int
var m3 map[bool]string          // bool
var m4 map[float64]string       // float64
var m5 map[complex128]string    // complex128
var m6 map[[2]int]string        // array
var m7 map[struct{id int}]string // struct with comparable fields

// ❌ Invalid key types (not comparable)
var m8 map[[]int]string         // slice
var m9 map[map[string]int]string // map
var m10 map[func()]string       // function
```

### Example 7: Struct as Map Key

```go
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    // Map with struct keys
    distances := map[Point]float64{
        {0, 0}: 0.0,
        {1, 1}: 1.414,
        {2, 2}: 2.828,
    }

    // Add new point
    distances[Point{3, 4}] = 5.0

    // Access by struct key
    fmt.Println("Distance from origin to (1,1):", distances[Point{1, 1}])

    // Iterate over struct keys
    for point, distance := range distances {
        fmt.Printf("Point (%d,%d): distance %.3f\n", point.X, point.Y, distance)
    }
}
```

## Nil Maps vs Empty Maps

### Nil Maps
```go
var m map[string]int
fmt.Println(m == nil)        // true
fmt.Println(len(m))          // 0
// m["key"] = 1              // ❌ Panic: assignment to entry in nil map
```

### Empty Maps
```go
m := make(map[string]int)     // or map[string]int{}
fmt.Println(m == nil)        // false
fmt.Println(len(m))          // 0
m["key"] = 1                 // ✅ Works fine
```

### Example 8: Safe Map Operations

```go
package main

import "fmt"

func safeMapOperations() {
    var nilMap map[string]int
    emptyMap := make(map[string]int)

    // Check if map is nil before operations
    if nilMap == nil {
        fmt.Println("nilMap is nil, initializing...")
        nilMap = make(map[string]int)
    }

    // Safe operations
    nilMap["key1"] = 100
    emptyMap["key2"] = 200

    // Safe value retrieval
    if value, exists := nilMap["key1"]; exists {
        fmt.Println("Found in nilMap:", value)
    }

    if value, exists := emptyMap["key2"]; exists {
        fmt.Println("Found in emptyMap:", value)
    }

    // Check for non-existent key
    if value, exists := emptyMap["nonexistent"]; exists {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Key not found, zero value:", value)
    }
}
```

## Performance Considerations

### Time Complexity
- **Access**: O(1) average, O(n) worst case
- **Insert**: O(1) average, O(n) worst case
- **Delete**: O(1) average, O(n) worst case
- **Iteration**: O(n)

### Memory Usage
```go
// Maps have overhead compared to slices
// Use slices when you need ordered data
// Use maps when you need key-value lookups
```

### Example 9: Performance Comparison

```go
package main

import (
    "fmt"
    "time"
)

func sliceLookup(slice []string, target string) bool {
    for _, item := range slice {
        if item == target {
            return true
        }
    }
    return false
}

func mapLookup(m map[string]bool, target string) bool {
    return m[target]
}

func main() {
    // Create test data
    items := []string{"apple", "banana", "cherry", "date", "elderberry"}
    itemMap := map[string]bool{
        "apple":      true,
        "banana":     true,
        "cherry":     true,
        "date":       true,
        "elderberry": true,
    }

    target := "date"

    // Test slice lookup
    start := time.Now()
    for i := 0; i < 1000000; i++ {
        sliceLookup(items, target)
    }
    sliceTime := time.Since(start)

    // Test map lookup
    start = time.Now()
    for i := 0; i < 1000000; i++ {
        mapLookup(itemMap, target)
    }
    mapTime := time.Since(start)

    fmt.Printf("Slice lookup time: %v\n", sliceTime)
    fmt.Printf("Map lookup time: %v\n", mapTime)
    fmt.Printf("Map is %.2fx faster\n", float64(sliceTime)/float64(mapTime))
}
```

## Common Patterns and Best Practices

### Pattern 1: Map as Set
```go
func main() {
    // Using map as a set
    set := make(map[string]bool)
    
    // Add elements
    set["apple"] = true
    set["banana"] = true
    set["cherry"] = true
    
    // Check membership
    if set["apple"] {
        fmt.Println("apple is in set")
    }
    
    // Remove element
    delete(set, "banana")
}
```

### Pattern 2: Map with Default Values
```go
func getWithDefault(m map[string]int, key string, defaultValue int) int {
    if value, exists := m[key]; exists {
        return value
    }
    return defaultValue
}
```

### Pattern 3: Map Merging
```go
func mergeMaps(m1, m2 map[string]int) map[string]int {
    result := make(map[string]int)
    
    // Copy m1
    for k, v := range m1 {
        result[k] = v
    }
    
    // Copy m2 (overwrites m1 values if keys exist)
    for k, v := range m2 {
        result[k] = v
    }
    
    return result
}
```

## Error Handling

### Example 10: Safe Map Operations with Error Handling

```go
package main

import (
    "errors"
    "fmt"
)

var ErrKeyNotFound = errors.New("key not found")

func safeGet(m map[string]int, key string) (int, error) {
    if value, exists := m[key]; exists {
        return value, nil
    }
    return 0, ErrKeyNotFound
}

func main() {
    data := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    // Safe retrieval
    if value, err := safeGet(data, "a"); err == nil {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Error:", err)
    }

    if value, err := safeGet(data, "x"); err == nil {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Error:", err)
    }
}
```

## Conclusion

Maps in Go are powerful data structures that provide efficient key-value storage and retrieval. They are essential for:

1. **Fast lookups**: O(1) average time complexity
2. **Data organization**: Associating related data
3. **Caching**: Storing computed results
4. **Counting**: Frequency analysis
5. **Sets**: Unique value collections

Understanding maps is crucial for writing efficient Go programs, especially when dealing with data that needs to be accessed by keys rather than indices.
