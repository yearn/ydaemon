package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
)

/**************************************************************************************************
** RPC stores Ethereum client connections for each chain ID.
** This map allows for easy access to RPC clients across the application.
**************************************************************************************************/
var RPC = map[uint64]*ethclient.Client{}

/**************************************************************************************************
** WS stores WebSocket client connections for each chain ID.
** This map allows for easy access to WebSocket clients across the application.
**************************************************************************************************/
var WS = map[uint64]*ethclient.Client{}

/**************************************************************************************************
** GetRPC returns the current RPC connection for a specific chain.
**
** This function provides access to the ethclient for a given chain ID, which can be used
** to interact with the blockchain via JSON-RPC.
**
** @param chainID The ID of the blockchain to get the connection for
** @return *ethclient.Client The Ethereum client for the specified chain
**************************************************************************************************/
func GetRPC(chainID uint64) *ethclient.Client {
	return RPC[chainID]
}

/**************************************************************************************************
** GetRPCURI returns the URI used to connect to the node for a specific chain ID.
**
** This function retrieves the RPC URI from the environment configuration for the specified chain.
** The URI is required to establish connections with the blockchain node.
**
** @param chainID The ID of the blockchain to get the URI for
** @return string The RPC URI for the specified chain, or empty string if chain not found
**************************************************************************************************/
func GetRPCURI(chainID uint64) string {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return ""
	}
	return chain.RpcURI
}

/**************************************************************************************************
** GetWSEnvURI returns the WebSocket URI used to connect to the node for a specific chain ID.
**
** This function first checks for a chain-specific WebSocket URI in the environment variables
** (format: WS_URI_FOR_[chainID]). If not found, it falls back to the chain's RPC URI from
** the environment configuration.
**
** @param chainID The ID of the blockchain to get the WebSocket URI for
** @return string The WebSocket URI for the specified chain, or empty string if chain not found
**************************************************************************************************/
func GetWSEnvURI(chainID uint64) string {
	wsFromEnv, exists := os.LookupEnv("WS_URI_FOR_" + strconv.FormatUint(chainID, 10))
	if !exists {
		chain, ok := env.GetChain(chainID)
		if !ok {
			return ""
		}
		return chain.RpcURI
	}
	return wsFromEnv
}

/**************************************************************************************************
** MulticallClientForChainID holds multicall client instances for each chain ID.
**
** This map stores the multicall clients that are used to batch multiple contract calls
** into a single RPC request, improving efficiency for read operations.
**************************************************************************************************/
var MulticallClientForChainID = make(map[uint64]TEthMultiCaller)

/**************************************************************************************************
** randomSigner generates a fake signer for the Ethereum client.
**
** This function creates a TransactOpts instance with a randomly generated private key,
** which can be used for transaction signing. The signer is configured not to actually
** send transactions (NoSend = true) and has maximum gas limits with zero gas pricing.
**
** @return *bind.TransactOpts A transaction signer with random credentials and zero gas settings
**************************************************************************************************/
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

/**************************************************************************************************
** GetWSClient returns a WebSocket connection for a specific chain.
**
** This function attempts to establish a WebSocket connection for the specified chain ID.
** It performs URL transformations to convert HTTP(S) URLs to WS(S) as needed, handling
** special cases for different providers (Infura, Chainstack, etc).
**
** If a connection already exists in the WS map, it returns that connection.
** Otherwise, it creates a new connection with a 10-second timeout.
**
** @param chainID The ID of the blockchain to get the WebSocket connection for
** @param shouldRetry Whether to retry the connection attempt if it times out
** @return *ethclient.Client The WebSocket client for the specified chain
** @return error An error if the connection fails
**************************************************************************************************/
func GetWSClient(chainID uint64, shouldRetry bool) (*ethclient.Client, error) {
	chain, ok := env.GetChain(chainID)
	if !ok {
		return nil, errors.New("chain not found")
	}
	if !chain.CanUseWebsocket {
		return nil, errors.New("chain cannot use websocket")
	}

	if WS[chainID] == nil {
		uriString := GetWSEnvURI(chainID)
		uri, _ := url.Parse(uriString)
		if strings.HasPrefix(uri.Host, `nd-`) {
			uri.Host = strings.Replace(uri.Host, `nd-`, `ws-nd-`, 1)
		}
		if strings.Contains(uri.Host, `infura.io`) && uri.Scheme == `https` {
			uri.Path = strings.Replace(uri.Path, `v3`, `ws/v3`, 1)
		}
		if strings.Contains(uri.Host, `chainstack.com`) && uri.Scheme == `https` {
			uri.Path = `ws` + uri.Path
		}

		switch uri.Scheme {
		case `https`:
			uri.Scheme = `wss`
		case `http`:
			uri.Scheme = `ws`
		}

		// contextTimeout, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
		// defer cancel()

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := ethclient.DialContext(ctx, uri.String())
		if err != nil {
			if shouldRetry && err.Error() == `i/o timeout` {
				logs.Warning(fmt.Sprintf("Chain %d - Timeout while opening WS client with RPC %s: %v",
					chainID, uri.String(), err))
				return GetWSClient(chainID, false)
			}
			logs.Error(fmt.Sprintf("Chain %d - Error while opening WS client with RPC %s: %v",
				chainID, uri.String(), err))
			return nil, err
		}
		WS[chainID] = client
	}
	return WS[chainID], nil
}
