package vaults

import (
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/types/common"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var YearnVaultABI, _ = contracts.Yvault043MetaData.GetAbi()
var YearnVaultRegistryV2, _ = contracts.Yregistryv2MetaData.GetAbi()

func getPricePerShare(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("pricePerShare")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `pricePerShare`,
		CallData: parsedData,
		Name:     name,
	}
}
func getAPIVersion(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("apiVersion")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `apiVersion`,
		CallData: parsedData,
		Name:     name,
	}
}
func getPerformanceFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("performanceFee")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `performanceFee`,
		CallData: parsedData,
		Name:     name,
	}
}
func getManagementFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("managementFee")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `managementFee`,
		CallData: parsedData,
		Name:     name,
	}
}
func getEmergencyShutdown(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("emergencyShutdown")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `emergencyShutdown`,
		CallData: parsedData,
		Name:     name,
	}
}
func getActivation(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("activation")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `activation`,
		CallData: parsedData,
		Name:     name,
	}
}
func getGovernance(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("governance")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `governance`,
		CallData: parsedData,
		Name:     name,
	}
}
func getGuardian(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("guardian")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `guardian`,
		CallData: parsedData,
		Name:     name,
	}
}
func getManagement(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("management")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `management`,
		CallData: parsedData,
		Name:     name,
	}
}
func getRewards(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("rewards")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `rewards`,
		CallData: parsedData,
		Name:     name,
	}
}
func getTotalAssets(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("totalAssets")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `totalAssets`,
		CallData: parsedData,
		Name:     name,
	}
}
func getDepositLimit(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("depositLimit")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `depositLimit`,
		CallData: parsedData,
		Name:     name,
	}
}
func getAvailableDepositLimit(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("availableDepositLimit")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `availableDepositLimit`,
		CallData: parsedData,
		Name:     name,
	}
}
func getToken(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("token")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress.Address,
			Abi:      YearnVaultABI,
			Method:   `token`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      YearnVaultABI,
		Method:   `token`,
		CallData: parsedData,
		Name:     name,
	}
}
