package vaults

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/utils"
)

/**************************************************************************************************
** Filter all transfer events to retrieve the transfer value and store it in an array of
** TEventBlock
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: the address of the vault we are working on
** - activation: the block number at which the strategy was activated
** - asyncMapLastReports: the ptr to the async map to store the blockTimestamp
** - wg: the async ptr to the WaitGroup to sync the goroutines
**
** Returns nothing as asyncMapLastReports is updated via a pointer
**************************************************************************************************/
func filterStrategyReported(
	chainID uint64,
	vaultAddress common.Address,
	activation uint64,
	asyncMapLastReports *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := currentVault.FilterStrategyReported(&bind.FilterOpts{Start: activation}, nil); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			blockTimestamp := ethereum.GetBlockTime(chainID, log.Event.Raw.BlockNumber)
			eventKey := (vaultAddress.String() + `-` +
				log.Event.Strategy.String() + `-` +
				strconv.FormatUint(log.Event.Raw.BlockNumber, 10))
			asyncMapLastReports.LoadOrStore(eventKey, blockTimestamp)
		}
	} else {
		traces.
			Capture(`error`, `impossible to FilterStrategyReported for Yvault043 `+vaultAddress.Hex()).
			SetEntity(`vault`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
			SetTag(`vaultAddress`, vaultAddress.Hex()).
			Send()
	}
}

/**********************************************************************************************
** Retrieve all transfers from vaults to strategies. This can only happen in one situation: the
** vault is sending strategist fees to the strategy for them to be taken by the strategist.
** We need that to be able to calculate the strategist fees as many variable could make the
** offchain calculation wrong.
** Thanks to this number, from offchain totalFees calculation, we can deduct the treasury fees.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - strategies: list of all TStrategyAdded to work on
**
** Returns:
** - a map of vaultAddress -> strategyAddress -> blockNumber -> TEventBlock
**********************************************************************************************/
func RetrieveHarvests(
	chainID uint64,
	vaults map[common.Address]utils.TVaultsFromRegistry,
) map[common.Address]map[common.Address]map[uint64]uint64 {
	trace := traces.Init(`app.indexer.vaults.harvest_events`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `vaults`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** Concurrently retrieve all strategyReported from vaults to strategies, waiting for the end
	** of all goroutines via the wg before continuing.
	**********************************************************************************************/
	asyncMapLastReports := sync.Map{}
	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go filterStrategyReported(
			chainID,
			v.VaultsAddress,
			v.Activation,
			&asyncMapLastReports,
			wg,
		)
	}
	wg.Wait()

	/**********************************************************************************************
	** Once all transfers from vaults to strategies have been retrieved, we need to extract them
	** from the sync.Map.
	**
	** The syncMap variable is setup as follows:
	** - key: vaultAddress-strategyAddress-blockNumber
	** - value: []TEventBlock
	**********************************************************************************************/
	count := 0
	lastReportForStrategy := make(map[common.Address]map[common.Address]map[uint64]uint64)
	asyncMapLastReports.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddressParsed := common.HexToAddress(eventKey[0])
		strategyAddressParsed := common.HexToAddress(eventKey[1])
		blockNumber, _ := strconv.ParseUint(eventKey[2], 10, 64)

		if _, ok := lastReportForStrategy[vaultAddressParsed]; !ok {
			lastReportForStrategy[vaultAddressParsed] = make(map[common.Address]map[uint64]uint64)
		}
		if _, ok := lastReportForStrategy[vaultAddressParsed][strategyAddressParsed]; !ok {
			lastReportForStrategy[vaultAddressParsed][strategyAddressParsed] = make(map[uint64]uint64)
		}
		lastReportForStrategy[vaultAddressParsed][strategyAddressParsed][blockNumber] = value.(uint64)
		count++
		return true
	})

	return lastReportForStrategy
}
