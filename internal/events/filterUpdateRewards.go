package events

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/traces"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** filterUpdateRewards filters the UpdateRewards events emitted by the specified Yvault043
** contract starting from the specified activation block number, and stores the latest event's
** Rewards value in the asyncMapLastReports sync.Map under a key derived from the contract address
** and the block number.
**
** Arguments:
** -chainID the ID of the Ethereum network to connect to
** - vaultAddress the address of the Yvault043 contract to filter events from
** - activation the block number to start filtering events from
** - asyncMapLastReports the sync.Map to store the latest event's Rewards value in
** Returns nothing as asyncMapTransfers is updated via a pointer
*************************************************************************************************/
func filterUpdateRewards(
	chainID uint64,
	vaultAddress common.Address,
	opts *bind.FilterOpts,
	asyncMapLastReports *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := currentVault.FilterUpdateRewards(opts); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			eventKey := vaultAddress.Hex() + `-` + strconv.FormatUint(log.Event.Raw.BlockNumber, 10)
			asyncMapLastReports.LoadOrStore(eventKey, log.Event.Rewards.Hex())
		}
	} else {
		traces.
			Capture(`error`, `impossible to filterUpdateRewards for Yvault043 `+vaultAddress.Hex()).
			SetEntity(`vault`).
			SetExtra(`error`, err.Error()).
			SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
			SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
			SetTag(`vaultAddress`, vaultAddress.Hex()).
			Send()
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
	asyncMapLastReports := sync.Map{}
	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		opts := &bind.FilterOpts{Start: start, End: end}
		if start == 0 {
			opts = &bind.FilterOpts{Start: v.Activation, End: end}
		}

		go filterUpdateRewards(
			chainID,
			v.Address,
			opts,
			&asyncMapLastReports,
			wg,
		)
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
	asyncMapLastReports.Range(func(key, value interface{}) bool {
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
