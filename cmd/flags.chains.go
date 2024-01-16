package main

import (
	"strconv"
	"strings"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
)

func handleChainsInitialization(rawChains *string) []uint64 {
	/**********************************************************************************************
	** chainFlag is an array of strings, so we need to convert it to an array of uint64, excluding
	** any invalid chain IDs or duplicates.
	**********************************************************************************************/
	chainStr := strings.Split(*rawChains, `,`)

	for _, chainIDString := range chainStr {
		chainID, err := strconv.ParseUint(chainIDString, 10, 64)
		if err != nil {
			logs.Error(`Invalid chain ID: ` + chainIDString)
			continue
		}
		if _, ok := env.CHAINS[chainID]; !ok {
			logs.Error(`Unsupported chain ID: ` + chainIDString)
			continue
		}
		if helpers.Contains(chains, chainID) {
			logs.Error(`Duplicate chain ID: ` + chainIDString)
			continue
		}
		chains = append(chains, chainID)
	}

	if len(chains) == 0 {
		chains = env.SUPPORTED_CHAIN_IDS
	}
	return chains
}
