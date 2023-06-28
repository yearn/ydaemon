package multicalls

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var CurvePoolRegistryABI, _ = contracts.CurvePoolRegistryMetaData.GetAbi()
var CurvePoolFactoryABI, _ = contracts.CurvePoolFactoryMetaData.GetAbi()
var CurvePoolCoinABI, _ = contracts.CurvePoolCoinMetaData.GetAbi()
var CurveCoinABI, _ = contracts.ICurveFiMetaData.GetAbi()
var CTokenABI, _ = contracts.CTokenMetaData.GetAbi()
var ATokenV1ABI, _ = contracts.ATokenV1MetaData.GetAbi()
var ATokenV2ABI, _ = contracts.ATokenV2MetaData.GetAbi()
var LensABI, _ = contracts.OracleMetaData.GetAbi()
var CurveGaugeABI, _ = contracts.CurveGaugeMetaData.GetAbi()
var CVXBoosterABI, _ = contracts.CVXBoosterMetaData.GetAbi()
var CrvUSDABI, _ = contracts.CrvUSDMetaData.GetAbi()

func GetPriceUsdcRecommendedCall(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := LensABI.Pack("getPriceUsdcRecommended", tokenAddress)
	if err != nil {
		logs.Error("Error packing LensABI getPriceUsdcRecommended", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      LensABI,
		Method:   `getPriceUsdcRecommended`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetPriceCrvUsdcCall(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CrvUSDABI.Pack("price")
	if err != nil {
		logs.Error("Error packing LensABI price", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CrvUSDABI,
		Method:   `price`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetPoolFromLpToken(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := CurvePoolRegistryABI.Pack("get_pool_from_lp_token", tokenAddress)
	if err != nil {
		logs.Error("Error packing CurvePoolRegistryABI get_pool_from_lp_token", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolRegistryABI,
		Method:   `get_pool_from_lp_token`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCompoundUnderlying(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CTokenABI.Pack("underlying")
	if err != nil {
		logs.Error("Error packing CTokenABI underlying", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CTokenABI,
		Method:   `underlying`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetAaveV1Underlying(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ATokenV1ABI.Pack("underlyingAssetAddress")
	if err != nil {
		logs.Error("Error packing ATokenV1ABI underlyingAssetAddress", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ATokenV1ABI,
		Method:   `underlyingAssetAddress`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetAaveV2Underlying(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ATokenV2ABI.Pack("UNDERLYING_ASSET_ADDRESS")
	if err != nil {
		logs.Error("Error packing ATokenV2ABI UNDERLYING_ASSET_ADDRESS", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ATokenV2ABI,
		Method:   `UNDERLYING_ASSET_ADDRESS`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCurveFactoryPool(name string, contractAddress common.Address, index *big.Int) ethereum.Call {
	parsedData, err := CurvePoolFactoryABI.Pack("pool_list", index)
	if err != nil {
		logs.Error("Error packing CurvePoolFactoryABI pool_list", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolFactoryABI,
		Method:   `pool_list`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCurveFactoryCoin(name string, contractAddress common.Address, poolAddress common.Address) ethereum.Call {
	parsedData, err := CurvePoolFactoryABI.Pack("get_coins", poolAddress)
	if err != nil {
		logs.Error("Error packing CurvePoolFactoryABI get_coins", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolFactoryABI,
		Method:   `get_coins`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCurveMinter(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CurvePoolCoinABI.Pack("minter")
	if err != nil {
		logs.Error("Error packing CurvePoolCoinABI minter", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolCoinABI,
		Method:   `minter`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCurveCoin(name string, contractAddress common.Address, index *big.Int) ethereum.Call {
	parsedData, err := CurveCoinABI.Pack("coins", index)
	if err != nil {
		logs.Error("Error packing CurveCoinABI coins", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurveCoinABI,
		Method:   `coins`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetConvexLockIncentive(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CVXBoosterABI.Pack("lockIncentive")
	if err != nil {
		logs.Error("Error packing GetConvexLockIncentive lockIncentive", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CVXBoosterABI,
		Method:   `lockIncentive`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetConvexStakerIncentive(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CVXBoosterABI.Pack("stakerIncentive")
	if err != nil {
		logs.Error("Error packing GetConvexStakerIncentive stakerIncentive", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CVXBoosterABI,
		Method:   `stakerIncentive`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetConvexEarmarkIncentive(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CVXBoosterABI.Pack("earmarkIncentive")
	if err != nil {
		logs.Error("Error packing GetConvexEarmarkIncentive earmarkIncentive", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CVXBoosterABI,
		Method:   `earmarkIncentive`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetConvexPlatformFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CVXBoosterABI.Pack("platformFee")
	if err != nil {
		logs.Error("Error packing GetConvexPlatformFee platformFee", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CVXBoosterABI,
		Method:   `platformFee`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCurveWorkingBalance(name string, contractAddress common.Address, voter common.Address) ethereum.Call {
	parsedData, err := CurveGaugeABI.Pack("working_balances", voter)
	if err != nil {
		logs.Error("Error packing GetCurveWorkingBalance working_balances", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurveGaugeABI,
		Method:   `working_balances`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetCurveBalanceOf(name string, contractAddress common.Address, voter common.Address) ethereum.Call {
	parsedData, err := CurveGaugeABI.Pack("balanceOf", voter)
	if err != nil {
		logs.Error("Error packing GetCurveBalanceOf balanceOf", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurveGaugeABI,
		Method:   `balanceOf`,
		CallData: parsedData,
		Name:     name,
	}
}
