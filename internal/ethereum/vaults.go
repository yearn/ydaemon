package ethereum

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/majorfi/ydaemon/internal/contracts"
	"github.com/microgolang/logs"
)

func decodeYearnVaultEvent(client *ethclient.Client, instance *contracts.YearnVault, eventLog types.Log, vaultName string) {
	if log, err := instance.YearnVaultFilterer.ParseStrategyAdded(eventLog); err == nil {
		logs.Success(`Strategy ` + log.Strategy.String() + ` deployed for vault ` + vaultName)
	}
}

func getYearnVaultHistoricalEvents(client *ethclient.Client, instance *contracts.YearnVault, vaultAddress common.Address, block uint64, vaultName string) {
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		logs.Error(err)
		return
	}

	fromBlock := block
	toBlock := currentBlock
	numBlock := uint64(2000) //2000 blocks per request allowd by alchemy
	for i := fromBlock; i < toBlock; i += numBlock {
		time.Sleep(time.Second)
		eventLogs, err := client.FilterLogs(
			context.Background(),
			ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(i)),
				ToBlock:   big.NewInt(int64(i + numBlock)),
				Addresses: []common.Address{common.HexToAddress(vaultAddress.String())},
			},
		)
		if err != nil {
			logs.Error(err)
		} else {
			for _, eventLog := range eventLogs {
				decodeYearnVaultEvent(client, instance, eventLog, vaultName)
			}
		}
	}
}

func listenToYearnVaultEvents(client *ethclient.Client, instance *contracts.YearnVault, startblock uint64, vaultName string) {
	logChan := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(
		context.Background(),
		ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(REGISTRY)},
		},
		logChan,
	)
	if err != nil {
		logs.Error(err)
		return
	}

	logs.Success(`Listening to live events`)
	for {
		select {
		case err := <-sub.Err():
			logs.Error(err)
		case vLog := <-logChan:
			decodeYearnVaultEvent(client, instance, vLog, vaultName)
		}
	}
}
