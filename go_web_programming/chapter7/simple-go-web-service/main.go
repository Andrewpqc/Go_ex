package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/post/", handleRequest)
	fmt.Println("Server start at 127.0.0.1,listen on port 8000.")
	server.ListenAndServe()
}
