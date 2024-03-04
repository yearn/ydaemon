package main

import (
	"os"
	"strconv"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yearn/ydaemon/common/ethereum"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/internal"
	"github.com/yearn/ydaemon/processes/apr"
	"github.com/yearn/ydaemon/processes/initDailyBlock"
	"github.com/yearn/ydaemon/processes/vaultsMigrations"
)

func triggerTgMessage(message string) {
	telegramToken, ok := os.LookupEnv("TELEGRAM_BOT")
	if !ok {
		logs.Error(`TELEGRAM_BOT environment variable not set`)
		return
	}
	telegramChatStr, ok := os.LookupEnv("TELEGRAM_CHAT")
	if !ok {
		logs.Error(`TELEGRAM_CHAT environment variable not set`)
		return
	}
	telegramChat, err := strconv.ParseInt(telegramChatStr, 10, 64)
	if err != nil {
		logs.Error(`TELEGRAM_CHAT environment variable is not a valid chat ID`)
		return
	}
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		logs.Error(`Error initializing Telegram bot: ` + err.Error())
		return
	}
	m, err := bot.Send(tgbotapi.NewMessage(telegramChat, message))
	if err != nil {
		logs.Error(`Error sending message to Telegram: ` + err.Error())
	}
	_ = m
}

func main() {
	initFlags()
	summonDaemonsForAllChains(chains)
	var wg sync.WaitGroup

	logs.Info(`Running external processes...`)

	/** 🔵 - Yearn *************************************************************************************
	** This section of the code is responsible for running external processes. Each process is associated
	** with a specific chain ID and is run concurrently for efficiency. The processes include fetching
	** and updating prices, initializing V2, building token lists, running vault migrations, initializing
	** daily blocks, initializing APY, etc.
	**
	** The status of each process is tracked and updated for each chain ID. Once all processes for a
	** chain ID have completed, the status is set to "OK".
	**
	** The server is ready to accept requests once all processes have completed.
	**************************************************************************************************/
	switch process {
	case ProcessServer:
		logs.Info(`Running yDaemon server process...`)
		go NewRouter().Run(`:8080`)
		go triggerTgMessage(`💛 - yDaemon server is ready to accept requests: https://ydaemon.yearn.fi/`)

		for _, chainID := range chains {
			setStatusForChainID(chainID, "Loading")
			logs.Info(`Getting WS client for chain ` + strconv.FormatUint(chainID, 10))
			ethereum.GetWSClient(chainID)
			ethereum.InitBlockTimestamp(chainID)
			wg.Add(1)
			go func(chainID uint64) {
				internal.InitializeV2(chainID, nil)
				triggerTgMessage(`✅ - yDaemon V2 initialized for chain ` + strconv.FormatUint(chainID, 10))
				wg.Done()
			}(chainID)
		}
		wg.Wait()
		for _, chainID := range chains {
			setStatusForChainID(chainID, "OK")
		}

		logs.Success(`Server ready on port 8080 !`)
		select {}

	case ProcessVaultMigrations:
		logs.Info(`Running yDaemon vault migrations process...`)
		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				vaultsMigrations.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()

	case ProcessInitDailyBlock:
		logs.Info(`Running yDaemon DailyBlock Initializer process...`)
		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				initDailyBlock.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()

	case ProcessAPY:
		logs.Info(`Running yDaemon APY Initializer process...`)
		for _, chainID := range chains {
			wg.Add(1)
			go func(chainID uint64) {
				apr.Run(chainID)
				wg.Done()
			}(chainID)
		}
		wg.Wait()
	}
}
