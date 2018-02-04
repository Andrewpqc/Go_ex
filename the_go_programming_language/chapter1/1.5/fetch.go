//我操，感觉go语言好牛逼啊
package main

import (
	"io"
	"fmt"
	// "io/ioutil"
	"net/http"
	"os"
)

func main(){
	for _, url :=range os.Args[1:]{
		resp,err:=http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"error:%v",err)
			os.Exit(1)
		}
		b,err:=ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err!=nil{
			fmt.Fprintf(os.Stderr,"error:%v",err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)

	}
}