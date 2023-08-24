package strategies

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/tokens"
)

var stratGroupErrorAlreadySent = make(map[uint64]map[string]bool)
var stakingData = make(map[uint64]map[string]models.TStaking)

func getTVLImpact(tvlUSDC *bigNumber.Float) int {
	tvl, _ := tvlUSDC.Float32()
	switch {
	case tvl == 0:
		return 0
	case tvl < 1_000_000:
		return 1
	case tvl < 10_000_000:
		return 2
	case tvl < 50_000_000:
		return 3
	case tvl < 100_000_000:
		return 4
	default:
		return 5
	}
}

func getLongevityImpact(activation *bigNumber.Int) int {
	if activation == nil || activation.IsZero() {
		return 5
	}
	activationUnix := time.Unix(activation.Int64(), 0)
	longevity := time.Since(activationUnix)
	days := time.Hour * 24
	switch {
	case longevity > 240*days:
		return 1
	case longevity > 120*days:
		return 2
	case longevity > 30*days:
		return 3
	case longevity > 7*days:
		return 4
	default:
		return 5
	}
}

func getMedianScore(group models.TStrategyGroupFromRisk) int {
	if group.StackingRewardTVLScore > 0 {
		scores := stats.LoadRawData([]int{
			group.AuditScore,
			group.CodeReviewScore,
			group.ComplexityScore,
			group.ProtocolSafetyScore,
			group.TeamKnowledgeScore,
			group.TestingScore,
			group.StackingRewardTVLScore,
		})
		median, _ := stats.Median(scores)
		return int(median)
	}
	scores := stats.LoadRawData([]int{
		group.AuditScore,
		group.CodeReviewScore,
		group.ComplexityScore,
		group.ProtocolSafetyScore,
		group.TeamKnowledgeScore,
		group.TestingScore,
	})
	median, _ := stats.Median(scores)
	return int(median)
}

func getMedianAllocation(group models.TStrategyGroupFromRisk) *bigNumber.Float {
	median := getMedianScore(group)
	switch {
	case median <= 1:
		return bigNumber.NewFloat(100_000_000)
	case median <= 2:
		return bigNumber.NewFloat(50_000_000)
	case median <= 3:
		return bigNumber.NewFloat(10_000_000)
	case median <= 4:
		return bigNumber.NewFloat(1_000_000)
	default:
		return bigNumber.NewFloat(0)
	}
}

func getAllocationStatus(group models.TStrategyGroupFromRisk) string {
	median := getMedianScore(group)
	medianAllocation := getMedianAllocation(group)
	tvl := group.Allocation.CurrentTVL
	diff := bigNumber.NewFloat(0).Sub(medianAllocation, tvl)

	// under allocated - Green
	if diff.Sign() > 0 {
		return "Green"
	}
	diffAllowed := bigNumber.NewFloat(0)
	switch {
	case median <= 1:
		diffAllowed = bigNumber.NewFloat(100_000_000)
	case median <= 2:
		diffAllowed = bigNumber.NewFloat(50_000_000)
	case median <= 3:
		diffAllowed = bigNumber.NewFloat(40_000_000)
	case median <= 4:
		diffAllowed = bigNumber.NewFloat(9_000_000)
	case median <= 5:
		diffAllowed = bigNumber.NewFloat(1_000_000)
	}
	// over allocated but within allowed range - Yellow
	if diff.Sub(diffAllowed, diff).Sign() > 0 {
		return "Yellow"
	}
	// over allocated and out of range - Red
	return "Red"
}

func getStrategyGroup(chainID uint64, strategy *models.TStrategy) *models.TStrategyGroupFromRisk {
	toLowerName := strings.ToLower(strategy.Name)
	groups := ListStrategiesRiskGroups(chainID)

	//First check if the address is in any group
	for _, group := range groups {
		if helpers.Contains(helpers.AddressToString(group.Criteria.Strategies), strategy.Address.Hex()) {
			return group
		}
	}

	for _, group := range groups {
		// check if nameLike and exclude intersect
		if helpers.Intersects(group.Criteria.NameLike, group.Criteria.Exclude) {
			toLowerNameLike := helpers.ToLower(group.Criteria.NameLike)
			toLowerExclude := helpers.ToLower(group.Criteria.Exclude)
			for _, nameLike := range toLowerNameLike {
				// if the nameLike is more specific
				if strings.Contains(toLowerName, nameLike) &&
					helpers.ContainsSubString(toLowerExclude, nameLike) {
					return group
				}
			}
		}

		// check exclude
		toLowerExclude := helpers.ToLower(group.Criteria.Exclude)
		if helpers.ContainsSubString(toLowerExclude, toLowerName) {
			continue
		}
		if helpers.Contains(group.Criteria.Exclude, strings.ToLower(strategy.Address.String())) {
			continue
		}

		// check nameLike
		toLowerNameLike := helpers.ToLower(group.Criteria.NameLike)
		if helpers.ContainsSubString(toLowerNameLike, toLowerName) {
			return group
		}
	}
	return nil
}

