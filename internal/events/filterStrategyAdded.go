package events

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** Filter all StrategyAdded events and store them in a map of vaultAddress-strategyAddress
** => TEventBlock
** Version [0.2.2] - [0.3.0 & 0.3.1] - [>= 0.3.2] have different events structure and thus need
** to be handled differently.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - syncMap: the async ptr to the map of vaultAddr -> strategyAddr -> blockNumber
**   -> TEventBlock
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func filterStrategyAdded(
	vault models.TVaultsFromRegistry,
	vaultsLastBlockSync map[common.Address]uint64,
	start uint64,
	end *uint64,
	syncMap *sync.Map,
) {
	client := ethereum.GetRPC(vault.ChainID)

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	/******************************************************************************************
	** Then, we need to know when to start our log fetching. By default, we will fetch from the
	** activation block in order to get all the vaults that were ever deployed since it was
	** deployed.
	** However, if the external database is used, we may want to only fetch the logs that were
	** emitted after the last time we fetched the logs. In that case, we will use the last
	** event block number + 1 as the start block number.
	** If another start block number is specified, we will use it instead.
	******************************************************************************************/
	if start == 0 {
		lastEvent := vaultsLastBlockSync[vault.Address]
		if lastEvent > 0 {
			start = lastEvent + 1 //Adding one to get the next event
		} else {
			start = vault.Activation
		}
	}

	/******************************************************************************************
	** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
	** avoid hitting some external node providers' rate limits.
	******************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[vault.ChainID] {
		chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[vault.ChainID]
		if chunkEnd > *end {
			chunkEnd = *end
		}
		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}

		switch vault.APIVersion {
		case `0.2.2`:
			contract, _ := contracts.NewYvault022(vault.Address, client)
			if log, err := contract.FilterStrategyAdded(opts, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					eventKey := vault.Address.Hex() + `-` + log.Event.Strategy.Hex()
					syncMap.Store(eventKey, models.TStrategyAdded{
						VaultAddress:    vault.Address,
						StrategyAddress: log.Event.Strategy,
						TxHash:          log.Event.Raw.TxHash,
						BlockNumber:     log.Event.Raw.BlockNumber,
						TxIndex:         log.Event.Raw.TxIndex,
						LogIndex:        log.Event.Raw.Index,
						DebtLimit:       bigNumber.SetInt(log.Event.DebtLimit),
						RateLimit:       bigNumber.SetInt(log.Event.RateLimit),
						PerformanceFee:  bigNumber.SetInt(log.Event.PerformanceFee),
					})
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault022 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
			}
		case `0.3.0`, `0.3.1`:
			contract, _ := contracts.NewYvault030(vault.Address, client)
			if log, err := contract.FilterStrategyAdded(opts, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}
					eventKey := vault.Address.Hex() + `-` + log.Event.Strategy.Hex()
					syncMap.Store(eventKey, models.TStrategyAdded{
						VaultAddress:    vault.Address,
						StrategyAddress: log.Event.Strategy,
						TxHash:          log.Event.Raw.TxHash,
						BlockNumber:     log.Event.Raw.BlockNumber,
						TxIndex:         log.Event.Raw.TxIndex,
						LogIndex:        log.Event.Raw.Index,
						DebtRatio:       bigNumber.SetInt(log.Event.DebtRatio),
						RateLimit:       bigNumber.SetInt(log.Event.RateLimit),
						PerformanceFee:  bigNumber.SetInt(log.Event.PerformanceFee),
					})
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault030 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
			}
		default: // case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
			contract, _ := contracts.NewYvault043(vault.Address, client)
			if log, err := contract.FilterStrategyAdded(opts, nil); err == nil {
				for log.Next() {
					if log.Error() != nil {
						continue
					}

					eventKey := vault.Address.Hex() + `-` + log.Event.Strategy.Hex()
					syncMap.Store(eventKey, models.TStrategyAdded{
						VaultAddress:      vault.Address,
						StrategyAddress:   log.Event.Strategy,
						TxHash:            log.Event.Raw.TxHash,
						BlockNumber:       log.Event.Raw.BlockNumber,
						TxIndex:           log.Event.Raw.TxIndex,
						LogIndex:          log.Event.Raw.Index,
						DebtRatio:         bigNumber.SetInt(log.Event.DebtRatio),
						MaxDebtPerHarvest: bigNumber.SetInt(log.Event.MaxDebtPerHarvest),
						MinDebtPerHarvest: bigNumber.SetInt(log.Event.MinDebtPerHarvest),
						PerformanceFee:    bigNumber.SetInt(log.Event.PerformanceFee),
					})
				}
			} else {
				logs.Error(`impossible to FilterStrategyAdded for NewYvault043 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
			}
		}
	}

	/**********************************************************************************************
	** We are storing in the DB the sync status, indicating we went up to the block number to check
	** for new vaults.
	**********************************************************************************************/
	go store.StoreSyncStrategiesAdded(vault.ChainID, vault.Address, end)
}

