package vaults

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/meta"
	"github.com/yearn/ydaemon/internal/strategies"
	"github.com/yearn/ydaemon/internal/tokens"
	"github.com/yearn/ydaemon/internal/utils"
)

// TTVL holds the info about the value locked in a vault
type TTVL struct {
	TotalAssets          *bigNumber.Int `json:"total_assets"`
	TotalDelegatedAssets *bigNumber.Int `json:"total_delegated_assets"`
	TVLDeposited         float64        `json:"tvl_deposited"`
	TVLDelegated         float64        `json:"tvl_delegated"`
	TVL                  float64        `json:"tvl"`
	Price                float64        `json:"price"`
}

// TAPYFees holds the fees information about this vault.
type TAPYFees struct {
	Performance float64 `json:"performance"`
	Withdrawal  float64 `json:"withdrawal"`
	Management  float64 `json:"management"`
	KeepCRV     float64 `json:"keep_crv"`
	CvxKeepCRV  float64 `json:"cvx_keep_crv"`
}

// TAPYPoints holds the points information about this vault.
type TAPYPoints struct {
	WeekAgo   float64 `json:"week_ago"`
	MonthAgo  float64 `json:"month_ago"`
	Inception float64 `json:"inception"`
}

// TAPYComposite holds the points information about this vault.
type TAPYComposite struct {
	Boost      float64 `json:"boost"`
	PoolAPY    float64 `json:"pool_apy"`
	BoostedAPR float64 `json:"boosted_apr"`
	BaseAPR    float64 `json:"base_apr"`
	CvxAPR     float64 `json:"cvx_apr"`
	RewardsAPR float64 `json:"rewards_apr"`
}

// TAPY contains all the information useful about the APY, APR, fees and breakdown.
type TAPY struct {
	Type              string        `json:"type"`
	GrossAPR          float64       `json:"gross_apr"`
	NetAPY            float64       `json:"net_apy"`
	StakingRewardsAPR float64       `json:"staking_rewards_apr"`
	Fees              TAPYFees      `json:"fees"`
	Points            TAPYPoints    `json:"points"`
	Composite         TAPYComposite `json:"composite"`
}

// TMigration helps us to know if a vault is in the process of being migrated.
type TMigration struct {
	Available bool           `json:"available"`
	Address   common.Address `json:"address"`
	Contract  common.Address `json:"contract"`
}

// TVault is the main structure returned by the API when trying to get all the vaults for a specific network
type TVault struct {
	Address               common.Address     `json:"address"`
	Management            common.Address     `json:"management"`
	Governance            common.Address     `json:"governance"`
	Guardian              common.Address     `json:"guardian"`
	Rewards               common.Address     `json:"rewards"`
	WithdrawalQueue       []common.Address   `json:"withdrawalQueue"` //Ignored from json
	PricePerShare         *bigNumber.Int     `json:"pricePerShare"`
	DepositLimit          *bigNumber.Int     `json:"depositLimit"`
	AvailableDepositLimit *bigNumber.Int     `json:"availableDepositLimit,omitempty"`
	TotalAssets           *bigNumber.Int     `json:"total_assets"`
	Type                  utils.TVaultType   `json:"type"`
	Symbol                string             `json:"symbol"`
	DisplaySymbol         string             `json:"display_symbol"`
	FormatedSymbol        string             `json:"formated_symbol"`
	Name                  string             `json:"name"`
	DisplayName           string             `json:"display_name"`
	FormatedName          string             `json:"formated_name"`
	Icon                  string             `json:"icon"`
	Version               string             `json:"version"`
	ChainID               uint64             `json:"chainID"`
	Inception             uint64             `json:"inception"`  //Timestamp
	Activation            uint64             `json:"activation"` //BlockNumber
	Decimals              uint64             `json:"decimals"`
	PerformanceFee        uint64             `json:"performanceFee"`
	ManagementFee         uint64             `json:"managementFee"`
	Endorsed              bool               `json:"endorsed"`
	EmergencyShutdown     bool               `json:"emergency_shutdown"`
	Token                 tokens.TERC20Token `json:"token"`
}

