package main

import (
	"fmt"
	"net/http"
)

func getHeader(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,r.Header)
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	http.HandleFunc("/header",getHeader)
	server.ListenAndServe()
}