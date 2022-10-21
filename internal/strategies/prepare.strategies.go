package strategies

import (
	"strconv"

	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

func buildDelegated(delegatedBalanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, delegatedTVL := helpers.FormatAmount(delegatedBalanceToken.String(), decimals)
	fHumanizedTVLPrice, _ := bigNumber.NewFloat().Mul(delegatedTVL, humanizedPrice).Float64()
	return fHumanizedTVLPrice
}

// BuildStrategies prepare the TStrategy struct based of a bunch of data
func BuildStrategies(
	chainID uint64,
	withStrategiesDetails bool,
	withStrategiesRisk bool,
	strategiesCondition string,
	humanizedTokenPrice *bigNumber.Float,
	vaultFromGraph models.TVaultFromGraph,
) []TStrategy {
	strategies := []TStrategy{}
	strategiesFromMeta := meta.Store.StrategiesFromMeta[chainID]
	strategiesFromMulticall := Store.StrategyMultiCallData[chainID]
	strategiesFromRisk := Store.StrategiesFromRisk[chainID]

	for _, strategy := range vaultFromGraph.Strategies {
		multicallData, ok := strategiesFromMulticall[strategy.Address]
		if !ok {
			multicallData = models.TStrategyMultiCallData{}
		}

		currentStrategy := TStrategy{
			Address:     strategy.Address,
			Name:        strategy.Name,
			Description: strategiesFromMeta[strategy.Address].Description,
		}
		debtLimit := multicallData.DebtLimit.Uint64()

		//Non exported fields, used for internal purposes
		withdrawalQueuePosition := bigNumber.NewInt().Safe(multicallData.WithdrawalQueuePosition, bigNumber.NewInt(-1)).Int64()
		currentStrategy.TotalDebt = multicallData.TotalDebt
		currentStrategy.DelegatedAssets = multicallData.DelegatedAssets
		currentStrategy.IsActive = multicallData.IsActive
		currentStrategy.InQueue = withdrawalQueuePosition > -1
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
			currentStrategy.Details.Protocols = strategiesFromMeta[strategy.Address].Protocols
			currentStrategy.Details.Keeper = strategy.Keeper
			currentStrategy.Details.Strategist = strategy.Strategist
			currentStrategy.Details.Rewards = strategy.Rewards
			currentStrategy.Details.HealthCheck = common.HexToAddress(strategy.HealthCheck)
			currentStrategy.Details.Version = strategy.ApiVersion
			currentStrategy.Details.InQueue = strategy.InQueue
			currentStrategy.Details.DoHealthCheck = strategy.DoHealthCheck
			currentStrategy.Details.EmergencyExit = strategy.EmergencyExit
			currentStrategy.Details.DebtLimit = debtLimit
			currentStrategy.Details.IsActive = multicallData.IsActive

			currentStrategy.Details.CreditAvailable = multicallData.CreditAvailable
			currentStrategy.Details.DebtOutstanding = multicallData.DebtOutstanding
			currentStrategy.Details.ExpectedReturn = multicallData.ExpectedReturn
			currentStrategy.Details.RateLimit = multicallData.RateLimit
			currentStrategy.Details.MinDebtPerHarvest = multicallData.MinDebtPerHarvest
			currentStrategy.Details.MaxDebtPerHarvest = multicallData.MaxDebtPerHarvest
			currentStrategy.Details.EstimatedTotalAssets = multicallData.EstimatedTotalAssets
			currentStrategy.Details.TotalGain = multicallData.TotalGain
			currentStrategy.Details.TotalLoss = multicallData.TotalLoss
			currentStrategy.Details.TotalDebt = currentStrategy.TotalDebt
			currentStrategy.Details.DelegatedAssets = currentStrategy.DelegatedAssets
			currentStrategy.Details.DelegatedValue = bigNumber.NewInt().SetString(currentStrategy.DelegatedValue)

			currentStrategy.Details.PerformanceFee = (multicallData.PerformanceFee).Uint64()
			currentStrategy.Details.Activation = (multicallData.Activation).Uint64()
			currentStrategy.Details.DebtRatio = (multicallData.DebtRatio).Uint64()
			currentStrategy.Details.KeepCRV = (multicallData.KeepCRV).Uint64()
			currentStrategy.Details.LastReport = (multicallData.LastReport).Uint64()

			currentStrategy.Details.APR = 0.0
			currentStrategy.Details.WithdrawalQueuePosition = withdrawalQueuePosition

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
			riskData, ok := strategiesFromRisk[strategy.Address]
			if !ok {
				riskData = TStrategyFromRisk{}
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
			currentStrategy.Risk.Allocation = riskData.Allocation
		}

		if strategiesCondition == `absolute` &&
			currentStrategy.InQueue &&
			currentStrategy.IsActive &&
			!currentStrategy.TotalDebt.IsZero() {
			strategies = append(strategies, currentStrategy)
		} else if strategiesCondition == `inQueue` && currentStrategy.InQueue {
			strategies = append(strategies, currentStrategy)
		} else if strategiesCondition == `debtLimit` && debtLimit == 0 {
			strategies = append(strategies, currentStrategy)
		}

	}
	return strategies
}
