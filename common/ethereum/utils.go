package ethereum

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
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
		return env.RPC_ENDPOINTS[1]
	case 10:
		return env.RPC_ENDPOINTS[10]
	case 250:
		return env.RPC_ENDPOINTS[250]
	case 42161:
		return env.RPC_ENDPOINTS[42161]
	}
	return ``
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
