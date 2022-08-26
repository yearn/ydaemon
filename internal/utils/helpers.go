package utils

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/majorfi/ydaemon/internal/logs"
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
