package apr

import (
	"math/big"
	"os"
	"strconv"
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
	return vault.Kind == models.VaultKindMultiple || vault.Kind == models.VaultKindSingle || versionMajor == `3`
}

func computeVaultV3ForwardAPR(
	vault models.TVault,
	allStrategiesForVault map[common.Address]models.TStrategy,
) TForwardAPR {
	netAPR := bigNumber.NewFloat(0)
	oracleContract := env.CHAINS[vault.ChainID].APROracleContract.Address
	if oracleContract == common.HexToAddress(``) {
		logs.Error("No APR oracle contract address for chain " + strconv.FormatUint(vault.ChainID, 10))
		return TForwardAPR{}
	}
	oracle, err := contracts.NewYVaultsV3APROracleCaller(oracleContract, ethereum.GetRPC(vault.ChainID))
	if err != nil {
		logs.Error(err)
		return TForwardAPR{}
	}

	/**********************************************************************************************
	** If the vault is a single strategy vault, we can use the oracle directly to get the APR of
	** the vault as expected APR
	**********************************************************************************************/
	var hasError error
	if vault.Kind == models.VaultKindSingle {
		expected, err := oracle.GetStrategyApr(nil, vault.Address, big.NewInt(0))
		if err == nil {
			netAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
		} else {
			hasError = err
		}
	}

	if vault.Kind == models.VaultKindMultiple || hasError != nil {
		expected, err := oracle.GetCurrentApr(nil, vault.Address)
		if err == nil {
			netAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
		}
	}

	/**********************************************************************************************
	** Otherwise we can do the classic calculation of the net APR by summing the APR of each
	** strategy weighted by the debt ratio of each strategy.
	**********************************************************************************************/
	shouldUseV2Method := false
	if shouldUseV2Method && vault.Kind == models.VaultKindMultiple {
		netAPR = bigNumber.NewFloat(0)
		for _, strategy := range allStrategiesForVault {
			if strategy.LastDebtRatio == nil || strategy.LastDebtRatio.IsZero() {
				if os.Getenv("ENVIRONMENT") == "dev" {
					logs.Info("Skipping strategy " + strategy.Address.Hex() + " for vault " + vault.Address.Hex() + " because debt ratio is zero")
				}
				continue
			}

			expected, err := oracle.GetStrategyApr(nil, strategy.Address, big.NewInt(0))
			if err != nil {
				logs.Error(err)
				continue
			}
			humanizedAPR := helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
			debtRatio := helpers.ToNormalizedAmount(strategy.LastDebtRatio, 4)
			scaledStrategyAPR := bigNumber.NewFloat(0).Mul(humanizedAPR, debtRatio)
			scaledStrategyAPR = bigNumber.NewFloat(0).Mul(scaledStrategyAPR, bigNumber.NewFloat(0.8))
			netAPR = bigNumber.NewFloat(0).Add(netAPR, scaledStrategyAPR)
		}
	}

	return TForwardAPR{
		Type:   `v3:onchainOracle`,
		NetAPR: netAPR,
	}
}
