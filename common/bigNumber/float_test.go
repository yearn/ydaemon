package bigNumber

import (
	"encoding/json"
	"math/big"
	"strconv"
	"testing"
)

/**************************************************************************************************
** TestNewFloat tests the creation of new Float instances using various initialization approaches.
** This validates that Float objects are properly created with expected values using:
** - Default initialization (zero value)
** - Initialization with specific float64 values
** - Multiple parameter handling
**
** Proper initialization is fundamental to ensure arithmetic operations start with correct values.
**************************************************************************************************/
func TestNewFloat(t *testing.T) {
	tests := []struct {
		name     string
		args     []float64
		expected string
	}{
		{
			name:     "Default initialization (zero)",
			args:     []float64{},
			expected: "0",
		},
		{
			name:     "Initialization with positive value",
			args:     []float64{123.456},
			expected: "123.456",
		},
		{
			name:     "Initialization with negative value",
			args:     []float64{-456.789},
			expected: "-456.789",
		},
		{
			name:     "Initialization with multiple args (should use first)",
			args:     []float64{789.012, 123.456},
			expected: "789.012",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Float
			if len(tt.args) == 0 {
				result = NewFloat()
			} else {
				result = NewFloat(tt.args...)
			}

			// For float comparisons, we should check if the difference is within an acceptable range
			// but for simplicity, we'll just check the string representation
			if result.String() != tt.expected {
				t.Errorf("NewFloat(%v) = %v, expected %v", tt.args, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestSetFloat tests the creation of new Float instances from standard library big.Float values.
** This validates that Float objects are properly created from big.Float pointers with:
** - Default initialization (zero value) when no value is provided
** - Initialization from a valid big.Float pointer
** - Nil pointer handling
**
** This ensures compatibility between the standard library and the custom Float wrapper.
**************************************************************************************************/
func TestSetFloat(t *testing.T) {
	// Create test big.Float values
	bigVal1 := big.NewFloat(123.456)
	var nilBigFloat *big.Float

	tests := []struct {
		name     string
		args     []*big.Float
		expected string
	}{
		{
			name:     "Default initialization (zero)",
			args:     []*big.Float{},
			expected: "0",
		},
		{
			name:     "Initialization from big.Float",
			args:     []*big.Float{bigVal1},
			expected: "123.456",
		},
		{
			name:     "Initialization from nil big.Float",
			args:     []*big.Float{nilBigFloat},
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Float
			if len(tt.args) == 0 {
				result = SetFloat()
			} else {
				result = SetFloat(tt.args...)
			}
			if result.String() != tt.expected {
				t.Errorf("SetFloat(%v) = %v, expected %v", tt.args, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestSetString tests the SetString method of the Float type.
** This validates that Float values can be correctly set from string representations:
** - Valid floating-point strings
** - Empty strings
** - Quoted empty strings (common in JSON)
**
** String parsing is critical for handling user input and data from external sources.
**************************************************************************************************/
func TestFloatSetString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Valid positive number", "123.456", "123.456"},
		{"Valid negative number", "-456.789", "-456.789"},
		{"Valid integer", "789", "789"},
		{"Valid zero", "0", "0"},
		{"Scientific notation", "1.23e4", "12300"},
		{"Empty string", "", "0"},
		{"Quoted empty string", "\"\"", "0"},
		{"Very small number", "0.0000000000123", "0.0000000000123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFloat().SetString(tt.input)
			// Convert both strings to float64 for comparison to avoid formatting differences
			expected, _ := strconv.ParseFloat(tt.expected, 64)
			actual, _ := strconv.ParseFloat(result.String(), 64)

			// Compare with tolerance for floating point imprecision
			if expected != actual {
				t.Errorf("SetString(%s) = %s, expected %s", tt.input, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestFloatArithmetic tests the basic arithmetic operations (Add, Sub, Mul, Quo, Div) of the Float type.
** This validates that arithmetic operations produce correct results with various inputs,
** including special cases like division by zero.
**
** These tests ensure the correctness of the core arithmetic functionality upon which more
** complex calculations depend.
**************************************************************************************************/
func TestFloatArithmetic(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		a := NewFloat(5.5)
		b := NewFloat(3.2)
		c := NewFloat(0)
		c.Add(a, b)

		expected := 8.7
		actual, _ := strconv.ParseFloat(c.String(), 64)
		if actual != expected {
			t.Errorf("5.5 + 3.2 = %s, expected 8.7", c.String())
		}
	})

	t.Run("Subtraction", func(t *testing.T) {
		a := NewFloat(10.8)
		b := NewFloat(3.3)
		c := NewFloat(0)
		c.Sub(a, b)

		// Use a small epsilon for floating point comparison to handle precision issues
		expected := 7.5
		actual, _ := strconv.ParseFloat(c.String(), 64)
		epsilon := 0.0000001
		if actual < expected-epsilon || actual > expected+epsilon {
			t.Errorf("10.8 - 3.3 = %s, expected 7.5", c.String())
		}
	})

	t.Run("Multiplication", func(t *testing.T) {
		a := NewFloat(4.5)
		b := NewFloat(2.0)
		c := NewFloat(0)
		c.Mul(a, b)

		expected := 9.0
		actual, _ := strconv.ParseFloat(c.String(), 64)
		if actual != expected {
			t.Errorf("4.5 * 2.0 = %s, expected 9.0", c.String())
		}
	})

	t.Run("Division (Quo)", func(t *testing.T) {
		a := NewFloat(20.0)
		b := NewFloat(4.0)
		c := NewFloat(0)
		c.Quo(a, b)

		expected := 5.0
		actual, _ := strconv.ParseFloat(c.String(), 64)
		if actual != expected {
			t.Errorf("20.0 / 4.0 = %s, expected 5.0", c.String())
		}
	})

	t.Run("Division (Div)", func(t *testing.T) {
		a := NewFloat(20.0)
		b := NewFloat(4.0)
		c := NewFloat(0)
		c.Div(a, b)

		expected := 5.0
		actual, _ := strconv.ParseFloat(c.String(), 64)
		if actual != expected {
			t.Errorf("20.0 / 4.0 = %s, expected 5.0", c.String())
		}

		// Division by zero
		d := NewFloat(10.0)
		zero := NewFloat(0.0)
		e := NewFloat(0)
		e.Div(d, zero)

		if !e.IsZero() {
			t.Errorf("Division by zero should return 0, got %s", e.String())
		}
	})
}

/**************************************************************************************************
** TestFloatComparisons tests the comparison operations (Gt, Gte, Lt, Lte, Eq, Not, IsZero)
** of the Float type. This validates that comparisons between Float values produce correct results.
**
** These tests are essential because comparisons are fundamental to conditional logic in
** financial calculations, token prices, and other blockchain-related operations.
**************************************************************************************************/
func TestFloatComparisons(t *testing.T) {
	a := NewFloat(10.5)
	b := NewFloat(5.25)
	c := NewFloat(10.5)
	zero := NewFloat(0.0)

	// Greater than
	if !a.Gt(b) {
		t.Errorf("Expected %s > %s", a.String(), b.String())
	}
	if b.Gt(a) {
		t.Errorf("Expected %s not > %s", b.String(), a.String())
	}
	if a.Gt(c) {
		t.Errorf("Expected %s not > %s (equal)", a.String(), c.String())
	}

	// Greater than or equal
	if !a.Gte(b) {
		t.Errorf("Expected %s >= %s", a.String(), b.String())
	}
	if !a.Gte(c) {
		t.Errorf("Expected %s >= %s (equal)", a.String(), c.String())
	}
	if b.Gte(a) {
		t.Errorf("Expected %s not >= %s", b.String(), a.String())
	}

	// Less than
	if !b.Lt(a) {
		t.Errorf("Expected %s < %s", b.String(), a.String())
	}
	if a.Lt(b) {
		t.Errorf("Expected %s not < %s", a.String(), b.String())
	}
	if a.Lt(c) {
		t.Errorf("Expected %s not < %s (equal)", a.String(), c.String())
	}

	// Less than or equal
	if !b.Lte(a) {
		t.Errorf("Expected %s <= %s", b.String(), a.String())
	}
	if !a.Lte(c) {
		t.Errorf("Expected %s <= %s (equal)", a.String(), c.String())
	}
	if a.Lte(b) {
		t.Errorf("Expected %s not <= %s", a.String(), b.String())
	}

	// Equal
	if !a.Eq(c) {
		t.Errorf("Expected %s == %s", a.String(), c.String())
	}
	if a.Eq(b) {
		t.Errorf("Expected %s != %s", a.String(), b.String())
	}

	// Not equal
	if !a.Not(b) {
		t.Errorf("Expected %s != %s", a.String(), b.String())
	}
	if a.Not(c) {
		t.Errorf("Expected %s == %s", a.String(), c.String())
	}

	// IsZero
	if !zero.IsZero() {
		t.Errorf("Expected %s to be zero", zero.String())
	}
	if a.IsZero() {
		t.Errorf("Expected %s not to be zero", a.String())
	}
}

/**************************************************************************************************
** TestFloatPow tests the Pow method of the Float type.
** This validates that exponentiation correctly calculates base^exponent with various inputs:
** - Positive exponents
** - Zero exponent
** - Negative base values
**
** Exponentiation is important for financial calculations such as compound interest and APY.
**************************************************************************************************/
func TestFloatPow(t *testing.T) {
	tests := []struct {
		name     string
		base     *Float
		exp      uint64
		expected float64
	}{
		{
			name:     "2.0^3",
			base:     NewFloat(2.0),
			exp:      3,
			expected: 8.0,
		},
		{
			name:     "3.5^2",
			base:     NewFloat(3.5),
			exp:      2,
			expected: 12.25,
		},
		{
			name:     "Any number^0",
			base:     NewFloat(123.456),
			exp:      0,
			expected: 1.0,
		},
		{
			name:     "Negative base^odd exponent",
			base:     NewFloat(-2.0),
			exp:      3,
			expected: -8.0,
		},
		{
			name:     "Negative base^even exponent",
			base:     NewFloat(-2.0),
			exp:      2,
			expected: 4.0,
		},
		{
			name:     "Fractional base^integer exponent",
			base:     NewFloat(0.5),
			exp:      3,
			expected: 0.125,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFloat().Pow(tt.base, tt.exp)
			actual, _ := strconv.ParseFloat(result.String(), 64)

			// Use a small epsilon for floating point comparison
			epsilon := 0.0000001
			if actual < tt.expected-epsilon || actual > tt.expected+epsilon {
				t.Errorf("%s^%d = %s, expected %v", tt.base.String(), tt.exp, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestFloatConversions tests the conversion methods of the Float type:
** - String(): converts to a string representation
** - Int(): converts to a bigNumber.Int
** - SetInt(): sets value from a bigNumber.Int
**
** Proper conversion between types is critical for interoperability and for performing
** operations that require different numeric representations.
**************************************************************************************************/
func TestFloatConversions(t *testing.T) {
	// Test String conversion
	t.Run("String conversion", func(t *testing.T) {
		tests := []struct {
			value    float64
			expected string
		}{
			{0, "0"},
			{123.456, "123.456"},
			{-456.789, "-456.789"},
		}

		for _, tt := range tests {
			result := NewFloat(tt.value).String()

			// Convert both to float64 to ignore formatting differences
			expected, _ := strconv.ParseFloat(tt.expected, 64)
			actual, _ := strconv.ParseFloat(result, 64)

			// Check if they're equal with a small epsilon
			epsilon := 0.0000001
			if actual < expected-epsilon || actual > expected+epsilon {
				t.Errorf("String conversion of %v = %s, expected %s", tt.value, result, tt.expected)
			}
		}
	})

	// Test Int conversion
	t.Run("Int conversion", func(t *testing.T) {
		tests := []struct {
			value    float64
			expected int64
		}{
			{0, 0},
			{123.456, 123},   // Should truncate toward zero
			{-456.789, -456}, // Should truncate toward zero
			{789.999, 789},   // Should truncate, not round
		}

		for _, tt := range tests {
			result := NewFloat(tt.value).Int()
			if result.String() != strconv.FormatInt(tt.expected, 10) {
				t.Errorf("Int conversion of %v = %s, expected %d", tt.value, result.String(), tt.expected)
			}
		}
	})

	// Test SetInt
	t.Run("SetInt", func(t *testing.T) {
		intVal := NewInt(123)
		floatVal := NewFloat().SetInt(intVal)

		expected := "123"
		if floatVal.String() != expected {
			t.Errorf("SetInt(%s) = %s, expected %s", intVal.String(), floatVal.String(), expected)
		}
	})
}

/**************************************************************************************************
** TestFloatJSONSerialization tests the JSON serialization and deserialization of Float values.
** This validates that Float objects can be correctly:
** - Marshaled to JSON (as strings to prevent JavaScript numeric limitations)
** - Unmarshaled from JSON (from string representation)
** - Handled properly in various edge cases (nil, zero)
**
** Correct JSON serialization is critical for APIs and data exchange with frontend applications.
**************************************************************************************************/
func TestFloatJSONSerialization(t *testing.T) {
	tests := []struct {
		name  string
		value *Float
	}{
		{"Zero value", NewFloat(0.0)},
		{"Positive value", NewFloat(123.456)},
		{"Negative value", NewFloat(-456.789)},
		{"Large value", NewFloat(1e20)},
		{"Small value", NewFloat(1e-20)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonBytes, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatalf("Failed to marshal %s: %v", tt.value.String(), err)
			}

			// Unmarshal from JSON
			var result Float
			err = json.Unmarshal(jsonBytes, &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal %s: %v", string(jsonBytes), err)
			}

			// Compare values
			original, _ := strconv.ParseFloat(tt.value.String(), 64)
			unmarshaled, _ := strconv.ParseFloat(result.String(), 64)

			// Use a small epsilon for floating point comparison
			epsilon := 0.0000001
			if unmarshaled < original-epsilon || unmarshaled > original+epsilon {
				t.Errorf("JSON round-trip: original %s, got %s", tt.value.String(), result.String())
			}
		})
	}

	// Test unmarshaling edge cases
	t.Run("Unmarshal edge cases", func(t *testing.T) {
		var result Float

		// Unmarshal from null
		err := json.Unmarshal([]byte("null"), &result)
		if err != nil {
			t.Errorf("Failed to unmarshal null: %v", err)
		}
		// Result should remain zero
		if !result.IsZero() {
			t.Errorf("Expected zero after unmarshaling null, got %s", result.String())
		}

		// Unmarshal from invalid value
		err = json.Unmarshal([]byte("\"not a number\""), &result)
		if err == nil {
			t.Errorf("Expected error when unmarshaling invalid number")
		}
	})
}

/**************************************************************************************************
** TestFloatCSVSerialization tests the CSV serialization of Float values.
** This validates that Float objects can be correctly marshaled to CSV format.
**
** CSV serialization is important for data export and integration with spreadsheet applications.
**************************************************************************************************/
func TestFloatCSVSerialization(t *testing.T) {
	tests := []struct {
		name     string
		value    *Float
		expected string
	}{
		{"Zero value", NewFloat(0.0), "0"},
		{"Positive value", NewFloat(123.456), "123.456"},
		{"Negative value", NewFloat(-456.789), "-456.789"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.value.MarshalCSV()
			if err != nil {
				t.Fatalf("Failed to marshal to CSV: %v", err)
			}

			// Convert both to float64 to ignore formatting differences
			expected, _ := strconv.ParseFloat(tt.expected, 64)
			actual, _ := strconv.ParseFloat(result, 64)

			// Check if they're equal with a small epsilon
			epsilon := 0.0000001
			if actual < expected-epsilon || actual > expected+epsilon {
				t.Errorf("CSV marshaling of %v = %s, expected %s", tt.value, result, tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestFloatCloneAndSet tests the Clone and Set methods of the Float type.
** This validates that Float values can be correctly:
** - Cloned from one Float to another
** - Set from a standard library big.Float
** - Handled properly with nil pointers
**
** These operations are important for safely manipulating Float values without
** accidentally modifying the original.
**************************************************************************************************/
func TestFloatCloneAndSet(t *testing.T) {
	t.Run("Clone method", func(t *testing.T) {
		original := NewFloat(123.456)
		clone := NewFloat().Clone(original)

		// Verify clone has the same value
		if !original.Eq(clone) {
			t.Errorf("Clone value %s doesn't match original %s", clone.String(), original.String())
		}

		// Modify clone and verify original is unchanged
		clone.Add(clone, NewFloat(1.0))
		if original.Eq(clone) {
			t.Errorf("Original should not change when clone is modified")
		}

		// Test cloning nil
		nilResult := NewFloat().Clone(nil)
		if nilResult.String() != "0" {
			t.Errorf("Cloning nil should leave value unchanged, got %s", nilResult.String())
		}
	})

	t.Run("Set method", func(t *testing.T) {
		original := big.NewFloat(123.456)
		floatVal := NewFloat().Set(original)

		// Verify set has the correct value
		expected, _ := original.Float64()
		actual, _ := strconv.ParseFloat(floatVal.String(), 64)

		// Use a small epsilon for floating point comparison
		epsilon := 0.0000001
		if actual < expected-epsilon || actual > expected+epsilon {
			t.Errorf("Set value %s doesn't match expected %v", floatVal.String(), expected)
		}

		// Test setting from nil
		nilResult := NewFloat(123.456).Set(nil)
		if nilResult.String() != "0" {
			t.Errorf("Setting from nil should set value to 0, got %s", nilResult.String())
		}
	})
}

/**************************************************************************************************
** TestFloatSafeMethod tests the Safe method of the Float type.
** This validates that the Safe method correctly handles various inputs:
** - Valid Float objects
** - Nil Float pointers
** - Default values for nil cases
**
** The Safe method is important for preventing nil pointer panics in production code.
**************************************************************************************************/
func TestFloatSafeMethod(t *testing.T) {
	// Setup test values
	validFloat := NewFloat(123.456)
	var nilFloat *Float
	defaultFloat := NewFloat(999.999)

	tests := []struct {
		name     string
		input    *Float
		default_ []*Float
		expected float64
	}{
		{"Valid Float", validFloat, []*Float{}, 123.456},
		{"Nil Float with no default", nilFloat, []*Float{}, 0.0},
		{"Nil Float with default", nilFloat, []*Float{defaultFloat}, 999.999},
		{"Nil Float with nil default", nilFloat, []*Float{nil}, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummy := NewFloat(0.0) // Just needed to call the method
			var result *Float
			if len(tt.default_) == 0 {
				result = dummy.Safe(tt.input)
			} else {
				result = dummy.Safe(tt.input, tt.default_...)
			}

			actual, _ := strconv.ParseFloat(result.String(), 64)

			// Use a small epsilon for floating point comparison
			epsilon := 0.0000001
			if actual < tt.expected-epsilon || actual > tt.expected+epsilon {
				t.Errorf("Safe(%v, %v) = %s, expected %v",
					tt.input, tt.default_, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestToFloat tests the ToFloat function for converting a bigNumber.Float back to a standard big.Float.
** This validates that Float values can be correctly converted back to the standard library type:
** - With valid Float objects
** - With nil Float pointers (should return a zero big.Float)
**
** This conversion is important for interoperability with standard library functions.
**************************************************************************************************/
func TestToFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    *Float
		expected float64
	}{
		{"Valid Float", NewFloat(123.456), 123.456},
		{"Nil Float", nil, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToFloat(tt.input)
			actual, _ := result.Float64()

			// Use a small epsilon for floating point comparison
			epsilon := 0.0000001
			if actual < tt.expected-epsilon || actual > tt.expected+epsilon {
				t.Errorf("ToFloat(%v) = %v, expected %v", tt.input, actual, tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestFloatSetMode tests the SetMode method of the Float type.
** This validates that precision mode can be correctly set for Float calculations.
**
** Setting appropriate precision is important for financial calculations where high precision
** may be required.
**************************************************************************************************/
func TestFloatSetMode(t *testing.T) {
	// Test setting different precision modes
	f := NewFloat(1.0)
	f.SetMode(big.ToNearestEven)

	// Verify we can chain operations after setting mode
	result := f.Div(NewFloat(2.0), NewFloat(3.0))

	// Check the result is reasonable (2.0/3.0 ≈ 0.6666...)
	actual, _ := strconv.ParseFloat(result.String(), 64)
	expected := 0.6666666666666666

	// Use a small epsilon for floating point comparison
	epsilon := 0.0000001
	if actual < expected-epsilon || actual > expected+epsilon {
		t.Errorf("SetMode and calculation = %v, expected approximately %v", actual, expected)
	}
}
