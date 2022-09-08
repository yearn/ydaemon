package meta

// TVaultFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/vaults/all
type TVaultFromMeta struct {
	DisplayName          string  `json:"displayName"`
	MigrationTargetVault string  `json:"migrationTargetVault"`
	MigrationContract    string  `json:"migrationContract"`
	Address              string  `json:"address"`
	Comment              string  `json:"comment"`
	APYTypeOverride      string  `json:"apyTypeOverride"`
	Order                float32 `json:"order"`
	HideAlways           bool    `json:"hideAlways"`
	DepositsDisabled     bool    `json:"depositsDisabled"`
	WithdrawalsDisabled  bool    `json:"withdrawalsDisabled"`
	MigrationAvailable   bool    `json:"migrationAvailable"`
	AllowZapIn           bool    `json:"allowZapIn"`
	AllowZapOut          bool    `json:"allowZapOut"`
	Retired              bool    `json:"retired"`
}

// TTokenFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/tokens/all
type TTokenFromMeta struct {
	Address      string         `json:"address"`
	Name         string         `json:"name"`
	Symbol       string         `json:"symbol"`
	Description  string         `json:"description"`
	Website      string         `json:"website"`
	Categories   []string       `json:"categories"`
	Localization *TLocalization `json:"localization,omitempty"`
}

// TStrategyFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/strategies/all
type TStrategyFromMeta struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Addresses    []string       `json:"addresses"`
	Protocols    []string       `json:"protocols"`
	Localization *TLocalization `json:"localization,omitempty"`
}

// TProtocolsFromMeta is the structure of data for the protocols metadata stored in data/meta/protocols
type TProtocolsFromMeta struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Localization *TLocalization `json:"localization,omitempty"`
}

// TLocalization contains the all the enabled languages for the protocol
type TLocalization struct {
	En TLocalizationDetails `json:"en"`
	Fr TLocalizationDetails `json:"fr"`
	Es TLocalizationDetails `json:"es"`
	De TLocalizationDetails `json:"de"`
	Pt TLocalizationDetails `json:"pt"`
	El TLocalizationDetails `json:"el"`
	Tr TLocalizationDetails `json:"tr"`
	Vi TLocalizationDetails `json:"vi"`
	Zh TLocalizationDetails `json:"zh"`
	Hi TLocalizationDetails `json:"hi"`
	Ja TLocalizationDetails `json:"ja"`
	Id TLocalizationDetails `json:"id"`
	Ru TLocalizationDetails `json:"ru"`
}

// TLocalizationDetails contains the localization details for a specific language
type TLocalizationDetails struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}
