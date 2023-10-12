package strategies

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal/models"
)

func graphQLRequestForReports(strategyAddress string, c *gin.Context) *graphql.Request {
	return graphql.NewRequest(`{
		strategy(id: "` + strings.ToLower(strategyAddress) + `") {
			` + helpers.GetStrategyReports() + `
		}
	}`)
}

// GetReports will, for a given strategy on a given chainID, return a list of reports
func (y Controller) GetReports(c *gin.Context) {
	chainID, ok := helpers.AssertChainID(c.Param("chainID"))
	if !ok {
		c.String(http.StatusBadRequest, "invalid chainID")
		return
	}
	address, ok := helpers.AssertAddress(c.Param("address"), chainID)
	if !ok {
		c.String(http.StatusBadRequest, "invalid address")
		return
	}

	graphQLEndpoint := env.CHAINS[chainID].SubgraphURI
	if graphQLEndpoint == "" {
		logs.Error("No graph endpoint for chainID", chainID)
		c.String(http.StatusInternalServerError, "impossible to fetch subgraph")
		return
	}
	client := graphql.NewClient(graphQLEndpoint)
	request := graphQLRequestForReports(address.Hex(), c)
	var responseRaw models.TReportsFromGraph
	if err := client.Run(context.Background(), request, &responseRaw); err != nil {
		logs.Error(err)
		c.String(http.StatusInternalServerError, "invalid graphQL response")
		return
	}

	var response models.TReportsFromGraphToClient
	response.Strategy.Reports = []models.TReport{}
	for _, report := range responseRaw.Strategy.Reports {
		timestamp, err := strconv.ParseUint(report.Timestamp, 10, 64)
		if err != nil {
			timestamp = 0
		}
		newReport := models.TReport{
			ID:        report.ID,
			DebtAdded: report.DebtAdded,
			DebtLimit: report.DebtLimit,
			TotalDebt: report.TotalDebt,
			Gain:      report.Gain,
			TotalGain: report.TotalGain,
			Loss:      report.Loss,
			TotalLoss: report.TotalLoss,
			DebtPaid:  report.DebtPaid,
			Timestamp: timestamp,
			Results:   []models.TReportResult{},
		}
		for _, allResults := range report.Results {
			duration, err := strconv.ParseUint(allResults.Duration, 10, 64)
			if err != nil {
				duration = 0
			}
			durationPr, err := strconv.ParseUint(allResults.DurationPr, 10, 64)
			if err != nil {
				durationPr = 0
			}
			apr, err := strconv.ParseFloat(allResults.Apr, 64)
			if err != nil {
				apr = 0
			}

			newResult := models.TReportResult{
				Duration:   duration,
				DurationPR: durationPr,
				APR:        apr,
			}
			newReport.Results = append(newReport.Results, newResult)
		}
		response.Strategy.Reports = append(response.Strategy.Reports, newReport)

	}

	c.JSON(http.StatusOK, response.Strategy.Reports)
}
