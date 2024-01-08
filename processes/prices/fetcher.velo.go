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

var VELO_SUGAR_ADDRESS = common.HexToAddress(`0x4D996E294B00cE8287C16A2b9A4e637ecA5c939f`)
var VELO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0x395942c2049604a314d39f370dfb8d87aac89e16`)
var OPT_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var OPT_OP_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000042`)
var OPT_USDC_ADDRESS = common.HexToAddress(`0x7F5c764cBc14f9669B88837ca1490cCa17c31607`)
var OPT_RATE_CONNECTORS = []common.Address{
	common.HexToAddress(`0x1e925de1c68ef83bd98ee3e130ef14a50309c01b`),
	common.HexToAddress(`0xdc6ff44d5d932cbd77b52e5612ba0529dc6226f1`),
	common.HexToAddress(`0x747e42eb0591547a0ab429b3627816208c734ea7`),
	common.HexToAddress(`0x1f514a61bcde34f94bc39731235690ab9da737f7`),
	common.HexToAddress(`0xc03b43d492d904406db2d7d57e67c7e8234ba752`),
	common.HexToAddress(`0xeb466342c4d449bc9f53a865d5cb90586f405215`),
	common.HexToAddress(`0x894134a25a5fac1c2c26f1d8fbf05111a3cb9487`),
	common.HexToAddress(`0xbc7b1ff1c6989f006a1185318ed4e7b5796e66e1`),
	common.HexToAddress(`0xb396b31599333739a97951b74652c117be86ee1d`),
	common.HexToAddress(`0x01b8b6384298d4848e3be63d4c9d17830eee488a`),
	common.HexToAddress(`0xff733b2a3557a7ed6697007ab5d11b79fdd1b76b`),
	common.HexToAddress(`0x00e1724885473b63bce08a9f0a52f35b0979e35a`),
	common.HexToAddress(`0x3ee6107d9c93955acbb3f39871d32b02f82b78ab`),
	common.HexToAddress(`0x6af3cb766d6cd37449bfd321d961a61b0515c1bc`),
	common.HexToAddress(`0x33bca143d9b41322479e8d26072a00a352404721`),
	common.HexToAddress(`0x9dabae7274d28a45f0b65bf8ed201a5731492ca0`),
	common.HexToAddress(`0x0b2c639c533813f4aa9d7837caf62653d097ff85`),
	common.HexToAddress(`0x58b9cb810a68a7f3e1e4f8cb45d1b9b3c79705e8`),
	common.HexToAddress(`0x3b08fcd15280e7b5a6e404c4abb87f7c774d1b2e`),
	common.HexToAddress(`0x7d14206c937e70e19e3a5b94011faf0d5b3928e2`),
	common.HexToAddress(`0x259c1c2ed264402b5ed2f02bc7dc25a15c680c18`),
	common.HexToAddress(`0x1574564fcfd15bccb3fe04e9818f61131ea74066`),
	common.HexToAddress(`0xc55e93c62874d8100dbd2dfe307edc1036ad5434`),
	common.HexToAddress(`0x3b6564b5da73a41d3a66e6558a98fd0e9e1e77ad`),
	common.HexToAddress(`0x323665443cef804a3b5206103304bd4872ea4253`),
	common.HexToAddress(`0x3eb398fec5f7327c6b15099a9681d9568ded2e82`),
	common.HexToAddress(`0x348fdfe2c35934a96c1353185f09d0f9efbada86`),
	common.HexToAddress(`0x4f604735c1cf31399c6e711d5962b2b3e0225ad3`),
	common.HexToAddress(`0x2dd1b4d4548accea497050619965f91f78b3b532`),
	common.HexToAddress(`0x819845b60a192167ed1139040b4f8eca31834f27`),
	common.HexToAddress(`0x9560e827af36c94d2ac33a39bce1fe78631088db`),
	common.HexToAddress(`0x4200000000000000000000000000000000000042`),
	common.HexToAddress(`0x4200000000000000000000000000000000000006`),
	common.HexToAddress(`0x9bcef72be871e61ed4fbbc7630889bee758eb81d`),
	common.HexToAddress(`0x2e3d870790dc77a83dd1d18184acc7439a53f475`),
	common.HexToAddress(`0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9`),
	common.HexToAddress(`0x1f32b1c2345538c0c6f582fcb022739c4a194ebb`),
	common.HexToAddress(`0xbfd291da8a403daaf7e5e9dc1ec0aceacd4848b9`),
	common.HexToAddress(`0xc3864f98f2a61a7caeb95b039d031b4e2f55e0e9`),
	common.HexToAddress(`0x9485aca5bbbe1667ad97c7fe7c4531a624c8b1ed`),
	common.HexToAddress(`0xda10009cbd5d07dd0cecc66161fc93d7c9000da1`),
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
	// sugarOracle, _ := contracts.NewVeloSugarOracleCaller(VELO_SUGAR_ORACLE_ADDRESS, client)
	for i := 0; i < 4; i++ {
		start := int64(i * 1000)
		allSugar, err := sugar.All(nil, big.NewInt(1000), big.NewInt(start), common.Address{})
		if len(allSugar) == 0 || err != nil {
			break
		}

		for _, pair := range allSugar {
			calls := []ethereum.Call{}
			ratesConnector := []common.Address{pair.Token0, pair.Token1}
			ratesConnector = append(ratesConnector, OPT_RATE_CONNECTORS...)
			calls = append(calls, multicalls.GetManyRatesWithConnectors(VELO_SUGAR_ORACLE_ADDRESS.Hex(), VELO_SUGAR_ORACLE_ADDRESS, 2, ratesConnector))
			calls = append(calls, multicalls.GetDecimals(pair.Token0.Hex(), pair.Token0))
			calls = append(calls, multicalls.GetDecimals(pair.Token1.Hex(), pair.Token1))
			response := multicalls.Perform(chainID, calls, nil)
			prices := helpers.DecodeBigInts(response[VELO_SUGAR_ORACLE_ADDRESS.Hex()+`getManyRatesWithConnectors`])

			if len(prices) != 2 {
				continue
			}
			token0Decimals := helpers.DecodeUint64(response[pair.Token0.Hex()+`decimals`])
			token1Decimals := helpers.DecodeUint64(response[pair.Token1.Hex()+`decimals`])
			token0Price := prices[0]
			token1Price := prices[1]

			if token0Price.IsZero() && addresses.Equals(pair.Token0, OPT_USDC_ADDRESS) {
				token0Price = bigNumber.NewInt(1e18)
			}
			if token1Price.IsZero() && addresses.Equals(pair.Token1, OPT_USDC_ADDRESS) {
				token1Price = bigNumber.NewInt(1e18)
			}

			token0PriceUSD := helpers.ToNormalizedAmount(token0Price, 18)
			token1PriceUSD := helpers.ToNormalizedAmount(token1Price, 18)
			token0totalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Reserve0), uint64(token0Decimals))
			token1totalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.Reserve1), uint64(token1Decimals))
			pairTotalSupply := helpers.ToNormalizedAmount(bigNumber.SetInt(pair.TotalSupply), 18)

			token0ValueInPair := bigNumber.NewFloat(0).Mul(token0PriceUSD, token0totalSupply)
			token1ValueInPair := bigNumber.NewFloat(0).Mul(token1PriceUSD, token1totalSupply)
			valueInPair := bigNumber.NewFloat(0).Add(token0ValueInPair, token1ValueInPair)

			pairPrice := bigNumber.NewFloat(0).Div(valueInPair, pairTotalSupply)
			pairPrice = bigNumber.NewFloat(0).Mul(pairPrice, bigNumber.NewFloat(1e6))
			token0Price = bigNumber.NewFloat(0).Mul(token0PriceUSD, bigNumber.NewFloat(1e6)).Int()
			token1Price = bigNumber.NewFloat(0).Mul(token1PriceUSD, bigNumber.NewFloat(1e6)).Int()

			if pairToken, _ := storage.GetERC20(chainID, pair.PairAddress); !pairToken.IsVaultLike() {
				if !pairPrice.IsZero() {
					humanizedPrice := helpers.ToNormalizedAmount(pairPrice.Int(), 6)
					priceMap[pair.PairAddress] = models.TPrices{
						Address:        pair.PairAddress,
						Price:          pairPrice.Int(),
						HumanizedPrice: humanizedPrice,
						Source:         `veloSugar`,
					}
				}
			}

			if token0, _ := storage.GetERC20(chainID, pair.Token0); !token0.IsVaultLike() {
				if !token0Price.IsZero() {
					humanizedPrice := helpers.ToNormalizedAmount(token0Price, 6)
					priceMap[pair.Token0] = models.TPrices{
						Address:        pair.Token0,
						Price:          token0Price,
						HumanizedPrice: humanizedPrice,
						Source:         `veloSugar`,
					}
				}
			}

			if token1, _ := storage.GetERC20(chainID, pair.Token1); !token1.IsVaultLike() {
				if !token1Price.IsZero() {
					humanizedPrice := helpers.ToNormalizedAmount(token1Price, 6)
					priceMap[pair.Token1] = models.TPrices{
						Address:        pair.Token1,
						Price:          token1Price,
						HumanizedPrice: humanizedPrice,
						Source:         `veloSugar`,
					}
				}
			}
		}
	}

	return priceMap
}
