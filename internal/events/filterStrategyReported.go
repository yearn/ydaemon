package events

import (
	"context"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

type TMimicStrategyReportBase struct {
	Vault     common.Address
	Strategy  common.Address
	Token     common.Address
	Gain      *big.Int
	Loss      *big.Int
	DebtPaid  *big.Int
	TotalGain *big.Int
	TotalLoss *big.Int
	TotalDebt *big.Int
	DebtAdded *big.Int
	DebtRatio *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

/**************************************************************************************************
** filterStrategyReportedFor031To043 will filter all the StrategyReported events for the vaults
** between 0.3.1 and 0.4.3 and store them in an async map to be decoded later. The key is set but
** is to be ignored as the value is the event itself.
**************************************************************************************************/
func filterStrategyReportedFor031To043(vault models.TVault, start uint64, end *uint64, syncMap *sync.Map) {
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
	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[vault.ChainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[vault.ChainID].MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := currentVault.FilterStrategyReported(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				formatedLogs := TMimicStrategyReportBase{
					Vault:     vault.Address,
					Token:     vault.Token.Address,
					Strategy:  log.Event.Strategy,
					Gain:      log.Event.Gain,
					Loss:      log.Event.Loss,
					DebtPaid:  log.Event.DebtPaid,
					TotalGain: log.Event.TotalGain,
					TotalLoss: log.Event.TotalLoss,
					TotalDebt: log.Event.TotalDebt,
					DebtAdded: log.Event.DebtAdded,
					DebtRatio: log.Event.DebtRatio,
					Raw:       log.Event.Raw,
				}
				eventKey := log.Event.Strategy.Hex() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10) + `-` + strconv.FormatUint(uint64(log.Event.Raw.TxIndex), 10) + `-` + strconv.FormatUint(uint64(log.Event.Raw.Index), 10)
				syncMap.Store(eventKey, formatedLogs)
			}
		} else {
			logs.Error(`impossible to FilterStrategyReported for Vault ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** filterStrategyReportedFor030 will filter all the StrategyReported events for the vaults 0.3.0
** and store them in an async map to be decoded later. The key is set but is to be ignored as the
** value is the event itself.
** DebtPaid is set to 0 as it was not present in the event.
**************************************************************************************************/
func filterStrategyReportedFor030(vault models.TVault, start uint64, end *uint64, syncMap *sync.Map) {
	client := ethereum.GetRPC(vault.ChainID)
	currentVault, _ := contracts.NewYvault030(vault.Address, client)
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
	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[vault.ChainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[vault.ChainID].MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := currentVault.FilterStrategyReported(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				formatedLogs := TMimicStrategyReportBase{
					Vault:     vault.Address,
					Token:     vault.Token.Address,
					Strategy:  log.Event.Strategy,
					Gain:      log.Event.Gain,
					Loss:      log.Event.Loss,
					DebtPaid:  big.NewInt(0),
					TotalGain: log.Event.TotalGain,
					TotalLoss: log.Event.TotalLoss,
					TotalDebt: log.Event.TotalDebt,
					DebtAdded: log.Event.DebtAdded,
					DebtRatio: log.Event.DebtRatio,
					Raw:       log.Event.Raw,
				}

				eventKey := log.Event.Strategy.Hex() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10) + `-` + strconv.FormatUint(uint64(log.Event.Raw.TxIndex), 10) + `-` + strconv.FormatUint(uint64(log.Event.Raw.Index), 10)
				syncMap.Store(eventKey, formatedLogs)
			}
		} else {
			logs.Error(`impossible to FilterStrategyReported for Vault ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** filterStrategyReportedFor022 will filter all the StrategyReported events for the vaults 0.2.2
** and store them in an async map to be decoded later. The key is set but is to be ignored as the
** value is the event itself.
** DebtPaid is set to 0 as it was not present in the event.
** DebtRatio is set to 0 as it was not present in the event.
**************************************************************************************************/
func filterStrategyReportedFor022(vault models.TVault, start uint64, end *uint64, syncMap *sync.Map) {
	client := ethereum.GetRPC(vault.ChainID)
	currentVault, _ := contracts.NewYvault022(vault.Address, client)
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
	for chunkStart := start; chunkStart < *end; chunkStart += env.CHAINS[vault.ChainID].MaxBlockRange {
		chunkEnd := chunkStart + env.CHAINS[vault.ChainID].MaxBlockRange
		if chunkEnd > *end {
			chunkEnd = *end
		}

		opts := &bind.FilterOpts{Start: chunkStart, End: &chunkEnd}
		if log, err := currentVault.FilterStrategyReported(opts, nil); err == nil {
			for log.Next() {
				if log.Error() != nil {
					continue
				}
				formatedLogs := TMimicStrategyReportBase{
					Vault:     vault.Address,
					Token:     vault.Token.Address,
					Strategy:  log.Event.Strategy,
					Gain:      log.Event.Gain,
					Loss:      log.Event.Loss,
					DebtPaid:  big.NewInt(0),
					TotalGain: log.Event.TotalGain,
					TotalLoss: log.Event.TotalLoss,
					TotalDebt: log.Event.TotalDebt,
					DebtAdded: log.Event.DebtAdded,
					DebtRatio: big.NewInt(0),
					Raw:       log.Event.Raw,
				}

				eventKey := log.Event.Strategy.Hex() + `-` + strconv.FormatUint(uint64(log.Event.Raw.BlockNumber), 10) + `-` + strconv.FormatUint(uint64(log.Event.Raw.TxIndex), 10) + `-` + strconv.FormatUint(uint64(log.Event.Raw.Index), 10)
				syncMap.Store(eventKey, formatedLogs)
			}
		} else {
			logs.Error(`impossible to FilterStrategyReported for Vault ` + vault.Address.Hex() + ` on chain ` + strconv.FormatUint(vault.ChainID, 10) + `: ` + err.Error())
		}
	}
}

/**************************************************************************************************
** HandleStrategyReported will loop over a map of vaults and fetch all the strategy reported events
** from start to end. Based on the version of the vault, the function will call the correct
** function to fetch the events.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultsMap: the map of vaults we want to fetch the fee for
** - start: the block number to start fetching from
** - end: the block number to stop fetching from
**
** Returns:
** - An array of TMimicStrategyReportBase via the T any. This hack is used to avoid import circles.
**************************************************************************************************/
func HandleStrategyReported(
	chainID uint64,
	vaultsMap map[common.Address]models.TVault,
	start uint64,
	end *uint64,
) []TMimicStrategyReportBase {
	timeBefore := time.Now()
	syncMap := sync.Map{}

	wg := &sync.WaitGroup{}
	for _, vault := range vaultsMap {
		wg.Add(1)
		go func(vault models.TVault) {
			defer wg.Done()
			switch vault.Version {
			case `0.2.2`:
				filterStrategyReportedFor022(
					vault,
					start,
					end,
					&syncMap,
				)
			case `0.3.0`:
				filterStrategyReportedFor030(
					vault,
					start,
					end,
					&syncMap,
				)
			default: //case `0.3.1`, `0.3.2`, `0.3.3`, `0.3.4`, `0.3.5`, `0.4.2`, `0.4.3`:
				filterStrategyReportedFor031To043(
					vault,
					start,
					end,
					&syncMap,
				)
			}
		}(vault)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once we have all the logs, we will dump them from the sync.Map to the array of
	** TMimicStrategyReportBase to use them in the rest of the process.
	**********************************************************************************************/
	events := []TMimicStrategyReportBase{}
	syncMap.Range(func(_, value interface{}) bool {
		events = append(events, value.(TMimicStrategyReportBase))
		return true
	})

	logs.Success(`It tooks`, time.Since(timeBefore), `to retrieve the all the harvest events`)
	return events
}
