package indexer

import (
	"context"
	"math/big"
	"strconv"
	"sync"
	"time"

	goEth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

func filterStakingPools(
	chainID uint64,
	registry env.TContractData,
	start uint64,
	end *uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (lastBlock uint64) {
	/**********************************************************************************************
	** As we cannot use the WS connection, we will fallback to the regular RPC connection. This
	** means that we will need to filter the logs from the lastSyncedBlock to the latest block
	** every x seconds to check if there are new vaults.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[chainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[chainID].MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
			lastBlock = chunkEnd
		}

		if chunkEnd >= *end && !isDone && wg != nil {
			wg.Done()
		}
		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}

		currentRegistry, _ := contracts.NewYOptimismStakingRewardRegistry(registry.Address, client)
		if log, err := currentRegistry.FilterStakingPoolAdded(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				storage.StoreStakingPool(chainID, models.TStakingPoolAdded{
					StackingPoolAddress: log.Event.StakingPool,
					VaultAddress:        log.Event.Token,
				})
			}
		} else {
			logs.Error(`impossible to FilterStakingPools for YRegistryV4 ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	}
	return lastBlock
}

/**************************************************************************************************
** Watch for new vaults added to the registry. This function is called by the infinite loop in
** indexStakingPoolWrapper. It uses the WS connection to listen to the events.
** It will catch the new vaults and start all the subsequent processes to add them correctly in
** yDaemon.
** On errror, it will return the lastSyncedBlock, a boolean to indicate if we should retry and the
** error.
**************************************************************************************************/
func watchStakingPool(
	chainID uint64,
	registry env.TContractData,
	lastSyncedBlock uint64,
	wg *sync.WaitGroup,
	isDone bool,
) (uint64, bool, error) {
	/**********************************************************************************************
	** First thing is to connect to the node via a WS connection. We need to use a WS connection
	** because we need to listen to new events as they are emitted via the node.
	** Not all nodes support WS connections, so we need to check if the node supports it.
	**********************************************************************************************/
	client, err := ethereum.GetWSClient(chainID)
	if err != nil {
		if wg != nil && !isDone {
			wg.Done()
		}
		return 0, false, err
	}

	currentRegistry, _ := contracts.NewYOptimismStakingRewardRegistry(registry.Address, client)
	etherReader := ethereum.Reader{Backend: client}
	contractABI, _ := contracts.YOptimismStakingRewardRegistryMetaData.GetAbi()
	topics, _ := abi.MakeTopics([][]interface{}{{contractABI.Events[`StakingPoolAdded`].ID}}...)
	query := goEth.FilterQuery{
		FromBlock: big.NewInt(int64(registry.Block)),
		Addresses: []common.Address{registry.Address},
		Topics:    topics,
	}
	stream, sub, history, err := etherReader.QueryWithHistory(context.Background(), &query)
	if err != nil {
		logs.Error(err)
		if wg != nil && !isDone {
			wg.Done()
		}
		return 0, false, err
	}
	defer sub.Unsubscribe()

	/** 🔵 - Yearn *************************************************************************************
	** Handle historical events
	**************************************************************************************************/
	for _, log := range history {
		value, err := currentRegistry.ParseStakingPoolAdded(log)
		if err != nil {
			continue
		}
		storage.StoreStakingPool(chainID, models.TStakingPoolAdded{
			StackingPoolAddress: value.StakingPool,
			VaultAddress:        value.Token,
		})
	}
	if wg != nil && !isDone {
		wg.Done()
	}

	/**********************************************************************************************
	** Listen and handle new events
	**********************************************************************************************/
	for {
		select {
		case log := <-stream:
			value, err := currentRegistry.ParseStakingPoolAdded(log)
			if err != nil {
				continue
			}
			lastSyncedBlock = value.Raw.BlockNumber
			storage.StoreStakingPool(chainID, models.TStakingPoolAdded{
				StackingPoolAddress: value.StakingPool,
				VaultAddress:        value.Token,
			})
		case err := <-sub.Err():
			logs.Error(err)
			return lastSyncedBlock, true, err
		}
	}
}

/**************************************************************************************************
** As WS connections can be lost, can be instable or can just not be supported by the node, we need
** to have a fallback to the regular RPC connection.
** This function is called by the infinite loop and will try to index the new vaults using the WS
** but will also provide a way to catch the new vaults via the RPC connection.
**************************************************************************************************/
func indexStakingPoolWrapper(
	chainID uint64,
	registry env.TContractData,
	higherIndexedBlockNumber uint64,
	wg *sync.WaitGroup,
) {
	isDone := false
	lastSyncedBlock := higherIndexedBlockNumber

	/**********************************************************************************************
	** Initialize the infinite loop to listen to new events. This is a wrapper around the actual
	** indexer to be able to catch the errors, to restart the indexer or just to exit the loop to
	** fallback to another method.
	**********************************************************************************************/
	shouldRetry := true
	for {
		/******************************************************************************************
		** Just checking if the connection is alive, if not, we will fallback to the filter method.
		** We must close the client we openned.
		******************************************************************************************/
		var err error
		if _, err := ethereum.GetWSClient(chainID); err != nil {
			/**********************************************************************************************
			** Default method: use the RPC connection to filter the events from the lastSyncedBlock to the
			** latest block. This is a fallback method in case the WS connection is not available.
			** It's only triggered if delay is greater than 0, allowing us to try to retry this whole
			** function under certain conditions while avoiding multiple calls to the same function.
			**********************************************************************************************/
			for {
				lastBlock := filterStakingPools(
					chainID,
					registry,
					lastSyncedBlock,
					nil,
					wg,
					isDone,
				)
				isDone = true
				lastSyncedBlock = lastBlock
				time.Sleep(1 * time.Minute)
			}
		}

		lastSyncedBlock, shouldRetry, err = watchStakingPool(
			chainID,
			registry,
			lastSyncedBlock,
			wg,
			isDone,
		)
		isDone = true
		if err != nil {
			logs.Error(`error while indexing staking pools from registry ` + registry.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
		if shouldRetry {
			time.Sleep(30 * time.Second)
			continue
		}
		break
	}
}

/** 🔵 - Yearn *************************************************************************************
** The function `IndexStakingPools` ...
**************************************************************************************************/
func IndexStakingPools(chainID uint64) (stakingPoolsFromRegistry map[common.Address]models.TStakingPoolAdded) {
	registry := env.CHAINS[chainID].StackingRewardContract
	if (registry.Address == common.Address{}) {
		return
	}
	wg := sync.WaitGroup{}

	/** 🔵 - Yearn *****************************************************************************
	** After retrieving the highest block number we can proceed to index new pools.
	******************************************************************************************/
	wg.Add(1)
	go indexStakingPoolWrapper(chainID, registry, registry.Block, &wg)
	wg.Wait()

	/** 🔵 - Yearn *********************************************************************************
	** This part is only accessed once, once the `historical` pools, aka that are already deployed
	** and indexed, are retrieved. New deployed pools will not hit this part, but will be handle
	** in the above functions directly
	**********************************************************************************************/
	stakingPoolsFromRegistry, _ = storage.ListStakingPools(chainID, storage.PerPool)
	return stakingPoolsFromRegistry
}
