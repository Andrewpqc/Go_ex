package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	//返回一个通道，定期发送事件，像一个节拍器一样，每一个事件是一个时间戳
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		fmt.Println(<-tick)
	}
	// launch()
}
