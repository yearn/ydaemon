package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/yearn/ydaemon/common/env"
	"github.com/yearn/ydaemon/common/helpers"
	"github.com/yearn/ydaemon/common/logs"
	"github.com/yearn/ydaemon/processes/prices"
)

var initializedCounter = 0

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

func triggerInitializedStatus(chainID uint64) {
	initializedCounter++
	triggerTgMessage(`✅ - yDaemon initialized for chain ` + strconv.FormatUint(chainID, 10) + ` (` + strconv.Itoa(initializedCounter) + `/` + strconv.Itoa(len(chains)) + `)`)
}

func listenToSignals() {
	telegramToken, ok := os.LookupEnv("TELEGRAM_BOT")
	if !ok {
		logs.Error(`TELEGRAM_BOT environment variable not set`)
		return
	}
	telegramWhitelistedUsers, ok := os.LookupEnv("TELEGRAM_WHITELIST")
	if !ok {
		logs.Error(`TELEGRAM_WHITELIST environment variable not set`)
		return
	}
	telegramWhitelistedUsersArray := strings.Split(telegramWhitelistedUsers, `,`)
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		logs.Error(`Error initializing Telegram bot: ` + err.Error())
		return
	}
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		lowercaseUserName := strings.ToLower(update.Message.From.UserName)
		if !helpers.Contains(telegramWhitelistedUsersArray, lowercaseUserName) {
			logs.Error(`Unauthorized user: ` + update.Message.From.UserName)
			msg.Text = "You are not authorized to use this bot."
			bot.Send(msg)
			continue
		}
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			triggerTgMessage(`Available commands:
- /help: Show this help message
- /restart: Restart the daemon
- /update: Update yDaemon with the latest version
- /upd_prices <chainID>: Update the prices for a given chain
- /origins: Get the origins of access`)
		case "restart":
			triggerTgMessage(`🔴 - ` + update.Message.From.UserName + ` asked for a restart`)
			os.Exit(1)
		case "update":
			//this might be useless
			reason := ` without a reason`
			arguments := update.Message.CommandArguments()
			if arguments != "" {
				reason = ` because: ` + arguments
			}
			triggerTgMessage(`♻️ - ` + update.Message.From.UserName + ` asked to update yDaemon away from v` + getVersion() + reason)

			//Grabbing the current executable name
			execName, _ := os.Executable()

			//Checkout local changes
			cmd := exec.Command("git", "checkout", "--", ".")
			cmd.Dir = filepath.Dir(execName)
			if err := cmd.Run(); err != nil {
				triggerTgMessage(`🔴 - Error checking out local changes: ` + err.Error())
				continue
			}

			//Pulling changes
			cmd = exec.Command("git", "pull")
			cmd.Dir = filepath.Dir(execName)
			if err := cmd.Run(); err != nil {
				triggerTgMessage(`🔴 - Error pulling changes: ` + err.Error())
				continue
			}

			//service ydaemon restart
			os.Exit(1)
		case "origins":
			listOfOrigins := []string{}
			itemsInLimiter := limiterSet.Items()
			for item := range itemsInLimiter {
				listOfOrigins = append(listOfOrigins, item)
			}
			triggerTgMessage(`👀 - Origins of access:` + "\n" + strings.Join(listOfOrigins, "\n"))
		case "upd_prices":
			arguments := update.Message.CommandArguments()
			if arguments == "" {
				triggerTgMessage(`🔴 - Incorrect format. Should be /upd_prices <chainID>`)
				continue
			}
			chainID, err := strconv.ParseUint(arguments, 10, 64)
			if err != nil {
				triggerTgMessage(`🔴 - Incorrect format. Should be /upd_prices <chainID> (number)`)
				continue
			}
			if _, ok := env.CHAINS[chainID]; !ok {
				triggerTgMessage(`🔴 - Chain not supported`)
				continue
			}
			triggerTgMessage(`💰 - ` + update.Message.From.UserName + ` asked for a price update for chain ` + strconv.FormatUint(chainID, 10))
			prices.UpdatePrices(chainID)
		default:
			msg.Text = "I don't know that command"
			bot.Send(msg)
		}
	}
}
