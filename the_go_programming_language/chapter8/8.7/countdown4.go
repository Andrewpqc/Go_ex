package main


import (
	"time"
	"os"
	"fmt"
)

func main(){
	abort:=make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte,1))
		abort<-struct{}{}
	}()

	tick:=time.Tick(1*time.Second)
	for countdown:=10;countdown>0;countdown--{
		fmt.Println(countdown)
		select{
		case <-tick:
			//
		case <-abort:
			fmt.Println("Launch abort!")
			return
		}
	}
	fmt.Println("Launched!")
}