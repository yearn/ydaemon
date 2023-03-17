package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fees"
	"github.com/yearn/ydaemon/internal/indexer"
	bribes "github.com/yearn/ydaemon/internal/indexer.bribes"
	partnerTracker "github.com/yearn/ydaemon/internal/indexer.partnerTracker"
	"github.com/yearn/ydaemon/internal/prices"
	"github.com/yearn/ydaemon/internal/registries"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/tokensList"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

var AllHarvests = make(map[common.Address][]vaults.THarvest)

var STRATLIST = []strategies.TStrategy{}

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
func runRetrieveAllVaults(chainID uint64, vaultsMap map[common.Address]utils.TVaultsFromRegistry, wg *sync.WaitGroup, delay time.Duration) {
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
		indexer.PostProcessStrategies(chainID)
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
	InitializePartnerTracker(chainID)
	os.Exit(0)

	defer wg.Done()
	go InitializeBribes(chainID)
	go InitializePartnerTracker(chainID)

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

	go tokensList.BuildTokenList(chainID)
	strategiesAddedList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsMap)

	//From our list of strategies, perform a multicall to get all strategies data -> Should be done every 5(?) minutes for all strategies
	internalWG.Add(1)
	go runRetrieveAllStrategies(chainID, strategiesAddedList, &internalWG, 5*time.Minute)
	internalWG.Wait()

	go registries.IndexNewVaults(chainID)

	Initialize(chainID)
}

func InitializeBribes(chainID uint64) {
	bribes.RetrieveAllRewardsAdded(chainID)
}
func InitializePartnerTracker(chainID uint64) {
	/**********************************************************************************************
	** First we need to catch all the events that happened in the past to be able to calculate the
	** current state of the partner tracker
	**********************************************************************************************/
	logs.Info(`Loading all referral balance increase events`)
	timeBefore := time.Now()
	partnerTracker.RetrieveAllReffererBalanceIncrease(chainID)
	logs.Success(`Loaded all referral balance increase events tooks`, time.Since(timeBefore))

	/**********************************************************************************************
	** Once we got all the events, we can check how many unique referrer, referee and vaults we
	** have and start building our relation tree:
	** [chainID][vault][referrer][referee][amount]
	**********************************************************************************************/
	allEvents := partnerTracker.ListReferralBalanceIncrease(chainID)
	uniqueReferrer := map[common.Address]uint64{}
	uniqueReferee := map[common.Address]bool{}
	uniqueVaults := map[common.Address]bool{}
	for _, event := range allEvents {
		uniqueReferrer[event.PartnerID.Address]++
		uniqueReferee[event.Depositer.Address] = true
		uniqueVaults[event.Vault.Address] = true
	}
	logs.Success(`Found`, len(uniqueReferrer), `unique referrer`)
	logs.Success(`Found`, len(uniqueReferee), `unique referee`)
	logs.Success(`Found`, len(uniqueVaults), `unique vaults`)

	/**********************************************************************************************
	** Now that we have all the unique referrer, referee and vaults, we can start building our
	** relation tree
	**********************************************************************************************/
	partnerTrackerTree := map[uint64]map[common.Address]map[common.Address]map[common.Address]*bigNumber.Int{}
	for _, event := range allEvents {
		/******************************************************************************************
		** Ugly go code to avoid crash because of nil pointer
		******************************************************************************************/
		if partnerTrackerTree[chainID] == nil {
			partnerTrackerTree[chainID] = map[common.Address]map[common.Address]map[common.Address]*bigNumber.Int{}
		}
		if partnerTrackerTree[chainID][event.Vault.ToAddress()] == nil {
			partnerTrackerTree[chainID][event.Vault.ToAddress()] = map[common.Address]map[common.Address]*bigNumber.Int{}
		}
		if partnerTrackerTree[chainID][event.Vault.ToAddress()][event.PartnerID.ToAddress()] == nil {
			partnerTrackerTree[chainID][event.Vault.ToAddress()][event.PartnerID.ToAddress()] = map[common.Address]*bigNumber.Int{}
		}
		if partnerTrackerTree[chainID][event.Vault.ToAddress()][event.PartnerID.ToAddress()][event.Depositer.ToAddress()] == nil {
			partnerTrackerTree[chainID][event.Vault.ToAddress()][event.PartnerID.ToAddress()][event.Depositer.ToAddress()] = bigNumber.NewInt(0)
		}

		/******************************************************************************************
		** Actual code to add the amount to the tree
		******************************************************************************************/
		currentAmount := partnerTrackerTree[chainID][event.Vault.ToAddress()][event.PartnerID.ToAddress()][event.Depositer.ToAddress()]

		partnerTrackerTree[chainID][event.Vault.ToAddress()][event.PartnerID.ToAddress()][event.Depositer.ToAddress()] = currentAmount.Add(
			currentAmount,
			event.Amount,
		)
	}

	//save that in a json file name partnerTrackerTree.json
	jsonData, err := json.Marshal(partnerTrackerTree)
	if err != nil {
		logs.Error(err)
	}
	err = ioutil.WriteFile("partnerTrackerTree.json", jsonData, 0644)
	if err != nil {
		logs.Error(err)
	}

	// logs.Pretty(
	// 	len(partnerTracker.ListReferralBalanceIncrease(chainID)),
	// 	partnerTracker.ListReferralBalanceIncrease(chainID)[0],
	// 	time.Since(timeBefore),
	// )

}

