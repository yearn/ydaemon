package vaults

import (
	"strings"

	"github.com/gin-gonic/gin"
)

/**************************************************************************************************
** Controller is the main struct for handling vault-related API endpoints.
** It provides methods for retrieving various types of vaults based on different filters.
**************************************************************************************************/
type Controller struct{}

/**************************************************************************************************
** STRATEGIES_CONDITIONS contains the possible conditions to determine which strategies should
** be used in the return value.
** If the strategy is `absolute`, an active strategy will be isActive, inQueue and with a debt > 0
** If the strategy is `inQueue`, we will check if the vault has the strategy in it's withdrawal queue
** If the strategy is `debtLimit`, we will check if the vault has allocated a debtLimit to the strategy
** If the strategy is `debtRatio`, we will check if the vault has allocated a debtRatio to the strategy
** If the strategy is `all`, all strategies associated with the vault will be included
**************************************************************************************************/
var STRATEGIES_CONDITIONS = []string{`inQueue`, `debtLimit`, `debtRatio`, `absolute`, `all`}

/**************************************************************************************************
** selectStrategiesCondition determines the appropriate strategy condition to use based on the input.
**
** @param s string - The strategy condition string from the request
** @return string - The validated strategy condition or the default "debtRatio" if invalid
**************************************************************************************************/
func selectStrategiesCondition(s string) string {
	if s == `` {
		return `debtRatio`
	}
	for _, c := range STRATEGIES_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `debtRatio`
}

/**************************************************************************************************
** MIGRABLE_CONDITIONS contains the possible conditions to determine if migrable vaults should be
** included in the return value.
** If the condition is `none`, they will not be included if they also are hideAlways or retired
** If the condition is `all`, they will all be included
** If the condition is `nodust`, they will be included if they have a TVL > 100$
** If the condition is `ignore`, any vault with a migration available will be ignored
**************************************************************************************************/
var MIGRABLE_CONDITIONS = []string{`none`, `all`, `nodust`, `ignore`}

/**************************************************************************************************
** selectMigrableCondition determines the appropriate migrable condition to use based on the input.
**
** @param s string - The migrable condition string from the request
** @return string - The validated migrable condition or the default "none" if invalid
**************************************************************************************************/
func selectMigrableCondition(s string) string {
	if s == `` {
		return `none`
	}
	for _, c := range MIGRABLE_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `none`
}

/**************************************************************************************************
** getQuery retrieves a query parameter from a Gin context in a case-insensitive manner.
**
** @param c *gin.Context - The Gin context containing the request
** @param targetKey string - The name of the query parameter to retrieve
** @return string - The value of the query parameter or an empty string if not found
**************************************************************************************************/
func getQuery(c *gin.Context, targetKey string) string {
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if strings.EqualFold(targetKey, key) {
			return strings.Join(values, ",")
		}
	}
	return ""
}
