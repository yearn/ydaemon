package tokens

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()
var YearnVaultABI, _ = contracts.Yvault043MetaData.GetAbi()
var CurvePoolRegistryABI, _ = contracts.CurvePoolRegistryMetaData.GetAbi()
var CurvePoolFactoryABI, _ = contracts.CurvePoolFactoryMetaData.GetAbi()
var CurvePoolCoinABI, _ = contracts.CurvePoolCoinMetaData.GetAbi()
var CurveCoinABI, _ = contracts.ICurveFiMetaData.GetAbi()
var CTokenABI, _ = contracts.CTokenMetaData.GetAbi()
var ATokenV1ABI, _ = contracts.ATokenV1MetaData.GetAbi()
var ATokenV2ABI, _ = contracts.ATokenV2MetaData.GetAbi()

func getName(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("name")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `name`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `name`,
		CallData: parsedData,
		Name:     name,
	}
}
func getSymbol(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("symbol")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `symbol`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `symbol`,
		CallData: parsedData,
		Name:     name,
	}
}
func getDecimals(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("decimals")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `decimals`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `decimals`,
		CallData: parsedData,
		Name:     name,
	}
}
func getToken(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("token")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      YearnVaultABI,
			Method:   `token`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `token`,
		CallData: parsedData,
		Name:     name,
	}
}
func getPoolFromLpToken(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := CurvePoolRegistryABI.Pack("get_pool_from_lp_token", tokenAddress)
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      CurvePoolRegistryABI,
			Method:   `get_pool_from_lp_token`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolRegistryABI,
		Method:   `get_pool_from_lp_token`,
		CallData: parsedData,
		Name:     name,
	}
}
func getCompoundUnderlying(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CTokenABI.Pack("underlying")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      CTokenABI,
			Method:   `underlying`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CTokenABI,
		Method:   `underlying`,
		CallData: parsedData,
		Name:     name,
	}
}
func getAaveV1Underlying(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ATokenV1ABI.Pack("underlyingAssetAddress")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ATokenV1ABI,
			Method:   `underlyingAssetAddress`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ATokenV1ABI,
		Method:   `underlyingAssetAddress`,
		CallData: parsedData,
		Name:     name,
	}
}
func getAaveV2Underlying(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ATokenV2ABI.Pack("UNDERLYING_ASSET_ADDRESS")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ATokenV2ABI,
			Method:   `UNDERLYING_ASSET_ADDRESS`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ATokenV2ABI,
		Method:   `UNDERLYING_ASSET_ADDRESS`,
		CallData: parsedData,
		Name:     name,
	}
}

func getCurveFactoryPool(name string, contractAddress common.Address, index *big.Int) ethereum.Call {
	parsedData, _ := CurvePoolFactoryABI.Pack("pool_list", index)
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolFactoryABI,
		Method:   `pool_list`,
		CallData: parsedData,
		Name:     name,
	}
}
func getCurveFactoryCoin(name string, contractAddress common.Address, poolAddress common.Address) ethereum.Call {
	parsedData, _ := CurvePoolFactoryABI.Pack("get_coins", poolAddress)
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolFactoryABI,
		Method:   `get_coins`,
		CallData: parsedData,
		Name:     name,
	}
}
func getCurveMinter(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := CurvePoolCoinABI.Pack("minter")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurvePoolCoinABI,
		Method:   `minter`,
		CallData: parsedData,
		Name:     name,
	}
}
func getCurveCoin(name string, contractAddress common.Address, index *big.Int) ethereum.Call {
	parsedData, _ := CurveCoinABI.Pack("coins", index)
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CurveCoinABI,
		Method:   `coins`,
		CallData: parsedData,
		Name:     name,
	}
}
