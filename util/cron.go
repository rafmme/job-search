package util

import (
	"fmt"
	"os"

	"github.com/rafmme/job-search/types"
	"gopkg.in/robfig/cron.v2"
)

func RunCronJobs() {
	recipientAddr := os.Getenv("RCV_ADDR")
	c := cron.New()

	c.AddFunc("@every 15s", func() {
		fmt.Println("Every hour on the half hour")
		sqData := new(types.SearchQueryData)
		sqData.From = 2
		searchResult := sqData.CreateJobSearchQuery().Execute().FormatJobList()
		SendEmail(recipientAddr, searchResult.SearchQuery, searchResult.CreateEmailHTMLString())
	})

	c.Start()
}
