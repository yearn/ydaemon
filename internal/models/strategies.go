package models

import (
	"encoding/json"
	"strconv"

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
	NetAPR             float64          `json:"netAPR"`                  // The net APR of the strategy
	APRType            TStrategyAPRType `json:"aprType"`                 // The type of APR of the strategy
	Protocols          []string         `json:"protocols"`               // The protocols used by the strategy
}

/**************************************************************************************************
** TStrategyReportDB represents a strategy report from the database
**
** This type is used to store the data retrieved from the database for strategy reports
**************************************************************************************************/
type TStrategyReportDB struct {
	ChainID            uint64          `json:"chainId"`
	Address            string          `json:"address"`
	EventName          string          `json:"eventName"`
	Profit             string          `json:"profit"`
	Loss               string          `json:"loss"`
	DebtPayment        string          `json:"debtPayment"`
	DebtOutstanding    string          `json:"debtOutstanding"`
	ProtocolFees       string          `json:"protocolFees"`
	PerformanceFees    string          `json:"performanceFees"`
	APR                json.RawMessage `json:"apr"`
	ProfitUSD          string          `json:"profitUsd"`
	LossUSD            string          `json:"lossUsd"`
	DebtPaymentUSD     string          `json:"debtPaymentUsd"`
	DebtOutstandingUSD string          `json:"debtOutstandingUsd"`
	ProtocolFeesUSD    string          `json:"protocolFeesUsd"`
	PerformanceFeesUSD string          `json:"performanceFeesUsd"`
	PriceUSD           string          `json:"priceUsd"`
	PriceSource        string          `json:"priceSource"`
	BlockNumber        uint64          `json:"blockNumber"`
	BlockTime          uint64          `json:"blockTime"`
	LogIndex           uint64          `json:"logIndex"`
	TransactionHash    string          `json:"transactionHash"`
}

// TStrategyCmsMetadataSchema represents the strategy metadata structure from ycms
type TStrategyCmsMetadataSchema struct {
	ChainID     uint64         `json:"chainId"`
	Address     common.Address `json:"address"`
	IsRetired   bool           `json:"isRetired"`
	DisplayName *string        `json:"displayName,omitempty"`
	Description *string        `json:"description,omitempty"`
	Protocols   []string       `json:"protocols"`
}




type KongLastReportDetail struct {
	ChainID         int     `json:"chainId"`
	Address         string  `json:"address"`
	BlockNumber     string  `json:"blockNumber"`
	BlockTime       string  `json:"blockTime"`
	TransactionHash string  `json:"transactionHash"`
	Profit          *string `json:"profit"`
	ProfitUsd       *float64 `json:"profitUsd"`
	Loss            *string `json:"loss"`
	LossUsd         *float64 `json:"lossUsd"`
	APR             *struct {
		Gross *float64 `json:"gross"`
		Net   *float64 `json:"net"`
	} `json:"apr"`
}

type KongLenderStatus struct {
	Name    string  `json:"name"`
	Assets  *string `json:"assets"`
	Rate    *string `json:"rate"`
	Address string  `json:"address"`
}


type KongClaim struct {
	ChainID    int      `json:"chainId"`
	Address    string   `json:"address"`
	Name       string   `json:"name"`
	Symbol     string   `json:"symbol"`
	Decimals   int      `json:"decimals"`
	Balance    *string  `json:"balance"`
	BalanceUsd *float64 `json:"balanceUsd"`
}


type KongRisk struct {
	Label               string   `json:"label"`
	AuditScore          *float64 `json:"auditScore"`
	CodeReviewScore     *float64 `json:"codeReviewScore"`
	ComplexityScore     *float64 `json:"complexityScore"`
	ProtocolSafetyScore *float64 `json:"protocolSafetyScore"`
	TeamKnowledgeScore  *float64 `json:"teamKnowledgeScore"`
	TestingScore        *float64 `json:"testingScore"`
}

type KongMeta struct {
	DisplayName *string  `json:"displayName"`
	Description *string  `json:"description"`
	Protocols   []string `json:"protocols"`
}


