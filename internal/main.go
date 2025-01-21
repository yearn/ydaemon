package internal

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-co-op/gocron"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/processes/apr"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/prices"
	"github.com/yearn/ydaemon/processes/risks"
)

var STRATLIST = []models.TStrategy{}

func initStakingPools(chainID uint64, scheduler *gocron.Scheduler) {
	/**********************************************************************************************
	** Start the Staking Indexing process and schedule it to run every hour. This indexer
	** will fetch the relevent data for the OP Staking contracts as well as thoose responsible
	** for the APR calculations.
	**********************************************************************************************/
	indexer.IndexStakingPools(chainID)
	indexer.IndexVeYFIStakingContract(chainID)
	indexer.IndexJuicedStakingContract(chainID)
	indexer.IndexV3StakingContract(chainID)
	logs.Success(chainID, `-`, `InitStakingPools ✅`)

	scheduler.Every(1).Hours().WaitForSchedule().Do(func() {
		indexer.IndexStakingPools(chainID)
		indexer.IndexVeYFIStakingContract(chainID)
		indexer.IndexJuicedStakingContract(chainID)
		indexer.IndexV3StakingContract(chainID)
	})
}

func initVaults(chainID uint64, scheduler *gocron.Scheduler) (
	map[common.Address]models.TVaultsFromRegistry,
	map[common.Address]models.TVault,
	map[common.Address]models.TERC20Token,
) {
	/** 🔵 - Yearn *************************************************************************************
	** InitializeV2 is only called on initialization. It's first job is to retrieve the initial data:
	** - The registries vaults
	** - The vaults
	** - The strategies
	** - The tokens
	**************************************************************************************************/
	indexer.IndexYearnXPoolTogetherVaults(chainID)
	registries := indexer.IndexNewVaults(chainID)
	logs.Success(chainID, `-`, `InitRegistries ✅`, len(registries))
	vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
	logs.Success(chainID, `-`, `InitVaults ✅`, len(vaultMap))
	tokenMap := fetcher.RetrieveAllTokens(chainID, vaultMap)
	logs.Success(chainID, `-`, `InitTokens ✅`, len(tokenMap))

	scheduler.Every(2).Hours().WaitForSchedule().Do(func() {
		indexer.IndexYearnXPoolTogetherVaults(chainID)
		registries := indexer.IndexNewVaults(chainID)
		logs.Success(chainID, `-`, `InitRegistries ✅`, len(registries))
		vaultMap := fetcher.RetrieveAllVaults(chainID, registries)
		logs.Success(chainID, `-`, `InitVaults ✅`, len(vaultMap))
		fetcher.RetrieveAllTokens(chainID, vaultMap)
		logs.Success(chainID, `-`, `InitTokens ✅`, len(tokenMap))
	})
	return registries, vaultMap, tokenMap
}

func initStrategies(chainID uint64, scheduler *gocron.Scheduler, vaultMap map[common.Address]models.TVault) {
	/** 🔵 - Yearn *************************************************************************************
	** initStrategies is only called on initialization. It's first job is to retrieve the strategies
	** and then to schedule the retrieval of the strategies every 3 hours
	**************************************************************************************************/
	strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
	fetcher.RetrieveAllStrategies(chainID, strategiesMap)
	logs.Success(chainID, `-`, `InitStrategies ✅`, len(strategiesMap))

	scheduler.Every(3).Hours().WaitForSchedule().Do(func() {
		strategiesMap := indexer.IndexNewStrategies(chainID, vaultMap)
		fetcher.RetrieveAllStrategies(chainID, strategiesMap)
		logs.Success(chainID, `-`, `InitStrategies ✅`, len(strategiesMap))
	})
}

func InitializeV2(chainID uint64, wg *sync.WaitGroup, scheduler *gocron.Scheduler) {
	if wg != nil {
		defer wg.Done()
	}

	_, vaultMap, tokenMap := initVaults(chainID, scheduler)

	initStakingPools(chainID, scheduler)
	initStrategies(chainID, scheduler, vaultMap)

	/**********************************************************************************************
	** Retrieving prices and strategies for all the given token and strategies on that chain.
	** This is done in parallel to speed up the process and reduce the time it takes to complete.
	** The scheduler is used to retrieve the strategies every 15 minutes.
	** Computing APRS
	**********************************************************************************************/
	prices.RetrieveAllPrices(chainID, tokenMap)
	logs.Success(chainID, `-`, `RetrieveAllPrices ✅`)
	apr.ComputeChainAPY(chainID)
	logs.Success(chainID, `-`, `ComputeChainAPY ✅`)

	scheduler.Every(1).Days().At("01:00").StartImmediately().Do(func() {
		risks.RetrieveAllRiskScores(chainID, vaultMap)
	})

	scheduler.Every(30).Minute().WaitForSchedule().Do(func() {
		currentTokenMap, _ := storage.ListERC20(chainID)
		prices.RetrieveAllPrices(chainID, currentTokenMap)
		apr.ComputeChainAPY(chainID)
	})

	/**********************************************************************************************
	** Do some background work
	**********************************************************************************************/
	scheduler.Every(10).Hours().StartImmediately().At("12:10").Do(func() {
		initDailyBlock.Run(chainID)
	})
	scheduler.StartAsync()
}
