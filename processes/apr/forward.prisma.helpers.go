package apr

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** Check if the strategy is a prisma strategy. This is a check based on the strategy name. What
** could go wrong.
**************************************************************************************************/
func isPrismaStrategy(strategy models.TStrategy) bool {
	return strings.Contains(strings.ToLower(strategy.Name), `prisma`)
}

/**************************************************************************************************
** Retrieve the prismaReceiver address from the strategy. If the strategy is not a prisma strategy
** or if the value is empty, return an empty address, meaning no prisma rewards.
**************************************************************************************************/
func fetchPrismaReceiver(chainID uint64, strategyAddress common.Address) common.Address {
	strategy, _ := contracts.NewYStrategyPrismaCaller(strategyAddress, ethereum.GetRPC(chainID))
	receiver, err := strategy.PrismaReceiver(nil)
	if err != nil {
		return common.Address{}
	}
	return receiver
}

/**************************************************************************************************
** Calculate the prisma APY for a given vault and strategy. This is done by getting the rewardRate
** from the prismaReceiver contract and dividing it by the total supply of the prismaReceiver.
**************************************************************************************************/
func getPrismaAPY(chainID uint64, prismaReceiver common.Address) (*bigNumber.Float, *bigNumber.Float) {
	receiver, _ := contracts.NewYPrismaReceiverCaller(prismaReceiver, ethereum.GetRPC(chainID))
	rewardRate, err := receiver.RewardRate(nil, big.NewInt(0))
	if err != nil {
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0)
	}
	totalSupply, err := receiver.TotalSupply(nil)
	if err != nil {
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0)
	}
	lpToken, err := receiver.LpToken(nil)
	if err != nil {
		return bigNumber.NewFloat(0), bigNumber.NewFloat(0)
	}
	rate := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(rewardRate), 18)
	supply := helpers.ToNormalizedAmount(bigNumber.NewInt(0).Set(totalSupply), 18)
	prismaPrice := bigNumber.NewFloat(0)
	prismaTokenAddress := common.HexToAddress(`0xdA47862a83dac0c112BA89c6abC2159b95afd71C`)
	if tokenPrice, ok := storage.GetPrice(chainID, prismaTokenAddress); ok {
		prismaPrice = tokenPrice.HumanizedPrice
	}

	lpTokenPrice := bigNumber.NewFloat(0)
	if tokenPrice, ok := storage.GetPrice(chainID, lpToken); ok {
		lpTokenPrice = tokenPrice.HumanizedPrice
	}

	prismaAPR := bigNumber.NewFloat(0).Mul(rate, prismaPrice)
	prismaAPR = bigNumber.NewFloat(0).Mul(prismaAPR, bigNumber.NewFloat(31536000))
	prismaAPR = bigNumber.NewFloat(0).Div(prismaAPR, supply)
	prismaAPR = bigNumber.NewFloat(0).Div(prismaAPR, lpTokenPrice)

	// Convert APR to APY
	const compoundingPeriodsPerYear = 365
	prismaAPY := bigNumber.NewFloat(0).Div(prismaAPR, bigNumber.NewFloat(compoundingPeriodsPerYear))
	prismaAPY = bigNumber.NewFloat(0).Add(prismaAPY, bigNumber.NewFloat(1))
	prismaAPY = bigNumber.NewFloat(0).Pow(prismaAPY, compoundingPeriodsPerYear)
	prismaAPY = bigNumber.NewFloat(0).Sub(prismaAPY, bigNumber.NewFloat(1))

	return prismaAPR, prismaAPY
}
