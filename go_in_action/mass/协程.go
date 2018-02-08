package main 

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start")

	//打印可以使用的物理处理器的数量
	fmt.Println(runtime.NumCPU())

	go func(){
		defer wg.Done()
		for index:=0;index<100;index++{
			fmt.Println("A",index)
		}
	}()


	go func(){
		defer wg.Done()
		for index:=0;index<100;index++{
			fmt.Println("B",index)
		}
	}()

	fmt.Println("waiting goroutine finish")
	wg.Wait()
	fmt.Println("gorutine is over")

}