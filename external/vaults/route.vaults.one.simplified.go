package vaults

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

// GetSimplifiedVault will, for a given chainID, return a list of all vaults
func (y Controller) GetSimplifiedVault(c *gin.Context) {
	/** 🔵 - Yearn *************************************************************************************
	** chainID: The chain IDs for which the vaults are to be returned. It is obtained from the
	** 'chainIDs' query parameter in the request.
	**************************************************************************************************/
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** address: The address of the vault to be returned. It is obtained from the 'address' query
	** parameter in the request.
	**************************************************************************************************/
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	/** 🔵 - Yearn *************************************************************************************
	** strategiesCondition: A string that determines the condition for selecting strategies. It is
	** obtained from the 'strategiesCondition' query parameter in the request.
	**************************************************************************************************/
	strategiesCondition := selectStrategiesCondition(getQuery(c, `strategiesCondition`))

	/** 🔵 - Yearn *************************************************************************************
	** The following block of code will store the final vault to be returned in the response, which will
	** receive a bunch of mutation to be transformed to a simplified version of the vault.
	**************************************************************************************************/
	currentVault, ok := storage.GetVault(chainID, address)
	if !ok {
		c.String(http.StatusBadRequest, "invalid vault")
		return
	}
	newVault, err := NewVault().AssignTVault(currentVault)
	if err != nil {
		c.String(http.StatusBadRequest, "vault not found")
		return
	}
	APRAsFloat := 0.0
	if newVault.APR.NetAPR != nil {
		APRAsFloat, _ = newVault.APR.NetAPR.Float64()
	}
	newVault.FeaturingScore = newVault.TVL.TVL * APRAsFloat

	/** 🔵 - Yearn *************************************************************************************
	** For the vault, it performs the following operations:
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
	** 4. It appends the current strategies to the vault
	**
	** This effectively populates the 'Strategies' field for the vault with the appropriate strategies
	** and computes the risk score.
	**************************************************************************************************/
	vaultStrategies, _ := storage.ListStrategiesForVault(currentVault.ChainID, common.HexToAddress(newVault.Address))
	newVault.Strategies = []TStrategy{}
	for _, strategy := range vaultStrategies {
		var externalStrategy TStrategy
		strategyWithDetails := NewStrategy().AssignTStrategy(strategy)
		if !strategyWithDetails.ShouldBeIncluded(strategiesCondition) {
			continue
		}

		externalStrategy = strategyWithDetails
		newVault.Strategies = append(newVault.Strategies, externalStrategy)
	}

	vaultAsStrategy, ok := storage.GetStrategy(newVault.ChainID, common.HexToAddress(newVault.Address))
	if ok {
		simplified := toSimplifiedVersion(newVault, vaultAsStrategy)
		simplified.Description = newVault.Description
		if simplified.Description == "" {
			simplified.Description = vaultAsStrategy.Description
		}
		c.JSON(http.StatusOK, simplified)
		return
	}
	simplified := toSimplifiedVersion(newVault, models.TStrategy{})
	simplified.Description = newVault.Description
	c.JSON(http.StatusOK, simplified)
}
