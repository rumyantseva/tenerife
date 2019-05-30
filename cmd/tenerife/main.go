package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/rumyantseva/tenerife/internal/application"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.Info("Starting the application")

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Port is not provided")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler(logger))

	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Server stopped with an error: %v", err)
	}
}