func getDefaultRiskGroup() models.TStrategyFromRisk {
	return models.TStrategyFromRisk{
		RiskGroup: "Others",
		RiskScore: 5,
		RiskDetails: models.TStrategyFromRiskRiskDetails{
			TVLImpact:           5,
			AuditScore:          5,
			CodeReviewScore:     5,
			ComplexityScore:     5,
			LongevityImpact:     5,
			ProtocolSafetyScore: 5,
			TeamKnowledgeScore:  5,
			TestingScore:        5,
		},
		Allocation: &models.TStrategyAllocation{
			Status:          "Green",
			CurrentTVL:      bigNumber.NewFloat(0),
			AvailableTVL:    bigNumber.NewFloat(0),
			CurrentAmount:   bigNumber.NewFloat(0),
			AvailableAmount: bigNumber.NewFloat(0),
		},
	}
}

func GetStakingData(chainID uint64, vaultAddress common.Address) models.TStaking {
	if _, ok := stakingData[chainID]; !ok {
		return models.TStaking{}
	}
	if score, ok := stakingData[chainID][vaultAddress.Hex()]; ok {
		return score
	}
	return models.TStaking{}
}

/**************************************************************************************************
** RetrieveAllRisksGroupsFromFiles will read all files in the /risks/network directory for a given
** chainID and store them in the _strategyRiskGroupMap global variable.
** The goal of this function is to get a list of all the group risk information for a given
** chainID.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**************************************************************************************************/
func RetrieveAllRisksGroupsFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/risks/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		traces.
			Capture(`warn`, `impossible to read risks files on chain `+chainIDStr).
			SetEntity(`strategy`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			Send()
		return
	}
	for _, content := range content {
		riskGroup := models.TStrategyGroupFromRisk{}
		if err := json.Unmarshal(content, &riskGroup); err != nil {
			traces.
				Capture(`warn`, `impossible to unmarshall risks files response body `+chainIDStr).
				SetEntity(`strategy`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetExtra(`content`, string(content)).
				Send()
			return
		}
		riskGroup.ChainID = chainID
		setRiskGroupInMap(chainID, &riskGroup)
	}
}

func computeRiskGroupAllocation(chainID uint64) {
	if stratGroupErrorAlreadySent[chainID] == nil {
		stratGroupErrorAlreadySent[chainID] = map[string]bool{}
	}
	//This will ensure we are working with clean data
	groups := ListStrategiesRiskGroups(chainID)
	for _, group := range groups {
		group.Allocation = &models.TStrategyAllocation{}
		group.Allocation.AvailableTVL = getMedianAllocation(*group)
		setRiskGroupInMap(chainID, group)
	}

	strategies := ListStrategies(chainID)
	for _, strategy := range strategies {
		strategyGroup := getStrategyGroup(chainID, strategy)
		if strategyGroup == nil {
			if !stratGroupErrorAlreadySent[chainID][strategy.Name] {
				traces.
					Capture(`warn`, `impossible to find stratGroup for strategy `+strategy.Name).
					SetEntity(`strategy`).
					SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
					SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
					SetTag(`strategyAddress`, strategy.Address.Hex()).
					SetTag(`strategyName`, strategy.Name).
					Send()
				stratGroupErrorAlreadySent[chainID][strategy.Name] = true
			}
			continue
		}

		tokenData, ok := tokens.FindUnderlyingForVault(chainID, strategy.VaultAddress)
		if !ok {
			traces.
				Capture(`warn`, `impossible to find token for vault `+strategy.VaultAddress.Hex()).
				SetEntity(`strategy`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`strategyAddress`, strategy.Address.Hex()).
				SetTag(`strategyName`, strategy.Name).
				Send()
			continue
		}

		tokenPrice, ok := prices.FindPrice(chainID, tokenData.Address)
		if !ok {
			traces.
				Capture(`warn`, `impossible to find tokenPrice for token `+tokenData.Address.Hex()+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`strategy`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`strategyAddress`, strategy.Address.Hex()).
				SetTag(`strategyName`, strategy.Name).
				Send()
			tokenPrice = bigNumber.NewInt(0)
		}
		_, price := helpers.FormatAmount(tokenPrice.String(), 6)
		_, amount := helpers.FormatAmount(strategy.EstimatedTotalAssets.String(), int(tokenData.Decimals))
		tvl := bigNumber.NewFloat(0).Mul(amount, price)

		// Underlying was hacked, force TVL to 0
		if addresses.Equals(strategy.VaultAddress, `0x718AbE90777F5B778B52D553a5aBaa148DD0dc5D`) {
			tvl = bigNumber.NewFloat(0)
		}

		if stackingData, hasStackingData := stakingData[chainID][strategy.VaultAddress.Hex()]; hasStackingData {
			impactedStrategies := []common.Address{
				common.HexToAddress("0x877D377b289AEa8e0F8aC2470252D226417D9e21"),
				common.HexToAddress("0xe82DEb62412DB78D00Cae77BE3d1334e26034Cf6"),
				common.HexToAddress("0x1d3BE43C5621bdA29aB24101c54d434B9b3f29c8"),
				common.HexToAddress("0xf8aD69d578bd58c7d3Ff643A22355f0BFd5cA12A"),
			}
			if helpers.Contains(impactedStrategies, strategy.Address) {
				strategyGroup.StackingRewardTVLScore = stackingData.Risk
			}
		}
		strategyGroup.Allocation.CurrentTVL = bigNumber.NewFloat(0).Add(strategyGroup.Allocation.CurrentTVL, tvl)
		strategyGroup.Allocation.CurrentAmount = bigNumber.NewFloat(0).Add(strategyGroup.Allocation.CurrentAmount, amount)
		if tokenPrice.Sign() <= 0 {
			strategyGroup.Allocation.AvailableTVL = bigNumber.NewFloat(0)
			strategyGroup.Allocation.AvailableAmount = bigNumber.NewFloat(0)
		} else {
			strategyGroup.Allocation.AvailableTVL = bigNumber.NewFloat(0).Sub(strategyGroup.Allocation.AvailableTVL, tvl)
			strategyGroup.Allocation.AvailableAmount = bigNumber.NewFloat(0).Quo(strategyGroup.Allocation.AvailableTVL, price)
		}
		strategyGroup.Allocation.Status = getAllocationStatus(*strategyGroup)
		setRiskGroupInMap(chainID, strategyGroup)
	}
}