type TLegacyAPIAPY struct {
	Type              string  `json:"type"`
	GrossAPR          float64 `json:"gross_apr"`
	NetAPY            float64 `json:"net_apy"`
	StakingRewardsAPR float64 `json:"staking_rewards_apr"`
	Fees              struct {
		Performance float64 `json:"performance"`
		Withdrawal  float64 `json:"withdrawal"`
		Management  float64 `json:"management"`
		KeepCRV     float64 `json:"keep_crv"`
		CvxKeepCRV  float64 `json:"cvx_keep_crv"`
	} `json:"fees"`
	Points struct {
		WeekAgo   float64 `json:"week_ago"`
		MonthAgo  float64 `json:"month_ago"`
		Inception float64 `json:"inception"`
	} `json:"points"`
	Composite struct {
		Boost      float64 `json:"boost"`
		PoolAPY    float64 `json:"pool_apy"`
		BoostedAPR float64 `json:"boosted_apr"`
		BaseAPR    float64 `json:"base_apr"`
		CvxAPR     float64 `json:"cvx_apr"`
		RewardsAPR float64 `json:"rewards_apr"`
	} `json:"composite"`
}
type TLegacyAPI struct {
	Address common.MixedcaseAddress
	APY     TLegacyAPIAPY
}

type TAggregatedVault struct {
	Address       common.MixedcaseAddress
	LegacyAPY     TLegacyAPIAPY
	PricePerShare bigNumber.Int
}

func (t *TVault) BuildNames(metaVaultName string) {
	name := strings.Replace(t.Name, "\"", "", -1)
	displayName := t.Name
	formatedName := t.Token.Name

	// If the meta file has a display name, use it
	if metaVaultName != "" {
		displayName = metaVaultName
	}
	// If the formated name is missing yVault suffix, add it
	if !strings.HasSuffix(formatedName, "yVault") {
		formatedName = formatedName + " yVault"
	}
	// If a display name exist, use it for the formating.
	if displayName != "" && !strings.HasSuffix(displayName, "yVault") {
		formatedName = displayName + " yVault"
	}
	// If the name is empty, use the displayName instead
	if name == "" {
		name = displayName
	}
	// If the name is still empty, use the formated name instead
	if name == "" {
		name = formatedName
	}

	if displayName == "" && strings.HasSuffix(name, " Auto-Compounding yVault") {
		displayName = strings.Replace(name, " Auto-Compounding yVault", "", -1)
	}
	if strings.HasSuffix(displayName, " Auto-Compounding") {
		displayName = strings.Replace(displayName, " Auto-Compounding", "", -1)
	}

	t.Name = strings.Replace(name, "-f-f", "-f", -1)
	t.DisplayName = strings.Replace(displayName, "-f-f", "-f", -1)
	t.FormatedName = strings.Replace(formatedName, "-f-f", "-f", -1)
}

func (t *TVault) BuildSymbol(metaVaultSymbol string) {
	symbol := strings.Replace(t.Symbol, "\"", "", -1)
	formatedSymbol := t.Token.Symbol
	displaySymbol := metaVaultSymbol

	//If the formated symbol is missing yv prefix, add it
	if !strings.HasPrefix(formatedSymbol, "yv") {
		formatedSymbol = "yv" + formatedSymbol
	}
	// If a display name exist, use it for the formating.
	if displaySymbol != "" && !strings.HasPrefix(displaySymbol, "yv") {
		formatedSymbol = "yv" + displaySymbol
	}
	symbol = helpers.SafeString(symbol, displaySymbol)
	symbol = helpers.SafeString(symbol, formatedSymbol)
	displaySymbol = helpers.SafeString(displaySymbol, symbol)

	t.Symbol = strings.Replace(symbol, "-f-f", "-f", -1)
	t.DisplaySymbol = strings.Replace(displaySymbol, "-f-f", "-f", -1)
	t.FormatedSymbol = strings.Replace(formatedSymbol, "-f-f", "-f", -1)
}

func (t *TVault) BuildMigration() TMigration {
	migration := TMigration{
		Available: false,
		Address:   t.Address,
	}

	if vaultFromMeta, ok := meta.GetMetaVault(t.ChainID, t.Address); ok {
		migrationAddress := t.Address
		migrationContract := common.Address{}
		migrationAvailable := vaultFromMeta.MigrationAvailable
		if vaultFromMeta.MigrationAvailable {
			migrationAddress = vaultFromMeta.MigrationTargetVault
			migrationContract = vaultFromMeta.MigrationContract
		}
		migration = TMigration{
			Available: migrationAvailable,
			Address:   migrationAddress,
			Contract:  migrationContract,
		}
	}
	return migration
}

