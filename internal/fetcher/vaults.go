package fetcher

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
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
	chunkSize := 50
	vaultList = []models.TVault{}
	vaultSlice := []models.TVault{}
	for _, vault := range vaultMap {
		vaultSlice = append(vaultSlice, vault)
	}

	/**********************************************************************************************
	** We prepare the data by chunking the vaults to avoid overloading the multicall.
	**********************************************************************************************/
	chunks := [][]models.TVault{}
	for i := 0; i < len(vaultSlice); i += chunkSize {
		end := i + chunkSize
		if end > len(vaultSlice) {
			end = len(vaultSlice)
		}
		chunks = append(chunks, vaultSlice[i:end])
	}

	totalChunks := len(chunks)
	for chunkIndex, chunk := range chunks {
		logs.Info(`Processing vault chunk ` + strconv.Itoa(chunkIndex+1) + `/` + strconv.Itoa(totalChunks) + ` (` + strconv.Itoa(len(chunk)) + ` vaults)`)
		/**********************************************************************************************
		** The first step is to prepare the multicall, connecting to the multicall instance and
		** preparing the array of calls to send. All calls for all vaults will be send in a single
		** multicall and will later be accessible via a concatened string `vaultAddress + methodName`.
		**********************************************************************************************/
		calls := []ethereum.Call{}
		for _, vault := range chunk {
			// Adding an exception for the vault that is retired but we still want to compute. Alchemix related
			isException := addresses.Equals(vault.Address, "0xaD17A225074191d5c8a37B50FdA1AE278a2EE6A2") ||
				addresses.Equals(vault.Address, "0x5B977577Eb8a480f63e11FC615D6753adB8652Ae") ||
				addresses.Equals(vault.Address, "0x65343F414FFD6c97b0f6add33d16F6845Ac22BAc") ||
				addresses.Equals(vault.Address, "0xFaee21D0f0Af88EE72BB6d68E54a90E6EC2616de")
			if vault.Metadata.IsRetired && !isException {
				continue
			}
			versionMajor := strings.Split(vault.Version, `.`)[0]
			if versionMajor == `3` {
				calls = append(calls, getV3VaultCalls(vault)...)
			} else {
				calls = append(calls, getV2VaultCalls(vault)...)
			}
		}

		/**********************************************************************************************
		** Then we can proceed the responses. Some date will already be available from the list of
		** tokens, so we can already play with that.
		**********************************************************************************************/
		response := multicalls.Perform(chainID, calls, nil)
		for _, vault := range chunk {
			// Adding an exception for the vault that is retired but we still want to compute. Alchemix related
			isException := addresses.Equals(vault.Address, "0xaD17A225074191d5c8a37B50FdA1AE278a2EE6A2") ||
				addresses.Equals(vault.Address, "0x5B977577Eb8a480f63e11FC615D6753adB8652Ae") ||
				addresses.Equals(vault.Address, "0x65343F414FFD6c97b0f6add33d16F6845Ac22BAc") ||
				addresses.Equals(vault.Address, "0xFaee21D0f0Af88EE72BB6d68E54a90E6EC2616de")
			if vault.Metadata.IsRetired && !isException {
				continue
			}
			newVault := vault
			newVault.ChainID = chainID
			versionMajor := strings.Split(vault.Version, `.`)[0]
			if versionMajor == `3` {
				newVault = handleV3VaultCalls(vault, response)
			} else {
				newVault = handleV2VaultCalls(vault, response)
			}
			vaultList = append(vaultList, newVault)
		}
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
	method TProcessNewVaultMethod,
) map[common.Address]models.TVault {
	chain, ok := env.GetChain(chainID)
	if !ok {
		logs.Error(chainID, `-`, `RetrieveAllVaults`, `Chain not found`)
		return nil
	}
	
	vaultCount := len(vaults)
	logs.Info(`Fetching details for ` + strconv.Itoa(vaultCount) + ` vaults on chain ` + strconv.FormatUint(chainID, 10))

	/**********************************************************************************************
	** First, try to retrieve the list of vaults from the database and populate our updatedVaultMap
	** with it.
	**********************************************************************************************/
	vaultMap, _ := storage.ListVaults(chainID)
	metadata := storage.GetVaultsJsonMetadata(chainID)
	shouldRefresh := metadata.ShouldRefresh
	updatedVaultMap := vaultMap
	if method == ProcessNewVaultMethodAppend {
		updatedVaultMap = map[common.Address]models.TVault{}
	}

	/**********************************************************************************************
	** From the vault registry we have the first batch of vaults. In order to proceed, we will
	** create a map of TVault which will be stored in the DB. This action should only be performed
	** for every refresh period, per vault per chainID.
	**********************************************************************************************/
	for _, currentVault := range vaults {
		if _, ok := vaultMap[currentVault.Address]; !ok || shouldRefresh {
			if (currentVault.TokenAddress == common.Address{}) {
				continue
			}
			isEndorsed := (currentVault.Type == models.TokenTypeStandardVault || currentVault.Type == models.TokenTypeAutomatedVault)
			kind := currentVault.Kind
			if currentVault.Kind == `` {
				kind = models.VaultKindLegacy
			}
			newVault := models.TVault{
				Address:      currentVault.Address,
				AssetAddress: currentVault.TokenAddress,
				Version:      currentVault.APIVersion,
				ChainID:      chainID,
				Endorsed:     isEndorsed,
				Type:         currentVault.Type,
				Kind:         kind,
				Activation:   currentVault.BlockNumber,
			}
			
			// Assign Kong debts to vault
			if kongDebts, ok := storage.GetKongDebts(chainID, currentVault.Address); ok {
				if kongDebtsJSON, err := json.Marshal(kongDebts); err == nil {
					newVault.KongDebts = string(kongDebtsJSON)
				}
			}
			updatedVaultMap[currentVault.Address] = newVault
		}
	}

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the vault data for them.
	** We will add them manually here.
	**********************************************************************************************/
	if method == ProcessNewVaultMethodReplace {
		for _, currentVault := range chain.ExtraVaults {
			if _, ok := updatedVaultMap[currentVault.Address]; !ok || shouldRefresh {
				kind := currentVault.Kind
				if currentVault.Kind == `` {
					kind = models.VaultKindLegacy
				}
				newVault := models.TVault{
					Address:      currentVault.Address,
					AssetAddress: currentVault.TokenAddress,
					Version:      currentVault.APIVersion,
					ChainID:      chainID,
					Activation:   currentVault.BlockNumber,
					Type:         currentVault.Type,
					Kind:         kind,
				}
				
				// Assign Kong debts to vault
				if kongDebts, ok := storage.GetKongDebts(chainID, currentVault.Address); ok {
					if kongDebtsJSON, err := json.Marshal(kongDebts); err == nil {
						newVault.KongDebts = string(kongDebtsJSON)
					}
				}
				updatedVaultMap[currentVault.Address] = newVault
			}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations.
	**********************************************************************************************/
	newVaultList := fetchVaultsBasicInformations(chainID, updatedVaultMap)

	/**********************************************************************************************
	** Once everything is setup we can re-store the elements to save them in our storage
	**********************************************************************************************/
	for _, vault := range newVaultList {
		vault.ChainID = chainID
		/******************************************************************************************
		** In some situation, a vault can be added to multiple registries: by default the public
		** one, and sometime another one on top which should be considered as the actual one.
		** To properly handle this, we will check assign the registry address we got when the
		** stored one is the zero address or when the registry is the public one.
		******************************************************************************************/
		if addresses.Equals(vault.RegistryAddress, common.Address{}) {
			vault.RegistryAddress = vaults[vault.Address].RegistryAddress
		} else if env.IsRegistryFromPublicERC4626(chainID, vault.RegistryAddress) {
			vault.RegistryAddress = vaults[vault.Address].RegistryAddress
		}

		/******************************************************************************************
		** We need to check if the associated registry is marked as hidden. If so, we need to mark
		** the vault as hidden as well.
		******************************************************************************************/
		isRegistryHidden := false
		for _, registry := range chain.Registries {
			if addresses.Equals(registry.Address, vault.RegistryAddress) {
				if registry.Tag == `STEALTH` {
					isRegistryHidden = true
					break
				}
			}
		}
		if isRegistryHidden {
			vault.Metadata.IsHidden = true
		}

		/******************************************************************************************
		** If no migration target is set, we will set it to the vault address as a default value,
		** and put the available flag to false.
		******************************************************************************************/
		if (vault.Metadata.Migration.Target == common.Address{}) {
			vault.Metadata.Migration = models.TMigration{
				Available: false,
				Target:    vault.Address,
			}
		}

		/******************************************************************************************
		** If the vault has some automated Type, we need to mark it as automated in the metadata.
		******************************************************************************************/
		if vault.Type == models.TokenTypeLegacyAutomatedVault || vault.Type == models.TokenTypeAutomatedVault {
			vault.Metadata.IsAutomated = true
		}

		/******************************************************************************************
		** If no stability is set, we will set it to unknown as a default value.
		******************************************************************************************/
		if vault.Metadata.Stability.Stability == `` {
			vault.Metadata.Stability.Stability = models.VaultStabilityUnknown
		}

		/******************************************************************************************
		** If the vault has no category, we will set it to automatic as a default value.
		******************************************************************************************/
		if vault.Metadata.Category == `` {
			vault.Metadata.Category = models.VaultCategoryAutomatic
		}

		storage.StoreVault(chainID, vault)
	}

	/**********************************************************************************************
	** Somehow, some properties are not properly updated. Let's make sure we have them properly
	** set.
	** If the registry address for a given vault is set to the zero address, we will set it to the
	** registry address of the vault we have in the vaults map.
	**********************************************************************************************/
	vaultMapFromStorage, _ := storage.ListVaults(chainID)
	for _, vault := range vaultMap {
		/******************************************************************************************
		** In some situation, a vault can be added to multiple registries: by default the public
		** one, and sometime another one on top which should be considered as the actual one.
		** To properly handle this, we will check assign the registry address we got when the
		** stored one is the zero address or when the registry is the public one.
		******************************************************************************************/
		if addresses.Equals(vault.RegistryAddress, common.Address{}) {
			vault.RegistryAddress = vaults[vault.Address].RegistryAddress
		} else if env.IsRegistryFromPublicERC4626(chainID, vault.RegistryAddress) {
			vault.RegistryAddress = vaults[vault.Address].RegistryAddress
		}

		/******************************************************************************************
		** If the inclusion is not set, we will set it based on the registry address. Some specific
		** rules exist, but overall, it's all based on the registry
		******************************************************************************************/
		if !vault.Metadata.Inclusion.IsSet {
			vault.Metadata.Inclusion.IsYearn = env.IsRegistryFromYearnCore(chainID, vault.RegistryAddress)
			vault.Metadata.Inclusion.IsYearnJuiced = env.IsRegistryFromJuiced(chainID, vault.RegistryAddress)
			vault.Metadata.Inclusion.IsPublicERC4626 = env.IsRegistryFromPublicERC4626(chainID, vault.RegistryAddress)
			vault.Metadata.Inclusion.IsPoolTogether = env.IsRegistryFromPoolTogether(chainID, vault.RegistryAddress)
			vault.Metadata.Inclusion.IsCove = env.IsRegistryFromCove(chainID, vault.RegistryAddress)
			vault.Metadata.Inclusion.IsMorpho = false //False by default
			vault.Metadata.Inclusion.IsKatana = false //False by default
			vault.Metadata.Inclusion.IsGimme = false  //False by default

			isYearn := vault.Metadata.Inclusion.IsYearn || vault.Metadata.Inclusion.IsYearnJuiced || vault.Metadata.Inclusion.IsGimme
			isPublic := vault.Metadata.Inclusion.IsPublicERC4626

			if isYearn {
				vault.Endorsed = true
				vault.Metadata.IsHidden = false
			}
			if isPublic && !isYearn {
				vault.Endorsed = false
				vault.Metadata.IsHidden = true
				vault.Metadata.IsHighlighted = false
			}

			vault.Metadata.Inclusion.IsSet = true
		}

		/******************************************************************************************
		** Finally, if the category is set to automatic, we will try to determine the category and
		** set it properly, updating the JSON with the new value in the end. This should be a one
		** time operation.
		******************************************************************************************/
		if vault.Metadata.Category == models.VaultCategoryAutomatic || vault.Metadata.Category == `` {
			strategies, _ := storage.ListStrategiesForVault(vault.ChainID, vault.Address)
			category := BuildVaultCategory(vault, strategies)

			if category != `Volatile` {
				vault.Metadata.Category = models.TVaultCategoryType(category)
			}
		}
		vaultMapFromStorage[vault.Address] = vault
	}
	storage.StoreVaultsToJson(chainID, vaultMapFromStorage)

	returnedVaultMap := map[common.Address]models.TVault{}

	// If the method is append, we only want the vaults that we asked to be added in the Vaults args
	if method == ProcessNewVaultMethodAppend {
		for _, vault := range vaults {
			returnedVaultMap[vault.Address] = vaultMapFromStorage[vault.Address]
		}
		return returnedVaultMap
	}
	return vaultMapFromStorage
}
