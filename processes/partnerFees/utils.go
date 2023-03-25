package partnerFees

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gocarina/gocsv"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/prices"
)

func formatWithPriceOnBlock(chainID uint64, blockNumber uint64, tokenAddress common.Address, amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	price, ok := prices.GetPricesOnBlock(chainID, blockNumber, tokenAddress)
	if !ok {
		logs.Warning(`No price found for`, tokenAddress.Hex(), `on block`, blockNumber)
	}
	return toNormalizedValue(amount, price, decimals)
}

func toNormalizedAmount(amount *bigNumber.Int, decimals uint64) *bigNumber.Float {
	return bigNumber.NewFloat(0).Quo(
		bigNumber.NewFloat(0).SetInt(amount),
		bigNumber.NewFloat(0).SetInt(
			bigNumber.NewInt(0).Exp(bigNumber.NewInt(10), bigNumber.NewInt(int64(decimals)), nil),
		),
	)
}

func toNormalizedRatio(amountA *bigNumber.Int, amountB *bigNumber.Int, decimals uint64) *bigNumber.Float {
	normalizedA := toNormalizedAmount(amountA, decimals)
	normalizedB := toNormalizedAmount(amountB, decimals)
	normalizedRatio := bigNumber.NewFloat(0).Quo(normalizedA, normalizedB)
	percent := bigNumber.NewFloat(0).Mul(normalizedRatio, bigNumber.NewFloat(100))
	return percent
}

func toNormalizedValue(amount *bigNumber.Int, price *bigNumber.Int, decimals uint64) *bigNumber.Float {
	normalizedAmount := toNormalizedAmount(amount, decimals)
	normalizedPrice := toNormalizedAmount(price, 6)
	normalizedValue := bigNumber.NewFloat(0).Mul(normalizedAmount, normalizedPrice)
	return normalizedValue
}

func partnerTVLTier(value *bigNumber.Float) *bigNumber.Float {
	if value.Lt(bigNumber.NewFloat(1_000_000)) {
		return bigNumber.NewFloat(0)
	}
	if value.Lt(bigNumber.NewFloat(5_000_000)) {
		return bigNumber.NewFloat(0.1)
	}
	if value.Lt(bigNumber.NewFloat(10_000_000)) {
		return bigNumber.NewFloat(0.15)
	}
	if value.Lt(bigNumber.NewFloat(50_000_000)) {
		return bigNumber.NewFloat(0.2)
	}
	if value.Lt(bigNumber.NewFloat(100_000_000)) {
		return bigNumber.NewFloat(0.25)
	}
	if value.Lt(bigNumber.NewFloat(200_000_000)) {
		return bigNumber.NewFloat(0.3)
	}
	if value.Lt(bigNumber.NewFloat(400_000_000)) {
		return bigNumber.NewFloat(0.35)
	}
	if value.Lt(bigNumber.NewFloat(700_000_000)) {
		return bigNumber.NewFloat(0.4)
	}
	if value.Lt(bigNumber.NewFloat(1_000_000_000)) {
		return bigNumber.NewFloat(0.45)
	}
	return bigNumber.NewFloat(0.5)
}

func exportCSV[T any](fileName string, data T) {
	os.Remove(fileName)
	outputFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logs.Error(err)
	}
	defer outputFile.Close()
	if err := gocsv.MarshalFile(&data, outputFile); err != nil {
		logs.Error(err)
	}
}

func getChainStartBlock(chainID uint64) uint64 {
	switch chainID {
	case 1:
		return 14166636
	case 10:
		return 29675215
	case 250:
		return 40499061
	case 42161:
		return 30385403
	default:
		return 0
	}
}
