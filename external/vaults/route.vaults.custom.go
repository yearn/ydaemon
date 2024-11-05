package vaults

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

type TRotkiVaultToken struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint64 `json:"decimals"`
	Icon     string `json:"icon"`
}

type TRotkiVaultStaking struct {
	Address   string `json:"address"`
	Available bool   `json:"available"`
}

type TRotkiVaults struct {
	Address        string             `json:"address"`
	Name           string             `json:"name"`
	Symbol         string             `json:"symbol"`
	Decimals       uint64             `json:"decimals"`
	Version        string             `json:"version"`
	Icon           string             `json:"icon"`
	FeaturingScore float64            `json:"-"` // Not exported
	Type           models.TTokenType  `json:"type"`
	Token          TRotkiVaultToken   `json:"token"`
	Staking        TRotkiVaultStaking `json:"staking"`
}

/**************************************************************************************************
** GetVaultsForRotki is a gin handler function to retrieve all the vaults matching the
** inclusion.IsYearn filter.
**************************************************************************************************/
func (y Controller) GetVaultsForRotki(c *gin.Context) []TRotkiVaults {
	/** 🔵 - Yearn *************************************************************************************
	** orderBy: A string that determines the order in which the vaults are returned. It is obtained
	** from the 'orderBy' query parameter in the request. If the parameter is not provided,
	** it defaults to 'details.order'.
	**
	** orderDirection: A string that determines the direction of the ordering of the vaults. It is
	** obtained from the 'orderDirection' query parameter in the request. If the parameter is not
	** provided, it defaults to 'asc'.
	**************************************************************************************************/
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `featuringScore`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)

	/** 🔵 - Yearn *************************************************************************************
	** page: A uint64 value that represents the page number for pagination. It is obtained from the
	** 'page' query parameter in the request. If the parameter is not provided, it defaults to 1.
	**
	** skip: A uint64 value that represents the number of vaults to skip. It is obtained from the
	** 'skip' query parameter in the request. If the parameter is not provided, it defaults to 0.
	**************************************************************************************************/
	skip := helpers.SafeStringToUint64(getQuery(c, `skip`), 0)
	limit := helpers.SafeStringToUint64(getQuery(c, `limit`), 200)

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
	chainsStr := strings.Split(getQuery(c, `chainIDs`), `,`)
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

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code initializes two empty slices, 'data' and 'allVaults'. 'data' will
	** store the final list of vaults to be returned in the response, while 'allVaults' is a temporary
	** storage for all vaults across all chains.
	**
	** The first for loop iterates over the 'SUPPORTED_CHAIN_IDS' from the environment variables. For
	** each chain ID, it calls the 'ListVaults' function to get a list of all vaults for that chain.
	**
	** The second nested for loop then iterates over each vault in 'vaultsForChain'. If the vault's
	** address is not in the 'BLACKLISTED_VAULTS' for the current chain ID, it is appended to the
	** 'allVaults' slice.
	** It will also check if, based on the provided parameters, the vault should be included in the
	** response. If not, it will be skipped.
	**
	** This process effectively gathers all vaults across all supported chains, excluding those that
	** are blacklisted, and stores them in 'allVaults' for further processing.
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
			newVault, err := NewVault().AssignTVault(currentVault)
			if err != nil {
				continue
			}

			APRAsFloat := 0.0
			if newVault.APR.NetAPR != nil {
				APRAsFloat, _ = newVault.APR.NetAPR.Float64()
			}
			newVault.FeaturingScore = newVault.TVL.TVL * APRAsFloat
			if newVault.Details.IsHighlighted {
				newVault.FeaturingScore = newVault.FeaturingScore * 1e18
			}

			vaultForRotki := TRotkiVaults{
				Address:        newVault.Address,
				Name:           newVault.Name,
				Symbol:         newVault.Symbol,
				Decimals:       newVault.Decimals,
				Version:        newVault.Version,
				Icon:           newVault.Icon,
				Type:           newVault.Type,
				FeaturingScore: newVault.FeaturingScore,
				Token: TRotkiVaultToken{
					Address:  newVault.Token.Address,
					Name:     newVault.Token.Name,
					Symbol:   newVault.Token.Symbol,
					Decimals: newVault.Token.Decimals,
					Icon:     newVault.Token.Icon,
				},
				Staking: TRotkiVaultStaking{
					Address:   newVault.Staking.Address,
					Available: newVault.Staking.Available,
				},
			}

			data = append(data, vaultForRotki)
		}
	}

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code sorts the 'data' slice based on the 'orderBy' and 'orderDirection'
	** parameters. The 'SortBy' function from the 'sort' package is used for this purpose.
	**
	** After sorting, it applies pagination to the 'data' slice based on the 'page' and 'limit'
	** parameters. The start index for the slice is calculated as (page - 1) * limit, and the end
	** index is calculated as page * limit. If the end index is greater than the length of the 'data'
	** slice, it is set to the length of the 'data' slice.
	**
	** The 'data' slice is then updated to only include the vaults between the start and end indices.
	** This effectively returns only the vaults for the requested page with the specified limit.
	**************************************************************************************************/
	sort.SortBy(orderBy, orderDirection, data)
	start := skip
	end := skip + limit
	if end > uint64(len(data)) {
		end = uint64(len(data))
	}
	data = data[start:end]

	return data
}

/**************************************************************************************************
** CountVaultsForRotki counts the number of vaults for Rotki integration
**
** This function takes a Gin context as input and returns an integer representing the count of
** vaults that meet the specified criteria. It performs the following steps:
**
** 1. Parses the 'chainIDs' query parameter to determine which chains to include
** 2. Iterates through the specified chains
** 3. For each chain, retrieves the list of vaults
** 4. Filters out blacklisted vaults and non-Yearn vaults
** 5. Calculates a featuring score for each valid vault
** 6. Increments the count for each valid vault
**
** The function uses various helper functions and environment variables to perform its operations.
** It's designed to work with the Rotki integration, providing a count of relevant vaults across
** specified blockchain networks.
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
	chainsStr := strings.Split(getQuery(c, `chainIDs`), `,`)
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
			_, err := NewVault().AssignTVault(currentVault)
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
