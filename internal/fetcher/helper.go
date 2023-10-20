package fetcher

import (
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**********************************************************************************************
** Get the price of the underlying asset. This is tricky because of the decimals. The prices are fetched
** using the lens oracle daemon, stored in the TokenPrices map, and based on the USDC token, aka with 6
** decimals. We first need to parse the Int Price to a float64, then divide by 10^6 to get the price
** in an human readable USDC format.
**********************************************************************************************/
func getHumanizedTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	price, ok := storage.GetPrice(chainID, tokenAddress)
	if ok {
		fHumanizedPrice, _ := price.HumanizedPrice.Float64()
		return price.HumanizedPrice, fHumanizedPrice
	}
	return bigNumber.NewFloat(), 0.0
}

/**********************************************************************************************
** Get the total assets locked in this vault. This is tricky because of the decimals. The total asset value
** is a string which will be formated as a float64 and scaled with the underlying token decimals. With that
** we should have a human readable Total Asset value, and we should be able to get the Total Value Locked
** in the vault thanks to the above humanizedPrice value.
**********************************************************************************************/
func getHumanizedValue(balanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, humanizedValue := helpers.FormatAmount(balanceToken.String(), decimals)
	fHumanizedValue, _ := bigNumber.NewFloat().Mul(humanizedValue, humanizedPrice).Float64()
	return fHumanizedValue
}

