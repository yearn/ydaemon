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
	"unicode/utf8"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/external/tokens"
)

// lensABI takes the ABI of the lens contract and prepare it for the multicall
var lensABI, _ = contracts.OracleMetaData.GetAbi()

// getPriceUsdcRecommendedCall will pack the transaction with it's argument and return the
// ethereum.Call object that can be used to execute the transaction.
func getPriceUsdcRecommendedCall(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := lensABI.Pack("getPriceUsdcRecommended", tokenAddress.Address)
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress.Address,
			Abi:      lensABI,
			Method:   `getPriceUsdcRecommended`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      lensABI,
		Method:   `getPriceUsdcRecommended`,
		CallData: parsedData,
		Name:     name,
	}
}

type TCurveFactoriesPoolData struct {
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Address     string  `json:"address"`
	LPAddress   string  `json:"lpTokenAddress"`
	TotalSupply string  `json:"totalSupply"`
	USDTotal    float64 `json:"usdTotal"`
	Coins       []struct {
		Address  string          `json:"address"`
		Decimals json.RawMessage `json:"decimals"`
		Symbol   string          `json:"symbol"`
		USDPrice float64         `json:"usdPrice"`
	} `json:"coins"`
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
		logs.Error(err)
		return []TCurveFactoriesPoolData{}
	}
	var factories TCurveFactories
	if err := json.Unmarshal(body, &factories); err != nil {
		logs.Error(err)
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
	for _, url := range env.CURVE_FACTORY_URI[chainID] {
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
		//First add the underlying tokens
		for _, coin := range fact.Coins {
			coinAddress := common.HexToAddress(coin.Address)
			if Store.TokenPrices[chainID][coinAddress] == nil {
				Store.TokenPrices[chainID][coinAddress] = bigNumber.NewInt()
			}

			if Store.TokenPrices[chainID][coinAddress].Cmp(big.NewInt(0)) == 0 {

				coinPrice := bigNumber.NewFloat(coin.USDPrice)
				coinPrice = coinPrice.Mul(coinPrice, bigNumber.NewFloat(1e6))
				coinPriceBigInt := coinPrice.Int()
				Store.TokenPrices[chainID][coinAddress] = coinPriceBigInt

				//Store the price in the token struct
				if tokens.Store.Tokens[chainID][coinAddress] == nil {
					//Hack because decimals can be string or uint
					decimals := uint64(18)
					if utf8.Valid(coin.Decimals) {
						if i, err := strconv.Atoi(string(coin.Decimals)); err != nil {
							decimals = uint64(i)
						} else {
							decimals, _ = strconv.ParseUint(string(coin.Decimals), 10, 64)
						}
					}
					tokens.Store.Tokens[chainID][coinAddress] = &tokens.TERC20Token{
						Address:  common.HexToAddress(coin.Address),
						Name:     coin.Symbol,
						Symbol:   coin.Symbol,
						Decimals: decimals,
						IsVault:  false,
					}
				}

				humanizedPrice, _ := helpers.FormatAmount(coinPriceBigInt.String(), 6)
				tokens.Store.Tokens[chainID][common.HexToAddress(coin.Address)].Price = humanizedPrice
			}
		}

		pricePerToken := 0.0
		formatedTotalSupply, _ := helpers.FormatAmount(fact.TotalSupply, 18)
		if formatedTotalSupply > 0 && fact.USDTotal > 0 {
			pricePerToken = fact.USDTotal / formatedTotalSupply
			pricePerTokenBigFloat := bigNumber.NewFloat(pricePerToken)
			pricePerTokenBigFloat = pricePerTokenBigFloat.Mul(pricePerTokenBigFloat, bigNumber.NewFloat(1e6))
			pricePerTokenBigInt := pricePerTokenBigFloat.Int()
			addressToUse := fact.LPAddress
			if addressToUse == `` {
				addressToUse = fact.Address
			}
			if Store.TokenPrices[chainID][common.HexToAddress(addressToUse)] == nil {
				Store.TokenPrices[chainID][common.HexToAddress(addressToUse)] = bigNumber.NewInt(0)
			}
			if Store.TokenPrices[chainID][common.HexToAddress(addressToUse)].IsZero() {
				Store.TokenPrices[chainID][common.HexToAddress(addressToUse)] = pricePerTokenBigInt

				if tokens.Store.Tokens[chainID][common.HexToAddress(addressToUse)] == nil {
					tokens.Store.Tokens[chainID][common.HexToAddress(addressToUse)] = &tokens.TERC20Token{
						Address:  common.HexToAddress(addressToUse),
						Name:     fact.Name,
						Symbol:   fact.Symbol,
						Decimals: 18,
						IsVault:  false,
					}
				}

				//Store the price in the token struct
				humanizedPrice, _ := helpers.FormatAmount(pricePerTokenBigInt.String(), 6)
				tokens.Store.Tokens[chainID][common.HexToAddress(addressToUse)].Price = humanizedPrice
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
		if !helpers.AddressIsValid(key, chainID) {
			continue
		}
		if value.Cmp(big.NewInt(0)) > 0 {
			continue
		}

		if allVaultsData[key] != nil && allVaultsData[key].IsVault { // Is this address a vault?
			// pricePerShare := bigNumber.NewInt().Clone(&vaults.Store.AggregatedVault[chainID][key].PricePerShare)
			pricePerShare := bigNumber.NewInt()

			//PricePerShare is in 10^decimals, we need to convert it to 10^6
			decimals := allVaultsData[key].Decimals
			if decimals > 0 {
				pricePerShare = pricePerShare.Mul(pricePerShare, bigNumber.NewInt().Exp(
					bigNumber.NewInt(10), bigNumber.NewInt(6), nil,
				))
				pricePerShare = pricePerShare.Div(pricePerShare, bigNumber.NewInt().Exp(
					bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil,
				))
			}

			underlyingTokenAddress := allVaultsData[key].UnderlyingTokenAddress
			underlyingTokenPrice, ok := Store.TokenPrices[chainID][underlyingTokenAddress]

			if ok && !underlyingTokenPrice.IsZero() {
				underlyingTokenPrice = bigNumber.NewInt().Clone(underlyingTokenPrice)
				vaultValue := bigNumber.NewInt().Mul(pricePerShare, underlyingTokenPrice)
				vaultValue = vaultValue.Div(vaultValue, bigNumber.NewInt(1e6))

				if vaultValue.IsZero() {
					// logs.Info(chainID, key.String()+`: `+value.String())
					continue
				}

				Store.TokenPrices[chainID][key] = vaultValue
				humanizedPrice, _ := helpers.FormatAmount(vaultValue.String(), 6)
				tokens.Store.Tokens[chainID][key].Price = humanizedPrice
			}
		}
	}
}

// FetchLens will fetch the prices of the yvTokens and the tokens that are listed in the
// corresponding subgraph, execute a multicall on the given chainID to retreive them,
// and will store the prices in the TokenPrices map.
func FetchLens(chainID uint64) {
	caller := ethereum.MulticallClientForChainID[chainID]
	lensAddress, ok := env.LENS_ADDRESSES[chainID]
	if !ok {
		logs.Error(`Lens address not found for chainID: `, chainID)
		return
	}

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
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	if Store.TokenPrices[chainID] == nil {
		Store.TokenPrices[chainID] = make(map[common.Address]*bigNumber.Int)
	}
	for key, value := range response {
		address := strings.TrimSuffix(key, "getPriceUsdcRecommended")
		if !helpers.AddressIsValid(common.HexToAddress(address), chainID) {
			continue
		}

		price := bigNumber.SetInt(value[0].(*big.Int))
		Store.TokenPrices[chainID][common.HexToAddress(address)] = price
		go store.SaveInDB(chainID, store.TABLES.PRICES, address, price)

		//Store the price in the token struct
		humanizedPrice, _ := helpers.FormatAmount(price.String(), 6)
		tokens.Store.Tokens[chainID][common.HexToAddress(address)].Price = humanizedPrice
	}

	setCurveFactoriesPrices(chainID)
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
}

// LoadLens will reload the lens data store from the last state of the local Badger Database
func LoadLens(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := make(map[common.Address]*bigNumber.Int)
	store.Interate(chainID, store.TABLES.PRICES, &temp)
	if temp != nil && (len(temp) > 0) {
		Store.TokenPrices[chainID] = temp
		logs.Success("Data loaded for the lens data store for chainID: " + strconv.FormatUint(chainID, 10))
	}
}
