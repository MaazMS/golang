package main

import (
	"testing"
	"time"
)

// TestNewProgramInfo tests the NewProgramInfo function.
// It verifies that the function creates a ProgramInfo instance
// with the correct values and current timestamp.
func TestNewProgramInfo(t *testing.T) {
	name := "test-program"
	version := "1.0.0"

	info := NewProgramInfo(name, version)

	// Test name
	if info.Name != name {
		t.Errorf("Expected name %s, got %s", name, info.Name)
	}

	// Test version
	if info.Version != version {
		t.Errorf("Expected version %s, got %s", version, info.Version)
	}

	// Test description
	expectedDesc := "Go documentation demonstration package"
	if info.Description != expectedDesc {
		t.Errorf("Expected description %s, got %s", expectedDesc, info.Description)
	}

	// Test start time is recent (within last second)
	now := time.Now()
	if info.StartTime.After(now) || info.StartTime.Before(now.Add(-time.Second)) {
		t.Errorf("StartTime should be recent, got %s", info.StartTime)
	}
}

// TestProgramInfoDisplayInfo tests the DisplayInfo method.
// It verifies that the method doesn't panic and produces output.
func TestProgramInfoDisplayInfo(t *testing.T) {
	info := NewProgramInfo("test", "1.0.0")

	// This test mainly ensures the method doesn't panic
	// In a real test, you might capture output and verify it
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DisplayInfo panicked: %v", r)
		}
	}()

	info.DisplayInfo()
}

// ExampleNewProgramInfo demonstrates how to use NewProgramInfo.
func ExampleNewProgramInfo() {
	info := NewProgramInfo("example", "2.0.0")
	fmt.Printf("Program: %s v%s\n", info.Name, info.Version)
	// Output: Program: example v2.0.0
}

// ExampleProgramInfo_DisplayInfo demonstrates how to use the DisplayInfo method.
func ExampleProgramInfo_DisplayInfo() {
	info := NewProgramInfo("example", "2.0.0")
	info.DisplayInfo()
	// Output:
	// Program: example v2.0.0
	// Started: 2024-01-15 10:30:45
	// Description: Go documentation demonstration package
}

// BenchmarkNewProgramInfo benchmarks the NewProgramInfo function.
func BenchmarkNewProgramInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewProgramInfo("benchmark", "1.0.0")
	}
}
