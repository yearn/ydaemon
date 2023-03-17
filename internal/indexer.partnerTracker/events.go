package partnerTracker

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/common/types/common"
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
func filterReferrerBalanceIncrease(
	chainID uint64,
	asyncRewardAdded *sync.Map,
) {
	client := ethereum.GetRPC(chainID)
	partnerContract := PARTNER_TRACKERS_ADDRESSES[chainID]
	partnerContractAddress := partnerContract.Address.ToAddress()
	currentVault, _ := contracts.NewYPartnerTracker(partnerContractAddress, client)
	opts := &bind.FilterOpts{
		Start: partnerContract.Block,
		End:   nil,
	}

	if log, err := currentVault.FilterReferredBalanceIncreased(opts, nil, nil, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			asyncRewardAdded.Store(log.Event.Raw.BlockNumber, TEventReferredBalanceIncreased{
				Amount:         bigNumber.SetInt(log.Event.AmountAdded),
				TotalDeposited: bigNumber.SetInt(log.Event.TotalDeposited),
				PartnerID:      common.FromAddress(log.Event.PartnerId),
				Vault:          common.FromAddress(log.Event.Vault),
				Depositer:      common.FromAddress(log.Event.Depositer),
				Timestamp:      ethereum.GetBlockTime(chainID, log.Event.Raw.BlockNumber),
				TxHash:         log.Event.Raw.TxHash,
				BlockNumber:    log.Event.Raw.BlockNumber,
				TxIndex:        log.Event.Raw.TxIndex,
				LogIndex:       log.Event.Raw.Index,
			})
		}
	} else {
		traces.
			Capture(`error`, `impossible to FilterReferrerBalanceIncrease for YBribeV3 `+partnerContractAddress.Hex()).
			SetEntity(`bribes`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
			SetTag(`bribeAddress`, partnerContractAddress.Hex()).
			Send()
	}
}

/**********************************************************************************************
** In order to get the list, or a feed, of all the RewardAdded events, we need to filter the
** events from the blockchain and store them in a map. This function will do that.
**********************************************************************************************/
func RetrieveAllReffererBalanceIncrease(chainID uint64) map[uint64]TEventReferredBalanceIncreased {
	trace := traces.Init(`app.indexer.partnertracker.referred_balance_increased`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `partnerTracker`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** Concurrently retrieve all first updateManagement events, waiting for the end of all
	** goroutines via the wg before continuing.
	**********************************************************************************************/
	asyncRewardAdded := sync.Map{}
	filterReferrerBalanceIncrease(chainID, &asyncRewardAdded)

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
	rewardAddedMap := make(map[uint64]TEventReferredBalanceIncreased)
	asyncRewardAdded.Range(func(key, value interface{}) bool {
		event := value.(TEventReferredBalanceIncreased)
		rewardAddedMap[key.(uint64)] = event
		SetInReferralBalanceIncreaseMap(chainID, event.BlockNumber, &event)
		count++
		return true
	})
	return rewardAddedMap
}
