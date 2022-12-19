package internal

import (
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/bribes"
	"github.com/yearn/ydaemon/internal/fees"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

var AllHarvests = make(map[ethcommon.Address][]vaults.THarvest)

var STRATLIST = []strategies.TStrategy{}

func postProcessStrategies(chainID uint64) {
	// Overwrite the st-yCRV strategy to set debtRatio to 100% and totalDebt to vault supply
	if chainID == 1 {
		styCRVStrategy, ok := strategies.FindStrategy(chainID, ethcommon.HexToAddress(`0xE7863292dd8eE5d215eC6D75ac00911D06E59B2d`))
		if ok {
			styCRVVault, ok := vaults.FindVault(chainID, common.FromAddress(styCRVStrategy.VaultAddress))
			if ok {
				styCRVStrategy.DebtRatio = bigNumber.NewUint64(10000)
				styCRVStrategy.TotalDebt = styCRVVault.TotalAssets
			}
		}
	}
}

func runRetrieveAllPrices(chainID uint64, wg *sync.WaitGroup, delay time.Duration) {
	isDone := false
	for {
		prices.RetrieveAllPrices(chainID)
		if !isDone {
			isDone = true
			wg.Done()
		}
		if delay == 0 {
			return
		}
		time.Sleep(delay)
	}
}
func runRetrieveAllVaults(chainID uint64, vaultsMap map[ethcommon.Address]utils.TVaultsFromRegistry, wg *sync.WaitGroup, delay time.Duration) {
	isDone := false
	for {
		vaults.RetrieveAllVaults(chainID, vaultsMap)
		logs.Debug(`DONE`)
		if !isDone {
			isDone = true
			wg.Done()
		}
		if delay == 0 {
			return
		}
		time.Sleep(delay)
	}
}
func runRetrieveAllStrategies(chainID uint64, strategiesAddedList []strategies.TStrategyAdded, wg *sync.WaitGroup, delay time.Duration) {
	isDone := false
	for {
		strategies.RetrieveAllStrategies(chainID, strategiesAddedList)
		if !isDone {
			isDone = true
			wg.Done()
		}
		if delay == 0 {
			return
		}
		time.Sleep(delay)
	}
}

func InitializeV2(chainID uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	go InitializeBribes(chainID)

	var internalWG sync.WaitGroup
	//From the events in the registries, retrieve all the vaults -> Should only be done on init or when a new vault is added
	vaultsMap := registries.RetrieveAllVaults(chainID, 0)

	// From our list of vaults, retrieve the ERC20 data for each shareToken, underlyingToken and the underlying of those tokens
	// -> Data store will not change unless extreme event, so this should only be done on init or when a new vault is added
	tokens.RetrieveAllTokens(chainID, vaultsMap)

	// From our list of tokens, retieve the price for each one of them -> Should be done every 1(?) minute for all tokens
	internalWG.Add(1)
	go runRetrieveAllPrices(chainID, &internalWG, 1*time.Minute)
	internalWG.Wait()

	//From our list of vault, perform a multicall to get all vaults data -> Should be done every 5(?) minutes for all vaults
	internalWG.Add(1)
	go runRetrieveAllVaults(chainID, vaultsMap, &internalWG, 5*time.Minute)
	internalWG.Wait()

	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsMap)

	//From our list of strategies, perform a multicall to get all strategies data -> Should be done every 5(?) minutes for all strategies
	internalWG.Add(1)
	go runRetrieveAllStrategies(chainID, strategiesAddedList, &internalWG, 5*time.Minute)
	internalWG.Wait()

	postProcessStrategies(chainID)

	go registries.IndexNewVaults(chainID)
}

func InitializeBribes(chainID uint64) {
	bribes.RetrieveAllRewardsAdded(chainID)
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
	vaultsList := registries.RetrieveAllVaults(chainID, 0)

	strategiesMap := map[ethcommon.Address]map[ethcommon.Address]strategies.TStrategyAdded{}
	transfersFromVaultsToTreasury := map[ethcommon.Address]map[uint64][]utils.TEventBlock{}
	transfersFromVaultsToStrategies := map[ethcommon.Address]map[ethcommon.Address]map[uint64][]utils.TEventBlock{}
	managementFees := map[ethcommon.Address]map[uint64][]utils.TEventBlock{}
	performanceFees := map[ethcommon.Address]map[uint64][]utils.TEventBlock{}
	strategiesPerformanceFees := map[ethcommon.Address]map[ethcommon.Address]map[uint64][]utils.TEventBlock{}
	allHarvests := map[ethcommon.Address]map[ethcommon.Address]map[uint64]uint64{}

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
		default: //case `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
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

	logs.Success(`It tooks`, time.Since(timeBefore), `to process`, count, `harvests`)
}
