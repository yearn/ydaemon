package vaults

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/strategies"
)

var metaVaultFileErrorAlreadySent = make(map[uint64]map[common.Address]bool)
var metaVaultTokenErrorAlreadySent = make(map[uint64]map[common.Address]bool)

/**************************************************************************************************
** fetchBasicInformations will, for a list of addresses, fetch all the relevant basic information
** for the related vaults.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: a list of addresses of the vaults we want to fetch the information for
**
** Returns:
** - a list of TVault containing the basic information for the vaults
**************************************************************************************************/
func fetchBasicInformations(
	chainID uint64,
	vaults []common.Address,
) (vaultList []models.TVault) {
	if metaVaultFileErrorAlreadySent[chainID] == nil {
		metaVaultFileErrorAlreadySent[chainID] = make(map[common.Address]bool)
		metaVaultTokenErrorAlreadySent[chainID] = make(map[common.Address]bool)
	}

	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `vaultAddress + methodName`.
	**********************************************************************************************/
	maxStrategiesPerVault := 20
	calls := []ethereum.Call{}
	for _, vault := range vaults {
		calls = append(calls, multicalls.GetPricePerShare(vault.Hex(), vault))
		calls = append(calls, multicalls.GetAPIVersion(vault.Hex(), vault))
		calls = append(calls, multicalls.GetPerformanceFee(vault.Hex(), vault))
		calls = append(calls, multicalls.GetManagementFee(vault.Hex(), vault))
		calls = append(calls, multicalls.GetToken(vault.Hex(), vault))
		calls = append(calls, multicalls.GetActivation(vault.Hex(), vault))
		calls = append(calls, multicalls.GetEmergencyShutdown(vault.Hex(), vault))
		calls = append(calls, multicalls.GetGovernance(vault.Hex(), vault))
		calls = append(calls, multicalls.GetGuardian(vault.Hex(), vault))
		calls = append(calls, multicalls.GetManagement(vault.Hex(), vault))
		calls = append(calls, multicalls.GetRewards(vault.Hex(), vault))
		calls = append(calls, multicalls.GetTotalAssets(vault.Hex(), vault))
		calls = append(calls, multicalls.GetDepositLimit(vault.Hex(), vault))
		calls = append(calls, multicalls.GetAvailableDepositLimit(vault.Hex(), vault))
		for i := 0; i < maxStrategiesPerVault; i++ {
			calls = append(calls, multicalls.GetVaultWithdrawalQueue(vault.Hex(), int64(i), vault))
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, vault := range vaults {
		rawPricePerShare := response[vault.Hex()+`pricePerShare`]
		rawApiVersion := response[vault.Hex()+`apiVersion`]
		rawPerformanceFee := response[vault.Hex()+`performanceFee`]
		rawManagementFee := response[vault.Hex()+`managementFee`]
		rawUnderlying := response[vault.Hex()+`token`]
		rawEmergencyShutdown := response[vault.Hex()+`emergencyShutdown`]
		rawActivation := response[vault.Hex()+`activation`]
		rawGovernance := response[vault.Hex()+`governance`]
		rawGuardian := response[vault.Hex()+`guardian`]
		rawManagement := response[vault.Hex()+`management`]
		rawRewards := response[vault.Hex()+`rewards`]
		rawTotalAssets := response[vault.Hex()+`totalAssets`]
		rawDepositLimit := response[vault.Hex()+`depositLimit`]
		rawAvailableDepositLimit := response[vault.Hex()+`availableDepositLimit`]

		shareTokenData, ok := store.GetERC20(chainID, vault)
		if !ok {
			shareTokenData = models.TERC20Token{}
			logs.Warning(`impossible to retrieve share token for vault ` + vault.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10))
		}

		underlyingTokenData, ok := store.GetERC20(chainID, helpers.DecodeAddress(rawUnderlying))
		if !ok {
			tokenAddress := helpers.DecodeAddress(rawUnderlying)
			underlyingTokenData = models.TERC20Token{}
			if !metaVaultTokenErrorAlreadySent[chainID][tokenAddress] {
				logs.Warning(`impossible to retrieve underlying token for vault ` + vault.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10))
				metaVaultTokenErrorAlreadySent[chainID][tokenAddress] = true
			}
		}

		vaultData, ok := meta.GetMetaVault(chainID, vault)
		if !ok {
			if !metaVaultFileErrorAlreadySent[chainID][vault] {
				logs.Warning(`impossible to retrieve meta file for vault ` + vault.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10))
				metaVaultFileErrorAlreadySent[chainID][vault] = true
			}
		}

		/******************************************************************************************
		** Preparing our new TVault object
		******************************************************************************************/
		newVault := models.TVault{
			ChainID:               chainID,
			Address:               vault,
			Name:                  shareTokenData.Name,
			Symbol:                shareTokenData.Symbol,
			Decimals:              shareTokenData.Decimals,
			Icon:                  env.BASE_ASSET_URL + strconv.FormatUint(chainID, 10) + `/` + vault.Hex() + `/logo-128.png`,
			Type:                  models.TokenTypeStandardVault,
			Endorsed:              false, //Default to false, will be updated later
			Inception:             helpers.DecodeBigInt(rawActivation).Uint64(),
			Version:               helpers.DecodeString(rawApiVersion),
			PricePerShare:         helpers.DecodeBigInt(rawPricePerShare),
			Management:            helpers.DecodeAddress(rawManagement),
			Governance:            helpers.DecodeAddress(rawGovernance),
			Guardian:              helpers.DecodeAddress(rawGuardian),
			Rewards:               helpers.DecodeAddress(rawRewards),
			TotalAssets:           helpers.DecodeBigInt(rawTotalAssets),
			DepositLimit:          helpers.DecodeBigInt(rawDepositLimit),
			AvailableDepositLimit: helpers.DecodeBigInt(rawAvailableDepositLimit),
			PerformanceFee:        helpers.DecodeBigInt(rawPerformanceFee).Uint64(),
			ManagementFee:         helpers.DecodeBigInt(rawManagementFee).Uint64(),
			EmergencyShutdown:     helpers.DecodeBool(rawEmergencyShutdown),
			Token: models.TERC20Token{
				Address:       underlyingTokenData.Address,
				Name:          underlyingTokenData.Name,
				DisplayName:   underlyingTokenData.DisplayName,
				DisplaySymbol: underlyingTokenData.DisplaySymbol,
				Symbol:        underlyingTokenData.Symbol,
				Description:   underlyingTokenData.Description,
				Decimals:      underlyingTokenData.Decimals,
				Icon:          underlyingTokenData.Icon,
				ChainID:       underlyingTokenData.ChainID,
			},
		}

		if vaultData != nil && vaultData.DisplayName != `` {
			newVault = BuildNames(newVault, vaultData.DisplayName)
		} else {
			newVault = BuildNames(newVault, shareTokenData.DisplayName)
		}
		newVault = BuildSymbol(newVault, shareTokenData.DisplaySymbol)

		/******************************************************************************************
		** Adding the withdrawal queue stuffs
		******************************************************************************************/
		withdrawQueueForStrategies := []common.Address{}
		for i := 0; i < maxStrategiesPerVault; i++ {
			result := response[vault.Hex()+strconv.FormatInt(int64(i), 10)+`withdrawalQueue`]
			if len(result) == 1 {
				strategyAddress := helpers.DecodeAddress(result)
				if helpers.AddressIsValid(strategyAddress, chainID) {
					withdrawQueueForStrategies = append(withdrawQueueForStrategies, strategyAddress)
				}
			}
		}
		newVault.WithdrawalQueue = withdrawQueueForStrategies
		strategies.SetInStrategiesWithdrawalQueue(chainID, vault, withdrawQueueForStrategies)
		vaultList = append(vaultList, newVault)
	}

	return vaultList
}

