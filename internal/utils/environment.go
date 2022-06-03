package utils

import (
	"log"
	"os"
	"strconv"
)

var (
	// EthNodeURI : Uri of the ethereum node
	EthNodeURI string
	// EthPrivateKey : PrivateKey used for this demonstration
	EthPrivateKey string
	// EthChainID : chain on which we will work
	EthChainID int64
	// StartBlock : block from which we should start working
	StartBlock uint64
	// DSN : id to connect to the database
	DSN string
)

func initEthCore() {
	var exists bool
	var err error

	EthNodeURI, exists = os.LookupEnv("ETH_NODE_URI")
	if !exists {
		log.Fatal("ETH_NODE_URI environment variable not set")
	}
	EthChainIDString, exists := os.LookupEnv("ETH_CHAIN_ID")
	if !exists {
		log.Fatal("ETH_CHAIN_ID environment variable not set")
	}
	EthChainID, err = strconv.ParseInt(EthChainIDString, 10, 64)
	if err != nil {
		log.Fatal("ETH_CHAIN_ID environment variable not set")
	}
	StartBlockString, exists := os.LookupEnv("START_BLOCK")
	if !exists {
		log.Fatal("START_BLOCK environment variable not set")
	}
	StartBlock, err = strconv.ParseUint(StartBlockString, 10, 64)
	if err != nil {
		log.Fatal("START_BLOCK environment variable not set")
	}
}

// InitEnvironment initializes variables from environment
func InitEnvironment() {
	initEthCore()
}
