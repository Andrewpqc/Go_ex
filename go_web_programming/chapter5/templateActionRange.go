package main

import (
	"net/http"
	"html/template"
)


func rangeAction(w http.ResponseWriter,r*http.Request){
	t:=template.Must(template.ParseFiles("range.html"))
	t.Execute(w,[]string{"新年快乐","万事如意","大吉大利"})
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	http.HandleFunc("/range/",rangeAction)
	server.ListenAndServe()
}