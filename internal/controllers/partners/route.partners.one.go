package partnersController

import (
	"math"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/contracts"
	"github.com/yearn/ydaemon/internal/ethereum"
	"github.com/yearn/ydaemon/internal/logs"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/store"
)

func formatAmount(balanceToken string, decimals int) float64 {
	fTotalAssets := new(big.Float)
	fTotalAssets.SetString(balanceToken)
	humanizedBalance := new(big.Float).Quo(fTotalAssets, big.NewFloat(math.Pow10(decimals)))
	fhumanizedBalance, _ := humanizedBalance.Float64()
	return fhumanizedBalance
}

func BalanceOf(c *models.TPartners) *big.Int {
	balance := big.NewInt(0)

	for _, wrapper := range c.Wrappers {
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
			bal, err := BalanceOf2(1, wrapper.Vault, wrapper.Wrapper)
			if err != nil {
				logs.Error(`Error creating YearnVault contract`, err)
			}
			store.Tokens[0] = make(map[common.Address]models.TERC20Token)
			decimals := store.Tokens[1][wrapper.Vault].Decimals
			logs.Info(`Balance of`, wrapper.Name, ` : `, formatAmount(bal.String(), int(decimals)))
		}
	}

	return balance
}

func BalanceOf2(chainID uint64, vaultAddress common.Address, userAddress common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(ethereum.GetRPCURI(chainID))
	if err != nil {
		logs.Error("Failed to connect to Ethereum node")
	}

	yearnVault, err := contracts.NewYearnVault(
		vaultAddress,
		client,
	)
	if err != nil {
		logs.Error(`Error creating YearnVault contract`, err)
	}
	bal, err := yearnVault.BalanceOf(nil, userAddress)
	if err != nil {
		logs.Error(`Error creating YearnVault contract`, err)
	}

	return bal, err
}

// GetPartner will, for a given address on given chainID, return the meta informations for the partner.
func (y Controller) GetPartner(c *gin.Context) {
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	partnerAddressOrName := c.Param("addressOrName")
	if partnerAddressOrName == `` {
		c.String(http.StatusBadRequest, "invalid address or name")
		return
	}
	partner, ok := store.PartnersByName[chainID][partnerAddressOrName]
	if !ok {
		partner, ok = store.PartnersByAddress[chainID][common.HexToAddress(partnerAddressOrName)]
		if !ok {
			c.String(http.StatusBadRequest, "no data available")
			return
		}
	}

	BalanceOf(&partner)

	c.JSON(http.StatusOK, partner)
}
