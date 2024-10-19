package helpers

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

// GetFIFOForUser construct the basic graphQL request to get all the deposits and withdrawals for a user
// for a given vault
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
