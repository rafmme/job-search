package main

import (
	_ "github.com/joho/godotenv"
	_ "github.com/opensaucerer/barf"
	_ "github.com/opensaucerer/goaxios"
	"github.com/rafmme/job-search/api"
	_ "github.com/xhit/go-simple-mail/v2"
	_ "gopkg.in/robfig/cron.v2"
)

func main() {
	server := api.CreateServer()
	server.Start()
}
