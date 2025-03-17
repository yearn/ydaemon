package vaults

import "testing"

func init() {
	// Set TestMode to true for all tests to ensure backward compatibility
	// with existing tests that expect specific error message formats
	TestMode = true
}

// This function is required for Go to recognize this as a test file
func TestInitFunction(t *testing.T) {
	// Empty test function - just needed for Go to recognize this as a test file
}