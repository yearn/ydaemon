package kong

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

const (
	DefaultTimeout = 30 * time.Second
)

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

type KongAsset struct {
	ChainID  int    `json:"chainId"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

type KongDebt struct {
	Strategy           string   `json:"strategy"`           // String
	PerformanceFee     *string  `json:"performanceFee"`     // BigInt (string) or null
	Activation         *string  `json:"activation"`         // BigInt (string) or null
	DebtRatio          *string  `json:"debtRatio"`          // BigInt (string) or null
	MinDebtPerHarvest  *string  `json:"minDebtPerHarvest"`  // BigInt (string) or null
	MaxDebtPerHarvest  *string  `json:"maxDebtPerHarvest"`  // BigInt (string) or null
	LastReport         *string  `json:"lastReport"`         // BigInt (string) or null
	TotalDebt          *string  `json:"totalDebt"`          // BigInt (string) or null
	TotalDebtUsd       *float64 `json:"totalDebtUsd"`       // Float or null
	TotalGain          *string  `json:"totalGain"`          // BigInt (string) or null
	TotalGainUsd       *float64 `json:"totalGainUsd"`       // Float or null
	TotalLoss          *string  `json:"totalLoss"`          // BigInt (string) or null
	TotalLossUsd       *float64 `json:"totalLossUsd"`       // Float or null
	CurrentDebt        *string  `json:"currentDebt"`        // BigInt (string) or null
	CurrentDebtUsd     *float64 `json:"currentDebtUsd"`     // Float or null
	MaxDebt            *string  `json:"maxDebt"`            // BigInt (string) or null
	MaxDebtUsd         *float64 `json:"maxDebtUsd"`         // Float or null
	TargetDebtRatio    *float64 `json:"targetDebtRatio"`    // Float or null
	MaxDebtRatio       *float64 `json:"maxDebtRatio"`       // Float or null
}

type KongTVL struct {
	Close float64 `json:"close"`
}

type KongLastReportDetail struct {
	ChainID         int     `json:"chainId"`
	Address         string  `json:"address"`
	BlockNumber     string  `json:"blockNumber"`
	BlockTime       string  `json:"blockTime"`
	TransactionHash string  `json:"transactionHash"`
	Profit          *string `json:"profit"`
	ProfitUsd       *float64 `json:"profitUsd"`
	Loss            *string `json:"loss"`
	LossUsd         *float64 `json:"lossUsd"`
	APR             *struct {
		Gross *float64 `json:"gross"`
		Net   *float64 `json:"net"`
	} `json:"apr"`
}

type KongLenderStatus struct {
	Name    string   `json:"name"`
	Assets  *string  `json:"assets"`
	Rate    *float64 `json:"rate"`
	Address string   `json:"address"`
}

type KongClaim struct {
	ChainID    int      `json:"chainId"`
	Address    string   `json:"address"`
	Name       string   `json:"name"`
	Symbol     string   `json:"symbol"`
	Decimals   int      `json:"decimals"`
	Balance    *string  `json:"balance"`
	BalanceUsd *float64 `json:"balanceUsd"`
}

type KongRisk struct {
	Label               string   `json:"label"`
	AuditScore          *float64 `json:"auditScore"`
	CodeReviewScore     *float64 `json:"codeReviewScore"`
	ComplexityScore     *float64 `json:"complexityScore"`
	ProtocolSafetyScore *float64 `json:"protocolSafetyScore"`
	TeamKnowledgeScore  *float64 `json:"teamKnowledgeScore"`
	TestingScore        *float64 `json:"testingScore"`
}

type KongMeta struct {
	DisplayName *string  `json:"displayName"`
	Description *string  `json:"description"`
	Protocols   []string `json:"protocols"`
}

type KongStrategy struct {
	ChainID               int                   `json:"chainId"`
	Address               string                `json:"address"`
	APIVersion            *string               `json:"apiVersion"`
	BalanceOfWant         *string               `json:"balanceOfWant"`
	BaseFeeOracle         *string               `json:"baseFeeOracle"`
	CreditThreshold       *string               `json:"creditThreshold"`
	Crv                   *string               `json:"crv"`
	CurveVoter            *string               `json:"curveVoter"`
	DelegatedAssets       *string               `json:"delegatedAssets"`
	DoHealthCheck         *bool                 `json:"doHealthCheck"`
	EmergencyExit         *bool                 `json:"emergencyExit"`
	Erc4626               *bool                 `json:"erc4626"`
	EstimatedTotalAssets  *string               `json:"estimatedTotalAssets"`
	ForceHarvestTriggerOnce *bool               `json:"forceHarvestTriggerOnce"`
	Gauge                 *string               `json:"gauge"`
	HealthCheck           *string               `json:"healthCheck"`
	InceptTime            *string               `json:"inceptTime"`
	InceptBlock           *string               `json:"inceptBlock"`
	IsActive              *bool                 `json:"isActive"`
	IsBaseFeeAcceptable   *bool                 `json:"isBaseFeeAcceptable"`
	IsOriginal            *bool                 `json:"isOriginal"`
	Keeper                *string               `json:"keeper"`
	LocalKeepCRV          *string               `json:"localKeepCRV"`
	MaxReportDelay        *string               `json:"maxReportDelay"`
	MetadataURI           *string               `json:"metadataURI"`
	MinReportDelay        *string               `json:"minReportDelay"`
	Name                  *string               `json:"name"`
	Proxy                 *string               `json:"proxy"`
	Rewards               *string               `json:"rewards"`
	StakedBalance         *string               `json:"stakedBalance"`
	Strategist            *string               `json:"strategist"`
	TradeFactory          *string               `json:"tradeFactory"`
	Vault                 string                `json:"vault"`
	Want                  *string               `json:"want"`
	DOMAIN_SEPARATOR      *string               `json:"DOMAIN_SEPARATOR"`
	FACTORY               *string               `json:"FACTORY"`
	MAX_FEE               *int                  `json:"MAX_FEE"`
	MIN_FEE               *int                  `json:"MIN_FEE"`
	Decimals              *int                  `json:"decimals"`
	FullProfitUnlockDate  *string               `json:"fullProfitUnlockDate"`
	IsShutdown            *bool                 `json:"isShutdown"`
	LastReport            *string               `json:"lastReport"`
	LastReportDetail      *KongLastReportDetail `json:"lastReportDetail"`
	Management            *string               `json:"management"`
	PendingManagement     *string               `json:"pendingManagement"`
	PerformanceFee        *string               `json:"performanceFee"`
	PerformanceFeeRecipient *string             `json:"performanceFeeRecipient"`
	PricePerShare         *string               `json:"pricePerShare"`
	ProfitMaxUnlockTime   *string               `json:"profitMaxUnlockTime"`
	ProfitUnlockingRate   *string               `json:"profitUnlockingRate"`
	Symbol                *string               `json:"symbol"`
	TotalAssets           *string               `json:"totalAssets"`
	TotalDebt             *string               `json:"totalDebt"`
	TotalIdle             *string               `json:"totalIdle"`
	TotalSupply           *string               `json:"totalSupply"`
	TotalDebtUsd          *float64              `json:"totalDebtUsd"`
	LenderStatuses        []KongLenderStatus    `json:"lenderStatuses"`
	Claims                []KongClaim           `json:"claims"`
	Risk                  *KongRisk             `json:"risk"`
	Meta                  *KongMeta             `json:"meta"`
	V3                    bool                  `json:"v3"`
	Yearn                 bool                  `json:"yearn"`
}

type KongVault struct {
	Address           string           `json:"address"`
	ChainID           int              `json:"chainId"`
	Asset             KongAsset        `json:"asset"`
	Registry          string           `json:"registry"`
	InceptBlock       string           `json:"inceptBlock"`
	APIVersion        string           `json:"apiVersion"`
	WithdrawalQueue   []string         `json:"withdrawalQueue"`
	GetDefaultQueue   []string         `json:"get_default_queue"`
	Strategies        []string         `json:"strategies"`
	Debts             []KongDebt       `json:"debts"`
	TVL               *KongTVL         `json:"tvl"`
	APY               models.KongAPY   `json:"apy"`
}

type VaultsResponse struct {
	Vaults []KongVault `json:"vaults"`
}

type StrategiesResponse struct {
	Strategies []KongStrategy `json:"strategies"`
}

type Client struct {
	httpClient *http.Client
	endpoint   string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		endpoint: env.KONG_API_URL,
	}
}

func (c *Client) executeQuery(ctx context.Context, query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	request := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var graphQLResp GraphQLResponse
	if err := json.Unmarshal(body, &graphQLResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(graphQLResp.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL errors: %v", graphQLResp.Errors)
	}

	return &graphQLResp, nil
}

func (c *Client) FetchVaultsForChain(ctx context.Context, chainID uint64) ([]KongVault, error) {
	query := `
		query FetchVaults($chainId: Int!) {
			vaults(chainId: $chainId, yearn: true) {
				address
				chainId
				asset {
					chainId
					address
					decimals
					name
					symbol
				}
				registry
				inceptBlock
				apiVersion
				withdrawalQueue
				get_default_queue
				strategies
				debts {
					strategy
					performanceFee
					activation
					debtRatio
					minDebtPerHarvest
					maxDebtPerHarvest
					lastReport
					totalDebt
					totalDebtUsd
					totalGain
					totalGainUsd
					totalLoss
					totalLossUsd
					currentDebt
					currentDebtUsd
					maxDebt
					maxDebtUsd
					targetDebtRatio
					maxDebtRatio
				}
				tvl {
					close
				}
				apy {
					pricePerShare
					weeklyNet
					weeklyPricePerShare
					monthlyNet
					monthlyPricePerShare
					inceptionNet
					blockNumber
					blockTime
				}
			}
		}
	`

	variables := map[string]interface{}{
		"chainId": int(chainID),
	}

	resp, err := c.executeQuery(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vaults for chain %d: %w", chainID, err)
	}

	var vaultsResp VaultsResponse
	if err := json.Unmarshal(resp.Data, &vaultsResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vaults response: %w", err)
	}

	return vaultsResp.Vaults, nil
}

func (c *Client) FetchAllVaults(ctx context.Context) (map[uint64][]KongVault, error) {
	query := `
		query FetchAllVaults {
			vaults(yearn: true) {
				address
				chainId
				asset {
					chainId
					address
					decimals
					name
					symbol
				}
				registry
				inceptBlock
				apiVersion
				withdrawalQueue
				get_default_queue
				strategies
				debts {
					strategy
					performanceFee
					activation
					debtRatio
					minDebtPerHarvest
					maxDebtPerHarvest
					lastReport
					totalDebt
					totalDebtUsd
					totalGain
					totalGainUsd
					totalLoss
					totalLossUsd
					currentDebt
					currentDebtUsd
					maxDebt
					maxDebtUsd
					targetDebtRatio
					maxDebtRatio
				}
				tvl {
					close
				}
				apy {
					pricePerShare
					weeklyNet
					weeklyPricePerShare
					monthlyNet
					monthlyPricePerShare
					inceptionNet
					blockNumber
					blockTime
				}
			}
		}
	`

	resp, err := c.executeQuery(ctx, query, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all vaults: %w", err)
	}

	var vaultsResp VaultsResponse
	if err := json.Unmarshal(resp.Data, &vaultsResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vaults response: %w", err)
	}

	vaultsByChain := make(map[uint64][]KongVault)
	for _, vault := range vaultsResp.Vaults {
		chainID := uint64(vault.ChainID)
		vaultsByChain[chainID] = append(vaultsByChain[chainID], vault)
	}

	return vaultsByChain, nil
}

func (v *KongVault) GetAddress() common.Address {
	return common.HexToAddress(v.Address)
}

func (v *KongVault) GetAPY() models.KongAPY {
	apy := v.APY
	// Include token decimals from asset for PPS normalization
	apy.Decimals = uint64(v.Asset.Decimals)
	return apy
}

func (v *KongVault) GetRegistry() common.Address {
	if v.Registry == "" {
		return common.Address{}
	}
	return common.HexToAddress(v.Registry)
}

func (v *KongVault) GetBlockNumber() uint64 {
	if v.InceptBlock == "" {
		return 0
	}
	blockNum, err := strconv.ParseUint(v.InceptBlock, 10, 64)
	if err != nil {
		return 0
	}
	return blockNum
}

func (v *KongVault) GetAssetAddress() common.Address {
	if v.Asset.Address == "" {
		return common.Address{}
	}
	return common.HexToAddress(v.Asset.Address)
}

func (v *KongVault) GetStrategies() []common.Address {
	strategySet := make(map[common.Address]bool)
	var strategies []common.Address

	// Combine strategies from all available sources (not prioritized fallback)
	// This matches the original yDaemon approach of getting all strategies

	// Add from WithdrawalQueue
	if v.WithdrawalQueue != nil && len(v.WithdrawalQueue) > 0 {
		for _, s := range v.WithdrawalQueue {
			if s != "" && s != "0x0000000000000000000000000000000000000000" {
				addr := common.HexToAddress(s)
				if !strategySet[addr] {
					strategySet[addr] = true
					strategies = append(strategies, addr)
				}
			}
		}
	}

	// Add from GetDefaultQueue
	if v.GetDefaultQueue != nil && len(v.GetDefaultQueue) > 0 {
		for _, s := range v.GetDefaultQueue {
			if s != "" && s != "0x0000000000000000000000000000000000000000" {
				addr := common.HexToAddress(s)
				if !strategySet[addr] {
					strategySet[addr] = true
					strategies = append(strategies, addr)
				}
			}
		}
	}

	// Add from Strategies
	if v.Strategies != nil && len(v.Strategies) > 0 {
		for _, s := range v.Strategies {
			if s != "" && s != "0x0000000000000000000000000000000000000000" {
				addr := common.HexToAddress(s)
				if !strategySet[addr] {
					strategySet[addr] = true
					strategies = append(strategies, addr)
				}
			}
		}
	}

	return strategies
}

type KongVaultData struct {
	Vault      KongVault
	Strategies []common.Address
	Debts      []KongDebt
	TVL        float64
	APY        models.KongAPY
}

func (v *KongVault) GetTVL() float64 {
	if v.TVL == nil {
		return 0
	}
	return v.TVL.Close
}

func (v *KongVault) GetDebts() []KongDebt {
	if v.Debts == nil {
		return []KongDebt{}
	}
	return v.Debts
	
}


func (c *Client) FetchStrategiesForChain(ctx context.Context, chainID uint64) ([]KongStrategy, error) {
	query := `
		query Strategies($chainId: Int!) {
			strategies(chainId: $chainId) {
				chainId
				address
				apiVersion
				balanceOfWant
				baseFeeOracle
				creditThreshold
				crv
				curveVoter
				delegatedAssets
				doHealthCheck
				emergencyExit
				erc4626
				estimatedTotalAssets
				forceHarvestTriggerOnce
				gauge
				healthCheck
				inceptTime
				inceptBlock
				isActive
				isBaseFeeAcceptable
				isOriginal
				keeper
				localKeepCRV
				maxReportDelay
				metadataURI
				minReportDelay
				name
				proxy
				rewards
				stakedBalance
				strategist
				tradeFactory
				vault
				want
				DOMAIN_SEPARATOR
				FACTORY
				MAX_FEE
				MIN_FEE
				decimals
				fullProfitUnlockDate
				isShutdown
				lastReport
				lastReportDetail {
					chainId
					address
					blockNumber
					blockTime
					transactionHash
					profit
					profitUsd
					loss
					lossUsd
					apr {
						gross
						net
					}
				}
				management
				pendingManagement
				performanceFee
				performanceFeeRecipient
				pricePerShare
				profitMaxUnlockTime
				profitUnlockingRate
				symbol
				totalAssets
				totalDebt
				totalIdle
				totalSupply
				totalDebtUsd
				lenderStatuses {
					name
					assets
					rate
					address
				}
				claims {
					chainId
					address
					name
					symbol
					decimals
					balance
					balanceUsd
				}
				risk {
					label
					auditScore
					codeReviewScore
					complexityScore
					protocolSafetyScore
					teamKnowledgeScore
					testingScore
				}
				meta {
					displayName
					description
					protocols
				}
				v3
				yearn
			}
		}
	`

	variables := map[string]interface{}{
		"chainId": int(chainID),
	}

	resp, err := c.executeQuery(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch strategies for chain %d: %w", chainID, err)
	}

	var strategiesResp StrategiesResponse
	if err := json.Unmarshal(resp.Data, &strategiesResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal strategies response: %w", err)
	}

	return strategiesResp.Strategies, nil
}

func (s *KongStrategy) GetAddress() common.Address {
	return common.HexToAddress(s.Address)
}

func (s *KongStrategy) GetVaultAddress() common.Address {
	return common.HexToAddress(s.Vault)
}

func (s *KongStrategy) GetAPIVersion() string {
	if s.APIVersion == nil {
		return ""
	}
	return *s.APIVersion
}

func (s *KongStrategy) GetName() string {
	if s.Name == nil {
		return ""
	}
	return *s.Name
}

func (s *KongStrategy) GetIsActive() bool {
	if s.IsActive == nil {
		return false
	}
	return *s.IsActive
}

func (s *KongStrategy) GetDoHealthCheck() bool {
	if s.DoHealthCheck == nil {
		return false
	}
	return *s.DoHealthCheck
}

func FetchVaultsFromKong(chainID uint64) (map[common.Address]KongVaultData, error) {
	ctx := context.Background()
	client := NewClient()

	vaults, err := client.FetchVaultsForChain(ctx, chainID)
	if err != nil {
		logs.Error(chainID, `-`, `Failed to fetch vaults from Kong: %v`, err)
		return nil, err
	}

	vaultData := make(map[common.Address]KongVaultData)
	for _, vault := range vaults {
		vaultAddr := vault.GetAddress()
		strategies := vault.GetStrategies()
		vaultData[vaultAddr] = KongVaultData{
			Vault:      vault,
			Strategies: strategies,
			APY:        vault.GetAPY(),
			Debts:      vault.GetDebts(),
			TVL:        vault.GetTVL(),
		}
	}

	logs.Success(chainID, `-`, `Fetched %d vaults from Kong`, len(vaults))
	return vaultData, nil
}

// FetchStrategiesFromKong fetches all strategies for a given chain from Kong GraphQL API
// and returns them mapped by vault address
func FetchStrategiesFromKong(chainID uint64) (map[common.Address][]KongStrategy, error) {
	ctx := context.Background()
	client := NewClient()

	strategies, err := client.FetchStrategiesForChain(ctx, chainID)
	if err != nil {
		logs.Error(chainID, `-`, `Failed to fetch strategies from Kong: %v`, err)
		return nil, err
	}

	// Map strategies by vault address as requested
	strategiesByVault := make(map[common.Address][]KongStrategy)
	for _, strategy := range strategies {
		vaultAddr := strategy.GetVaultAddress()
		strategiesByVault[vaultAddr] = append(strategiesByVault[vaultAddr], strategy)
	}

	logs.Success(chainID, `-`, `Fetched %d strategies from Kong, mapped to %d vaults`, len(strategies), len(strategiesByVault))
	return strategiesByVault, nil
}
