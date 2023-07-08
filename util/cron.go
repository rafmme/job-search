package util

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/robfig/cron.v2"
)

func RunCronJobs() {
	recipientAddr := os.Getenv("RCV_ADDR")
	tgRecipientUserId, err := strconv.ParseInt(os.Getenv("TG_ID"), 10, 64)

	if err != nil {
		log.Printf("Invalid Recipient Telegram ID: %v", err)
	}

	c := cron.New()
	c.AddFunc("@every 8h10m", func() {
		fmt.Println("For All Jobs")
		var emailContent string
		var tgMessageList []string

		emailContent, tgMessageList = GetMyJobs()

		for len(tgMessageList) < 1 {
			emailContent, tgMessageList = GetMyJobs()
		}

		if len(tgMessageList) > 0 {
			for _, tgMsg := range tgMessageList {
				SendTGBotMessage(tgMsg, tgRecipientUserId)
			}

			SendEmail(
				recipientAddr, "Posted Jobs List",
				emailContent, "ext",
			)
		}

	})

	c.Start()
}
