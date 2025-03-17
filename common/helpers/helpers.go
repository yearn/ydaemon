package helpers

import (
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The helpers package provides a collection of utility functions used throughout yDaemon for
** common operations such as array manipulation, type conversion, address validation, and numeric
** calculations. These functions abstract away repetitive tasks and provide consistent behavior
** across the application.
**
** This package includes utilities for both general-purpose programming needs as well as
** blockchain-specific operations like working with Ethereum addresses and handling token amounts
** with proper decimal precision.
**************************************************************************************************/

/**************************************************************************************************
** Intersects checks if two string arrays have at least one element in common.
** This function is useful for comparing tag lists, categories, or other string-based identifiers.
**
** @param arr1 The first array of strings to check
** @param arr2 The second array of strings to check
** @return bool True if there is at least one common element, false otherwise
**************************************************************************************************/
func Intersects(arr1 []string, arr2 []string) bool {
	for _, v := range arr1 {
		if ContainsSubString(arr2, v) {
			return true
		}
	}
	return false
}

/**************************************************************************************************
** UniqueArrayAddress removes duplicate elements from an array of comparable items.
** This is particularly useful for deduplicating arrays of Ethereum addresses or other identifiers.
**
** The function uses a map to track which elements have been seen, ensuring O(n) time complexity.
**
** @param arr The array to deduplicate
** @return []T A new array containing only the unique elements from the input array
**************************************************************************************************/
func UniqueArrayAddress[T comparable](arr []T) []T {
	occurred := map[T]bool{}
	result := []T{}
	for i := range arr {
		if !occurred[arr[i]] {
			occurred[arr[i]] = true
			result = append(result, arr[i])
		}
	}
	return result
}

/**************************************************************************************************
** RemoveFromArray creates a new array with all occurrences of a specified element removed.
** This function provides a convenient way to filter out unwanted elements from an array.
**
** @param arr The original array to filter
** @param element The element to remove from the array
** @return []T A new array with all occurrences of the element removed
**************************************************************************************************/
func RemoveFromArray[T comparable](arr []T, element T) []T {
	result := []T{}
	for i := range arr {
		if arr[i] != element {
			result = append(result, arr[i])
		}
	}
	return result
}

/**************************************************************************************************
** ReadAllFilesInDir reads the content of all files in a directory that match a specific suffix.
** This function is commonly used to load configuration files, data files, or other resources
** from a directory based on their file extension.
**
** @param directory The path to the directory to read files from
** @param suffix The file suffix to filter by (e.g., ".json", ".csv")
** @return [][]byte An array of file contents as byte arrays
** @return []string An array of filenames corresponding to the contents
** @return error An error if directory reading fails
**************************************************************************************************/
func ReadAllFilesInDir(directory string, suffix string) ([][]byte, []string, error) {
	filenames := []string{}
	dest := [][]byte{}
	outputDirRead, err := os.Open(directory)
	if err != nil {
		return dest, filenames, err
	}
	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		return dest, filenames, err
	}
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]
		outputFileName := outputFileHere.Name()
		if strings.HasSuffix(outputFileName, suffix) {
			content, err := os.ReadFile(directory + outputFileName)
			if err != nil {
				logs.Error(`error`, `impossible to read files in `+directory)
				continue
			}
			filenames = append(filenames, outputFileName)
			dest = append(dest, content)
		}
	}
	return dest, filenames, nil
}

