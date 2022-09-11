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

func getTVLImpact(strategyData models.TStrategyMultiCallData) float64 {
	if strategyData.EstimatedTotalAssets.Cmp(big.NewInt(0)) == 0 {
		return 0
	} else if strategyData.EstimatedTotalAssets.Cmp(big.NewInt(1_000_000)) < 0 {
		return 1
	} else if strategyData.EstimatedTotalAssets.Cmp(big.NewInt(10_000_000)) < 0 {
		return 2
	} else if strategyData.EstimatedTotalAssets.Cmp(big.NewInt(50_000_000)) < 0 {
		return 3
	} else if strategyData.EstimatedTotalAssets.Cmp(big.NewInt(100_000_000)) < 0 {
		return 4
	}
	return 5
}

func getLongevityImpact(strategyData models.TStrategyMultiCallData) float64 {
	if strategyData.Activation == nil {
		return 5
	}
	var longevity big.Int
	now := big.NewInt(time.Now().Unix())
	longevity.Sub(now, strategyData.Activation)

	var days int64 = 60 * 60 * 24
	if longevity.Cmp(big.NewInt(240*days)) > 0 {
		return 1
	} else if longevity.Cmp(big.NewInt(120*days)) > 0 {
		return 2
	} else if longevity.Cmp(big.NewInt(30*days)) > 0 {
		return 3
	} else if longevity.Cmp(big.NewInt(7*days)) > 0 {
		return 4
	}
	return 5
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

	groupLoop:
		for _, group := range groups {
			// Exclude strategies
			if len(group.Criteria.Exclude) > 0 {
				for _, stratExclude := range group.Criteria.Exclude {
					if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(stratExclude)) {
						continue groupLoop
					}
				}
			}
			// Include strategies
			if len(group.Criteria.Strategies) > 0 {
				for _, stratInclude := range group.Criteria.Strategies {
					if helpers.ContainsAddress(strat.Addresses, stratInclude) {
						stratGroup = group
						break groupLoop
					}
				}
			}

			// Include nameLikes
			if len(group.Criteria.NameLike) > 0 {
				for _, nameLike := range group.Criteria.NameLike {
					if strings.Contains(strings.ToLower(strat.Name), strings.ToLower(nameLike)) {
						stratGroup = group
						break groupLoop
					}
				}
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
