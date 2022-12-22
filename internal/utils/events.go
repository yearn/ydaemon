package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

type TVaultType string

const (
	VaultTypeExperimental TVaultType = "Experimental"
	VaultTypeStandard     TVaultType = "Standard"
	VaultTypeAutomated    TVaultType = "Automated"
)

type TVaultsFromRegistry struct {
	RegistryAddress common.Address
	VaultsAddress   common.Address
	TokenAddress    common.Address
	BlockHash       common.Hash
	Type            TVaultType
	APIVersion      string
	BlockNumber     uint64
	Activation      uint64
	ManagementFee   uint64
	TxIndex         uint
	LogIndex        uint
}