type KongStrategy struct {
	ChainID               int                   `json:"chainId"`
	Address               string                `json:"address"`
	APIVersion            *string               `json:"apiVersion"`
	BalanceOfWant         *string               `json:"balanceOfWant"`
	BaseFeeOracle         *string               `json:"baseFeeOracle"`
	CreditThreshold       *string               `json:"creditThreshold"`
	Crv                   *string               `json:"crv"`
	CurveVoter            *string               `json:"curveVoter"`
	DelegatedAssets       *string               `json:"delegatedAssets"`
	DoHealthCheck         *bool                 `json:"doHealthCheck"`
	EmergencyExit         *bool                 `json:"emergencyExit"`
	Erc4626               *bool                 `json:"erc4626"`
	EstimatedTotalAssets  *string               `json:"estimatedTotalAssets"`
	ForceHarvestTriggerOnce *bool               `json:"forceHarvestTriggerOnce"`
	Gauge                 *string               `json:"gauge"`
	HealthCheck           *string               `json:"healthCheck"`
	InceptTime            *string               `json:"inceptTime"`
	InceptBlock           *string               `json:"inceptBlock"`
	IsActive              *bool                 `json:"isActive"`
	IsBaseFeeAcceptable   *bool                 `json:"isBaseFeeAcceptable"`
	IsOriginal            *bool                 `json:"isOriginal"`
	Keeper                *string               `json:"keeper"`
	LocalKeepCRV          *string               `json:"localKeepCRV"`
	MaxReportDelay        *string               `json:"maxReportDelay"`
	MetadataURI           *string               `json:"metadataURI"`
	MinReportDelay        *string               `json:"minReportDelay"`
	Name                  *string               `json:"name"`
	Proxy                 *string               `json:"proxy"`
	Rewards               *string               `json:"rewards"`
	StakedBalance         *string               `json:"stakedBalance"`
	Strategist            *string               `json:"strategist"`
	TradeFactory          *string               `json:"tradeFactory"`
	Vault                 string                `json:"vault"`
	Want                  *string               `json:"want"`
	DOMAIN_SEPARATOR      *string               `json:"DOMAIN_SEPARATOR"`
	FACTORY               *string               `json:"FACTORY"`
	MAX_FEE               *int                  `json:"MAX_FEE"`
	MIN_FEE               *int                  `json:"MIN_FEE"`
	Decimals              *int                  `json:"decimals"`
	FullProfitUnlockDate  *string               `json:"fullProfitUnlockDate"`
	IsShutdown            *bool                 `json:"isShutdown"`
	LastReport            *string               `json:"lastReport"`
	LastReportDetail      *KongLastReportDetail `json:"lastReportDetail"`
	Management            *string               `json:"management"`
	PendingManagement     *string               `json:"pendingManagement"`
	PerformanceFee        *int              `json:"performanceFee"`
	PerformanceFeeRecipient *string             `json:"performanceFeeRecipient"`
	PricePerShare         *string               `json:"pricePerShare"`
	ProfitMaxUnlockTime   *string               `json:"profitMaxUnlockTime"`
	ProfitUnlockingRate   *string               `json:"profitUnlockingRate"`
	Symbol                *string               `json:"symbol"`
	TotalAssets           *string               `json:"totalAssets"`
	TotalDebt             *string               `json:"totalDebt"`
	TotalIdle             *string               `json:"totalIdle"`
	TotalSupply           *string               `json:"totalSupply"`
	TotalDebtUsd          *float64              `json:"totalDebtUsd"`
	LenderStatuses        []KongLenderStatus    `json:"lenderStatuses"`
	Claims                []KongClaim           `json:"claims"`
	Risk                  *KongRisk             `json:"risk"`
	Meta                  *KongMeta             `json:"meta"`
	V3                    bool                  `json:"v3"`
	Yearn                 bool                  `json:"yearn"`
}


func (s KongStrategy) GetAddress() common.Address {
	return common.HexToAddress(s.Address)
}

func (s KongStrategy) GetVaultAddress() common.Address {
	return common.HexToAddress(s.Vault)
}

func (s KongStrategy) GetAPIVersion() string {
	if s.APIVersion == nil {
		return ""
	}
	return *s.APIVersion
}

func (s KongStrategy) GetDisplayName() string {
	if s.Meta == nil || s.Meta.DisplayName == nil {
		return s.GetName()
	}
	return *s.Meta.DisplayName
}

func (s KongStrategy) GetName() string {
	if s.Name == nil {
		return ""
	}
	return *s.Name
}

func (s KongStrategy) GetIsActive() bool {
	if s.IsActive == nil {
		return false
	}
	return *s.IsActive
}

func (s KongStrategy) GetDoHealthCheck() bool {
	if s.DoHealthCheck == nil {
		return false
	}
	return *s.DoHealthCheck
}

func (s KongStrategy) GetInceptBlock() uint64 {
	if s.InceptBlock == nil {
		return 0
	}
	blockNum, err := strconv.ParseUint(*s.InceptBlock, 10, 64)
	if err != nil {
		return 0
	}
	return blockNum
}

func (s KongStrategy) ToTStrategy() TStrategy {
	strategy := TStrategy{
		Address:       s.GetAddress(),
		VaultAddress:  s.GetVaultAddress(),
		Name:          s.GetDisplayName(),
		DisplayName:   s.GetDisplayName(),
		VaultVersion:  s.GetAPIVersion(),
		ChainID:       uint64(s.ChainID),
		DoHealthCheck: s.GetDoHealthCheck(),
		IsActive:      s.GetIsActive(),
		Activation:    uint64(s.GetInceptBlock()),
	}

	// DisplayName and Description from Meta
	if s.Meta != nil {
		if s.Meta.DisplayName != nil {
			strategy.DisplayName = *s.Meta.DisplayName
		}
		if s.Meta.Description != nil {
			strategy.Description = *s.Meta.Description
		}
		if s.Meta.Protocols != nil {
			strategy.Protocols = s.Meta.Protocols
		}
	}

	// LastReport
	if s.LastReport != nil {
		strategy.LastReport = bigNumber.NewInt().SetString(*s.LastReport)
	}

	// LastTotalDebt from TotalDebt
	if s.TotalDebt != nil {
		strategy.LastTotalDebt = bigNumber.NewInt().SetString(*s.TotalDebt)
	}

	// LastReportDetail data
	if s.LastReportDetail != nil {
		if s.LastReportDetail.Profit != nil {
			strategy.LastTotalGain = bigNumber.NewInt().SetString(*s.LastReportDetail.Profit)
		}
		if s.LastReportDetail.Loss != nil {
			strategy.LastTotalLoss = bigNumber.NewInt().SetString(*s.LastReportDetail.Loss)
		}
		if s.LastReportDetail.APR != nil && s.LastReportDetail.APR.Net != nil {
			strategy.NetAPR = *s.LastReportDetail.APR.Net
		}
	}

	// PerformanceFee
	if s.PerformanceFee != nil {
		strategy.LastPerformanceFee = bigNumber.NewInt(int64(*s.PerformanceFee))
	}

	// LocalKeepCRV
	if s.LocalKeepCRV != nil {
		strategy.KeepCRV = bigNumber.NewInt().SetString(*s.LocalKeepCRV)
	}

	// IsRetired based on IsShutdown
	if s.IsShutdown != nil {
		strategy.IsRetired = *s.IsShutdown
	}

	return strategy
}