package partnerTracker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
)

// event ReferredBalanceIncreased(address indexed partnerId, address indexed vault, address indexed depositer, uint amountAdded, uint totalDeposited);

// TEventReferredBalanceIncreased contains the ReferredBalanceIncreased event data
// for the partner tracker contract
type TEventReferredBalanceIncreased struct {
	Amount         *bigNumber.Int `json:"amount"`
	TotalDeposited *bigNumber.Int `json:"totalDeposited"`
	PartnerID      common.Address `json:"partnerID"`
	Vault          common.Address `json:"vault"`
	Depositer      common.Address `json:"depositer"`
	TxHash         common.Hash    `json:"txHash"`
	Timestamp      uint64         `json:"timestamp"`
	BlockNumber    uint64         `json:"blockNumber"`
	TxIndex        uint           `json:"-"`
	LogIndex       uint           `json:"-"`
}

/**********************************************************************************************
** Set of functions to store and retrieve the referralBalances from the cache and/or database
** and being able to access them from the rest of the application.
** The _partnerReferralBalanceEvent variable is not exported and is only used internally by the
** functions below.
** The _partnerReferralBalanceEvent variable is a map of:
** chainID -> blockNumber -> TEventReferredBalanceIncreased
**********************************************************************************************/
var _partnerReferralBalanceEvent = make(map[uint64]map[uint64]*TEventReferredBalanceIncreased)

/**********************************************************************************************
** ListReferralBalanceIncrease will, for a given chainID, return the list of all the events
** stored in _partnerReferralBalanceEvent.
**********************************************************************************************/
func ListReferralBalanceIncrease(chainID uint64) []*TEventReferredBalanceIncreased {
	var events []*TEventReferredBalanceIncreased
	for _, event := range _partnerReferralBalanceEvent[chainID] {
		events = append(events, event)
	}
	return events
}

/**********************************************************************************************
** SetInReferralBalanceIncreaseMap will allow us to safely populate the
** _partnerReferralBalanceEvent variable.
**********************************************************************************************/
func SetInReferralBalanceIncreaseMap(chainID uint64, blockNumber uint64, event *TEventReferredBalanceIncreased) {
	if _, ok := _partnerReferralBalanceEvent[chainID]; !ok {
		_partnerReferralBalanceEvent[chainID] = make(map[uint64]*TEventReferredBalanceIncreased)
	}
	_partnerReferralBalanceEvent[chainID][blockNumber] = event
}
