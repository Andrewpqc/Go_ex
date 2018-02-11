package main

import(
	"net"
	"log"
	"io"
	"os"
)


//ch用来同步主goroutine和函数中的goroutine
func main(){
	ch:=make(chan struct{})
	conn,err:=net.Dial("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	go func(){
		io.Copy(os.Stdout,conn)
		log.Println("done")
		ch<-struct{}{}
	}()

	io.Copy(conn,os.Stdin)
	conn.Close()
	<-ch
}