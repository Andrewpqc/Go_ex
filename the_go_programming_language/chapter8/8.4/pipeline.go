package main

import(
	"fmt"
)

func main(){
	naturals:=make(chan int)
	squares:=make(chan int)

	go func(){
		for x:=0;;x++{
			naturals<-x
		}
	}()
	
	go func(){
		for x:=0;;x++{
				n:=<-naturals
				squares<-n*n
			}
	}()

	for{
		fmt.Println(<-squares)
	}

}