package ethereum

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/utils"
)

// GetRPCURI returns the URI to use to connect to the node for a specific chainID
func GetRPCURI(chainID uint64) string {
	switch chainID {
	case 1:
		return utils.RPCURIFor1
	case 250:
		return utils.RPCURIFor250
	case 42161:
		return utils.RPCURIFor42161
	}
	return ``
}

// GetGraphURI returns the URI to use to connect to the graph provider for a specific chainID
func GetGraphURI(chainID uint64) string {
	switch chainID {
	case 1:
		return `https://api.thegraph.com/subgraphs/name/0xkofee/yearn-vaults-v2`
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
		return common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`)
	case 250:
		return common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`)
	case 42161:
		return common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`)
	}
	return common.HexToAddress(`0`)
}

// randomSigner will generate a fake signer for the ethereum client. We don't need to
// sign anything, but we need to provide a signer to the ethereum client.
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
