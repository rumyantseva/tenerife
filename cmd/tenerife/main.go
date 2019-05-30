package main

import (
	"github.com/rumyantseva/tenerife/internal/diagnostics"
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

	r.HandleFunc("/healthz", diagnostics.LivenessHandler(logger))
	r.HandleFunc("/readyz", diagnostics.ReadinessHandler(logger))


	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Server stopped with an error: %v", err)
	}
}
