package events

import (
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
func filterStakingPoolAdded(
	chainID uint64,
	opts *bind.FilterOpts,
	asyncMap *sync.Map,
) {
	client := ethereum.GetRPC(chainID)

	stackingRewardAddress, ok := env.STACKING_REWARD_ADDRESSES[chainID]
	if !ok {
		logs.Error("No stacking reward address found for chainID", chainID)
		return
	}
	currentVault, err := contracts.NewYOptimismStakingReward(stackingRewardAddress, client)
	if err != nil {
		logs.Error("Error while fetching StakingPoolAdded event", err)
		return
	}

	if log, err := currentVault.FilterStakingPoolAdded(opts, nil, nil); err == nil {
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
func HandleStakingPoolAdded(
	chainID uint64,
	start uint64,
	end *uint64,
) []TStakingPoolAdded {
	timeBefore := time.Now()
	asyncStackingPoolAddedMap := sync.Map{}

	filterStakingPoolAdded(
		chainID,
		&bind.FilterOpts{Start: start, End: end},
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
