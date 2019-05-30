package main

import (
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()

	server := http.Server{
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	server.ListenAndServe()
}
