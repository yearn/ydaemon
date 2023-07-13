package curveVirtualPrice

import (
	"math"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
)

var _virtualPrices = make(map[uint64]*sync.Map)

func Run(chainID uint64) {
	retrieveStoredVirtualPrices(chainID)
	retrieveNewVirtualPrices(chainID)
}

func checkValue(response []interface{}, decimals uint64) (*bigNumber.Int, float64, bool) {
	rawVirtualPrice := response
	virtualPrice := helpers.DecodeBigInt(rawVirtualPrice)
	if virtualPrice.IsZero() {
		return nil, 0, false
	}
	humanizedPricePerShare := helpers.ToNormalizedAmount(virtualPrice, decimals)
	humanizedFloat, _ := humanizedPricePerShare.Float64()
	return virtualPrice, humanizedFloat, true
}

/**********************************************************************************************
** retrieveNewVirtualPrices is used to fetch the new _virtualPrices since the last run. This
** function will fetch the _virtualPrices starting from the last day minus 7 days and store
** them in the _virtualPrices map.
**********************************************************************************************/
func retrieveNewVirtualPrices(chainID uint64) {
	timeBlockMap := map[time.Time]uint64{}
	syncMap := _virtualPrices[chainID]
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
	** to process. We need to exclude the non-curve vaults from our list of vaults.
	**********************************************************************************************/
	allVaults := vaults.ListVaults(chainID)
	unDeployedVaults := []models.TVault{}
	deployedVaults := []models.TVault{}

	for _, vault := range allVaults {
		tokenInstance, ok := tokens.FindToken(chainID, vault.Token.Address)
		if !ok {
			continue
		}
		if tokenInstance.Type == models.TokenTypeCurveLP {
			unDeployedVaults = append(unDeployedVaults, vault)
		}
	}

	/**********************************************************************************************
	** Sort our timeBlockMap by blockNumber to get them ordered.
	**********************************************************************************************/
	keys := []time.Time{}
	for k := range timeBlockMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	logs.Info(`Retrieving new virtual prices started for chain ` + strconv.FormatUint(chainID, 10))
	itemToSave := []store.DBHistoricalValue{}
	for _, time := range keys {
		canUseMulticall := true
		block := timeBlockMap[time]
		if env.MULTICALL_ADDRESSES[chainID].Block > block {
			canUseMulticall = false
		}
		/******************************************************************************************
		** For each unDeployedVault, check if the vault was deployed since the previous time range.
		** If it was, add it to the deployedVaults list and remove it from the unDeployedVaults
		** list.
		** Theses lists are persisted between time ranges.
		******************************************************************************************/
		for _, vault := range unDeployedVaults {
			if (vault.Activation > 0 && vault.Activation <= block) || (vault.Inception <= uint64(time.Unix())) {
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
		** check on the virtualPrice method to get the current price per share of the vault.
		** Here, current is the block we are processing.
		******************************************************************************************/
		calls := []ethereum.Call{}
		if canUseMulticall {
			for _, vault := range updatedDeployedVaults {
				calls = append(calls, multicalls.GetVirtualPrice(vault.Token.Address.Hex(), vault.Token.Address))
				for _, registry := range env.CURVE_REGISTRIES[chainID] {
					if registry.Block > block {
						continue
					}
					calls = append(calls, multicalls.GetVirtualPriceFromLP(vault.Token.Address.Hex()+`_`+registry.Address.Hex(), registry.Address, vault.Token.Address))
				}
			}
		}

		var response map[string][]interface{}
		if canUseMulticall {
			response = multicalls.Perform(chainID, calls, big.NewInt(int64(block)))
		}
		for _, vault := range updatedDeployedVaults {
			found := false
			if canUseMulticall {
				virtualPrice, humanizedFloat, isFound := checkValue(
					response[vault.Token.Address.Hex()+`get_virtual_price`],
					vault.Decimals,
				)
				if isFound {
					found = true
					itemToSave = append(itemToSave, store.DBHistoricalValue{
						UUID:           `will-be-overwritten`,
						Token:          vault.Token.Address.Hex(),
						Value:          virtualPrice.String(),
						HumanizedValue: humanizedFloat,
						Time:           uint64(time.Unix()),
						Block:          block,
						ChainID:        chainID,
					})
				} else {
					for _, registry := range env.CURVE_REGISTRIES[chainID] {
						if registry.Block > block {
							continue
						}
						virtualPrice, humanizedFloat, isFound := checkValue(
							response[vault.Token.Address.Hex()+`_`+registry.Address.Hex()+`get_virtual_price_from_lp_token`],
							vault.Decimals,
						)
						if isFound {
							found = true
							itemToSave = append(itemToSave, store.DBHistoricalValue{
								UUID:           `will-be-overwritten`,
								Token:          vault.Token.Address.Hex(),
								Value:          virtualPrice.String(),
								HumanizedValue: humanizedFloat,
								Time:           uint64(time.Unix()),
								Block:          block,
								ChainID:        chainID,
							})
							break
						}
					}
				}
				if !found {
					LPToken, err := contracts.NewCurveToken(vault.Token.Address, ethereum.GetRPC(chainID))
					if err != nil {
						continue
					}
					virtualPrice, err := LPToken.GetVirtualPrice(nil)
					if err == nil { //Try with the minter
						minter, err := LPToken.Minter(nil)
						if err != nil {
							continue
						}
						minterContract, err := contracts.NewCurveToken(minter, ethereum.GetRPC(chainID))
						if err != nil {
							continue
						}
						virtualPrice, err = minterContract.GetVirtualPrice(nil)
						if err != nil {
							continue
						}
					}
					if bigNumber.SetInt(virtualPrice).IsZero() {
						continue
					}
					found = true
					humanizedPricePerShare := helpers.ToNormalizedAmount(bigNumber.SetInt(virtualPrice), vault.Decimals)
					humanizedFloat, _ := humanizedPricePerShare.Float64()
					itemToSave = append(itemToSave, store.DBHistoricalValue{
						UUID:           `will-be-overwritten`,
						Token:          vault.Token.Address.Hex(),
						Value:          virtualPrice.String(),
						HumanizedValue: humanizedFloat,
						Time:           uint64(time.Unix()),
						Block:          block,
						ChainID:        chainID,
					})
					break
				}
			} else {
				for _, registry := range env.CURVE_REGISTRIES[chainID] {
					if registry.Block > block {
						continue
					}
					registryContract, err := contracts.NewCurveRegistry(registry.Address, ethereum.GetRPC(chainID))
					if err != nil {
						continue
					}
					virtualPrice, err := registryContract.GetVirtualPriceFromLpToken(nil, vault.Token.Address)
					if err != nil {
						continue
					}
					if bigNumber.SetInt(virtualPrice).IsZero() {
						continue
					}
					found = true
					humanizedPricePerShare := helpers.ToNormalizedAmount(bigNumber.SetInt(virtualPrice), vault.Decimals)
					humanizedFloat, _ := humanizedPricePerShare.Float64()
					itemToSave = append(itemToSave, store.DBHistoricalValue{
						UUID:           `will-be-overwritten`,
						Token:          vault.Token.Address.Hex(),
						Value:          virtualPrice.String(),
						HumanizedValue: humanizedFloat,
						Time:           uint64(time.Unix()),
						Block:          block,
						ChainID:        chainID,
					})
					break
				}
			}
		}
		store.StoreVirtualPrice(chainID, itemToSave)
	}
	logs.Success(`Retrieving new virtual prices finished for chain ` + strconv.FormatUint(chainID, 10))

	/**********************************************************************************************
	** Just a for loop to avoid killing the process while some requests are still pending.
	**********************************************************************************************/
	for !store.StoreRateLimiter().Allow() {
		time.Sleep(1 * time.Second)
	}
}

/**********************************************************************************************
** retrieveStoredVirtualPrices is used to populate or repopulate the _virtualPrices sync map
** for a given chain.
** We need to call this function prior to any run which will look for the _virtualPrices to
** avoid refetching previous values
**********************************************************************************************/
func retrieveStoredVirtualPrices(chainID uint64) {
	lastItem := store.DBHistoricalValue{}
	store.DATABASE.
		Table(`db_curve_virtual_prices`).
		Where(`chain_id = ?`, chainID).
		Order(`time DESC`).
		First(&lastItem)

	syncMap := _virtualPrices[chainID]
	if syncMap == nil {
		_virtualPrices[chainID] = &sync.Map{}
		syncMap = _virtualPrices[chainID]
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