/**********************************************************************************************
** Prepare the multicall to get the basic informations for the V3 vaults
**********************************************************************************************/
func getV3VaultCalls(vault models.TVault) []ethereum.Call {
	metadata := storage.GetStrategiesJsonMetadata(vault.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh

	calls := []ethereum.Call{}
	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetPricePerShare(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetTotalAssets(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetDefaultQueue(vault.Address.Hex(), vault.Address))

	if time.Since(lastUpdate).Hours() > 1 || shouldRefresh {
		// If the last vault update was more than 1 hour ago, we will do a partial update
		calls = append(calls, multicalls.GetShutdown(vault.Address.Hex(), vault.Address))
	}
	if time.Since(lastUpdate).Hours() > 24 || shouldRefresh {
		// If the last vault update was more than 24 hour ago, we will do a full update
		calls = append(calls, multicalls.GetAsset(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetV3APIVersion(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetRoleManager(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetAccountant(vault.Address.Hex(), vault.Address))
	}
	return calls
}

/**********************************************************************************************
** Prepare the multicall to get the basic informations for the V2 and earlier vaults
**********************************************************************************************/
func getV2VaultCalls(vault models.TVault) []ethereum.Call {
	metadata := storage.GetStrategiesJsonMetadata(vault.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh

	calls := []ethereum.Call{}
	maxStrategiesPerVault := 20

	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetPricePerShare(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetTotalAssets(vault.Address.Hex(), vault.Address))
	for i := 0; i < maxStrategiesPerVault; i++ {
		calls = append(calls, multicalls.GetVaultWithdrawalQueue(vault.Address.Hex(), int64(i), vault.Address))
	}
	if time.Since(lastUpdate).Hours() > 1 || shouldRefresh {
		// If the last vault update was more than 1 hour ago, we will do a partial update
		calls = append(calls, multicalls.GetPerformanceFee(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetManagementFee(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetEmergencyShutdown(vault.Address.Hex(), vault.Address))
	}
	if time.Since(lastUpdate).Hours() > 24 || shouldRefresh {
		// If the last vault update was more than 24 hour ago, we will do a full update
		calls = append(calls, multicalls.GetAPIVersion(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetV3APIVersion(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetToken(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetAccountant(vault.Address.Hex(), vault.Address))
	}
	return calls
}

/**********************************************************************************************
** Prepare the multicall to get the basic informations for the V2 and earlier strategies
**********************************************************************************************/
func getV2StrategyCalls(strat models.TStrategy) []ethereum.Call {
	metadata := storage.GetStrategiesJsonMetadata(strat.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh

	calls := []ethereum.Call{}
	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetStrategies(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
	calls = append(calls, multicalls.GetStategyEstimatedTotalAsset(strat.Address.Hex(), strat.Address, strat.VaultVersion))
	if time.Since(lastUpdate).Hours() > 1 || shouldRefresh {
		// If the last strat update was more than 1 hour ago, we will do a partial update
		calls = append(calls, multicalls.GetStategyKeepCRV(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCRVPercent(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCVX(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetEmergencyExit(strat.Address.Hex(), strat.Address, strat.VaultVersion))
	}
	if time.Since(lastUpdate).Hours() > 24 || shouldRefresh {
		// If the last strat update was more than 24 hour ago, we will do a full update
		calls = append(calls, multicalls.GetStategyIsActive(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStrategyName(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetDoHealthCheck(strat.Address.Hex(), strat.Address, strat.VaultVersion))
	}
	return calls
}

/**********************************************************************************************
** Prepare the multicall to get the basic informations for the V3 strategies
**********************************************************************************************/
func getV3StrategyCalls(strat models.TStrategy) []ethereum.Call {
	metadata := storage.GetStrategiesJsonMetadata(strat.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh

	calls := []ethereum.Call{}
	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetV3Strategies(strat.Address.Hex(), strat.VaultAddress, strat.Address, strat.VaultVersion))
	calls = append(calls, multicalls.GetPerformanceFee(strat.Address.Hex(), strat.Address))
	calls = append(calls, multicalls.GetTotalAssets(strat.VaultAddress.Hex(), strat.VaultAddress))
	calls = append(calls, multicalls.GetTotalAssets(strat.Address.Hex(), strat.Address))
	if time.Since(lastUpdate).Hours() > 1 || shouldRefresh {
		// If the last strat update was more than 1 hour ago, we will do a partial update
		calls = append(calls, multicalls.GetStategyKeepCRV(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCRVPercent(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCVX(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetIsShutdown(strat.Address.Hex(), strat.Address, strat.VaultVersion))
	}
	if time.Since(lastUpdate).Hours() > 24 || shouldRefresh {
		// If the last strat update was more than 24 hour ago, we will do a full update
		calls = append(calls, multicalls.GetStrategyName(strat.Address.Hex(), strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetDoHealthCheck(strat.Address.Hex(), strat.Address, strat.VaultVersion))
	}
	return calls
}

/**********************************************************************************************
** Handle the calls result
**********************************************************************************************/
func handleV3VaultCalls(vault models.TVault, response map[string][]interface{}) models.TVault {
	rawPricePerShare := response[vault.Address.Hex()+`pricePerShare`]
	rawTotalAssets := response[vault.Address.Hex()+`totalAssets`]
	rawDefaultQueue := response[vault.Address.Hex()+`get_default_queue`]
	rawShutdown := response[vault.Address.Hex()+`shutdown`]
	rawUnderlying := response[vault.Address.Hex()+`asset`]
	rawApiVersion := response[vault.Address.Hex()+`api_version`]
	rawAccountant := response[vault.Address.Hex()+`accountant`]

	vault.LastPricePerShare = helpers.DecodeBigInt(rawPricePerShare)
	vault.LastTotalAssets = helpers.DecodeBigInt(rawTotalAssets)
	vault.LastActiveStrategies = helpers.DecodeAddresses(rawDefaultQueue)
	if len(rawShutdown) > 0 {
		vault.EmergencyShutdown = helpers.DecodeBool(rawShutdown)
	}
	if len(rawUnderlying) > 0 {
		vault.AssetAddress = helpers.DecodeAddress(rawUnderlying)
	}
	if len(rawApiVersion) > 0 {
		vault.Version = helpers.DecodeString(rawApiVersion)
	}
	if len(rawAccountant) > 0 {
		vault.Accountant = helpers.DecodeAddress(rawAccountant)
	}

	return vault
}

func handleV2VaultCalls(vault models.TVault, response map[string][]interface{}) models.TVault {
	rawPricePerShare := response[vault.Address.Hex()+`pricePerShare`]
	rawApiVersion := response[vault.Address.Hex()+`apiVersion`]
	rawPerformanceFee := response[vault.Address.Hex()+`performanceFee`]
	rawManagementFee := response[vault.Address.Hex()+`managementFee`]
	rawUnderlying := response[vault.Address.Hex()+`token`]
	rawEmergencyShutdown := response[vault.Address.Hex()+`emergencyShutdown`]
	rawTotalAssets := response[vault.Address.Hex()+`totalAssets`]

	vault.LastPricePerShare = helpers.DecodeBigInt(rawPricePerShare)
	vault.LastTotalAssets = helpers.DecodeBigInt(rawTotalAssets)

	if len(rawPerformanceFee) > 0 {
		vault.PerformanceFee = helpers.DecodeBigInt(rawPerformanceFee).Uint64()
	}
	if len(rawManagementFee) > 0 {
		vault.ManagementFee = helpers.DecodeBigInt(rawManagementFee).Uint64()
	}
	if len(rawEmergencyShutdown) > 0 {
		vault.EmergencyShutdown = helpers.DecodeBool(rawEmergencyShutdown)
	}
	if len(rawUnderlying) > 0 {
		vault.AssetAddress = helpers.DecodeAddress(rawUnderlying)
	}
	if len(rawApiVersion) > 0 {
		vault.Version = helpers.DecodeString(rawApiVersion)
	}

	maxStrategiesPerVault := 20
	withdrawQueueForStrategies := []common.Address{}
	for i := 0; i < maxStrategiesPerVault; i++ {
		result := response[vault.Address.Hex()+strconv.FormatInt(int64(i), 10)+`withdrawalQueue`]
		if len(result) == 1 {
			strategyAddress := helpers.DecodeAddress(result)
			if helpers.AddressIsValid(strategyAddress, vault.ChainID) {
				withdrawQueueForStrategies = append(withdrawQueueForStrategies, strategyAddress)
			}
		}
	}
	vault.LastActiveStrategies = withdrawQueueForStrategies

	return vault
}

func handleV2StrategyCalls(strat models.TStrategy, response map[string][]interface{}) models.TStrategy {
	rawStrategies := response[strat.Address.Hex()+`strategies`]
	rawIsActive := response[strat.Address.Hex()+`isActive`]
	rawKeepCRV := response[strat.Address.Hex()+`keepCRV`]
	rawKeepCRVPercent := response[strat.Address.Hex()+`keepCrvPercent`]
	rawKeepCVX := response[strat.Address.Hex()+`keepCVX`]
	rawName := response[strat.Address.Hex()+`name`]
	rawDoHealthCheck := response[strat.Address.Hex()+`doHealthCheck`]
	rawEmergencyExit := response[strat.Address.Hex()+`emergencyExit`]
	rawEstimatedTotalAssets := response[strat.Address.Hex()+`estimatedTotalAssets`]

	strat.LastEstimatedTotalAssets = helpers.DecodeBigInt(rawEstimatedTotalAssets)
	if strat.VaultVersion == `0.2.2` && len(rawStrategies) == 8 {
		strat.LastReport = bigNumber.SetInt(rawStrategies[4].(*big.Int))
		strat.LastTotalDebt = bigNumber.SetInt(rawStrategies[5].(*big.Int))
		strat.LastTotalGain = bigNumber.SetInt(rawStrategies[6].(*big.Int))
		strat.LastTotalLoss = bigNumber.SetInt(rawStrategies[7].(*big.Int))
		strat.LastPerformanceFee = bigNumber.SetInt(rawStrategies[0].(*big.Int))
		strat.TimeActivated = bigNumber.SetInt(rawStrategies[1].(*big.Int))
	} else if (strat.VaultVersion == `0.3.0` || strat.VaultVersion == `0.3.1`) && len(rawStrategies) == 8 {
		strat.LastReport = bigNumber.SetInt(rawStrategies[4].(*big.Int))
		strat.LastTotalDebt = bigNumber.SetInt(rawStrategies[5].(*big.Int))
		strat.LastTotalGain = bigNumber.SetInt(rawStrategies[6].(*big.Int))
		strat.LastTotalLoss = bigNumber.SetInt(rawStrategies[7].(*big.Int))
		strat.LastPerformanceFee = bigNumber.SetInt(rawStrategies[0].(*big.Int))
		strat.LastDebtRatio = bigNumber.SetInt(rawStrategies[2].(*big.Int))
		strat.TimeActivated = bigNumber.SetInt(rawStrategies[1].(*big.Int))
	} else if len(rawStrategies) == 9 {
		strat.LastReport = bigNumber.SetInt(rawStrategies[5].(*big.Int))
		strat.LastTotalDebt = bigNumber.SetInt(rawStrategies[6].(*big.Int))
		strat.LastTotalGain = bigNumber.SetInt(rawStrategies[7].(*big.Int))
		strat.LastTotalLoss = bigNumber.SetInt(rawStrategies[8].(*big.Int))
		strat.LastPerformanceFee = bigNumber.SetInt(rawStrategies[0].(*big.Int))
		strat.LastDebtRatio = bigNumber.SetInt(rawStrategies[2].(*big.Int))
		strat.TimeActivated = bigNumber.SetInt(rawStrategies[1].(*big.Int))
	}

	if len(rawKeepCRV) > 0 {
		strat.KeepCRV = helpers.DecodeBigInt(rawKeepCRV)
	}
	if len(rawKeepCRVPercent) > 0 {
		strat.KeepCRVPercent = helpers.DecodeBigInt(rawKeepCRVPercent)
	}
	if len(rawKeepCVX) > 0 {
		strat.KeepCVX = helpers.DecodeBigInt(rawKeepCVX)
	}
	if len(rawName) > 0 {
		strat.Name = helpers.DecodeString(rawName)
	}
	if len(rawIsActive) > 0 {
		strat.IsActive = helpers.DecodeBool(rawIsActive)
	}
	if len(rawDoHealthCheck) > 0 {
		strat.DoHealthCheck = helpers.DecodeBool(rawDoHealthCheck)
	}
	if !strat.IsActive || (len(rawEmergencyExit) > 0 && helpers.DecodeBool(rawEmergencyExit)) {
		strat.IsRetired = true
	}

	return strat
}

func handleV3StrategyCalls(strat models.TStrategy, response map[string][]interface{}) models.TStrategy {
	type typeOfRawStrategies = struct {
		Activation  *big.Int "json:\"activation\""
		LastReport  *big.Int "json:\"last_report\""
		CurrentDebt *big.Int "json:\"current_debt\""
		MaxDebt     *big.Int "json:\"max_debt\""
	}

	rawStrategies := response[strat.Address.Hex()+`strategies`]
	rawKeepCRV := response[strat.Address.Hex()+`keepCRV`]
	rawKeepCRVPercent := response[strat.Address.Hex()+`keepCrvPercent`]
	rawKeepCVX := response[strat.Address.Hex()+`keepCVX`]
	rawName := response[strat.Address.Hex()+`name`]
	rawVaultTotalAssets := response[strat.VaultAddress.Hex()+`totalAssets`]
	rawEstimatedTotalAssets := response[strat.Address.Hex()+`totalAssets`]
	rawDoHealthCheck := response[strat.Address.Hex()+`doHealthCheck`]
	rawIsShutdown := response[strat.Address.Hex()+`isShutdown`]
	rawPerformanceFee := response[strat.Address.Hex()+`performanceFee`]

	strat.LastEstimatedTotalAssets = helpers.DecodeBigInt(rawEstimatedTotalAssets)
	strat.LastPerformanceFee = helpers.DecodeBigInt(rawPerformanceFee)
	strat.LastTotalDebt = bigNumber.SetInt(rawStrategies[0].(typeOfRawStrategies).CurrentDebt)
	strat.TimeActivated = bigNumber.SetInt(rawStrategies[0].(typeOfRawStrategies).Activation)
	strat.LastReport = bigNumber.SetInt(rawStrategies[0].(typeOfRawStrategies).LastReport)
	strat.LastTotalGain = bigNumber.NewInt(0) //Not available in V3
	strat.LastTotalLoss = bigNumber.NewInt(0) //Not available in V3
	vaultTotalAssets := helpers.DecodeBigInt(rawVaultTotalAssets)

	// Debt ratio should be a int between 0 and 10000, 10000 being 100%
	// At this point, stratDebtRatioBigFloat is a number between 0 and 1
	// We need to multiply it by 10000 to get the debt ratio
	// Then, we need to convert it to a int, but first we need to round it
	stratDebtRatioBigFloat := bigNumber.NewFloat(0).Div(bigNumber.NewFloat().SetInt(strat.LastTotalDebt), bigNumber.NewFloat().SetInt(vaultTotalAssets))
	stratDebtRatioFloat, _ := stratDebtRatioBigFloat.Float64()
	stratDebtRatioFloat = stratDebtRatioFloat * 10000
	stratDebtRatioInt := int64(math.Round(stratDebtRatioFloat))
	strat.LastDebtRatio = bigNumber.NewUint64(uint64(stratDebtRatioInt))

	if len(rawKeepCRV) > 0 {
		strat.KeepCRV = helpers.DecodeBigInt(rawKeepCRV)
	}
	if len(rawKeepCRVPercent) > 0 {
		strat.KeepCRVPercent = helpers.DecodeBigInt(rawKeepCRVPercent)
	}
	if len(rawKeepCVX) > 0 {
		strat.KeepCVX = helpers.DecodeBigInt(rawKeepCVX)
	}
	if len(rawName) > 0 {
		strat.Name = helpers.DecodeString(rawName)
	}
	if len(rawDoHealthCheck) > 0 {
		strat.DoHealthCheck = helpers.DecodeBool(rawDoHealthCheck)
	}
	if len(rawIsShutdown) > 0 {
		strat.IsActive = helpers.DecodeBool(rawIsShutdown)
	}
	if !strat.IsActive || helpers.DecodeBool(rawIsShutdown) {
		strat.IsRetired = true
	}

	return strat
}
