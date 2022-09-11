package strategies

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/models"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// FetchStrategiesFromRisk fetches the strategies information from the Risk Framework for a given chainID
// and store the result to the global variable StrategiesFromRisk for later use.
func FetchStrategiesFromRisk(chainID uint64) {
	// Get the information from the Risk Framework
	resp, err := http.Get(utils.RISK_BASE_URL)
	if err != nil {
		logs.Error("Error fetching information from the Risk Framework", err)
		return
	}

	// Defer the closing of the response body to avoid memory leaks
	defer resp.Body.Close()

	// Read the response body and store it in the body variable
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("Error reading response body from the Risk Framework", err)
		return
	}

	// Unmarshal the response body into the variable StrategyGroupsFromRisk. Body is a byte array,
	groups := []models.TStrategyGroupFromRisk{}
	if err := json.Unmarshal(body, &groups); err != nil {
		logs.Error("Error unmarshalling response body from the Risk Framework", err)
		return
	}

	// For each strategy in the meta, parse the strategy group scores
	if store.StrategiesFromRisk[chainID] == nil {
		store.StrategiesFromRisk[chainID] = make(map[common.Address]models.TStrategyFromRisk)
	}
	strategies, ok := store.StrategiesFromMeta[chainID]
	if !ok {
		logs.Error("Error fetching strategies from meta")
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
					if utils.ContainsAddress(strat.Addresses, stratInclude) {
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
			addressHex := common.HexToAddress(address)

			// Fetch activation and tvl from multicall
			data, ok := store.StrategyMultiCallData[chainID][addressHex]
			if !ok {
				logs.Error("Error fetching strategy data from multicall")
				return
			}

			// TVL Impact
			var TVLImpact float64 = 5
			if data.EstimatedTotalAssets.Cmp(big.NewInt(0)) == 0 {
				TVLImpact = 0
			} else if data.EstimatedTotalAssets.Cmp(big.NewInt(1_000_000)) < 0 {
				TVLImpact = 1
			} else if data.EstimatedTotalAssets.Cmp(big.NewInt(10_000_000)) < 0 {
				TVLImpact = 2
			} else if data.EstimatedTotalAssets.Cmp(big.NewInt(50_000_000)) < 0 {
				TVLImpact = 3
			} else if data.EstimatedTotalAssets.Cmp(big.NewInt(100_000_000)) < 0 {
				TVLImpact = 4
			}

			// Longevity Impact
			var LongevityImpact float64 = 5
			if data.Activation != nil {
				var longevity big.Int
				now := big.NewInt(time.Now().Unix())
				longevity.Sub(now, data.Activation)

				var days int64 = 60 * 60 * 24
				if longevity.Cmp(big.NewInt(240*days)) > 0 {
					LongevityImpact = 1
				} else if longevity.Cmp(big.NewInt(120*days)) > 0 {
					LongevityImpact = 2
				} else if longevity.Cmp(big.NewInt(30*days)) > 0 {
					LongevityImpact = 3
				} else if longevity.Cmp(big.NewInt(7*days)) > 0 {
					LongevityImpact = 4
				}
			}

			// Store strategy to DB
			strategy := models.TStrategyFromRisk{
				RiskScores: models.TStrategyFromRiskRiskScores{
					TVLImpact:           TVLImpact,
					AuditScore:          stratGroup.AuditScore,
					CodeReviewScore:     stratGroup.CodeReviewScore,
					ComplexityScore:     stratGroup.ComplexityScore,
					LongevityImpact:     LongevityImpact,
					ProtocolSafetyScore: stratGroup.ProtocolSafetyScore,
					TeamKnowledgeScore:  stratGroup.TeamKnowledgeScore,
					TestingScore:        stratGroup.TestingScore,
				},
			}
			store.StrategiesFromRisk[chainID][addressHex] = strategy
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
