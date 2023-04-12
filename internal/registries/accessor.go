package registries

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/events"
	"github.com/yearn/ydaemon/internal/utils"
	"github.com/yearn/ydaemon/internal/vaults"
)

func RegisterAllVaults(
	chainID uint64,
	start uint64,
	end *uint64,
) map[common.Address]utils.TVaultsFromRegistry {
	standardVaultList, experimentalVaultList := events.HandleNewVaults(chainID, start, end)

	/**********************************************************************************************
	** Once we got all the vaults, we need to remove the duplicates. This is because some vaults
	** were deployed first as experimental vaults and then as standard vaults. In that case, we
	** keep the standard vault.
	**********************************************************************************************/
	uniqueVaultsList := make(map[common.Address]utils.TVaultsFromRegistry)
	for _, v := range standardVaultList {
		uniqueVaultsList[v.VaultsAddress] = v
	}

	for _, v := range experimentalVaultList {
		if _, ok := uniqueVaultsList[v.VaultsAddress]; !ok {
			uniqueVaultsList[v.VaultsAddress] = v
		}
	}

	for _, v := range env.EXTRA_VAULTS[chainID] {
		if _, ok := uniqueVaultsList[v.VaultsAddress]; !ok {
			uniqueVaultsList[v.VaultsAddress] = v
		}
	}

	return vaults.RetrieveActivationForAllVaults(chainID, uniqueVaultsList)
}