func (t *TVault) BuildAPY(aggregatedVault *TAggregatedVault, hasLegacyAPY bool) TAPY {
	apy := TAPY{}
	vaultFromMeta, okMeta := meta.GetMetaVault(t.ChainID, t.Address)

	if hasLegacyAPY {
		apy = TAPY{
			Type:              aggregatedVault.LegacyAPY.Type,
			GrossAPR:          aggregatedVault.LegacyAPY.GrossAPR,
			NetAPY:            aggregatedVault.LegacyAPY.NetAPY,
			StakingRewardsAPR: aggregatedVault.LegacyAPY.StakingRewardsAPR,
			Points: TAPYPoints{
				WeekAgo:   aggregatedVault.LegacyAPY.Points.WeekAgo,
				MonthAgo:  aggregatedVault.LegacyAPY.Points.MonthAgo,
				Inception: aggregatedVault.LegacyAPY.Points.Inception,
			},
			Composite: TAPYComposite{
				Boost:      aggregatedVault.LegacyAPY.Composite.Boost,
				PoolAPY:    aggregatedVault.LegacyAPY.Composite.PoolAPY,
				BoostedAPR: aggregatedVault.LegacyAPY.Composite.BoostedAPR,
				BaseAPR:    aggregatedVault.LegacyAPY.Composite.BaseAPR,
				CvxAPR:     aggregatedVault.LegacyAPY.Composite.CvxAPR,
				RewardsAPR: aggregatedVault.LegacyAPY.Composite.RewardsAPR,
			},
			Fees: TAPYFees{
				Performance: aggregatedVault.LegacyAPY.Fees.Performance,
				Management:  aggregatedVault.LegacyAPY.Fees.Management,
				Withdrawal:  aggregatedVault.LegacyAPY.Fees.Withdrawal,
				KeepCRV:     aggregatedVault.LegacyAPY.Fees.KeepCRV,
				CvxKeepCRV:  aggregatedVault.LegacyAPY.Fees.CvxKeepCRV,
			},
		}
		if okMeta && vaultFromMeta.APYTypeOverride != `` {
			apy.Type = vaultFromMeta.APYTypeOverride
		}
	} else if okMeta && vaultFromMeta.APYTypeOverride != `` {
		logs.Error(`Missing APY vault data for chainID: ` + strconv.FormatUint(t.ChainID, 10) + ` and address: ` + t.Address.Hex())
		apy.Type = vaultFromMeta.APYTypeOverride
	}
	return apy
}

func (t *TVault) BuildTVL() TTVL {
	humanizedPrice, fHumanizedPrice := getHumanizedTokenPrice(t.ChainID, t.Token.Address)
	fHumanizedTVLPrice := getHumanizedValue(t.TotalAssets, int(t.Decimals), humanizedPrice)
	strategiesList := strategies.ListStrategiesForVault(t.ChainID, t.Address)
	delegatedTokenAsBN := bigNumber.NewInt(0)
	fDelegatedValue := 0.0

	for _, strat := range strategiesList {
		delegatedValue := getHumanizedValue(strat.DelegatedAssets, int(t.Decimals), humanizedPrice)
		delegatedTokenAsBN = delegatedTokenAsBN.Add(delegatedTokenAsBN, strat.DelegatedAssets)
		fDelegatedValue += delegatedValue
	}

	tvl := TTVL{
		TotalAssets:          t.TotalAssets,
		TotalDelegatedAssets: delegatedTokenAsBN,
		TVL:                  fHumanizedTVLPrice - fDelegatedValue,
		TVLDeposited:         fHumanizedTVLPrice,
		TVLDelegated:         fDelegatedValue,
		Price:                fHumanizedPrice,
	}
	return tvl
}