/**************************************************************************************************
** Filter all strategiesMigrated events. Theses events are then stored in a map for easy access:
** vaultAddress-oldAddress-newAddress => TEventBlock.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - syncMap: the async ptr to the map of vaultAddr -> oldStrategyAddr ->
**   newStrategyAddr -> models.TStrategyMigrated
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func filterStrategiesMigrated(
	vault models.TVaultsFromRegistry,
	vaultsLastBlockSync map[common.Address]uint64,
	start uint64,
	end *uint64,
	syncMap *sync.Map,
) {
	/**************************************************************************************************
	** First we make sure to connect with our RPC client and get the vault contract.
	**************************************************************************************************/
	client := ethereum.GetRPC(vault.ChainID)
	contract, _ := contracts.NewYvault043(vault.Address, client)

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	/******************************************************************************************
	** Then, we need to know when to start our log fetching. By default, we will fetch from the
	** activation block in order to get all the vaults that were ever deployed since it was
	** deployed.
	******************************************************************************************/
	if start == 0 {
		lastEvent := vaultsLastBlockSync[vault.Address]
		if lastEvent > 0 {
			start = lastEvent + 1 //Adding one to get the next event
		} else {
			start = vault.Activation
		}
	}

	/******************************************************************************************
	** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
	** avoid hitting some external node providers' rate limits.
	******************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[vault.ChainID] {
		chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[vault.ChainID]
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := contract.FilterStrategyMigrated(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				oldAddress := log.Event.OldVersion
				newAddress := log.Event.NewVersion

				eventKey := vault.Address.Hex() + `-` + oldAddress.Hex() + `-` + newAddress.Hex()
				syncMap.Store(eventKey, models.TStrategyMigrated{
					VaultAddress:       vault.Address,
					OldStrategyAddress: oldAddress,
					NewStrategyAddress: newAddress,
					TxHash:             log.Event.Raw.TxHash,
					BlockNumber:        log.Event.Raw.BlockNumber,
					TxIndex:            log.Event.Raw.TxIndex,
					LogIndex:           log.Event.Raw.Index,
				})
			}
		} else {
			logs.Error(`impossible to FilterStrategyMigrated for NewYvault043 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
		}
	}

	/**********************************************************************************************
	** We are storing in the DB the sync status, indicating we went up to the block number to check
	** for new vaults.
	**********************************************************************************************/
	go store.StoreSyncStrategiesAdded(vault.ChainID, vault.Address, end)
}

