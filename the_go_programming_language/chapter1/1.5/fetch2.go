package main

import (
	"io"
	"strings"
	"net/http"
	"fmt"
	"os"
)

func main(){
	for _,url := range os.Args[1:]{
		if !strings.HasPrefix(url,"http://"){
			url="http://"+url
		}
		resp,err:=http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"%v",nil)
			os.Exit(1)
		}

		io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
	}
}