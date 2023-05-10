package bribes

import "github.com/yearn/ydaemon/internal/models"

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _yBribeRewardAdded variable is not exported and is only used internally by the
** functions below.
**********************************************************************************************/
var _yBribeRewardAdded = make(map[uint64][]models.TEventRewardAdded)

/**********************************************************************************************
** ListRewardAdded will, for a given chainID, return the list of all the events stored in
** _yBribeRewardAdded.
**********************************************************************************************/
func ListRewardAdded(chainID uint64) []models.TEventRewardAdded {
	var events []models.TEventRewardAdded
	events = append(events, _yBribeRewardAdded[chainID]...)
	return events
}

/**********************************************************************************************
** SetInRewardAddedMap will allow us to safely populate the _yBribeRewardAdded variable.
**********************************************************************************************/
func SetInRewardAddedMap(chainID uint64, event models.TEventRewardAdded) {
	if _, ok := _yBribeRewardAdded[chainID]; !ok {
		_yBribeRewardAdded[chainID] = []models.TEventRewardAdded{}
	}
	_yBribeRewardAdded[chainID] = append(_yBribeRewardAdded[chainID], event)
}
