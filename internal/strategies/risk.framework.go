package strategies

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/montanaflynn/stats"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
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
	scoreValues := []int{
		group.AuditScore,
		group.CodeReviewScore,
		group.ComplexityScore,
		group.ProtocolSafetyScore,
		group.TeamKnowledgeScore,
		group.TestingScore,
	}
	if group.StackingRewardTVLScore > 0 {
		scoreValues = append(scoreValues, group.StackingRewardTVLScore)
	}
	scores := stats.LoadRawData(scoreValues)
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

func getMediamDiffAllowed(group models.TStrategyGroupFromRisk) *bigNumber.Float {
	median := getMedianScore(group)
	switch {
	case median <= 1:
		return bigNumber.NewFloat(100_000_000)
	case median <= 2:
		return bigNumber.NewFloat(50_000_000)
	case median <= 3:
		return bigNumber.NewFloat(40_000_000)
	case median <= 4:
		return bigNumber.NewFloat(9_000_000)
	case median <= 5:
		return bigNumber.NewFloat(1_000_000)
	default:
		return bigNumber.NewFloat(0)
	}
}

/**********************************************************************************************
** The getAllocationStatus function takes a strategy group as input and calculates the median
** allocation and the allowed median difference for the group. It then retrieves the current
** Total Value Locked (TVL) for the group and calculates the difference between the median
** allocation and the TVL.
** If the difference is positive, the function returns "Green".
** If the difference is greater than the allowed difference, the function returns "Yellow".
** In all other cases, the function returns "Red".
** This function is used to determine the allocation status of a strategy group based on its
** current TVL and median allocation.
**********************************************************************************************/
func getAllocationStatus(group models.TStrategyGroupFromRisk) string {
	medianAllocation := getMedianAllocation(group)
	diffAllowed := getMediamDiffAllowed(group)
	tvl := group.Allocation.CurrentTVL
	diff := bigNumber.NewFloat(0).Sub(medianAllocation, tvl)

	if diff.Sign() > 0 {
		return "Green"
	}
	if diff.Sub(diffAllowed, diff).Sign() > 0 {
		return "Yellow"
	}
	return "Red"
}

/**********************************************************************************************
** The retrieveStrategyGroup function takes a chainID and a strategy as input. It lists all the
** risk groups for the given chainID and checks if the strategy's address is in any of the groups.
** If it finds a match, it returns the group. If no match is found, it returns nil.
** This function is used to retrieve the risk group that a strategy belongs to.
**********************************************************************************************/
func retrieveStrategyGroup(chainID uint64, strategy *models.TStrategy) *models.TStrategyGroupFromRisk {
	groups := ListStrategiesRiskGroups(chainID)
	for _, group := range groups {
		if helpers.Contains(helpers.AddressToString(group.Strategies), strategy.Address.Hex()) {
			return group
		}
	}

	return nil
}

/**********************************************************************************************
** The useStrategyGroup function takes a list of strategy groups and a strategy as input. It
** iterates over the strategy groups and checks if the strategy's address is in any of the groups.
** If it finds a match, it returns the group. If no match is found, it returns nil.
** This function is used to find the risk group that a strategy belongs to.
**********************************************************************************************/
func useStrategyGroup(groups []*models.TStrategyGroupFromRisk, strategy *models.TStrategy) *models.TStrategyGroupFromRisk {
	for _, group := range groups {
		if helpers.Contains(helpers.AddressToString(group.Strategies), strategy.Address.Hex()) {
			return group
		}
	}
	return nil
}

/**********************************************************************************************
** The getDefaultRiskGroup function returns a default risk group with a risk score of 5 and a
** risk group name of "Others". This function is used when a strategy does not belong to any
** existing risk group. The returned risk group has all its risk details set to 5 and its
** allocation status set to "Green". The current, available TVL and amount are all set to 0.
**********************************************************************************************/
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

/**********************************************************************************************
** RetrieveAllRisksGroupsFromFiles reads all risk group files for a given chainID from the
** /risks/networks/ directory.
** It unmarshals the content of each file into a TStrategyGroupFromRisk model and stores it in
** the _strategyRiskGroupMap global variable.
** If any error occurs during the reading or unmarshalling process, it logs the error and
** returns.
** The function also sets the ChainID of each risk group to the given chainID.
*********************************************************************************************
 */
