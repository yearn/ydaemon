package env

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

/**************************************************************************************************
** TestChainStructures tests the chain-related data structures to ensure they can be properly
** initialized and their fields accessed. This test validates:
** - TExtraStakingContracts structure initialization and field access
** - TChainCurve structure initialization and field access
** - TChainExtraURI structure initialization and field access
**
** These structures are fundamental for configuring chain-specific behavior.
**************************************************************************************************/
func TestChainStructures(t *testing.T) {
	// Test TExtraStakingContracts
	t.Run("TExtraStakingContracts", func(t *testing.T) {
		stakingContract := TExtraStakingContracts{
			VaultAddress:   common.HexToAddress("0x1111111111111111111111111111111111111111"),
			StakingAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
			Tag:            "JUICED",
		}

		if stakingContract.VaultAddress.Hex() != "0x1111111111111111111111111111111111111111" {
			t.Errorf("VaultAddress mismatch, got %s, expected 0x1111111111111111111111111111111111111111",
				stakingContract.VaultAddress.Hex())
		}

		if stakingContract.StakingAddress.Hex() != "0x2222222222222222222222222222222222222222" {
			t.Errorf("StakingAddress mismatch, got %s, expected 0x2222222222222222222222222222222222222222",
				stakingContract.StakingAddress.Hex())
		}

		if stakingContract.Tag != "JUICED" {
			t.Errorf("Tag mismatch, got %s, expected JUICED", stakingContract.Tag)
		}
	})

	// Test TChainCurve
	t.Run("TChainCurve", func(t *testing.T) {
		chainCurve := TChainCurve{
			RegistryAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
			FactoryAddress:  common.HexToAddress("0x4444444444444444444444444444444444444444"),
			PoolsURIs:       []string{"https://api.curve.fi/api/getPoolList", "https://api.curve.fi/api/getFactoryPoolList"},
			GaugesURI:       "https://api.curve.fi/api/getGauges",
		}

		if chainCurve.RegistryAddress.Hex() != "0x3333333333333333333333333333333333333333" {
			t.Errorf("RegistryAddress mismatch, got %s, expected 0x3333333333333333333333333333333333333333",
				chainCurve.RegistryAddress.Hex())
		}

		if chainCurve.FactoryAddress.Hex() != "0x4444444444444444444444444444444444444444" {
			t.Errorf("FactoryAddress mismatch, got %s, expected 0x4444444444444444444444444444444444444444",
				chainCurve.FactoryAddress.Hex())
		}

		if len(chainCurve.PoolsURIs) != 2 {
			t.Errorf("PoolsURIs length mismatch, got %d, expected 2", len(chainCurve.PoolsURIs))
		}

		if chainCurve.GaugesURI != "https://api.curve.fi/api/getGauges" {
			t.Errorf("GaugesURI mismatch, got %s, expected https://api.curve.fi/api/getGauges",
				chainCurve.GaugesURI)
		}
	})

	// Test TChainExtraURI
	t.Run("TChainExtraURI", func(t *testing.T) {
		chainExtraURI := TChainExtraURI{
			GammaMerklURI:      "https://api.gamma.xyz/merkl",
			GammaHypervisorURI: []string{"https://api.gamma.xyz/hypervisor1", "https://api.gamma.xyz/hypervisor2"},
			PendleCoreURI:      "https://api.pendle.finance/core",
		}

		if chainExtraURI.GammaMerklURI != "https://api.gamma.xyz/merkl" {
			t.Errorf("GammaMerklURI mismatch, got %s, expected https://api.gamma.xyz/merkl",
				chainExtraURI.GammaMerklURI)
		}

		if len(chainExtraURI.GammaHypervisorURI) != 2 {
			t.Errorf("GammaHypervisorURI length mismatch, got %d, expected 2",
				len(chainExtraURI.GammaHypervisorURI))
		}

		if chainExtraURI.PendleCoreURI != "https://api.pendle.finance/core" {
			t.Errorf("PendleCoreURI mismatch, got %s, expected https://api.pendle.finance/core",
				chainExtraURI.PendleCoreURI)
		}
	})
}

