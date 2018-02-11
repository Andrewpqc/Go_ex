package main

import(
	"net/http"
	"fmt"
)

func handler(writer http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method,req.URL.Path,req.Host,req.URL)
	fmt.Fprintf(writer,"hello from %s",req.URL.Path[1:])
}

func main(){
	http.HandleFunc("/",handler)
	fmt.Println("Server start at 127.0.0.1:8080!")
	http.ListenAndServe(":8080",nil)
	
}