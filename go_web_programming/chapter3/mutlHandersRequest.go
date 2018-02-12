//多处理器

package main

import(
	"net/http"
	"fmt"
)

type HelloHandler struct{}

func (hh *HelloHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Hello")
}

type WorldHandler struct{}

func (wh *WorldHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"World")
}


func main(){
	hello:=HelloHandler{}
	world:=WorldHandler{}

	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	http.Handle("/hello/",&hello)
	http.Handle("/world/",&world)

	server.ListenAndServe()
}