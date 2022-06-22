package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/majorfi/ydaemon/internal/daemons"
	"github.com/majorfi/ydaemon/internal/logs"
	"github.com/majorfi/ydaemon/internal/utils"
)

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
