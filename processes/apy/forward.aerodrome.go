package apy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

var AERO_STAKING_POOLS_REGISTRY = common.HexToAddress(`0x16613524e02ad97eDfeF371bC883F2F5d6C480A5`)

/**************************************************************************************************
** Check if the vault is a aerodrome vault. In order to check this, we need to check if the token
** has a staking pool in aerodrome
**************************************************************************************************/
func isAeroVault(chainID uint64, vault models.TVault) (common.Address, bool) {
	if chainID != 8453 {
		return common.Address{}, false
	}

	aeroVoter, err := contracts.NewAerodromeVoterRegistryCaller(AERO_STAKING_POOLS_REGISTRY, ethereum.GetRPC(chainID))
	if err != nil {
		logs.Error(err)
		return common.Address{}, false
	}
	gaugeAddressForVoter, err := aeroVoter.Gauges(nil, vault.AssetAddress)
	if err != nil {
		logs.Error(err)
		return common.Address{}, false
	}

	return gaugeAddressForVoter, gaugeAddressForVoter != common.Address{}
}
