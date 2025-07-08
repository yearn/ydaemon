package fetcher

import (
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** getHumanizedTokenPrice retrieves and formats a token's price in human-readable form.
**
** This function fetches the price of a token from storage and returns it in two formats:
** 1. As a bigNumber.Float object for precise mathematical operations
** 2. As a float64 for easy display and simple calculations
**
** Token prices in the system are stored in USDC terms (6 decimal places). This function
** handles the necessary conversion to make prices human-readable and usable for calculations.
**
** @param chainID uint64 - The blockchain network ID
** @param tokenAddress common.Address - The address of the token to get the price for
** @return *bigNumber.Float - The humanized price as a BigNumber for precise operations
** @return float64 - The humanized price as a regular float for display purposes
**************************************************************************************************/
func getHumanizedTokenPrice(chainID uint64, tokenAddress common.Address) (*bigNumber.Float, float64) {
	price, ok := storage.GetPrice(chainID, tokenAddress)
	if ok {
		fHumanizedPrice, _ := price.HumanizedPrice.Float64()
		return price.HumanizedPrice, fHumanizedPrice
	}
	return bigNumber.NewFloat(), 0.0
}

/**************************************************************************************************
** getHumanizedValue calculates the USD value of a token amount.
**
** This function converts a raw token balance into its equivalent USD value by:
** 1. Converting the raw token amount to a human-readable form using the token's decimals
** 2. Multiplying the human-readable token amount by its USD price
**
** This is essential for calculating TVL (Total Value Locked) and other USD-denominated metrics
** across different tokens with varying decimal precisions.
**
** @param balanceToken *bigNumber.Int - The raw token balance
** @param decimals int - The number of decimal places for the token
** @param humanizedPrice *bigNumber.Float - The humanized price of the token in USD
** @return float64 - The USD value of the token amount
**************************************************************************************************/
func getHumanizedValue(balanceToken *bigNumber.Int, decimals int, humanizedPrice *bigNumber.Float) float64 {
	_, humanizedValue := helpers.FormatAmount(balanceToken.String(), decimals)
	fHumanizedValue, _ := bigNumber.NewFloat().Mul(humanizedValue, humanizedPrice).Float64()
	return fHumanizedValue
}

/**************************************************************************************************
** getV3VaultCalls prepares multicall requests for fetching data from V3 vaults.
**
** This function builds an array of Ethereum calls to retrieve various metrics and properties
** from V3 vault contracts. The information retrieved follows a tiered approach based on time
** since last update:
**
** 1. Core metrics (every request):
**    - Price per share (with convertToAssets fallback for proper decimal handling)
**    - Decimals (for accurate value calculations)
**    - Total assets (for TVL calculations)
**    - Default queue (list of active strategies)
**    - API version (for compatibility checks)
**    - Shutdown status (V3 equivalent of emergency shutdown)
**
** 2. Daily updates (if more than 24 hours since last update or forced refresh):
**    - Asset (underlying token address)
**    - Role manager address (for permission checks)
**    - Accountant address (for fee configurations)
**    - Vault-kind specific fee information:
**      - For multi-strategy vaults: default fee config from accountant
**      - For single-strategy vaults: performance fee
**
** The function handles special vault kinds differently, particularly for fee retrieval.
** For multi-strategy vaults, it attempts to get the accountant address directly if not
** available in storage, to ensure fee data can be retrieved.
**
** @param vault models.TVault - The vault to build calls for
** @return []ethereum.Call - Array of Ethereum calls to be executed in a multicall
**************************************************************************************************/
func getV3VaultCalls(vault models.TVault) []ethereum.Call {
	metadata := storage.GetVaultsJsonMetadata(vault.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh
	calls := []ethereum.Call{}

	asEthers := big.NewInt(10)
	asEthers = asEthers.Exp(asEthers, big.NewInt(18), nil)

	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetPricePerShare(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetConvertPricePerShare(vault.Address.Hex(), vault.Address, asEthers))
	calls = append(calls, multicalls.GetDecimals(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetTotalAssets(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetDefaultQueue(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetAPIVersion(vault.Address.Hex(), vault.Address))
	calls = append(calls, multicalls.GetIsShutdown(vault.Address.Hex(), vault.Address, ``))

	if time.Since(lastUpdate).Hours() > 24 || shouldRefresh {
		// If the last vault update was more than 24 hour ago, we will do a full update
		calls = append(calls, multicalls.GetAsset(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetRoleManager(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetAccountant(vault.Address.Hex(), vault.Address))

		switch vault.Kind {
		case models.VaultKindMultiple:
			// If the vault is a multi strategy vault, we need to get the default fees from the accountant
			existingVault, ok := storage.GetVault(vault.ChainID, vault.Address)
			if ok && (existingVault.Accountant != nil) && (existingVault.Accountant.Hex() != common.Address{}.Hex()) {
				calls = append(calls, multicalls.GetDefaultFeeConfig(vault.Address.Hex(), *existingVault.Accountant))
			} else {
				//get accountant now
				client, err := ethclient.Dial(ethereum.GetRPCURI(vault.ChainID))
				if err == nil {
					vaultContract, err := contracts.NewYvault300(vault.Address, client)
					if err == nil {
						accountant, err := vaultContract.Accountant(nil)
						if err == nil {
							calls = append(calls, multicalls.GetDefaultFeeConfig(vault.Address.Hex(), accountant))
						}
					}
				}

			}
		case models.VaultKindSingle:
			calls = append(calls, multicalls.GetPerformanceFee(vault.Address.Hex(), vault.Address))
		}
	}
	return calls
}

/**************************************************************************************************
** getV2VaultCalls prepares multicall requests for fetching data from V2 and earlier vaults.
**
** This function builds an array of Ethereum calls to retrieve various metrics and properties
** from vault contracts. The information retrieved follows a tiered approach based on time
** since last update:
**
** 1. Core metrics (every request):
**    - Price per share (for APY calculations and share value)
**    - Total assets (for TVL and allocation calculations)
**    - Withdrawal queue (strategy addresses in order of withdrawal priority)
**
** 2. Hourly updates (if more than 1 hour since last update or forced refresh):
**    - Performance fee (for yield calculations)
**    - Management fee (for yield calculations)
**    - Emergency shutdown status (for vault availability)
**
** 3. Daily updates (if more than 24 hours since last update or forced refresh):
**    - API version (for compatibility checks)
**    - Token (underlying asset address)
**    - Accountant address (for fee configuration)
**
** @param vault models.TVault - The vault to build calls for
** @return []ethereum.Call - Array of Ethereum calls to be executed in a multicall
**************************************************************************************************/
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
		calls = append(calls, multicalls.GetToken(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetAccountant(vault.Address.Hex(), vault.Address))
	}
	return calls
}

/**************************************************************************************************
** getV2StrategyCalls prepares multicall requests for fetching data from V2 strategies.
**
** This function builds an array of Ethereum calls to retrieve various metrics and properties
** from V2 strategy contracts. The information retrieved varies based on the time since the
** last update, following a tiered approach:
**
** 1. Core metrics (every request):
**    - Retrieves basic strategy information from the vault
**    - Checks if the strategy is active
**
** 2. Hourly updates (if more than 1 hour since last update or forced refresh):
**    - Retrieves CRV-related settings (keepCRV, keepCRVPercent)
**    - Retrieves CVX-related settings (keepCVX)
**    - Checks emergency exit status
**
** 3. Daily updates (if more than 24 hours since last update or forced refresh):
**    - Retrieves strategy name
**    - Checks health check configuration
**
** @param strat models.TStrategy - The strategy to build calls for
** @return []ethereum.Call - Array of Ethereum calls to be executed in a multicall
**************************************************************************************************/
func getV2StrategyCalls(strat models.TStrategy) []ethereum.Call {
	metadata := storage.GetStrategiesJsonMetadata(strat.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh
	strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()

	calls := []ethereum.Call{}
	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetStrategies(strategyKey, strat.VaultAddress, strat.Address, strat.VaultVersion))
	calls = append(calls, multicalls.GetStategyIsActive(strategyKey, strat.Address, strat.VaultVersion))
	if time.Since(lastUpdate).Hours() > 1 || shouldRefresh {
		// If the last strat update was more than 1 hour ago, we will do a partial update
		calls = append(calls, multicalls.GetStategyKeepCRV(strategyKey, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCRVPercent(strategyKey, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCVX(strategyKey, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetEmergencyExit(strategyKey, strat.Address, strat.VaultVersion))
	}
	// Always fetch strategy name and health check
	calls = append(calls, multicalls.GetStrategyName(strategyKey, strat.Address, strat.VaultVersion))
	calls = append(calls, multicalls.GetDoHealthCheck(strategyKey, strat.Address, strat.VaultVersion))
	return calls
}

/**************************************************************************************************
** getV3StrategyCalls prepares multicall requests for fetching data from V3 strategies.
**
** This function builds an array of Ethereum calls to retrieve various metrics and properties
** from V3 strategy contracts. V3 strategies have a different interface and data structure
** compared to V2, requiring specialized call patterns.
**
** The information retrieved follows a tiered approach based on update frequency:
**
** 1. Core metrics (every request):
**    - Retrieves strategy information from the vault
**    - Gets performance fee
**    - Fetches total assets from the vault
**    - Checks if the strategy is shut down
**
** 2. Hourly updates (if more than 1 hour since last update or forced refresh):
**    - Retrieves CRV-related settings (keepCRV, keepCRVPercent)
**    - Retrieves CVX-related settings (keepCVX)
**
** 3. Daily updates (if more than 24 hours since last update or forced refresh):
**    - Retrieves strategy name
**    - Checks health check configuration
**
** @param strat models.TStrategy - The strategy to build calls for
** @return []ethereum.Call - Array of Ethereum calls to be executed in a multicall
**************************************************************************************************/
func getV3StrategyCalls(strat models.TStrategy) []ethereum.Call {
	metadata := storage.GetStrategiesJsonMetadata(strat.ChainID)
	lastUpdate := metadata.LastUpdate
	shouldRefresh := metadata.ShouldRefresh
	strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()

	calls := []ethereum.Call{}
	//For every loop we need at least to update theses
	calls = append(calls, multicalls.GetV3Strategies(strategyKey, strat.VaultAddress, strat.Address, strat.VaultVersion))
	calls = append(calls, multicalls.GetPerformanceFee(strategyKey, strat.Address))
	calls = append(calls, multicalls.GetTotalAssets(strat.VaultAddress.Hex(), strat.VaultAddress))
	calls = append(calls, multicalls.GetIsShutdown(strategyKey, strat.Address, strat.VaultVersion))
	if time.Since(lastUpdate).Hours() > 1 || shouldRefresh {
		// If the last strat update was more than 1 hour ago, we will do a partial update
		calls = append(calls, multicalls.GetStategyKeepCRV(strategyKey, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCRVPercent(strategyKey, strat.Address, strat.VaultVersion))
		calls = append(calls, multicalls.GetStategyKeepCVX(strategyKey, strat.Address, strat.VaultVersion))
	}
	// Always fetch strategy name and health check
	calls = append(calls, multicalls.GetStrategyName(strategyKey, strat.Address, strat.VaultVersion))
	calls = append(calls, multicalls.GetDoHealthCheck(strategyKey, strat.Address, strat.VaultVersion))
	return calls
}

/**************************************************************************************************
** handleV2VaultCalls processes the multicall responses for V2 and earlier vaults.
**
** This function takes the raw responses from the multicall and updates the vault model with
** the extracted data. It processes essential vault metrics including:
**
** 1. Financial data:
**    - Price per share (for user share value calculations)
**    - Total assets (for TVL calculations)
**    - Performance and management fees (for yield projections)
**
** 2. Operational data:
**    - Emergency shutdown status (for vault availability)
**    - API version (for compatibility checks)
**    - Underlying asset address (for token identification)
**
** 3. Strategy data:
**    - Withdrawal queue (list of active strategies in priority order)
**
** The function handles response decoding with appropriate type conversions and safely
** processes each field only if data is available in the response.
**
** @param vault models.TVault - The vault model to update
** @param response map[string][]interface{} - The multicall responses keyed by call identifiers
** @return models.TVault - The updated vault with values from the response
**************************************************************************************************/
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

/**************************************************************************************************
** handleV3VaultCalls processes the multicall responses for V3 vaults.
**
** This function extracts and transforms data from multicall responses to update a V3 vault model.
** V3 vaults have a different structure and behavior compared to earlier versions, requiring
** specialized data processing.
**
** Key operations include:
** 1. Processing price per share with fallback logic:
**    - Primarily uses direct pricePerShare value if available
**    - Falls back to converted asset value with proper decimal adjustment if needed
**
** 2. Extracting core vault metrics:
**    - Total assets (for TVL calculations)
**    - Active strategies list (from the default queue)
**    - Emergency shutdown status (as isShutdown in V3)
**    - Underlying asset address
**    - API version
**    - Accountant address for fees
**
** 3. Fee processing based on vault kind:
**    - Multiple-strategy vaults: extracts fees from the accountant contract
**    - Single-strategy vaults: uses direct performance fee from the vault
**
** The price per share calculation is particularly complex, handling decimal normalization
** between different token decimal standards and the 18-decimal standard used internally.
**
** @param vault models.TVault - The vault model to update
** @param response map[string][]interface{} - The multicall responses keyed by call identifiers
** @return models.TVault - The updated vault with values from the response
**************************************************************************************************/
func handleV3VaultCalls(vault models.TVault, response map[string][]interface{}) models.TVault {
	rawPricePerShare := response[vault.Address.Hex()+`pricePerShare`]
	rawConvertPricePerShare := response[vault.Address.Hex()+`convertToAssets`]
	rawTotalAssets := response[vault.Address.Hex()+`totalAssets`]
	rawDefaultQueue := response[vault.Address.Hex()+`get_default_queue`]
	rawShutdown := response[vault.Address.Hex()+`isShutdown`]
	rawUnderlying := response[vault.Address.Hex()+`asset`]
	rawApiVersion := response[vault.Address.Hex()+`apiVersion`]
	rawAccountant := response[vault.Address.Hex()+`accountant`]
	rawDecimals := response[vault.Address.Hex()+`decimals`]

	vault.LastPricePerShare = helpers.DecodeBigInt(rawPricePerShare)
	if vault.LastPricePerShare.IsZero() {
		//Try to use rawConvertPricePerShare. However, rawConvertPricePerShare is on base 10^18 while the token decimals are on base 10^decimals.
		//We need to convert rawConvertPricePerShare to the same base as the token decimals
		decimals := helpers.DecodeUint64(rawDecimals)
		converted := helpers.DecodeBigInt(rawConvertPricePerShare)

		//Assuming decimals is 6, the code below is doing `convertedValue * 10^6 / 10^18`
		vault.LastPricePerShare = bigNumber.NewInt(0).Div(
			bigNumber.NewInt(0).Mul(
				converted,
				bigNumber.NewInt(10).Exp(
					bigNumber.NewInt(10),
					bigNumber.NewInt(int64(decimals)),
					nil,
				),
			),
			bigNumber.NewInt(10).Exp(
				bigNumber.NewInt(10),
				bigNumber.NewInt(18),
				nil,
			),
		)
	}

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
		accountant := helpers.DecodeAddress(rawAccountant)
		vault.Accountant = &accountant
	}

	switch vault.Kind {
	case models.VaultKindMultiple:
		rawDefaultFees := response[vault.Address.Hex()+`defaultConfig`]
		if len(rawDefaultFees) >= 2 {
			defaultFees := helpers.DecodeUint16s(rawDefaultFees)
			bigMngFee := uint64(defaultFees[0])
			bigPerfFee := uint64(defaultFees[1])
			vault.PerformanceFee = bigPerfFee
			vault.ManagementFee = bigMngFee
		}
	case models.VaultKindSingle:
		rawPerformanceFee := response[vault.Address.Hex()+`performanceFee`]
		if len(rawPerformanceFee) > 0 {
			vault.PerformanceFee = helpers.DecodeBigInt(rawPerformanceFee).Uint64()
		}
	}
	return vault
}

/**************************************************************************************************
** handleV2StrategyCalls processes the multicall responses for V2 strategies.
**
** This function takes the raw responses from the multicall and maps them to the appropriate
** fields in the strategy model. The response processing is version-aware, handling different
** data structures depending on the vault version (0.2.2, 0.3.0, 0.3.1, or later).
**
** Key data processed includes:
** - Financial metrics: total debt, total gain, total loss, performance fee
** - Operational settings: keepCRV, keepCRVPercent, keepCVX
** - Status indicators: isActive, isRetired, emergencyExit
** - Configuration data: name, doHealthCheck
**
** The function handles version-specific differences in the strategies response array:
** - 0.2.2: 8-element array with different field positions
** - 0.3.0/0.3.1: 8-element array with debt ratio at position 2
** - Later versions: 9-element array with slightly different structure
**
** @param strat models.TStrategy - The strategy model to update with response data
** @param response map[string][]interface{} - The multicall responses
** @return models.TStrategy - The updated strategy with values from the response
**************************************************************************************************/
func handleV2StrategyCalls(strat models.TStrategy, response map[string][]interface{}) models.TStrategy {
	strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()
	rawStrategies := response[strategyKey+`strategies`]
	rawIsActive := response[strategyKey+`isActive`]
	rawKeepCRV := response[strategyKey+`keepCRV`]
	rawKeepCRVPercent := response[strategyKey+`keepCrvPercent`]
	rawKeepCVX := response[strategyKey+`keepCVX`]
	rawName := response[strategyKey+`name`]
	rawDoHealthCheck := response[strategyKey+`doHealthCheck`]
	rawEmergencyExit := response[strategyKey+`emergencyExit`]

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
		// if !strat.IsActive {
		// 	strat.IsRetired = true
		// }
	}
	if len(rawDoHealthCheck) > 0 {
		strat.DoHealthCheck = helpers.DecodeBool(rawDoHealthCheck)
	}
	if len(rawEmergencyExit) > 0 {
		isEmergencyExit := helpers.DecodeBool(rawEmergencyExit)
		if isEmergencyExit {
			strat.IsRetired = true
		}
	}

	return strat
}

/**************************************************************************************************
** handleV3StrategyCalls processes the multicall responses for V3 strategies.
**
** This function takes the raw responses from the multicall and maps them to the appropriate
** fields in the strategy model for Yearn V3 vaults. V3 strategies have a different structure
** and behavior compared to earlier versions.
**
** Key operations:
** 1. Processes core strategy data (debt, activation time, last report)
** 2. Sets total gain and loss fields (note: these are not directly available in V3)
** 3. Calculates debt ratio as a percentage of vault's total assets
** 4. Processes operational settings (keepCRV, keepCRVPercent, keepCVX)
** 5. Updates status flags (isActive, isRetired) based on the shutdown state
**
** The debt ratio calculation is particularly important as it:
** - Converts raw debt values to a percentage (0-10000, where 10000 = 100%)
** - Uses mathematical rounding to ensure integer precision
** - Properly handles cases where vault total assets might be zero
**
** @param strat models.TStrategy - The strategy model to update with response data
** @param response map[string][]interface{} - The multicall responses
** @return models.TStrategy - The updated strategy with values from the response
**************************************************************************************************/
func handleV3StrategyCalls(strat models.TStrategy, response map[string][]interface{}) models.TStrategy {
	type typeOfRawStrategies = struct {
		Activation  *big.Int "json:\"activation\""
		LastReport  *big.Int "json:\"last_report\""
		CurrentDebt *big.Int "json:\"current_debt\""
		MaxDebt     *big.Int "json:\"max_debt\""
	}

	strategyKey := strat.Address.Hex() + `_` + strat.VaultAddress.Hex()
	rawStrategies := response[strategyKey+`strategies`]
	rawKeepCRV := response[strategyKey+`keepCRV`]
	rawKeepCRVPercent := response[strategyKey+`keepCrvPercent`]
	rawKeepCVX := response[strategyKey+`keepCVX`]
	rawName := response[strategyKey+`name`]
	rawVaultTotalAssets := response[strat.VaultAddress.Hex()+`totalAssets`]
	rawDoHealthCheck := response[strategyKey+`doHealthCheck`]
	rawIsShutdown := response[strategyKey+`isShutdown`]
	rawPerformanceFee := response[strategyKey+`performanceFee`]

	if (len(rawPerformanceFee) > 0) && (len(rawStrategies) > 0) {
		strat.LastPerformanceFee = helpers.DecodeBigInt(rawPerformanceFee)
	} else {
		strat.LastPerformanceFee = bigNumber.NewInt(0) // Default to 1000, aka 10%
	}
	if len(rawStrategies) > 0 {
		strat.LastTotalDebt = bigNumber.SetInt(rawStrategies[0].(typeOfRawStrategies).CurrentDebt)
		strat.TimeActivated = bigNumber.SetInt(rawStrategies[0].(typeOfRawStrategies).Activation)
		strat.LastReport = bigNumber.SetInt(rawStrategies[0].(typeOfRawStrategies).LastReport)
	}
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
		strat.IsActive = !helpers.DecodeBool(rawIsShutdown)
		if !strat.IsActive {
			strat.IsRetired = true
		} else {
			strat.IsRetired = false
		}
	}

	return strat
}

type TProcessNewVaultMethod string

const (
	ProcessNewVaultMethodAppend  TProcessNewVaultMethod = "append"
	ProcessNewVaultMethodReplace TProcessNewVaultMethod = "replace"
)
