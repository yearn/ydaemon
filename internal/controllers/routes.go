package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/machinebox/graphql"
	"github.com/majorfi/ydaemon/internal/daemons"
	"github.com/majorfi/ydaemon/internal/ethereum"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/models"
	"github.com/majorfi/ydaemon/internal/store"
	"github.com/majorfi/ydaemon/internal/utils"
)

type controller struct{}

//GetAllVaults will, for a given chainID, return a list of all vaults
func (y controller) GetAllVaults(c *gin.Context) {
	// get the chainID from the URI
	chainID, err := strconv.ParseUint(c.Param("chainID"), 10, 64)
	if err != nil {
		chainID = 1
		c.String(http.StatusBadRequest, "Invalid chainID")
		return
	}

	client := graphql.NewClient(ethereum.GetGraphURI(chainID))
	request := graphql.NewRequest(`
        {
			vaults(first: 1000) {
				id
				activation
				apiVersion
				classification
				managementFeeBps
				performanceFeeBps
				balanceTokens
				latestUpdate {
					timestamp
				}
				shareToken {
					name
					symbol
					id
					decimals
				}
				token {
					name
					symbol
					id
					decimals
				}
				strategies(first: 40) {
					address
					name
					inQueue
					debtLimit
				}
			}
        }
    `)
	var response models.TGraphQueryResponseForVaults
	if err := client.Run(context.Background(), request, &response); err != nil {
		logs.Error(err)
		c.String(http.StatusBadRequest, "Invalid chainID")
		return
	}

	data := []*models.TVault{}
	for _, vaultFromGraph := range response.Vaults {
		vaultFromMeta := store.VaultsFromMeta[chainID][common.HexToAddress(vaultFromGraph.Id).String()]
		tokenFromMeta := store.TokensFromMeta[chainID][common.HexToAddress(vaultFromGraph.Token.Id).String()]
		apyFromAPIV1 := store.VaultsFromAPIV1[chainID][common.HexToAddress(vaultFromGraph.Id).String()]
		strategiesFromMeta := store.StrategiesFromMeta[chainID]
		pricesForChainID := store.TokenPrices[chainID]

		data = append(data, prepareVaultSchema(
			vaultFromGraph,
			vaultFromMeta,
			tokenFromMeta,
			strategiesFromMeta,
			apyFromAPIV1,
			pricesForChainID,
		))
	}
	c.JSON(http.StatusOK, data)
}

//TriggerMetaRefreshWebhook will do trigger a webhook from github
func (y controller) TriggerMetaRefreshWebhook(c *gin.Context) {
	//Check if the webhook has correct data and secret
	hook, _ := github.New(github.Options.Secret(utils.WebhookSecret))
	payload, err := hook.Parse(c.Request, github.DeploymentStatusEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := payload.(github.DeploymentStatusPayload)

	//Perform some extra checks
	isSuccess := data.DeploymentStatus.State == "success"
	isProduction := data.Deployment.Environment == "Production"
	isVercelBot := data.Sender.Login == "vercel[bot]"
	isBot := data.Sender.Type == "Bot"

	logs.Pretty(isSuccess, isProduction, isVercelBot, isBot)
	if isSuccess && isProduction && isVercelBot && isBot {
		logs.Info("Triggering meta refresh in 1 minute")
		go func() {
			//Let's wait one minute just to be sure
			time.Sleep(time.Minute)
			//Update meta information
			go daemons.FetchStrategiesFromMeta(1)
			go daemons.FetchStrategiesFromMeta(250)
			go daemons.FetchStrategiesFromMeta(42161)

			//Update Strategies information
			go daemons.FetchTokensFromMeta(1)
			go daemons.FetchTokensFromMeta(250)
			go daemons.FetchTokensFromMeta(42161)

			//Update Vaults information
			go daemons.FetchVaultsFromMeta(1)
			go daemons.FetchVaultsFromMeta(250)
			go daemons.FetchVaultsFromMeta(42161)
		}()
	}
	c.JSON(http.StatusAccepted, nil)
}
