package vaults

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/fetcher"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

/**************************************************************************************************
** TRotkiVaultToken represents the underlying token for a vault in Rotki format.
**
** This structure contains simplified token information specifically formatted for Rotki,
** a portfolio tracking application. It includes only the essential token properties needed
** for integration with Rotki's tracking system.
**
** @field Address string - The token's Ethereum address
** @field Name string - The human-readable name of the token
** @field Symbol string - The token's symbol (e.g., DAI, USDC)
** @field Decimals uint64 - The number of decimal places the token uses
** @field Icon string - URL to the token's icon image
**************************************************************************************************/
type TRotkiVaultToken struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint64 `json:"decimals"`
	Icon     string `json:"icon"`
}

/**************************************************************************************************
** TRotkiVaults represents vault information in Rotki format.
**
** This structure contains the vault data specifically formatted for integration with Rotki,
** a portfolio tracking application. It presents a streamlined representation with only
** the information necessary for tracking vault positions.
**
** @field Address string - The vault's Ethereum address
** @field Name string - The human-readable name of the vault
** @field Symbol string - The vault's symbol
** @field Decimals uint64 - The vault's token decimal places
** @field Version string - The vault's version identifier
** @field Icon string - URL to the vault's icon image
** @field FeaturingScore float64 - Internal score used for sorting (not exported to JSON)
** @field Type models.TTokenType - The vault's type classification
** @field Token TRotkiVaultToken - The underlying token information
** @field StakingAddress string - Optional staking contract address if available
**************************************************************************************************/
type TRotkiVaults struct {
	Address        string            `json:"address"`
	Name           string            `json:"name"`
	Symbol         string            `json:"symbol"`
	Decimals       uint64            `json:"decimals"`
	Version        string            `json:"version"`
	Icon           string            `json:"icon"`
	FeaturingScore float64           `json:"-"` // Not exported
	Type           models.TTokenType `json:"type"`
	Token          TRotkiVaultToken  `json:"token"`
	StakingAddress string            `json:"staking,omitempty"`
}

