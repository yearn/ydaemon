package registries

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

/**************************************************************************************************
** Watch for new vaults added to the registry. This function is called by the infinite loop in
** indexNewVaultsWrapper. It uses the WS connection to listen to the events.
** It will catch the new vaults and start all the subsequent processes to add them correctly in
** yDaemon.
** On errror, it will return the lastSyncedBlock, a boolean to indicate if we should retry and the
** error.
**
** Arguments:
** - chainID: the chain ID of the network we are listening to
** - registryAddress: the address of the registry we are listening to
** - registryVersion: the version of the registry we are listening to
** - lastSyncedBlock: the last block we have synced
**
** Returns:
** - lastSyncedBlock: the last block we have synced
** - shouldRetry: a boolean to indicate if we should retry
** - err: the error if any
**************************************************************************************************/
func indexNewVaults(
	chainID uint64,
	registryAddress ethcommon.Address,
	registryVersion uint64,
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

	switch registryVersion {
	case 1, 2:
		/**********************************************************************************************
		** Connect to the Yearn Registry with general configuration, starting from the lastSyncedBlock.
		** No error should happen here, but if it does, we just return.
		**********************************************************************************************/
		currentVault, _ := contracts.NewYregistryv2(registryAddress, client)
		channel := make(chan *contracts.Yregistryv2NewVault)
		watcher, err := currentVault.WatchNewVault(
			&bind.WatchOpts{Start: &currentBlock},
			channel,
			nil,
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
				newVault := utils.TVaultsFromRegistry{
					RegistryAddress: registryAddress,
					VaultsAddress:   log.Vault,
					TokenAddress:    log.Token,
					APIVersion:      log.ApiVersion,
					BlockNumber:     log.Raw.BlockNumber,
					Activation:      log.Raw.BlockNumber,
					ManagementFee:   200,
					BlockHash:       log.Raw.BlockHash,
					TxIndex:         log.Raw.TxIndex,
					LogIndex:        log.Raw.Index,
					Type:            utils.VaultTypeStandard,
				}
				logs.Info(`Got vault ` + log.Vault.Hex() + ` from registry ` + registryAddress.Hex())

				newVaultList := map[ethcommon.Address]utils.TVaultsFromRegistry{
					newVault.VaultsAddress: newVault,
				}
				vaults.RetrieveActivationForAllVaults(chainID, newVaultList)
				indexer.ProcessNewVault(chainID, newVaultList)
			case err := <-watcher.Err():
				return lastSyncedBlock, true, err
			}
		}
	case 3:
		/**********************************************************************************************
		** Connect to the Yearn Registry with general configuration, starting from the lastSyncedBlock.
		** No error should happen here, but if it does, we just return.
		**********************************************************************************************/
		currentVault, _ := contracts.NewYRegistryV3(registryAddress, client)
		channel := make(chan *contracts.YRegistryV3NewVault)
		watcher, err := currentVault.WatchNewVault(
			&bind.WatchOpts{Start: &currentBlock},
			channel,
			nil,
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
				newVault := utils.TVaultsFromRegistry{
					RegistryAddress: registryAddress,
					VaultsAddress:   log.Vault,
					TokenAddress:    log.Token,
					APIVersion:      log.ApiVersion,
					BlockNumber:     log.Raw.BlockNumber,
					Activation:      log.Raw.BlockNumber,
					ManagementFee:   200,
					BlockHash:       log.Raw.BlockHash,
					TxIndex:         log.Raw.TxIndex,
					LogIndex:        log.Raw.Index,
					Type:            utils.VaultTypeStandard,
				}
				if log.VaultType.Cmp(big.NewInt(2)) == 0 {
					newVault.Type = utils.VaultTypeAutomated
				}
				logs.Info(`Got vault ` + log.Vault.Hex() + ` from registry ` + registryAddress.Hex())

				newVaultList := map[ethcommon.Address]utils.TVaultsFromRegistry{
					newVault.VaultsAddress: newVault,
				}
				vaults.RetrieveActivationForAllVaults(chainID, newVaultList)
				indexer.ProcessNewVault(chainID, newVaultList)
			case err := <-watcher.Err():
				return lastSyncedBlock, true, err
			}
		}
	}
	return lastSyncedBlock, true, err
}

/**************************************************************************************************
** As WS connections can be lost, can be instable or can just not be supported by the node, we need
** to have a fallback to the regular RPC connection.
** This function is called by the infinite loop and will try to index the new vaults using the WS
** but will also provide a way to catch the new vaults via the RPC connection.
**
** Arguments:
** - chainID: the chain ID of the network we are listening to
** - registryAddress: the address of the registry we are listening to
** - registryVersion: the version of the registry we are listening to
** - registryActivation: the activation block of the registry
** - delay: the delay between two standard RPC retries
**
** Returns nothing
**************************************************************************************************/
func indexNewVaultsWrapper(
	chainID uint64,
	registryAddress ethcommon.Address,
	registryVersion uint64,
	registryActivation uint64,
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
	lastSyncedBlock := registryActivation
	shouldRetry := true
	err := error(nil)
	for {
		lastSyncedBlock, shouldRetry, err = indexNewVaults(
			chainID,
			registryAddress,
			registryVersion,
			lastSyncedBlock,
		)
		if err != nil {
			traces.
				Capture(`error`, `error while indexing NewVault from registry `+registryAddress.Hex()+` on chain `+strconv.FormatUint(chainID, 10)).
				SetEntity(`registry`).
				SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
				SetTag(`registryAddress`, registryAddress.Hex()).
				SetExtra(`error`, err.Error()).
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
			vaultsList := []utils.TVaultsFromRegistry{}
			filterNewVaults(chainID, common.FromAddress(registryAddress), registryVersion, lastSyncedBlock, &vaultsList, nil)
			uniqueVaultsList := make(map[ethcommon.Address]utils.TVaultsFromRegistry)
			for _, v := range vaultsList {
				uniqueVaultsList[v.VaultsAddress] = v
				if v.BlockNumber > lastSyncedBlock {
					lastSyncedBlock = v.BlockNumber
				}
			}
			vaults.RetrieveActivationForAllVaults(chainID, uniqueVaultsList)
			if shouldRetry {
				indexNewVaultsWrapper(chainID, registryAddress, registryVersion, lastSyncedBlock, 0)
				if delay > 60 {
					time.Sleep(delay - 1*time.Minute)
				}
			} else {
				time.Sleep(delay)
			}
		}
	}
}

func IndexNewVaults(chainID uint64) {
	for _, registry := range YEARN_REGISTRIES[chainID] {
		go indexNewVaultsWrapper(chainID, registry.Address.ToAddress(), registry.Version, registry.Activation, 1*time.Minute)
	}

	logs.Success(`Indexer Daemon has started. Let's wait for the first vaults to be indexed.`)
}
