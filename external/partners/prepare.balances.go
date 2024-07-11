package partners

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
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
			bBalance, err := computeDefaultBalance(1, wrapper.Vault, wrapper.Wrapper)
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
	vaultAddress string,
	userAddress string,
) (*bigNumber.Int, error) {
	yearnVault, _ := contracts.NewYearnVault(common.HexToAddress(vaultAddress), ethereum.GetRPC(chainID))
	balance, err := yearnVault.BalanceOf(nil, common.HexToAddress(userAddress))
	if err != nil {
		logs.Error(err, `Error getting balance`, err)
		return nil, err
	}
	return bigNumber.SetInt(balance), err
}
