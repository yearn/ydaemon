package partnerTracker

import "github.com/yearn/ydaemon/internal/models"

/**********************************************************************************************
** Set of functions to store and retrieve the referralBalances from the cache and/or database
** and being able to access them from the rest of the application.
** The _partnerReferralBalanceEvent variable is not exported and is only used internally by the
** functions below.
** The _partnerReferralBalanceEvent variable is a map of:
** chainID -> blockNumber -> TEventReferredBalanceIncreased
**********************************************************************************************/
var _partnerReferralBalanceEvent = make(map[uint64]map[uint64]*models.TEventReferredBalanceIncreased)

/**********************************************************************************************
** ListReferralBalanceIncrease will, for a given chainID, return the list of all the events
** stored in _partnerReferralBalanceEvent.
**********************************************************************************************/
func ListReferralBalanceIncrease(chainID uint64) []*models.TEventReferredBalanceIncreased {
	var events []*models.TEventReferredBalanceIncreased
	for _, event := range _partnerReferralBalanceEvent[chainID] {
		events = append(events, event)
	}
	return events
}

/**********************************************************************************************
** SetInReferralBalanceIncreaseMap will allow us to safely populate the
** _partnerReferralBalanceEvent variable.
**********************************************************************************************/
func SetInReferralBalanceIncreaseMap(chainID uint64, blockNumber uint64, event *models.TEventReferredBalanceIncreased) {
	if _, ok := _partnerReferralBalanceEvent[chainID]; !ok {
		_partnerReferralBalanceEvent[chainID] = make(map[uint64]*models.TEventReferredBalanceIncreased)
	}
	_partnerReferralBalanceEvent[chainID][blockNumber] = event
}
