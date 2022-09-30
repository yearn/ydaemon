package common

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

type Address struct{ common.Address }

func FromAddress(a common.Address) Address {
	return Address{a}
}

func HexToAddress(a string) Address {
	return Address{common.HexToAddress(a)}
}

func IsHexAddress(a string) bool {
	return common.IsHexAddress(a)
}

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Hex())
}
