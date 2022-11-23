package strategies

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/utils"
)

var alreadyRunningIndexers = make(map[uint64]map[ethcommon.Address]bool)

/**************************************************************************************************
** Watch for the StrategyAdded events in each vaults. This function is called by the infinite loop
** in indexStrategyAddedWrapper. It uses the WS connection to listen to the events.
** It will catch the strategies added and start all the subsequent processes to add them correctly
** in yDaemon. One process should be open for each vault.
** On error, it will return the lastSyncedBlock, a boolean to indicate if we should retry and the
** error.
**
** This function is splitted in multiple functions to be able to handle the different versions of
** the vaults, which have different events.
** Arguments:
** - chainID: the chain ID of the network we are listening to
** - vaultAddress: the address of the vault we are listening to
** - vaultVersion: the version of the vault we are listening to
** - lastSyncedBlock: the last block we have synced
**
** Returns:
** - lastSyncedBlock: the last block we have synced
** - shouldRetry: a boolean to indicate if we should retry
** - err: the error if any
**************************************************************************************************/
func indexStrategyAddedV022(
	chainID uint64,
	vaultAddress ethcommon.Address,
	vaultVersion string,
	lastSyncedBlock uint64,
) (uint64, bool, error) {
	/**********************************************************************************************
	** First thing is to connect to the node via a WS connection. We need to use a WS connection
	** because we need to listen to new events as they are emitted via the node.
	** Not all nodes support WS connections, so we need to check if the node supports it.
	**********************************************************************************************/
	client, err := ethereum.GetWSClient(chainID)
	if err != nil {
		return lastSyncedBlock, true, err
	}

	/**********************************************************************************************
	** We will need to store the current BlockNumber to be able to filter the events from that
	** specific block number. It's a security measure to avoid missing events. If the node goes
	** down or the WS connection is lost, we can just proceed with the regular RPC connection by
	** filtering logs from that block to the latest block.
	**********************************************************************************************/
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		return lastSyncedBlock, true, err
	}

	if currentBlock > lastSyncedBlock {
		currentBlock = lastSyncedBlock
	}

	/**********************************************************************************************
	** Connect to the Yearn Registry with general configuration, starting from the lastSyncedBlock.
	** No error should happen here, but if it does, we just return.
	**********************************************************************************************/
	currentVault, _ := contracts.NewYvault022(vaultAddress, client)
	channel := make(chan *contracts.Yvault022StrategyAdded)
	watcher, err := currentVault.WatchStrategyAdded(
		&bind.WatchOpts{Start: &currentBlock},
		channel,
		nil,
	)
	if err != nil {
		return currentBlock, false, err
	}

	/**********************************************************************************************
	** Listen to the channel forever, and handle the new events as they are emitted. Once an event
	** is received, we also update the lastSyncedBlock to be able to filter the next events from
	** that block.
	** On error, fallback to the default setup
	**********************************************************************************************/
	for {
		select {
		case log := <-channel:
			lastSyncedBlock = log.Raw.BlockNumber
			newStrategyAdded := TStrategyAdded{
				VaultAddress:    vaultAddress,
				StrategyAddress: log.Strategy,
				TxHash:          log.Raw.TxHash,
				BlockNumber:     log.Raw.BlockNumber,
				TxIndex:         log.Raw.TxIndex,
				LogIndex:        log.Raw.Index,
				DebtLimit:       bigNumber.SetInt(log.DebtLimit),
				RateLimit:       bigNumber.SetInt(log.RateLimit),
				PerformanceFee:  bigNumber.SetInt(log.PerformanceFee),
			}
			logs.Info(`New Strategy Added V0.2.2 detected by indexer!`)
			logs.Pretty(newStrategyAdded)

			processNewStrategies(chainID, []TStrategyAdded{newStrategyAdded})
		case err := <-watcher.Err():
			return lastSyncedBlock, true, err
		}
	}
}

