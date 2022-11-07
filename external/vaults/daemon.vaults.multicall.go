package vaults

import (
	"math"
	"math/big"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/contracts"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/common/store"
	"github.com/yearn/ydaemon/common/types/common"
	"github.com/yearn/ydaemon/internal/tokens"
)

// yearnVaultABI takes the ABI of the standard Yearn Vault contract and prepare it for the multicall
var yearnVaultABI, _ = contracts.YearnVaultMetaData.GetAbi()

func getPricePerShare(contractAddress common.Address) ethereum.Call {
	parsedData, _ := yearnVaultABI.Pack("pricePerShare")
	return ethereum.Call{
		Target:   contractAddress.Address,
		Abi:      yearnVaultABI,
		Method:   `pricePerShare`,
		CallData: parsedData,
		Name:     contractAddress.String(),
	}
}

// FetchVaultMulticallData will perform a multicall to get some specific data from on-chain for a specific list of vaults
func FetchVaultMulticallData(chainID uint64) {
	// First we connect to the multicall client, stored in memory via the initializer.
	caller := ethereum.MulticallClientForChainID[chainID]

	// Then, for each token listed for our current chainID, we prepare the calls
	var calls = make([]ethereum.Call, 0)
	for _, vault := range tokens.ListVaults(chainID) {
		calls = append(calls, getPricePerShare(common.FromAddress(vault.Address)))
	}

	if len(calls) == 0 {
		logs.Error("No vault found.")
		return
	}

	// Then, we execute the multicall and store the prices in the TokenPrices map
	maxBatch := math.MaxInt64
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	if Store.AggregatedVault[chainID] == nil {
		Store.AggregatedVault[chainID] = make(map[common.Address]*TAggregatedVault)
	}

	for _, vault := range tokens.ListVaults(chainID) {
		tokenAddress := common.FromAddress(vault.Address)
		pricePerShareRaw := response[vault.Address.String()+`pricePerShare`]
		pricePerShare := bigNumber.NewInt()
		if len(pricePerShareRaw) == 1 {
			pricePerShare = bigNumber.SetInt(pricePerShareRaw[0].(*big.Int))
		}
		if Store.AggregatedVault[chainID][tokenAddress] == nil {
			Store.AggregatedVault[chainID][tokenAddress] = &TAggregatedVault{
				Address:   tokenAddress,
				LegacyAPY: TLegacyAPIAPY{},
			}
		}
		Store.AggregatedVault[chainID][tokenAddress].SetPricePerShare(bigNumber.NewInt().Clone(pricePerShare))
		go store.SaveInDB(
			chainID,
			store.TABLES.VAULTS,
			tokenAddress.String(),
			Store.AggregatedVault[chainID][tokenAddress],
		)
	}
}
