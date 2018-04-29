package main

import (
	"time"
	"fmt"
)

func PrintLetters(){
	for i:='A';i< 'A'+10;i++{
		fmt.Printf("%c ",i)
	}
}
+

func PrintNumbers(){
	for i:=1;i<10;i++{
		fmt.Printf("%d ",i)
	}
}


func Print1(){
	PrintLetters()
	PrintNumbers()
}


func Print2(){
	go PrintLetters()
	go PrintNumbers()
	time.Sleep(1*time.Second)
}


func main(){
	Print1()
	Print2()
}
