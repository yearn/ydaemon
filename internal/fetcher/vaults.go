package fetcher

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** fetchVaultsBasicInformations will, for a list of addresses, fetch all the relevant basic information
** for the related vaults.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: a list of addresses of the vaults we want to fetch the information for
**
** Returns:
** - a list of TVault containing the basic information for the vaults
**************************************************************************************************/
func fetchVaultsBasicInformations(
	chainID uint64,
	vaultMap map[common.Address]models.TVault,
) (vaultList []models.TVault) {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `vaultAddress + methodName`.
	**********************************************************************************************/
	maxStrategiesPerVault := 20
	calls := []ethereum.Call{}
	for _, vault := range vaultMap {
		if vault.IsRetired {
			continue
		}
		if time.Since(vault.LastUpdate).Hours() > 24 || true {
			// If the last vault update was more than 24 hour ago, we will do a full update
			calls = append(calls, multicalls.GetPricePerShare(vault.Address.String(), vault.Address))
			calls = append(calls, multicalls.GetAPIVersion(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetPerformanceFee(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetManagementFee(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetToken(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetEmergencyShutdown(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetGovernance(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetGuardian(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetManagement(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetRewards(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetTotalAssets(vault.Address.Hex(), vault.Address))
			for i := 0; i < maxStrategiesPerVault; i++ {
				calls = append(calls, multicalls.GetVaultWithdrawalQueue(vault.Address.Hex(), int64(i), vault.Address))
			}
		} else if time.Since(vault.LastUpdate).Hours() > 1 {
			// If the last vault update was more than 1 hour ago, we will do a partial update
			calls = append(calls, multicalls.GetPerformanceFee(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetManagementFee(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetEmergencyShutdown(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetPricePerShare(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetTotalAssets(vault.Address.Hex(), vault.Address))
			for i := 0; i < maxStrategiesPerVault; i++ {
				calls = append(calls, multicalls.GetVaultWithdrawalQueue(vault.Address.Hex(), int64(i), vault.Address))
			}
		} else {
			// Else, for every loop, do at least this
			calls = append(calls, multicalls.GetPricePerShare(vault.Address.Hex(), vault.Address))
			calls = append(calls, multicalls.GetTotalAssets(vault.Address.Hex(), vault.Address))
			for i := 0; i < maxStrategiesPerVault; i++ {
				calls = append(calls, multicalls.GetVaultWithdrawalQueue(vault.Address.Hex(), int64(i), vault.Address))
			}
		}
	}

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := multicalls.Perform(chainID, calls, nil)
	for _, vault := range vaultMap {
		if vault.IsRetired {
			continue
		}
		rawPricePerShare := response[vault.Address.Hex()+`pricePerShare`]
		rawApiVersion := response[vault.Address.Hex()+`apiVersion`]
		rawPerformanceFee := response[vault.Address.Hex()+`performanceFee`]
		rawManagementFee := response[vault.Address.Hex()+`managementFee`]
		rawUnderlying := response[vault.Address.Hex()+`token`]
		rawEmergencyShutdown := response[vault.Address.Hex()+`emergencyShutdown`]
		rawGovernance := response[vault.Address.Hex()+`governance`]
		rawGuardian := response[vault.Address.Hex()+`guardian`]
		rawManagement := response[vault.Address.Hex()+`management`]
		rawRewards := response[vault.Address.Hex()+`rewards`]
		rawTotalAssets := response[vault.Address.Hex()+`totalAssets`]

		/******************************************************************************************
		** Preparing our new TVault object
		******************************************************************************************/
		newVault := vault
		newVault.LastPricePerShare = helpers.DecodeBigInt(rawPricePerShare)
		newVault.LastTotalAssets = helpers.DecodeBigInt(rawTotalAssets)
		newVault.LastUpdate = time.Now()

		if len(rawManagement) > 0 {
			newVault.Management = helpers.DecodeAddress(rawManagement)
		}
		if len(rawGovernance) > 0 {
			newVault.Governance = helpers.DecodeAddress(rawGovernance)
		}
		if len(rawGuardian) > 0 {
			newVault.Guardian = helpers.DecodeAddress(rawGuardian)
		}
		if len(rawRewards) > 0 {
			newVault.Rewards = helpers.DecodeAddress(rawRewards)
		}
		if len(rawPerformanceFee) > 0 {
			newVault.PerformanceFee = helpers.DecodeBigInt(rawPerformanceFee).Uint64()
		}
		if len(rawManagementFee) > 0 {
			newVault.ManagementFee = helpers.DecodeBigInt(rawManagementFee).Uint64()
		}
		if len(rawEmergencyShutdown) > 0 {
			newVault.EmergencyShutdown = helpers.DecodeBool(rawEmergencyShutdown)
		}
		if len(rawUnderlying) > 0 {
			newVault.AssetAddress = helpers.DecodeAddress(rawUnderlying)
		}
		if len(rawApiVersion) > 0 {
			newVault.Version = helpers.DecodeString(rawApiVersion)
		}

		/******************************************************************************************
		** Adding the withdrawal queue stuffs
		******************************************************************************************/
		withdrawQueueForStrategies := []common.Address{}
		for i := 0; i < maxStrategiesPerVault; i++ {
			result := response[vault.Address.Hex()+strconv.FormatInt(int64(i), 10)+`withdrawalQueue`]
			if len(result) == 1 {
				strategyAddress := helpers.DecodeAddress(result)
				if helpers.AddressIsValid(strategyAddress, chainID) {
					withdrawQueueForStrategies = append(withdrawQueueForStrategies, strategyAddress)
				}
			}
		}
		newVault.LastActiveStrategies = withdrawQueueForStrategies
		vaultList = append(vaultList, newVault)
	}

	return vaultList
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
	** First, try to retrieve the list of vaults from the database and populate our updatedVaultMap
	** with it.
	**********************************************************************************************/
	vaultMap, _ := storage.ListVaults(chainID)
	updatedVaultMap := vaultMap

	/**********************************************************************************************
	** From the vault registry we have the first batch of vaults. In order to proceed, we will
	** create a map of TVault which will be stored in the DB. This action should only be performed
	** for every refresh period, per vault per chainID.
	**********************************************************************************************/
	for _, currentVault := range vaults {
		if _, ok := vaultMap[currentVault.Address]; !ok {
			if (currentVault.TokenAddress == common.Address{}) {
				continue
			}
			isEndorsed := (currentVault.Type == models.TokenTypeStandardVault || currentVault.Type == models.TokenTypeAutomatedVault)
			updatedVaultMap[currentVault.Address] = models.TVault{
				Address:    currentVault.Address,
				Endorsed:   isEndorsed,
				Type:       currentVault.Type,
				Activation: currentVault.BlockNumber,
			}
		}
	}

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the vault data for them.
	** We will add them manually here.
	**********************************************************************************************/
	for _, vault := range env.CHAINS[chainID].ExtraVaults {
		if _, ok := vaultMap[vault.Address]; !ok {
			if _, ok := updatedVaultMap[vault.Address]; !ok {
				updatedVaultMap[vault.Address] = models.TVault{Address: vault.Address}
			}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations.
	**********************************************************************************************/
	newVaultMap := fetchVaultsBasicInformations(chainID, updatedVaultMap)

	/**********************************************************************************************
	** Once everything is setup, we will store each token in the DB. The storage is set as a map
	** of vaultAddress -> TTokens. All vaults will be retrievable from the store.Interate() func.
	**********************************************************************************************/
	for _, vault := range newVaultMap {
		if (vault.Migration.Target == common.Address{}) {
			vault.Migration = models.TMigration{
				Available: false,
				Target:    vault.Address,
			}
		}
		storage.StoreVault(chainID, vault)
	}
	vaultMap, _ = storage.ListVaults(chainID)
	storage.StoreVaultsToJson(chainID, vaultMap)
	return vaultMap
}
