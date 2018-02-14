package main

import (
	"fmt"
	"net/http"
)

func handlerf(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.PostForm)
	fmt.Fprint(w,"hello,world!")
}


func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	http.HandleFunc("/form/",handlerf)
	server.ListenAndServe()
}