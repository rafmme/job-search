package util

import (
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/go-telegram/bot"
	_ "github.com/go-telegram/bot/models"
)

var (
	TgBot *tgbotapi.BotAPI
)

func SetupTGBot() {
	var err error
	TgBot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))

	if err != nil {
		log.Printf("Error with Telegram Bot: %v", err)
	}

	TgBot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := TgBot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if strings.ToLower(update.Message.Text) == "jobs" || strings.ToLower(update.Message.Text) == "/jobs" {
			var tgMessageList [][]string
			myTGId, err := strconv.ParseInt(os.Getenv("TG_ID"), 10, 64)

			if err != nil {
				log.Printf("Invalid Recipient Telegram ID: %v", err)
			}

			if update.Message.Chat.ID == myTGId {
				_, tgMessageList = GetMyJobs("ser-api")
			} else {
				_, tgMessageList = GetMyJobs("g-cse")
			}

			for _, tgMsg := range tgMessageList {
				for _, msg := range tgMsg {
					ReplyTGBotMessage(
						msg,
						update.Message.Chat.ID, update.Message.MessageID,
					)
				}
			}

			return
		}

		ReplyTGBotMessage(
			"😟 Sorry, Still work in progress!",
			update.Message.Chat.ID, update.Message.MessageID,
		)
	}

}

func ReplyTGBotMessage(message string, recipientId int64, messageId int) {
	msg := tgbotapi.NewMessage(recipientId, message)
	msg.ReplyToMessageID = messageId

	if _, err := TgBot.Send(msg); err != nil {
		log.Printf("Error with Telegram Bot: %v", err)
	}
}

func SendTGBotMessage(message string, recipientId int64) {
	msg := tgbotapi.NewMessage(recipientId, message)

	if _, err := TgBot.Send(msg); err != nil {
		log.Printf("Error with Telegram Bot: %v", err)
	}
}