/**************************************************************************************************
** TestGetChainsFunction tests the GetChains function to ensure it returns the global CHAINS map.
** This test validates:
** - The function returns a map with the expected number of entries
** - The returned map entries match the global CHAINS map entries
**
** This is a basic test to ensure the accessor function works correctly.
**************************************************************************************************/
func TestGetChainsFunction(t *testing.T) {
	// Get the chains map
	chains := GetChains()

	// The map should be non-nil
	if chains == nil {
		t.Error("GetChains returned nil")
		return
	}

	// Check that the returned map has the same number of entries as CHAINS
	if len(chains) != len(CHAINS) {
		t.Errorf("GetChains returned map with %d entries, expected %d", len(chains), len(CHAINS))
	}

	// Check that each chain in CHAINS is present in the returned map
	for chainID, chainConfig := range CHAINS {
		if returnedChain, exists := chains[chainID]; !exists {
			t.Errorf("Chain ID %d missing from returned map", chainID)
		} else if returnedChain.ID != chainConfig.ID {
			t.Errorf("Chain ID mismatch for chain %d, got %d", chainID, returnedChain.ID)
		}
	}
}

/**************************************************************************************************
** TestGetChainFunction tests the GetChain function for retrieving specific chain configurations.
** This test validates:
** - The function returns the correct chain configuration for valid chain IDs
** - The function returns false for non-existent chain IDs
**
** This ensures that chain-specific configurations can be correctly retrieved.
**************************************************************************************************/
func TestGetChainFunction(t *testing.T) {
	// The function should work for all supported chains
	for _, chainID := range SUPPORTED_CHAIN_IDS {
		t.Run("Valid chain ID: "+string(rune(chainID)), func(t *testing.T) {
			chain, exists := GetChain(chainID)

			// The chain should exist
			if !exists {
				t.Errorf("GetChain reported chain ID %d as non-existent", chainID)
				return
			}

			// The returned chain should have the correct ID
			if chain.ID != chainID {
				t.Errorf("GetChain returned chain with ID %d, expected %d", chain.ID, chainID)
			}
		})
	}

	// Test with an invalid chain ID
	t.Run("Invalid chain ID", func(t *testing.T) {
		invalidChainID := uint64(999999)
		_, exists := GetChain(invalidChainID)

		// The chain should not exist
		if exists {
			t.Errorf("GetChain reported non-existent chain ID %d as existent", invalidChainID)
		}
	})
}

/**************************************************************************************************
** TestSupportedChainIDs tests the SUPPORTED_CHAIN_IDS global variable to ensure it contains
** all the chain IDs present in the CHAINS map. This test validates:
** - SUPPORTED_CHAIN_IDS has the same length as the CHAINS map
** - All chain IDs in CHAINS are present in SUPPORTED_CHAIN_IDS
**
** This ensures that the list of supported chain IDs is kept in sync with the actual chains map.
**************************************************************************************************/
func TestSupportedChainIDs(t *testing.T) {
	// Check that SUPPORTED_CHAIN_IDS has the expected number of elements
	if len(SUPPORTED_CHAIN_IDS) != len(CHAINS) {
		t.Errorf("SUPPORTED_CHAIN_IDS has %d elements, expected %d",
			len(SUPPORTED_CHAIN_IDS), len(CHAINS))
	}

	// Check that each chain ID in CHAINS is present in SUPPORTED_CHAIN_IDS
	for chainID := range CHAINS {
		found := false
		for _, supportedChainID := range SUPPORTED_CHAIN_IDS {
			if chainID == supportedChainID {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Chain ID %d is in CHAINS but not in SUPPORTED_CHAIN_IDS", chainID)
		}
	}

	// Check that each chain ID in SUPPORTED_CHAIN_IDS is present in CHAINS
	for _, supportedChainID := range SUPPORTED_CHAIN_IDS {
		if _, exists := CHAINS[supportedChainID]; !exists {
			t.Errorf("Chain ID %d is in SUPPORTED_CHAIN_IDS but not in CHAINS", supportedChainID)
		}
	}
}
