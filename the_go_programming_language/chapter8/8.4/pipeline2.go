//改进的循环通道

package main

import(
	"fmt"
)


func main(){
	naturals:=make(chan int)
	squares:=make(chan int)

	go func(){
		for i:=0;i<100;i++{
			naturals<-i
		}
		close(naturals)
	}()

	go func(){
		for{
			n,ok:=<-naturals
			if !ok{
				break
			}
			squares<-n*n
		}
		close(squares)
	}()

	for{
		d,ok:=<-squares
		if !ok{
			break
		}
		fmt.Println(d)
	}
}