package models

//TVaultFromGraphToken holds the info about a specific token or shareToken
type TVaultFromGraphToken struct {
	Id       string
	Name     string
	Symbol   string
	Decimals uint64
}

//TVaultFromGraphStrategy holds the info about a specific strategy for this vault
type TVaultFromGraphStrategy struct {
	Address   string
	Name      string
	DebtLimit string
	InQueue   bool
}

//TVaultFromGraphVaultDayData holds the daily information about the vault status
type TVaultFromGraphVaultDayData struct {
	PricePerShare string
	Timestamp     string
}

//TVaultFromGraph holds the response data for a graphql request for the vaults
type TVaultFromGraph struct {
	Id                string
	Activation        string
	Classification    string
	ApiVersion        string
	BalanceTokens     string
	ManagementFeeBps  uint64
	PerformanceFeeBps uint64
	ShareToken        TVaultFromGraphToken
	Token             TVaultFromGraphToken
	Strategies        []TVaultFromGraphStrategy
	VaultDayData      []TVaultFromGraphVaultDayData
	LatestUpdate      struct {
		Timestamp string
	}
}

//TGraphQueryResponseForVaults is the response from the graphql query when we ask for the vaults
type TGraphQueryResponseForVaults struct {
	Vaults []TVaultFromGraph
}

//TGraphQueryResponseForVault is the response from the graphql query when we ask for one specific vault
type TGraphQueryResponseForVault struct {
	Vault TVaultFromGraph
}
