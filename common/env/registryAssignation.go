package env

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
)

/**************************************************************************************************
** IsRegistryFromYearnCore will check if the registry is from the Yearn Core type of vault. This
** means the vaults that should be displayed in the Yearn.fi website.
** Some vaults/registries are Yearn, but for side projects.
**************************************************************************************************/
func IsRegistryFromYearnCore(chainID uint64, registryAddress common.Address) bool {
	chains, ok := GetChain(chainID)
	if !ok {
		return false
	}
	registries := chains.Registries
	for _, registry := range registries {
		if registry.Address.Hex() == registryAddress.Hex() {
			return registry.Label == "YEARN"
		}
	}
	return false
}

/**************************************************************************************************
** IsRegistryFromJuiced will check if the registry is for Juiced vaults. They will be displayed on
** the Juiced.app website.
**************************************************************************************************/
func IsRegistryFromJuiced(chainID uint64, registryAddress common.Address) bool {
	chains, ok := GetChain(chainID)
	if !ok {
		return false
	}
	registries := chains.Registries
	for _, registry := range registries {
		if registry.Address.Hex() == registryAddress.Hex() {
			return registry.Label == "JUICED"
		}
	}
	return false
}

/**************************************************************************************************
** IsRegistryFromPublicERC4626 will check if the registry is for public ERC4626 vaults. They will
** not be displayed by default in the Yearn.fi website.
**************************************************************************************************/
func IsRegistryFromPublicERC4626(chainID uint64, registryAddress common.Address) bool {
	chains, ok := GetChain(chainID)
	if !ok {
		return false
	}
	registries := chains.Registries
	for _, registry := range registries {
		if registry.Address.Hex() == registryAddress.Hex() {
			return registry.Label == "PUBLIC_ERC4626"
		}
	}
	return false
}

/**************************************************************************************************
** IsRegistryFromPoolTogether will check if the registry is for Yearn X Pooltogether vaults.
** They will not be displayed by default in the Yearn.fi website.
**************************************************************************************************/
func IsRegistryFromPoolTogether(chainID uint64, registryAddress common.Address) bool {
	return chainID == 10 && addresses.Equals(
		registryAddress,
		common.HexToAddress(`0x0c379e9b71ba7079084ada0d1c1aeb85d24dfd39`),
	)
}

func IsRegistryDisabled(chainID uint64, registryAddress common.Address) bool {
	chains, ok := GetChain(chainID)
	if !ok {
		return false
	}
	registries := chains.Registries
	for _, registry := range registries {
		if registry.Address.Hex() == registryAddress.Hex() {
			return registry.Tag == "DISABLED"
		}
	}
	return false
}