func InitRiskScore(chainID uint64) {
	if chainID == 10 {
		allStackingPoolAdded := events.HandleStakingPoolAdded(chainID, 0, nil)
		calls := []ethereum.Call{}
		for _, pool := range allStackingPoolAdded {
			currentToken, ok := tokens.FindToken(chainID, pool.VaultAddress)
			if !ok {
				logs.Warning(`[InitRiskScore]`, `impossible to find token for pool address`, pool.StackingPoolAddress.Hex())
				continue
			}
			if !tokens.IsVault(currentToken) {
				logs.Warning(`[InitRiskScore]`, `token is not a vault`, pool.VaultAddress.Hex())
				continue
			}
			calls = append(calls, multicalls.GetTotalSupply(pool.StackingPoolAddress.Hex(), pool.StackingPoolAddress))
			calls = append(calls, multicalls.GetDecimals(pool.StackingPoolAddress.Hex(), pool.VaultAddress))
			calls = append(calls, multicalls.GetPriceUsdcRecommendedCall(pool.StackingPoolAddress.Hex(), env.LENS_ADDRESSES[chainID], pool.VaultAddress))
		}

		response := multicalls.Perform(chainID, calls, nil)
		for _, pool := range allStackingPoolAdded {
			totalSupply := helpers.DecodeBigInt(response[pool.StackingPoolAddress.Hex()+`totalSupply`])
			tokenPrice := helpers.DecodeBigInt(response[pool.StackingPoolAddress.Hex()+`getPriceUsdcRecommended`])
			decimals := helpers.DecodeUint64(response[pool.StackingPoolAddress.Hex()+`decimals`])

			_, price := helpers.FormatAmount(tokenPrice.String(), 6)
			_, amount := helpers.FormatAmount(totalSupply.String(), int(decimals))

			if _, ok := stakingData[chainID]; !ok {
				stakingData[chainID] = map[string]models.TStaking{}
			}
			tvl, _ := bigNumber.NewFloat(0).Mul(amount, price).Float64()
			stakingData[chainID][pool.VaultAddress.Hex()] = models.TStaking{
				Available: true,
				Address:   pool.StackingPoolAddress,
				Risk:      getTVLImpact(bigNumber.NewFloat(0).Mul(amount, price)),
				TVL:       tvl,
			}
			// logs.Pretty(`[InitRiskScore]`, pool.Token.Hex(), stakingData[chainID][pool.Token.Hex()].Risk, bigNumber.NewFloat(0).Mul(amount, price).String(), amount, price)
		}
	}
	computeRiskGroupAllocation(chainID)
}
