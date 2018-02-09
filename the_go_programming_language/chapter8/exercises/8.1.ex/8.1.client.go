package main

import (
	"time"
	"io"
	"log"
	"flag"
	"net"
	"os"
	"fmt"
)

var NewYork = flag.String("NewYork", "", "get the time of NewYork")
var London = flag.String("London", "", "get the time of London")
var Tokyo = flag.String("Tokyo", "", "get the time of Tokyo")

func main() {
	flag.Parse()
	if *NewYork != "" {
		fmt.Println(*NewYork)
		go client(*NewYork)
	}
	if *London != "" {
		go client(*London)
	}
	if *Tokyo != "" {
		go client(*Tokyo)
	}	
	time.Sleep(1*time.Minute)
}


func client(addr string){
	conn,err:=net.Dial("tcp",addr)
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout,conn)
}

func mustCopy(dst io.Writer,src io.Reader){
	if _,err:=io.Copy(dst,src);err!=nil{
		log.Fatal(err)
	}
}