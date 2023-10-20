package apr

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

func retrieveCurveGauges(chainID uint64) []models.CurveGauge {
	resp, err := http.Get(CURVE_GAUGES_URI[chainID])
	if err != nil {
		logs.Error(err)
		return []models.CurveGauge{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return []models.CurveGauge{}
	}
	var gauges models.TCurveGauges
	if err := json.Unmarshal(body, &gauges); err != nil {
		logs.Error(err)
		return []models.CurveGauge{}
	}
	pools := []models.CurveGauge{}
	for _, gauge := range gauges.Data {
		pools = append(pools, gauge)
	}
	return pools
}

func retrieveCurveGetPools(chainID uint64) []models.CurvePool {
	URIsToFetch := env.CHAINS[chainID].Curve.PoolsURIs
	pools := []models.CurvePool{}
	for _, uri := range URIsToFetch {
		resp, err := http.Get(uri)
		if err != nil {
			logs.Error(err)
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
			continue
		}
		var getPools models.TCurvePools
		if err := json.Unmarshal(body, &getPools); err != nil {
			logs.Error(err)
			continue
		}
		pools = append(pools, getPools.Data.PoolData...)
	}
	return pools
}

func retrieveCurveSubgraphData(chainID uint64) []models.CurveSubgraphData {
	if v, ok := CURVE_SUBGRAPHDATA_URI[chainID]; !ok || v == `` {
		return []models.CurveSubgraphData{}
	}
	resp, err := http.Get(CURVE_SUBGRAPHDATA_URI[chainID])
	if err != nil {
		logs.Error(err)
		return []models.CurveSubgraphData{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return []models.CurveSubgraphData{}
	}
	var subgraphData models.TCurveSubgraphData
	if err := json.Unmarshal(body, &subgraphData); err != nil {
		logs.Error(err)
		return []models.CurveSubgraphData{}
	}
	data := []models.CurveSubgraphData{}
	data = append(data, subgraphData.Data.PoolList...)
	return data
}

func findGaugeForVault(tokenAddress common.Address, pools []models.CurveGauge) models.CurveGauge {
	for _, pool := range pools {
		if common.HexToAddress(pool.SwapToken) == tokenAddress {
			return pool
		}
	}
	return models.CurveGauge{}
}

func findPoolForVault(tokenAddress common.Address, pools []models.CurvePool) models.CurvePool {
	for _, pool := range pools {
		if common.HexToAddress(pool.LPTokenAddress) == tokenAddress {
			return pool
		}
	}
	return models.CurvePool{}
}
func findSubgraphItemForVault(poolAddress common.Address, pools []models.CurveSubgraphData) models.CurveSubgraphData {
	for _, pool := range pools {
		if common.HexToAddress(pool.Address) == poolAddress {
			return pool
		}
	}
	return models.CurveSubgraphData{}
}
