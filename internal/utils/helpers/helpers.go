package helpers

import (
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

//ContainsAddress returns true if address exists in addresses
func ContainsAddress(addresses []common.Address, address common.Address) bool {
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

// ValueWithFallback is used to return the first value if it is not empty, otherwise
// it returns the second value
func ValueWithFallback(s string, defaultValue string) string {
	if s == "" || s == "\"\"" {
		return defaultValue
	}
	return s
}

// BValueWithFallbackUint64 is used to return the first value as uint64 if it is
// not empty, otherwise it returns the second value
func BValueWithFallbackUint64(s *big.Int, defaultValue uint64) uint64 {
	if s == nil {
		return defaultValue
	}
	return s.Uint64()
}

// BValueWithFallbackString is used to return the first value as string if it is
// not empty, otherwise it returns the second value
func BValueWithFallbackString(s *big.Int, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return s.String()
}

// StrToUint is used to convert a string to a uint64, and returns the second value
// if the first one if invalid
func StrToUint(s string, defaultValue uint64) uint64 {
	if s == "" {
		return defaultValue
	}
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
