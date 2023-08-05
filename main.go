package main

import (
	"github.com/rafmme/job-search/api"
)

func main() {
	server := api.CreateServer()
	//util.RunCronJobs()
	server.Start()
}
