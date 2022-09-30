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
		debtLimit := helpers.SafeBigInt(multicallData.DebtLimit).Uint64()

		//Non exported fields, used for internal purposes
		currentStrategy.TotalDebt = helpers.SafeBigInt(multicallData.TotalDebt).String()
		currentStrategy.DelegatedAssets = helpers.SafeBigInt(multicallData.DelegatedAssets).String()
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
			currentStrategy.Details.CreditAvailable = helpers.SafeBigInt(multicallData.CreditAvailable).String()
			currentStrategy.Details.DebtOutstanding = helpers.SafeBigInt(multicallData.DebtOutstanding).String()
			currentStrategy.Details.ExpectedReturn = helpers.SafeBigInt(multicallData.ExpectedReturn).String()
			currentStrategy.Details.PerformanceFee = helpers.SafeBigInt(multicallData.PerformanceFee).Uint64()
			currentStrategy.Details.Activation = helpers.SafeBigInt(multicallData.Activation).Uint64()
			currentStrategy.Details.DebtRatio = helpers.SafeBigInt(multicallData.DebtRatio).Uint64()
			currentStrategy.Details.RateLimit = helpers.SafeBigInt(multicallData.RateLimit).String()
			currentStrategy.Details.MinDebtPerHarvest = helpers.SafeBigInt(multicallData.MinDebtPerHarvest).String()
			currentStrategy.Details.MaxDebtPerHarvest = helpers.SafeBigInt(multicallData.MaxDebtPerHarvest).String()
			currentStrategy.Details.EstimatedTotalAssets = helpers.SafeBigInt(multicallData.EstimatedTotalAssets).String()
			currentStrategy.Details.DelegatedAssets = helpers.SafeString(currentStrategy.DelegatedAssets, ``)
			currentStrategy.Details.DelegatedValue = helpers.SafeString(currentStrategy.DelegatedValue, ``)
			currentStrategy.Details.KeepCRV = helpers.SafeBigInt(multicallData.KeepCRV).Uint64()
			currentStrategy.Details.LastReport = helpers.SafeBigInt(multicallData.LastReport).Uint64()
			currentStrategy.Details.TotalDebt = currentStrategy.TotalDebt
			currentStrategy.Details.TotalGain = helpers.SafeBigInt(multicallData.TotalGain).String()
			currentStrategy.Details.TotalLoss = helpers.SafeBigInt(multicallData.TotalLoss).String()
			currentStrategy.Details.APR = 0.0
			currentStrategy.Details.WithdrawalQueuePosition = helpers.SafeBigInt(multicallData.WithdrawalQueuePosition, big.NewInt(-1)).Int64()

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
