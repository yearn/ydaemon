package models

//TVaultFromGraphToken holds the info about a specific token or shareToken
type TVaultFromGraphToken struct {
	Id       string
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
	Address       string                           `json:"address"`
	Name          string                           `json:"name"`
	ApiVersion    string                           `json:"apiVersion,omitempty"`
	Keeper        string                           `json:"keeper,omitempty"`
	Strategist    string                           `json:"strategist,omitempty"`
	Rewards       string                           `json:"rewards,omitempty"`
	HealthCheck   string                           `json:"healthCheck"`
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
	Id                    string                        `json:"id"`
	Activation            string                        `json:"activation"`
	Classification        string                        `json:"classification"`
	ApiVersion            string                        `json:"apiVersion"`
	BalanceTokens         string                        `json:"balanceTokens"`
	Guardian              string                        `json:"guardian,omitempty"`
	Management            string                        `json:"management,omitempty"`
	Governance            string                        `json:"governance,omitempty"`
	Rewards               string                        `json:"rewards,omitempty"`
	AvailableDepositLimit string                        `json:"availableDepositLimit,omitempty"`
	DepositLimit          string                        `json:"depositLimit,omitempty"`
	BalanceTokensIdle     string                        `json:"balanceTokensIdle,omitempty"`
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
			ID        string `json:"id"`
			Timestamp string `json:"timestamp"`
			DebtAdded string `json:"debtAdded"`
			DebtLimit string `json:"debtLimit"`
			TotalDebt string `json:"totalDebt"`
			Gain      string `json:"gain"`
			TotalGain string `json:"totalGain"`
			Loss      string `json:"loss"`
			TotalLoss string `json:"totalLoss"`
			DebtPaid  string `json:"debtPaid"`
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
