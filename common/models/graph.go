package models

import (
	"github.com/yearn/ydaemon/common/bigNumber"
)

// TReportsFromGraph holds the reports data for a graphql request
type TReportsFromGraph struct {
	Strategy struct {
		Reports []struct {
			ID        string         `json:"id"`
			DebtAdded *bigNumber.Int `json:"debtAdded"`
			DebtLimit *bigNumber.Int `json:"debtLimit"`
			TotalDebt *bigNumber.Int `json:"totalDebt"`
			Gain      *bigNumber.Int `json:"gain"`
			TotalGain *bigNumber.Int `json:"totalGain"`
			Loss      *bigNumber.Int `json:"loss"`
			TotalLoss *bigNumber.Int `json:"totalLoss"`
			DebtPaid  *bigNumber.Int `json:"debtPaid"`
			Timestamp string         `json:"timestamp"`
			Results   []struct {
				Duration   string `json:"duration"`
				DurationPr string `json:"durationPR"`
				Apr        string `json:"APR"`
			} `json:"results"`
		} `json:"reports"`
	} `json:"strategy"`
}

// TGraphQLHarvestRequestForOneVault is the request for the graphql query when we ask for the harvests for one specific vault
type TGraphQLHarvestRequestForOneVault struct {
	Harvests []struct {
		Transaction struct {
			Hash string `json:"hash"`
		}
		Vault struct {
			Id    string `json:"id"`
			Token struct {
				Id       string `json:"id"`
				Decimals int    `json:"decimals"`
			}
		}
		Strategy struct {
			Id string `json:"id"`
		}
		Timestamp string `json:"timestamp"`
		Profit    string `json:"profit"`
		Loss      string `json:"loss"`
	} `json:"harvests"`
}

// TFIFOForUserForVault is the request for the graphql query when we ask for the earnings for one specific vault
type TFIFOForUserForVault struct {
	AccountVaultPositions []struct {
		Updates []struct {
			Deposits       string
			Withdrawals    string
			SharesBurnt    string
			SharesMinted   string
			SharesSent     string
			SharesReceived string
			TokensSent     string
			TokensReceived string
		} `json:"updates"`
		Vault struct {
			Id         string `json:"id"`
			ShareToken struct {
				Decimals int64 `json:"decimals"`
			} `json:"shareToken"`
			LatestUpdate struct {
				PricePerShare string `json:"pricePerShare"`
			} `json:"latestUpdate"`
		} `json:"vault"`
	}
}
