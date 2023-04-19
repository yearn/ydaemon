package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

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
	VaultVersion      string
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

type TStrategyInitialization struct {
	TxHash      common.Hash
	BlockNumber uint64
	TxIndex     uint
	LogIndex    uint
}

type TStrategy struct {
	Address                 common.Address          `json:"address"`
	VaultAddress            common.Address          `json:"vaultAddress"`
	KeeperAddress           common.Address          `json:"keeperAddress"`
	StrategistAddress       common.Address          `json:"strategistAddress"`
	RewardsAddress          common.Address          `json:"rewardsAddress"`
	HealthCheckAddress      common.Address          `json:"healthCheckAddress"`
	VaultVersion            string                  `json:"vaultVersion"`
	Name                    string                  `json:"name"`
	DisplayName             string                  `json:"displayName"`
	GroupName               string                  `json:"groupName"`
	Description             string                  `json:"description"`
	APIVersion              string                  `json:"apiVersion"`
	Protocols               []string                `json:"protocols"`
	CreditAvailable         *bigNumber.Int          `json:"creditAvailable"`
	DebtOutstanding         *bigNumber.Int          `json:"debtOutstanding"`
	ExpectedReturn          *bigNumber.Int          `json:"expectedReturn"`
	PerformanceFee          *bigNumber.Int          `json:"performanceFee"`
	Activation              *bigNumber.Int          `json:"activation"`
	DebtRatio               *bigNumber.Int          `json:"debtRatio,omitempty"`         // Only > 0.2.2
	DebtLimit               *bigNumber.Int          `json:"debtLimit,omitempty"`         // Only = 0.2.2
	RateLimit               *bigNumber.Int          `json:"rateLimit,omitempty"`         // Only < 0.3.2
	MinDebtPerHarvest       *bigNumber.Int          `json:"minDebtPerHarvest,omitempty"` // Only >= 0.3.2
	MaxDebtPerHarvest       *bigNumber.Int          `json:"maxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastReport              *bigNumber.Int          `json:"lastReport"`
	TotalDebt               *bigNumber.Int          `json:"totalDebt"`
	TotalGain               *bigNumber.Int          `json:"totalGain"`
	TotalLoss               *bigNumber.Int          `json:"totalLoss"`
	EstimatedTotalAssets    *bigNumber.Int          `json:"estimatedTotalAssets"`
	KeepCRV                 *bigNumber.Int          `json:"keepCRV"`
	DelegatedAssets         *bigNumber.Int          `json:"delegatedAssets"`
	WithdrawalQueuePosition *bigNumber.Int          `json:"withdrawalQueuePosition"`
	ChainID                 uint64                  `json:"chainID"`
	DoHealthCheck           bool                    `json:"doHealthCheck"`
	EmergencyExit           bool                    `json:"emergencyExit"`
	IsActive                bool                    `json:"isActive"`
	IsInQueue               bool                    `json:"isInQueue"`
	Initialization          TStrategyInitialization `json:"-"`
}

// TStrategyFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/strategies/all
type TStrategyFromMeta struct {
	Address          common.Address   `json:"address"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	RelatedAddresses []common.Address `json:"addresses"`
	Protocols        []string         `json:"protocols"`
	ChainID          uint64           `json:"chainID"`
	Localization     *TLocalization   `json:"localization,omitempty"`
}
