//串联处理器
package  main

import (
	"fmt"
	"net/http"
	// "runtime"
	// "reflect"
)

type HelloHandler struct{}

func (h *HelloHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"hello")
}

func log(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		fmt.Printf("%T called!",h)
		h.ServeHTTP(w,r)
	})
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}

	hello:= HelloHandler{}
	http.Handle("/hello/",log(&hello))
	server.ListenAndServe()
}