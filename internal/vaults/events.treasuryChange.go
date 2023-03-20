package vaults

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/traces"
)

var treasuryChangesForChain = make(map[uint64]map[common.Address]map[uint64]common.Address)
var defaultTreasury = common.HexToAddress(`0x93a62da5a14c80f265dabc077fcee437b1a0efde`)

/**************************************************************************************************
** filterTreasuryChange filters the UpdateRewards events emitted by the specified Yvault043
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
func filterTreasuryChange(
	chainID uint64,
	vaultAddress common.Address,
	activation uint64,
	asyncMapLastReports *sync.Map,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	client := ethereum.GetRPC(chainID)
	currentVault, _ := contracts.NewYvault043(vaultAddress, client)
	if log, err := currentVault.FilterUpdateRewards(&bind.FilterOpts{Start: activation}); err == nil {
		for log.Next() {
			if log.Error() != nil {
				continue
			}
			eventKey := (vaultAddress.String() + `-` + strconv.FormatUint(log.Event.Raw.BlockNumber, 10))
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
** RetrieveTresuryChange retrieves all treasury change events for the provided vaults, waiting
** for the end of all goroutines before continuing.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - vaults: a map of all TVault to work on
**
** Returns:
** - a map of vaultAddress -> blockNumber -> treasuryAddress
**********************************************************************************************/
func RetrieveTresuryChange(
	chainID uint64,
	vaults map[common.Address]*TVault,
) map[common.Address]map[uint64]common.Address {
	trace := traces.Init(`app.indexer.vaults.treasury_change_events`).
		SetTag(`chainID`, strconv.FormatUint(chainID, 10)).
		SetTag(`rpcURI`, ethereum.GetRPCURI(chainID)).
		SetTag(`entity`, `vaults`).
		SetTag(`subsystem`, `daemon`)
	defer trace.Finish()

	/**********************************************************************************************
	** Concurrently retrieve all treasuryChange events for the vaults, waiting for the end
	** of all goroutines via the wg before continuing.
	**********************************************************************************************/
	asyncMapLastReports := sync.Map{}
	wg := &sync.WaitGroup{}
	for _, v := range vaults {
		wg.Add(1)
		go filterTreasuryChange(
			chainID,
			v.Address,
			v.Activation,
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

	treasuryChangesForChain[chainID] = treasuryChangeForVault
	return treasuryChangeForVault
}

/**************************************************************************************************
** This function is used to find the treasury address for a given vault at a particular block
** number on a specified chain.
** Parameters:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: address of the vault for which the treasury is being searched
** - blockNumber: the blockNumber at which the treasury is to be retrieved.
**
** The function first initializes a default treasury address. If there are no changes to the
** treasury for the given chainID, the default treasury address is returned. If there are no
** changes to the treasury for the given vaultAddress, the default treasury address is returned.
** Otherwise, the function iterates over the treasury changes for the given vaultAddress and checks
** if the block number for the current treasury is less than or equal to the given blockNumber. If
** it is, the treasury variable is updated with the current treasury address. The function then
** returns the updated treasury address.
**************************************************************************************************/
func findTreasuryAtBlock(chainID uint64, vaultAddress common.Address, blockNumber uint64) common.Address {
	treasury := defaultTreasury
	if _, ok := treasuryChangesForChain[chainID]; !ok {
		return treasury
	}
	if _, ok := treasuryChangesForChain[chainID][vaultAddress]; !ok {
		return treasury
	}
	for block, treasuryForBlock := range treasuryChangesForChain[chainID][vaultAddress] {
		if block <= blockNumber {
			treasury = treasuryForBlock
		}
	}
	return treasury
}

/**************************************************************************************************
** This function is used to find all of the treasury addresses that have been used for a given
** vault on a specified chain.
** Parameters:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: address of the vault for which the treasuries are being searched
**
** The function first initializes a default treasury address with the vaultAddress. If there are
** no changes to the treasury for the given chainID, the default treasury address is returned. If
** there are no changes to the treasury for the given vaultAddress, the default treasury address
** is returned. Otherwise, the function iterates over the treasury changes for the given
** vaultAddress and adds each treasury address to the treasuries list if it is not already in the
** list. The function then returns the list of updated treasuries.
**************************************************************************************************/
func findTreasuriesForVault(chainID uint64, vaultAddress common.Address) []common.Address {
	treasuries := []common.Address{vaultAddress}
	if _, ok := treasuryChangesForChain[chainID]; !ok {
		return treasuries
	}
	if _, ok := treasuryChangesForChain[chainID][vaultAddress]; !ok {
		return treasuries
	}
	for _, treasuryForBlock := range treasuryChangesForChain[chainID][vaultAddress] {
		if !helpers.Contains(treasuries, treasuryForBlock) {
			treasuries = append(treasuries, treasuryForBlock)
		}
	}
	return treasuries
}

/**************************************************************************************************
** This function is used to check if the treasury address for a given vault at a particular block
** number on a specified chain is the same as the vault address.
** Parameters:
** - chainID: the chain ID of the network we are working on
** - vaultAddress: address of the vault for which the treasury is being checked
** - blockNumber: the blockNumber at which the treasury is to be checked.
**
** The function first calls the findTreasuryAtBlock function to get the treasury address at the
** specified block. The function then checks if the retrieved treasury address is the same as the
** vault address. If it is, the function returns true, otherwise it returns false.
**************************************************************************************************/
func isTreasuryAtBlock(chainID uint64, vaultAddress common.Address, blockNumber uint64) bool {
	treasury := findTreasuryAtBlock(chainID, vaultAddress, blockNumber)
	return treasury == vaultAddress
}
