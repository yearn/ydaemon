package ethereum

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/majorfi/ydaemon/internal/contracts"
	"github.com/majorfi/ydaemon/internal/utils"
	"github.com/microgolang/logs"
)

var REGISTRY = "0x50c1a2eA0a861A967D9d0FFE2AE4012c2E053804"
var REGISTRY_DEPLOYED_AT = "12045555"

func decodeEvent(
	client *ethclient.Client,
	instance *contracts.Registry,
	eventLog types.Log,
) {
	blockDetails, err := client.HeaderByHash(context.Background(), eventLog.BlockHash)
	if err != nil {
		logs.Error(err)
		return
	}
	if log, err := instance.RegistryFilterer.ParseNewVault(eventLog); err == nil {
		erc20TokenInstance, err := contracts.NewERC20(log.Token, client)
		if err != nil {
			logs.Error(err)
			return
		}
		vaultInstance, err := contracts.NewYearnVault(log.Vault, client)
		if err != nil {
			logs.Error(err)
			return
		}

		tokenName, _ := erc20TokenInstance.Name(nil)
		vaultName, _ := vaultInstance.Name(nil)
		vaultVersion, _ := vaultInstance.ApiVersion(nil)
		logs.Success(`New v` + vaultVersion + ` vault ` + vaultName + ` deployed at ` + log.Vault.String() + ` for token ` + tokenName + ` at block ` + blockDetails.Number.String())

		go getYearnVaultHistoricalEvents(client, vaultInstance, log.Vault, eventLog.BlockNumber, vaultName)

	}
}

func getHistoricalEvents(
	client *ethclient.Client,
	instance *contracts.Registry,
) {
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		logs.Error(err)
		return
	}

	baseBlock, _ := strconv.ParseUint(REGISTRY_DEPLOYED_AT, 10, 64)
	fromBlock := baseBlock
	toBlock := currentBlock
	numBlock := uint64(2000) //2000 blocks per request allowd by alchemy
	for i := fromBlock; i < toBlock; i += numBlock {
		time.Sleep(time.Second)
		eventLogs, err := client.FilterLogs(
			context.Background(),
			ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(i)),
				ToBlock:   big.NewInt(int64(i + numBlock)),
				Addresses: []common.Address{common.HexToAddress(REGISTRY)},
			},
		)
		if err != nil {
			logs.Error(err)
		} else {
			for _, eventLog := range eventLogs {
				decodeEvent(client, instance, eventLog)
			}
		}
	}
}

func listenToEvents(
	client *ethclient.Client,
	instance *contracts.Registry,
) {
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
			decodeEvent(client, instance, vLog)
		}
	}
}

func InitDaemon(shouldFetchHistoricalEvents bool) {
	client, err := ethclient.Dial(utils.EthNodeURI)
	if err != nil {
		logs.Error(err)
		return
	}

	instance, err := contracts.NewRegistry(common.HexToAddress(REGISTRY), client)
	if err != nil {
		logs.Error(err)
		return
	}

	if shouldFetchHistoricalEvents {
		go getHistoricalEvents(client, instance)
	}
	listenToEvents(client, instance)
}
