package apr

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

func isV3Vault(vault models.TVault) bool {
	versionMajor := strings.Split(vault.Version, `.`)[0]
	return vault.Kind == models.VaultKindMultiple || vault.Kind == models.VaultKindSingle || versionMajor == `3` || versionMajor == `~3`
}

func computeVaultV3ForwardAPY(
	vault models.TVault,
	allStrategiesForVault map[string]models.TStrategy,
) TForwardAPY {
	oracleAPR := bigNumber.NewFloat(0)
	chain, ok := env.GetChain(vault.ChainID)
	if !ok {
		return TForwardAPY{}
	}
	oracleContract := chain.APROracleContract.Address
	if oracleContract == common.HexToAddress(``) {
		return TForwardAPY{}
	}
	oracle, err := contracts.NewYVaultsV3APROracleCaller(oracleContract, ethereum.GetRPC(vault.ChainID))
	if err != nil {
		logs.Error(err)
		return TForwardAPY{}
	}

	/**********************************************************************************************
	** Use the oracle to get the APR of the vault. The oracle automatically handles:
	** - Single strategy vaults: Returns strategy APR
	** - Multi-strategy vaults: Returns weighted average with performance fees applied
	**********************************************************************************************/
	expected, err := oracle.GetStrategyApr(nil, vault.Address, big.NewInt(0))
	if err != nil {
		logs.Error(`GetStrategyApr failed for vault ` + vault.Address.Hex() + `: ` + err.Error())
		return TForwardAPY{}
	}
	oracleAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)

	/**********************************************************************************************
	** Use the oracle APR as the primary APR (no manual calculation needed)
	**********************************************************************************************/
	primaryAPR := oracleAPR

	primaryAPRFloat64, _ := primaryAPR.Float64()
	primaryAPY := bigNumber.NewFloat(0).SetFloat64(convertFloatAPRToAPY(primaryAPRFloat64, 52))

	return TForwardAPY{
		Type:   `v3:onchainOracle`,
		NetAPY: primaryAPY,
		Composite: TCompositeData{
			V3OracleCurrentAPR:    primaryAPY,
			V3OracleStratRatioAPR: bigNumber.NewFloat(0),
		},
	}
}
