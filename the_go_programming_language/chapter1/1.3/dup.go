package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)//读取输入，以行或单词为单位分开
	for input.Scan() { //Ctrl+D停止输入
		count[input.Text()]++
	}
	fmt.Println("dfgh")
	for index, size := range count {
		if size > 1 {
			fmt.Println(index, size)
		}
	}
}
