//处理器函数
package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"hello")
}


func world(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"world")
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}

	http.HandleFunc("/hello/",hello)
	http.HandleFunc("/world/",world)

	server.ListenAndServe()
}