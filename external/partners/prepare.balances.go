package partners

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/tokens"
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
			c.Wrappers[index].BalanceOf = bBalance
		}
	}

	return c
}

func computeDefaultBalance(
	chainID uint64,
	vaultAddress common.MixedcaseAddress,
	userAddress common.MixedcaseAddress,
) (*bigNumber.Int, *bigNumber.Float, error) {
	yearnVault, _ := contracts.NewYearnVault(vaultAddress.Address(), ethereum.GetRPC(chainID))
	balance, err := yearnVault.BalanceOf(nil, userAddress.Address())
	if err != nil {
		logs.Error(err, `Error getting balance`, err)
		return nil, nil, err
	}

	decimals := 18
	if token, ok := tokens.FindToken(chainID, vaultAddress.Address()); ok {
		decimals = int(token.Decimals)
	}

	_, fBalance := helpers.FormatAmount(balance.String(), int(decimals))

	return bigNumber.SetInt(balance), fBalance, err
}
