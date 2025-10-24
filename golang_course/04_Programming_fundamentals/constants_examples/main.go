package main

import "fmt"

// Example 1: Basic Constants
func basicConstants() {
	fmt.Println("=== Basic Constants Example ===")

	const pi = 3.14159
	const appName = "Go Learning App"
	const maxUsers = 1000
	const isActive = true

	fmt.Printf("Pi: %.5f\n", pi)
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Max Users: %d\n", maxUsers)
	fmt.Printf("Is Active: %t\n", isActive)
	fmt.Println()
}

// Example 2: Multiple Constants Declaration
func multipleConstants() {
	fmt.Println("=== Multiple Constants Example ===")

	const (
		companyName = "TechCorp"
		version     = "2.1.0"
		maxRetries  = 3
		timeout     = 30
		debugMode   = false
	)

	fmt.Printf("Company: %s\n", companyName)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Max Retries: %d\n", maxRetries)
	fmt.Printf("Timeout: %d seconds\n", timeout)
	fmt.Printf("Debug Mode: %t\n", debugMode)
	fmt.Println()
}

// Example 3: Typed vs Untyped Constants
func typedVsUntyped() {
	fmt.Println("=== Typed vs Untyped Constants ===")

	// Typed constants
	const typedInt int = 42
	const typedFloat float64 = 3.14
	const typedString string = "Hello"

	// Untyped constants
	const untypedInt = 42
	const untypedFloat = 3.14
	const untypedString = "Hello"

	fmt.Printf("Typed int: %d (Type: %T)\n", typedInt, typedInt)
	fmt.Printf("Untyped int: %d (Type: %T)\n", untypedInt, untypedInt)
	fmt.Printf("Typed float: %.2f (Type: %T)\n", typedFloat, typedFloat)
	fmt.Printf("Untyped float: %.2f (Type: %T)\n", untypedFloat, untypedFloat)
	fmt.Printf("Typed string: %s (Type: %T)\n", typedString, typedString)
	fmt.Printf("Untyped string: %s (Type: %T)\n", untypedString, untypedString)
	fmt.Println()
}

// Example 4: Constants in Functions
func calculateArea(radius float64) float64 {
	const pi = 3.14159
	return pi * radius * radius
}

func constantsInFunctions() {
	fmt.Println("=== Constants in Functions ===")

	radius := 5.0
	area := calculateArea(radius)

	fmt.Printf("Area of circle with radius %.1f: %.2f\n", radius, area)
	fmt.Println()
}

// Example 5: HTTP Status Codes
func httpStatusExample() {
	fmt.Println("=== HTTP Status Codes Example ===")

	const (
		statusOK       = 200
		statusCreated  = 201
		statusNotFound = 404
		statusError    = 500
	)

	responseCode := 200

	switch responseCode {
	case statusOK:
		fmt.Println("✅ Request successful!")
	case statusCreated:
		fmt.Println("✅ Resource created successfully!")
	case statusNotFound:
		fmt.Println("❌ Resource not found")
	case statusError:
		fmt.Println("❌ Internal server error")
	default:
		fmt.Println("❓ Unknown status code")
	}
	fmt.Println()
}

// Example 6: Application Configuration
func appConfigExample() {
	fmt.Println("=== Application Configuration Example ===")

	const (
		appName     = "MyGoApp"
		version     = "1.0.0"
		environment = "development"
		maxRetries  = 3
		timeout     = 30
		debugMode   = true
	)

	fmt.Printf("Application: %s v%s\n", appName, version)
	fmt.Printf("Environment: %s\n", environment)
	fmt.Printf("Max Retries: %d\n", maxRetries)
	fmt.Printf("Timeout: %d seconds\n", timeout)
	fmt.Printf("Debug Mode: %t\n", debugMode)
	fmt.Println()
}

// Example 7: Mathematical Constants
func mathConstantsExample() {
	fmt.Println("=== Mathematical Constants Example ===")

	const (
		pi     = 3.14159
		e      = 2.71828
		golden = 1.61803
	)

	radius := 7.0
	circumference := 2 * pi * radius
	area := pi * radius * radius

	fmt.Printf("Pi: %.5f\n", pi)
	fmt.Printf("Euler's number: %.5f\n", e)
	fmt.Printf("Golden ratio: %.5f\n", golden)
	fmt.Printf("Circumference of circle (r=%.1f): %.2f\n", radius, circumference)
	fmt.Printf("Area of circle (r=%.1f): %.2f\n", radius, area)
	fmt.Println()
}

// Example 8: Constants vs Variables Comparison
func constantsVsVariables() {
	fmt.Println("=== Constants vs Variables ===")

	// Constants
	const maxSize = 100
	const appName = "Calculator"

	// Variables
	var currentSize = 50
	var userName = "John"

	fmt.Printf("Max size (constant): %d\n", maxSize)
	fmt.Printf("Current size (variable): %d\n", currentSize)
	fmt.Printf("App name (constant): %s\n", appName)
	fmt.Printf("User name (variable): %s\n", userName)

	// Variables can be changed
	currentSize = 75
	userName = "Jane"

	fmt.Printf("Updated current size: %d\n", currentSize)
	fmt.Printf("Updated user name: %s\n", userName)

	// Constants cannot be changed (this would cause a compilation error):
	// maxSize = 200  // ❌ Cannot assign to maxSize
	fmt.Println()
}

func main() {
	fmt.Println("🚀 Go Constants - Complete Examples")
	fmt.Println("=====================================")
	fmt.Println()

	basicConstants()
	multipleConstants()
	typedVsUntyped()
	constantsInFunctions()
	httpStatusExample()
	appConfigExample()
	mathConstantsExample()
	constantsVsVariables()

	fmt.Println("✨ All examples completed!")
}