func indexStrategyAddedV030(
	chainID uint64,
	vaultAddress ethcommon.Address,
	vaultVersion string,
	lastSyncedBlock uint64,
) (uint64, bool, error) {
	/**********************************************************************************************
	** First thing is to connect to the node via a WS connection. We need to use a WS connection
	** because we need to listen to new events as they are emitted via the node.
	** Not all nodes support WS connections, so we need to check if the node supports it.
	**********************************************************************************************/
	client, err := ethereum.GetWSClient(chainID)
	if err != nil {
		return lastSyncedBlock, true, err
	}

	/**********************************************************************************************
	** We will need to store the current BlockNumber to be able to filter the events from that
	** specific block number. It's a security measure to avoid missing events. If the node goes
	** down or the WS connection is lost, we can just proceed with the regular RPC connection by
	** filtering logs from that block to the latest block.
	**********************************************************************************************/
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		return lastSyncedBlock, true, err
	}

	if currentBlock > lastSyncedBlock {
		currentBlock = lastSyncedBlock
	}

	/**********************************************************************************************
	** Connect to the Yearn Registry with general configuration, starting from the lastSyncedBlock.
	** No error should happen here, but if it does, we just return.
	**********************************************************************************************/
	currentVault, _ := contracts.NewYvault030(vaultAddress, client)
	channel := make(chan *contracts.Yvault030StrategyAdded)
	watcher, err := currentVault.WatchStrategyAdded(
		&bind.WatchOpts{Start: &currentBlock},
		channel,
		nil,
	)
	if err != nil {
		return currentBlock, false, err
	}

	/**********************************************************************************************
	** Listen to the channel forever, and handle the new events as they are emitted. Once an event
	** is received, we also update the lastSyncedBlock to be able to filter the next events from
	** that block.
	** On error, fallback to the default setup
	**********************************************************************************************/
	for {
		select {
		case log := <-channel:
			lastSyncedBlock = log.Raw.BlockNumber
			newStrategyAdded := TStrategyAdded{
				VaultAddress:    vaultAddress,
				StrategyAddress: log.Strategy,
				TxHash:          log.Raw.TxHash,
				BlockNumber:     log.Raw.BlockNumber,
				TxIndex:         log.Raw.TxIndex,
				LogIndex:        log.Raw.Index,
				DebtRatio:       bigNumber.SetInt(log.DebtRatio),
				RateLimit:       bigNumber.SetInt(log.RateLimit),
				PerformanceFee:  bigNumber.SetInt(log.PerformanceFee),
			}
			logs.Info(`New Strategy Added V0.3.0 detected by indexer!`)
			logs.Pretty(newStrategyAdded)

			processNewStrategies(chainID, []TStrategyAdded{newStrategyAdded})
		case err := <-watcher.Err():
			return lastSyncedBlock, true, err
		}
	}
}

func indexStrategyAddedV043(
	chainID uint64,
	vaultAddress ethcommon.Address,
	vaultVersion string,
	lastSyncedBlock uint64,
) (uint64, bool, error) {
	/**********************************************************************************************
	** First thing is to connect to the node via a WS connection. We need to use a WS connection
	** because we need to listen to new events as they are emitted via the node.
	** Not all nodes support WS connections, so we need to check if the node supports it.
	**********************************************************************************************/
	client, err := ethereum.GetWSClient(chainID)
	if err != nil {
		return lastSyncedBlock, true, err
	}

	/**********************************************************************************************
	** We will need to store the current BlockNumber to be able to filter the events from that
	** specific block number. It's a security measure to avoid missing events. If the node goes
	** down or the WS connection is lost, we can just proceed with the regular RPC connection by
	** filtering logs from that block to the latest block.
	**********************************************************************************************/
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		return lastSyncedBlock, true, err
	}

	if currentBlock > lastSyncedBlock {
		currentBlock = lastSyncedBlock
	}

	/**********************************************************************************************
	** Connect to the Yearn Registry with general configuration, starting from the lastSyncedBlock.
	** No error should happen here, but if it does, we just return.
	**********************************************************************************************/
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	channel := make(chan *contracts.Yvault043StrategyAdded)
	watcher, err := currentVault.WatchStrategyAdded(
		&bind.WatchOpts{Start: &currentBlock},
		channel,
		nil,
	)
	if err != nil {
		return currentBlock, false, err
	}

	/**********************************************************************************************
	** Listen to the channel forever, and handle the new events as they are emitted. Once an event
	** is received, we also update the lastSyncedBlock to be able to filter the next events from
	** that block.
	** On error, fallback to the default setup
	**********************************************************************************************/
	for {
		select {
		case log := <-channel:
			lastSyncedBlock = log.Raw.BlockNumber
			newStrategyAdded := TStrategyAdded{
				VaultAddress:      vaultAddress,
				StrategyAddress:   log.Strategy,
				TxHash:            log.Raw.TxHash,
				BlockNumber:       log.Raw.BlockNumber,
				TxIndex:           log.Raw.TxIndex,
				LogIndex:          log.Raw.Index,
				DebtRatio:         bigNumber.SetInt(log.DebtRatio),
				MaxDebtPerHarvest: bigNumber.SetInt(log.MaxDebtPerHarvest),
				MinDebtPerHarvest: bigNumber.SetInt(log.MinDebtPerHarvest),
				PerformanceFee:    bigNumber.SetInt(log.PerformanceFee),
			}
			logs.Info(`New Strategy Added >= V0.3.2 detected by indexer!`)
			logs.Pretty(newStrategyAdded)

			processNewStrategies(chainID, []TStrategyAdded{newStrategyAdded})
		case err := <-watcher.Err():
			return lastSyncedBlock, true, err
		}
	}
}

