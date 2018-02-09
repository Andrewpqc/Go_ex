//回声服务器

package main

import(
	"net"
	// "os"
	"log"
	"time"
	"fmt"
	"strings"
	"bufio"
)

func main(){
	lister,err:=net.Listen("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	defer lister.Close()
	for{
		conn,err:=lister.Accept()
		if err!=nil{
			log.Fatal(err)
		}
		go connHandler(conn)//使其可以处理多个连接
	}
}


func connHandler(c net.Conn){
	input:=bufio.NewScanner(c)
	for input.Scan(){
		echo(c,input.Text(),1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn,shout string,delay time.Duration){
	fmt.Fprintln(c,"   "+strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"   "+shout)
	time.Sleep(delay)
	fmt.Fprintln(c,"   "+strings.ToLower(shout))
}