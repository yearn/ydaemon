package prices

import (
	"encoding/json"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

var VELO_SUGAR_ADDRESS = common.HexToAddress(`0x3e532BC1998584fe18e357B5187897ad0110ED3A`)
var VELO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0xcA97e5653d775cA689BED5D0B4164b7656677011`)
var OPT_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var OPT_OP_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000042`)
var OPT_USDC_ADDRESS = common.HexToAddress(`0x7F5c764cBc14f9669B88837ca1490cCa17c31607`)
var OPT_RATE_CONNECTORS = []common.Address{
	common.HexToAddress(`0x9560e827aF36c94D2Ac33a39bCE1Fe78631088Db`),
	common.HexToAddress(`0x4200000000000000000000000000000000000042`),
	common.HexToAddress(`0x4200000000000000000000000000000000000006`),
	common.HexToAddress(`0x9bcef72be871e61ed4fbbc7630889bee758eb81d`),
	common.HexToAddress(`0x2e3d870790dc77a83dd1d18184acc7439a53f475`),
	common.HexToAddress(`0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9`),
	common.HexToAddress(`0x1f32b1c2345538c0c6f582fcb022739c4a194ebb`),
	common.HexToAddress(`0xbfd291da8a403daaf7e5e9dc1ec0aceacd4848b9`),
	common.HexToAddress(`0xc3864f98f2a61a7caeb95b039d031b4e2f55e0e9`),
	common.HexToAddress(`0x9485aca5bbbe1667ad97c7fe7c4531a624c8b1ed`),
	common.HexToAddress(`0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1`),
	common.HexToAddress(`0x73cb180bf0521828d8849bc8cf2b920918e23032`),
	common.HexToAddress(`0x6806411765af15bddd26f8f544a34cc40cb9838b`),
	common.HexToAddress(`0x6c2f7b6110a37b3b0fbdd811876be368df02e8b0`),
	common.HexToAddress(`0xc5b001dc33727f8f26880b184090d3e252470d45`),
	common.HexToAddress(`0x6c84a8f1c29108f47a79964b5fe888d4f4d0de40`),
	common.HexToAddress(`0xc40f949f8a4e094d1b49a23ea9241d289b7b2819`),
	common.HexToAddress(`0x94b008aa00579c1307b0ef2c499ad98a8ce58e58`),
	common.HexToAddress(`0x0b2c639c533813f4aa9d7837caf62653d097ff85`),
	common.HexToAddress(`0x7f5c764cbc14f9669b88837ca1490cca17c31607`),
}

func fetchVelo(url string) []TVeloPairData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(`impossible to get velo URL`, err.Error())
		return []TVeloPairData{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(`impossible to read velo Get body`, err.Error())
		return []TVeloPairData{}
	}
	var factories TVeloPairs
	if err := json.Unmarshal(body, &factories); err != nil {
		logs.Error(`impossible to unmarshal velo Get body`, err.Error())
		return []TVeloPairData{}
	}
	return factories.Data
}

