package events

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/helpers"
)

var treasuryChangesForChain = make(map[uint64]map[common.Address]map[uint64]common.Address)
var defaultTreasury = common.HexToAddress(`0x93a62da5a14c80f265dabc077fcee437b1a0efde`)

/**********************************************************************************************
** StoreUpdateRewards is used to store the treasury changes for a given chain in a reusable
** map
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - updates: a map of vaultAddress -> blockNumber -> treasuryAddress
**********************************************************************************************/
func StoreUpdateRewards(chainID uint64, updates map[common.Address]map[uint64]common.Address) {
	treasuryChangesForChain[chainID] = updates
}

/**************************************************************************************************
** This function is used to find the treasury address for a given vault at a particular block
** number on a specified chain.
** Parameters:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: address of the vault for which the treasury is being searched
** - blockNumber: the blockNumber at which the treasury is to be retrieved.
**
** The function first initializes a default treasury address. If there are no changes to the
** treasury for the given chainID, the default treasury address is returned. If there are no
** changes to the treasury for the given vaultAddress, the default treasury address is returned.
** Otherwise, the function iterates over the treasury changes for the given vaultAddress and checks
** if the block number for the current treasury is less than or equal to the given blockNumber. If
** it is, the treasury variable is updated with the current treasury address. The function then
** returns the updated treasury address.
**************************************************************************************************/
func FindTreasuryAtBlock(chainID uint64, vaultAddress common.Address, blockNumber uint64) common.Address {
	treasury := defaultTreasury
	if _, ok := treasuryChangesForChain[chainID]; !ok {
		return treasury
	}
	if _, ok := treasuryChangesForChain[chainID][vaultAddress]; !ok {
		return treasury
	}
	for block, treasuryForBlock := range treasuryChangesForChain[chainID][vaultAddress] {
		if block <= blockNumber {
			treasury = treasuryForBlock
		}
	}
	return treasury
}

/**************************************************************************************************
** This function is used to find all of the treasury addresses that have been used for a given
** vault on a specified chain.
** Parameters:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: address of the vault for which the treasuries are being searched
**
** The function first initializes a default treasury address with the vaultAddress. If there are
** no changes to the treasury for the given chainID, the default treasury address is returned. If
** there are no changes to the treasury for the given vaultAddress, the default treasury address
** is returned. Otherwise, the function iterates over the treasury changes for the given
** vaultAddress and adds each treasury address to the treasuries list if it is not already in the
** list. The function then returns the list of updated treasuries.
**************************************************************************************************/
func FindTreasuriesForVault(chainID uint64, vaultAddress common.Address) []common.Address {
	treasuries := []common.Address{vaultAddress}
	if _, ok := treasuryChangesForChain[chainID]; !ok {
		return treasuries
	}
	if _, ok := treasuryChangesForChain[chainID][vaultAddress]; !ok {
		return treasuries
	}
	for _, treasuryForBlock := range treasuryChangesForChain[chainID][vaultAddress] {
		if !helpers.Contains(treasuries, treasuryForBlock) {
			treasuries = append(treasuries, treasuryForBlock)
		}
	}
	return treasuries
}

/**************************************************************************************************
** This function is used to check if the treasury address for a given vault at a particular block
** number on a specified chain is the same as the vault address.
** Parameters:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: address of the vault for which the treasury is being checked
** - blockNumber: the blockNumber at which the treasury is to be checked.
**
** The function first calls the FindTreasuryAtBlock function to get the treasury address at the
** specified block. The function then checks if the retrieved treasury address is the same as the
** vault address. If it is, the function returns true, otherwise it returns false.
**************************************************************************************************/
func IsTreasuryAtBlock(chainID uint64, vaultAddress common.Address, blockNumber uint64) bool {
	treasury := FindTreasuryAtBlock(chainID, vaultAddress, blockNumber)
	return treasury == vaultAddress
}
