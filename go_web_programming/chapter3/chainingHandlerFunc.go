//串联处理器函数

package main

import (
	"net/http"
	"fmt"
	"runtime"
	"reflect"
)

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"hello")
}

func log(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		name:=runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println(name+" called!")
		h(w,r)
	}
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}

	http.HandleFunc("/hello/",log(hello))

	server.ListenAndServe()
}