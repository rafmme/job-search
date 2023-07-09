package main

import (
	"github.com/rafmme/job-search/api"
	"github.com/rafmme/job-search/util"
)

func main() {
	server := api.CreateServer()
	util.RunCronJobs()
	server.Start()
}
