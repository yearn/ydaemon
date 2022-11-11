package vaults

import (
	"math"
	"math/big"
	"strconv"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/external/meta"
	"github.com/yearn/ydaemon/external/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
)

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
	vaults []ethcommon.Address,
) (vaultList []*TVault) {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all vaults will be send in a single
	** multicall and will later be accessible via a concatened string `vaultAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, vault := range vaults {
		calls = append(calls, getPricePerShare(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getAPIVersion(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getPerformanceFee(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getManagementFee(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getToken(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getActivation(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getEmergencyShutdown(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getGovernance(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getGuardian(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getManagement(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getRewards(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getTotalAssets(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getDepositLimit(vault.String(), common.FromAddress(vault)))
		calls = append(calls, getAvailableDepositLimit(vault.String(), common.FromAddress(vault)))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	maxBatch := math.MaxInt64
	if chainID == 250 {
		maxBatch = 3
	}

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

		shareTokenData, ok := tokens.FindToken(chainID, common.FromAddress(vault))
		if !ok {
			shareTokenData = &tokens.TERC20Token{}
			logs.Warning(`NO SHARE TOKEN FOUND FOR VAULT`, vault.String())
		}

		underlyingTokenData, ok := tokens.FindToken(chainID, common.FromAddress(rawUnderlying[0].(ethcommon.Address)))
		if !ok {
			underlyingTokenData = &tokens.TERC20Token{}
			logs.Warning(`NO UNDERLYING TOKEN FOUND FOR VAULT`, vault.String())
		}

		vaultFromMeta, ok := meta.Store.VaultsFromMeta[chainID][common.FromAddress(vault)]
		if !ok {
			// logs.Warning(`No meta file found for vault`, vault.String())
			// If the vault file is missing, we set the default values for its fields
			vaultFromMeta = meta.TVaultFromMeta{
				Order:               1000000000,
				HideAlways:          false,
				DepositsDisabled:    false,
				WithdrawalsDisabled: false,
				MigrationAvailable:  false,
				AllowZapIn:          true,
				AllowZapOut:         true,
				Retired:             false,
			}
		}

		/******************************************************************************************
		** THIS PART MUST BE CLEANED
		******************************************************************************************/
		humanizedPrice, fHumanizedPrice := buildTokenPrice(
			chainID,
			common.FromAddress(rawUnderlying[0].(ethcommon.Address)),
		)
		fHumanizedTVLPrice := buildTVL(
			bigNumber.SetInt(rawTotalAssets[0].(*big.Int)),
			int(shareTokenData.Decimals),
			humanizedPrice,
		)
		delegatedTokenAsBN := bigNumber.NewInt(0)
		fDelegatedValue := 0.0

		strategies := []strategies.TStrategy{}
		for _, strat := range strategies {
			stratDelegatedValueAsFloat, err := strconv.ParseFloat(strat.DelegatedValue, 64)
			if err == nil {
				delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, strat.DelegatedAssets)
				fDelegatedValue += stratDelegatedValueAsFloat
			}
		}
		/******************************************************************************************
		** ! THIS PART MUST BE CLEANED
		******************************************************************************************/

		/******************************************************************************************
		** Preparing our new TVault object
		******************************************************************************************/
		newVault := &TVault{
			Address:            vault,
			Name:               shareTokenData.Name,
			Symbol:             shareTokenData.Symbol,
			Decimals:           shareTokenData.Decimals,
			Icon:               env.GITHUB_ICON_BASE_URL + strconv.FormatUint(chainID, 10) + `/` + vault.Hex() + `/logo-128.png`,
			Type:               `v2`,  //Always v2 for now
			Endorsed:           false, //Default to false, will be updated later
			Inception:          rawActivation[0].(*big.Int).Uint64(),
			Version:            rawApiVersion[0].(string),
			PricePerShare:      *bigNumber.SetInt(rawPricePerShare[0].(*big.Int)),
			Emergency_shutdown: rawEmergencyShutdown[0].(bool),
			Strategies:         strategies,
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
			TVL: TTVL{
				TotalAssets:          bigNumber.SetInt(rawTotalAssets[0].(*big.Int)),
				TotalDelegatedAssets: delegatedTokenAsBN,
				TVL:                  fHumanizedTVLPrice - fDelegatedValue,
				TVLDeposited:         fHumanizedTVLPrice,
				TVLDelegated:         fDelegatedValue,
				Price:                fHumanizedPrice,
			},
			Details: &TVaultDetails{
				Management:            common.FromAddress(rawManagement[0].(ethcommon.Address)),
				Governance:            common.FromAddress(rawGovernance[0].(ethcommon.Address)),
				Guardian:              common.FromAddress(rawGuardian[0].(ethcommon.Address)),
				Rewards:               common.FromAddress(rawRewards[0].(ethcommon.Address)),
				DepositLimit:          bigNumber.SetInt(rawDepositLimit[0].(*big.Int)),
				AvailableDepositLimit: bigNumber.SetInt(rawAvailableDepositLimit[0].(*big.Int)),
				PerformanceFee:        uint64(rawPerformanceFee[0].(*big.Int).Uint64()),
				ManagementFee:         uint64(rawManagementFee[0].(*big.Int).Uint64()),
				Comment:               vaultFromMeta.Comment,
				Order:                 vaultFromMeta.Order,
				DepositsDisabled:      vaultFromMeta.DepositsDisabled,
				WithdrawalsDisabled:   vaultFromMeta.WithdrawalsDisabled,
				AllowZapIn:            vaultFromMeta.AllowZapIn,
				AllowZapOut:           vaultFromMeta.AllowZapOut,
				Retired:               vaultFromMeta.Retired,
				APYTypeOverride:       vaultFromMeta.APYTypeOverride,
				APYOverride:           vaultFromMeta.APYOverride,
			},
		}

		newVault.BuildNames(shareTokenData.DisplayName)
		newVault.BuildSymbol(shareTokenData.DisplaySymbol)
		newVault.BuildMigration(chainID)
		newVault.BuildAPY(chainID)
		newVault.APY.Fees.Performance = float64(uint64(rawPerformanceFee[0].(*big.Int).Uint64()) / 10000)
		newVault.APY.Fees.Management = float64(uint64(rawManagementFee[0].(*big.Int).Uint64()) / 10000)

		if vaultFromMeta.APYTypeOverride != `` {
			newVault.APY.Type = vaultFromMeta.APYTypeOverride
		}
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
	vaultMap map[ethcommon.Address]*TVault,
) map[ethcommon.Address]*TVault {
	newMap := make(map[ethcommon.Address]*TVault)
	vaultListAddresses := []ethcommon.Address{}
	for vaultAddress := range vaultMap {
		vaultListAddresses = append(vaultListAddresses, vaultAddress)
	}

	newVaultMap := fetchBasicInformations(chainID, vaultListAddresses)
	for _, vault := range newVaultMap {
		vault.Endorsed = vaultMap[vault.Address].Endorsed
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
	vaults map[ethcommon.Address]utils.TVaultsFromRegistry,
) map[ethcommon.Address]*TVault {
	timeBefore := time.Now()

	/**********************************************************************************************
	** First, try to retrieve the list of vaults from the database to exclude the one existing
	** from the upcoming calls
	**********************************************************************************************/
	vaultMap := make(map[ethcommon.Address]*TVault)
	store.Iterate(chainID, store.TABLES.VAULTS, &vaultMap)

	/**********************************************************************************************
	** From the vault registry we have the first batch of vaults. In order to proceed, we will
	** create a map of TVault which will be stored in the DB. This action should only be performed
	** for every refresh period, per vault per chainID.
	** TODO: Optimize this part to split possible changes (pricePerShare) from almost impossible
	** changes (name, symbol)
	**********************************************************************************************/
	updatedVaultMap := make(map[ethcommon.Address]*TVault)
	for _, currentVault := range vaults {
		updatedVaultMap[currentVault.VaultsAddress] = &TVault{
			Address:  currentVault.TokenAddress,
			Endorsed: currentVault.Type == `Standard`,
		}
	}

	/**********************************************************************************************
	** Somehow, some vaults are not in the registries, but we still need the vault data for them.
	** We will add them manually here.
	**********************************************************************************************/
	extraVaults := []string{
		// `0x34fe2a45D8df28459d7705F37eD13d7aE4382009`, // yvWBTC
	}
	for _, vaultAddress := range extraVaults {
		vaultAddress := ethcommon.HexToAddress(vaultAddress)
		if _, ok := vaultMap[vaultAddress]; !ok {
			updatedVaultMap[vaultAddress] = &TVault{
				Address: vaultAddress,
			}
		}
	}

	/**********************************************************************************************
	** Once we got out basic list, we will fetch theses basics informations. This includes:
	** - the name
	** - the symbol
	** - the decimals
	** Based on the type of token (aave, compound, curve, etc), we will also fetch the underlyings
	** vaults.
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
				store.SaveInDB(
					chainID,
					store.TABLES.VAULTS,
					_token.Address.String(),
					_token,
				)
			}(token)
		}
		wg.Wait()
		store.Iterate(chainID, store.TABLES.VAULTS, &vaultMap)
	}

	logs.Success(`It took`, time.Since(timeBefore), `to retrieve`, len(vaultMap), `vaults`)
	_vaultMap[chainID] = vaultMap
	return vaultMap
}
