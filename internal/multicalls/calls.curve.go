package multicalls

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

var CURVE_REGISTRY_ABI, _ = contracts.CurveRegistryMetaData.GetAbi()
var CURVE_ABI, _ = contracts.ICurveFiMetaData.GetAbi()

func GetVirtualPriceFromLP(name string, contractAddress common.Address, tokenAddress common.Address) ethereum.Call {
	parsedData, err := CURVE_REGISTRY_ABI.Pack("get_virtual_price_from_lp_token", tokenAddress)
	if err != nil {
		logs.Error("Error packing CURVE_REGISTRY_ABI get_virtual_price_from_lp_token", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CURVE_REGISTRY_ABI,
		Method:   `get_virtual_price_from_lp_token`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetVirtualPrice(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := CURVE_ABI.Pack("get_virtual_price")
	if err != nil {
		logs.Error("Error packing CURVE_ABI get_virtual_price", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      CURVE_ABI,
		Method:   `get_virtual_price`,
		CallData: parsedData,
		Name:     name,
	}
}
