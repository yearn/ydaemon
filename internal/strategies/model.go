package strategies

import (
	"strconv"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

type TStrategyAdded struct {
	VaultAddress      ethcommon.Address
	StrategyAddress   ethcommon.Address
	TxHash            ethcommon.Hash
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
	VaultAddress       ethcommon.Address
	OldStrategyAddress ethcommon.Address
	NewStrategyAddress ethcommon.Address
	TxHash             ethcommon.Hash
	BlockNumber        uint64
	TxIndex            uint
	LogIndex           uint
}

type TStrategyInitialization struct {
	TxHash      ethcommon.Hash
	BlockNumber uint64
	TxIndex     uint
	LogIndex    uint
}

type TStrategy struct {
	Address                 ethcommon.Address       `json:"address"`
	VaultAddress            ethcommon.Address       `json:"vaultAddress"`
	KeeperAddress           ethcommon.Address       `json:"keeperAddress"`
	StrategistAddress       ethcommon.Address       `json:"strategistAddress"`
	RewardsAddress          ethcommon.Address       `json:"rewardsAddress"`
	HealthCheckAddress      ethcommon.Address       `json:"healthCheckAddress"`
	VaultVersion            string                  `json:"vaultVersion"`
	Name                    string                  `json:"name"`
	DisplayName             string                  `json:"displayName"`
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

func (t *TStrategy) BuildRiskScore() *TStrategyFromRisk {
	strategyAddress := common.FromAddress(t.Address)
	vaultAddress := common.FromAddress(t.VaultAddress)
	strategyGroup := getStrategyGroup(t.ChainID, t)
	strategyRiskScore := getDefaultRiskGroup()
	strategyRiskScore.Address = strategyAddress
	strategyRiskScore.ChainID = t.ChainID
	if strategyGroup == nil {
		traces.
			Capture(`warn`, `impossible to find strategyGroup for group `+t.Name).
			SetEntity(`strategy`).
			SetTag(`chainID`, strconv.FormatUint(t.ChainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(t.ChainID)).
			SetTag(`strategyAddress`, t.Address.Hex()).
			SetTag(`strategyName`, t.Name).
			Send()
		return &strategyRiskScore
	}

	// Fetch activation and tvl from multicall
	strategyRiskScore.RiskScores.LongevityImpact = getLongevityImpact(t.Activation)

	// Fetch tvl of strategy
	tokenData, ok := tokens.FindUnderlyingForVault(t.ChainID, vaultAddress)
	if !ok {
		traces.
			Capture(`warn`, `impossible to find token for vault `+t.VaultAddress.Hex()).
			SetEntity(`strategy`).
			SetTag(`chainID`, strconv.FormatUint(t.ChainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(t.ChainID)).
			SetTag(`strategyAddress`, strategyAddress.Hex()).
			SetTag(`strategyName`, t.Name).
			Send()
		return &strategyRiskScore
	}

	_tokenPrice, ok := prices.FindPrice(t.ChainID, common.FromAddress(tokenData.Address))
	if !ok {
		traces.
			Capture(`warn`, `impossible to find tokenPrice for token `+tokenData.Address.Hex()).
			SetEntity(`strategy`).
			SetTag(`chainID`, strconv.FormatUint(t.ChainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(t.ChainID)).
			SetTag(`strategyAddress`, strategyAddress.Hex()).
			SetTag(`strategyName`, t.Name).
			Send()
	}
	_, price := helpers.FormatAmount(_tokenPrice.String(), 6)
	_, amount := helpers.FormatAmount(t.EstimatedTotalAssets.String(), int(tokenData.Decimals))
	tvl := bigNumber.NewFloat(0).Mul(amount, price)

	// Updating the default schema
	strategyRiskScore.RiskGroup = strategyGroup.Label
	strategyRiskScore.RiskScores.AuditScore = strategyGroup.AuditScore
	strategyRiskScore.RiskScores.CodeReviewScore = strategyGroup.CodeReviewScore
	strategyRiskScore.RiskScores.ComplexityScore = strategyGroup.ComplexityScore
	strategyRiskScore.RiskScores.ProtocolSafetyScore = strategyGroup.ProtocolSafetyScore
	strategyRiskScore.RiskScores.TeamKnowledgeScore = strategyGroup.TeamKnowledgeScore
	strategyRiskScore.RiskScores.TestingScore = strategyGroup.TestingScore
	strategyRiskScore.RiskScores.TVLImpact = getTVLImpact(tvl)
	strategyRiskScore.Allocation = strategyGroup.Allocation
	return &strategyRiskScore
}

/**********************************************************************************************
** Set of functions to store and retrieve the strategies from the cache and/or database and
** being able to access them from the rest of the application.
** The _strategyMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _strategyMap = make(map[uint64]map[ethcommon.Address]*TStrategy)
var _strategyWithdrawalQueueMap = make(map[uint64]map[ethcommon.Address][]common.Address)

/**********************************************************************************************
** ListStrategies will, for a given chainID, return the list of all the strategies stored in
** the _strategyMap.
**********************************************************************************************/
func ListStrategies(chainID uint64) []*TStrategy {
	var strategies []*TStrategy
	for _, strategy := range _strategyMap[chainID] {
		strategies = append(strategies, strategy)
	}
	return strategies
}

/**********************************************************************************************
** ListStrategiesForVault will, for a given chainID and a given vault address, return the list
** of all the strategies stored in the _strategyMap.
**********************************************************************************************/
func ListStrategiesForVault(chainID uint64, vaultAddress common.Address) []*TStrategy {
	var strategies []*TStrategy
	for _, strategy := range _strategyMap[chainID] {
		if strategy.VaultAddress.Hex() == vaultAddress.Hex() {
			strategies = append(strategies, strategy)
		}
	}
	return strategies
}

/**********************************************************************************************
** ListStrategiesAddresses will, for a given chainID, return the list of addresses of all the
** strategies stored in _strategyMap.
**********************************************************************************************/
func ListStrategiesAddresses(chainID uint64) []ethcommon.Address {
	var addresses []ethcommon.Address
	for address := range _strategyMap[chainID] {
		addresses = append(addresses, address)
	}
	return addresses
}

/**********************************************************************************************
** FindStrategy will, for a given chainID, try to find the provided strategyAddress stored in
** the _strategyMap. It will return the strategy if found, and a boolean indicating if the
** strategy was found or not.
**********************************************************************************************/
func FindStrategy(chainID uint64, strategyAddress ethcommon.Address) (*TStrategy, bool) {
	strategy, ok := _strategyMap[chainID][strategyAddress]
	if !ok {
		return nil, false
	}
	return strategy, true
}

/**********************************************************************************************
** SlipStrategiesAddedPerVault will transform a list of TStrategyAdded into a map of
** TStrategyAdded per vault address:
** - [vaultAddress] - [strategyAddress] - TStrategyAdded
**********************************************************************************************/
func SlipStrategiesAddedPerVault(strategiesAddedList []TStrategyAdded) map[ethcommon.Address]map[ethcommon.Address]TStrategyAdded {
	strategiesAddedPerVault := make(map[ethcommon.Address]map[ethcommon.Address]TStrategyAdded)
	for _, strategy := range strategiesAddedList {
		if _, ok := strategiesAddedPerVault[strategy.VaultAddress]; !ok {
			strategiesAddedPerVault[strategy.VaultAddress] = make(map[ethcommon.Address]TStrategyAdded)
		}
		strategiesAddedPerVault[strategy.VaultAddress][strategy.StrategyAddress] = strategy
	}
	return strategiesAddedPerVault
}

/**********************************************************************************************
** SetInStrategiesWithdrawalQueue will mirror a vault withdrawal queue in the strategies
** package to avoid import circles.
**********************************************************************************************/
func SetInStrategiesWithdrawalQueue(chainID uint64, vaultAddress ethcommon.Address, queue []common.Address) {
	if _, ok := _strategyWithdrawalQueueMap[chainID]; !ok {
		_strategyWithdrawalQueueMap[chainID] = make(map[ethcommon.Address][]common.Address)
	}
	_strategyWithdrawalQueueMap[chainID][vaultAddress] = queue
}

/**********************************************************************************************
** FindWithdrawalQueueForVault will retrieve the withdrawal queue for a given vault address.
**********************************************************************************************/
func FindWithdrawalQueueForVault(chainID uint64, vaultAddress ethcommon.Address) ([]common.Address, bool) {
	if _, ok := _strategyWithdrawalQueueMap[chainID]; !ok {
		return nil, false
	}
	return _strategyWithdrawalQueueMap[chainID][vaultAddress], true
}
