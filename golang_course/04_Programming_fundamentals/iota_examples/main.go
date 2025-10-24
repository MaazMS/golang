package main

import "fmt"

// Example 1: Basic IOTA Usage
func basicIotaExample() {
	fmt.Println("=== Basic IOTA Example ===")

	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
		d = iota // 3
	)

	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("c = %d\n", c)
	fmt.Printf("d = %d\n", d)
	fmt.Println()
}

// Example 2: Simplified IOTA Syntax
func simplifiedIotaExample() {
	fmt.Println("=== Simplified IOTA Syntax ===")

	const (
		first  = iota // 0
		second        // 1 (automatically iota)
		third         // 2 (automatically iota)
		fourth        // 3 (automatically iota)
	)

	fmt.Printf("first = %d\n", first)
	fmt.Printf("second = %d\n", second)
	fmt.Printf("third = %d\n", third)
	fmt.Printf("fourth = %d\n", fourth)
	fmt.Println()
}

// Example 3: IOTA Resets
func iotaResetsExample() {
	fmt.Println("=== IOTA Resets Example ===")

	const (
		group1First  = iota // 0
		group1Second = iota // 1
		group1Third  = iota // 2
	)

	const (
		group2First  = iota // 0 (resets!)
		group2Second = iota // 1
		group2Third  = iota // 2
	)

	fmt.Println("Group 1:")
	fmt.Printf("group1First = %d\n", group1First)
	fmt.Printf("group1Second = %d\n", group1Second)
	fmt.Printf("group1Third = %d\n", group1Third)

	fmt.Println("\nGroup 2 (IOTA resets):")
	fmt.Printf("group2First = %d\n", group2First)
	fmt.Printf("group2Second = %d\n", group2Second)
	fmt.Printf("group2Third = %d\n", group2Third)
	fmt.Println()
}

// Example 4: Days of the Week
func daysOfWeekExample() {
	fmt.Println("=== Days of the Week Example ===")

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Days of the week:")
	fmt.Printf("Sunday = %d\n", Sunday)
	fmt.Printf("Monday = %d\n", Monday)
	fmt.Printf("Tuesday = %d\n", Tuesday)
	fmt.Printf("Wednesday = %d\n", Wednesday)
	fmt.Printf("Thursday = %d\n", Thursday)
	fmt.Printf("Friday = %d\n", Friday)
	fmt.Printf("Saturday = %d\n", Saturday)
	fmt.Println()
}

// Example 5: User Roles
func userRolesExample() {
	fmt.Println("=== User Roles Example ===")

	const (
		Guest = iota
		User
		Moderator
		Admin
		SuperAdmin
	)

	userRole := User

	fmt.Printf("Current user role: %d\n", userRole)

	switch userRole {
	case Guest:
		fmt.Println("👋 Welcome, guest!")
	case User:
		fmt.Println("👤 Welcome back, user!")
	case Moderator:
		fmt.Println("🛡️ Moderator privileges active")
	case Admin:
		fmt.Println("⚙️ Admin panel access granted")
	case SuperAdmin:
		fmt.Println("🔐 Full system access")
	}
	fmt.Println()
}

// Example 6: File Permissions (Bit Flags)
func filePermissionsExample() {
	fmt.Println("=== File Permissions Example ===")

	const (
		Read    = 1 << iota // 1 (binary: 001)
		Write               // 2 (binary: 010)
		Execute             // 4 (binary: 100)
		Delete              // 8 (binary: 1000)
	)

	// Combining permissions
	readWrite := Read | Write                   // 3 (binary: 011)
	allPerms := Read | Write | Execute | Delete // 15 (binary: 1111)

	fmt.Printf("Read permission: %d (binary: %04b)\n", Read, Read)
	fmt.Printf("Write permission: %d (binary: %04b)\n", Write, Write)
	fmt.Printf("Execute permission: %d (binary: %04b)\n", Execute, Execute)
	fmt.Printf("Delete permission: %d (binary: %04b)\n", Delete, Delete)
	fmt.Printf("Read + Write: %d (binary: %04b)\n", readWrite, readWrite)
	fmt.Printf("All permissions: %d (binary: %04b)\n", allPerms, allPerms)

	// Checking permissions
	if allPerms&Read != 0 {
		fmt.Println("✅ Read permission is set")
	}
	if allPerms&Write != 0 {
		fmt.Println("✅ Write permission is set")
	}
	if allPerms&Execute != 0 {
		fmt.Println("✅ Execute permission is set")
	}
	if allPerms&Delete != 0 {
		fmt.Println("✅ Delete permission is set")
	}
	fmt.Println()
}

