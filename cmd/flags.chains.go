package main

import (
	"strconv"
	"strings"

	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
)

type arrayFlags []string

var chainFlag arrayFlags

func (i *arrayFlags) String() string {
	defaultSupportedChains := env.SUPPORTED_CHAIN_IDS
	supportedChainsString := ``
	for _, chainID := range defaultSupportedChains {
		supportedChainsString += strconv.FormatUint(chainID, 10) + ` `
	}
	supportedChainsString = strings.Trim(supportedChainsString, ` `)
	return supportedChainsString
}
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func handleChainsInitialization() []uint64 {
	/**********************************************************************************************
	** chainFlag is an array of strings, so we need to convert it to an array of uint64, excluding
	** any invalid chain IDs or duplicates.
	**********************************************************************************************/
	for _, chainIDString := range chainFlag {
		chainID, err := strconv.ParseUint(chainIDString, 10, 64)
		if err != nil {
			continue
		}
		if !helpers.Contains(env.SUPPORTED_CHAIN_IDS, chainID) {
			continue
		}
		if helpers.Contains(chains, chainID) {
			continue
		}
		chains = append(chains, chainID)
	}

	if len(chains) == 0 {
		chains = env.SUPPORTED_CHAIN_IDS
	}
	return chains
}
