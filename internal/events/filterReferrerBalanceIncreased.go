package events

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** Filter all the RewardAdded events for a the yBribe contract.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - asyncRewardAdded: the async ptr to the map of blockNumber -> TEventReferredBalanceIncreased
**
** Returns nothing as the asyncRewardAdded is updated via a pointer
**************************************************************************************************/
func filterReferrerBalanceIncreased(chainID uint64, fromBlock uint64, toBlock *uint64, asyncRewardAdded *sync.Map) {
	client := ethereum.GetRPC(chainID)
	partnerContract := env.PARTNER_TRACKERS_ADDRESSES[chainID]
	partnerContractAddress := partnerContract.Address
	currentVault, _ := contracts.NewYPartnerTracker(partnerContractAddress, client)

	//Note: we don't use fromBlock here because we want the full history previous to the toblock
	//to ensure the balance is correct
	opts := &bind.FilterOpts{Start: partnerContract.Block, End: toBlock}

	if log, err := currentVault.FilterReferredBalanceIncreased(opts, nil, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			asyncRewardAdded.Store(log.Event.Raw.BlockNumber, models.TEventReferredBalanceIncreased{
				Amount:         bigNumber.SetInt(log.Event.AmountAdded),
				TotalDeposited: bigNumber.SetInt(log.Event.TotalDeposited),
				PartnerID:      log.Event.PartnerId,
				Vault:          log.Event.Vault,
				Depositer:      log.Event.Depositer,
				TxHash:         log.Event.Raw.TxHash,
				BlockNumber:    log.Event.Raw.BlockNumber,
				TxIndex:        log.Event.Raw.TxIndex,
				LogIndex:       log.Event.Raw.Index,
			})
		}
	}
}

/**********************************************************************************************
** HandleRefererBalanceIncrease will retrieve all the RewardAdded events for the yBribeV3 contract
** In order to get the list, or a feed, of all the RewardAdded events, we need to filter the
** events from the blockchain and store them in a map. This function will do that.
**********************************************************************************************/
func HandleRefererBalanceIncrease(chainID uint64, fromBlock uint64, toBlock *uint64) map[uint64]models.TEventReferredBalanceIncreased {
	/**********************************************************************************************
	** Concurrently retrieve all first updateManagement events, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	asyncRewardAdded := sync.Map{}
	filterReferrerBalanceIncreased(chainID, fromBlock, toBlock, &asyncRewardAdded)

	/**********************************************************************************************
	** Once we got all the reward added blocks, we need to extract them from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: blockNumber
	** - value: TEventReferredBalanceIncreased
	**
	** We need to update, for each corresponding event, the activation block to the TEventReferredBalanceIncreased.
	**********************************************************************************************/
	count := 0
	rewardAddedMap := make(map[uint64]models.TEventReferredBalanceIncreased)
	asyncRewardAdded.Range(func(key, value interface{}) bool {
		event := value.(models.TEventReferredBalanceIncreased)
		rewardAddedMap[key.(uint64)] = event
		count++
		return true
	})
	return rewardAddedMap
}
