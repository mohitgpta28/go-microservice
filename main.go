package main

import (
	"github.com/eatliftcode/go-microservice/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
