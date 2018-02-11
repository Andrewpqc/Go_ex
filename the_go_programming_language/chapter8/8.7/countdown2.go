package main

import (
	"time"
	"os"
	"fmt"
)

func main(){
	abort:=make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte,1))//读取单个字节
		abort<-struct{}{}
	}()
	select{
	case <-abort:
		fmt.Println("Launch aborted!")
		return 
	case <-time.After(10*time.Second):
		//
	}

	fmt.Println("Launched!")
}