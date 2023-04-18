package events

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** filterUpdateRewards filters the UpdateRewards events emitted by the specified Yvault043
** contract starting from the specified activation block number, and stores the latest event's
** Rewards value in the reportMap sync.Map under a key derived from the contract address
** and the block number.
**
** Arguments:
** -chainID the ID of the Ethereum network to connect to
** - vaultAddress the address of the Yvault043 contract to filter events from
** - activation the block number to start filtering events from
** - reportMap the sync.Map to store the latest event's Rewards value in
** Returns nothing as asyncMapTransfers is updated via a pointer
*************************************************************************************************/
func filterUpdateRewards(vault *models.TVault, start uint64, end *uint64, reportMap *sync.Map) {
	client := ethereum.GetRPC(vault.ChainID)
	currentVault, _ := contracts.NewYvault043(vault.Address, client)

	/**********************************************************************************************
	** First, we need to know when to stop our log fetching. By default, we will fetch until the
	** current block number, aka nil.
	** As using nil may cause some issues, we will specify the current block number instead.
	**********************************************************************************************/
	if end == nil {
		blockEnd, _ := client.BlockNumber(context.Background())
		end = &blockEnd
	}

	/******************************************************************************************
	** Then, we need to know when to start our log fetching. By default, we will fetch from the
	** activation block in order to get all the vaults that were ever deployed since it was
	** deployed.
	******************************************************************************************/
	if start == 0 {
		start = vault.Activation
	}

	/******************************************************************************************
	** Finally, we will fetch the logs in chunks of MAX_BLOCK_RANGE blocks. This is done to
	** avoid hitting some external node providers' rate limits.
	******************************************************************************************/
	for chunkStart := start; chunkStart < *end; chunkStart += env.MAX_BLOCK_RANGE[vault.ChainID] {
		chunkEnd := chunkStart + env.MAX_BLOCK_RANGE[vault.ChainID]
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := currentVault.FilterUpdateRewards(opts); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				eventKey := vault.Address.Hex() + `-` + strconv.FormatUint(log.Event.Raw.BlockNumber, 10)
				reportMap.LoadOrStore(eventKey, log.Event.Rewards.Hex())
			}
		} else {
			logs.Error(`impossible to FilterUpdateRewards for YRegistryV2 ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
		}
	}
}

/**********************************************************************************************
** HandleUpdateRewards retrieves all treasury change events for the provided vaults, waiting
** for the end of all goroutines before continuing.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: a map of all TVault to work on
**
** Returns:
** - a map of vaultAddress -> blockNumber -> treasuryAddress
**********************************************************************************************/
func HandleUpdateRewards(
	chainID uint64,
	vaults map[common.Address]*models.TVault,
	start uint64,
	end *uint64,
) map[common.Address]map[uint64]common.Address {
	/**********************************************************************************************
	** Concurrently retrieve all treasuryChange events for the vaults, waiting for the end
	** of all goroutines via the wg before continuing.
	**********************************************************************************************/
	reportMap := sync.Map{}
	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go func(v *models.TVault) {
			defer wg.Done()
			filterUpdateRewards(v, start, end, &reportMap)
		}(v)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once all the treasury changes have been catched, we need to extract them from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-blockNumber
	** - value: treasuryAddress as common.Address
	**********************************************************************************************/
	count := 0
	treasuryChangeForVault := make(map[common.Address]map[uint64]common.Address)
	reportMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddressParsed := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[1], 10, 64)

		if _, ok := treasuryChangeForVault[vaultAddressParsed]; !ok {
			treasuryChangeForVault[vaultAddressParsed] = make(map[uint64]common.Address)
		}
		treasuryChangeForVault[vaultAddressParsed][blockNumber] = common.HexToAddress(value.(string))
		count++
		return true
	})

	return treasuryChangeForVault
}