/**************************************************************************************************
** GetVaultsForRotki retrieves vaults formatted specifically for Rotki integration.
**
** This endpoint provides vault data in a format specifically designed for the Rotki portfolio
** tracking application. It includes a streamlined representation with only the information
** necessary for tracking vault positions and returns data in a format compatible with Rotki's
** integration needs.
**
** The endpoint accepts parameters for sorting, pagination, and filtering by chain IDs:
** - orderBy: Field to sort results by (default: 'featuringScore')
** - orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
** - skip/limit: Pagination controls (defaults: skip 0, limit 200)
** - chainIDs: Comma-separated list of chain IDs to include (default: all supported chains)
**
** Endpoint: GET /vaults/rotki
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return []TRotkiVaults - The formatted list of vaults for Rotki
**************************************************************************************************/
func (y Controller) GetVaultsForRotki(c *gin.Context) []TRotkiVaults {
	/** 🔵 - Yearn *************************************************************************************
	** Validate and process query parameters for sorting, pagination, and chain filtering.
	**
	** orderBy: Field to sort results by (default: 'featuringScore')
	** orderDirection: Sort direction, 'asc' or 'desc' (default: 'asc')
	** skip: Number of records to skip for pagination (default: 0)
	** limit: Maximum number of records to return (default: 200)
	** chainIDs: Comma-separated list of chain IDs to include (default: all supported chains)
	**************************************************************************************************/
	// Define valid fields for sorting
	validOrderFields := []string{
		"featuringScore", "name", "symbol", "decimals", "type", "version", "address",
	}
	orderBy := ValidateStringChoiceQuery(c, "orderBy", "featuringScore", validOrderFields, "GetVaultsForRotki")

	// Define valid sort directions
	validDirections := []string{"asc", "desc"}
	orderDirection := ValidateStringChoiceQuery(c, "orderDirection", "asc", validDirections, "GetVaultsForRotki")

	// Validate pagination parameters
	skip := ValidateNumericQuery(c, "skip", 0, 0, 10000, "GetVaultsForRotki")
	limit := ValidateNumericQuery(c, "limit", 200, 1, 1000, "GetVaultsForRotki")

	// Process chain IDs parameter
	chainsStr := strings.Split(getQueryParam(c, `chainIDs`), `,`)
	chains := []uint64{}
	if len(chainsStr) == 0 || (len(chainsStr) == 1 && chainsStr[0] == ``) {
		chains = env.SUPPORTED_CHAIN_IDS
	} else {
		for _, chainStr := range chainsStr {
			chain, ok := helpers.AssertChainID(chainStr)
			if !ok {
				HandleError(c, fmt.Errorf("invalid chain ID: %s", chainStr),
					http.StatusBadRequest, "Invalid chain ID", "GetVaultsForRotki")
				continue
			}
			chains = append(chains, chain)
		}
		// If no valid chains were provided, use all supported chains
		if len(chains) == 0 {
			chains = env.SUPPORTED_CHAIN_IDS
		}
	}

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code initializes an empty array of TRotkiVaults objects to 'data'.
	** 'data' will store the final list of vaults to be returned in the response
	**
	** The first for loop iterates over the 'chains' array of chain IDs. For each chain ID,
	** it calls the 'ListVaults' function to get a list of all vaults for that chain.
	**
	** The second nested for loop then iterates over each vault in 'vaultsForChain'. If the vault's
	** address is not in the 'BLACKLISTED_VAULTS' for the current chain ID, it attempts to process the vault.
	**
	** For each valid vault, it retrieves the token and staking information, and constructs a
	** TRotkiVaults object with the appropriate data. This object is then appended to the 'data' slice.
	**
	** This process effectively gathers all non-blacklisted vaults across all specified chains
	** and formats them in the Rotki-specific structure.
	**************************************************************************************************/
	data := []TRotkiVaults{}
	for _, chainID := range chains {
		vaultsForChain, _ := storage.ListVaults(chainID)
		for _, currentVault := range vaultsForChain {
			/******************************************************************************************
			** We want to ignore all non Yearn vaults
			******************************************************************************************/
			chain, ok := env.GetChain(chainID)
			if !ok {
				continue
			}
			if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
				continue
			}
			newVault, err := CreateExternalVault(currentVault)
			if err != nil {
				continue
			}

			stakingData := assignStakingData(chainID, currentVault.Address)
			stakingAddress := ""
			if stakingData.Available {
				stakingAddress = stakingData.Address
			}

			vaultToken, ok := storage.GetERC20(chainID, currentVault.Address)
			if !ok {
				continue
			}
			name, displayName, _ := fetcher.BuildVaultNames(currentVault, currentVault.Metadata.DisplayName)
			symbol, displaySymbol, _ := fetcher.BuildVaultSymbol(currentVault, currentVault.Metadata.DisplaySymbol)

			vaultName := displayName
			if vaultName == "" {
				vaultName = name
			}
			vaultSymbol := displaySymbol
			if vaultSymbol == "" {
				vaultSymbol = symbol
			}

			APRAsFloat := 0.0
			if newVault.APR.NetAPR != nil {
				APRAsFloat, _ = newVault.APR.NetAPR.Float64()
			}
			newVault.FeaturingScore = newVault.TVL.TVL * APRAsFloat
			if newVault.Details.IsHighlighted {
				newVault.FeaturingScore = newVault.FeaturingScore * HIGHLIGHTING_MULTIPLIER
			}

			vaultForRotki := TRotkiVaults{
				Address:        newVault.Address,
				Name:           vaultName,
				Symbol:         vaultSymbol,
				Decimals:       newVault.Decimals,
				Version:        newVault.Version,
				Icon:           newVault.Icon,
				Type:           newVault.Type,
				FeaturingScore: newVault.FeaturingScore,
				Token: TRotkiVaultToken{
					Address:  vaultToken.Address.String(),
					Name:     vaultToken.Name,
					Symbol:   vaultToken.Symbol,
					Decimals: vaultToken.Decimals,
					Icon:     vaultToken.Icon,
				},
			}

			if stakingAddress != "" {
				vaultForRotki.StakingAddress = stakingAddress
			}

			data = append(data, vaultForRotki)
		}
	}

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code sorts the 'data' slice based on the 'orderBy' and 'orderDirection'
	** parameters. The 'SortBy' function from the 'sort' package is used for this purpose.
	**
	** After sorting, it applies pagination to the 'data' slice based on the 'skip' and 'limit'
	** parameters. If 'skip' is greater than the length of the 'data' slice, an empty slice is returned.
	** Otherwise, the 'data' slice is updated to only include the vaults starting from the 'skip' index
	** and up to 'limit' vaults.
	**
	** This effectively returns only the vaults for the requested pagination with the specified limit.
	**************************************************************************************************/
	sort.SortBy(orderBy, orderDirection, data)
	if skip >= uint64(len(data)) {
		return []TRotkiVaults{}
	}
	end := skip + limit
	if end > uint64(len(data)) {
		end = uint64(len(data))
	}
	data = data[skip:end]

	return data
}

