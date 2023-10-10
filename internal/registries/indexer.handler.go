package registries

import (
	"math/big"

	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

func handleV2Vault(chainID uint64, value *contracts.YRegistryV2NewVault) models.TVaultsFromRegistry {
	newVault := models.TVaultsFromRegistry{
		ChainID:         chainID,
		RegistryAddress: value.Raw.Address,
		Address:         value.Vault,
		TokenAddress:    value.Token,
		APIVersion:      value.ApiVersion,
		BlockNumber:     value.Raw.BlockNumber,
		Type:            models.TokenTypeStandardVault,
	}
	return newVault
}

func handleV3Vault(chainID uint64, value *contracts.YRegistryV3NewVault) models.TVaultsFromRegistry {
	newVault := models.TVaultsFromRegistry{
		ChainID:         chainID,
		RegistryAddress: value.Raw.Address,
		Address:         value.Vault,
		TokenAddress:    value.Token,
		APIVersion:      value.ApiVersion,
		BlockNumber:     value.Raw.BlockNumber,
		Type:            models.TokenTypeStandardVault,
	}
	if value.VaultType.Cmp(big.NewInt(2)) == 0 {
		newVault.Type = models.TokenTypeAutomatedVault
	}
	return newVault
}

func handleV4Vault(chainID uint64, value *contracts.YRegistryV4NewEndorsedVault) models.TVaultsFromRegistry {
	newVault := models.TVaultsFromRegistry{
		ChainID:         chainID,
		RegistryAddress: value.Raw.Address,
		Address:         value.Vault,
		TokenAddress:    value.Asset,
		APIVersion:      value.ReleaseVersion.Text(10),
		BlockNumber:     value.Raw.BlockNumber,
		Type:            models.TokenTypeStandardVault,
	}
	logs.Info(`Got V4 vault ` + value.Vault.Hex() + ` from registry ` + value.Raw.Address.Hex())
	return newVault
}
