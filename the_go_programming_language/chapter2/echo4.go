//flag包好牛逼啊
package main

import(
	"fmt"
	"flag"
	// "os"
	"strings"
)

var n=flag.Bool("n",false,"omit trailing newline")
var sep=flag.String("s"," ","separator")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n{
		fmt.Print("\n")
	}
}