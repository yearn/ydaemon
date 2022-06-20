package models

// TVaultFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/vaults/all
type TVaultFromMeta struct {
	Comment              string
	HideAlways           bool
	DepositsDisabled     bool
	WithdrawalsDisabled  bool
	Order                float32
	MigrationAvailable   bool
	AllowZapIn           bool
	AllowZapOut          bool
	Retired              bool
	DisplayName          string
	MigrationTargetVault string
	MigrationContract    string
	Address              string
}

// TTokenFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/tokens/all
type TTokenFromMeta struct {
	Address     string
	Name        string
	Symbol      string
	Description string
	Website     string
	Categories  []string
}

// TStrategyFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/strategies/all
type TStrategyFromMeta struct {
	Name        string
	Description string
	Addresses   []string
	Protocols   []string
}
