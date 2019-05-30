package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/rumyantseva/tenerife/internal/application"
)

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler)

	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	server.ListenAndServe()
}
