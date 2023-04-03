package vaults

import (
	"math"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
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
) (vaultList []*TVault) {
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
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, vault := range vaults {
		calls = append(calls, getPricePerShare(vault.String(), vault))
		calls = append(calls, getAPIVersion(vault.String(), vault))
		calls = append(calls, getPerformanceFee(vault.String(), vault))
		calls = append(calls, getManagementFee(vault.String(), vault))
		calls = append(calls, getToken(vault.String(), vault))
		calls = append(calls, getActivation(vault.String(), vault))
		calls = append(calls, getEmergencyShutdown(vault.String(), vault))
		calls = append(calls, getGovernance(vault.String(), vault))
		calls = append(calls, getGuardian(vault.String(), vault))
		calls = append(calls, getManagement(vault.String(), vault))
		calls = append(calls, getRewards(vault.String(), vault))
		calls = append(calls, getTotalAssets(vault.String(), vault))
		calls = append(calls, getDepositLimit(vault.String(), vault))
		calls = append(calls, getAvailableDepositLimit(vault.String(), vault))
		for i := 0; i < maxStrategiesPerVault; i++ {
			calls = append(calls, getVaultWithdrawalQueue(vault.String(), int64(i), vault))
		}
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64

	/**********************************************************************************************
	** Then we can proceed the responses. Some date will already be available from the list of
	** tokens, so we can already play with that.
	**********************************************************************************************/
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, vault := range vaults {
		rawPricePerShare := response[vault.String()+`pricePerShare`]
		rawApiVersion := response[vault.String()+`apiVersion`]
		rawPerformanceFee := response[vault.String()+`performanceFee`]
		rawManagementFee := response[vault.String()+`managementFee`]
		rawUnderlying := response[vault.String()+`token`]
		rawEmergencyShutdown := response[vault.String()+`emergencyShutdown`]
		rawActivation := response[vault.String()+`activation`]
		rawGovernance := response[vault.String()+`governance`]
		rawGuardian := response[vault.String()+`guardian`]
		rawManagement := response[vault.String()+`management`]
		rawRewards := response[vault.String()+`rewards`]
		rawTotalAssets := response[vault.String()+`totalAssets`]
		rawDepositLimit := response[vault.String()+`depositLimit`]
		rawAvailableDepositLimit := response[vault.String()+`availableDepositLimit`]

		shareTokenData, ok := tokens.FindToken(chainID, vault)
		if !ok {
			shareTokenData = &tokens.TERC20Token{}
			traces.
				Capture(`warn`, `impossible to retrieve share token for vault `+vault.Hex()+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`meta`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`vaultAddress`, vault.Hex()).
				Send()
		}

		underlyingTokenData, ok := tokens.FindToken(chainID, helpers.DecodeAddress(rawUnderlying))
		if !ok {
			tokenAddress := helpers.DecodeAddress(rawUnderlying)
			underlyingTokenData = &tokens.TERC20Token{}
			if !metaVaultTokenErrorAlreadySent[chainID][tokenAddress] {
				traces.
					Capture(`warn`, `impossible to retrieve underlying token for vault `+vault.Hex()+` on chain `+strconv.FormatUint(chainID, 10)).
					SetEntity(`meta`).
					SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
					SetTag(`vaultAddress`, vault.Hex()).
					SetTag(`underlyingAddress`, helpers.DecodeAddress(rawUnderlying).Hex()).
					Send()
				metaVaultTokenErrorAlreadySent[chainID][tokenAddress] = true
			}
		}

		vaultData, ok := meta.GetMetaVault(chainID, vault)
		if !ok {
			if !metaVaultFileErrorAlreadySent[chainID][vault] {
				traces.
					Capture(`warn`, `impossible to retrieve meta file for vault `+vault.Hex()+` on chain `+strconv.FormatUint(chainID, 10)).
					SetEntity(`meta`).
					SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
					SetTag(`vaultAddress`, vault.Hex()).
					Send()
				metaVaultFileErrorAlreadySent[chainID][vault] = true
			}
		}

		/******************************************************************************************
		** Preparing our new TVault object
		******************************************************************************************/
		newVault := &TVault{
			ChainID:               chainID,
			Address:               vault,
			Name:                  shareTokenData.Name,
			Symbol:                shareTokenData.Symbol,
			Decimals:              shareTokenData.Decimals,
			Icon:                  env.GITHUB_ICON_BASE_URL + strconv.FormatUint(chainID, 10) + `/` + vault.Hex() + `/logo-128.png`,
			Type:                  utils.VaultTypeStandard,
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
			Token: tokens.TERC20Token{
				Address:       underlyingTokenData.Address,
				Name:          underlyingTokenData.Name,
				DisplayName:   underlyingTokenData.DisplayName,
				DisplaySymbol: underlyingTokenData.DisplaySymbol,
				Symbol:        underlyingTokenData.Symbol,
				Description:   underlyingTokenData.Description,
				Decimals:      underlyingTokenData.Decimals,
				Icon:          underlyingTokenData.Icon,
			},
		}

		if vaultData != nil && vaultData.DisplayName != `` {
			newVault.BuildNames(vaultData.DisplayName)
		} else {
			newVault.BuildNames(shareTokenData.DisplayName)
		}
		newVault.BuildSymbol(shareTokenData.DisplaySymbol)

		/******************************************************************************************
		** Adding the withdrawal queue stuffs
		******************************************************************************************/
		withdrawQueueForStrategies := []common.Address{}
		for i := 0; i < maxStrategiesPerVault; i++ {
			result := response[vault.String()+strconv.FormatInt(int64(i), 10)+`withdrawalQueue`]
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
	vaultMap map[common.Address]*TVault,
) map[common.Address]*TVault {
	newMap := make(map[common.Address]*TVault)
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
	vaults map[common.Address]utils.TVaultsFromRegistry,
) map[common.Address]*TVault {
	trace := traces.Init(`app.indexer.vaults.multicall_data`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `vaults`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** First, try to retrieve the list of vaults from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	vaultMap := make(map[common.Address]*TVault)
	store.ListFromBadgerDB(chainID, store.TABLES.VAULTS, &vaultMap)

	/**********************************************************************************************
	** From the vault registry we have the first batch of vaults. In order to proceed, we will
	** create a map of TVault which will be stored in the DB. This action should only be performed
	** for every refresh period, per vault per chainID.
	** TODO: Optimize this part to split possible changes (pricePerShare) from almost impossible
	** changes (name, symbol)
	**********************************************************************************************/
	updatedVaultMap := make(map[common.Address]*TVault)
	for _, currentVault := range vaults {
		updatedVaultMap[currentVault.VaultsAddress] = &TVault{
			Address:    currentVault.TokenAddress,
			Endorsed:   (currentVault.Type == utils.VaultTypeStandard || currentVault.Type == utils.VaultTypeAutomated) && currentVault.TokenAddress != common.Address{},
			Type:       currentVault.Type,
			Activation: currentVault.Activation,
		}
	}

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the vault data for them.
	** We will add them manually here.
	**********************************************************************************************/
	extraVaults := map[uint64][]string{
		1:     {},
		10:    {},
		250:   {},
		42161: {},
	}
	for _, vaultAddress := range extraVaults[chainID] {
		vaultAddress := common.HexToAddress(vaultAddress)
		if _, ok := vaultMap[vaultAddress]; !ok {
			updatedVaultMap[vaultAddress] = &TVault{
				Address: vaultAddress,
			}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations.
	**********************************************************************************************/
	if len(updatedVaultMap) > 0 {
		updatedVaultMap = findAllVaults(chainID, updatedVaultMap)

		/**********************************************************************************************
		** Once everything is setup, we will store each token in the DB. The storage is set as a map
		** of vaultAddress -> TTokens. All vaults will be retrievable from the store.Interate() func.
		**********************************************************************************************/
		wg := sync.WaitGroup{}
		wg.Add(len(updatedVaultMap))
		for _, token := range updatedVaultMap {
			go func(_token *TVault) {
				defer wg.Done()
				store.SaveInBadgerDB(
					chainID,
					store.TABLES.VAULTS,
					_token.Address.String(),
					_token,
				)
			}(token)
		}
		wg.Wait()
		store.ListFromBadgerDB(chainID, store.TABLES.VAULTS, &vaultMap)
	}

	_vaultMap[chainID] = vaultMap
	return vaultMap
}
