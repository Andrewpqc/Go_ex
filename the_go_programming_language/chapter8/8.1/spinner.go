package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	// time.Sleep(5*time.Second)
	fmt.Println(fib(45))
}

func spinner(delay time.Duration) {
	for {
		for _, d := range `-\|/` {
			fmt.Printf("\r%c", d)
			time.Sleep(delay)
		}
	}
}

func fib(a int) int {
	if a < 2 {
		return a
	} else {
		return fib(a-1) + fib(a-2)
	}
}
