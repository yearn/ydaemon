package vaultPricePerShare

import (
	"math"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/vaults"
)

var _pricePerShare = make(map[uint64]*sync.Map)

func Run(chainID uint64) {
	retrieveStoredPricePerShare(chainID)
	retrieveHistoricalPricePerShare(chainID)
}

func retrieveHistoricalPricePerShare(chainID uint64) {
	timeBlockMap := map[time.Time]uint64{}
	syncMap := _pricePerShare[chainID]
	if syncMap == nil {
		return
	}
	syncMap.Range(func(key, value interface{}) bool {
		timeBlockMap[key.(time.Time)] = value.(uint64)
		return true
	})

	/**********************************************************************************************
	** The first steps are to init out environment. We need to fetch all our vaults, strategies,
	** tokens, prices, etc. This will allow us to have an exhaustive list of all the data we need
	** to process.
	**********************************************************************************************/
	allVaults := vaults.ListVaults(chainID)

	deployedVaults := []models.TVault{}
	unDeployedVaults := allVaults

	/**********************************************************************************************
	** Sort our timeBlockMap by time. This will allow us to process the blocks in order and select
	** only the vaults that were deployed at the time of the block.
	**********************************************************************************************/
	keys := []time.Time{}
	for k := range timeBlockMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	for _, time := range keys {
		block := timeBlockMap[time]
		/******************************************************************************************
		** For each unDeployedVault, check if the vault was deployed since the previous time range.
		** If it was, add it to the deployedVaults list and remove it from the unDeployedVaults
		** list.
		** Theses lists are persisted between time ranges.
		******************************************************************************************/
		for _, vault := range unDeployedVaults {
			if vault.Activation <= block {
				deployedVaults = append(deployedVaults, vault)
				for i, v := range unDeployedVaults {
					if v.Address.Hex() == vault.Address.Hex() {
						unDeployedVaults = append(unDeployedVaults[:i], unDeployedVaults[i+1:]...)
						break
					}
				}
			}
		}

		/******************************************************************************************
		** Because somehow some deployed vaults are duplicated, we need to remove the duplicates
		** to avoid a way too long multicall (and useless)
		******************************************************************************************/
		updatedDeployedVaults := []models.TVault{}
		for _, vault := range deployedVaults {
			unique := true
			for _, v := range updatedDeployedVaults {
				if v.Address.Hex() == vault.Address.Hex() {
					unique = false
					break
				}
			}
			if unique {
				updatedDeployedVaults = append(updatedDeployedVaults, vault)
			}
		}
		deployedVaults = updatedDeployedVaults

		/******************************************************************************************
		** Once we got the list of vaults to check, we can start the multicall. It's a simple
		** check on the pricePerShare method to get the current price per share of the vault.
		** Here, current is the block we are processing.
		******************************************************************************************/
		calls := []ethereum.Call{}
		for _, vault := range deployedVaults {
			calls = append(calls, multicalls.GetPricePerShare(vault.Address.Hex(), vault.Address))
		}
		response := multicalls.Perform(chainID, calls, big.NewInt(int64(block)))
		itemToSave := []store.DBHistoricalValue{}
		for _, vault := range deployedVaults {
			rawPricePerShare := response[vault.Address.Hex()+`pricePerShare`]
			pricePerShare := helpers.DecodeBigInt(rawPricePerShare)
			if pricePerShare.IsZero() {
				continue
			}
			humanizedPricePerShare := helpers.ToNormalizedAmount(pricePerShare, vault.Decimals)
			humanizedFloat, _ := humanizedPricePerShare.Float64()
			itemToSave = append(itemToSave, store.DBHistoricalValue{
				UUID:           `will-be-overwritten`,
				Token:          vault.Address.Hex(),
				Value:          pricePerShare.String(),
				HumanizedValue: humanizedFloat,
				Time:           uint64(time.Unix()),
				Block:          block,
				ChainID:        chainID,
			})
		}
		store.StorePricePerShare(chainID, itemToSave)
	}

	logs.Success(`Process finished for chain ` + strconv.FormatUint(chainID, 10))

	/**********************************************************************************************
	** Just a for loop to avoid killing the process while some requests are still pending.
	**********************************************************************************************/
	for !store.StoreRateLimiter().Allow() {
		time.Sleep(1 * time.Second)
	}
}

/**********************************************************************************************
** retrieveStoredPricePerShare is used to populate or repopulate the _pricePerShare sync map
** for a given chain.
** We need to call this function prior to any run which will look for the _pricePerShare to
** avoid refetching previous values
**********************************************************************************************/
func retrieveStoredPricePerShare(chainID uint64) {
	lastItem := store.DBHistoricalValue{}
	store.DATABASE.
		Table(`db_vaults_v2_price_per_share`).
		Where(`chain_id = ?`, chainID).
		Order(`time DESC`).
		First(&lastItem)

	syncMap := _pricePerShare[chainID]
	if syncMap == nil {
		_pricePerShare[chainID] = &sync.Map{}
		syncMap = _pricePerShare[chainID]
	}
	/**********************************************************************************************
	** The first step is to find the earliest relevant deployment for yearn on this chain. For us
	** this means the earliest block where we have a registry contract deployed.
	**********************************************************************************************/
	earliestBlock := uint64(lastItem.Block)
	if (earliestBlock == 0) || (earliestBlock == math.MaxUint64) {
		earliestBlock = uint64(math.MaxUint64)
		for _, registry := range env.YEARN_REGISTRIES[chainID] {
			if registry.Block < earliestBlock {
				earliestBlock = registry.Block
			}
		}
	}

	/**********************************************************************************************
	** Then we will play with noon UTC time, starting from the earliest block. We will loop every
	** 24 hours until we reach now. For each day, we will get the block number at noon UTC and
	** store it in the dailyBlockNumber map.
	**********************************************************************************************/
	timestamp := ethereum.GetBlockTime(chainID, earliestBlock)
	timestampAsTime := time.Unix(int64(timestamp), 0)
	noonUTC := time.Date(timestampAsTime.Year(), timestampAsTime.Month(), timestampAsTime.Day(), 12, 0, 0, 0, time.UTC)
	now := time.Now()

	/**********************************************************************************************
	** And here we loop, first checking if we already have the block number for this day, but if
	** not, we will fetch it from DeFiLlama. If we get it, we will store it in the dailyBlockNumber
	** map.
	**********************************************************************************************/
	for noonUTC.Before(now) {
		nextDayNoonUTCTimestamp := noonUTC.Unix()
		data, ok := store.GetTimeBlock(chainID, uint64(nextDayNoonUTCTimestamp))
		if ok {
			syncMap.Store(noonUTC, data)
			noonUTC = noonUTC.AddDate(0, 0, 1)
		} else {
			logs.Pretty(`MISSING`, noonUTC)
		}
	}
}
