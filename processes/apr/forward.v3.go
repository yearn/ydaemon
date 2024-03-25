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
	allStrategiesForVault map[string]models.TStrategy,
) TForwardAPR {
	oracleAPR := bigNumber.NewFloat(0)
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
			oracleAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
		} else {
			hasError = err
		}
	}

	if vault.Kind == models.VaultKindMultiple || hasError != nil {
		expected, err := oracle.GetCurrentApr(nil, vault.Address)
		if err == nil {
			oracleAPR = helpers.ToNormalizedAmount(bigNumber.SetInt(expected), 18)
		}
	}

	/**********************************************************************************************
	** Otherwise we can do the classic calculation of the net APR by summing the APR of each
	** strategy weighted by the debt ratio of each strategy.
	**********************************************************************************************/
	debtRatioAPR := bigNumber.NewFloat(0)
	if vault.Kind == models.VaultKindMultiple {
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

			// Scaling based on the performance fee
			// Retrieve the ratio we should use to take into account the performance fee. If the performance fee is 10%, the ratio is 0.9
			// 10_000 is the precision. Ex: 1 - (1000 / 10_000)
			performanceFeeFloat := bigNumber.NewFloat(0).SetInt(strategy.LastPerformanceFee)
			performanceFee := bigNumber.NewFloat(0).Div(performanceFeeFloat, bigNumber.NewFloat(10_000))
			performanceFee = bigNumber.NewFloat(0).Sub(bigNumber.NewFloat(1), performanceFee)
			scaledStrategyAPR = bigNumber.NewFloat(0).Mul(scaledStrategyAPR, performanceFee)

			debtRatioAPR = bigNumber.NewFloat(0).Add(debtRatioAPR, scaledStrategyAPR)
		}

		//Reduce the APR by 10% to account for the fees/slippage and other factors
		debtRatioAPR = bigNumber.NewFloat(0).Mul(debtRatioAPR, bigNumber.NewFloat(0.9))
	}

	/**********************************************************************************************
	** Define which APR we want to use as "Net APR".
	**********************************************************************************************/
	primaryAPR := oracleAPR
	if vault.Metadata.ShouldUseV2APR {
		primaryAPR = debtRatioAPR
	}
	return TForwardAPR{
		Type:   `v3:onchainOracle`,
		NetAPR: primaryAPR,
		Composite: TCompositeData{
			V3OracleCurrentAPR:    oracleAPR,
			V3OracleStratRatioAPR: debtRatioAPR,
		},
	}
}
