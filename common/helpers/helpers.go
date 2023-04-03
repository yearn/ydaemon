package helpers

import (
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/traces"
)

// Intersects returns true if both arrays have at least one element in common
func Intersects(arr1 []string, arr2 []string) bool {
	for _, v := range arr1 {
		if ContainsSubString(arr2, v) {
			return true
		}
	}
	return false
}

// UniqueArrayAddress is used to find and remove duplicate from an array of
// ethereum addresses. No such utility function exists with go, so we have to
// implement it ourselves.
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

// RemoveFromArray is used to find and remove an element from an array
func RemoveFromArray[T comparable](arr []T, element T) []T {
	result := []T{}
	for i := range arr {
		if arr[i] != element {
			result = append(result, arr[i])
		}
	}
	return result
}

// ReadAllFilesInDir is used to grab the content of all the files in a specific directory,
// matching the suffix condition
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
			content, err := ioutil.ReadFile(directory + outputFileName)
			if err != nil {
				traces.
					Capture(`error`, `impossible to read files in `+directory).
					SetEntity(`files`).
					SetExtra(`error`, err.Error()).
					SetTag(`directory`, directory).
					SetTag(`file`, outputFileName).
					Send()
				continue
			}
			filenames = append(filenames, outputFileName)
			dest = append(dest, content)
		}
	}
	return dest, filenames, nil
}

// FormatUint64 is used to convert a string to a uint64, and returns the second value
// if the first one if invalid
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

// FormatAmount is used to convert a uint256 string to a float64, based on the provided decimal
func FormatAmount(amount string, decimals int) (float64, *bigNumber.Float) {
	fTotalAssets := bigNumber.NewFloat().SetString(amount)
	fDecimals := bigNumber.NewFloat(math.Pow10(decimals))
	humanizedBalance := bigNumber.NewFloat().Quo(fTotalAssets, fDecimals)
	fhumanizedBalance, _ := humanizedBalance.Float64()
	return fhumanizedBalance, humanizedBalance
}

// Contains returns true if value exists in arr
func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsSubString returns true if value exists in arr
// arr are assumed to contain the substrings of value
func ContainsSubString(arr []string, value string) bool {
	for _, v := range arr {
		if strings.Contains(value, v) {
			return true
		}
	}
	return false
}

func AssertChainID(chainIDStr string) (uint64, bool) {
	chainID, err := strconv.ParseUint(chainIDStr, 10, 64)
	if err != nil {
		return 0, false
	}
	if chainID == 1337 {
		return 1, true
	}
	if !Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
		return 0, false
	}
	return chainID, true
}

func AssertAddress(addressStr string, chainID uint64) (common.Address, bool) {
	if !common.IsHexAddress(addressStr) {
		return common.Address{}, false
	}
	address := common.HexToAddress(addressStr)
	if chainID > 0 && Contains(env.BLACKLISTED_VAULTS[chainID], address) {
		return common.Address{}, false
	}
	return address, true
}

func AddressIsValid(address common.Address, chainID uint64) bool {
	if (address == common.Address{} || addresses.Equals(address, "0x0000000000000000000000000000000000000000")) {
		return false
	}
	if chainID > 0 && Contains(env.BLACKLISTED_VAULTS[chainID], address) {
		return false
	}
	return true
}

func StringToBool(str string) bool {
	return str == "true"
}

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

func ToLower(arr []string) []string {
	for i, v := range arr {
		arr[i] = strings.ToLower(v)
	}
	return arr
}

func AddressToString(arr []common.Address) []string {
	arrStr := make([]string, len(arr))
	for i, v := range arr {
		arrStr[i] = v.String()
	}
	return arrStr
}

func DecodeString(something []interface{}) string {
	if len(something) == 0 {
		return ""
	}
	return something[0].(string)
}
func DecodeUint64(something []interface{}) uint64 {
	if len(something) == 0 {
		return 0
	}
	return uint64(something[0].(uint8))
}
func DecodeBigInt(something []interface{}) *bigNumber.Int {
	if len(something) == 0 {
		return bigNumber.NewInt(0)
	}
	return bigNumber.SetInt(something[0].(*big.Int))
}
func DecodeBool(something []interface{}) bool {
	if len(something) == 0 {
		return false
	}
	return something[0].(bool)
}
func DecodeAddress(something []interface{}) common.Address {
	if len(something) == 0 {
		return common.Address{}
	}
	return something[0].(common.Address)
}

func ToNormalizedAmount(amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	return bigNumber.NewFloat(0).Quo(
		bigNumber.NewFloat(0).SetInt(amount),
		bigNumber.NewFloat(0).SetInt(
			bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil),
		),
	)
}

func ToNormalizedValue(amount *bigNumber.Int, price *bigNumber.Int, decimals uint64) *bigNumber.Float {
	normalizedAmount := ToNormalizedAmount(amount, decimals)
	normalizedPrice := ToNormalizedAmount(price, 6)
	normalizedValue := bigNumber.NewFloat(0).Mul(normalizedAmount, normalizedPrice)
	return normalizedValue
}
