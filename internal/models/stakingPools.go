package models

import "github.com/ethereum/go-ethereum/common"

/**************************************************************************************************
** TStakingPoolAdded contains the staking pool information for a given pool
**************************************************************************************************/
type TStakingPoolAdded struct {
	StackingPoolAddress common.Address
	VaultAddress        common.Address
}
