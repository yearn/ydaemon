package main

import (
	"strconv"
	"strings"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/helpers"
)

type arrayFlags string

func handleChainsInitialization(rawChains *string) []uint64 {
	/**********************************************************************************************
	** chainFlag is an array of strings, so we need to convert it to an array of uint64, excluding
	** any invalid chain IDs or duplicates.
	**********************************************************************************************/
	chainStr := strings.Split(*rawChains, `,`)

	for _, chainIDString := range chainStr {
		chainID, err := strconv.ParseUint(chainIDString, 10, 64)
		if err != nil {
			continue
		}
		if _, ok := env.CHAINS[chainID]; !ok {
			continue
		}
		if helpers.Contains(chains, chainID) {
			continue
		}
		chains = append(chains, chainID)
		ethereum.GetWSClient(chainID)
	}

	if len(chains) == 0 {
		chains = env.SUPPORTED_CHAIN_IDS
	}
	return chains
}
