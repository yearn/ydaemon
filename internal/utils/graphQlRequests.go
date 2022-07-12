package utils

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
				totalDebt
				totalLoss
				totalGain
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
			}
			doHealthCheck
			healthCheck
		`
	}
	return (`
		strategies(first: ` + numStr + `) {
			address
			name
			inQueue
			debtLimit
			` + details + `
		}
	`)
}
