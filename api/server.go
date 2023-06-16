package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/opensaucerer/barf"
)

type Server struct {
	listenAddr string
}

func CreateServer() *Server {
	err := godotenv.Load()
	port := os.Getenv("PORT")

	if err != nil || port == "" {
		log.Println("Error loading .env file or PORT variable empty!")
		port = "7590"
	}

	return &Server{
		listenAddr: port,
	}
}

func (server *Server) Start() {
	allow := true

	if err := barf.Stark(barf.Augment{
		Port:     server.listenAddr,
		Logging:  &allow, // enable request logging
		Recovery: &allow, // enable panic recovery so barf returns a 500 error instead of crashing
	}); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}

	barf.Get("/", Home)
	barf.Post("/api/v1/search", SearchJobs)

	if err := barf.Beck(); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}
}
