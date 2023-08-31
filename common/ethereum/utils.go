package ethereum

import (
	"context"
	"math"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

var RPC = map[uint64]*ethclient.Client{}
var WS = map[uint64]*ethclient.Client{}

// GetRPC returns the current connection for a specific chain
func GetRPC(chainID uint64) *ethclient.Client {
	return RPC[chainID]
}

// GetRPCURI returns the URI to use to connect to the node for a specific chainID
func GetRPCURI(chainID uint64) string {
	return env.RPC_ENDPOINTS[chainID]
}

// GetWSEnvURI returns the URI to use to connect to the node for a specific chainID
func GetWSEnvURI(chainID uint64) string {
	wsFromEnv, exists := os.LookupEnv("WS_URI_FOR_" + strconv.FormatUint(chainID, 10))
	if !exists {
		return env.RPC_ENDPOINTS[chainID]
	} else {
		return wsFromEnv
	}
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

// GetWSClient returns the current ws connection for a specific chain
func GetWSClient(chainID uint64) (*ethclient.Client, error) {
	if WS[chainID] == nil {
		uriString := GetWSEnvURI(chainID)
		// uriString := GetRPCURI(chainID)
		// uriString := `http://localhost:8545`
		uri, _ := url.Parse(uriString)
		if uri.Scheme == `https` {
			uri.Scheme = `wss`
		} else {
			uri.Scheme = `ws`
		}

		client, err := ethclient.Dial(uri.String())
		if err != nil {
			logs.Error(`error while openning ws client for chain ` + strconv.FormatUint(chainID, 10) + ` with RPC ` + uri.String())
			return nil, err
		}
		WS[chainID] = client
	}
	return WS[chainID], nil
}
