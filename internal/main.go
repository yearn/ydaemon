package internal

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fees"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

var AllHarvests = make(map[common.Address][]vaults.THarvest)

var STRATLIST = []strategies.TStrategy{}

func InitializeV2(chainID uint64) {
	vaultsList := registries.RetrieveAllVaults(chainID)
	tokens.RetrieveAllTokens(chainID, vaultsList)
	prices.RetrieveAllPrices(chainID)
	vaults.RetrieveAllVaults(chainID, vaultsList)
	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsList)
	strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
}

func Initialize(chainID uint64) {
	timeBefore := time.Now()
	/**********************************************************************************************
	** All vaults from Yearn are registered in the registries contracts. A vault can be either
	** Standard or Experimental.
	** From the registries, we are fetching all vaults along with the block in which they were
	** added to the registry, and we remove the duplicates only to keep the latest version of a
	** same vault. Duplicate can happen when a vault is moved from Experimental to Standard.
	**********************************************************************************************/
	vaultsList := registries.RetrieveAllVaults(chainID)

	strategiesMap := map[common.Address]map[common.Address]strategies.TStrategyAdded{}
	transfersFromVaultsToTreasury := map[common.Address]map[uint64][]utils.TEventBlock{}
	transfersFromVaultsToStrategies := map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock{}
	managementFees := map[common.Address]map[uint64][]utils.TEventBlock{}
	performanceFees := map[common.Address]map[uint64][]utils.TEventBlock{}
	strategiesPerformanceFees := map[common.Address]map[common.Address]map[uint64][]utils.TEventBlock{}
	allHarvests := map[common.Address]map[common.Address]map[uint64]uint64{}

	/**********************************************************************************************
	** Retrieve all tokens used by Yearn, along with the underlying tokens. The tokens are only
	** retrieved for the new vaults, as the old vaults should have been already registered in the
	** database.
	** The function returns a map of token address to token data.
	** The function store the tokens in a table [chainID] [token address] [token data], in the
	** data/store/[chainID]/tokens folder.
	**********************************************************************************************/
	tokens.RetrieveAllTokens(chainID, vaultsList)
	vaults.RetrieveAllVaults(chainID, vaultsList)

	/**********************************************************************************************
	** Fetching all the strategiesList and relevant transfers to proceed
	**********************************************************************************************/
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()

		/**********************************************************************************************
		** Retrieve all the strategies ever attached to a vault. This will be used in the next step
		** to retrieve the transfer events for the strategists fees.
		** With this process, we are retrieving the standard blockEvents elements and all the arguments
		** from the `StrategyAdded` event.
		**********************************************************************************************/
		strategiesList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsList)
		strategiesMap = strategies.SlipStrategiesAddedPerVault(strategiesList)

		/**********************************************************************************************
		** Retrieve all transfers from vaults to strategies. This can only happen in one situation: the
		** vault is sending strategist fees to the strategy for them to be taken by the strategist.
		** We need that to be able to calculate the strategist fees as many variable could make the
		** offchain calculation wrong.
		** Thanks to this number, from offchain totalFees calculation, we can deduct the treasury fees
		**********************************************************************************************/
		transfersFromVaultsToStrategies = vaults.RetrieveAllTransferFromVaultsToStrategies(chainID, strategiesMap)

		/**********************************************************************************************
		** For each vault we need to know the fee per block, which is the percentage of gains after each
		** harvest that will be sent to the governance. This is a dynamic value, and it can be changed
		** by the governance. We need to fetch all the events of type `UpdateManagementFee`,
		** `UpdateStrategyPerformanceFee` and `UpdatePerformanceFee` and build an historical mapping of
		** the fee per block, knowing for each block which fee to use.
		**********************************************************************************************/
		managementFees, performanceFees, strategiesPerformanceFees = fees.RetrieveAllFeesBPS(
			chainID,
			vaultsList,
			strategiesMap,
		)
	}()

	go func() {
		defer wg.Done()
		/**********************************************************************************************
		** Retrieve all transfers from vaults to treasury. This can only happen in one situation: the
		** vault is sending managements fees to the treasury after a harvest.
		** We need that to be able to calculate the actual fees as many variable could make the
		** offchain calculation wrong.
		**********************************************************************************************/
		transfersFromVaultsToTreasury = vaults.RetrieveAllTransferFromVaultsToTreasury(chainID, vaultsList)
	}()

	go func() {
		defer wg.Done()
		/**********************************************************************************************
		** Retrieve all harvest events for a vault. This will enable us to know where to look and to
		** compute the gains, losses and the fees.
		**********************************************************************************************/
		allHarvests = vaults.RetrieveHarvests(chainID, vaultsList)
	}()
	wg.Wait()

	logs.Success("Initialization done in", time.Since(timeBefore))

	timeBefore = time.Now()
	syncGroup := &sync.WaitGroup{}
	harvests := []vaults.THarvest{}
	for _, vault := range vaultsList {
		switch vault.APIVersion {
		case `0.2.2`:
		case `0.3.0`:
			continue //SKIP
		case `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
			syncGroup.Add(1)
			go vaults.HandleEvenStrategyReportedFor031To043(
				chainID,
				vault,
				managementFees[vault.VaultsAddress],
				performanceFees[vault.VaultsAddress],
				strategiesPerformanceFees[vault.VaultsAddress],
				transfersFromVaultsToStrategies[vault.VaultsAddress],
				transfersFromVaultsToTreasury[vault.VaultsAddress],
				allHarvests[vault.VaultsAddress],
				syncGroup,
				&harvests,
			)
		}
	}
	syncGroup.Wait()

	count := 0
	for _, v := range harvests {
		AllHarvests[v.Vault] = append(AllHarvests[v.Vault], v)
		count++
	}

	logs.Success(`It took`, time.Since(timeBefore), `to process`, count, `harvests`)
}
