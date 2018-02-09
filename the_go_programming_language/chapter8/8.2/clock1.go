//并发时钟服务器

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println(listener.Addr())
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handConn(conn)
	}
}

func handConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(3 * time.Second)
	}
}
