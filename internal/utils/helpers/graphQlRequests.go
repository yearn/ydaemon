package helpers

import (
	"strconv"
)

// GetGraphRequestVault construct the basic graphQL request for the vault part
func GetGraphRequestVault() string {
	return (`
		id
		activation
		apiVersion
		classification
		managementFeeBps
		performanceFeeBps
		balanceTokens
		management
		governance
		guardian
		rewards
		depositLimit
		latestUpdate {
			timestamp
		}
		shareToken {
			name
			symbol
			id
			decimals
		}
		token {
			name
			symbol
			id
			decimals
		}
	`)
}

// GetGraphRequestStrategies construct the basic graphQL request for the strategies part.
// Based on withDetails or not, this can change
func GetGraphRequestStrategies(num int64, withDetails bool) string {
	numStr := strconv.FormatInt(int64(num), 10)
	details := ``
	if withDetails {
		details = `
			apiVersion
			emergencyExit
			keeper
			strategist
			rewards
			reports(first: 10, orderBy: timestamp, orderDirection: desc) {
				id
				results {
					apr
					duration
					durationPr
				}
			}
			doHealthCheck
			healthCheck`
	}
	return (`
		strategies(first: ` + numStr + `) {
			address
			name
			inQueue` + details + `
		}`)
}

// GetStrategyReports construct the basic graphQL request to get the reports
func GetStrategyReports() string {
	return (`reports(first: 10, orderBy: timestamp, orderDirection: desc) {
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

// GetHarvestsForVaults construct the basic graphQL request to get the harvests for a vault with
// all the minimal elements to format the data
func GetHarvestsForVaults() string {
	return (`{
		vault {
			id
			token {
				id
				decimals
			}
		}
		transaction {
			hash
		}
		timestamp
		profit
		loss
	}`)
}
