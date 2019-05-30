package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/rumyantseva/tenerife/internal/application"
	"github.com/rumyantseva/tenerife/internal/diagnostics"
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

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan error, 1)

	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		shutdown <- err
	}()

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			logger.Print("Got SIGINT...")
		case syscall.SIGTERM:
			logger.Print("Got SIGTERM...")
		}
	case <-shutdown:
		logger.Printf("Got an error...")
	}

	//server.Shutdown()
}
