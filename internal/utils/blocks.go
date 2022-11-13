package utils

import (
	"errors"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
)

type TEventBlock struct {
	EventType   string
	TxHash      common.Hash
	BlockNumber uint64
	TxIndex     uint
	LogIndex    uint
	Value       *bigNumber.Int
}

/**************************************************************************************************
** findInBlock try to find the event in the block that matches the txIndex and logIndex.
** All events are in the same block. We need to find the one before or equal to the provided
** txIndex and, if the txIndex is the same, the one before or equal to the provided logIndex.
**************************************************************************************************/
func findInBlock(blocks []TEventBlock, lookingForTxIndex uint, lookingForLogIndex uint) (TEventBlock, error) {
	/**********************************************************************************************
	** Blocks are not ordered by txIndex and logIndex, so we need to sort them first while removing
	** all events that are after the one we are looking for (txIndex > lookingForTxIndex)
	**********************************************************************************************/
	blockEvents := []TEventBlock{}
	for _, block := range blocks {
		if block.TxIndex > lookingForTxIndex {
			continue
		}
		blockEvents = append(blockEvents, block)
	}

	/**********************************************************************************************
	** If we have no events left, we can't find the one we are looking for
	**********************************************************************************************/
	if len(blockEvents) == 0 {
		return TEventBlock{}, errors.New("no event found")
	}

	/**********************************************************************************************
	** Sort the events by txIndex and logIndex, with the highest txIndex first and the highest
	** logIndex first
	**********************************************************************************************/
	sort.Slice(blockEvents, func(i, j int) bool {
		if blockEvents[i].TxIndex == blockEvents[j].TxIndex {
			return blockEvents[i].LogIndex > blockEvents[j].LogIndex
		}
		return blockEvents[i].TxIndex > blockEvents[j].TxIndex
	})

	for _, event := range blockEvents {
		if event.TxIndex < lookingForTxIndex {
			return event, nil
		} else if event.TxIndex == lookingForTxIndex && event.LogIndex <= lookingForLogIndex {
			return event, nil
		}
	}
	return TEventBlock{}, errors.New(`no previous block found`)
}

/**************************************************************************************************
** analyzeBlocks will, for a given TEventBlock mapping, try to find the block just before the block
** that matches the blockNumber, txIndex and logIndex of the event. If it finds it, it will return
** it. If it doesn't find it, it will return an empty TEventBlock.
**************************************************************************************************/
func analyzeBlocks(m map[uint64][]TEventBlock, lookingForBlock uint64, lookingForTxIndex uint, lookingForLogIndex uint) TEventBlock {
	/**********************************************************************************************
	** First step is to order all blocks by blockNumber, from the highest to the lowest (last
	** first) and remove all thooses that are higher than the blockNumber of the event.
	**********************************************************************************************/
	blockNumbers := make([]int, 0, len(m))
	for k := range m {
		if k <= lookingForBlock {
			blockNumbers = append(blockNumbers, int(k))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(blockNumbers)))
	if len(blockNumbers) == 0 {
		return TEventBlock{}
	}

	/**********************************************************************************************
	** Then, we have two cases:
	** 1. The blockNumber of the event is in the map. In this case, we need to find, in the list of
	**    events of this block, the one with a txIndex just before the one of the event then a
	**    logIndex just before the one of the event.
	** 2. The blockNumber of the event is not in the map. In this case, we need to find the block
	**    just before the one of the event.
	**********************************************************************************************/
	if blockNumbers[0] == int(lookingForBlock) {
		//case 1
		blocks := m[lookingForBlock]
		block, err := findInBlock(blocks, lookingForTxIndex, lookingForLogIndex)
		if err != nil {
			return TEventBlock{}
		}
		return block
	} else {
		//case 2
		blocks := m[uint64(blockNumbers[0])]
		block, err := findInBlock(blocks, ^uint(0), ^uint(0))
		if err != nil {
			helpers.LogAndCaptureError(err, `NOT FOUND`)
			return TEventBlock{}
		}
		return block
	}
}

/**************************************************************************************************
** FindEventBefore will, for a given TEventBlock mapping, try to find the value amount to use for
** this block. This is done by analyzing all fee change event and determining which one is the
** most recent one based on the event sent as parameter.
**************************************************************************************************/
func FindEventBefore(feeUpdatesBlocks map[uint64][]TEventBlock, eventBlock TEventBlock) *bigNumber.Int {
	previousBlocks := analyzeBlocks(feeUpdatesBlocks, eventBlock.BlockNumber, eventBlock.TxIndex, eventBlock.LogIndex)
	return previousBlocks.Value
}

/**************************************************************************************************
** FindPreviousBlock will, for a given blockNumber mapping, try to find the blockNumber just before
** and return it. If it can't find it, it will return 0.
**************************************************************************************************/
func FindPreviousBlock(blocks map[uint64]uint64, blockNumber uint64) (previous uint64) {
	keys := make([]int, 0, len(blocks))
	for k := range blocks {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		if k >= int(blockNumber) {
			return
		}
		previous = blocks[uint64(k)]
	}
	return
}