// fetchPricesFromSugar is used to fetch prices from the sugar APR (velo).
func fetchPricesFromSugar(chainID uint64, blockNumber *uint64, tokens []models.TERC20Token) map[common.Address]models.TPrices {
	priceMap := make(map[common.Address]models.TPrices)
	if chainID != 10 {
		return priceMap
	}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	sugar, _ := contracts.NewVeloSugarCaller(VELO_SUGAR_ADDRESS, client)
	for i := 0; i < 30; i++ {
		start := int64(i * 25)
		allSugar, err := sugar.All(nil, big.NewInt(50), big.NewInt(start))
		if len(allSugar) == 0 || err != nil {
			if err != nil {
				logs.Error(`error fetching velo sugar`, err)
			}
			break
		}

		/**********************************************************************************************
		** -> Retrieve the ratesConnector
		** We need to know the prices of the tokens in the pairs. We will use the velo sugar oracle
		** to get the prices, via the `getManyRatesWithConnectors` method. It will returns an array
		** of prices.
		** Then, we need to prepare the prices to be used in the next step. We will create a map of
		** prices with the token address as key and the price as value.
		**********************************************************************************************/
		tokensInVelo := []common.Address{}
		for _, pair := range allSugar {
			if !helpers.Contains(tokensInVelo, pair.Token0) {
				tokensInVelo = append(tokensInVelo, pair.Token0)
			}
			if !helpers.Contains(tokensInVelo, pair.Token1) {
				tokensInVelo = append(tokensInVelo, pair.Token1)
			}
		}
		veloSugarContract, err := contracts.NewVeloSugarOracleCaller(VELO_SUGAR_ORACLE_ADDRESS, client)
		if err != nil {
			logs.Error(`error fetching velo sugar contract`, err)
			return priceMap
		}
		prices, err := veloSugarContract.GetManyRatesWithConnectors(
			nil,
			uint8(len(tokensInVelo)),
			append(tokensInVelo, OPT_RATE_CONNECTORS...),
		)
		if err != nil {
			logs.Error(`error fetching velo sugar prices`, err)
			return priceMap
		}
		priceObj := make(map[common.Address]*bigNumber.Int)
		for i, rate := range tokensInVelo {
			priceObj[rate] = bigNumber.SetInt(prices[i])
		}

		/**********************************************************************************************
		** -> Retrieve tokens decimals
		**
		**********************************************************************************************/
		calls := []ethereum.Call{}
		for _, token := range tokensInVelo {
			if pairToken, ok := storage.GetERC20(chainID, token); !ok || pairToken.Decimals == 0 {
				calls = append(calls, multicalls.GetDecimals(token.Hex(), token))
			}
		}
		decimalsForTokens := multicalls.Perform(chainID, calls, nil)

		/**********************************************************************************************
		** We can now loop over the pairs and calculate the price of the pair and tokens
		**********************************************************************************************/
		for _, pair := range allSugar {
			//Assign token0 decimals
			token0Decimals := uint64(0)
			if pairToken, ok := storage.GetERC20(chainID, pair.Token0); ok {
				token0Decimals = pairToken.Decimals
			} else if decimals, ok := decimalsForTokens[pair.Token0.Hex()+`decimals`]; ok {
				token0Decimals = helpers.DecodeUint64(decimals)
			}

			//Assign token1 decimals
			token1Decimals := uint64(0)
			if pairToken, ok := storage.GetERC20(chainID, pair.Token1); ok {
				token1Decimals = pairToken.Decimals
			} else if decimals, ok := decimalsForTokens[pair.Token1.Hex()+`decimals`]; ok {
				token1Decimals = helpers.DecodeUint64(decimals)
			}

			//Assign token0 price
			token0Price := priceObj[pair.Token0]
			if token0Price.IsZero() && addresses.Equals(pair.Token0, OPT_USDC_ADDRESS) {
				token0Price = bigNumber.NewInt(1e18)
			}

			//Assign token1 price
			token1Price := priceObj[pair.Token1]
			if token1Price.IsZero() && addresses.Equals(pair.Token1, OPT_USDC_ADDRESS) {
				token1Price = bigNumber.NewInt(1e18)
			}

			token0PriceUSD := helpers.ToNormalizedAmount(token0Price, 18)
			token1PriceUSD := helpers.ToNormalizedAmount(token1Price, 18)
			token0totalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Reserve0), uint64(token0Decimals))
			token1totalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Reserve1), uint64(token1Decimals))
			pairTotalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Liquidity), 18)

			token0ValueInPair := bigNumber.NewFloat(0).Mul(token0PriceUSD, token0totalSupply)
			token1ValueInPair := bigNumber.NewFloat(0).Mul(token1PriceUSD, token1totalSupply)
			valueInPair := bigNumber.NewFloat(0).Add(token0ValueInPair, token1ValueInPair)

			pairPrice := bigNumber.NewFloat(0).Div(valueInPair, pairTotalSupply)
			pairPrice = bigNumber.NewFloat(0).Mul(pairPrice, bigNumber.NewFloat(1e6))
			token0Price = bigNumber.NewFloat(0).Mul(token0PriceUSD, bigNumber.NewFloat(1e6)).Int()
			token1Price = bigNumber.NewFloat(0).Mul(token1PriceUSD, bigNumber.NewFloat(1e6)).Int()

			if pairToken, _ := storage.GetERC20(chainID, pair.Lp); !pairToken.IsVaultLike() {
				if !pairPrice.IsZero() {
					humanizedPrice := helpers.ToNormalizedAmount(pairPrice.Int(), 6)
					priceMap[pair.Lp] = models.TPrices{
						Address:        pair.Lp,
						Price:          pairPrice.Int(),
						HumanizedPrice: humanizedPrice,
						Source:         `veloSugarPair`,
					}
				}
			}

			if token0, _ := storage.GetERC20(chainID, pair.Token0); !token0.IsVaultLike() {
				if !token0Price.IsZero() {
					humanizedPrice := helpers.ToNormalizedAmount(token0Price, 6)
					newPriceItem := models.TPrices{
						Address:        pair.Token0,
						Price:          token0Price,
						HumanizedPrice: humanizedPrice,
						Source:         `veloSugarToken0`,
					}

					if existing, ok := priceMap[pair.Token1]; ok {
						if existing.Source != `veloSugarPair` {
							priceMap[pair.Token0] = newPriceItem
						}
					} else {
						priceMap[pair.Token0] = newPriceItem
					}
				}
			}

			if token1, _ := storage.GetERC20(chainID, pair.Token1); !token1.IsVaultLike() {
				if !token1Price.IsZero() {
					humanizedPrice := helpers.ToNormalizedAmount(token1Price, 6)
					newPriceItem := models.TPrices{
						Address:        pair.Token1,
						Price:          token1Price,
						HumanizedPrice: humanizedPrice,
						Source:         `veloSugarToken1`,
					}

					if existing, ok := priceMap[pair.Token1]; ok {
						if existing.Source != `veloSugarPair` {
							priceMap[pair.Token0] = newPriceItem
						}
					} else {
						priceMap[pair.Token0] = newPriceItem
					}
				}
			}
		}
	}

	return priceMap
}
