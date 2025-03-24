package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TStrategyAdded struct {
	Address      common.Address `json:"address"`
	VaultAddress common.Address `json:"vaultAddress"`
	VaultVersion string         `json:"vaultVersion"`
	ChainID      uint64         `json:"chainID"`
	BlockNumber  uint64         `json:"blockNumber"`
}

type TStrategyMigrated struct {
	VaultAddress       common.Address `json:"vaultAddress"`
	OldStrategyAddress common.Address `json:"oldStrategyAddress"`
	NewStrategyAddress common.Address `json:"newStrategyAddress"`
	ChainID            uint64         `json:"chainID"`
	BlockNumber        uint64         `json:"blockNumber"`
}

type TStrategyAPRType string

const (
	APRTypeCurrent TStrategyAPRType = "current"
	APRTypeForward TStrategyAPRType = "forward"
)

/**************************************************************************************************
** TStrategyStatus represents the operational status of a strategy.
**
** This type is used to categorize strategies based on their current state in the system:
** - Active: The strategy is currently active and has funds allocated to it
** - NotActive: The strategy is either retired or explicitly marked as inactive
** - Unallocated: The strategy is supposed to be active but currently has no funds (LastTotalDebt = 0)
**************************************************************************************************/
type TStrategyStatus string

const (
	StrategyStatusActive      TStrategyStatus = "active"
	StrategyStatusNotActive   TStrategyStatus = "not_active"
	StrategyStatusUnallocated TStrategyStatus = "unallocated"
)

type TStrategy struct {
	Address            common.Address   `json:"address"`      // The address of the strategy
	VaultAddress       common.Address   `json:"vaultAddress"` // The address of the vault
	Name               string           `json:"name"`
	VaultVersion       string           `json:"vaultVersion"` // The version of the vault
	DisplayName        string           `json:"displayName"`  // The name of the strategy
	Description        string           `json:"description"`  // The description of the strategy
	Activation         uint64           `json:"activation"`
	ChainID            uint64           `json:"chainID"`
	DoHealthCheck      bool             `json:"doHealthCheck"`
	IsActive           bool             `json:"isActive"`
	IsInQueue          bool             `json:"isInQueue"`
	IsRetired          bool             `json:"isRetired"`               // If false, will bypass the `IsActive` variable
	ShouldRefresh      bool             `json:"shouldRefresh,omitempty"` // Will be refreshed no matter what
	Status             TStrategyStatus  `json:"status"`                  // Categorized status: active, not_active, or unallocated
	TimeActivated      *bigNumber.Int   `json:"-"`                       // When the strategy was activated. Only used internaly to compute the longevityImpact.
	KeepCRV            *bigNumber.Int   `json:"keepCRV"`
	KeepCRVPercent     *bigNumber.Int   `json:"keepCRVPercent"`
	KeepCVX            *bigNumber.Int   `json:"keepCVX"`
	LastTotalDebt      *bigNumber.Int   `json:"lastTotalDebt"`           // Used to filter strategies and by the FE
	LastTotalLoss      *bigNumber.Int   `json:"lastTotalLoss"`           // Used by the FE
	LastTotalGain      *bigNumber.Int   `json:"lastTotalGain"`           // Used by the FE
	LastPerformanceFee *bigNumber.Int   `json:"lastPerformanceFee"`      // Used for APR calculation and by the FE
	LastReport         *bigNumber.Int   `json:"lastReport"`              // Used by the FE
	LastDebtRatio      *bigNumber.Int   `json:"lastDebtRatio,omitempty"` // Only > 0.2.2 | Used by the APY process
	NetAPR             *bigNumber.Float `json:"netAPR"`                  // The net APR of the strategy
	APRType            TStrategyAPRType `json:"aprType"`                 // The type of APR of the strategy
	Protocols          []string         `json:"protocols"`               // The protocols used by the strategy
}
