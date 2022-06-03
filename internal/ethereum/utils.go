package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/majorfi/ydaemon/internal/utils"
)

// GetClient initializes a connection to the Ethereum Network
func GetClient() *ethclient.Client {
	nodeURL := utils.EthNodeURI
	client, _ := ethclient.Dial(nodeURL)
	return client
}
