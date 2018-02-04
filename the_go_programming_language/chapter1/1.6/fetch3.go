//并发获取url
package main

import(
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main (){
	start:=time.Now()
	ch:=make(chan string)
	for _,url:=range os.Args[1:]{
		if !strings.HasPrefix(url,"http://"){
			url="http://"+url
		}
		go fetch(url,ch)
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Println("time:",time.Since(start).Seconds())
}

func fetch(url string,ch chan string){
	start:=time.Now()
	resp,err:=http.Get(url)
	if err!=nil{
		// fmt.Fprintf(os.Stderr,"%v",err)
		ch<-fmt.Sprintf("%v",err)
		return 
	}
	n,err:=io.Copy(ioutil.Discard,resp.Body)
	if err!=nil{
		ch<-fmt.Sprintf("while read %s:%v",url,err)
		return
	}
	deltatime:=time.Since(start).Seconds()
	ch<-fmt.Sprintf("%.2fs %7d %s",deltatime,n,url)
	return 
}