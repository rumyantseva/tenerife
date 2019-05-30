package main

import (
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	server := http.Server{
		Addr: net.JoinHostPort("", port),
	}

	server.ListenAndServe()
}
