package vaults

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
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
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `pricePerShare`,
		CallData: parsedData,
		Name:     name,
	}
}
func getAPIVersion(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("apiVersion")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `apiVersion`,
		CallData: parsedData,
		Name:     name,
	}
}
func getPerformanceFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("performanceFee")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `performanceFee`,
		CallData: parsedData,
		Name:     name,
	}
}
func getManagementFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("managementFee")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `managementFee`,
		CallData: parsedData,
		Name:     name,
	}
}
func getEmergencyShutdown(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("emergencyShutdown")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `emergencyShutdown`,
		CallData: parsedData,
		Name:     name,
	}
}
func getActivation(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("activation")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `activation`,
		CallData: parsedData,
		Name:     name,
	}
}
func getGovernance(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("governance")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `governance`,
		CallData: parsedData,
		Name:     name,
	}
}
func getGuardian(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("guardian")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `guardian`,
		CallData: parsedData,
		Name:     name,
	}
}
func getManagement(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("management")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `management`,
		CallData: parsedData,
		Name:     name,
	}
}
func getRewards(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("rewards")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `rewards`,
		CallData: parsedData,
		Name:     name,
	}
}
func getTotalAssets(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("totalAssets")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `totalAssets`,
		CallData: parsedData,
		Name:     name,
	}
}
func getDepositLimit(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("depositLimit")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `depositLimit`,
		CallData: parsedData,
		Name:     name,
	}
}
func getAvailableDepositLimit(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("availableDepositLimit")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `availableDepositLimit`,
		CallData: parsedData,
		Name:     name,
	}
}
func getToken(name string, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("token")
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `token`,
		CallData: parsedData,
		Name:     name,
	}
}
func getVaultWithdrawalQueue(name string, index int64, contractAddress common.Address) ethereum.Call {
	parsedData, _ := YearnVaultABI.Pack("withdrawalQueue", big.NewInt(index))
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `withdrawalQueue`,
		CallData: parsedData,
		Name:     name + strconv.FormatInt(int64(index), 10),
	}
}
