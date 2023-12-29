package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/opensaucerer/barf"
)

type Server struct {
	listenAddr string
}

func requestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		log.Printf("%s %s - %v", r.Method, r.URL.Path, duration)
	})
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
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()

	mux.Handle("/", fs)

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui"))))
	mux.HandleFunc("/docs", SwaggerAPIDocHandler)
	mux.HandleFunc("/jobs", FetchJobs)
	mux.HandleFunc("/api/v1/search", SearchJobs)

	//util.SetupTGBot()

	loggedMux := requestLoggerMiddleware(mux)
	port := fmt.Sprintf(":%s", server.listenAddr)
	barf.Logger().Info(fmt.Sprintf("🆙 Server up on PORT %s", port))
	err := http.ListenAndServe(port, loggedMux)

	if err != nil {
		barf.Logger().Error("Could'nt start the server. " + err.Error())
		os.Exit(1)
	}
}
