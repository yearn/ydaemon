package partnerFees

import (
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

/**************************************************************************************************
** getVaultsSupplyAtBlock will get the total supply of the vaults at the block number of the
** harvests. It will return a map of vaults with a map of vault, block number and total supply:
** map -> [vaultAddress][blockNumber] -> totalSupply
**************************************************************************************************/
func getVaultsSupplyAtBlock(
	allVaults []*models.TVault,
	allHarvests map[common.Address]map[common.Address][]models.THarvest,
) map[common.Address]map[uint64]*big.Int {
	syncMap := sync.Map{}
	for _, vault := range allVaults {
		vaultContract, _ := contracts.NewERC20(vault.Address, ethereum.GetRPC(vault.ChainID))

		/******************************************************************************************
		** First we transform the map of harvests into a slice of block numbers to be able to
		** parallelize the calls to the blockchain.
		******************************************************************************************/
		blockNumbers := []uint64{}
		for _, harvests := range allHarvests[vault.Address] {
			for _, harvest := range harvests {
				blockNumbers = append(blockNumbers, harvest.BlockNumber)
			}
		}

		/******************************************************************************************
		** Then we call the blockchain to get the total supply of the vault at the block number of
		** the harvests, all in gorooutines, and we store the result in a sync map.
		******************************************************************************************/
		wg := sync.WaitGroup{}
		for _, blockNumber := range blockNumbers {
			wg.Add(1)
			go func(blockNumber uint64) {
				defer wg.Done()
				totalSupply, err := vaultContract.TotalSupply(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNumber))})
				if err != nil {
					logs.Error(err)
					return
				}
				eventKey := vault.Address.Hex() + `-` + strconv.FormatUint(blockNumber, 10)
				syncMap.Store(eventKey, totalSupply)
			}(blockNumber)
		}
		wg.Wait()
	}

	/******************************************************************************************
	** As we are working with a sync map, we need to transform it into a regular map to be able
	** to return it. We just range over the sync map and store the values in a regular map of
	** map -> [vaultAddress][blockNumber] -> totalSupply
	******************************************************************************************/
	vaultsSupplyAtBlock := make(map[common.Address]map[uint64]*big.Int)
	syncMap.Range(func(key, value interface{}) bool {
		eventKey := strings.Split(key.(string), `-`)
		vaultAddress := common.HexToAddress(eventKey[0])
		blockNumber, _ := strconv.ParseUint(eventKey[1], 10, 64)

		if vaultsSupplyAtBlock[vaultAddress] == nil {
			vaultsSupplyAtBlock[vaultAddress] = make(map[uint64]*big.Int)
		}
		vaultsSupplyAtBlock[vaultAddress][blockNumber] = value.(*big.Int)
		return true
	})

	return vaultsSupplyAtBlock
}
