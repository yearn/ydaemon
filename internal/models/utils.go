package models

import "github.com/ethereum/go-ethereum/common"

type TContractData struct {
	Address    common.Address // Address of the contract
	Block      uint64         // Block number where the contract was deployed
	Version    uint64         // Version of the contract. May be empty.
	Activation uint64         // Timestamp of the contract activation. May be empty.
}
