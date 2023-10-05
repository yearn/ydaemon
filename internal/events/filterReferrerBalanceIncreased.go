package events

import (
	"context"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** Filter all the RewardAdded events for a the yBribe contract.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - syncMap: the async ptr to the map of blockNumber -> TEventReferredBalanceIncreased
**
** Returns nothing as the syncMap is updated via a pointer
**************************************************************************************************/
func filterReferrerBalanceIncreased(chainID uint64, start uint64, end *uint64, syncMap *sync.Map) {
	client := ethereum.GetRPC(chainID)
	partnerContract := env.CHAINS[chainID].PartnerContract
	partnerContractAddress := partnerContract.Address
	currentPartnerContract, _ := contracts.NewYPartnerTracker(partnerContractAddress, client)

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
	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[chainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[chainID].MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
		}
		opts := &bind.FilterOpts{Start: partnerContract.Block, End: &chunkEnd}
		if log, err := currentPartnerContract.FilterReferredBalanceIncreased(opts, nil, nil, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				syncMap.Store(log.Event.Raw.BlockNumber, models.TEventReferredBalanceIncreased{
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
		} else {
			logs.Error(`impossible to FilterReferredBalanceIncreased for PartnerContract ` + partnerContractAddress.Hex() + ` on chain ` + strconv.FormatUint(chainID, 10) + `: ` + err.Error())
		}
	}

}

/**********************************************************************************************
** HandleRefererBalanceIncrease will retrieve all the RewardAdded events for the yBribeV3 contract
** In order to get the list, or a feed, of all the RewardAdded events, we need to filter the
** events from the blockchain and store them in a map. This function will do that.
**********************************************************************************************/
func HandleRefererBalanceIncrease(chainID uint64, start uint64, end *uint64) map[uint64]models.TEventReferredBalanceIncreased {
	/**********************************************************************************************
	** Concurrently retrieve all first updateManagement events, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	syncMap := sync.Map{}
	filterReferrerBalanceIncreased(chainID, start, end, &syncMap)

	/**********************************************************************************************
	** Once we got all the reward added blocks, we need to extract them from the sync.Map.
	** The syncMap variable is setup as follows:
	** - key: blockNumber
	** - value: TEventReferredBalanceIncreased
	**
	** We need to update, for each corresponding event, the activation block to the
	** TEventReferredBalanceIncreased.
	**********************************************************************************************/
	count := 0
	rewardAddedMap := make(map[uint64]models.TEventReferredBalanceIncreased)
	syncMap.Range(func(key, value interface{}) bool {
		event := value.(models.TEventReferredBalanceIncreased)
		rewardAddedMap[key.(uint64)] = event
		count++
		return true
	})
	return rewardAddedMap
}
