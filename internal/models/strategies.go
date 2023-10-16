package models

import (
	"time"

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

type TStrategy struct {
	// Immutable elements. They won't change
	Address       common.Address `json:"address"`      // The address of the strategy
	VaultAddress  common.Address `json:"vaultAddress"` // The address of the vault
	Name          string         `json:"name"`
	VaultVersion  string         `json:"vaultVersion"` // The version of the vault
	Activation    uint64         `json:"activation"`
	ChainID       uint64         `json:"chainID"`
	DoHealthCheck bool           `json:"doHealthCheck"`
	IsActive      bool           `json:"isActive"`
	IsInQueue     bool           `json:"isInQueue"`
	TimeActivated *bigNumber.Int `json:"-"` // When the strategy was activated. Only used internaly to compute the longevityImpact.

	// Semi-mutable eelements. They can change but rarely
	KeepCRV        *bigNumber.Int `json:"keepCRV"`
	KeepCRVPercent *bigNumber.Int `json:"keepCRVPercent"`
	KeepCVX        *bigNumber.Int `json:"keepCVX"`
	EmergencyExit  bool           `json:"emergencyExit"`

	// Mutable elements. They will often change
	LastCreditAvailable      *bigNumber.Int `json:"lastCreditAvailable"`
	LastDebtOutstanding      *bigNumber.Int `json:"lastDebtOutstanding"`
	LastExpectedReturn       *bigNumber.Int `json:"lastExpectedReturn"`
	LastReport               *bigNumber.Int `json:"lastReport"`
	LastTotalDebt            *bigNumber.Int `json:"lastTotalDebt"`
	LastTotalGain            *bigNumber.Int `json:"lastTotalGain"`
	LastTotalLoss            *bigNumber.Int `json:"lastTotalLoss"`
	LastEstimatedTotalAssets *bigNumber.Int `json:"lastEstimatedTotalAssets"`
	LastDelegatedAssets      *bigNumber.Int `json:"lastDelegatedAssets"` //Amount of asset delegated
	LastPerformanceFee       *bigNumber.Int `json:"lastPerformanceFee"`
	LastDebtRatio            *bigNumber.Int `json:"lastDebtRatio,omitempty"`         // Only > 0.2.2
	LastDebtLimit            *bigNumber.Int `json:"lastDebtLimit,omitempty"`         // Only = 0.2.2
	LastRateLimit            *bigNumber.Int `json:"lastRateLimit,omitempty"`         // Only < 0.3.2
	LastMinDebtPerHarvest    *bigNumber.Int `json:"lastMinDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastMaxDebtPerHarvest    *bigNumber.Int `json:"lastMaxDebtPerHarvest,omitempty"` // Only >= 0.3.2
	LastUpdate               time.Time      `json:"lastUpdate"`                      // When this struct was last updated

	// Manual elements. They are manually set by the team
	IsRetired   bool     `json:"isRetired"`   // If false, will bypass the `IsActive` variable
	DisplayName string   `json:"displayName"` // The name of the strategy
	Description string   `json:"description"` // The description of the strategy
	Protocols   []string `json:"protocols"`   // The protocols used by the strategy

	// Extra fields, used for control purpose. They are not stored, but can be added to trigger some actions
	ShouldRefresh bool `json:"shouldRefresh,omitempty"` // Will be refreshed no matter what
}
