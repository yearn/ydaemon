package vaults

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/vaults"
)

// GetAllVaultsForAllChains will return a list of all vaults for all chains
func (y Controller) GetAllVaultsForAllChains(c *gin.Context) {
	/** 🔵 - Yearn *************************************************************************************
	** orderBy: A string that determines the order in which the vaults are returned. It is obtained
	** from the 'orderBy' query parameter in the request. If the parameter is not provided,
	** it defaults to 'details.order'.
	**
	** orderDirection: A string that determines the direction of the ordering of the vaults. It is
	** obtained from the 'orderDirection' query parameter in the request. If the parameter is not
	** provided, it defaults to 'asc'.
	**************************************************************************************************/
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `FeaturingScore`)
	orderDirection := helpers.SafeString(getQuery(c, `orderDirection`), `asc`)

	/** 🔵 - Yearn *************************************************************************************
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**
	** withStrategiesDetails: A boolean value that determines whether to include details of the
	** strategies in the response. It is obtained from the 'strategiesDetails' query parameter in
	** the request. If the parameter is 'withDetails', this value is true.
	**************************************************************************************************/
	strategiesCondition := selectStrategiesCondition(getQuery(c, `strategiesCondition`))
	withStrategiesDetails := strings.EqualFold(getQuery(c, `strategiesDetails`), `withDetails`)

	/** 🔵 - Yearn *************************************************************************************
	** hideAlways: A boolean value that determines whether to hide certain vaults. It is obtained
	** from the 'hideAlways' query parameter in the request. If the parameter is not provided,
	** it defaults to 'false'.
	**
	** migrable: A string that determines the condition for selecting migrable vaults. It is
	** obtained from the 'migrable' query parameter in the request.
	**************************************************************************************************/
	hideAlways := helpers.StringToBool(helpers.SafeString(getQuery(c, `hideAlways`), `false`))
	migrable := selectMigrableCondition(getQuery(c, `migrable`))
	if migrable != `none` && hideAlways {
		c.String(http.StatusBadRequest, `migrable and hideAlways cannot be true at the same time`)
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** page: A uint64 value that represents the page number for pagination. It is obtained from the
	** 'page' query parameter in the request. If the parameter is not provided, it defaults to 1.
	**
	** limit: A uint64 value that represents the number of vaults to be returned per page. It is
	** obtained from the 'limit' query parameter in the request. If the parameter is not provided,
	** it defaults to 50.
	**************************************************************************************************/
	page := helpers.SafeStringToUint64(getQuery(c, `page`), 1)
	limit := helpers.SafeStringToUint64(getQuery(c, `limit`), 50)

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
	if len(chains) == 0 {
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
	data := []TExternalVault{}
	allVaults := []*TExternalVault{}
	for _, chainID := range chains {
		vaultsForChain := vaults.ListVaults(chainID)
		for _, currentVault := range vaultsForChain {
			if helpers.Contains(env.BLACKLISTED_VAULTS[chainID], currentVault.Address) {
				continue
			}
			newVault := NewVault().AssignTVault(currentVault)
			if migrable == `none` && (newVault.Details.HideAlways || newVault.Details.Retired) && hideAlways {
				continue
			} else if migrable == `nodust` && (newVault.TVL.TVL < 100 || !newVault.Migration.Available) {
				continue
			} else if migrable == `all` && !newVault.Migration.Available {
				continue
			} else if migrable == `ignore` && (newVault.Migration.Available || newVault.Details.HideAlways || newVault.Details.Retired) {
				continue
			}
			// featuringScore =  ((v.tvl.tvl || 0) * (v?.apy?.net_apy || 0))
			newVault.FeaturingScore = newVault.TVL.TVL * newVault.APY.NetAPY
			allVaults = append(allVaults, newVault)
		}
	}

	for _, currentVault := range allVaults {
		vaultStrategies := strategies.ListStrategiesForVault(currentVault.ChainID, common.HexToAddress(currentVault.Address))
		currentVault.Strategies = []*TStrategy{}
		for _, strategy := range vaultStrategies {
			var externalStrategy *TStrategy
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			if withStrategiesDetails {
				externalStrategy = strategyWithDetails
				externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(strategies.BuildRiskScore(strategy))
			} else {
				externalStrategy = &TStrategy{
					Address:     strategy.Address.Hex(),
					Name:        strategy.Name,
					DisplayName: strategy.DisplayName,
					Description: strategy.Description,
				}
			}
			currentVault.Strategies = append(currentVault.Strategies, externalStrategy)
		}
		if withStrategiesDetails {
			currentVault.RiskScore = currentVault.ComputeRiskScore()
		}

		data = append(data, *currentVault)
	}

	//Sort by details.order by default
	sort.SortBy(orderBy, orderDirection, data)
	c.JSON(http.StatusOK, data)
}
