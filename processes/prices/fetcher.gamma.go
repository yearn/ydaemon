package prices

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type TGammaData struct {
	PoolAddress string   `json:"poolAddress"`
	Name        string   `json:"name"`
	Token0      string   `json:"token0"`
	Token1      string   `json:"token1"`
	Decimals0   int      `json:"decimals0"`
	Decimals1   int      `json:"decimals1"`
	Tvl0        float64  `json:"tvl0"`
	Tvl1        float64  `json:"tvl1"`
	TvlUSD      string   `json:"tvlUSD"`
	PoolTvlUSD  string   `json:"poolTvlUSD"`
	PoolFeesUSD string   `json:"poolFeesUSD"`
	TotalSupply *big.Int `json:"totalSupply"`
}

func fetchGammaData() map[string]TGammaData {
	resp, err := http.Get(`https://wire2.gamma.xyz/quickswap/polygon/hypervisors/allData`)
	if err != nil {
		logs.Error(`impossible to get Gamma URL`, err.Error())
		return map[string]TGammaData{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(`impossible to read Gamma Get body`, err.Error())
		return map[string]TGammaData{}
	}
	var factories map[string]TGammaData
	if err := json.Unmarshal(body, &factories); err != nil {
		logs.Error(`impossible to unmarshal Gamma Get body`, err.Error())
		return map[string]TGammaData{}
	}
	return factories
}

// getGammaLPPricesFromAPI
func getGammaLPPricesFromAPI(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)

	if chainID != 137 {
		return priceMap
	}

	pools := fetchGammaData()
	for addr, poolData := range pools {
		totalTVLFloat := bigNumber.NewFloat(0).SetString(poolData.PoolTvlUSD)
		totalSupplyFloat := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(poolData.TotalSupply), 18)
		tokenPrice := bigNumber.NewFloat(0).Div(totalTVLFloat, totalSupplyFloat)
		rawPrice := bigNumber.NewFloat(0).Mul(tokenPrice, bigNumber.NewFloat(1e6)).Int()
		priceMap[common.HexToAddress(addr)] = models.TPrices{
			Address:        common.HexToAddress(addr),
			Price:          rawPrice,
			HumanizedPrice: tokenPrice,
			Source:         `gammaAPI`,
		}
	}
	return priceMap
}
