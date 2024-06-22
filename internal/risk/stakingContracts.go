package risk

import (
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/multicalls"
	"github.com/yearn/ydaemon/internal/storage"
)

/**********************************************************************************************
** This is the staking contracts strategy. It is responsible for calculating metrics related
** to staking contracts. It only operates on chain ID 10. The metrics calculated include the
** total value locked in each staking pool, the amount of tokens in each pool, and the price
** of each token. This information is used to assess the risk associated with each staking
** pool.
**********************************************************************************************/
func calculateStakingMetrics(chainID uint64) {
	if chainID != 10 {
		return
	}

	/**********************************************************************************************
	** This block of code is responsible for handling the addition of staking pools. It first
	** retrieves all staking pools that have been added. Then, for each pool, it finds the
	** corresponding token.
	** If the token is not found or is not a vault, a warning is logged and the loop continues to
	** the next pool.
	** If the token is found and is a vault, several calls are made to retrieve the total supply,
	** decimals, and recommended price in USDC for the pool. These calls are appended to a list of
	** calls to be performed via a multicall.
	**********************************************************************************************/
	_, allStackingPoolAdded := storage.ListOPStaking(chainID, storage.PerPool)
	calls := []ethereum.Call{}
	for _, pool := range allStackingPoolAdded {
		currentToken, ok := storage.GetERC20(chainID, pool.VaultAddress)
		if !ok {
			logs.Warning(`[retriveStakingContracts] impossible to find token for pool address`, pool.StakingAddress.Hex())
			continue
		}
		if !currentToken.IsVaultLike() {
			logs.Warning(`[retriveStakingContracts] token is not a vault`, pool.VaultAddress.Hex())
			continue
		}
		calls = append(calls, multicalls.GetTotalSupply(pool.StakingAddress.Hex(), pool.StakingAddress))
		calls = append(calls, multicalls.GetDecimals(pool.StakingAddress.Hex(), pool.VaultAddress))
		calls = append(calls, multicalls.GetPriceUsdcRecommendedCall(
			pool.StakingAddress.Hex(),
			env.CHAINS[chainID].LensContract.Address,
			pool.VaultAddress,
		))
	}
	response := multicalls.Perform(chainID, calls, nil)

	/**********************************************************************************************
	** The following block of code is responsible for handling the staking pools. It iterates over
	** each pool and retrieves the total supply, token price, and decimals from the response of the
	** multicall performed earlier.
	** It then formats the token price and total supply to a standard format.
	** If the staking data for the current chain ID does not exist, it initializes it.
	** It then calculates the Total Value Locked (TVL) by multiplying the amount of tokens in the
	** pool by the price of the token. This TVL is then stored in the staking data for the current
	** chain ID and the vault address of the pool.
	**********************************************************************************************/
	for _, pool := range allStackingPoolAdded {
		totalSupply := helpers.DecodeBigInt(response[pool.StakingAddress.Hex()+`totalSupply`])
		tokenPrice := helpers.DecodeBigInt(response[pool.StakingAddress.Hex()+`getPriceUsdcRecommended`])
		decimals := helpers.DecodeUint64(response[pool.StakingAddress.Hex()+`decimals`])

		_, price := helpers.FormatAmount(tokenPrice.String(), 6)
		_, amount := helpers.FormatAmount(totalSupply.String(), int(decimals))

		if _, ok := stakingData[chainID]; !ok {
			stakingData[chainID] = map[string]models.TStaking{}
		}
		tvl, _ := bigNumber.NewFloat(0).Mul(amount, price).Float64()
		stakingData[chainID][pool.VaultAddress.Hex()] = models.TStaking{
			Available:    true,
			Address:      pool.StakingAddress,
			VaultAddress: pool.VaultAddress,
			Risk:         getTVLImpact(bigNumber.NewFloat(0).Mul(amount, price)),
			TVL:          tvl,
			Amount:       amount,
			Price:        price,
		}
	}
}
