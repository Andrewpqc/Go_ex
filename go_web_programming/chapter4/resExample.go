package main

import (
	"net/http"
	"encoding/json"
)

type AA struct{
	Name string
	Password string
	aa []string
}

func WriteExample(w http.ResponseWriter,r *http.Request){
	src:=`<html>
	<head><title>Write</title></head>
	<body>
	<h1>hello,world!</h1>
	</body>
	</html>`

	w.Write([]byte(src))
}

func WriteHeaderExample(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(501)
	w.Write([]byte("No such service,try next door\n"))
}


func HeaderExample(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Location","https://www.baidu.com")
	w.WriteHeader(302)
	// w.Write([]byte("No such service,try next door\n"))
}

func WriteJSON(w http.ResponseWriter,r * http.Request){
	a:=AA{Name:"aaaa",Password:"ppppp",aa:[]string{"a","b","c"}}
	json,_:=json.Marshal(a)
	w.Header().Set("Content-Type","application/json")
	w.Write(json)
}


func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	
	http.HandleFunc("/write_example/",WriteExample)
	http.HandleFunc("/write_header/",WriteHeaderExample)
	http.HandleFunc("/header/",HeaderExample)
	http.HandleFunc("/json/",WriteJSON)
	server.ListenAndServe()
}