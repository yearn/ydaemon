package multicalls

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

var YearnVaultABI, _ = contracts.Yvault043MetaData.GetAbi()
var YearnVaultV3ABI, _ = contracts.Yvault300MetaData.GetAbi()
var AccountantABI, _ = contracts.AccountantMetaData.GetAbi()

func GetToken(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("token")
	if err != nil {
		logs.Error("Error packing YearnVaultABI token", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `token`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetPricePerShare(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("pricePerShare")
	if err != nil {
		logs.Error("Error packing YearnVaultABI pricePerShare", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `pricePerShare`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetConvertPricePerShare(name string, contractAddress common.Address, value *big.Int) ethereum.Call {
	parsedData, err := YearnVaultV3ABI.Pack("convertToAssets", value)
	if err != nil {
		logs.Error("Error packing YearnVaultV3ABI convertToAssets", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultV3ABI,
		Method:   `convertToAssets`,
		CallData: parsedData,
		Name:     name,
	}
}

func GetAPIVersion(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("apiVersion")
	if err != nil {
		logs.Error("Error packing YearnVaultABI apiVersion", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `apiVersion`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetPerformanceFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("performanceFee")
	if err != nil {
		logs.Error("Error packing YearnVaultABI performanceFee", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `performanceFee`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetManagementFee(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("managementFee")
	if err != nil {
		logs.Error("Error packing YearnVaultABI managementFee", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `managementFee`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetEmergencyShutdown(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("emergencyShutdown")
	if err != nil {
		logs.Error("Error packing YearnVaultABI emergencyShutdown", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `emergencyShutdown`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetActivation(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("activation")
	if err != nil {
		logs.Error("Error packing YearnVaultABI activation", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `activation`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetGovernance(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("governance")
	if err != nil {
		logs.Error("Error packing YearnVaultABI governance", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `governance`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetRoleManager(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultV3ABI.Pack("role_manager")
	if err != nil {
		logs.Error("Error packing YearnVaultV3ABI role_manager", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultV3ABI,
		Method:   `role_manager`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetAccountant(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultV3ABI.Pack("accountant")
	if err != nil {
		logs.Error("Error packing YearnVaultV3ABI accountant", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultV3ABI,
		Method:   `accountant`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetDefaultFeeConfig(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := AccountantABI.Pack("defaultConfig")
	if err != nil {
		logs.Error("Error packing AccountantABI defaultConfig", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      AccountantABI,
		Method:   `defaultConfig`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetGuardian(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("guardian")
	if err != nil {
		logs.Error("Error packing YearnVaultABI guardian", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `guardian`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetManagement(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("management")
	if err != nil {
		logs.Error("Error packing YearnVaultABI management", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `management`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetRewards(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("rewards")
	if err != nil {
		logs.Error("Error packing YearnVaultABI rewards", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `rewards`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetTotalAssets(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("totalAssets")
	if err != nil {
		logs.Error("Error packing YearnVaultABI totalAssets", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `totalAssets`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetDepositLimit(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("depositLimit")
	if err != nil {
		logs.Error("Error packing YearnVaultABI depositLimit", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `depositLimit`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetAvailableDepositLimit(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("availableDepositLimit")
	if err != nil {
		logs.Error("Error packing YearnVaultABI availableDepositLimit", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `availableDepositLimit`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetVaultWithdrawalQueue(name string, index int64, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("withdrawalQueue", big.NewInt(index))
	if err != nil {
		logs.Error("Error packing YearnVaultABI index", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `withdrawalQueue`,
		CallData: parsedData,
		Name:     name + strconv.FormatInt(int64(index), 10),
	}
}
func GetDefaultQueue(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultV3ABI.Pack("get_default_queue")
	if err != nil {
		logs.Error("Error packing YearnVaultV3ABI get_default_queue", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultV3ABI,
		Method:   `get_default_queue`,
		CallData: parsedData,
		Name:     name,
	}
}
func GetCreditAvailable(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("creditAvailable0", strategyAddress)
	if err != nil {
		logs.Error("Error packing YearnVaultABI creditAvailable0", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `creditAvailable0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}
func GetDebtOutstanding(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("debtOutstanding0", strategyAddress)
	if err != nil {
		logs.Error("Error packing YearnVaultABI debtOutstanding0", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `debtOutstanding0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}
func GetStrategies(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
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

	parsedData, err := abiToUse.Pack("strategies", strategyAddress)
	if err != nil {
		logs.Error("Error packing YearnVaultABI strategies", err)
	}

	return ethereum.Call{
		Name:     name,
		Target:   contractAddress,
		Abi:      abiToUse,
		Method:   `strategies`,
		CallData: parsedData,
		Version:  version,
	}
}

func GetV3Strategies(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnVaultV3ABI.Pack("strategies", strategyAddress)
	if err != nil {
		logs.Error("Error packing YearnVaultV3ABI strategies", err)
	}

	return ethereum.Call{
		Name:     name,
		Target:   contractAddress,
		Abi:      YearnVaultV3ABI,
		Method:   `strategies`,
		CallData: parsedData,
		Version:  version,
	}
}

func GetExpectedReturn(name string, contractAddress common.Address, strategyAddress common.Address, version string) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("expectedReturn0", strategyAddress)
	if err != nil {
		logs.Error("Error packing YearnVaultABI expectedReturn0", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `expectedReturn0`,
		CallData: parsedData,
		Name:     name,
		Version:  version,
	}
}
func GetDebtRatio(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := YearnVaultABI.Pack("debtRatio")
	if err != nil {
		logs.Error("Error packing YearnVaultABI debtRatio", err)
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      YearnVaultABI,
		Method:   `debtRatio`,
		CallData: parsedData,
		Name:     name,
	}
}
