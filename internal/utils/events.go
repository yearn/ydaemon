package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
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

type TStrategyAdded struct {
	VaultAddress      common.Address
	StrategyAddress   common.Address
	TxHash            common.Hash
	DebtRatio         *bigNumber.Int // >= 0.3.0
	MaxDebtPerHarvest *bigNumber.Int // >= 0.3.2
	MinDebtPerHarvest *bigNumber.Int // >= 0.3.2
	PerformanceFee    *bigNumber.Int // >= 0.2.2
	DebtLimit         *bigNumber.Int // == 0.2.2
	RateLimit         *bigNumber.Int // == 0.2.2 - 0.3.0
	BlockNumber       uint64
	TxIndex           uint
	LogIndex          uint
}

type TStrategyMigrated struct {
	VaultAddress       common.Address
	OldStrategyAddress common.Address
	NewStrategyAddress common.Address
	TxHash             common.Hash
	BlockNumber        uint64
	TxIndex            uint
	LogIndex           uint
}
