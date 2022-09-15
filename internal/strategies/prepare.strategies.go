package strategies

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
)

func buildDelegated(delegatedBalanceToken string, decimals int, humanizedPrice *big.Float) float64 {
	_, delegatedTVL := helpers.FormatAmount(delegatedBalanceToken, decimals)
	fHumanizedTVLPrice, _ := big.NewFloat(0).Mul(delegatedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

// BuildStrategies prepare the TStrategy struct based of a bunch of data
func BuildStrategies(
	chainID uint64,
	withStrategiesDetails bool,
	withStrategiesRisk bool,
	strategiesCondition string,
	humanizedTokenPrice *big.Float,
	vaultFromGraph models.TVaultFromGraph,
) []TStrategy {
	strategies := []TStrategy{}
	strategiesFromMeta := meta.Store.StrategiesFromMeta[chainID]
	strategiesFromMulticall := Store.StrategyMultiCallData[chainID]
	strategiesFromRisk := Store.StrategiesFromRisk[chainID]

	for _, strategy := range vaultFromGraph.Strategies {
		multicallData, ok := strategiesFromMulticall[common.HexToAddress(strategy.Address)]
		if !ok {
			multicallData = models.TStrategyMultiCallData{}
		}

		currentStrategy := TStrategy{
			Address:     common.HexToAddress(strategy.Address).String(),
			Name:        strategy.Name,
			Description: strategiesFromMeta[common.HexToAddress(strategy.Address)].Description,
		}
		debtLimit := helpers.BValueWithFallbackUint64(multicallData.DebtLimit, 0)

		//Non exported fields, used for internal purposes
		currentStrategy.TotalDebt = helpers.BValueWithFallbackString(multicallData.TotalDebt, `0`)
		currentStrategy.DelegatedAssets = helpers.BValueWithFallbackString(multicallData.DelegatedAssets, `0`)
		currentStrategy.IsActive = multicallData.IsActive
		currentStrategy.InQueue = strategy.InQueue
		currentStrategy.DelegatedValue = strconv.FormatFloat(
			buildDelegated(
				currentStrategy.DelegatedAssets,
				int(vaultFromGraph.Token.Decimals),
				humanizedTokenPrice,
			), 'f', -1, 64,
		)

		//Compute the details about the strategy
		if withStrategiesDetails {
			currentStrategy.Details = &TStrategyDetails{}
			currentStrategy.Details.Protocols = strategiesFromMeta[common.HexToAddress(strategy.Address)].Protocols
			currentStrategy.Details.Keeper = common.HexToAddress(strategy.Keeper).String()
			currentStrategy.Details.Strategist = common.HexToAddress(strategy.Strategist).String()
			currentStrategy.Details.Rewards = common.HexToAddress(strategy.Rewards).String()
			currentStrategy.Details.HealthCheck = common.HexToAddress(strategy.HealthCheck).String()
			currentStrategy.Details.Version = strategy.ApiVersion
			currentStrategy.Details.InQueue = strategy.InQueue
			currentStrategy.Details.DoHealthCheck = strategy.DoHealthCheck
			currentStrategy.Details.EmergencyExit = strategy.EmergencyExit
			currentStrategy.Details.DebtLimit = debtLimit
			currentStrategy.Details.IsActive = multicallData.IsActive
			currentStrategy.Details.CreditAvailable = helpers.BValueWithFallbackString(multicallData.CreditAvailable, `0`)
			currentStrategy.Details.DebtOutstanding = helpers.BValueWithFallbackString(multicallData.DebtOutstanding, `0`)
			currentStrategy.Details.ExpectedReturn = helpers.BValueWithFallbackString(multicallData.ExpectedReturn, `0`)
			currentStrategy.Details.PerformanceFee = helpers.BValueWithFallbackUint64(multicallData.PerformanceFee, 0)
			currentStrategy.Details.Activation = helpers.BValueWithFallbackUint64(multicallData.Activation, 0)
			currentStrategy.Details.DebtRatio = helpers.BValueWithFallbackUint64(multicallData.DebtRatio, 0)
			currentStrategy.Details.RateLimit = helpers.BValueWithFallbackString(multicallData.RateLimit, `0`)
			currentStrategy.Details.MinDebtPerHarvest = helpers.BValueWithFallbackString(multicallData.MinDebtPerHarvest, `0`)
			currentStrategy.Details.MaxDebtPerHarvest = helpers.BValueWithFallbackString(multicallData.MaxDebtPerHarvest, `0`)
			currentStrategy.Details.EstimatedTotalAssets = helpers.BValueWithFallbackString(multicallData.EstimatedTotalAssets, `0`)
			currentStrategy.Details.DelegatedAssets = currentStrategy.DelegatedAssets
			currentStrategy.Details.DelegatedValue = currentStrategy.DelegatedValue
			currentStrategy.Details.KeepCRV = helpers.BValueWithFallbackUint64(multicallData.KeepCRV, 0)
			currentStrategy.Details.LastReport = helpers.BValueWithFallbackUint64(multicallData.LastReport, 0)
			currentStrategy.Details.TotalDebt = currentStrategy.TotalDebt
			currentStrategy.Details.TotalGain = helpers.BValueWithFallbackString(multicallData.TotalGain, `0`)
			currentStrategy.Details.TotalLoss = helpers.BValueWithFallbackString(multicallData.TotalLoss, `0`)
			currentStrategy.Details.APR = 0.0
			currentStrategy.Details.WithdrawalQueuePosition = helpers.BValueWithFallbackInt64(multicallData.WithdrawalQueuePosition, -1)

			if len(strategy.Reports) > 0 {
				var totalAPR float64
				for _, report := range strategy.Reports {
					if len(report.Results) > 0 {
						floatAPR, err := strconv.ParseFloat(report.Results[0].APR, 64)
						if err != nil {
							logs.Warning(err)
							floatAPR = 0
						}
						totalAPR += floatAPR
					}
				}
				floatTotalAPR := (totalAPR / float64(len(strategy.Reports))) * 100
				currentStrategy.Details.APR = floatTotalAPR
			}
		}

		//Compute the risk data
		if withStrategiesRisk {
			riskData, ok := strategiesFromRisk[common.HexToAddress(strategy.Address)]
			if !ok {
				riskData = models.TStrategyFromRisk{}
			}

			currentStrategy.Risk = &TStrategyRisk{}
			currentStrategy.Risk.RiskGroup = riskData.RiskGroup
			currentStrategy.Risk.TVLImpact = int(riskData.RiskScores.TVLImpact)
			currentStrategy.Risk.AuditScore = int(riskData.RiskScores.AuditScore)
			currentStrategy.Risk.CodeReviewScore = int(riskData.RiskScores.CodeReviewScore)
			currentStrategy.Risk.ComplexityScore = int(riskData.RiskScores.ComplexityScore)
			currentStrategy.Risk.LongevityImpact = int(riskData.RiskScores.LongevityImpact)
			currentStrategy.Risk.ProtocolSafetyScore = int(riskData.RiskScores.ProtocolSafetyScore)
			currentStrategy.Risk.TeamKnowledgeScore = int(riskData.RiskScores.TeamKnowledgeScore)
			currentStrategy.Risk.TestingScore = int(riskData.RiskScores.TestingScore)
		}

		if strategiesCondition == `absolute` &&
			currentStrategy.InQueue &&
			currentStrategy.IsActive &&
			currentStrategy.TotalDebt != `0` {
			strategies = append(strategies, currentStrategy)
		} else if strategiesCondition == `inQueue` && currentStrategy.InQueue {
			strategies = append(strategies, currentStrategy)
		} else if strategiesCondition == `debtLimit` && debtLimit == 0 {
			strategies = append(strategies, currentStrategy)
		}

	}
	return strategies
}
