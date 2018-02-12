package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"hello")
}

func main(){
	handler:=http.HandlerFunc(hello)
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	http.Handle("/hello/",&handler)
	server.ListenAndServe()
}