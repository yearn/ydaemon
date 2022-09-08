package ethereum

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/internal/utils/helpers"
	"github.com/yearn/ydaemon/internal/utils/logs"
)

var RPC = map[uint64]*ethclient.Client{}

// GetRPC returns the current connection for a specific chain
func GetRPC(chainID uint64) *ethclient.Client {
	return RPC[chainID]
}

// GetRPCURI returns the URI to use to connect to the node for a specific chainID
func GetRPCURI(chainID uint64) string {
	switch chainID {
	case 1:
		return helpers.RPCURIFor1
	case 10:
		return helpers.RPCURIFor10
	case 250:
		return helpers.RPCURIFor250
	case 42161:
		return helpers.RPCURIFor42161
	}
	return ``
}

// GetGraphURI returns the URI to use to connect to the graph provider for a specific chainID
func GetGraphURI(chainID uint64) string {
	switch chainID {
	case 1:
		return `https://api.thegraph.com/subgraphs/name/rareweasel/yearn-vaults-v2-subgraph-mainnet`
	case 10:
		return `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-optimism`
	case 250:
		return `https://api.thegraph.com/subgraphs/name/bsamuels453/yearn-fantom-validation-grafted`
	case 42161:
		return `https://api.thegraph.com/subgraphs/name/yearn/yearn-vaults-v2-arbitrum`
	}
	return ``
}

// GetLensAddress returns the address of the Lens oracle for a specific chainID
func GetLensAddress(chainID uint64) common.Address {
	switch chainID {
	case 1:
		return common.HexToAddress(`0x83d95e0D5f402511dB06817Aff3f9eA88224B030`)
	case 10:
		return common.HexToAddress(`0xB082d9f4734c535D9d80536F7E87a6f4F471bF65`)
	case 250:
		return common.HexToAddress(`0x57AA88A0810dfe3f9b71a9b179Dd8bF5F956C46A`)
	case 42161:
		return common.HexToAddress(`0x043518AB266485dC085a1DB095B8d9C2Fc78E9b9`)
	}
	return common.HexToAddress(`0`)
}

// GetMulticallAddress returns the address of the multicall2 contract for a specific chainID
func GetMulticallAddress(chainID uint64) common.Address {
	switch chainID {
	case 1:
		return common.HexToAddress(`0x5ba1e12693dc8f9c48aad8770482f4739beed696`)
	case 10:
		return common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`)
	case 250:
		return common.HexToAddress(`0x470ADB45f5a9ac3550bcFFaD9D990Bf7e2e941c9`)
	case 42161:
		return common.HexToAddress(`0x842eC2c7D803033Edf55E478F461FC547Bc54EB2`)
	}
	return common.HexToAddress(`0`)
}

// MulticallClientForChainID holds the multicall client for a specific chainID
var MulticallClientForChainID = make(map[uint64]TEthMultiCaller)

// randomSigner will generate a fake signer for the ethereum client.
// We don't need to sign anything, but we need to provide a signer
// to the ethereum client.
func randomSigner() *bind.TransactOpts {
	privateKey, err := crypto.GenerateKey()
	if err != nil { // Should never happen
		logs.Error("Failed to generate key")
		time.Sleep(time.Second)
		return randomSigner()
	}

	signer, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	signer.NoSend = true
	signer.Context = context.Background()
	signer.GasLimit = math.MaxUint64
	signer.GasFeeCap = big.NewInt(0)
	signer.GasTipCap = big.NewInt(0)
	signer.GasPrice = big.NewInt(0)
	return signer
}

func init() {
	for _, chainID := range helpers.SUPPORTED_CHAIN_IDS {
		client, err := ethclient.Dial(GetRPCURI(chainID))
		if err != nil {
			logs.Error("Failed to connect to Ethereum node")
			continue
		}
		RPC[chainID] = client
	}

	MulticallClientForChainID[1] = NewMulticall(
		GetRPCURI(1),
		GetMulticallAddress(1),
	)
	MulticallClientForChainID[10] = NewMulticall(
		GetRPCURI(10),
		GetMulticallAddress(10),
	)
	MulticallClientForChainID[250] = NewMulticall(
		GetRPCURI(250),
		GetMulticallAddress(250),
	)
	MulticallClientForChainID[42161] = NewMulticall(
		GetRPCURI(42161),
		GetMulticallAddress(42161),
	)
}
