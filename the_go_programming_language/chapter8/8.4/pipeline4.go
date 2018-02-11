package main

import (
	"fmt"
)

func counter(out chan <- int){
	for i:=0;i<100;i++{
		out<-i
	}
	close(out)
}

func squarer(out chan <- int,in <-chan int){
	for i:=range in{
		out<-i*i
	}
	close(out)
}


func printer(out <-chan int){
	for i:=range out{
		fmt.Println(i)
	}
}


func main(){
	naturals:=make(chan int)
	squarers:=make(chan int)
	go counter(naturals)
	go squarer(squarers,naturals)
	printer(squarers)
}