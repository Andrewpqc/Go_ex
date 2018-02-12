package main
import (
	"net/http"
	"fmt"
)

type MyHandler struct{}

func (h *MyHandler)ServeHTTP(http.ResponseWriter,*http.Request){
	fmt.Fprintf(w,"Hello,world,%d",1)
}
	
func main(){
	var handler MyHandler
	server:=http.Server{
		Addr:"127.0.0.1:8000",
		Handler:&handler
	}

	server.ListenAndServe()
	
}