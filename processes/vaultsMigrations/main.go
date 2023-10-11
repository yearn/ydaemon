package vaultsMigrations

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/indexer"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/vaults"
	"github.com/yearn/ydaemon/processes/prices"
)

func Run(chainID uint64) {
	/**********************************************************************************************
	** The first steps are to init out environment. We need to fetch all our vaults, strategies,
	** tokens, prices, etc. This will allow us to have an exhaustive list of all the data we need
	** to process.
	**********************************************************************************************/
	initYearnEcosystem(chainID)
	allVaults := vaults.ListVaults(chainID)

	/**********************************************************************************************
	** Once all the vaults are here, we can start checking if the vault is disabled or not and if
	** he has a migration or not.
	** First, we will simply query the blockchain to see if the vault is disabled or not.
	**********************************************************************************************/
	calls := []ethereum.Call{}
	for _, vault := range allVaults {
		calls = append(calls, multicalls.GetDecimals(vault.Address.Hex(), vault.Address))
		calls = append(calls, multicalls.GetDepositLimit(vault.Address.Hex(), vault.Address))
	}
	response := multicalls.Perform(chainID, calls, nil)
	for _, vault := range allVaults {
		vaultChain := []common.Address{}
		currentVault := vault.Address
		for {
			vaultChain = append(vaultChain, currentVault)
			toVault, depositLimit, isOk, hasWarning, hasError := checkVaultMigrationStatus(currentVault, response, chainID)
			if hasWarning || hasError {
				str := strconv.FormatUint(chainID, 10) + ` - `
				for _, v := range vaultChain {
					str += v.Hex() + ` -> `
				}
				if hasError {
					str += `🔴 ERROR`
				} else if hasWarning {
					str += `🟠 WARNING`
					break
				}
				logs.Info(str)
				break
			} else if (isOk && toVault != common.Address{}) {
				currentVault = toVault
				continue
			} else {
				if len(vaultChain) == 1 {
					break
				}
				str := strconv.FormatUint(chainID, 10) + ` - `
				for _, v := range vaultChain {
					str += v.Hex() + ` -> `
				}
				str += `🟢 OK (depositLimit: ` + depositLimit.String() + `)`
				// logs.Info(str)
				break
			}
		}
	}
	logs.Success(`Process finished for chain ` + strconv.FormatUint(chainID, 10))
}

func checkVaultMigrationStatus(vaultAddress common.Address, depositLimitMulticallResponse map[string][]any, chainID uint64) (common.Address, *bigNumber.Float, bool, bool, bool) {
	rawDepositLimit := depositLimitMulticallResponse[vaultAddress.Hex()+`depositLimit`]
	rawDecimals := depositLimitMulticallResponse[vaultAddress.Hex()+`decimals`]
	depositLimit := helpers.DecodeBigInt(rawDepositLimit)
	decimals := helpers.DecodeUint64(rawDecimals)
	depositLimitNormalized := helpers.ToNormalizedAmount(depositLimit, decimals)
	if vault, ok := storage.GetVault(chainID, vaultAddress); ok {
		if depositLimit.Gt(bigNumber.Zero) {
			return common.Address{}, depositLimitNormalized, true, false, false
		}

		if depositLimit.Eq(bigNumber.Zero) { //Deposit limit 0
			if vault.Migration.Available {
				return vault.Migration.Target, depositLimitNormalized, true, false, false
			}
			if (vault.IsRetired) && !vault.Migration.Available {
				return common.Address{}, depositLimitNormalized, false, true, false
			}
			if !vault.Migration.Available {
				return common.Address{}, depositLimitNormalized, false, false, true
			}
		}
		// never reached
	}
	// logs.Error(strconv.FormatUint(chainID, 10) + `- Vault ` + vaultAddress.Hex() + ` not found in the meta files`)
	return common.Address{}, depositLimitNormalized, false, false, true
}

/**************************************************************************************************
** initYearnEcosystem is used to initialize the yearn ecosystem data, aka fetching all the vaults,
** strategies, tokens, prices, etc.
** Based on that, we have everything ready to compute the fees for each partner.
**************************************************************************************************/
func initYearnEcosystem(chainID uint64) {
	historicalVaults := indexer.IndexNewVaults(chainID)
	tokens.RetrieveAllTokens(chainID, historicalVaults)
	prices.Run(chainID)
	vaults.RetrieveAllVaults(chainID, historicalVaults)
}
