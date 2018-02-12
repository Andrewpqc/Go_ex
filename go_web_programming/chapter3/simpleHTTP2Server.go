package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
	
)



type MyHandler struct{}


func (h *MyHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"hello")
}
func main(){
	handler:=MyHandler{}
	server:=http.Server{
		Addr:"127.0.0.1:8000",
		Handler:&handler,
	}

	http2.ConfigureServer(&server,&http2.Server{})
	server.ListenAndServeTLS("cert.pem","key.pem")
}