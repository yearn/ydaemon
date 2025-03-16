package addresses

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

/**************************************************************************************************
** TestToAddress validates the ToAddress function by testing conversions from different types:
** - string representation of addresses
** - common.MixedcaseAddress objects
** - common.Address objects
** - invalid types
**
** These tests ensure that address conversion works correctly for all supported formats and
** that it handles edge cases appropriately.
**************************************************************************************************/
func TestToAddress(t *testing.T) {
	// Test data
	validAddrHex := "0x6B175474E89094C44Da98b954EedeAC495271d0F" // DAI token address
	validAddr := common.HexToAddress(validAddrHex)
	mixedCaseAddr := common.NewMixedcaseAddress(validAddr)

	// Test cases
	tests := []struct {
		name     string
		input    interface{}
		expected common.Address
	}{
		{
			name:     "Convert from string",
			input:    validAddrHex,
			expected: validAddr,
		},
		{
			name:     "Convert from lowercase string",
			input:    "0x6b175474e89094c44da98b954eedeac495271d0f",
			expected: validAddr,
		},
		{
			name:     "Convert from MixedcaseAddress",
			input:    mixedCaseAddr,
			expected: validAddr,
		},
		{
			name:     "Convert from Address",
			input:    validAddr,
			expected: validAddr,
		},
		{
			name:     "Convert from empty string",
			input:    "",
			expected: common.HexToAddress("0x0"),
		},
		{
			name:     "Convert from invalid type",
			input:    123, // integer is not a valid address type
			expected: common.Address{},
		},
		{
			name:     "Convert from nil",
			input:    nil,
			expected: common.Address{},
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToAddress(tt.input)
			if result != tt.expected {
				t.Errorf("ToAddress(%v) = %v, expected %v", tt.input, result.Hex(), tt.expected.Hex())
			}
		})
	}
}

/**************************************************************************************************
** TestEquals validates the Equals function by comparing various combinations of addresses:
** - Same address in different formats
** - Different addresses
** - Edge cases like empty or zero addresses
**
** These tests ensure that address equality checks work correctly regardless of the input format.
**************************************************************************************************/
func TestEquals(t *testing.T) {
	// Test data
	addr1Hex := "0x6B175474E89094C44Da98b954EedeAC495271d0F" // DAI token address
	addr1 := common.HexToAddress(addr1Hex)
	addr1Mixed := common.NewMixedcaseAddress(addr1)

	addr2Hex := "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984" // UNI token address
	addr2 := common.HexToAddress(addr2Hex)

	zeroAddr := common.Address{}

	// Test cases
	tests := []struct {
		name     string
		a        interface{}
		b        interface{}
		expected bool
	}{
		{
			name:     "Same address - both strings",
			a:        addr1Hex,
			b:        addr1Hex,
			expected: true,
		},
		{
			name:     "Same address - string and Address",
			a:        addr1Hex,
			b:        addr1,
			expected: true,
		},
		{
			name:     "Same address - string and MixedcaseAddress",
			a:        addr1Hex,
			b:        addr1Mixed,
			expected: true,
		},
		{
			name:     "Same address - Address and MixedcaseAddress",
			a:        addr1,
			b:        addr1Mixed,
			expected: true,
		},
		{
			name:     "Same address - case difference",
			a:        addr1Hex,
			b:        "0x6b175474e89094c44da98b954eedeac495271d0f", // lowercase
			expected: true,
		},
		{
			name:     "Different addresses - both strings",
			a:        addr1Hex,
			b:        addr2Hex,
			expected: false,
		},
		{
			name:     "Different addresses - mixed types",
			a:        addr1,
			b:        addr2,
			expected: false,
		},
		{
			name:     "Zero address comparison",
			a:        "0x0000000000000000000000000000000000000000",
			b:        zeroAddr,
			expected: true,
		},
		{
			name:     "Empty string and zero address",
			a:        "",
			b:        zeroAddr,
			expected: true,
		},
		{
			name:     "Invalid type and zero address",
			a:        123, // integer is not a valid address type
			b:        zeroAddr,
			expected: true, // Both convert to zero address
		},
		{
			name:     "Two invalid types",
			a:        123,
			b:        456,
			expected: true, // Both convert to zero address
		},
		{
			name:     "Nil and zero address",
			a:        nil,
			b:        zeroAddr,
			expected: true, // Nil converts to zero address
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Equals(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equals(%v, %v) = %v, expected %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

/**************************************************************************************************
** TestToAddressReflectType validates that ToAddress correctly handles different reflect types.
** This test specifically checks that the function properly identifies and processes different
** types according to their reflection characteristics.
**************************************************************************************************/
func TestToAddressReflectType(t *testing.T) {
	// Test nil values
	result := ToAddress(nil)
	if result != (common.Address{}) {
		t.Errorf("ToAddress(nil) should return empty address, got %s", result.Hex())
	}

	// Test pointer types that are valid
	addrValue := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	addrPtr := &addrValue
	result = ToAddress(addrPtr)
	if result == (common.Address{}) {
		t.Errorf("ToAddress should handle pointer to Address, got zero address")
	}

	// Test special handling for unsupported pointer types
	var someInt int = 123
	result = ToAddress(&someInt)
	if result != (common.Address{}) {
		t.Errorf("ToAddress(pointer to int) should return empty address, got %s", result.Hex())
	}

	// Test handling of interface values
	var iface interface{} = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	expected := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	result = ToAddress(iface)
	if result != expected {
		t.Errorf("ToAddress(interface with string) = %v, expected %v", result.Hex(), expected.Hex())
	}

	// Test with nil pointers to common types
	var nilAddrPtr *common.Address
	result = ToAddress(nilAddrPtr)
	if result != (common.Address{}) {
		t.Errorf("ToAddress(nil *common.Address) should return empty address, got %s", result.Hex())
	}

	var nilStringPtr *string
	result = ToAddress(nilStringPtr)
	if result != (common.Address{}) {
		t.Errorf("ToAddress(nil *string) should return empty address, got %s", result.Hex())
	}
}

/**************************************************************************************************
** TestEqualsEdgeCases tests additional edge cases for the Equals function to ensure robust
** behavior in all scenarios, particularly focusing on combinations of edge cases and
** unusual input types.
**************************************************************************************************/
func TestEqualsEdgeCases(t *testing.T) {
	// Test with one valid address and one nil
	validAddr := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	if Equals(validAddr, nil) {
		t.Errorf("Equals(validAddr, nil) should be false")
	}

	// Test with a valid address and an empty string (converts to zero address)
	if Equals(validAddr, "") {
		t.Errorf("Equals(validAddr, \"\") should be false")
	}

	// Test with pointer types
	addr1 := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	addr2 := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	if !Equals(&addr1, &addr2) {
		t.Errorf("Equals(&addr1, &addr2) should be true for same address")
	}

	// Test with invalid hex string
	invalidHex := "0xNOTHEX"
	zeroAddr := common.Address{}
	if !Equals(invalidHex, zeroAddr) {
		t.Errorf("Equals(invalidHex, zeroAddr) should be true as invalid hex becomes zero address")
	}
}