/**************************************************************************************************
** FormatUint64 converts a string to a uint64 with a default fallback value.
** This function safely handles string-to-integer conversion, returning a default value
** if the string is empty or contains invalid numeric data.
**
** @param s The string to convert to a uint64
** @param defaultValue The fallback value to return if conversion fails
** @return uint64 The converted uint64 value or the default value if conversion fails
**************************************************************************************************/
func FormatUint64(s string, defaultValue uint64) uint64 {
	if s == "" {
		return defaultValue
	}
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

/**************************************************************************************************
** FormatAmount converts a raw token amount string to a human-readable float value based on
** the provided decimal places. This is useful for displaying token amounts in user interfaces.
**
** @param amount The raw token amount as a string
** @param decimals The number of decimal places for the token
** @return float64 The humanized amount as a float64
** @return *bigNumber.Float The humanized amount as a precise big.Float
**************************************************************************************************/
func FormatAmount(amount string, decimals int) (float64, *bigNumber.Float) {
	fTotalAssets := bigNumber.NewFloat().SetString(amount)
	fDecimals := bigNumber.NewFloat(math.Pow10(decimals))
	humanizedBalance := bigNumber.NewFloat().Quo(fTotalAssets, fDecimals)
	fhumanizedBalance, _ := humanizedBalance.Float64()
	return fhumanizedBalance, humanizedBalance
}

/**************************************************************************************************
** Contains checks if a specific value exists in an array.
** This generic function works with any comparable type, making it useful for checking
** existence in arrays of strings, integers, addresses, etc.
**
** @param arr The array to search in
** @param value The value to search for
** @return bool True if the value is found in the array, false otherwise
**************************************************************************************************/
func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

/**************************************************************************************************
** EndsWithSubstring checks if a value ends with any of the strings in the provided array.
** This is useful for checking file extensions, URL paths, or other string pattern matching.
**
** @param arr The array of potential suffixes to check
** @param value The string to check against the suffixes
** @return bool True if the value ends with any of the strings in the array
**************************************************************************************************/
func EndsWithSubstring(arr []string, value string) bool {
	for _, v := range arr {
		if strings.HasSuffix(value, v) {
			return true
		}
	}
	return false
}

/**************************************************************************************************
** ContainsSubString checks if a value contains any of the substrings in the provided array.
** This function is useful for keyword matching, filtering, or pattern detection in strings.
**
** @param arr The array of substrings to search for
** @param value The string to check for the presence of substrings
** @return bool True if the value contains any of the substrings in the array
**************************************************************************************************/
func ContainsSubString(arr []string, value string) bool {
	for _, v := range arr {
		if strings.Contains(value, v) {
			return true
		}
	}
	return false
}

/**************************************************************************************************
** AssertChainID validates a chain ID string and converts it to a uint64.
** This function checks if the chain ID is valid and supported by the system.
** Special handling is provided for local development chains (1337 is mapped to 1).
**
** @param chainIDStr The chain ID as a string
** @return uint64 The validated chain ID as a uint64
** @return bool True if the chain ID is valid and supported, false otherwise
**************************************************************************************************/
func AssertChainID(chainIDStr string) (uint64, bool) {
	chainID, err := strconv.ParseUint(chainIDStr, 10, 64)
	if err != nil {
		return 0, false
	}
	if chainID == 1337 {
		return 1, true
	}
	if _, ok := env.GetChain(chainID); !ok {
		return 0, false
	}
	return chainID, true
}

/**************************************************************************************************
** AssertAddress validates an Ethereum address string for a specific chain.
** This function checks if the address is a valid hex address and not blacklisted for the chain.
**
** @param addressStr The address as a string
** @param chainID The chain ID to check the address against
** @return common.Address The validated Ethereum address
** @return bool True if the address is valid and not blacklisted, false otherwise
**************************************************************************************************/
func AssertAddress(addressStr string, chainID uint64) (common.Address, bool) {
	if !common.IsHexAddress(addressStr) {
		return common.Address{}, false
	}
	address := common.HexToAddress(addressStr)
	chain, ok := env.GetChain(chainID)
	if !ok {
		return common.Address{}, false
	}
	if chainID > 0 && Contains(chain.BlacklistedVaults, address) {
		return common.Address{}, false
	}
	return address, true
}

/**************************************************************************************************
** AddressIsValid checks if an Ethereum address is valid and not blacklisted for a specific chain.
** This function complements AssertAddress but works with existing common.Address objects
** instead of address strings.
**
** @param address The Ethereum address to validate
** @param chainID The chain ID to check the address against
** @return bool True if the address is valid and not blacklisted, false otherwise
**************************************************************************************************/
func AddressIsValid(address common.Address, chainID uint64) bool {
	if (address == common.Address{} || addresses.Equals(address, "0x0000000000000000000000000000000000000000")) {
		return false
	}
	chain, ok := env.GetChain(chainID)
	if !ok {
		return false
	}
	if chainID > 0 && Contains(chain.BlacklistedVaults, address) {
		return false
	}
	return true
}

/**************************************************************************************************
** StringToBool converts a string to a boolean value.
** This simple utility function treats the string "true" as true and any other value as false.
**
** @param str The string to convert to a boolean
** @return bool True if the string equals "true", false otherwise
**************************************************************************************************/
func StringToBool(str string) bool {
	return str == "true"
}

/**************************************************************************************************
** GetHumanizedValue calculates a human-readable value of a token amount considering its price.
** This function is useful for displaying token values in fiat or another reference currency.
**
** @param amount The raw token amount
** @param decimals The number of decimal places for the token
** @param price The token price in a reference currency (with 6 decimals)
** @return float64 The humanized value (amount * price) as a float64
**************************************************************************************************/
func GetHumanizedValue(
	amount *bigNumber.Int,
	decimals int,
	price *bigNumber.Int,
) float64 {
	_, humanizedValue := FormatAmount(amount.String(), decimals)
	_, humanizedPrice := FormatAmount(price.String(), 6)
	fHumanizedValue, _ := bigNumber.NewFloat().Mul(humanizedValue, humanizedPrice).Float64()
	return fHumanizedValue
}

/**************************************************************************************************
** ToLower converts all strings in an array to lowercase.
** This function is useful for case-insensitive comparisons or standardization of string arrays.
**
** @param arr The array of strings to convert to lowercase
** @return []string A new array with all strings converted to lowercase
**************************************************************************************************/
func ToLower(arr []string) []string {
	for i, v := range arr {
		arr[i] = strings.ToLower(v)
	}
	return arr
}

/**************************************************************************************************
** AddressToString converts an array of Ethereum addresses to their string representations.
** This function is useful for serialization, logging, or display purposes.
**
** @param arr The array of Ethereum addresses to convert
** @return []string An array of address strings in hexadecimal format
**************************************************************************************************/
func AddressToString(arr []common.Address) []string {
	arrStr := make([]string, len(arr))
	for i, v := range arr {
		arrStr[i] = v.Hex()
	}
	return arrStr
}

/**************************************************************************************************
** DecodeString extracts a string from an array of interface{} values.
** This function is used to safely access string values from dynamic data structures,
** typically returned from contract calls or external APIs.
**
** @param something An array of interface{} values
** @return string The first element as a string, or empty string if the array is empty
**************************************************************************************************/
func DecodeString(something []interface{}) string {
	if len(something) == 0 {
		return ""
	}
	return something[0].(string)
}

/**************************************************************************************************
** DecodeUint64 extracts a uint64 from an array of interface{} values.
** This function is used to safely access numeric values from dynamic data structures.
**
** @param something An array of interface{} values
** @return uint64 The first element as a uint64, or 0 if the array is empty
**************************************************************************************************/
func DecodeUint64(something []interface{}) uint64 {
	if len(something) == 0 {
		return 0
	}
	return uint64(something[0].(uint8))
}

/**************************************************************************************************
** DecodeUint16s extracts an array of uint16 values from an array of interface{} values.
** This function handles conversion of dynamic values to a strongly typed array.
**
** @param something An array of interface{} values
** @return []uint16 An array of uint16 values, or empty array if the input is empty
**************************************************************************************************/
func DecodeUint16s(something []interface{}) []uint16 {
	if len(something) == 0 {
		return []uint16{}
	}
	uints := []uint16{}
	for _, elem := range something {
		uints = append(uints, elem.(uint16))
	}
	return uints
}

/**************************************************************************************************
** DecodeBigInt extracts a big integer from an array of interface{} values.
** This function is commonly used when working with token amounts or other large numbers
** from contract calls.
**
** @param something An array of interface{} values
** @return *bigNumber.Int The first element as a big integer, or 0 if the array is empty
**************************************************************************************************/
func DecodeBigInt(something []interface{}) *bigNumber.Int {
	if len(something) == 0 {
		return bigNumber.NewInt(0)
	}
	return bigNumber.SetInt(something[0].(*big.Int))
}

/**************************************************************************************************
** DecodeBigInts extracts an array of big integers from an array of interface{} values.
** This function is useful for handling multiple token amounts or other numeric values
** from contract calls.
**
** @param something An array of interface{} values
** @return []*bigNumber.Int An array of big integers, or empty array if the input is empty
**************************************************************************************************/
func DecodeBigInts(something []interface{}) []*bigNumber.Int {
	if len(something) == 0 {
		return []*bigNumber.Int{}
	}
	bigNumbers := []*bigNumber.Int{}
	somethingArr := something[0].([]*big.Int)
	for _, v := range somethingArr {
		bigNumbers = append(bigNumbers, bigNumber.SetInt(v))
	}
	return bigNumbers
}

/**************************************************************************************************
** DecodeBool extracts a boolean value from an array of interface{} values.
** This function is used to safely access boolean flags from dynamic data structures.
**
** @param something An array of interface{} values
** @return bool The first element as a boolean, or false if the array is empty
**************************************************************************************************/
func DecodeBool(something []interface{}) bool {
	if len(something) == 0 {
		return false
	}
	return something[0].(bool)
}

/**************************************************************************************************
** DecodeAddress extracts an Ethereum address from an array of interface{} values.
** This function is commonly used when working with contract addresses or user addresses
** from external data sources.
**
** @param something An array of interface{} values
** @return common.Address The first element as an Ethereum address, or zero address if empty
**************************************************************************************************/
func DecodeAddress(something []interface{}) common.Address {
	if len(something) == 0 {
		return common.Address{}
	}
	return something[0].(common.Address)
}

/**************************************************************************************************
** DecodeAddresses extracts an array of Ethereum addresses from an array of interface{} values.
** This function is useful for handling lists of contracts, tokens, or user addresses.
**
** @param something An array of interface{} values
** @return []common.Address An array of Ethereum addresses, or empty array if the input is empty
**************************************************************************************************/
func DecodeAddresses(something []interface{}) []common.Address {
	if len(something) == 0 {
		return []common.Address{}
	}
	return something[0].([]common.Address)
}

/**************************************************************************************************
** ToRawAmount converts a normalized amount to a raw amount with the specified number of decimals.
** This function is used to convert human-readable token amounts to the raw values used in
** blockchain transactions and smart contracts.
**
** @param amount The normalized amount to convert
** @param decimals The number of decimal places for the token
** @return *bigNumber.Int The raw amount with the appropriate number of decimal places
**************************************************************************************************/
func ToRawAmount(amount *bigNumber.Int, decimals uint64) *bigNumber.Int {
	return bigNumber.NewInt(0).Mul(
		amount,
		bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil),
	)
}