/**************************************************************************************************
** CountVaultsForRotki returns the total count of vaults available for Rotki integration.
**
** This endpoint provides a simple count of all vaults that would be available through the
** GetVaultsForRotki endpoint. It's useful for pagination purposes and for Rotki to understand
** the total dataset size without retrieving all vault data.
**
** The endpoint accepts chainIDs as a filtering parameter:
** - chainIDs: Comma-separated list of chain IDs to include (default: all supported chains)
**
** Endpoint: GET /vaults/rotki/count
**
** @param c *gin.Context - The Gin context containing the HTTP request
** @return void - Response is sent directly via Gin with JSON containing the count
**************************************************************************************************/
func (y Controller) CountVaultsForRotki(c *gin.Context) {
	/** 🔵 - Yearn *************************************************************************************
	** chainsStr: A string that represents the chain IDs for which the vaults are to be returned. It is
	** obtained from the 'chainIDs' query parameter in the request. The string is split by commas to
	** obtain an array of chain IDs.
	**
	** chains: An array of uint64 values that represents the chain IDs for which the vaults are to be
	** returned. If the 'chains' query parameter is not provided or is empty, this array defaults to
	** all supported chain IDs.
	**
	** The 'chains' array is populated by iterating over the 'chainsStr' array and converting each
	** chain ID to a uint64 value. If the conversion is not successful, the chain ID is ignored.
	**
	** The 'chains' array is used to filter the vaults that are returned in the response.
	**************************************************************************************************/
	chainsStr := strings.Split(getQueryParam(c, `chainIDs`), `,`)
	chains := []uint64{}
	if len(chainsStr) == 0 || (len(chainsStr) == 1 && chainsStr[0] == ``) {
		chains = env.SUPPORTED_CHAIN_IDS
	} else {
		for _, chainStr := range chainsStr {
			chain, ok := helpers.AssertChainID(chainStr)
			if !ok {
				continue
			}
			chains = append(chains, chain)
		}
	}

	count := 0
	for _, chainID := range chains {
		vaultsForChain, _ := storage.ListVaults(chainID)
		for _, currentVault := range vaultsForChain {
			/******************************************************************************************
			** We want to ignore all non Yearn vaults
			******************************************************************************************/
			chain, ok := env.GetChain(chainID)
			if !ok {
				continue
			}
			if helpers.Contains(chain.BlacklistedVaults, currentVault.Address) {
				continue
			}
			_, err := CreateExternalVault(currentVault)
			if err != nil {
				continue
			}

			count += 1
		}
	}

	c.JSON(http.StatusOK, gin.H{
		`numberOfVaults`: count,
	})
}