func Initialize(chainID uint64) {
	timeBefore := time.Now()
	/**********************************************************************************************
	** All vaults from Yearn are registered in the registries contracts. A vault can be either
	** Standard, Experimental or Automated.
	** From the registries, we are fetching all vaults along with the block in which they were
	** added to the registry, and we remove the duplicates only to keep the latest version of a
	** same vault. Duplicate can happen when a vault is moved from Experimental to Standard.
	**********************************************************************************************/
	// vaultsList := registries.RetrieveAllVaults(chainID, 0)

	// strategiesMap := map[common.Address]map[common.Address]strategies.TStrategyAdded{}
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
	// tokens.RetrieveAllTokens(chainID, vaultsList)
	// vaults.RetrieveAllVaults(chainID, vaultsList)

	/**********************************************************************************************
	** Fetching all the strategiesList and relevant transfers to proceed
	**********************************************************************************************/
	vaultsMap := vaults.MapVaults(chainID)
	strategiesList := strategies.ListStrategies(chainID)

	// // only grab for that vault 0xe9dc63083c464d6edccff23444ff3cfc6886f6fb
	// vaultsMap := map[common.Address]*vaults.TVault{}
	// for k, v := range oldvaultsMap {
	// 	if v.Address.Hex() == common.HexToAddress("0xe9dc63083c464d6edccff23444ff3cfc6886f6fb").Hex() {
	// 		vaultsMap[v.Address] = oldvaultsMap[k]
	// 	}
	// }

	// // only grab for that strategy 0x126e4fdfa9dcea94f8f4157ef8ad533140c60fc7
	// strategiesList := []*strategies.TStrategy{}
	// for k, s := range oldstrategiesList {
	// 	if s.Address.Hex() == common.HexToAddress("0x126e4fdfa9dcea94f8f4157ef8ad533140c60fc7").Hex() {
	// 		strategiesList = append(strategiesList, oldstrategiesList[k])
	// 	}
	// }

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
		// strategiesList := strategies.RetrieveAllStrategiesAdded(chainID, vaultsList)
		strategiesMap := strategies.SplitStrategiesAddedPerVault(strategiesList)

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
			vaultsMap,
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
		transfersFromVaultsToTreasury = vaults.RetrieveAllTransferFromVaultsToTreasury(chainID, vaultsMap)
	}()

	go func() {
		defer wg.Done()
		/**********************************************************************************************
		** Retrieve all harvest events for a vault. This will enable us to know where to look and to
		** compute the gains, losses and the fees.
		**********************************************************************************************/
		allHarvests = vaults.RetrieveHarvests(chainID, vaultsMap)
	}()
	wg.Wait()
	logs.Success("Initialization done in", time.Since(timeBefore))

	syncGroup := &sync.WaitGroup{}
	harvests := []vaults.THarvest{}
	for _, vault := range vaultsMap {
		switch vault.Version {
		case `0.2.2`:
		case `0.3.0`:
			continue //SKIP
		default: //case `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
			syncGroup.Add(1)
			go vaults.HandleEvenStrategyReportedFor031To043(
				chainID,
				vault,
				managementFees[vault.Address],
				performanceFees[vault.Address],
				strategiesPerformanceFees[vault.Address],
				transfersFromVaultsToStrategies[vault.Address],
				transfersFromVaultsToTreasury[vault.Address],
				allHarvests[vault.Address],
				syncGroup,
				&harvests,
			)
		}
	}
	syncGroup.Wait()

	allHarvestPerStrategyPerVault := map[common.Address]map[common.Address][]vaults.THarvest{}
	count := 0
	for _, v := range harvests {
		AllHarvests[v.Vault] = append(AllHarvests[v.Vault], v)
		if allHarvestPerStrategyPerVault[v.Vault] == nil {
			allHarvestPerStrategyPerVault[v.Vault] = map[common.Address][]vaults.THarvest{}
		}
		allHarvestPerStrategyPerVault[v.Vault][v.Strategy] = append(allHarvestPerStrategyPerVault[v.Vault][v.Strategy], v)
		count++
	}

	// // For each harvest for each vault, compute the APY
	// type harvestAPY struct {
	// }
	// for _, vault := range vaultsMap {
	// 	allHarvestForThisVault := AllHarvests[vault.Address]
	// 	lastHarvestTime := vault.Inception
	// 	for _, harvest := range allHarvestForThisVault {
	// 		gains := harvest.Gain
	// 		losses := harvest.Loss
	// 		totalFees := harvest.Fees.TotalCollectedFee
	// 		harvestTime := harvest.Timestamp

	// 		lastHarvestTime = harvestTime

	// 	}
	// }
	curveSTETHVault := common.HexToAddress(`0x5B8C556B8b2a78696F0B9B830B3d67623122E270`)
	curveSTETHStrategy := common.HexToAddress(`0xaBec96AC9CdC6863446657431DD32F73445E80b1`)
	allRelevantHarvest := allHarvestPerStrategyPerVault[curveSTETHVault][curveSTETHStrategy]

	strategy, ok := strategies.FindStrategy(1, curveSTETHStrategy)
	if !ok {
		logs.Error("Strategy not found")
		return
	}

	lastHarvest := ethereum.GetBlockTime(1, strategy.Activation.Uint64())
	for _, harvest := range allRelevantHarvest {
		/**********************************************************************************************
		** For each harvest, retrieve the relevant data: gains, losses, fees and amount allocated. They
		** are all in WEI notation, so we need to convert them to normalized value with FormatedAmount.
		** If no gain or loss, skip.
		**********************************************************************************************/
		gains, _ := helpers.FormatAmount(harvest.Gain.String(), 18)
		losses, _ := helpers.FormatAmount(harvest.Loss.String(), 18)
		totalFees, _ := helpers.FormatAmount(harvest.Fees.TotalCollectedFee.String(), 18)
		amountAllocated, _ := helpers.FormatAmount(harvest.TotalDebt.String(), 18)
		gainsMinusLoss := gains - losses
		if gainsMinusLoss == 0 { //Skip, no gain or loss, no APY
			lastHarvest = harvest.Timestamp
			continue
		}

		/**********************************************************************************************
		** The result is simply the gains minus the losses minus the fees. We then divide by the amount
		** allocated to get the ratio.
		** Here we use both the result with fees and without fees.
		**********************************************************************************************/
		resultMinusFees := gainsMinusLoss - totalFees
		resultRatio := resultMinusFees / amountAllocated
		resultRatioNoFees := gainsMinusLoss / amountAllocated

		/**********************************************************************************************
		** We need to compute the time while this amount was allocated, aka result for this amount in
		** this time. With that we can extrapolate the APY for the year.
		**********************************************************************************************/
		lastHarvestTime := time.Unix(int64(lastHarvest), 0)
		harvestTimeTime := time.Unix(int64(harvest.Timestamp), 0)
		timeSinceHarvest := harvestTimeTime.Sub(lastHarvestTime)
		durationSinceHarvest := timeSinceHarvest.Hours()
		hourPerYear := float64(365 * 24)

		APR := (resultRatio * (hourPerYear / durationSinceHarvest) * 100)
		APRNoFees := (resultRatioNoFees * (hourPerYear / durationSinceHarvest) * 100)
		logs.Pretty(APR, APRNoFees)

		lastHarvest = harvest.Timestamp
	}

}
