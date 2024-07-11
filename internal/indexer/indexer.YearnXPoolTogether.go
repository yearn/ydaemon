package indexer

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/addresses"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/** 🔵 - Yearn *************************************************************************************
**
**************************************************************************************************/
func IndexYearnXPoolTogetherVaults(chainID uint64) (vaults []models.TVaultsFromRegistry) {
	if chainID != 10 {
		return vaults
	}

	/** 🔵 - Yearn *********************************************************************************
	** Loop over all the known vaults for the specified chain ID.
	**********************************************************************************************/
	client := ethereum.GetRPC(chainID)
	registry, err := contracts.NewPoolTogetherRegistry(
		common.HexToAddress(`0x0c379e9b71ba7079084ada0d1c1aeb85d24dfd39`),
		client,
	)
	if err != nil {
		return vaults
	}
	totalVaults, err := registry.TotalVaults(nil)
	if err != nil {
		return vaults
	}

	for i := uint64(0); i < totalVaults.Uint64(); i++ {
		vaultAddress, err := registry.AllVaults(nil, big.NewInt(int64(i)))
		if err != nil {
			continue
		}

		if addresses.Equals(vaultAddress, `0x3a14DdB934e785Cd1e29949EA814e8090D5F8b69`) {
			continue
		}
		if addresses.Equals(vaultAddress, `0x37E49c9dBf195F5D436C7A7610fe703cDcd8147B`) {
			continue
		}

		if _, ok := storage.GetVaultFromRegistry(chainID, vaultAddress); ok {
			continue
		}
		vaultContract, err := contracts.NewERC4626(vaultAddress, client)
		if err != nil {
			continue
		}
		vaultAsset, _ := vaultContract.Asset(nil)
		vaultName, _ := vaultContract.Name(nil)
		if !strings.Contains(strings.ToLower(vaultName), `yearn`) {
			continue
		}

		poolTogetherVault := models.TVaultsFromRegistry{
			ChainID:         chainID,
			Address:         vaultAddress,
			RegistryAddress: common.HexToAddress(`0x0c379e9b71ba7079084ada0d1c1aeb85d24dfd39`),
			TokenAddress:    vaultAsset,
			APIVersion:      `3.0.2`,
			BlockNumber:     119_536_717,
			Type:            models.TTokenType(models.VaultKindSingle),
			Kind:            models.VaultKindSingle,
		}
		vaults = append(vaults, poolTogetherVault)
		storage.StoreNewVaultToRegistry(chainID, poolTogetherVault)
	}

	return vaults
}