/**************************************************************************************************
** ToNormalizedAmount converts a raw token amount to a normalized floating-point value.
** For example, 1234564784154784147 wei would be converted to 1.234564784154784147 ETH.
**
** This function is essential for displaying token amounts in a human-readable format.
**
** @param amount The raw token amount to normalize
** @param decimals The number of decimal places for the token
** @return *bigNumber.Float The normalized amount as a precise floating-point number
**************************************************************************************************/
func ToNormalizedAmount(amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	return bigNumber.NewFloat(0).Quo(
		bigNumber.NewFloat(0).SetInt(amount),
		bigNumber.NewFloat(0).SetInt(
			bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil),
		),
	)
}

/**************************************************************************************************
** ToNormalizedFloat converts a raw token amount to a standard float64 value.
** This is similar to ToNormalizedAmount but returns a Go float64 instead of a custom big float.
**
** @param amount The raw token amount to normalize
** @param decimals The number of decimal places for the token
** @return float64 The normalized amount as a standard float64 value
**************************************************************************************************/
func ToNormalizedFloat(amount *bigNumber.Int, decimals uint64) float64 {
	floatValue, _ := bigNumber.NewFloat(0).Quo(
		bigNumber.NewFloat(0).SetInt(amount),
		bigNumber.NewFloat(0).SetInt(
			bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil),
		),
	).Float64()
	return floatValue
}

