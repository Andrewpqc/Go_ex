//输出命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s1, s2 string
	s1 = ""
	for i := 1; i < len(os.Args); i++ {
		s1 += os.Args[i]
		fmt.Println(os.Args[i])
	}
	s2 = s1 + "this is what appended to s1"
	fmt.Println(s2)
}
