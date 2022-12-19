package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

type TVaultsFromRegistry struct {
	RegistryAddress common.Address
	VaultsAddress   common.Address
	TokenAddress    common.Address
	BlockHash       common.Hash
	APIVersion      string
	Type            string
	BlockNumber     uint64
	Activation      uint64
	ManagementFee   uint64
	TxIndex         uint
	LogIndex        uint
}
