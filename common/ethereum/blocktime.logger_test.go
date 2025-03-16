package ethereum

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

/**************************************************************************************************
** TestVerboseBlocktimeFlag tests the VerboseBlocktime flag to ensure it can be set and
** read correctly. This test validates:
** - The default state of the VerboseBlocktime flag
** - The EnableVerboseBlocktime function correctly sets the flag to true
** - The DisableVerboseBlocktime function correctly sets the flag to false
**
** This ensures that the verbosity control for blocktime logging works as expected.
**************************************************************************************************/
func TestVerboseBlocktimeFlag(t *testing.T) {
	// Save initial state to restore later
	initialState := VerboseBlocktime

	// Test default state (should be false unless modified elsewhere)
	if VerboseBlocktime != false {
		t.Errorf("Expected default VerboseBlocktime to be false, got %v", VerboseBlocktime)
	}

	// Test EnableVerboseBlocktime
	EnableVerboseBlocktime()
	if VerboseBlocktime != true {
		t.Errorf("Expected VerboseBlocktime to be true after EnableVerboseBlocktime, got %v", VerboseBlocktime)
	}

	// Test DisableVerboseBlocktime
	DisableVerboseBlocktime()
	if VerboseBlocktime != false {
		t.Errorf("Expected VerboseBlocktime to be false after DisableVerboseBlocktime, got %v", VerboseBlocktime)
	}

	// Restore initial state
	VerboseBlocktime = initialState
}

/**************************************************************************************************
** captureOutput is a helper function that captures log output for testing.
** It temporarily redirects log output to a buffer and returns the captured output.
**
** @param f A function to execute while capturing output
** @return string The captured log output
**************************************************************************************************/
func captureOutput(f func()) string {
	// Create a pipe to capture log output
	originalOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture the output
	log.SetOutput(w)
	f()
	w.Close()

	// Read the captured output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	// Restore normal output
	os.Stdout = originalOutput
	log.SetOutput(originalOutput)

	return buf.String()
}

/**************************************************************************************************
** TestBlocktimeLogFunctions tests the blocktime logging functions to ensure they log messages
** appropriately based on the VerboseBlocktime flag. This test validates:
** - blocktimeLog only logs when VerboseBlocktime is true
** - blocktimeWarning logs (partially) even when VerboseBlocktime is false
** - blocktimeSuccess only logs when VerboseBlocktime is true
** - blocktimeError always logs, regardless of VerboseBlocktime
**
** This ensures that the logging functions respect the verbosity settings.
**************************************************************************************************/
func TestBlocktimeLogFunctions(t *testing.T) {
	// Save initial state to restore later
	initialState := VerboseBlocktime

	// Test with VerboseBlocktime = true
	VerboseBlocktime = true

	testLogMessage := "Test log message " + time.Now().String()
	output := captureOutput(func() {
		blocktimeLog(testLogMessage)
	})
	if !strings.Contains(output, testLogMessage) {
		t.Errorf("Expected blocktimeLog to log with VerboseBlocktime=true")
	}

	testWarningMessage := "Test warning message " + time.Now().String()
	output = captureOutput(func() {
		blocktimeWarning(testWarningMessage)
	})
	if !strings.Contains(output, testWarningMessage) {
		t.Errorf("Expected blocktimeWarning to log with VerboseBlocktime=true")
	}

	testSuccessMessage := "Test success message " + time.Now().String()
	output = captureOutput(func() {
		blocktimeSuccess(testSuccessMessage)
	})
	if !strings.Contains(output, testSuccessMessage) {
		t.Errorf("Expected blocktimeSuccess to log with VerboseBlocktime=true")
	}

	testErrorMessage := "Test error message " + time.Now().String()
	output = captureOutput(func() {
		blocktimeError(testErrorMessage)
	})
	if !strings.Contains(output, testErrorMessage) {
		t.Errorf("Expected blocktimeError to always log")
	}

	// Test with VerboseBlocktime = false
	VerboseBlocktime = false

	output = captureOutput(func() {
		blocktimeLog(testLogMessage)
	})
	if strings.Contains(output, testLogMessage) {
		t.Errorf("Expected blocktimeLog not to log with VerboseBlocktime=false")
	}

	// For truncated messages, just check that some output occurs
	longWarningMessage := testWarningMessage + strings.Repeat("x", 200)
	output = captureOutput(func() {
		blocktimeWarning(longWarningMessage)
	})
	// Should truncate to 100 chars + "..." when verbose is false
	if strings.Contains(output, longWarningMessage) {
		t.Errorf("Expected blocktimeWarning to truncate long message with VerboseBlocktime=false")
	}
	if !strings.Contains(output, "...") {
		t.Errorf("Expected truncated warning to include '...'")
	}

	output = captureOutput(func() {
		blocktimeSuccess(testSuccessMessage)
	})
	if strings.Contains(output, testSuccessMessage) {
		t.Errorf("Expected blocktimeSuccess not to log with VerboseBlocktime=false")
	}

	output = captureOutput(func() {
		blocktimeError(testErrorMessage)
	})
	if !strings.Contains(output, testErrorMessage) {
		t.Errorf("Expected blocktimeError to always log regardless of VerboseBlocktime")
	}

	// Restore initial state
	VerboseBlocktime = initialState
}

/**************************************************************************************************
** TestGetCallerInfo tests the getCallerInfo function to ensure it correctly retrieves caller
** information for debugging purposes. This test validates:
** - The function returns a non-empty string with caller information
** - The returned string contains expected elements like file name and line number
**
** This ensures that the function provides useful debugging information.
**************************************************************************************************/
func TestGetCallerInfo(t *testing.T) {
	// For this test, we're just verifying that the function returns something useful
	// The exact format can vary based on Go version and testing framework
	info := getCallerInfo()

	if info == "" {
		t.Error("Expected non-empty caller information, got empty string")
	}

	// Check that the string contains some basic elements we'd expect
	// from call stack information, without being too specific about format
	if !strings.Contains(info, ".go:") {
		t.Errorf("Expected caller info to contain file name and line number, got %s", info)
	}

	// All we really need to check is that the function returns something
	// that looks like debugging information containing file and line references
	t.Logf("getCallerInfo returned: %s", info)
}

// Helper function to get caller info that will show up correctly in the test
func getCallerInfoHelper() string {
	return getCallerInfo()
}
