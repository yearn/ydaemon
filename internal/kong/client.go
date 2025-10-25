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
	"github.com/yearn/ydaemon/common/bigNumber"
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
	TotalAssets       *bigNumber.Int   `json:"totalAssets"`
}

type VaultsResponse struct {
	Vaults []KongVault `json:"vaults"`
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
				totalAssets
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
				totalAssets
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
	TotalAssets *bigNumber.Int
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
			Vault:       vault,
			Strategies:  strategies,
			APY:         vault.GetAPY(),
			Debts:       vault.GetDebts(),
			TVL:         vault.GetTVL(),
			TotalAssets: vault.TotalAssets,
		}
	}

	logs.Success(chainID, `-`, `Fetched %d vaults from Kong`, len(vaults))
	return vaultData, nil
}
