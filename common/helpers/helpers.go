package helpers

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
)

// ContainsAddress returns true if address exists in addresses
func ContainsAddress[T comparable](addresses []T, address T) bool {
	for _, _address := range addresses {
		if _address == address {
			return true
		}
	}
	return false
}

// UniqueArrayAddress is used to find and remove duplicate from an array of
// ethereum addresses. No such utility function exists with go, so we have to
// implement it ourselves.
func UniqueArrayAddress(arr []common.Address) []common.Address {
	occurred := map[common.Address]bool{}
	result := []common.Address{}
	for i := range arr {
		if !occurred[arr[i]] {
			occurred[arr[i]] = true
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
				logs.Error(err)
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
func FormatAmount(balanceToken string, decimals int) (float64, *bigNumber.Float) {
	fTotalAssets := bigNumber.NewFloat().SetString(balanceToken)
	fDecimals := bigNumber.NewFloat(math.Pow10(decimals))
	humanizedBalance := bigNumber.NewFloat().Quo(fTotalAssets, fDecimals)
	fhumanizedBalance, _ := humanizedBalance.Float64()
	return fhumanizedBalance, humanizedBalance
}

// ContainsUint64 returns true if value exists in arr
func ContainsUint64(arr []uint64, value uint64) bool {
	for _, v := range arr {
		if v == value {
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
	if !ContainsUint64(env.SUPPORTED_CHAIN_IDS, chainID) {
		return 0, false
	}
	return chainID, true
}

func AssertAddress(addressStr string, chainID uint64) (common.Address, bool) {
	if !common.IsHexAddress(addressStr) {
		return common.Address{}, false
	}
	address := common.HexToAddress(addressStr)
	if chainID > 0 && ContainsAddress(env.BLACKLISTED_VAULTS[chainID], address) {
		return common.Address{}, false
	}
	return address, true
}

func AddressIsValid(address common.Address, chainID uint64) bool {
	if (address == common.Address{} || address == common.HexToAddress("0x0000000000000000000000000000000000000000")) {
		return false
	}
	if chainID > 0 && ContainsAddress(env.BLACKLISTED_VAULTS[chainID], address) {
		return false
	}
	return true
}

func StringToBool(str string) bool {
	return str == "true"
}
