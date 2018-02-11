package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: nil,
	}

	server.ListenAndServe()
}
