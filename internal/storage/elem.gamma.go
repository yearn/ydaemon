package storage

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** This struct represents the relevant response from the Gamma Merkl API & the Gamma allData APR
**************************************************************************************************/
type TGammaMerklAPIResp struct {
	Pools map[string]struct {
		ALM map[string]struct {
			ALMAPR     float64 `json:"almAPR"`
			ALMAddress string  `json:"almAddress"`
		} `json:"alm"`
	} `json:"pools"`
}

type TGammaDataAPIResp struct {
	Address     string   `json:"address"`
	PoolAddress string   `json:"poolAddress"`
	Token0      string   `json:"token0"`
	Token1      string   `json:"token1"`
	Name        string   `json:"name"`
	Decimals0   int      `json:"decimals0"`
	Decimals1   int      `json:"decimals1"`
	Tvl0        float64  `json:"tvl0"`
	Tvl1        float64  `json:"tvl1"`
	TvlUSD      string   `json:"tvlUSD"`
	PoolTvlUSD  string   `json:"poolTvlUSD"`
	PoolFeesUSD string   `json:"poolFeesUSD"`
	TotalSupply *big.Int `json:"totalSupply"`
	Returns     struct {
		Monthly struct {
			APR float64 `json:"feeApr"`
		} `json:"monthly"`
	} `json:"returns"`
}

/**************************************************************************************************
** As the API response is quite large, we cache it to avoid unnecessary requests.
** This will be refreshed everytime we globaly refresh the APR system
**************************************************************************************************/
var cachedGammaMerkl map[uint64]map[string]TGammaMerklAPIResp
var cachedGammaAllData map[uint64]map[string]TGammaDataAPIResp

func init() {
	cachedGammaMerkl = make(map[uint64]map[string]TGammaMerklAPIResp)
	cachedGammaAllData = make(map[uint64]map[string]TGammaDataAPIResp)
}
func RefreshGammaCalls(chainID uint64) {
	updatedMerklValues, ok := RetrieveGammaMerklData(chainID)
	if !ok {
		return
	}
	cachedGammaMerkl[chainID] = updatedMerklValues

	updatedAllDataValues, ok := RetrieveGammaAllData(chainID)
	if !ok {
		return
	}
	cachedGammaAllData[chainID] = updatedAllDataValues
}

func GetCachedGammaMerkl(chainID uint64) (map[string]TGammaMerklAPIResp, bool) {
	value, ok := cachedGammaMerkl[chainID]
	return value, ok
}

func GetCachedGammaAllData(chainID uint64) (map[string]TGammaDataAPIResp, bool) {
	value, ok := cachedGammaAllData[chainID]
	return value, ok
}

/**************************************************************************
** This function retrieves the Gamma Merkl API response.
**************************************************************************/
func RetrieveGammaMerklData(chainID uint64) (map[string]TGammaMerklAPIResp, bool) {
	if _, ok := cachedGammaMerkl[chainID]; !ok {
		cachedGammaMerkl[chainID] = map[string]TGammaMerklAPIResp{}
	}

	pools := map[string]TGammaMerklAPIResp{}
	if env.CHAINS[chainID].ExtraURI.GammaMerklURI == `` {
		return pools, false
	}
	resp, err := http.Get(env.CHAINS[chainID].ExtraURI.GammaMerklURI)
	if err != nil {
		logs.Error(err)
		return pools, false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return pools, false
	}
	var gammaMerkl map[string]TGammaMerklAPIResp
	if err := json.Unmarshal(body, &gammaMerkl); err != nil {
		logs.Error(err)
		return pools, false
	}
	cachedGammaMerkl[chainID] = gammaMerkl
	return gammaMerkl, true
}

/**************************************************************************
** This function retrieves the Gamma allData API response
**************************************************************************/
func RetrieveGammaAllData(chainID uint64) (map[string]TGammaDataAPIResp, bool) {
	if _, ok := cachedGammaAllData[chainID]; !ok {
		cachedGammaAllData[chainID] = map[string]TGammaDataAPIResp{}
	}

	pools := map[string]TGammaDataAPIResp{}
	if env.CHAINS[chainID].ExtraURI.GammaHypervisorURI == nil ||
		len(env.CHAINS[chainID].ExtraURI.GammaHypervisorURI) == 0 {
		return pools, false
	}

	var gammaAllData map[string]TGammaDataAPIResp
	for _, uri := range env.CHAINS[chainID].ExtraURI.GammaHypervisorURI {
		resp, err := http.Get(uri)
		if err != nil {
			logs.Error(err)
			return pools, false
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
			return pools, false
		}
		var gammaAllDataTmp map[string]TGammaDataAPIResp
		if err := json.Unmarshal(body, &gammaAllDataTmp); err != nil {
			logs.Error(err)
			return pools, false
		}

		gammaAllData = make(map[string]TGammaDataAPIResp)
		for address, pool := range gammaAllDataTmp {
			pool.Address = common.HexToAddress(address).Hex()
			gammaAllData[common.HexToAddress(address).Hex()] = pool
		}
		cachedGammaAllData[chainID] = gammaAllData
	}

	return gammaAllData, true
}
