package vaults

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/sort"
	"github.com/yearn/ydaemon/internal/risk"
	"github.com/yearn/ydaemon/internal/storage"
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
	orderBy := helpers.SafeString(getQuery(c, `orderBy`), `featuringScore`)
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
	data := []TExternalVault{}
	allVaults := []TExternalVault{}
	for _, chainID := range chains {
		vaultsForChain, _ := storage.ListVaults(chainID)
		for _, currentVault := range vaultsForChain {
			if helpers.Contains(env.CHAINS[chainID].BlacklistedVaults, currentVault.Address) {
				continue
			}
			newVault, err := NewVault().AssignTVault(currentVault)
			if err != nil {
				continue
			}
			if migrable == `none` && (newVault.Details.IsHidden || newVault.Details.IsRetired) && hideAlways {
				continue
			} else if migrable == `nodust` && (newVault.TVL.TVL < 100 || !newVault.Migration.Available) {
				continue
			} else if migrable == `all` && !newVault.Migration.Available {
				continue
			} else if migrable == `ignore` && (newVault.Migration.Available || newVault.Details.IsHidden || newVault.Details.IsRetired) {
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
			allVaults = append(allVaults, newVault)
		}
	}

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code is a loop that iterates over each vault in the 'allVaults' slice.
	** For each vault, it performs the following operations:
	**
	** 1. It calls the 'ListStrategiesForVault' function to get a list of all strategies for the
	**    current vault. The function takes two parameters: the chain ID of the vault and the address
	**    of the vault.
	**
	** 2. It initializes an empty slice of pointers to 'TStrategy' objects for the current vault.
	**
	** 3. It then enters a nested loop that iterates over each strategy in 'vaultStrategies'. For each
	**    strategy, it performs the following operations:
	**
	**    a. It initializes a pointer to a 'TStrategy' object.
	**
	**    b. It calls the 'NewStrategy' function to create a new 'TStrategy' object and assigns the
	**       current strategy to it using the 'AssignTStrategy' method.
	**
	**    c. It checks if the strategy should be included based on the 'strategiesCondition'. If not,
	**       it skips the current iteration of the loop.
	**
	**    d. If 'withStrategiesDetails' is true, it assigns the strategy with details to
	**       'externalStrategy' and assigns a risk score to it. Otherwise, it assigns a new 'TStrategy'
	**       object with the address, name, display name, and description of the current strategy to
	**       'externalStrategy'.
	**
	**    e. It appends 'externalStrategy' to the 'Strategies' field of 'currentVault'.
	**
	** 4. If 'withStrategiesDetails' is true, it computes the risk score for the current vault and
	**    assigns it to the 'RiskScore' field of 'currentVault'.
	**
	** 5. It appends the current vault to the 'data' slice.
	**
	** This loop effectively populates the 'Strategies' field of each vault in 'allVaults' with the
	** appropriate strategies and computes the risk score for each vault if 'withStrategiesDetails' is
	** true. It also populates the 'data' slice with the final list of vaults to be returned in the
	** response.
	**************************************************************************************************/
	for _, currentVault := range allVaults {
		vaultStrategies, _ := storage.ListStrategiesForVault(currentVault.ChainID, common.HexToAddress(currentVault.Address))
		currentVault.Strategies = []TStrategy{}
		for _, strategy := range vaultStrategies {
			var externalStrategy TStrategy
			strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
			if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
				continue
			}

			if withStrategiesDetails {
				externalStrategy = strategyWithDetails
				externalStrategy.Risk = NewRiskScore().AssignTStrategyFromRisk(risk.BuildRiskScore(strategy))
			} else {
				externalStrategy = strategyWithDetails
			}
			currentVault.Strategies = append(currentVault.Strategies, externalStrategy)
		}
		if withStrategiesDetails {
			currentVault.RiskScore = currentVault.ComputeRiskScore()
		}

		data = append(data, currentVault)
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
	start := (page - 1) * limit
	end := page * limit
	if end > uint64(len(data)) {
		end = uint64(len(data))
	}
	data = data[start:end]

	c.JSON(http.StatusOK, data)
}