func RetrieveAllRisksGroupsFromFiles(chainID uint64) {
	chainIDStr := strconv.FormatUint(chainID, 10)
	content, _, err := helpers.ReadAllFilesInDir(env.BASE_DATA_PATH+`/risks/networks/`+chainIDStr+`/`, `.json`)
	if err != nil {
		logs.Error(`[RetrieveAllRisksGroupsFromFiles] impossible to read risks files ` + chainIDStr)
		return
	}
	for _, content := range content {
		riskGroup := models.TStrategyGroupFromRisk{}
		if err := json.Unmarshal(content, &riskGroup); err != nil {
			logs.Error(`[RetrieveAllRisksGroupsFromFiles] impossible to unmarshal risks files ` + chainIDStr)
			return
		}
		riskGroup.ChainID = chainID
		setStrategyRiskGroup(chainID, &riskGroup)
	}
}

func populateAllocation(strategy *models.TStrategy, group *models.TStrategyGroupFromRisk) {
	tvl, amount, price := computeTVL(strategy)

	/**********************************************************************************************
	** Exception section: If the vault address matches the specified address, the TVL is forced to
	** 0.
	** This is done to handle specific cases where the underlying was hacked.
	**********************************************************************************************/
	if addresses.Equals(strategy.VaultAddress, `0x718AbE90777F5B778B52D553a5aBaa148DD0dc5D`) {
		tvl = bigNumber.NewFloat(0)
	}

	/**********************************************************************************************
	** The following block of code is responsible for updating the current Total Value Locked (TVL)
	** and the current amount of the strategy's vault. The TVL is calculated by multiplying the
	** amount of tokens in the vault by the price of the token. The current amount is the amount
	** of tokens in the vault. These values are crucial for risk assessment and allocation computations.
	**********************************************************************************************/
	group.Allocation.CurrentTVL = bigNumber.NewFloat(0).Add(group.Allocation.CurrentTVL, tvl)
	group.Allocation.CurrentAmount = bigNumber.NewFloat(0).Add(group.Allocation.CurrentAmount, amount)
	if price.Sign() <= 0 {
		group.Allocation.AvailableTVL = bigNumber.NewFloat(0)
		group.Allocation.AvailableAmount = bigNumber.NewFloat(0)
	} else {
		group.Allocation.AvailableTVL = bigNumber.NewFloat(0).Sub(group.Allocation.AvailableTVL, tvl)
		group.Allocation.AvailableAmount = bigNumber.NewFloat(0).Quo(group.Allocation.AvailableTVL, price)
	}
	group.Allocation.Status = getAllocationStatus(*group)
}

func populateRisk(chainID uint64) {
	if stratGroupErrorAlreadySent[chainID] == nil {
		stratGroupErrorAlreadySent[chainID] = map[string]bool{}
	}

	/**********************************************************************************************
	** This loop iterates over all the risk groups for the given chainID. For each group, it
	** initializes the Allocation field and sets the AvailableTVL to the median allocation of the group.
	**********************************************************************************************/
	groups := ListStrategiesRiskGroups(chainID)
	for _, group := range groups {
		group.Allocation = &models.TStrategyAllocation{}
		group.Allocation.AvailableTVL = getMedianAllocation(*group)
	}

	strategies := ListStrategies(chainID)
	for _, strategy := range strategies {
		group := useStrategyGroup(groups, strategy)
		if group == nil {
			if !stratGroupErrorAlreadySent[chainID][strategy.Name] {
				logs.Error(`[populateRisk] impossible to find a risk group for strategy ` + strategy.Address.Hex() + ` (` + strategy.Name + `) on chain ` + strconv.FormatUint(chainID, 10))
				stratGroupErrorAlreadySent[chainID][strategy.Name] = true
			}
			continue
		}
		populateAllocation(strategy, group)

		if stackingData, hasStackingData := stakingData[chainID][strategy.VaultAddress.Hex()]; hasStackingData {
			impactedStrategies := []common.Address{
				common.HexToAddress("0x877D377b289AEa8e0F8aC2470252D226417D9e21"),
				common.HexToAddress("0xe82DEb62412DB78D00Cae77BE3d1334e26034Cf6"),
				common.HexToAddress("0x1d3BE43C5621bdA29aB24101c54d434B9b3f29c8"),
				common.HexToAddress("0xf8aD69d578bd58c7d3Ff643A22355f0BFd5cA12A"),
			}
			if helpers.Contains(impactedStrategies, strategy.Address) {
				group.StackingRewardTVLScore = stackingData.Risk
			}
		}
	}

	/**********************************************************************************************
	** This loop iterates over all the risk groups for the given chainID. For each group, it
	** sets the strategy risk group. This is done after the risk group allocation computation
	** to ensure that the risk group data is up-to-date.
	**********************************************************************************************/
	for _, group := range groups {
		setStrategyRiskGroup(chainID, group)
	}
}

func InitRiskScore(chainID uint64) {
	calculateStakingMetrics(chainID)
	populateRisk(chainID)
}
