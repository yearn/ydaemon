package partners

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils/bigNumber"
	"github.com/yearn/ydaemon/internal/utils/contracts"
	"github.com/yearn/ydaemon/internal/utils/ethereum"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

func (c *TPartners) BalanceOf() *TPartners {
	for index, wrapper := range c.Wrappers {
		switch wrapper.Type {
		case `wildcard`:
			logs.Info(`wildcard`, wrapper.Name, wrapper.Wrapper)
		case `gearbox`:
			logs.Info(`gearbox`, wrapper.Name, wrapper.Wrapper)
		case `element`:
			logs.Info(`element`, wrapper.Name, wrapper.Wrapper)
		case `bentobox`:
			logs.Info(`bentobox`, wrapper.Name, wrapper.Wrapper)
		case `begenbox`:
			logs.Info(`begenbox`, wrapper.Name, wrapper.Wrapper)
		case `yapeswapFactory`:
			logs.Info(`yapeswapFactory`, wrapper.Name, wrapper.Wrapper)
		default:
			bBalance, _, err := computeDefaultBalance(1, wrapper.Vault, wrapper.Wrapper)
			if err != nil {
				c.Wrappers[index].BalanceOf = bigNumber.NewInt()
				continue
			}
			// decimals := store.Tokens[1][wrapper.Vault].Decimals
			// _, balance := helpers.FormatAmount(bBalance.String(), int(decimals))
			// logs.Info(wrapper.Name, ` : `, balance.String())
			c.Wrappers[index].BalanceOf = bigNumber.FromBigInt(bBalance)
		}
	}

	return c
}

func computeDefaultBalance(
	chainID uint64,
	vaultAddress common.Address,
	userAddress common.Address,
) (*big.Int, *bigNumber.BigFloat, error) {
	yearnVault, err := contracts.NewYearnVault(vaultAddress, ethereum.GetRPC(chainID))
	if err != nil {
		logs.Error("Failed to create YearnVault contract", err)
		return nil, nil, err
	}
	balance, err := yearnVault.BalanceOf(nil, userAddress)
	if err != nil {
		logs.Error(`Error getting balance`, err)
		return nil, nil, err
	}

	decimals := tokens.Store.Tokens[chainID][vaultAddress].Decimals
	_, fBalance := helpers.FormatAmount(balance.String(), int(decimals))

	return balance, fBalance, err
}