func (t *TVault) BuildCategory() string {
	category := `Volatile`
	baseForStableCurrencies := []string{`USD`, `EUR`, `AUD`, `CHF`, `KRW`, `GBP`, `JPY`}
	baseForCurve := []string{`curve`, `crv`}
	baseForBalancer := []string{`balancer`, `bal`}
	baseForVelodrom := []string{`velodrome`, `velo`}
	allNames := []string{
		strings.ToLower(t.FormatedName),
		strings.ToLower(t.Name),
		strings.ToLower(t.DisplayName),
	}

	if vaultFromMeta, ok := meta.GetMetaVault(t.ChainID, t.Address); ok {
		//Using meta classification to set the category
		if vaultFromMeta.Classification.Stability == `Volatile` {
			category = `Volatile`
		} else {
			if helpers.Contains(baseForStableCurrencies, vaultFromMeta.Classification.StableBaseAsset) {
				category = `Stablecoin`
			} else {
				category = `Volatile`
			}
		}
		if helpers.Intersects(allNames, baseForCurve) {
			category = `Curve`
		}
		if helpers.Intersects(allNames, baseForBalancer) {
			category = `Balancer`
		}
		if helpers.Intersects(allNames, baseForVelodrom) {
			category = `Velodrome`
		}
	} else {
		//No meta, back to custom classification
		baseForBitcoin := []string{`btc`, `bitcoin`}
		baseForEth := []string{`eth`, `ethereum`}
		baseForStableCoins := []string{`dai`, `rai`, `mim`, `dola`}

		for _, stable := range baseForStableCurrencies {
			baseForStableCoins = append(baseForStableCoins, strings.ToLower(stable))
		}

		if helpers.Intersects(allNames, baseForBitcoin) {
			category = `Volatile`
		}
		if helpers.Intersects(allNames, baseForEth) {
			category = `Volatile`
		}
		if helpers.Intersects(allNames, baseForStableCoins) {
			category = `Stablecoin`
		}
		if helpers.Intersects(allNames, baseForCurve) {
			category = `Curve`
		}
		if helpers.Intersects(allNames, baseForBalancer) {
			category = `Balancer`
		}
		if helpers.Intersects(allNames, baseForVelodrom) {
			category = `Velodrome`
		}
	}
	return category
}

/**********************************************************************************************
** Set of functions to store and retrieve the tokens from the cache and/or database and being
** able to access them from the rest of the application.
** The _vaultMap variable is not exported and is only used internally by the functions below.
**********************************************************************************************/
var _vaultMap = make(map[uint64]map[common.Address]*TVault)

/**********************************************************************************************
** GetVault will, for a given chainID, try to retrieve the vault from the _vaultMap variable.
** It will return the vault if found, and a boolean indicating if the vault was found or not.
**********************************************************************************************/
func GetVault(chainID uint64, vaultAddress common.Address) (*TVault, bool) {
	if vaultsForChain, ok := _vaultMap[chainID]; ok {
		if vault, ok := vaultsForChain[vaultAddress]; ok {
			return vault, true
		}
	}
	return nil, false
}

/**********************************************************************************************
** ListVaults will, for a given chainID, return the list of all the vaults stored in _vaultMap.
**********************************************************************************************/
func ListVaults(chainID uint64) []*TVault {
	var vaults []*TVault
	for _, vault := range _vaultMap[chainID] {
		vaults = append(vaults, vault)
	}
	return vaults
}

/**********************************************************************************************
** MapVaults will, for a given chainID, return the map of all the vaults stored in _vaultMap.
**********************************************************************************************/
func MapVaults(chainID uint64) map[common.Address]*TVault {
	var vaults map[common.Address]*TVault
	if vaultsForChain, ok := _vaultMap[chainID]; ok {
		vaults = vaultsForChain
	}

	return vaults
}

/**********************************************************************************************
** ListVaultsAddresses will, for a given chainID, return the list of addresses of all the
** vaults stored in _vaultMap.
**********************************************************************************************/
func ListVaultsAddresses(chainID uint64) []common.Address {
	var addresses []common.Address
	for address := range _vaultMap[chainID] {
		addresses = append(addresses, address)
	}
	return addresses
}

/**********************************************************************************************
** FindVault will, for a given chainID, try to find the provided vaultAddress stored in
** _vaultMap. It will return the vault if found, and a boolean indicating if the vault was
** found or not.
**********************************************************************************************/
func FindVault(chainID uint64, vaultAddress common.Address) (*TVault, bool) {
	token, ok := _vaultMap[chainID][vaultAddress]
	if !ok {
		return nil, false
	}
	return token, true
}
