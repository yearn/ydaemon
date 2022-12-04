package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGraphRequestVault(t *testing.T) {
	assert.Equal(t, GetGraphRequestVault(), `
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

func TestGetGraphRequestStrategies(t *testing.T) {
	assert.Equal(t, GetGraphRequestStrategies(1, true), `
		strategies(first: 1) {
			address
			name
			inQueue
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
			healthCheck
		}`)
}

func TestGetStrategyReports(t *testing.T) {
	assert.Equal(t, GetStrategyReports(), `reports(first: 10, orderBy: timestamp, orderDirection: desc) {
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
