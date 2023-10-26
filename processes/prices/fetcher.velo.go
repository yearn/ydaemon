package prices

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/multicalls"
)

var VELO_SUGAR_ADDRESS = common.HexToAddress(`0x4D996E294B00cE8287C16A2b9A4e637ecA5c939f`)
var VELO_SUGAR_ORACLE_ADDRESS = common.HexToAddress(`0x395942c2049604a314d39f370dfb8d87aac89e16`)
var OPT_WETH_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000006`)
var OPT_OP_ADDRESS = common.HexToAddress(`0x4200000000000000000000000000000000000042`)
var OPT_USDC_ADDRESS = common.HexToAddress(`0x7F5c764cBc14f9669B88837ca1490cCa17c31607`)
var OPT_RATE_CONNECTORS = []common.Address{
	common.HexToAddress(`0x9560e827af36c94d2ac33a39bce1fe78631088db`),
	common.HexToAddress(`0x46f21fda29f1339e0ab543763ff683d399e393ec`),
	common.HexToAddress(`0x3417e54a51924c225330f8770514ad5560b9098d`),
	common.HexToAddress(`0x7f5c764cbc14f9669b88837ca1490cca17c31607`),
	common.HexToAddress(`0x7ae97042a4a0eb4d1eb370c34bfec71042a056b7`),
	common.HexToAddress(`0x4200000000000000000000000000000000000006`),
	common.HexToAddress(`0x3e29d3a9316dab217754d13b28646b76607c5f04`),
	common.HexToAddress(`0x6806411765af15bddd26f8f544a34cc40cb9838b`),
	common.HexToAddress(`0x1f32b1c2345538c0c6f582fcb022739c4a194ebb`),
	common.HexToAddress(`0xfdb794692724153d1488ccdbe0c56c252596735f`),
	common.HexToAddress(`0x340fe1d898eccaad394e2ba0fc1f93d27c7b717a`),
	common.HexToAddress(`0xc40f949f8a4e094d1b49a23ea9241d289b7b2819`),
	common.HexToAddress(`0xda10009cbd5d07dd0cecc66161fc93d7c9000da1`),
	common.HexToAddress(`0xc5b001dc33727f8f26880b184090d3e252470d45`),
	common.HexToAddress(`0xdfa46478f9e5ea86d57387849598dbfb2e964b02`),
	common.HexToAddress(`0x10010078a54396f62c96df8532dc2b4847d47ed3`),
	common.HexToAddress(`0x4200000000000000000000000000000000000042`),
	common.HexToAddress(`0x9485aca5bbbe1667ad97c7fe7c4531a624c8b1ed`),
	common.HexToAddress(`0xca0e54b636db823847b29f506bffee743f57729d`),
	common.HexToAddress(`0xcb8fa9a76b8e203d8c3797bf438d8fb81ea3326a`),
	common.HexToAddress(`0x2e3d870790dc77a83dd1d18184acc7439a53f475`),
	common.HexToAddress(`0x73cb180bf0521828d8849bc8cf2b920918e23032`),
	common.HexToAddress(`0x50c5725949a6f0c72e6c4a641f24049a917db0cb`),
	common.HexToAddress(`0xb153fb3d196a8eb25522705560ac152eeec57901`),
	common.HexToAddress(`0x8ae125e8653821e851f12a49f7765db9a9ce7384`),
	common.HexToAddress(`0x4e720dd3ac5cfe1e1fbde4935f386bb1c66f4642`),
	common.HexToAddress(`0x94b008aa00579c1307b0ef2c499ad98a8ce58e58`),
	common.HexToAddress(`0x39fde572a18448f8139b7788099f0a0740f51205`),
	common.HexToAddress(`0x970d50d09f3a656b43e11b0d45241a84e3a6e011`),
	common.HexToAddress(`0x920cf626a271321c151d027030d5d08af699456b`),
	common.HexToAddress(`0x8700daec35af8ff88c16bdf0418774cb3d7599b4`),
	common.HexToAddress(`0x1db2466d9f5e10d7090e7152b68d62703a2245f0`),
	common.HexToAddress(`0x1610e3c85dd44af31ed7f33a63642012dca0c5a5`),
	common.HexToAddress(`0x9e5aac1ba1a2e6aed6b32689dfcf62a509ca96f3`),
	common.HexToAddress(`0xbfd291da8a403daaf7e5e9dc1ec0aceacd4848b9`),
	common.HexToAddress(`0xb0b195aefa3650a6908f15cdac7d92f8a5791b0b`),
	common.HexToAddress(`0x9bcef72be871e61ed4fbbc7630889bee758eb81d`),
	common.HexToAddress(`0x375488f097176507e39b9653b88fdc52cde736bf`),
	common.HexToAddress(`0x9e1028f5f1d5ede59748ffcee5532509976840e0`),
	common.HexToAddress(`0x79af5dd14e855823fa3e9ecacdf001d99647d043`),
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
	common.HexToAddress(`0x7f5c764cbc14f9669b88837ca1490cca17c31607`),
}

func fetchVelo(url string) []TVeloPairData {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(`impossible to get velo URL`, err.Error())
		return []TVeloPairData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
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
func fetchPricesFromSugar(chainID uint64, blockNumber *uint64, tokens []common.Address) map[common.Address]*bigNumber.Int {
	newPriceMap := make(map[common.Address]*bigNumber.Int)
	if chainID != 10 {
		return newPriceMap
	}
	newPairPriceMap := make(map[common.Address]*bigNumber.Int)
	newTokensPriceMap := make(map[common.Address]*bigNumber.Int)

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	sugar, _ := contracts.NewVeloSugarCaller(VELO_SUGAR_ADDRESS, client)
	// sugarOracle, _ := contracts.NewVeloSugarOracleCaller(VELO_SUGAR_ORACLE_ADDRESS, client)
	allSugar, _ := sugar.All(nil, big.NewInt(1000), big.NewInt(0), common.Address{})

	for _, pair := range allSugar {
		ratesConnector := []common.Address{pair.Token0, pair.Token1}
		ratesConnector = append(ratesConnector, OPT_RATE_CONNECTORS...)
		calls := []ethereum.Call{}
		calls = append(calls, multicalls.GetManyRatesWithConnectors(VELO_SUGAR_ORACLE_ADDRESS.Hex(), VELO_SUGAR_ORACLE_ADDRESS, 2, ratesConnector))
		calls = append(calls, multicalls.GetDecimals(pair.Token0.Hex(), pair.Token0))
		calls = append(calls, multicalls.GetDecimals(pair.Token1.Hex(), pair.Token1))
		response := multicalls.Perform(chainID, calls, nil)

		prices := helpers.DecodeBigInts(response[VELO_SUGAR_ORACLE_ADDRESS.Hex()+`getManyRatesWithConnectors`])
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

		if pairToken, _ := store.GetERC20(chainID, pair.PairAddress); !store.IsVaultLike(pairToken) {
			if !pairPrice.IsZero() {
				newPairPriceMap[pair.PairAddress] = pairPrice.Int()
			}
		}

		if token0, _ := store.GetERC20(chainID, pair.Token0); !store.IsVaultLike(token0) {
			if !token0Price.IsZero() {
				newTokensPriceMap[pair.Token0] = token0Price
			}
		}

		if token1, _ := store.GetERC20(chainID, pair.Token1); !store.IsVaultLike(token1) {
			if !token1Price.IsZero() {
				newTokensPriceMap[pair.Token1] = token1Price
			}
		}
	}

	for token, price := range newPairPriceMap {
		newPriceMap[token] = price
	}
	for token, price := range newTokensPriceMap {
		if _, ok := newPriceMap[token]; !ok {
			newPriceMap[token] = price
		}
	}

	return newPriceMap
}
