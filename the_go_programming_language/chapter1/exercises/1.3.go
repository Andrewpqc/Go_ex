package main

import(
	"fmt"
	"os"
	"time"
	"strings"
)

func low_effective(ch chan<-string) string{
	start:=time.Now()
	sep:=""
	s:=""
	for _,d:=range os.Args[1:]{
		s+=sep+d
		sep=" "
	}
	fmt.Println(s)
	ch<-fmt.Sprintf("low cost:%f",time.Since(start).Seconds())
	return s
}

func high_effective(ch chan<-string)string{
	start:=time.Now()
	s:=strings.Join(os.Args[1:]," ")
	fmt.Println(s)
	ch<-fmt.Sprintf("high cost:%f",time.Since(start).Seconds())
	return s
}

func main(){
	ch:=make(chan string)
	//下面两个过程几乎同时发生
	go low_effective(ch)
	go high_effective(ch)
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}