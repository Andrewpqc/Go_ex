//另一种形式的循环
package main	

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	i:=1
	for i<5{
		fmt.Println(i)
		i++
	}
}