/**************************************************************************************************
** As WS connections can be lost, can be instable or can just not be supported by the node, we need
** to have a fallback to the regular RPC connection.
** This function is called by the infinite loop and will try to index the strategies added using the
** WS but will also provide a way to catch the strategies added via the RPC connection.
**
** Arguments:
** - chainID: the chain ID of the network we are listening to
** - vaultAddress: the address of the vault we are listening to
** - vaultActivation: the activation block of the vault (start block to listen to events)
** - vaultVersion: the version of the vault we are listening to
** - delay: the delay between two standard RPC retries
**
** Returns nothing
**************************************************************************************************/
func indexStrategyAddedWrapper(
	chainID uint64,
	vaultAddress ethcommon.Address,
	vaultActivation uint64,
	vaultVersion string,
	delay time.Duration,
) {
	retrySleepInBetween := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		2 * time.Second,
		5 * time.Second,
		5 * time.Second,
		10 * time.Second,
		30 * time.Second,
	}

	/**********************************************************************************************
	** Initialize the infinite loop to listen to new events. This is a wrapper around the actual
	** indexer to be able to catch the errors, to restart the indexer or just to exit the loop to
	** fallback to another method.
	**********************************************************************************************/
	currentRetryIndex := 0
	lastSyncedBlock := vaultActivation
	shouldRetry := true
	err := error(nil)
	for {
		switch vaultVersion {
		case `0.2.2`:
			lastSyncedBlock, shouldRetry, err = indexStrategyAddedV022(
				chainID,
				vaultAddress,
				vaultVersion,
				lastSyncedBlock,
			)
		case `0.3.0`, `0.3.1`:
			lastSyncedBlock, shouldRetry, err = indexStrategyAddedV030(
				chainID,
				vaultAddress,
				vaultVersion,
				lastSyncedBlock,
			)
		default: //case `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
			lastSyncedBlock, shouldRetry, err = indexStrategyAddedV043(
				chainID,
				vaultAddress,
				vaultVersion,
				lastSyncedBlock,
			)
		}
		if err != nil {
			traces.
				Capture(`error`, `error while indexing StrategyAdded from registry `+vaultAddress.Hex()+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`strategy`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`vaultAddress`, vaultAddress.Hex()).
				Send()
		}
		if !shouldRetry {
			break
		}
		time.Sleep(retrySleepInBetween[currentRetryIndex])
		currentRetryIndex++
		if currentRetryIndex >= len(retrySleepInBetween) {
			currentRetryIndex = len(retrySleepInBetween) - 1
		}
	}

	/**********************************************************************************************
	** Default method: use the RPC connection to filter the events from the lastSyncedBlock to the
	** latest block. This is a fallback method in case the WS connection is not available.
	** It's only triggered if delay is greater than 0, allowing us to try to retry this whole
	** function under certain conditions while avoiding multiple calls to the same function.
	**********************************************************************************************/
	if delay > 0 {
		for {
			asyncStrategiesForVaults := &sync.Map{}
			strategiesAdded := getStrategiesAdded(
				chainID,
				vaultAddress,
				lastSyncedBlock,
				vaultVersion,
				asyncStrategiesForVaults,
				nil,
			)
			processNewStrategies(chainID, strategiesAdded)

			if shouldRetry {
				indexStrategyAddedWrapper(chainID, vaultAddress, lastSyncedBlock, vaultVersion, 0)
				if delay > 60 {
					time.Sleep(delay - 1*time.Minute)
				}
			} else {
				time.Sleep(delay)
			}
		}
	}
}

func IndexStrategyAdded(chainID uint64, vaultsMap map[ethcommon.Address]utils.TVaultsFromRegistry) {
	if alreadyRunningIndexers[chainID] == nil {
		alreadyRunningIndexers[chainID] = make(map[ethcommon.Address]bool)
	}

	for vaultAddress, vault := range vaultsMap {
		if alreadyRunningIndexers[chainID][vaultAddress] {
			continue
		}
		alreadyRunningIndexers[chainID][vaultAddress] = true
		go indexStrategyAddedWrapper(chainID, vaultAddress, vault.Activation, vault.APIVersion, 1*time.Minute)

	}

	logs.Success(`StrategyAdded Indexer Daemon has started. Let's wait for the first vaults to be indexed.`)
}
