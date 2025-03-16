package addresses

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The addresses package provides utility functions for working with Ethereum addresses across
** the application. It offers a standardized way to handle different address formats and perform
** common address operations, helping maintain consistency across the codebase.
**
** This package abstracts away the complexities of address type conversions and comparisons,
** making it easier to work with addresses regardless of their original format.
**************************************************************************************************/

/**************************************************************************************************
** ToAddress converts various address formats to a standard Ethereum address (common.Address).
** It handles multiple input types:
** - String: Converts hex strings to addresses
** - common.MixedcaseAddress: Extracts the underlying address
** - common.Address: Returns the address unchanged
** - *common.Address: Dereferences the pointer to return the address
**
** If an unsupported type is provided, a warning is logged and an empty address is returned.
**
** @param address The input to convert, which can be a string, MixedcaseAddress, or Address
** @return common.Address The standardized Ethereum address
**************************************************************************************************/
func ToAddress(address any) common.Address {
	// Handle nil input first
	if address == nil {
		logs.Warning(`nil address provided`)
		return common.Address{}
	}

	valueType := reflect.TypeOf(address)

	// Handle pointers
	if valueType.Kind() == reflect.Pointer {
		// Handle nil pointers to avoid panic when dereferencing
		if reflect.ValueOf(address).IsNil() {
			logs.Warning(`nil pointer address provided`)
			return common.Address{}
		}

		// Check for pointer to common.Address
		if valueType.Elem().String() == "common.Address" {
			return *(address.(*common.Address))
		}

		// For other pointer types, log warning and return empty address
		logs.Warning(`unsupported pointer type`, valueType.String())
		return common.Address{}
	}

	// Handle direct types
	if valueType.Kind() == reflect.String {
		return common.HexToAddress(address.(string))
	}
	if valueType.String() == `common.MixedcaseAddress` {
		addr := address.(common.MixedcaseAddress)
		return addr.Address()
	}
	if valueType.String() == `common.Address` {
		return address.(common.Address)
	}

	logs.Warning(`unknown type`, valueType.String())
	return common.Address{}
}

/**************************************************************************************************
** Equals compares two Ethereum addresses for equality, regardless of their input format.
** This function converts both inputs to standard Ethereum addresses using ToAddress(),
** then performs a case-insensitive comparison of their hexadecimal representations.
**
** This is useful for comparing addresses that might be in different formats but represent
** the same underlying Ethereum address.
**
** @param a The first address to compare (can be string, MixedcaseAddress, or Address)
** @param b The second address to compare (can be string, MixedcaseAddress, or Address)
** @return bool True if the addresses are equal, false otherwise
**************************************************************************************************/
func Equals(a, b any) bool {
	return ToAddress(a).Hex() == ToAddress(b).Hex()
}
