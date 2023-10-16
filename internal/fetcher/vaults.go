package fetcher

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
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
	calls := []ethereum.Call{}
	for _, vault := range vaultMap {
		if vault.IsRetired {
			continue
		}
		if vault.Version == `3.0.0` {
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
	for _, vault := range vaultMap {
		if vault.IsRetired {
			continue
		}
		newVault := vault
		newVault.ChainID = chainID
		newVault.LastUpdate = time.Now()
		if vault.Version == `3.0.0` {
			newVault = handleV3VaultCalls(vault, response)
		} else {
			newVault = handleV2VaultCalls(vault, response)
		}
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
				Address:      currentVault.Address,
				AssetAddress: currentVault.TokenAddress,
				Version:      currentVault.APIVersion,
				Endorsed:     isEndorsed,
				Type:         currentVault.Type,
				Activation:   currentVault.BlockNumber,
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
				updatedVaultMap[vault.Address] = models.TVault{
					Address:      vault.Address,
					AssetAddress: vault.TokenAddress,
					Version:      vault.APIVersion,
					Activation:   vault.BlockNumber,
					Type:         vault.Type,
				}
			}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations.
	**********************************************************************************************/
	newVaultMap := fetchVaultsBasicInformations(chainID, updatedVaultMap)

	/**********************************************************************************************
	** Once everything is setup we can re-store the elements to save them in our storage
	**********************************************************************************************/
	for _, vault := range newVaultMap {
		vault.ChainID = chainID
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