// Example 7: Starting from Different Value
func startingFromDifferentValue() {
	fmt.Println("=== Starting from Different Value ===")

	const (
		_        = iota // Skip 0
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
	fmt.Printf("April = %d\n", April)
	fmt.Printf("May = %d\n", May)
	fmt.Printf("June = %d\n", June)
	fmt.Println()
}

// Example 8: Skipping Values
func skippingValuesExample() {
	fmt.Println("=== Skipping Values Example ===")

	const (
		Small  = iota // 0
		Medium = iota // 1
		_             // Skip 2
		_             // Skip 3
		Large  = iota // 4
		XLarge = iota // 5
	)

	fmt.Printf("Small = %d\n", Small)
	fmt.Printf("Medium = %d\n", Medium)
	fmt.Printf("Large = %d\n", Large)
	fmt.Printf("XLarge = %d\n", XLarge)
	fmt.Println()
}

// Example 9: Custom Values with IOTA
func customValuesExample() {
	fmt.Println("=== Custom Values with IOTA ===")

	const (
		Apple  = (iota + 1) * 10 // 10
		Orange                   // 20
		Banana                   // 30
		Grape                    // 40
	)

	fmt.Printf("Apple = %d\n", Apple)
	fmt.Printf("Orange = %d\n", Orange)
	fmt.Printf("Banana = %d\n", Banana)
	fmt.Printf("Grape = %d\n", Grape)
	fmt.Println()
}

// Example 10: Log Levels
func logLevelsExample() {
	fmt.Println("=== Log Levels Example ===")

	const (
		LogDebug   = iota // 0
		LogInfo           // 1
		LogWarning        // 2
		LogError          // 3
		LogFatal          // 4
	)

	currentLogLevel := LogInfo

	fmt.Printf("Current log level: %d\n", currentLogLevel)

	if currentLogLevel >= LogInfo {
		fmt.Println("📝 This message will be logged (Info level)")
	}

	if currentLogLevel >= LogWarning {
		fmt.Println("⚠️ This warning will be logged")
	} else {
		fmt.Println("⚠️ Warning level not reached")
	}
	fmt.Println()
}

// Example 11: Database Operations
func databaseOperationsExample() {
	fmt.Println("=== Database Operations Example ===")

	const (
		OpCreate = iota // 0
		OpRead          // 1
		OpUpdate        // 2
		OpDelete        // 3
	)

	operation := OpRead

	fmt.Printf("Current operation: %d\n", operation)

	switch operation {
	case OpCreate:
		fmt.Println("➕ Creating new record...")
	case OpRead:
		fmt.Println("📖 Reading record...")
	case OpUpdate:
		fmt.Println("✏️ Updating record...")
	case OpDelete:
		fmt.Println("🗑️ Deleting record...")
	default:
		fmt.Println("❓ Unknown operation")
	}
	fmt.Println()
}

// Example 12: HTTP Status Categories
func httpStatusCategoriesExample() {
	fmt.Println("=== HTTP Status Categories Example ===")

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
		_                   = iota
		SuccessCategory     // 1
		ClientErrorCategory // 2
		ServerErrorCategory // 3
	)

	responseCode := 404

	fmt.Printf("Response code: %d\n", responseCode)

	if responseCode >= 200 && responseCode < 300 {
		fmt.Printf("✅ Success category: %d\n", SuccessCategory)
	} else if responseCode >= 400 && responseCode < 500 {
		fmt.Printf("❌ Client error category: %d\n", ClientErrorCategory)
	} else if responseCode >= 500 {
		fmt.Printf("💥 Server error category: %d\n", ServerErrorCategory)
	}
	fmt.Println()
}

// Example 13: String Constants with IOTA
func stringConstantsWithIotaExample() {
	fmt.Println("=== String Constants with IOTA ===")

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
	fmt.Println()
}

func main() {
	fmt.Println("🚀 Go IOTA - Complete Examples")
	fmt.Println("===============================")
	fmt.Println()

	basicIotaExample()
	simplifiedIotaExample()
	iotaResetsExample()
	daysOfWeekExample()
	userRolesExample()
	filePermissionsExample()
	startingFromDifferentValue()
	skippingValuesExample()
	customValuesExample()
	logLevelsExample()
	databaseOperationsExample()
	httpStatusCategoriesExample()
	stringConstantsWithIotaExample()

	fmt.Println("✨ All IOTA examples completed!")
}
