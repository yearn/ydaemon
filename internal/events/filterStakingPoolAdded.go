package events

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
)

type TStakingPoolAdded struct {
	StackingPool common.Address
	Token        common.Address
}

/**************************************************************************************************
** Filter all stakingPoolAdded events and store them in a map of eventKey -> TStackingPoolAdded
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - opts: the filter options
** - asyncMap: the async ptr to the map of tokenAddr -> stackingPoolAddr -> block
**
** Returns nothing as the asyncMap is updated via a pointer
**************************************************************************************************/
func filterStakingPoolAdded(chainID uint64, start uint64, end *uint64, asyncMap *sync.Map) {
	client := ethereum.GetRPC(chainID)
	stackingReward, ok := env.STACKING_REWARD_ADDRESSES[chainID]
	contract, err := contracts.NewYOptimismStakingReward(stackingReward.Address, client)
	if err != nil || !ok {
		logs.Error("Error while fetching StakingPoolAdded event", err)
		return
	}
	if start == 0 || start < stackingReward.Block {
		start = stackingReward.Block
	}

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
	** Then, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
	** avoid hitting some external node providers' rate limits.
	** Note: we don't use start here because we want the full history previous to the end
	** to ensure the balance is correct
	******************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[chainID] {
		chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[chainID]
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: start, End: &chunkEnd}
		if log, err := contract.FilterStakingPoolAdded(opts, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					logs.Error("Error while fetching StakingPoolAdded event", log.Error())
					continue
				}

				eventKey := log.Event.Token.Hex() + `-` + log.Event.StakingPool.Hex() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10)
				asyncMap.Store(eventKey, TStakingPoolAdded{
					StackingPool: log.Event.StakingPool,
					Token:        log.Event.Token,
				})
			}
		} else {
			logs.Error(`impossible to FilterStakingPoolAdded for StakingContract ` + stackingReward.Address.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** In order to compute the risk score, we need to know when a new Stacking Pool is added to the
** registry. This function will fetch all the events and store them.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
**
** Returns:
** - a map of vaultAddress -> strategyAddress -> blockNumber -> PerformanceFee
**************************************************************************************************/
func HandleStakingPoolAdded(chainID uint64, start uint64, end *uint64) []TStakingPoolAdded {
	timeBefore := time.Now()
	asyncStackingPoolAddedMap := sync.Map{}

	filterStakingPoolAdded(
		chainID,
		start,
		end,
		&asyncStackingPoolAddedMap,
	)

	/**********************************************************************************************
	** Once we have all the stackingPools added we need to extract them from the sync.Map in a
	** simple array, as we just need to iterate over them.
	**
	** The syncMap variable is setup as follows:
	** - key: tokenAddress-poolAddress-blockNumber
	** - value: TStakingPoolAdded
	**
	** We need to transform it into a map as follows:
	** []TStakingPoolAdded
	**********************************************************************************************/
	stackingPools := []TStakingPoolAdded{}
	asyncStackingPoolAddedMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		tokenAddress := common.HexToAddress(eventKey[0])
		poolAddress := common.HexToAddress(eventKey[1])

		stackingPools = append(stackingPools, TStakingPoolAdded{
			StackingPool: poolAddress,
			Token:        tokenAddress,
		})
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve the stakingPoolAdded updates`)
	return stackingPools
}
