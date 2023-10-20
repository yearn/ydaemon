package apr

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type TFraxMinMax struct {
	Min string `json:"min"`
	Max string `json:"max"`
}
type TFraxCoin struct {
	Address                string  `json:"address"`
	UsdPrice               float64 `json:"usdPrice"`
	Decimals               uint64  `json:"decimals"`
	Name                   string  `json:"name"`
	Symbol                 string  `json:"symbol"`
	StakingContractBalance string  `json:"stakingContractBalance"`

	//Added after
	RewardApr           float64 `json:"rewardApr"`
	MinBoostedRewardApr float64 `json:"minBoostedRewardApr"`
	MaxBoostedRewardApr float64 `json:"maxBoostedRewardApr"`
}
type TFraxPool struct {
	ID                          uint64  `json:"id"`
	Type                        string  `json:"type"`
	Name                        string  `json:"name"`
	UnderlyingTokenAddress      string  `json:"underlyingTokenAddress"`
	StakingAddress              string  `json:"stakingAddress"`
	StakingTokenAddress         string  `json:"stakingTokenAddress"`
	RewardsAddress              string  `json:"rewardsAddress"`
	UnderlyingTokenVirtualPrice float64 `json:"underlyingTokenVirtualPrice"`
	StakingTokenUsdPrice        float64
	StakingTokenUsdPriceRaw     any           `json:"stakingTokenUsdPrice"`
	StakingTokenSymbol          string        `json:"stakingTokenSymbol"`
	StakingTokenName            string        `json:"stakingTokenName"`
	StakingTokenDecimals        uint64        `json:"stakingTokenDecimals"`
	VeFXSMultiplier             uint64        `json:"veFXSMultiplier"`
	LockMaxMultiplier           float64       `json:"lockMaxMultiplier"`
	RewardCoins                 []TFraxCoin   `json:"rewardCoins"`
	RewardAPRs                  []string      `json:"rewardAprs"`
	BoostedRewardAprs           []TFraxMinMax `json:"boostedRewardAprs"`
	TotalRewardAPRs             TFraxMinMax   `json:"totalRewardAprs"`
}
type TFraxPools struct {
	Pools struct {
		AugmentedPoolData []TFraxPool `json:"augmentedPoolData"`
	} `json:"pools"`
}

const FRAX_POOL_API_URI = "https://frax.convexfinance.com/api/frax/pools"

func retrieveFraxPools(chainID uint64) []TFraxPool {
	resp, err := http.Get(FRAX_POOL_API_URI)
	if err != nil {
		logs.Error(err)
		return []TFraxPool{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return []TFraxPool{}
	}
	var pools TFraxPools
	if err := json.Unmarshal(body, &pools); err != nil {
		logs.Error(err)
		return []TFraxPool{}
	}
	poolsList := []TFraxPool{}
	for _, pool := range pools.Pools.AugmentedPoolData {
		if pool.Type != `convex` {
			continue
		}
		switch pool.StakingTokenUsdPriceRaw.(type) {
		case string:
			pool.StakingTokenUsdPrice, _ = strconv.ParseFloat(pool.StakingTokenUsdPriceRaw.(string), 64)
		case float64:
			pool.StakingTokenUsdPrice = pool.StakingTokenUsdPriceRaw.(float64)
		}

		for index := range pool.RewardCoins {
			pool.RewardCoins[index].RewardApr, _ = strconv.ParseFloat(pool.RewardAPRs[index], 64)
			pool.RewardCoins[index].MinBoostedRewardApr, _ = strconv.ParseFloat(pool.BoostedRewardAprs[index].Min, 64)
			pool.RewardCoins[index].MaxBoostedRewardApr, _ = strconv.ParseFloat(pool.BoostedRewardAprs[index].Max, 64)
		}
		poolsList = append(poolsList, pool)
	}
	return poolsList
}

func findFraxPoolForVault(tokenAddress common.Address, pools []TFraxPool) TFraxPool {
	for _, pool := range pools {
		if common.HexToAddress(pool.UnderlyingTokenAddress) == tokenAddress {
			return pool
		}
	}
	return TFraxPool{}
}

/**************************************************************************************************
** Check if the strategy is a frax strategy. This is a check based on the strategy name. What
** could go wrong.
**************************************************************************************************/
func isFraxStrategy(strategy models.TStrategy) bool {
	return strings.Contains(strings.ToLower(strategy.Name), `convexfrax`)
}
