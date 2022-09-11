package prices

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/contracts"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
	"github.com/yearn/ydaemon/internal/utils/store"
)

// lensABI takes the ABI of the lens contract and prepare it for the multicall
var lensABI, _ = contracts.OracleMetaData.GetAbi()

// getPriceUsdcRecommendedCall will pack the transaction with it's argument and return the
// ethereum.Call object that can be used to execute the transaction.
func getPriceUsdcRecommendedCall(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := lensABI.Pack("getPriceUsdcRecommended", tokenAddress)
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      lensABI,
			Method:   `getPriceUsdcRecommended`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      lensABI,
		Method:   `getPriceUsdcRecommended`,
		CallData: parsedData,
		Name:     name,
	}
}

type TCurveFactoriesPoolData struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	LPAddress   string  `json:"lpTokenAddress"`
	TotalSupply string  `json:"totalSupply"`
	USDTotal    float64 `json:"usdTotal"`
}
type TCurveFactories struct {
	Data struct {
		PoolData []TCurveFactoriesPoolData `json:"poolData"`
	} `json:"data"`
}

func fetchCurve(url string) []TCurveFactoriesPoolData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
		return []TCurveFactoriesPoolData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []TCurveFactoriesPoolData{}
	}
	var factories TCurveFactories
	if err := json.Unmarshal(body, &factories); err != nil {
		return []TCurveFactoriesPoolData{}
	}
	return factories.Data.PoolData
}

// setCurveFactoriesPrices is used after setting the prices from the multicall, aka the lens contract.
// Some missing prices may be for Curve LP tokens or Curve tokens.
// The Curve API provides the total supply and the total USD value of the LP tokens. Based on that
// we can calculate the price per token, and assign it to the token if the Store price is 0
func setCurveFactoriesPrices(chainID uint64) {
	curveFactoryPoolData := []TCurveFactoriesPoolData{}

	// Running a sync group to execute all fetch at the same time
	wg := sync.WaitGroup{}
	for _, url := range helpers.CURVE_FACTORY_URI[chainID] {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			curveFactoryPoolData = append(curveFactoryPoolData, fetchCurve(url)...)
		}(url)
	}
	wg.Wait()

	// For each pool, we calculate the price per token and assign it to the token
	// if the Store price is 0
	for _, fact := range curveFactoryPoolData {
		pricePerToken := 0.0
		formatedTotalSupply, _ := helpers.FormatAmount(fact.TotalSupply, 18)
		if formatedTotalSupply == 0 || fact.USDTotal == 0 {
			continue
		}
		pricePerToken = fact.USDTotal / formatedTotalSupply
		pricePerTokenBigFloat := big.NewFloat(0).SetFloat64(pricePerToken)
		pricePerTokenBigFloat = pricePerTokenBigFloat.Mul(pricePerTokenBigFloat, big.NewFloat(1e6))
		pricePerTokenBigInt, _ := pricePerTokenBigFloat.Int(nil)
		addressToUse := fact.LPAddress
		if addressToUse == `` {
			addressToUse = fact.Address
		}
		if Store.TokenPrices[chainID][common.HexToAddress(addressToUse)] != nil {
			if Store.TokenPrices[chainID][common.HexToAddress(addressToUse)].Cmp(big.NewInt(0)) == 0 {
				Store.TokenPrices[chainID][common.HexToAddress(addressToUse)] = pricePerTokenBigInt
			}
		}
	}
}

// setMissingYearnVaultPrices is used after setting the prices to assign the price of yvToken when
// it's missing. It's done by taking the PricePerShare from each vault and multiplying it by the
// price of the underlying token. If the underlying token price is 0, we skip it.
func setMissingYearnVaultPrices(chainID uint64) {
	allVaultsData := tokens.Store.Tokens[chainID]

	for key, value := range Store.TokenPrices[chainID] {
		if value.Cmp(big.NewInt(0)) == 0 {
			if allVaultsData[key] != nil && allVaultsData[key].IsVault { // Is this address a vault?
				pricePerShare := Store.VaultPricePerShare[chainID][key]

				//pricePerShare is ^18, we need to convert to ^6
				pricePerShare = pricePerShare.Div(pricePerShare, big.NewInt(1e12))

				underlyingTokenAddress := allVaultsData[key].UnderlyingTokenAddress
				underlyingTokenPrice := Store.TokenPrices[chainID][underlyingTokenAddress]
				vaultValue := big.NewInt(0).Mul(pricePerShare, underlyingTokenPrice)
				if vaultValue.Cmp(big.NewInt(0)) == 0 {
					// logs.Info(chainID, key.String()+`: `+value.String())
					continue
				}
				Store.TokenPrices[chainID][key] = vaultValue
			}
		}
	}
}

// FetchLens will fetch the prices of the yvTokens and the tokens that are listed in the
// corresponding subgraph, execute a multicall on the given chainID to retreive them,
// and will store the prices in the TokenPrices map.
func FetchLens(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := ethereum.MulticallClientForChainID[chainID]

	// Then, for the given chainID, we need to select the correct lens contract address,
	// aka the endpoint we will use to perform the read transaction.
	lensAddress := ethereum.GetLensAddress(chainID)

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, token := range tokens.Store.TokenList[chainID] {
		calls = append(calls, getPriceUsdcRecommendedCall(token.String(), lensAddress, token))
	}

	if len(calls) == 0 {
		logs.Error("No tokens found.")
		return
	}

	// Then, we execute the multicall and store the prices in the TokenPrices map
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}
	response := caller.ExecuteByBatch(calls, maxBatch)
	if Store.TokenPrices[chainID] == nil {
		Store.TokenPrices[chainID] = make(map[common.Address]*big.Int)
	}
	for key, value := range response {
		address := strings.TrimSuffix(key, "getPriceUsdcRecommended")
		Store.TokenPrices[chainID][common.HexToAddress(address)] = value[0].(*big.Int)
	}

	if chainID == 1 {
		setCurveFactoriesPrices(chainID)
	}

	setMissingYearnVaultPrices(chainID)

	//TODO: debug, list prices 0
	// isZero := 0
	// for key, value := range Store.TokenPrices[chainID] {
	// 	if value.Cmp(big.NewInt(0)) == 0 {
	// 		logs.Info(chainID, key.String()+`: `+value.String())
	// 		isZero++
	// 		continue
	// 	}
	// }
	// logs.Pretty(isZero)

	store.SaveInDBForChainID(`TokenPrices`, chainID, Store.TokenPrices[chainID])
}

// LoadLens will reload the lens data store from the last state of the local Badger Database
func LoadLens(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]*big.Int)
	if err := store.LoadFromDBForChainID(`TokenPrices`, chainID, &temp); err != nil {
		return
	}
	if temp != nil && (len(temp) > 0) {
		Store.TokenPrices[chainID] = temp
		logs.Success("Data loaded for the lens data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
