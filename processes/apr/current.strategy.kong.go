package apr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/yearn/ydaemon/common/bigNumber"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/internal/models"
	"github.com/yearn/ydaemon/internal/storage"
)

const KONG_API_URL = "https://kong.yearn.farm/api/gql"

/**************************************************************************************************
** getKongStrategyReport retrieves the latest strategy report from Kong API
** This function handles the setup and execution of the Kong HTTP request
**
** @param strategyAddress string - The address of the strategy to query
** @param chainID uint64 - The chain ID of the network
** @return models.TReportsFromKong - The response from Kong containing strategy report data
** @return error - Any error encountered during the request
**************************************************************************************************/
func getKongStrategyReport(strategyAddress string, chainID uint64) (models.TReportsFromKong, error) {
	query := fmt.Sprintf(`{"query":"query Query {\n  strategy(address: \"%s\", chainId: %d) {\n    lastReportDetail {\n      apr {\n        net\n        gross\n      }\n      loss\n      lossUsd\n      profit\n      profitUsd\n      blockTime\n      blockNumber\n    }\n    address\n  }\n}","variables":{"chainId":%d,"address":"%s"},"operationName":"Query"}`,
		strategyAddress, chainID, chainID, strategyAddress)

	req, err := http.NewRequest("POST", KONG_API_URL, strings.NewReader(query))
	if err != nil {
		return models.TReportsFromKong{}, err
	}

	// Set all required headers
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.TReportsFromKong{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.TReportsFromKong{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return models.TReportsFromKong{}, fmt.Errorf("Kong API returned status %d: %s", resp.StatusCode, string(body))
	}

	var kongResponseRaw models.TReportsFromKong
	if err := json.Unmarshal(body, &kongResponseRaw); err != nil {
		return models.TReportsFromKong{}, err
	}

	return kongResponseRaw, nil
}

/**************************************************************************************************
** GetCurrentStrategyAPRFromKong retrieves the latest APR for a specific strategy from Kong.
** It returns the APR value as a float64, or 0 if the APR cannot be determined.
**
** @param chainID The blockchain network ID
** @param strategyAddress The strategy address to check
** @return float64 The latest APR value or 0 if not available
** @return error An error if the request fails
**************************************************************************************************/
func GetCurrentStrategyAPRFromKong(chainID uint64, strategyAddress string) (*bigNumber.Float, error) {
	// Get chain configuration
	_, ok := env.GetChain(chainID)
	if !ok {
		return bigNumber.NewFloat(0), errors.New("chain not found")
	}

	// Get the latest report from Kong
	kongResponseRaw, err := getKongStrategyReport(strategyAddress, chainID)
	if err != nil {
		return bigNumber.NewFloat(0), err
	}

	return bigNumber.NewFloat(kongResponseRaw.Data.Strategy.LastReportDetail.Apr.Gross), nil
}

/**************************************************************************************************
** getStrategyReportFromDB retrieves the latest strategy report from the database
** This function handles the database query and returns the latest report
**
** @param strategyAddress string - The address of the strategy to query
** @param chainID uint64 - The chain ID of the network
** @return models.TStrategyReportDB - The response from the database containing strategy report data
** @return error - Any error encountered during the query
**************************************************************************************************/
func getStrategyReportFromDB(strategyAddress string, chainID uint64) (models.TStrategyReportDB, error) {
	var report models.TStrategyReportDB

	query := `
		SELECT
			chain_id AS "chainId",
			address,
			event_name AS "eventName",
			hook->'apr' AS "apr"
		FROM evmlog
		WHERE
			(chain_id = ?) AND (address = ?)
			AND (event_name = 'Reported' OR event_name = 'Harvested')
		ORDER BY
			block_time DESC, log_index DESC
		LIMIT 1`

	err := storage.GetDB().Raw(query, chainID, strategyAddress).Scan(&report).Error
	if err != nil {
		return models.TStrategyReportDB{}, err
	}

	return report, nil
}

/**************************************************************************************************
** GetCurrentStrategyAPRFromDB retrieves the latest APR for a specific strategy from the database.
** It returns the APR value as a float64, or 0 if the APR cannot be determined.
**
** @param chainID The blockchain network ID
** @param strategyAddress The strategy address to check
** @return float64 The latest APR value or 0 if not available
** @return error An error if the request fails
**************************************************************************************************/
func GetCurrentStrategyAPRFromDB(chainID uint64, strategyAddress string) (*bigNumber.Float, error) {
	// Get chain configuration
	_, ok := env.GetChain(chainID)
	if !ok {
		return bigNumber.NewFloat(0), errors.New("chain not found")
	}

	// Get the latest report from the database
	report, err := getStrategyReportFromDB(strategyAddress, chainID)
	if err != nil {
		return bigNumber.NewFloat(0), err
	}

	// Parse the APR from the JSON field
	var aprData struct {
		Gross float64 `json:"gross"`
		Net   float64 `json:"net"`
	}
	if err := json.Unmarshal(report.APR, &aprData); err != nil {
		return bigNumber.NewFloat(0), err
	}

	return bigNumber.NewFloat(aprData.Gross), nil
}
