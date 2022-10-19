package models

import (
	"github.com/yearn/ydaemon/internal/utils/types/bigNumber"
	"github.com/yearn/ydaemon/internal/utils/types/common"
)

//TVaultFromGraphToken holds the info about a specific token or shareToken
type TVaultFromGraphToken struct {
	Id       common.Address
	Name     string
	Symbol   string
	Decimals uint64
}

//TVaultFromGraphStrategyReportsResult holds the results for a given report
type TVaultFromGraphStrategyReportsResults struct {
	APR        string `json:"apr,omitempty"`
	Duration   string `json:"duration,omitempty"`
	DurationPr string `json:"durationPr,omitempty"`
}

//TVaultFromGraphStrategyReports holds the reports for a given strategy
type TVaultFromGraphStrategyReports struct {
	Id      string                                  `json:"id,omitempty"`
	Results []TVaultFromGraphStrategyReportsResults `json:"results,omitempty"`
}

//TVaultFromGraphStrategy holds the info about a specific strategy for this vault
type TVaultFromGraphStrategy struct {
	Address       common.Address                   `json:"address"`
	Keeper        common.Address                   `json:"keeper,omitempty"`
	Strategist    common.Address                   `json:"strategist,omitempty"`
	Rewards       common.Address                   `json:"rewards,omitempty"`
	HealthCheck   string                           `json:"healthCheck"`
	Name          string                           `json:"name"`
	ApiVersion    string                           `json:"apiVersion,omitempty"`
	DoHealthCheck bool                             `json:"doHealthCheck,omitempty"`
	InQueue       bool                             `json:"inQueue"`
	EmergencyExit bool                             `json:"emergencyExit,omitempty"`
	Reports       []TVaultFromGraphStrategyReports `json:"reports,omitempty"`
}

//TVaultFromGraphVaultDayData holds the daily information about the vault status
type TVaultFromGraphVaultDayData struct {
	PricePerShare string
	Timestamp     string
}

//TVaultFromGraph holds the response data for a graphql request for the vaults
type TVaultFromGraph struct {
	Id                    common.Address                `json:"id"`
	Guardian              common.Address                `json:"guardian,omitempty"`
	Management            common.Address                `json:"management,omitempty"`
	Governance            common.Address                `json:"governance,omitempty"`
	Rewards               common.Address                `json:"rewards,omitempty"`
	BalanceTokens         *bigNumber.Int                `json:"balanceTokens"`
	AvailableDepositLimit *bigNumber.Int                `json:"availableDepositLimit,omitempty"`
	DepositLimit          *bigNumber.Int                `json:"depositLimit,omitempty"`
	BalanceTokensIdle     *bigNumber.Int                `json:"balanceTokensIdle,omitempty"`
	Activation            string                        `json:"activation"`
	Classification        string                        `json:"classification"`
	ApiVersion            string                        `json:"apiVersion"`
	ManagementFeeBps      uint64                        `json:"managementFeeBps"`
	PerformanceFeeBps     uint64                        `json:"performanceFeeBps"`
	ShareToken            TVaultFromGraphToken          `json:"shareToken"`
	Token                 TVaultFromGraphToken          `json:"token"`
	Strategies            []TVaultFromGraphStrategy     `json:"strategies"`
	VaultDayData          []TVaultFromGraphVaultDayData `json:"vaultDayData"`
	LatestUpdate          struct {
		Timestamp string `json:"timestamp"`
	} `json:"latestUpdate"`
}

//TMetaFromGraph holds some meta information, aka on the graph itself, for a graphql request
type TMetaFromGraph struct {
	HasIndexingErrors bool `json:"hasIndexingErrors,omitempty"`
	Block             struct {
		Number int64 `json:"number,omitempty"`
	} `json:"block,omitempty"`
}

//TReportsFromGraph holds the reports data for a graphql request
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

//TGraphQueryResponseForVaults is the response from the graphql query when we ask for the vaults
type TGraphQueryResponseForVaults struct {
	Vaults []TVaultFromGraph
}

//TGraphQueryResponseForVault is the response from the graphql query when we ask for one specific vault
type TGraphQueryResponseForVault struct {
	Vault TVaultFromGraph
}

//TGraphQueryResponseForWatchVaults is the response from the graphql query when we ask for the vaults for Yearn Watch
type TGraphQueryResponseForWatchVaults struct {
	Meta   TMetaFromGraph `json:"_meta,omitempty"`
	Vaults []TVaultFromGraph
}

//TGraphQLHarvestRequestForOneVault is the request for the graphql query when we ask for the harvests for one specific vault
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
		Timestamp string `json:"timestamp"`
		Profit    string `json:"profit"`
		Loss      string `json:"loss"`
	} `json:"harvests"`
}
