package main

import (
	"net/http"
	"fmt"
)

func hello1(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"/hello")
}




func hello2(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"/ccc/bbb")
}


func hello3(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"/lll/aaa")
}



func hello4(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"/")
}
func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}

	http.HandleFunc("/",hello4)
	http.HandleFunc("/hello/",hello1)
	http.HandleFunc("/lll/aaa",hello3)
	http.HandleFunc("/ccc/bbb",hello2)

	server.ListenAndServe()
}
