package events

import (
	"context"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
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
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func filterStrategyAdded(
	chainID uint64,
	vaultAddress common.Address,
	vaultVersion string,
	opts *bind.FilterOpts,
	syncMap *sync.Map,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}

	client := ethereum.GetRPC(chainID)
	switch vaultVersion {
	case `0.2.2`:
		vault, _ := contracts.NewYvault022(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := vaultAddress.Hex() + `-` + log.Event.Strategy.Hex()
				syncMap.Store(eventKey, models.TStrategyAdded{
					VaultAddress:    vaultAddress,
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
			logs.Error(`impossible to FilterStrategyAdded for NewYvault022 ` + vaultAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	case `0.3.0`, `0.3.1`:
		vault, _ := contracts.NewYvault030(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := vaultAddress.Hex() + `-` + log.Event.Strategy.Hex()
				syncMap.Store(eventKey, models.TStrategyAdded{
					VaultAddress:    vaultAddress,
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
			logs.Error(`impossible to FilterStrategyAdded for NewYvault030 ` + vaultAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	default: // case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
		vault, _ := contracts.NewYvault043(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}

				eventKey := vaultAddress.Hex() + `-` + log.Event.Strategy.Hex()
				syncMap.Store(eventKey, models.TStrategyAdded{
					VaultAddress:      vaultAddress,
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
			logs.Error(`impossible to FilterStrategyAdded for NewYvault043 ` + vaultAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** Filter all strategiesMigrated events. Theses events are then stored in a map for easy access:
** vaultAddress-oldAddress-newAddress => TEventBlock.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - asyncStrategiesMigration: the async ptr to the map of vaultAddr -> oldStrategyAddr ->
**   newStrategyAddr -> models.TStrategyMigrated
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func filterStrategiesMigrated(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	opts *bind.FilterOpts,
	asyncStrategiesMigration *sync.Map,
	wg *sync.WaitGroup,
) {
	if wg != nil {
		defer wg.Done()
	}

	/**************************************************************************************************
	** First we make sure to connect with our RPC client and get the vault contract.
	**************************************************************************************************/
	client := ethereum.GetRPC(chainID)
	vault, _ := contracts.NewYvault043(vaultAddress, client)

	/**************************************************************************************************
	** As some RPC could run in crippled mode, we need to make sure to stay within the block range.
	** TLDR; querying a lot and a lot of block will probably fail because of communication issues. To
	** avoid this, we split the query in smaller chunks.
	** From initial set of tests, we found out that 9M blocks is the max range we can query without
	** getting any error with our powerful RPCs.
	**************************************************************************************************/
	fromBlock := vaultActivation
	toBlock, _ := client.BlockNumber(context.Background())
	maxBlockRange := uint64(9_000_000)

	/**************************************************************************************************
	** If some overrides are provided, we use them instead of the default values. However, we need to
	** make sure to stay within the block range allowed by the network.
	**************************************************************************************************/
	if opts.Start != vaultActivation || opts.End != nil {
		fromBlock = uint64(opts.Start)
		toBlock = *opts.End
	}

	/**************************************************************************************************
	** We then loop through the block range and query the events. We need to make sure to stay within
	** the block range allowed by the network.
	**************************************************************************************************/
	for blockStart := fromBlock; blockStart < toBlock; blockStart += maxBlockRange {
		end := uint64(blockStart) + maxBlockRange
		opts := &bind.FilterOpts{Start: blockStart, End: &end}
		if maxBlockRange == uint64(math.MaxUint64) {
			opts = &bind.FilterOpts{Start: blockStart}
		}

		if log, err := vault.FilterStrategyMigrated(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				oldAddress := log.Event.OldVersion
				newAddress := log.Event.NewVersion

				eventKey := vaultAddress.Hex() + `-` + oldAddress.Hex() + `-` + newAddress.Hex()
				asyncStrategiesMigration.Store(eventKey, models.TStrategyMigrated{
					VaultAddress:       vaultAddress,
					OldStrategyAddress: oldAddress,
					NewStrategyAddress: newAddress,
					TxHash:             log.Event.Raw.TxHash,
					BlockNumber:        log.Event.Raw.BlockNumber,
					TxIndex:            log.Event.Raw.TxIndex,
					LogIndex:           log.Event.Raw.Index,
				})
			}
		} else {
			logs.Error(`impossible to FilterStrategyMigrated for NewYvault043 ` + vaultAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
			traces.
				Capture(`error`, `impossible to FilterStrategyMigrated for Yvault043 `+vaultAddress.Hex()).
				SetEntity(`strategy`).
				SetExtra(`error`, err.Error()).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
				SetTag(`vaultAddress`, vaultAddress.Hex()).
				Send()
		}
	}
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
	** We will then listen to all events related to the strategies added or migrated to the vaults.
	** We will use a sync.Map to store the events and a WaitGroup to wait for all the goroutines.
	**********************************************************************************************/
	asyncStrategiesForVaults := &sync.Map{}
	asyncStrategiesMigrated := &sync.Map{}

	wg := &sync.WaitGroup{}
	for _, v := range vaultsMap {
		wg.Add(2)
		opts := &bind.FilterOpts{Start: start, End: end}
		if start == 0 {
			opts = &bind.FilterOpts{Start: v.Activation, End: end}
		}
		go filterStrategyAdded(
			chainID,
			v.VaultsAddress,
			v.APIVersion,
			opts,
			asyncStrategiesForVaults,
			wg,
		)
		go filterStrategiesMigrated(
			chainID,
			v.VaultsAddress,
			v.Activation,
			opts,
			asyncStrategiesMigrated,
			wg,
		)
	}
	wg.Wait()

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
		allStrategiesList = append(allStrategiesList, valueParsed)
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
			BlockNumber:       value.(models.TStrategyMigrated).BlockNumber,
			TxIndex:           value.(models.TStrategyMigrated).TxIndex,
			LogIndex:          value.(models.TStrategyMigrated).LogIndex,
		}
		newStrategy.VaultVersion = vaultsMap[vaultAddressParsed].APIVersion
		strategies[vaultAddressParsed][newStrategyAddressParsed] = newStrategy
		allStrategiesList = append(allStrategiesList, newStrategy)
		count++
		return true
	})
	return allStrategiesList
}
