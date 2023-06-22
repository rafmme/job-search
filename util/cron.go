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

	c.AddFunc("@every 6h", func() {
		fmt.Println("For NodeJS jobs")
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

	c.AddFunc("@every 6h", func() {
		fmt.Println("For JavaScript jobs")
		sqData := new(types.SearchQueryData)
		sqData.From = 2
		sqData.Include = []string{"javascript"}
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

	c.AddFunc("@every 8h", func() {
		fmt.Println("For Go jobs")
		sqData := new(types.SearchQueryData)
		sqData.From = 2
		sqData.Include = []string{"golang", "go"}
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

	c.AddFunc("@every 8h", func() {
		fmt.Println("For Rust jobs")
		sqData := new(types.SearchQueryData)
		sqData.From = 2
		sqData.Include = []string{"rust"}
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
