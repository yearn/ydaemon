package utils

import "github.com/ethereum/go-ethereum/common"

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
