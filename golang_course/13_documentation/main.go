// Package writedoc demonstrates Go documentation standards and best practices.
//
// This package provides examples of proper Go documentation including:
// - Package-level documentation
// - Function and method documentation
// - Type documentation
// - Variable documentation
// - Example functions
//
// Documentation in Go is generated from comments and follows specific
// conventions to create readable, searchable documentation.
//
// Usage:
//   go run main.go
//   go doc writedoc
//   godoc -http=:6060
//
// For more information about Go documentation:
// https://golang.org/doc/effective_go.html#commentary
package main

import (
	"fmt"
	"time"
)

// ProgramInfo holds information about the program.
type ProgramInfo struct {
	Name        string    // Program name
	Version     string    // Program version
	StartTime   time.Time // Program start time
	Description string    // Program description
}

// NewProgramInfo creates a new ProgramInfo instance with default values.
// It initializes the program information with the current timestamp
// and provided name and version.
//
// Parameters:
//   name: The name of the program
//   version: The version string (e.g., "1.0.0")
//
// Returns:
//   *ProgramInfo: A pointer to the initialized ProgramInfo struct
//
// Example:
//   info := NewProgramInfo("writedoc", "1.0.0")
//   fmt.Printf("Program: %s v%s\n", info.Name, info.Version)
func NewProgramInfo(name, version string) *ProgramInfo {
	return &ProgramInfo{
		Name:        name,
		Version:     version,
		StartTime:   time.Now(),
		Description: "Go documentation demonstration package",
	}
}

// DisplayInfo prints the program information to standard output.
// It formats and displays all the program information in a readable format.
//
// Example output:
//   Program: writedoc v1.0.0
//   Started: 2024-01-15 10:30:45
//   Description: Go documentation demonstration package
func (p *ProgramInfo) DisplayInfo() {
	fmt.Printf("Program: %s v%s\n", p.Name, p.Version)
	fmt.Printf("Started: %s\n", p.StartTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Description: %s\n", p.Description)
}

// main is the entry point of the writedoc demonstration program.
// It creates a ProgramInfo instance and displays the program information
// to demonstrate proper Go documentation practices.
//
// The function showcases:
//   - Package documentation standards
//   - Function documentation with examples
//   - Type documentation
//   - Method documentation
//
// Example usage:
//   go run main.go
//
// Expected output:
//   Program: writedoc v1.0.0
//   Started: 2024-01-15 10:30:45
//   Description: Go documentation demonstration package
func main() {
	// Create program info
	info := NewProgramInfo("writedoc", "1.0.0")

	// Display program information
	info.DisplayInfo()

	// Demonstrate documentation features
	fmt.Println("\nDocumentation Features Demonstrated:")
	fmt.Println("- Package-level documentation")
	fmt.Println("- Function documentation with examples")
	fmt.Println("- Type and method documentation")
	fmt.Println("- Comment formatting and conventions")
}
