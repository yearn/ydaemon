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

type TReportsFromKong struct {
	Data struct {
		Strategy struct {
			LastReportDetail struct {
				Apr struct {
					Net   float64 `json:"net"`
					Gross float64 `json:"gross"`
				} `json:"apr"`
				Loss        string  `json:"loss"`
				LossUsd     float64 `json:"lossUsd"`
				Profit      string  `json:"profit"`
				ProfitUsd   float64 `json:"profitUsd"`
				BlockTime   string  `json:"blockTime"`
				BlockNumber string  `json:"blockNumber"`
			} `json:"lastReportDetail"`
			Address string `json:"address"`
		} `json:"strategy"`
	} `json:"data"`
}

type TReportResult struct {
	Duration   uint64  `json:"duration"`
	DurationPR uint64  `json:"durationPR"`
	APR        float64 `json:"APR"`
}
type TReport struct {
	ID        string          `json:"id"`
	DebtAdded *bigNumber.Int  `json:"debtAdded"`
	DebtLimit *bigNumber.Int  `json:"debtLimit"`
	TotalDebt *bigNumber.Int  `json:"totalDebt"`
	Gain      *bigNumber.Int  `json:"gain"`
	TotalGain *bigNumber.Int  `json:"totalGain"`
	Loss      *bigNumber.Int  `json:"loss"`
	TotalLoss *bigNumber.Int  `json:"totalLoss"`
	DebtPaid  *bigNumber.Int  `json:"debtPaid"`
	Timestamp uint64          `json:"timestamp"`
	Results   []TReportResult `json:"results"`
}
type TReportsFromGraphToClient struct {
	Strategy struct {
		Reports []TReport `json:"reports"`
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