/**************************************************************************************************
** findAllVaults is simply a wrapper around our fetchBasicInformations function to make it easier
** to read. It will take the vaultMap, extract the individual addresses, and then call the
** fetchBasicInformations function to get the vault information, before assigning them to a new
** map.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultMap: our map of vaultAddress -> TVault
**
** Returns:
** - a map of vaultAddress -> TVault
**************************************************************************************************/
func findAllVaults(
	chainID uint64,
	vaultMap map[common.Address]models.TVault,
) map[common.Address]models.TVault {
	newMap := make(map[common.Address]models.TVault)
	vaultListAddresses := []common.Address{}
	for vaultAddress := range vaultMap {
		vaultListAddresses = append(vaultListAddresses, vaultAddress)
	}

	newVaultMap := fetchBasicInformations(chainID, vaultListAddresses)
	for _, vault := range newVaultMap {
		vault.Endorsed = vaultMap[vault.Address].Endorsed
		vault.Type = vaultMap[vault.Address].Type
		newMap[vault.Address] = vault
	}

	return newMap
}

/**************************************************************************************************
** The base of Yearn are the vaults. They are the smart contracts that are used to manage the
** deposits and the withdrawals of the users.
** The goal of this function is to get a list of all the Vaults living in the Yearn's Ecosystem.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**
** Returns:
** - a map of vaultAddress -> TVault
**************************************************************************************************/
func RetrieveAllVaults(
	chainID uint64,
	vaults map[common.Address]models.TVaultsFromRegistry,
) map[common.Address]models.TVault {
	/**********************************************************************************************
	** First, try to retrieve the list of vaults from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	vaultMap, _ := store.ListAllVaults(chainID)

	/**********************************************************************************************
	** From the vault registry we have the first batch of vaults. In order to proceed, we will
	** create a map of TVault which will be stored in the DB. This action should only be performed
	** for every refresh period, per vault per chainID.
	** TODO: Optimize this part to split possible changes (pricePerShare) from almost impossible
	** changes (name, symbol)
	**********************************************************************************************/
	updatedVaultMap := make(map[common.Address]models.TVault)
	for _, currentVault := range vaults {
		updatedVaultMap[currentVault.Address] = models.TVault{
			Address:    currentVault.TokenAddress,
			Endorsed:   (currentVault.Type == models.TokenTypeStandardVault || currentVault.Type == models.TokenTypeAutomatedVault) && currentVault.TokenAddress != common.Address{},
			Type:       currentVault.Type,
			Activation: currentVault.Activation,
		}
	}

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the vault data for them.
	** We will add them manually here.
	**********************************************************************************************/
	for _, vault := range env.CHAINS[chainID].ExtraVaults {
		vaultAddress := vault.Address
		if _, ok := vaultMap[vaultAddress]; !ok {
			updatedVaultMap[vaultAddress] = models.TVault{Address: vaultAddress}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations.
	**********************************************************************************************/
	updatedVaultMap = findAllVaults(chainID, updatedVaultMap)

	/**********************************************************************************************
	** Once everything is setup, we will store each token in the DB. The storage is set as a map
	** of vaultAddress -> TTokens. All vaults will be retrievable from the store.Interate() func.
	**********************************************************************************************/
	for _, vault := range updatedVaultMap {
		store.AppendToVaultMap(chainID, vault)
		if _, ok := vaultMap[vault.Address]; !ok {
			store.StoreVault(chainID, vault)
		}
	}
	vaultMap, _ = store.ListAllVaults(chainID)

	_vaultMap[chainID] = vaultMap
	return vaultMap
}