/**************************************************************************************************
** HandleStrategyAdded will loop over a map of vaults and fetch all the strategy added events
** from start to end. Based on the version of the vault, the function will call the correct
** function to fetch the events.
** This will also fetch the strategyMigrated events as they can be missed
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultsMap: the map of vaults we want to fetch the fee for
** - start: the block number to start fetching from
** - end: the block number to stop fetching from
**
** Returns:
** - An array of TMimicStrategyReportBase via the T any. This hack is used to avoid import circles.
**************************************************************************************************/
func HandleStrategyAdded(
	chainID uint64,
	vaultsMap map[common.Address]models.TVaultsFromRegistry,
	start uint64,
	end *uint64,
) []models.TStrategyAdded {
	/**********************************************************************************************
	** Our first action is to make sure we ignore the strategies we already have in our local
	** storage, which we loaded from the database.
	**********************************************************************************************/
	var strategyAddedSync []store.DBStrategyAddedSync
	store.DATABASE.Table("db_strategy_added_syncs").
		Where("chain_id = ?", chainID).
		Find(&strategyAddedSync)

	vaultsLastBlockSync := map[common.Address]uint64{}
	for _, syncEvent := range strategyAddedSync {
		vaultsLastBlockSync[common.HexToAddress(syncEvent.Vault)] = syncEvent.Block
	}
	allPreviouslyAddedStrategies, _ := store.ListAllStrategiess(chainID)

	/**********************************************************************************************
	** We will then listen to all events related to the strategies added or migrated to the vaults.
	** We will use a sync.Map to store the events and a WaitGroup to wait for all the goroutines.
	**********************************************************************************************/
	asyncStrategiesForVaults := &sync.Map{}
	asyncStrategiesMigrated := &sync.Map{}

	for _, v := range vaultsMap {
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go func(v models.TVaultsFromRegistry) {
			defer wg.Done()
			filterStrategyAdded(
				v,
				vaultsLastBlockSync,
				start,
				end,
				asyncStrategiesForVaults,
			)
		}(v)
		go func(v models.TVaultsFromRegistry) {
			defer wg.Done()
			filterStrategiesMigrated(
				v,
				vaultsLastBlockSync,
				start,
				end,
				asyncStrategiesMigrated,
			)
		}(v)
		wg.Wait()
	}

	/**********************************************************************************************
	** Once all vaults StrategyAdded updates have been retrieved, we need to extract them from the
	** sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-strategyAddress
	** - value: []models.TStrategyAdded
	**
	** We need to transform it into a map as follows:
	** - vaultAddress -> strategyAddress -> []TEventBlock
	**********************************************************************************************/
	count := 0
	allStrategiesList := []models.TStrategyAdded{}
	strategies := make(map[common.Address]map[common.Address]models.TStrategyAdded)
	asyncStrategiesForVaults.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddressParsed := common.HexToAddress(eventKey[0])
		strategyAddressParsed := common.HexToAddress(eventKey[1])

		if _, ok := strategies[vaultAddressParsed]; !ok {
			strategies[vaultAddressParsed] = make(map[common.Address]models.TStrategyAdded)
		}
		valueParsed := value.(models.TStrategyAdded)
		valueParsed.VaultVersion = vaultsMap[vaultAddressParsed].APIVersion
		strategies[vaultAddressParsed][strategyAddressParsed] = valueParsed
		valueParsed.ChainID = chainID
		allStrategiesList = append(allStrategiesList, valueParsed)
		go store.StoreStrategies(chainID, valueParsed)
		count++
		return true
	})

	/**********************************************************************************************
	** Once all vaults StrategyMigrated updates have been retrieved, we need to extract them from
	** the sync.Map to append them to the strategies map. Theses strategies should use the last
	** state of the previous strategy in order to keep the data consistent.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-oldStrategyAddress-newStrategyAddress
	** - value: []models.TStrategyMigrated
	**
	** We need to transform it into a map as follows:
	** - vaultAddress -> strategyAddress -> []TEventBlock
	**********************************************************************************************/
	asyncStrategiesMigrated.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddressParsed := common.HexToAddress(eventKey[0])
		oldStrategyAddressParsed := common.HexToAddress(eventKey[1])
		newStrategyAddressParsed := common.HexToAddress(eventKey[2])

		if _, ok := strategies[vaultAddressParsed]; !ok {
			strategies[vaultAddressParsed] = make(map[common.Address]models.TStrategyAdded)
		}
		oldStrategy := strategies[vaultAddressParsed][oldStrategyAddressParsed]
		newStrategy := models.TStrategyAdded{
			VaultAddress:      vaultAddressParsed,
			StrategyAddress:   newStrategyAddressParsed,
			TxHash:            value.(models.TStrategyMigrated).TxHash,
			DebtRatio:         oldStrategy.DebtRatio,
			MaxDebtPerHarvest: oldStrategy.MaxDebtPerHarvest,
			MinDebtPerHarvest: oldStrategy.MinDebtPerHarvest,
			PerformanceFee:    oldStrategy.PerformanceFee,
			DebtLimit:         oldStrategy.DebtLimit,
			RateLimit:         oldStrategy.RateLimit,
			ChainID:           chainID,
			BlockNumber:       value.(models.TStrategyMigrated).BlockNumber,
			TxIndex:           value.(models.TStrategyMigrated).TxIndex,
			LogIndex:          value.(models.TStrategyMigrated).LogIndex,
		}
		newStrategy.VaultVersion = vaultsMap[vaultAddressParsed].APIVersion
		strategies[vaultAddressParsed][newStrategyAddressParsed] = newStrategy
		allStrategiesList = append(allStrategiesList, newStrategy)
		go store.StoreStrategies(chainID, newStrategy)
		count++
		return true
	})

	for _, strat := range allPreviouslyAddedStrategies {
		allStrategiesList = append(allStrategiesList, strat)
	}

	return allStrategiesList
}
