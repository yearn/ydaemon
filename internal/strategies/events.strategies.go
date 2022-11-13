package strategies

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/getsentry/sentry-go"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
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
**   newStrategyAddr -> TStrategyMigrated
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

	client := ethereum.RPC[chainID]
	vault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := vault.FilterStrategyMigrated(&bind.FilterOpts{Start: vaultActivation}, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			oldAddress := log.Event.OldVersion
			newAddress := log.Event.NewVersion

			eventKey := vaultAddress.String() + `-` + oldAddress.String() + `-` + newAddress.String()
			asyncStrategiesMigration.Store(eventKey, TStrategyMigrated{
				VaultAddress:       vaultAddress,
				OldStrategyAddress: oldAddress,
				NewStrategyAddress: newAddress,
				TxHash:             log.Event.Raw.TxHash,
				BlockNumber:        log.Event.Raw.BlockNumber,
				TxIndex:            log.Event.Raw.TxIndex,
				LogIndex:           log.Event.Raw.Index,
			})
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

	client := ethereum.RPC[chainID]
	switch vaultVersion {
	case `0.2.2`:
		vault, _ := contracts.NewYvault022(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(&bind.FilterOpts{Start: vaultActivation}, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := vaultAddress.String() + `-` + log.Event.Strategy.String()
				asyncStrategiesForVaults.Store(eventKey, TStrategyAdded{
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
				eventKey := vaultAddress.String() + `-` + log.Event.Strategy.String()
				asyncStrategiesForVaults.Store(eventKey, TStrategyAdded{
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
	case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
		vault, _ := contracts.NewYvault043(vaultAddress, client)
		if log, err := vault.FilterStrategyAdded(&bind.FilterOpts{Start: vaultActivation}, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}

				eventKey := vaultAddress.String() + `-` + log.Event.Strategy.String()
				asyncStrategiesForVaults.Store(eventKey, TStrategyAdded{
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
) []TStrategyAdded {
	span := sentry.StartSpan(context.Background(), "app.fetch",
		sentry.TransactionName("Fetch Strategy Management Events"))
	span.SetTag("subsystem", "daemon")
	defer span.Finish()

	timeBefore := time.Now()

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
	** - value: []TStrategyAdded
	**
	** We need to transform it into a map as follows:
	** - vaultAddress -> strategyAddress -> []TEventBlock
	**********************************************************************************************/
	count := 0
	allStrategiesList := []TStrategyAdded{}
	strategies := make(map[common.Address]map[common.Address]TStrategyAdded)
	asyncStrategiesForVaults.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddressParsed := common.HexToAddress(eventKey[0])
		strategyAddressParsed := common.HexToAddress(eventKey[1])

		if _, ok := strategies[vaultAddressParsed]; !ok {
			strategies[vaultAddressParsed] = make(map[common.Address]TStrategyAdded)
		}
		valueParsed := value.(TStrategyAdded)
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
	** - value: []TStrategyMigrated
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
			strategies[vaultAddressParsed] = make(map[common.Address]TStrategyAdded)
		}
		oldStrategy := strategies[vaultAddressParsed][oldStrategyAddressParsed]
		newStrategy := TStrategyAdded{
			VaultAddress:      vaultAddressParsed,
			StrategyAddress:   newStrategyAddressParsed,
			TxHash:            value.(TStrategyMigrated).TxHash,
			DebtRatio:         oldStrategy.DebtRatio,
			MaxDebtPerHarvest: oldStrategy.MaxDebtPerHarvest,
			MinDebtPerHarvest: oldStrategy.MinDebtPerHarvest,
			PerformanceFee:    oldStrategy.PerformanceFee,
			DebtLimit:         oldStrategy.DebtLimit,
			RateLimit:         oldStrategy.RateLimit,
			BlockNumber:       value.(TStrategyMigrated).BlockNumber,
			TxIndex:           value.(TStrategyMigrated).TxIndex,
			LogIndex:          value.(TStrategyMigrated).LogIndex,
		}
		newStrategy.VaultVersion = vaults[vaultAddressParsed].APIVersion
		strategies[vaultAddressParsed][newStrategyAddressParsed] = newStrategy
		allStrategiesList = append(allStrategiesList, newStrategy)
		count++
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve`, len(allStrategiesList), `strategies from vaults events`)
	return allStrategiesList
}
