package strategies

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/types/common"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var YearnVaultABI, _ = contracts.Yvault043MetaData.GetAbi()
var YearnStrategyABI, _ = contracts.StrategyBaseMetaData.GetAbi()

func getCreditAvailable(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("creditAvailable0", strategyAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `creditAvailable0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getDebtOutstanding(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("debtOutstanding0", strategyAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `debtOutstanding0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStrategies(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	var abiToUse *abi.ABI

	switch version {
	case `0.2.2`:
		_abi, _ := abi.JSON(strings.NewReader(helpers.YEARN_VAULT_V022_ABI))
		abiToUse = &_abi
	case `0.3.0`, `0.3.1`:
		_abi, _ := abi.JSON(strings.NewReader(helpers.YEARN_VAULT_V030_ABI))
		abiToUse = &_abi
	default:
		abiToUse = YearnVaultABI
	}

	parsedData, _ := abiToUse.Pack("strategies", strategyAddress.Address)

	return ethereum.Call{
		Name:     name,
		Target:   contractAddress.Address,
		Abi:      abiToUse,
		Method:   `strategies`,
		CallData: parsedData,
		Version:  version,
	}
}

func getExpectedReturn(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("expectedReturn0", strategyAddress.Address)
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `expectedReturn0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyEstimatedTotalAsset(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("estimatedTotalAssets")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `estimatedTotalAssets`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyIsActive(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("isActive")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `isActive`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyKeepCRV(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("keepCRV")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `keepCRV`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStategyDelegatedAssets(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("delegatedAssets")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `delegatedAssets`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getName(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("name")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `name`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getKeeper(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("keeper")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `keeper`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getStrategist(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("strategist")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `strategist`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getRewards(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("rewards")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `rewards`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getHealthCheck(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("healthCheck")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `healthCheck`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getAPIVersion(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("apiVersion")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `apiVersion`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getDoHealthCheck(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("doHealthCheck")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `doHealthCheck`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func getEmergencyExit(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, _ := YearnStrategyABI.Pack("emergencyExit")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnStrategyABI,
		Method:   `emergencyExit`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}
