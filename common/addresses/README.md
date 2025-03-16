# Addresses Package

## Overview

The `addresses` package provides utility functions for handling Ethereum addresses across the yDaemon application. It standardizes address operations, making it easier to work with different address formats and ensuring consistent behavior throughout the codebase.

This package is particularly useful when working with addresses from different sources (user input, blockchain data, contract events) that may come in varying formats but need to be treated uniformly.

## Core Components

### Address Conversion

The package provides a universal conversion function for different address representations:

- **ToAddress**: Converts various input types to a standard `common.Address` type, handling:
    - Hex strings (e.g., "0x1234...")
    - `common.MixedcaseAddress` instances (case-sensitive Ethereum addresses)
    - `common.Address` values (passed through unchanged)

### Address Comparison

The package offers a type-agnostic way to compare addresses:

- **Equals**: Checks if two addresses are the same, regardless of their original format, by converting them to standard addresses first and comparing their hex representations

## Usage Examples

### Converting Different Address Formats

```go
import (
    "github.com/yearn/ydaemon/common/addresses"
    "github.com/ethereum/go-ethereum/common"
)

// From a string
stringAddr := "0x6b175474e89094c44da98b954eedeac495271d0f" // DAI token
standardAddr := addresses.ToAddress(stringAddr)

// From a MixedcaseAddress
mixedCaseAddr := common.NewMixedcaseAddress(common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"))
standardAddr = addresses.ToAddress(mixedCaseAddr)

// From an existing common.Address
existingAddr := common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
standardAddr = addresses.ToAddress(existingAddr)
```

### Comparing Addresses in Different Formats

```go
import (
    "fmt"
    "github.com/yearn/ydaemon/common/addresses"
)

// Different representations of the same address
addr1 := "0x6b175474e89094c44da98b954eedeac495271d0f"
addr2 := "0x6B175474E89094C44Da98b954EedeAC495271d0F" // Different case

// Check equality despite format differences
if addresses.Equals(addr1, addr2) {
    fmt.Println("Addresses are the same")
}

// Compare addresses from different sources
userInputAddr := userInput // String from UI
contractAddr := someContractEvent.Address // common.Address from contract event

if addresses.Equals(userInputAddr, contractAddr) {
    fmt.Println("User entered the contract address correctly")
}
```

## Integration with Other Packages

The `addresses` package is commonly used alongside:

- **ethereum**: For working with blockchain contracts and RPC calls that require address parameters
- **helpers**: For additional validation and conversion operations
- **models**: For populating address fields in structured data types

## Error Handling

The `ToAddress` function logs a warning when it encounters an unsupported type, but doesn't return an error. Instead, it returns an empty `common.Address{}`, which evaluates to the zero address (0x0000...0000). Code using this function should check for this case if it's critical to distinguish between valid and invalid addresses.

## Best Practices

1. **Always Use Package Functions**: Instead of direct type conversions or string comparisons, use this package's functions to ensure consistent handling across the codebase.

2. **Case Sensitivity**: Remember that while address comparisons via `Equals()` are case-insensitive, Ethereum addresses displayed to users should preserve their checksum case for improved readability and error detection.

3. **Type Safety**: When working with known address types, it's more efficient to use direct type assertions or conversions. Use this package primarily when dealing with addresses from external sources or with uncertain types.

4. **Error Checking**: The `ToAddress` function returns the zero address (0x0000...0000) for invalid inputs. Consider checking for this case when processing addresses from untrusted sources.

```go
addr := addresses.ToAddress(userInput)
if addr == (common.Address{}) {
    // Handle invalid address
}
```

## Beyond Basic Addresses

This package focuses on core address conversion and comparison. For more advanced address operations, consider using complementary functions from other packages:

- `helpers.AssertAddress` - Validates that a string is a properly formatted Ethereum address
- `helpers.AddressIsValid` - Checks if an address is valid for a specific chain (non-zero and not blacklisted)
- `ethereum.IsContract` - Determines if an address contains contract code
