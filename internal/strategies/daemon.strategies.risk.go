package strategies

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/store"
)

func excludeNameLike(strat meta.TStrategyFromMeta, group models.TStrategyGroupFromRisk) bool {
	if len(group.Criteria.Exclude) > 0 {
		for _, stratExclude := range group.Criteria.Exclude {
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(stratExclude)) {
				return true
			}
		}
	}
	return false
}

func includeAddress(strat meta.TStrategyFromMeta, group models.TStrategyGroupFromRisk) bool {
	if len(group.Criteria.Strategies) > 0 {
		for _, stratInclude := range group.Criteria.Strategies {
			if helpers.ContainsAddress(strat.Addresses, stratInclude) {
				return true
			}
		}
	}
	return false
}

func includeNameLike(strat meta.TStrategyFromMeta, group models.TStrategyGroupFromRisk) bool {
	if len(group.Criteria.NameLike) > 0 {
		for _, nameLike := range group.Criteria.NameLike {
			if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(nameLike)) {
				return true
			}
		}
	}
	return false
}

func getTVLImpact(strategyData models.TStrategyMultiCallData) float64 {
	tvl := strategyData.EstimatedTotalAssets.Int64()
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

func getLongevityImpact(strategyData models.TStrategyMultiCallData) float64 {
	if strategyData.Activation == nil || strategyData.Activation == big.NewInt(0) {
		return 5
	}
	activation := time.Unix(strategyData.Activation.Int64(), 0)
	longevity := time.Since(activation)
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

// FetchStrategiesFromRisk fetches the strategies information from the Risk Framework for a given chainID
// and store the result to the global variable StrategiesFromRisk for later use.
func FetchStrategiesFromRisk(chainID uint64) {
	// Read data from the risk framework json file
	content, err := ioutil.ReadFile(helpers.BASE_DATA_PATH + "/risks.json")
	if err != nil {
		logs.Warning("Error fetching information from the Risk Framework")
		return
	}
	groups := []models.TStrategyGroupFromRisk{}
	if err := json.Unmarshal(content, &groups); err != nil {
		logs.Warning("Error unmarshalling response body from the Risk Framework")
		return
	}

	// For each strategy in the meta, parse the strategy group scores
	if Store.StrategiesFromRisk[chainID] == nil {
		Store.StrategiesFromRisk[chainID] = make(map[common.Address]models.TStrategyFromRisk)
	}
	strategies, ok := meta.Store.StrategiesFromMeta[chainID]
	if !ok {
		logs.Warning("Error reading meta information from the Yearn Meta Files")
		return
	}
	for _, strat := range strategies {
		var stratGroup models.TStrategyGroupFromRisk
		for _, group := range groups {
			if excludeNameLike(strat, group) {
				continue
			}
			if includeAddress(strat, group) || includeNameLike(strat, group) {
				stratGroup = group
				break
			}
		}
		// Skip if no group was found
		if stratGroup.GroupID == "" {
			break
		}

		// Loop over addresses
		for _, address := range strat.Addresses {
			// Fetch activation and tvl from multicall
			addressHex := common.HexToAddress(address)
			data, ok := Store.StrategyMultiCallData[chainID][addressHex]
			if !ok {
				logs.Error("Error fetching strategy data from multicall")
				return
			}

			// Store strategy to DB
			strategy := models.TStrategyFromRisk{
				RiskScores: models.TStrategyFromRiskRiskScores{
					TVLImpact:           getTVLImpact(data),
					AuditScore:          stratGroup.AuditScore,
					CodeReviewScore:     stratGroup.CodeReviewScore,
					ComplexityScore:     stratGroup.ComplexityScore,
					LongevityImpact:     getLongevityImpact(data),
					ProtocolSafetyScore: stratGroup.ProtocolSafetyScore,
					TeamKnowledgeScore:  stratGroup.TeamKnowledgeScore,
					TestingScore:        stratGroup.TestingScore,
				},
			}
			Store.StrategiesFromRisk[chainID][addressHex] = strategy
		}
	}
	store.SaveInDBForChainID(`StrategiesFromRisk`, chainID, Store.StrategiesFromRisk[chainID])
}

// LoadRiskStrategies will reload the risk strategies data store from the last state of the local Badger Database
func LoadRiskStrategies(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]models.TStrategyFromRisk)
	if err := store.LoadFromDBForChainID(`StrategiesFromRisk`, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		Store.StrategiesFromRisk[chainID] = temp
		logs.Success("Data loaded for the risk data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
