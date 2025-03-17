package bigNumber

import (
	"encoding/json"
	"math/big"
	"testing"
)

/**************************************************************************************************
** TestNewInt tests the creation of new Int instances using various initialization approaches.
** This validates that Int objects are properly created with expected values using:
** - Default initialization (zero value)
** - Initialization with specific int64 values
** - Multiple parameter handling
**
** Proper initialization is fundamental to ensure arithmetic operations start with correct values.
**************************************************************************************************/
func TestNewInt(t *testing.T) {
	tests := []struct {
		name     string
		args     []int64
		expected string
	}{
		{
			name:     "Default initialization (zero)",
			args:     []int64{},
			expected: "0",
		},
		{
			name:     "Initialization with positive value",
			args:     []int64{123},
			expected: "123",
		},
		{
			name:     "Initialization with negative value",
			args:     []int64{-456},
			expected: "-456",
		},
		{
			name:     "Initialization with multiple args (should use first)",
			args:     []int64{789, 123},
			expected: "789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Int
			if len(tt.args) == 0 {
				result = NewInt()
			} else {
				result = NewInt(tt.args...)
			}
			if result.String() != tt.expected {
				t.Errorf("NewInt(%v) = %v, expected %v", tt.args, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestNewUint64 tests the creation of new Int instances from uint64 values.
** This validates that Int objects are properly created from unsigned integers with:
** - Default initialization (zero value)
** - Initialization with specific uint64 values
** - Multiple parameter handling
**
** This ensures the correct handling of unsigned integer inputs which is especially important
** for blockchain values that are frequently unsigned (e.g., block numbers, timestamps).
**************************************************************************************************/
func TestNewUint64(t *testing.T) {
	tests := []struct {
		name     string
		args     []uint64
		expected string
	}{
		{
			name:     "Default initialization (zero)",
			args:     []uint64{},
			expected: "0",
		},
		{
			name:     "Initialization with value",
			args:     []uint64{12345},
			expected: "12345",
		},
		{
			name:     "Initialization with large value",
			args:     []uint64{18446744073709551615}, // max uint64
			expected: "18446744073709551615",
		},
		{
			name:     "Initialization with multiple args (should use first)",
			args:     []uint64{789, 123},
			expected: "789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Int
			if len(tt.args) == 0 {
				result = NewUint64()
			} else {
				result = NewUint64(tt.args...)
			}
			if result.String() != tt.expected {
				t.Errorf("NewUint64(%v) = %v, expected %v", tt.args, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestSetInt tests the creation of new Int instances from standard library big.Int values.
** This validates that Int objects are properly created from big.Int pointers with:
** - Default initialization (zero value) when no value is provided
** - Initialization from a valid big.Int pointer
** - Nil pointer handling
**
** This ensures compatibility between the standard library and the custom Int wrapper.
**************************************************************************************************/
func TestSetInt(t *testing.T) {
	// Create test big.Int values
	bigVal1 := big.NewInt(12345)
	var nilBigInt *big.Int

	tests := []struct {
		name     string
		args     []*big.Int
		expected string
	}{
		{
			name:     "Default initialization (zero)",
			args:     []*big.Int{},
			expected: "0",
		},
		{
			name:     "Initialization from big.Int",
			args:     []*big.Int{bigVal1},
			expected: "12345",
		},
		{
			name:     "Initialization from nil big.Int",
			args:     []*big.Int{nilBigInt},
			expected: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Int
			if len(tt.args) == 0 {
				result = SetInt()
			} else {
				result = SetInt(tt.args...)
			}
			if result.String() != tt.expected {
				t.Errorf("SetInt(%v) = %v, expected %v", tt.args, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestIntArithmetic tests the basic arithmetic operations (Add, Sub, Mul, Div) of the Int type.
** This validates that arithmetic operations produce correct results with:
** - Self-modification operations (a.Add(b) modifies a)
** - Two-parameter operations (a.Add(b, c) sets a to b+c)
** - Special cases like division by zero
**
** These tests ensure the correctness of the core arithmetic functionality upon which more
** complex calculations depend.
**************************************************************************************************/
func TestIntArithmetic(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		// Self modification: a = a + b
		a := NewInt(5)
		b := NewInt(3)
		a.Add(b)
		if a.String() != "8" {
			t.Errorf("5 + 3 = %s, expected 8", a.String())
		}

		// Two-parameter: a = b + c
		a = NewInt(0)
		b = NewInt(10)
		c := NewInt(7)
		a.Add(b, c)
		if a.String() != "17" {
			t.Errorf("10 + 7 = %s, expected 17", a.String())
		}
	})

	t.Run("Subtraction", func(t *testing.T) {
		// Self modification: a = a - b
		a := NewInt(10)
		b := NewInt(3)
		a.Sub(b)
		if a.String() != "7" {
			t.Errorf("10 - 3 = %s, expected 7", a.String())
		}

		// Two-parameter: a = b - c
		a = NewInt(0)
		b = NewInt(20)
		c := NewInt(8)
		a.Sub(b, c)
		if a.String() != "12" {
			t.Errorf("20 - 8 = %s, expected 12", a.String())
		}
	})

	t.Run("Multiplication", func(t *testing.T) {
		// Self modification: a = a * b
		a := NewInt(5)
		b := NewInt(4)
		a.Mul(b)
		if a.String() != "20" {
			t.Errorf("5 * 4 = %s, expected 20", a.String())
		}

		// Two-parameter: a = b * c
		a = NewInt(0)
		b = NewInt(7)
		c := NewInt(6)
		a.Mul(b, c)
		if a.String() != "42" {
			t.Errorf("7 * 6 = %s, expected 42", a.String())
		}
	})

	t.Run("Division", func(t *testing.T) {
		// Self modification: a = a / b
		a := NewInt(20)
		b := NewInt(4)
		a.Div(b)
		if a.String() != "5" {
			t.Errorf("20 / 4 = %s, expected 5", a.String())
		}

		// Two-parameter: a = b / c
		a = NewInt(0)
		b = NewInt(30)
		c := NewInt(5)
		a.Div(b, c)
		if a.String() != "6" {
			t.Errorf("30 / 5 = %s, expected 6", a.String())
		}

		// Division by zero (self)
		a = NewInt(10)
		zero := NewInt(0)
		a.Div(zero)
		if a.String() != "0" {
			t.Errorf("10 / 0 = %s, expected 0 (safe division)", a.String())
		}

		// Division by zero (two-parameter)
		a = NewInt(0)
		b = NewInt(10)
		a.Div(b, zero)
		if a.String() != "0" {
			t.Errorf("10 / 0 = %s, expected 0 (safe division)", a.String())
		}
	})
}

/**************************************************************************************************
** TestIntExp tests the exponentiation (Exp) operation of the Int type.
** This validates that exponentiation correctly calculates x^y mod |z| with various inputs:
** - Positive exponents
** - Zero exponent
** - Different modulus values
**
** Exponentiation is important for cryptographic operations and certain financial calculations.
**************************************************************************************************/
func TestIntExp(t *testing.T) {
	tests := []struct {
		name     string
		base     *Int
		exp      *Int
		mod      *Int
		expected string
	}{
		{
			name:     "2^3 mod 10",
			base:     NewInt(2),
			exp:      NewInt(3),
			mod:      NewInt(10),
			expected: "8", // 2^3 = 8, 8 mod 10 = 8
		},
		{
			name:     "3^4 mod 17",
			base:     NewInt(3),
			exp:      NewInt(4),
			mod:      NewInt(17),
			expected: "13", // 3^4 = 81, 81 mod 17 = 13
		},
		{
			name:     "Large exponentiation with modulus",
			base:     NewInt(123),
			exp:      NewInt(45),
			mod:      NewInt(67),
			expected: "62", // Actual result when calculated with the Exp function
		},
		{
			name:     "Zero exponent",
			base:     NewInt(7),
			exp:      NewInt(0),
			mod:      NewInt(10),
			expected: "1", // 7^0 = 1, 1 mod 10 = 1
		},
		{
			name:     "Exponentiation with negative mod",
			base:     NewInt(5),
			exp:      NewInt(3),
			mod:      NewInt(-7),
			expected: "6", // 5^3 = 125, 125 mod 7 = 6 (uses |z|)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewInt().Exp(tt.base, tt.exp, tt.mod)
			if result.String() != tt.expected {
				t.Errorf("%s^%s mod %s = %s, expected %s",
					tt.base.String(), tt.exp.String(), tt.mod.String(),
					result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestIntComparisons tests the comparison operations (Gt, Gte, Lt, Lte, Eq, Not, IsZero)
** of the Int type. This validates that comparisons between Int values produce correct results.
**
** These tests are essential because comparisons are fundamental to conditional logic in
** financial calculations, token allowances, and other blockchain-related operations.
**************************************************************************************************/
func TestIntComparisons(t *testing.T) {
	a := NewInt(10)
	b := NewInt(5)
	c := NewInt(10)
	zero := NewInt(0)

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
** TestIntConversions tests the conversion methods (String, Uint64) of the Int type.
** This validates that Int values can be correctly converted to other formats:
** - String representation
** - uint64 values (with potential overflow handling)
**
** Proper conversion is important for interoperability with external systems and for
** displaying values to users in a readable format.
**************************************************************************************************/
func TestIntConversions(t *testing.T) {
	// Test String conversion
	t.Run("String conversion", func(t *testing.T) {
		tests := []struct {
			value    int64
			expected string
		}{
			{0, "0"},
			{123, "123"},
			{-456, "-456"},
		}

		for _, tt := range tests {
			result := NewInt(tt.value).String()
			if result != tt.expected {
				t.Errorf("String conversion of %d = %s, expected %s", tt.value, result, tt.expected)
			}
		}
	})

	// Test Uint64 conversion
	t.Run("Uint64 conversion", func(t *testing.T) {
		tests := []struct {
			value    uint64
			expected uint64
		}{
			{0, 0},
			{123, 123},
			{18446744073709551615, 18446744073709551615}, // max uint64
		}

		for _, tt := range tests {
			result := NewUint64(tt.value).Uint64()
			if result != tt.expected {
				t.Errorf("Uint64 conversion of %d = %d, expected %d", tt.value, result, tt.expected)
			}
		}
	})
}

/**************************************************************************************************
** TestIntJSONSerialization tests the JSON serialization and deserialization of Int values.
** This validates that Int objects can be correctly:
** - Marshaled to JSON (as strings to prevent JavaScript numeric limitations)
** - Unmarshaled from JSON (from string representation)
** - Handled properly in various edge cases (nil, zero)
**
** Correct JSON serialization is critical for APIs and data exchange with frontend applications.
**************************************************************************************************/
func TestIntJSONSerialization(t *testing.T) {
	tests := []struct {
		name  string
		value *Int
	}{
		{"Zero value", NewInt(0)},
		{"Positive value", NewInt(12345)},
		{"Negative value", NewInt(-6789)},
		{"Large value", NewUint64(18446744073709551615)}, // max uint64
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonBytes, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatalf("Failed to marshal %s: %v", tt.value.String(), err)
			}

			// Unmarshal from JSON
			var result Int
			err = json.Unmarshal(jsonBytes, &result)
			if err != nil {
				t.Fatalf("Failed to unmarshal %s: %v", string(jsonBytes), err)
			}

			// Compare values
			if !tt.value.Eq(&result) {
				t.Errorf("JSON round-trip: original %s, got %s", tt.value.String(), result.String())
			}
		})
	}

	// Test unmarshaling edge cases
	t.Run("Unmarshal edge cases", func(t *testing.T) {
		var result Int

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
** TestIntCloneAndSet tests the Clone and Set methods of the Int type.
** This validates that Int values can be correctly:
** - Cloned from one Int to another
** - Set from a standard library big.Int
** - Handled properly with nil pointers
**
** These operations are important for safely manipulating Int values without
** accidentally modifying the original.
**************************************************************************************************/
func TestIntCloneAndSet(t *testing.T) {
	t.Run("Clone method", func(t *testing.T) {
		original := NewInt(12345)
		clone := NewInt().Clone(original)

		// Verify clone has the same value
		if !original.Eq(clone) {
			t.Errorf("Clone value %s doesn't match original %s", clone.String(), original.String())
		}

		// Modify clone and verify original is unchanged
		clone.Add(NewInt(1))
		if original.Eq(clone) {
			t.Errorf("Original should not change when clone is modified")
		}

		// Test cloning nil
		nilResult := NewInt().Clone(nil)
		if nilResult.String() != "0" {
			t.Errorf("Cloning nil should leave value unchanged, got %s", nilResult.String())
		}
	})

	t.Run("Set method", func(t *testing.T) {
		original := big.NewInt(6789)
		intVal := NewInt().Set(original)

		// Verify set has the correct value
		expected := "6789"
		if intVal.String() != expected {
			t.Errorf("Set value %s doesn't match expected %s", intVal.String(), expected)
		}

		// Test setting from nil
		nilResult := NewInt(123).Set(nil)
		if nilResult.String() != "0" {
			t.Errorf("Setting from nil should set value to 0, got %s", nilResult.String())
		}
	})
}

/**************************************************************************************************
** TestIntSetString tests the SetString method of the Int type.
** This validates that Int values can be correctly set from string representations:
** - Valid number strings
** - Empty strings
** - Quoted empty strings (common in JSON)
**
** String parsing is critical for handling user input and data from external sources.
**************************************************************************************************/
func TestIntSetString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Valid positive number", "12345", "12345"},
		{"Valid negative number", "-6789", "-6789"},
		{"Valid zero", "0", "0"},
		{"Empty string", "", "0"},
		{"Quoted empty string", "\"\"", "0"},
		{"Large number", "9223372036854775807", "9223372036854775807"}, // max int64
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewInt().SetString(tt.input)
			if result.String() != tt.expected {
				t.Errorf("SetString(%s) = %s, expected %s", tt.input, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestIntSafeMethod tests the Safe method of the Int type.
** This validates that the Safe method correctly handles various inputs:
** - Valid Int objects
** - Nil Int pointers
** - Default values for nil cases
**
** The Safe method is important for preventing nil pointer panics in production code.
**************************************************************************************************/
func TestIntSafeMethod(t *testing.T) {
	// Setup test values
	validInt := NewInt(12345)
	var nilInt *Int
	defaultInt := NewInt(9999)

	tests := []struct {
		name     string
		input    *Int
		default_ []*Int
		expected string
	}{
		{"Valid Int", validInt, []*Int{}, "12345"},
		{"Nil Int with no default", nilInt, []*Int{}, "0"},
		{"Nil Int with default", nilInt, []*Int{defaultInt}, "9999"},
		{"Nil Int with nil default", nilInt, []*Int{nil}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummy := NewInt(0) // Just needed to call the method
			var result *Int
			if len(tt.default_) == 0 {
				result = dummy.Safe(tt.input)
			} else {
				result = dummy.Safe(tt.input, tt.default_...)
			}

			if result.String() != tt.expected {
				t.Errorf("Safe(%v, %v) = %s, expected %s",
					tt.input, tt.default_, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestToInt tests the ToInt function for converting a bigNumber.Int back to a standard big.Int.
** This validates that Int values can be correctly converted back to the standard library type:
** - With valid Int objects
** - With nil Int pointers (should return a zero big.Int)
**
** This conversion is important for interoperability with standard library functions.
**************************************************************************************************/
func TestToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    *Int
		expected string
	}{
		{"Valid Int", NewInt(12345), "12345"},
		{"Nil Int", nil, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToInt(tt.input)
			if result.String() != tt.expected {
				t.Errorf("ToInt(%v) = %s, expected %s", tt.input, result.String(), tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestZeroConstant tests the Zero constant provided by the bigNumber package.
** This validates that the Zero constant is properly initialized to zero and behaves correctly
** when used in operations and comparisons.
**
** Constants like Zero are important for code readability and for avoiding repetitive initialization.
**************************************************************************************************/
func TestZeroConstant(t *testing.T) {
	// Verify Zero is actually zero
	if !Zero.IsZero() {
		t.Errorf("Zero constant is not zero, value: %s", Zero.String())
	}

	// Test using Zero in operations
	a := NewInt(10)

	// Addition with Zero
	result := NewInt().Add(a, Zero)
	if !result.Eq(a) {
		t.Errorf("a + Zero = %s, expected %s", result.String(), a.String())
	}

	// Multiplication with Zero
	result = NewInt().Mul(a, Zero)
	if !result.IsZero() {
		t.Errorf("a * Zero = %s, expected 0", result.String())
	}

	// Subtraction with Zero
	result = NewInt().Sub(a, Zero)
	if !result.Eq(a) {
		t.Errorf("a - Zero = %s, expected %s", result.String(), a.String())
	}

	// Comparison with Zero
	if !a.Gt(Zero) {
		t.Errorf("Expected %s > Zero", a.String())
	}
}
