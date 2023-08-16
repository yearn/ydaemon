package multicalls

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var YearnStrategyABI, _ = contracts.StrategyBaseMetaData.GetAbi()
var YearnStrategyVeloABI, _ = contracts.YStrategyVeloMetaData.GetAbi()

func GetStategyEstimatedTotalAsset(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("estimatedTotalAssets")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI estimatedTotalAssets", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `estimatedTotalAssets`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStategyIsActive(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("isActive")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI isActive", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `isActive`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStategyKeepCRV(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("keepCRV")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI keepCRV", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `keepCRV`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStategyKeepCRVPercent(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("keepCrvPercent")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI keepCrvPercent", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `keepCrvPercent`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStategyKeepCVX(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("keepCVX")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI keepCVX", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `keepCVX`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStategyDelegatedAssets(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("delegatedAssets")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI delegatedAssets", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `delegatedAssets`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStrategyName(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("name")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI name", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `name`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetKeeper(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("keeper")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI keeper", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `keeper`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStrategist(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("strategist")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI strategist", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `strategist`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStrategyRewards(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("rewards")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI rewards", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `rewards`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetHealthCheck(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("healthCheck")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI healthCheck", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `healthCheck`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStrategyAPIVersion(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("apiVersion")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI apiVersion", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `apiVersion`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetDoHealthCheck(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("doHealthCheck")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI doHealthCheck", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `doHealthCheck`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetEmergencyExit(name string, contractAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnStrategyABI.Pack("emergencyExit")
	if err != nil {
		logs.Error("Error packing YearnStrategyABI emergencyExit", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyABI,
		Method:   `emergencyExit`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}

func GetStategyLocalKeepVelo(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnStrategyVeloABI.Pack("localKeepVELO")
	if err != nil {
		logs.Error("Error packing YearnStrategyVeloABI localKeepVELO", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnStrategyVeloABI,
		Method:   `localKeepVELO`,
		CallData: parsedData,
		Name:     name,
	}
}
