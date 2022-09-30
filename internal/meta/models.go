package meta

import (
	"github.com/yearn/ydaemon/internal/types/common"
	"github.com/yearn/ydaemon/internal/utils/models"
)

// TVaultFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/vaults/all
type TVaultFromMeta struct {
	Address              common.Address `json:"address"`
	MigrationTargetVault common.Address `json:"migrationTargetVault"`
	MigrationContract    common.Address `json:"migrationContract"`
	DisplayName          string         `json:"displayName"`
	Comment              string         `json:"comment"`
	APYTypeOverride      string         `json:"apyTypeOverride"`
	Order                float32        `json:"order"`
	HideAlways           bool           `json:"hideAlways"`
	DepositsDisabled     bool           `json:"depositsDisabled"`
	WithdrawalsDisabled  bool           `json:"withdrawalsDisabled"`
	MigrationAvailable   bool           `json:"migrationAvailable"`
	AllowZapIn           bool           `json:"allowZapIn"`
	AllowZapOut          bool           `json:"allowZapOut"`
	Retired              bool           `json:"retired"`
}

// TTokenFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/tokens/all
type TTokenFromMeta struct {
	Address      common.Address        `json:"address"`
	Name         string                `json:"name"`
	Symbol       string                `json:"symbol"`
	Description  string                `json:"description"`
	Website      string                `json:"website"`
	Categories   []string              `json:"categories"`
	Localization *models.TLocalization `json:"localization,omitempty"`
}

// TStrategyFromMeta is the structure of data we receive when calling meta.yearn.finance/api/1/strategies/all
type TStrategyFromMeta struct {
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Addresses    []common.Address      `json:"addresses"`
	Protocols    []string              `json:"protocols"`
	Localization *models.TLocalization `json:"localization,omitempty"`
}

// TProtocolsFromMeta is the structure of data for the protocols metadata stored in data/meta/protocols
type TProtocolsFromMeta struct {
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Localization *models.TLocalization `json:"localization,omitempty"`
}
