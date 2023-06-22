package main

import (
	"github.com/rafmme/job-search/api"
	"github.com/rafmme/job-search/util"
	_ "github.com/serpapi/google-search-results-golang"
)

func main() {
	server := api.CreateServer()
	util.RunCronJobs()
	server.Start()
}
