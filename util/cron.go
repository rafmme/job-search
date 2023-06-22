package util

import (
	"fmt"
	"log"
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

		if len(searchResult.Jobs) < 1 {
			log.Println("Search Result empty!")
			return
		}

		SendEmail(
			recipientAddr, searchResult.SearchQuery,
			searchResult.CreateEmailHTMLString(), "ext",
		)
	})

	c.Start()
}
