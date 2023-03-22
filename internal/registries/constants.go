package registries

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/utils"
)

type TRegistry struct {
	Address    common.Address
	Version    uint64
	Activation uint64
}

// YEARN_REGISTRIES is the list of registries used by Yearn across all the supported chains, with the version and the activation block
var YEARN_REGISTRIES = map[uint64][]TRegistry{
	1: {
		{Address: common.HexToAddress("0xe15461b18ee31b7379019dc523231c57d1cbc18c"), Version: 1, Activation: 11563389},
		{Address: common.HexToAddress("0x50c1a2eA0a861A967D9d0FFE2AE4012c2E053804"), Version: 2, Activation: 12045555},
		{Address: common.HexToAddress("0xaF1f5e1c19cB68B30aAD73846eFfDf78a5863319"), Version: 3, Activation: 16215519},
	},
	10: {
		{Address: common.HexToAddress("0x81291ceb9bB265185A9D07b91B5b50Df94f005BF"), Version: 3, Activation: 22450349},
		{Address: common.HexToAddress("0x79286Dd38C9017E5423073bAc11F53357Fc5C128"), Version: 3, Activation: 22451152},
	},
	250: {
		{Address: common.HexToAddress("0x727fe1759430df13655ddb0731dE0D0FDE929b04"), Version: 2, Activation: 18455565},
	},
	42161: {
		{Address: common.HexToAddress("0x3199437193625DCcD6F9C9e98BDf93582200Eb1f"), Version: 2, Activation: 4841854},
	},
}

// EXTRA_VAULTS is a list of vaults that are not registered in the registries, but are still used by Yearn
var EXTRA_VAULTS = map[uint64][]utils.TVaultsFromRegistry{
	1:  {},
	10: {},
	250: {
		{
			//yvMIM, alone in it's own registry, not work registering and listening to it
			RegistryAddress: common.HexToAddress(`0x265F7b1413F6B06654746cf2485082182389A5d0`),
			VaultsAddress:   common.HexToAddress(`0x0A0b23D9786963DE69CB2447dC125c49929419d8`),
			TokenAddress:    common.HexToAddress(`0x82f0b8b456c1a451378467398982d4834b6829c1`),
			APIVersion:      `0.4.3`,
			BlockNumber:     18309707,
			Activation:      18302860,
			ManagementFee:   200,
			BlockHash:       common.HexToHash(`0x00009ee300000d281b4c0169bb3320b32f435e3fd830fe1625adcfd4cf6410cb`),
			TxIndex:         0,
			LogIndex:        0,
			Type:            utils.VaultTypeStandard,
		},
	},
	42161: {},
}