/**************************************************************************************************
** ToNormalizedValue calculates the normalized value of a token amount considering its price.
** This combines the normalization of the amount with multiplication by a price to get
** a value in terms of another currency (typically USD).
**
** @param amount The raw token amount
** @param price The token price in a reference currency (with 6 decimals)
** @param decimals The number of decimal places for the token
** @return *bigNumber.Float The normalized value (amount * price) as a precise floating-point
**************************************************************************************************/
func ToNormalizedValue(amount *bigNumber.Int, price *bigNumber.Int, decimals uint64) *bigNumber.Float {
	normalizedAmount := ToNormalizedAmount(amount, decimals)
	normalizedPrice := ToNormalizedAmount(price, 6)
	normalizedValue := bigNumber.NewFloat(0).Mul(normalizedAmount, normalizedPrice)
	return normalizedValue
}

/**************************************************************************************************
** SafeString provides a safe way to handle potentially empty or invalid strings.
** This function returns the input string if it has content, otherwise it returns a default value.
** It also handles specific cases like quoted empty strings ("").
**
** @param s The string to check
** @param defaultValue The fallback value to return if the string is empty or invalid
** @return string The original string if valid, or the default value otherwise
**************************************************************************************************/
func SafeString(s string, defaultValue string) string {
	if s == "" || s == "\"\"" {
		return defaultValue
	}
	return s
}

/**************************************************************************************************
** SafeStringToUint64 converts a string to a uint64 with safety checks and a default fallback.
** This function handles empty strings, quoted empty strings, and conversion errors by
** returning a specified default value when the conversion fails.
**
** @param n The string to convert to uint64
** @param defaultValue The fallback value to return if conversion fails
** @return uint64 The converted uint64 value, or the default value if conversion fails
**************************************************************************************************/
func SafeStringToUint64(n string, defaultValue uint64) uint64 {
	if n == "" || n == "\"\"" {
		return defaultValue
	}
	result, err := strconv.ParseUint(n, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}
