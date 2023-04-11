package strategies

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
	"github.com/yearn/ydaemon/internal/utils"
)

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
func getStrategiesMigrated(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	asyncStrategiesMigration *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := ethereum.GetRPC(chainID)
	vault, _ := contracts.NewYvault043(vaultAddress, client)
	blockEnd, _ := client.BlockNumber(context.Background())
	maxBlockRange := uint64(math.MaxUint64)
	maxBlockRange = 9_000_000

	for blockStart := vaultActivation; blockStart < blockEnd; blockStart += maxBlockRange {
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
			logs.Error(err)
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
** Filter all StrategyAdded events and store them in a map of vaultAddress-strategyAddress
** => TEventBlock
** Version [0.2.2] - [0.3.0 & 0.3.1] - [>= 0.3.2] have different events structure and thus need
** to be handled differently.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - vaultActivation: the block number at which the vault was activated
** - asyncStrategiesForVaults: the async ptr to the map of vaultAddr -> strategyAddr -> blockNumber
**   -> TEventBlock
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as the asyncFeeMap is updated via a pointer
**************************************************************************************************/
func getStrategiesAdded(
	chainID uint64,
	vaultAddress common.Address,
	vaultActivation uint64,
	vaultVersion string,
	asyncStrategiesForVaults *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := ethereum.GetRPC(chainID)
	switch vaultVersion {
	case `0.2.2`:
		vault, _ := contracts.NewYvault022(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(&bind.FilterOpts{Start: vaultActivation}, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := vaultAddress.Hex() + `-` + log.Event.Strategy.Hex()
				asyncStrategiesForVaults.Store(eventKey, models.TStrategyAdded{
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
		}
	case `0.3.0`, `0.3.1`:
		vault, _ := contracts.NewYvault030(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(&bind.FilterOpts{Start: vaultActivation}, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := vaultAddress.Hex() + `-` + log.Event.Strategy.Hex()
				asyncStrategiesForVaults.Store(eventKey, models.TStrategyAdded{
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
		}
	default: // case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
		vault, _ := contracts.NewYvault043(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(&bind.FilterOpts{Start: vaultActivation}, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}

				eventKey := vaultAddress.Hex() + `-` + log.Event.Strategy.Hex()
				asyncStrategiesForVaults.Store(eventKey, models.TStrategyAdded{
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
		}
	}
}

/**********************************************************************************************
** Once we got all the vaults, we can track the events linked to the strategies to be able to
** get the corresponding data. Events are not the same for all the vaults versions, but should
** have some common data.
** Not all events are required and the minimal configuration to get the vaults is the
** StrategyAdded event.
**
** Other possible events:
** StrategyReported | StrategyUpdateDebtRatio | StrategyUpdateMinDebtPerHarvest |
** StrategyUpdateMaxDebtPerHarvest | StrategyMigrated | StrategyRevoked |
** StrategyUpdatePerformanceFee | StrategyRemovedFromQueue | StrategyAddedToQueue
**********************************************************************************************/
func RetrieveAllStrategiesAdded(
	chainID uint64,
	vaults map[common.Address]utils.TVaultsFromRegistry,
) []models.TStrategyAdded {
	trace := traces.Init(`app.indexer.strategies.activation_events`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `strategies`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** We will then listen to all events related to the strategies added or migrated to the vaults.
	** We will use a sync.Map to store the events and a WaitGroup to wait for all the goroutines.
	**********************************************************************************************/
	asyncStrategiesForVaults := &sync.Map{}
	asyncStrategiesMigrated := &sync.Map{}

	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(2)
		go getStrategiesAdded(
			chainID,
			v.VaultsAddress,
			v.Activation,
			v.APIVersion,
			asyncStrategiesForVaults,
			wg,
		)
		go getStrategiesMigrated(
			chainID,
			v.VaultsAddress,
			v.Activation,
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
		valueParsed.VaultVersion = vaults[vaultAddressParsed].APIVersion
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
		newStrategy.VaultVersion = vaults[vaultAddressParsed].APIVersion
		strategies[vaultAddressParsed][newStrategyAddressParsed] = newStrategy
		allStrategiesList = append(allStrategiesList, newStrategy)
		count++
		return true
	})
	return allStrategiesList
}
