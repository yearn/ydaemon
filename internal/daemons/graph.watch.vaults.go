package daemons

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
)

func graphQLRequestForAllVaults() *graphql.Request {
	return graphql.NewRequest(`{
		_meta {
			hasIndexingErrors
			block {
				number
			}
		}
		vaults(first: 1000) {
			id
			guardian
			management
			governance
			rewards
			availableDepositLimit
			depositLimit
			balanceTokens
			balanceTokensIdle
			managementFeeBps
			performanceFeeBps
			apiVersion
			shareToken {
				decimals
				id
				name
				symbol
			}
			token {
				decimals
				id
				name
				symbol
			}
			strategies {
				address
				name
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
				}
				doHealthCheck
				healthCheck
			}
		}
	}`)
}

// FetchExtendedVaultsFromSubgraph fetches the strategies information from the Yearn Meta API for a given chainID
// and store the result to the global variable StrategiesFromMeta for later use.
func FetchExtendedVaultsFromSubgraph(chainID uint64) {
	// Prepare and execute the GraphQL request
	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphQLRequestForAllVaults()
	var response models.TGraphQueryResponseForWatchVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		return
	}

	for _, vaultFromGraph := range response.Vaults {
		logs.Pretty(vaultFromGraph.Id)
	}
}

// // RunWatch is a goroutine that periodically fetches the strategies information from the
// // Yearn Subgraph to work with Yearn Watch
// func RunWatch(chainID uint64, wg *sync.WaitGroup) {
// 	isDone := false
// 	for {
// 		FetchExtendedVaultsFromSubgraph(chainID)
// 		if !isDone {
// 			isDone = true
// 			wg.Done()
// 		}
// 		time.Sleep(24 * time.Hour)
// 	}
// }
