package partnerFees

import (
	"github.com/yearn/ydaemon/common/bigNumber"
)

type TMainCSVExport struct {
	Block               uint64 `csv:"block"`
	Timestamp           uint64 `csv:"timestamp"`
	Partner             string `csv:"partner"`
	Vault               string `csv:"vault"`
	Tier                string `csv:"tier"`
	VaultPrice          string `csv:"vault_price"`
	Balance             string `csv:"delegated_balance"`
	BalanceUSD          string `csv:"delegated_balance_usd"`
	TotalSupply         string `csv:"total_supply"`
	TotalSupplyUSD      string `csv:"total_supply_usd"`
	TotalDelegatedValue string `csv:"total_delegated_value"`
	CollectedFee        string `csv:"collected_fee"`
	CollectedFeeUSD     string `csv:"collected_fee_usd"`
	PartnerFeeShare     string `csv:"partner_fee"`
	PartnerFeeShareUSD  string `csv:"partner_fee_usd"`
}

type TPartnerBreakdownCSVExport struct {
	Block      uint64 `csv:"block"`
	Partner    string `csv:"partner"`
	Vault      string `csv:"vault"`
	Depositor  string `csv:"depositor"`
	Balance    string `csv:"balance"`
	BalanceUSD string `csv:"balance_usd"`
}

type TBig struct {
	Raw        *bigNumber.Int
	Normalized *bigNumber.Float
	Value      *bigNumber.Float
}
