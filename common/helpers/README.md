# Helpers Package

> The helpers package provides a comprehensive collection of utility functions used throughout
> yDaemon for common operations such as array manipulation, type conversion, address validation,
> and numeric calculations. These functions abstract away repetitive tasks and provide consistent
> behavior across the application.

## Core Components

The helpers package is organized into several functional categories:

### Array and Collection Utilities

Functions for manipulating arrays and collections:

- **Deduplication**: `UniqueArrayAddress` removes duplicates from arrays
- **Filtering**: `RemoveFromArray` creates a new array without specified elements
- **Searching**: `Contains`, `ContainsSubString`, `EndsWithSubstring`, `Intersects`
- **Transformation**: `ToLower` converts all strings in an array to lowercase

```go
// Remove blacklisted addresses from an array
filteredAddresses := helpers.RemoveFromArray(addresses, blacklistedAddress)

// Check if arrays have any common elements
hasCommonTags := helpers.Intersects(vaultTags, userPreferences)
```

### Blockchain Utilities

Functions for working with Ethereum addresses and blockchain data:

- **Address Validation**: `AssertAddress`, `AddressIsValid`
- **Chain ID Handling**: `AssertChainID`
- **Address Conversion**: `AddressToString` converts addresses to hex strings

```go
// Validate an address for a specific chain
address, isValid := helpers.AssertAddress(addressStr, chainID)
if !isValid {
    // Handle invalid address
}
```

### Type Conversion and Formatting

Functions for safely converting between different data types:

- **String Conversion**: `FormatUint64`, `SafeString`, `SafeStringToUint64`, `StringToBool`
- **Interface Decoding**: `DecodeString`, `DecodeBigInt`, `DecodeAddress`, etc.
- **Token Amount Formatting**: `FormatAmount`, `GetHumanizedValue`

```go
// Safely convert a string to uint64 with a default fallback
blockNumber := helpers.SafeStringToUint64(blockStr, 0)

// Convert addresses to string representation
addressStrings := helpers.AddressToString(addressArray)
```

### Token Amount Handling

Functions for handling token amounts with proper decimal precision:

- **Normalization**: `ToNormalizedAmount`, `ToNormalizedFloat`, `ToNormalizedValue`
- **Raw Value Conversion**: `ToRawAmount`

```go
// Convert raw token amount to human-readable form
normalizedAmount := helpers.ToNormalizedAmount(tokenAmount, 18)

// Calculate USD value of token amount
usdValue := helpers.ToNormalizedValue(tokenAmount, tokenPrice, 18)
```

### File Operations

Functions for working with files and directories:

- **Directory Reading**: `ReadAllFilesInDir` reads files matching a specific suffix

```go
// Read all JSON files in a directory
jsonFiles, filenames, err := helpers.ReadAllFilesInDir("./data/", ".json")
```

### Contract Call Decoding

Functions for safely extracting data from contract call results:

- **Type-Specific Decoders**: `DecodeString`, `DecodeUint64`, `DecodeBigInt`, `DecodeBool`, `DecodeAddress`
- **Collection Decoders**: `DecodeUint16s`, `DecodeBigInts`, `DecodeAddresses`

```go
// Extract address from contract call result
vaultAddress := helpers.DecodeAddress(callResult)

// Extract array of big integers from contract call result
tokenAmounts := helpers.DecodeBigInts(callResult)
```

## Design Principles

The helpers package follows several key design principles:

1. **Type Safety**: Uses Go generics for type-safe operations where appropriate
2. **Default Values**: Provides sensible defaults when conversions fail
3. **Error Handling**: Gracefully handles edge cases to prevent panics
4. **Immutability**: Functions return new values rather than modifying inputs
5. **Consistency**: Follows consistent naming and behavior patterns

## HTTP Utilities

For HTTP-related utilities, the package provides specialized functions:

- **JSON Fetching**: `FetchJSON` and `FetchJSONWithReject` provide different error handling strategies
- **GraphQL Queries**: Pre-formatted query builders for The Graph protocol endpoints

## Usage Examples

### Working with Token Amounts

```go
// Convert a raw token amount (e.g., wei) to a normalized form (e.g., ETH)
rawAmount := bigNumber.NewInt(1000000000000000000) // 1 ETH in wei
normalizedAmount := helpers.ToNormalizedAmount(rawAmount, 18)
fmt.Println(normalizedAmount) // Outputs: 1.0

// Calculate the USD value of tokens using price data
tokenAmount := bigNumber.NewInt(5000000000000000000) // 5 tokens with 18 decimals
tokenPrice := bigNumber.NewInt(2000000) // $2.00 with 6 decimals for price
usdValue := helpers.ToNormalizedValue(tokenAmount, tokenPrice, 18)
fmt.Println(usdValue) // Outputs: 10.0 (5 tokens * $2.00)
```

### Safe Handling of External Data

```go
// Safely handle potentially missing values with defaults
blockNumber := helpers.SafeStringToUint64(queryParams["block"], 0)
searchTerm := helpers.SafeString(queryParams["search"], "")

// Validate chain ID and address inputs
chainID, validChain := helpers.AssertChainID(chainIDStr)
if !validChain {
    // Handle invalid chain ID
}

address, validAddress := helpers.AssertAddress(addressStr, chainID)
if !validAddress {
    // Handle invalid address
}
```

### Array Operations

```go
// Remove duplicates from an array of addresses
uniqueAddresses := helpers.UniqueArrayAddress(addresses)

// Filter an array to remove specific elements
activeStrategies := helpers.RemoveFromArray(allStrategies, deprecatedStrategy)

// Check if an array contains a specific element
isSupported := helpers.Contains(supportedChains, chainID)
```

This package serves as a foundation for common operations throughout the yDaemon codebase, promoting code reuse and ensuring consistent behavior across the application.
