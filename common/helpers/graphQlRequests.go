package helpers

import "strconv"

/**************************************************************************************************
** This file contains pre-formatted GraphQL query strings used throughout the application to
** retrieve data from The Graph protocol endpoints. These queries target Yearn-specific data
** such as strategy reports, vault harvests, and user transaction history.
**
** Each function returns a standardized GraphQL query fragment that can be incorporated into
** larger queries or used directly with GraphQL endpoints. The fragments are formatted as
** multi-line strings for better readability.
**************************************************************************************************/

/**************************************************************************************************
** GetStrategyReports constructs a GraphQL query fragment to retrieve the most recent strategy
** reports for a vault. This query targets key performance metrics for strategies including
** debt levels, gains, losses, and calculated APR.
**
** The query is limited to the 10 most recent reports, ordered by timestamp in descending order
** to prioritize the most recent data.
**
** @return string A formatted GraphQL query fragment for strategy reports
**************************************************************************************************/
func GetStrategyReports(first int) string {
	return (`reports(first: ` + strconv.Itoa(first) + `, orderBy: timestamp, orderDirection: desc) {
		id
		totalDebt
		totalLoss
		totalGain
		debtLimit
		debtPaid
		debtAdded
		loss
		gain
		timestamp
		results {
			apr
			duration
			durationPr
		}
	}`)
}

/**************************************************************************************************
**************************************************************************************************/
func GetKongStrategyReports(first int) string {
	return (`lastReportDetail {
		apr {
			net
			gross
		}
		loss
		lossUsd
		profit
		profitUsd
		blockTime
      	blockNumber
	}`)
}

/**************************************************************************************************
** GetHarvestsForVaults constructs a GraphQL query fragment to retrieve harvest events for a
** vault. This query includes essential data needed to analyze harvest performance, including
** vault and strategy identification, transaction details, and profit/loss information.
**
** This query is designed to be incorporated into a larger GraphQL query with the appropriate
** entity selection and filtering.
**
** @return string A formatted GraphQL query fragment for vault harvests
**************************************************************************************************/
func GetHarvestsForVaults() string {
	return (`{
		vault {
			id
			token {
				id
				decimals
			}
		}
		strategy {
			id
		}
		transaction {
			hash
		}
		timestamp
		profit
		loss
	}`)
}

/**************************************************************************************************
** GetFIFOForUser constructs a GraphQL query fragment to retrieve a user's deposit and withdrawal
** history for a specific vault. This query captures all token movements, including shares minted,
** burnt, sent, and received, as well as the underlying tokens involved in these transactions.
**
** The query also includes the vault's latest price per share, which is essential for calculating
** current portfolio value and returns.
**
** @return string A formatted GraphQL query fragment for user transaction history
**************************************************************************************************/
func GetFIFOForUser() string {
	return (`
		updates {
			deposits
			withdrawals
			sharesBurnt
			sharesMinted
			sharesSent
			sharesReceived
			tokensSent
			tokensReceived
		}
		vault {
			id
			shareToken {
				decimals
			}
			latestUpdate {
				pricePerShare
			}
		}
	`)
}
